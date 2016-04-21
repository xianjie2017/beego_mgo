package controllers

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Id    bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Phone string        `bson:"phone"`
}

type ResultMsg struct {
	Message string
}

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Get() {
	//this.Ctx.WriteString("hello world")
	name := this.Ctx.Input.Param(":name")
	//var persons []Person
	//id := this.Ctx.Input.Param(":id")

	var person Person
	query := func(c *mgo.Collection) error {
		//return c.Find(nil).All(&persons)
		return c.Find(bson.M{"name": name}).One(&person)
		//return c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&person)
	}
	err := Execute("col", query)
	if err != nil {
		//panic(err)
		this.Data["json"] = ResultMsg{Message: err.Error()}
	} else {
		//this.Data["json"] = persons
		this.Data["json"] = person
	}
	this.ServeJSON()
}

func (this *HelloController) Post() {
	name := this.Ctx.Input.Param(":name")
	phone := this.Ctx.Input.Param(":phone")
	insert := func(c *mgo.Collection) error {
		return c.Insert(&Person{Id: bson.NewObjectId(), Name: name, Phone: phone})
	}
	err := Execute("col", insert)
	if err != nil {
		//panic(err)
		this.Data["json"] = ResultMsg{Message: err.Error()}
	} else {
		this.Data["json"] = ResultMsg{Message: "success"}
	}
	this.ServeJSON()
}

func (this *HelloController) Delete() {
	name := this.Ctx.Input.Param(":name")
	remove := func(c *mgo.Collection) error {
		return c.Remove(bson.M{"name": name})
	}
	err := Execute("col", remove)
	if err != nil {
		//panic(err)
		this.Data["json"] = ResultMsg{Message: err.Error()}
	} else {
		this.Data["json"] = ResultMsg{Message: "success"}
	}

	this.ServeJSON()
}

func (this *HelloController) Put() {
	name := this.Ctx.Input.Param(":name")
	phone := this.Ctx.Input.Param(":phone")
	update := func(c *mgo.Collection) error {
		return c.Update(bson.M{"name": name}, bson.M{"$set": bson.M{"phone": phone}})
	}
	err := Execute("col", update)
	if err != nil {
		//panic(err)
		this.Data["json"] = ResultMsg{Message: err.Error()}
	} else {
		this.Data["json"] = ResultMsg{Message: "success"}
	}

	this.ServeJSON()
}
