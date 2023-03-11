# monkey-go
from the book "Writing An Interpreter in Go"

```
>> let newAdder = fn(x) { fn(y) { x + y}; };            
>> let addTwo = newAdder(2);
>> addTwo(3);
5
```
