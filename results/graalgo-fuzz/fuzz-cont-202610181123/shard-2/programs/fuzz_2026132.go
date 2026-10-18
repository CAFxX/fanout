package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint16) int64 {
    p0 = p0
    p0 += p0
    var v5 uint32 = uint32(p0)
    var v6 bool = true
    v7 := []int{}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 483
}

func main() {
    var v1 int32 = 0
    var v2 uint8 = 42
    var v3 string = ""
    var v4 uint16 = uint16(v2)
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int64 = 42
        _ = fi8
        _ = v9
    }
    fmt.Println(v4, v2, v1)
    v3 = "a"
    fmt.Println(v2, v1)
    for fi10 := 0; fi10 < 8; fi10++ {
        v4 = 279
        v4 -= (^(v4))
        v1 += -616
        _ = fi10
    }
    v4 -= 0
    var v11 uint32 = 2147483647
    v12 := []int{}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v13 := GenS1{F0: (v3 + v3)}
    for fi14 := 0; fi14 < 12; fi14++ {
        if (float64((3.14159 * (0.1 * 1e10))) == 1e10) {
            var v15 int8 = 127
            var v16 string = ("a" + "foo bar")
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    if true {
        v11 -= (^(v11))
    }
    v1 += -2147483648
    if true {
        var v17 bool = (!false)
        _ = v17
    } else {
        fmt.Println(v1)
        fmt.Println(v4)
        var v18 int16 = -273
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
