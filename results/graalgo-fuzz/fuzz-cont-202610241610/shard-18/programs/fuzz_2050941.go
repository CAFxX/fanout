package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint64 }
func genFn0(p0 uint64, p1 uint64, p2 int8) int32 {
    var v4 int = 856
    switch 0 {
    case -640:
        fmt.Println(p1)
    }
    v5 := []int{v4}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -525
}

func main() {
    var v1 int64 = 42
    var v2 uint = 18446744073709551615
    var v3 int32 = 2147483647
    if false {
        fmt.Println(v1, v1)
        var v6 int64 = (v1 / ((v1) | 1))
        v7 := []int{2147483647, -324, 0, 32767}
        _ = v6
        _ = v7
    } else {
        v2 -= 390
    }
    if (!(false || ((0.0 + (-2.25 * (3.14159 + 1.5))) > -2.25))) {
        v8 := GenS1{F0: "\n\t", F1: 42, F2: 2}
        var v9 uint8 = uint8(v2)
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        switch (fi10 * fi10) {
        case 42:
            var v11 uint = v2
            _ = v11
        case -421:
            fmt.Println(v1)
        }
        _ = fi10
    }
    var v12 int8 = 0
    fmt.Println(v1, v2)
    v1 -= v1
    if false {
        fmt.Println(v2, v12, v1)
    }
    fmt.Println(v3, v12)
    v13 := GenS1{F0: ("a" + "hello"), F1: 2147483647, F2: 4294967295}
    v12 = v12
    v14 := GenS1{F0: "\n\t", F1: 2147483647, F2: 42}
    v15 := []int{}
    var v16 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
