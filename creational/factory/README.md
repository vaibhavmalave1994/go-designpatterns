# Factory Design Pattern

Factory pattern is second most used design pattern.
Its purpose is to abstract user from the knowleged of struct they need to achieve for specific purpose.
The user only needs an interface that provides him this values.
By delegating descision to factory, a factory can provide a interface that fits user requirment.
It also eases the process of upgrading or downgrading of the implementation of the underlaying types if needed.

While using factory pattern we delegate the creation of families of object to another package to abstract us from knowleged of creating posible objects that we could use.


# when to use factory design pattern?
When a method return one of the several possible classes that share common parent class.
The class chosen at then run time.
1. When you don't know in future what class object you need
2. When the potential classed are in the same subclass hierarchy
3. To centralize class selection code
4. When you don't want user to know each subclass
5. To encapsulate object selection

