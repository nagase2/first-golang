

var objectArray: any[] = [
  {x: "hello", y: 1},
  {x: "world", y: 5},
  {x: "foo", y: 3},
  {x: "bar", y: 2},
];

var sortedArray: any[] = objectArray.sort((n1,n2) => {
    if (n1.y > n2.y) {
        return -1;
    }
    if (n1.y < n2.y) {
        return 1;
    }
    return 0;
});

console.log(sortedArray)
