# MaxStack-Golang
Implementation of the MaxStack Data Structure in Go

This repository contains the design of a max stack data structure that supports the stack operations and supports finding the stack's maximum element.

The code implements the MaxStack interface with the following functions:

1. Constructor() Initializes the stack object.
2. void push(int x) Pushes element x onto the stack.
3. int pop() Removes the element on top of the stack and returns it.
4. int top() Gets the element on the top of the stack without removing it.
5. int peekMax() Retrieves the maximum element in the stack without removing it.
6. int popMax() Retrieves the maximum element in the stack and removes it. If there is more than one maximum element, only remove the top-most one.
