package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() bool {
    if false {
        var v4 int32 = 0
        _ = v4
    } else {
        var v5 int8 = 100
        fmt.Println(v5)
        var v6 int8 = v5
        _ = v5
        _ = v6
    }
    var v7 int8 = 0
    _ = v7
    return ((!(0 < 255)) || false)
}
func genFn1(p0 int8) int64 {
    p0 = p0
    v8 := []int{}
    var v9 uint32 = 10
    _ = p0
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{65535, 1, 42}
    var v2 int64 = 9223372036854775807
    var v3 uint16 = 1
    v2 = (v2 / ((v2) | 1))
    var v10 bool = false
    var v11 uint = 255
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v11 -= 2
    fmt.Println(v10, v3)
    switch (("" + "a") + "foo bar") {
    case "a":
        for fi12 := 0; fi12 < 11; fi12++ {
            v2 = v2
            _ = len(v1)
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 bool = v10
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    default:
        v1 = append(v1, 2147483647)
    }
    v15 := GenS1{F0: (^(v11))}
    var v16 int8 = 23
    _ = len(v1)
    v11 = v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
