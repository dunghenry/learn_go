function displayNumber() {
    number = 0;
    return function () {
        number++;
        return number;
    };
}

num = displayNumber();
console.log(num())
console.log(num())
console.log(num())