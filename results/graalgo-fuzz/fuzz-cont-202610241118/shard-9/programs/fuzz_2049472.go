package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint32, p1 int16) uint32 {
    switch 127 {
    case 400:
        p0 += (p0 / ((p0) | 1))
    case -1:
        fmt.Println(p1, p1)
    case 2147483647:
        fmt.Println(p1)
    default:
        fmt.Println(p0, p0)
    }
    fmt.Println(p1, p1)
    if (0 > -1) {
        fmt.Println(p0, p0)
    } else {
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 uint16) int8 {
    for fi6 := 0; fi6 < 15; fi6++ {
        p0 = p0
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        p0 = (^(p0))
        p0 = 2
        v9 := GenS1{F0: ("" + "foo bar")}
        _ = fi8
        _ = v9
    }
    _ = p0
    return -128
}

func main() {
    var v1 string = "a"
    var v2 string = (v1 + "hello")
    var v3 int32 = -1
    v4 := make(map[uint32]string)
    var v5 bool = false
    var v10 uint64 = 1
    if (32767 >= 32767) {
        v3 = (v3 % ((42) | 1))
    }
    v11 := GenS1{F0: v2}
    v12 := GenS1{F0: "foo bar"}
    v10 = 0
    var v13 float64 = float64(0.0)
    v1 = v1
    fmt.Println(v5, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
