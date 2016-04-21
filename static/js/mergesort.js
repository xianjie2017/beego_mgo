/**
 * Created by simon on 2016/4/7.
 */
function ArrayList(){
    var array = [];
    this.insert = function(item){
        array.push(item);
    };
    this.toString=function() {
        return array.join();
    };
    this.mergeSort = function(){
        array=mergeSortRec(array);
    };
}

var mergeSortRec=function(array){
    var length=array.length;
    if(length===1){
        return array;
    }
    var mid = Math.floor(length/2),//floor() 方法执行的是向下取整计算，它返回的是小于或等于函数参数，并且与之最接近的整数。
        left = array.slice(0,mid),
        right=array.slice(mid,length);
    return merge(mergeSortRec(left),mergeSortRec(right));
};

var merge = function(left, right){
    var result = [],il=0,ir=0;
    while(il<left.length&&ir<right.length){
        if(left[il]<right[ir]){
            result.push(left[il++]);
        }else{
            result.push(right[ir++]);
        }
    }
    while(il<left.length){
        result.push(left[il++]);
    }
    while(ir<right.length){
        result.push(right[ir++]);
    }
    return result;
};

//测试
var array = new ArrayList();
array.insert(10);
array.insert(3);
array.insert(8);
array.insert(2);
array.insert(7);
array.insert(5);
array.insert(6);
array.insert(9);
array.insert(0);
array.insert(1);
array.mergeSort();
console.log(array.toString());