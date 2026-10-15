package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int64 = -241
    var v2 string = ("\n\t" + "foo bar")
    var v3 int = -998
    if false {
        v2 = (v2 + v2)
        fmt.Println(v1)
        if (v2 == "foo bar") {
            if false {
                fmt.Println(v2, v3)
            } else {
                v4 := GenS1{F0: 973}
                v1 += (v1 | 976)
                var v5 uint = 42
                _ = v4
                _ = v5
            }
            for fi6 := 0; fi6 < 11; fi6++ {
                v7 := GenS1{F0: 1}
                fmt.Println(v3, v1, v3)
                _ = fi6
                _ = v7
            }
        }
    }
    v2 = ("" + v2)
    v3 += 4294967295
    v2 = v2
    var v8 int8 = -36
    var v9 uint32 = 991
    v2 = v2
    v2 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
