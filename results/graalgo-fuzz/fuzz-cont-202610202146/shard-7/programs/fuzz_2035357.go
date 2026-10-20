package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int) string {
    var v5 int32 = 2147483647
    if true {
        var v6 bool = false
        _ = v6
    }
    _ = p0
    _ = v5
    return ("hello" + "a")
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int32 = 0
    var v3 bool = false
    var v4 uint16 = 255
    var v7 int16 = 42
    v7 = v7
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint = 202
        _ = fi8
        _ = v9
    }
    v10 := []int{v1, v1, (-(v1))}
    var v11 uint16 = v4
    for fi12 := 0; fi12 < 16; fi12++ {
        if len(v10) > 0 {
            v10[0] = int(v11)
        }
        _ = fi12
    }
    v10 = append(v10, -244)
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v7, v4, v1)
        fmt.Println(v7)
        _ = fi13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint32 = 334
        var v16 uint16 = (^(v11))
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        fmt.Println(v2, v11)
        if v3 {
            v18 := GenS1{F0: (false && (-2.25 == (3.14159 * -2.25)))}
            var v19 uint = 1
            _ = v18
            _ = v19
        } else {
            fmt.Println(v1, v3, v7)
            fmt.Println(v7, v7, v1)
        }
        v1 = (v1 / ((v1) | 1))
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
