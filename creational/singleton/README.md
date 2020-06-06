# Singletone Design pattern

Singletone design patterns will provide you with single instace of object, and gaurantees that there are no duplicate.

At the first call to use instance, it is created and then resused between all the parts in the application that need to use particular behaviour.


You will use the singletone pattern in many situations, like:
1. We need to use single connection to database to run all quries
2. we need to open SSH connection and use same connection to do all tasks
3. We need to limit access to varibale or space. Then we can use singletone as a door to this variable/ space
4. We need to limit the number of calls to places, then we create a singletone instance to make the call in accepted window.

# Objective
We consider using singletone pattern when the following rule applies
1. We need a single, shared value of particular type
2. We need to restrict object creation of some type to single unit along the entire program.

# Example - a unique counter
As an example of an object of which we must ensure that there is only one instance, we will write a counter that holds the number of times it has been called during program execution. It shouldn't matter how many instances we have of the counter, all of them must count the same value and it must be consistent between the instances.


# Requirements and acceptance criteria
There are some requirements and acceptance criteria to write the described single counter. They are as follows:
1. When no counter has been created before, a new one is created with the value 0
2. If a counter has already been created, return this instance that holds the actual count
3. If we call the method AddOne, the count must be incremented by 1.


# Notes 
1. Singleton package will give us power to have unique instace of struct in application, and no other package can make clone of this struct
2. Singleton pattern also hides complexity of creating object, incase it requires some computation, and pitfall of creating it every time you need it if all of them are similar. All this code writting, checking if the variable already exust, and storage, are encapsulated in singleton and we don't need to repeat it everywhere if you use a global variable.
