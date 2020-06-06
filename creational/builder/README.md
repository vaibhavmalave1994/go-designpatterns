# Builder Design pattern

Builder design patterns helps us construct complex objects without directly instantiating their structs, or writing complex logic they required.

Imagine an object that could have dozens of fields that themselves are complex structs. Now imagine you have many objects with such charecteristics, and you could have more. We don't need to write logic to create such objects in package that just need to be use the object.

# objective
1. Abstract complex creation so that object creation is separated from object user
2. Create an object step by step by filling its fields and creating embeded objects 
3. Reuse the object creation algoritham between many objects

# Example – vehicle manufacturing
The Builder design pattern has been commonly described as the relationship between a director, a few Builders, and the product they build. Continuing with our example of the car, we'll create a vehicle Builder. The process (widely described as the algorithm) of creating a vehicle (the product) is more or less the same for every kind of vehicle–choose vehicle type, assemble the structure, place the wheels, and place the seats. If you think about it, you could build a car and a motorbike (two Builders) with this description, so we are reusing the description to create cars in manufacturing. The director is represented by the ManufacturingDirector type in our example.

# Requirements and acceptance criteria
As far as we have described, we must dispose of a Builder of type Car and Motorbike and a unique director called ManufacturingDirector to take builders and construct products. So the requirements for a Vehicle builder example would be the following:
1. I must have a manufacturing type that constructs everything that a vehicle needs 
2. When using a car builder, the VehicleProduct with four wheels, five seats, and a structure defined as Car must be returned
3. When using a motorbike builder, the VehicleProduct with two wheels, two seats, and a structure defined as Motorbike must be returned
4. A VehicleProduct built by any BuildProcess builder must be open to modifications