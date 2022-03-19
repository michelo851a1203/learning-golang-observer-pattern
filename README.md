## golang observer pattern

#### [interface] observer (method following)
- getID() string : get observer ID
- Notify(string) : notify (such as send email to user ...)

#### [interface] subject
- notifyAll() : call all observer `Notify` method
- addObserver(observer) > (register) : add Observer to observer list
- removeObserver(observer) > (deRegister) : remove observer on observer list

#### [concrete observer] customer
> provided by interface
- GetID()
- Notify()

##### [concrete subject] product
+ observerList
- New > build instance to input Notify message here
- addObserver(observer) > append observer to observerList
- removeObserver(observer) > remove observer from observerList
- notifyAll() > for loop to implement observerList notify method

#### usage situation

``` usage situation
if product is on sale then email all customer
