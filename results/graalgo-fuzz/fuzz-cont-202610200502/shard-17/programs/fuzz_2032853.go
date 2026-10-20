package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int8 }

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 4294967295
    var v3 int = 128
    var v4 uint = 65535
    var v5 float64 = -2.25
    if true {
        var v6 int32 = 2147483647
        if false {
            var v7 uint16 = 65535
            var v8 float64 = (0.0 - 0.1)
            _ = v7
            _ = v8
        }
        v1 = (v1 + ((v1 + (("foo bar" + v1) + "hello")) + v1))
        _ = v6
    }
    var v9 string = (v1 + "a")
    switch (v3 & v3) {
    case -1:
        v10 := []int{(v3 | 906)}
        _ = v10
    }
    v11 := GenS1{F0: v1}
    v12 := GenS2{F0: (v2 ^ 1), F1: 0}
    v4 = 179
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
