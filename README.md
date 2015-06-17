# go-testhelpers
Simple methods, shortcuts, and patterns to fill in the gaps when working with go test.  The `testhelpers` library provides helper methods and patterns to easily provide clear and useful error messages without changing your testing methodology or practices.

This package is not a testing framework or an assertion library. 

# Overview
The draw for me to the testing frameworks/libraries was the nice output.

The library ships with a small number of helper functions that will output nicely:

    t.Error(NotEqualMsg(1234, 1234.1))

    example_test.go:1: 
           expected: 1234 
             actual: 1234.1

Because they just output strings, you can chain them or use other values:

    t.Error("Error Context",
      NotEqualMsg(1234, 1234.1))

    example_test.go:1: Error Context
           expected: 1234 
             actual: 1234.1

There are also raw strings you can use to compose your own messages:

    t.Error("Error Context",
      ES_COUNT, 5,
      NotEqualMsg(1234, 1234.1))

    example_test.go:1: Error Context
              count: 5
           expected: 1234 
             actual: 1234.1
      
This principle keeps your test code clean, consise, and is easily extentable.

## Methods
Because the raw consts are so easily composable, only a handful of helper methods ship out of the box.

These are __not__ assertion methods, only message formatting methods.  A typical usage pattern is:

    expected := "Ali Baba"
    actual := theif.Name
    if expected != actual  {
      t.Error(NotEqualMsg(expected, actual))
    }

This gives you the most usability without changing or adding any syntactic 'sugar' to `go test`.  If your devs know go, then they don't have another framework to spin up or learn to test, allowing them to write useful tests quickly.

### Not Equal: `NotEqualMsg(e, a)`

    t.Error("Error Context",
      NotEqualMsg(1234, 1234.1))

    example_test.go:1: Error Context
           expected: 1234 
             actual: 1234.1

### Type Not Equal: `TypeNotEqualMsg(e, a)`
This is a special wrapper that will display the output of `fmt.Sprintf("%T")` of both the expected and actual inputs.

    strongType, ok := interfaceType.(MyType)
    if !ok {
      t.Error("Wrong Type",
        TypeNotEqualMsg(MyType{}, interfaceType))
    }

    example_test.go:1: Wrong Type
           expected: MyType 
             actual: WhatEverTypeItActuallyWas

### SQL + Args: `SqlArgsMsg(sql, args)`
Display SQL statement and any args

    t.Error("DB Error",
      SqlArgsMsg(sql, args))

    example_test.go:1: DB Error
         sql: SELECT * FROM TABLE
        args: [1 2 1.1 true false]

### Unexpected Error: `UnexpectedErrMsg(err)`
    
    result, err := LoadCust()
    if err !+ nil {
      t.Error("LoadCust()",
        UnexpectedErrMsg(err))
    }

    example_test.go:1: LoadCust() Unexpected Error
                err: Table Customers does not exist

It should be noted that if you do not want the `Unexpected Error` moniker then just use the `ES_ERR` const:
    
    result, err := LoadCust()
    if err !+ nil {
      t.Error("LoadCust()",
        ES_ERR, err)
    }

    example_test.go:1: LoadCust() 
                err: Table Customers does not exist

## Error String Consts (`ES_`)
The error string consts are really where the power and simplicty of the library shine.  Each one can be cherry picked and composed with any other.

Here is an example showing all of them:

    t.Error("OH NOEZ!",
      ES_EXPECTED, "one",
      ES_GOT, "two",
      ES_ARGS, "three",
      ES_SQL, "four",
      ES_ERR, "five",
      ES_VALUE, "six",
      ES_COUNT, "seven")

    example_test.go:1: OH NOEZ!
            example: one
             actual: two
               args: three
                sql: four
                err: five
              value: six
              count: seven

They are also composable with the other helper methods, say you are need a count in your error message:

    t.Error(NotEqualMsg(1234, 1234.1),
      ES_COUNT, 5)

    example_test.go:1: 
           expected: 1234 
             actual: 1234.1
              count: 5
