const cal = (array, target) => {
    let count = 0
    for (let i = 0; i < array.length; i++) {
  
        for (let j = i+1; j < array.length; j++) {
            // console.log(array[i],array[j],array[i]*array[j])
            if (array[i]*array[j] == target ){
                console.log(array[i], array[j])
                count++
            }
        }
    }
    return count
 }
 

//  console.log(cal([0, 9, 11, 100], 0))
 console.log(cal([1, 2, 3, 4,5,6], 12))