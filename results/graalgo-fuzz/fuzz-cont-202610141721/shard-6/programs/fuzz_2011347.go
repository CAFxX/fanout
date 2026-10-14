package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
func genFn0(p0 bool) string {
    p0 = ((42 == 4294967295) && p0)
    fmt.Println(p0)
    fmt.Println(p0)
    v5 := GenS1{F0: 2147483647}
    if ((false || (p0 && false)) || (1e10 >= 0.0)) {
        fmt.Println(p0)
        v6 := GenS2{F0: 127}
        v7 := []int{-2147483648, 32767}
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v5
    return ""
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 42
    var v4 uint64 = 576
    for fi8 := 0; fi8 < 8; fi8++ {
        _ = len(v1)
        v9 := GenS3{F0: ""}
        _ = fi8
        _ = v9
    }
    var v10 uint64 = v4
    var v11 string = (("foo bar" + "foo bar") + "\n\t")
    if (!(v4 < (v4 - v4))) {
        var v12 uint64 = (v4 * v10)
        var v13 uint8 = 1
        _ = v12
        _ = v13
    }
    v11 = v11
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint64 = 0
        var v16 bool = true
        for fi17 := 0; fi17 < 2; fi17++ {
            v11 = "hello"
            _ = fi17
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v18 uint64 = (v10 / ((864) | 1))
    v4 -= v18
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v10, v11, v4)
    if false {
        var v19 int8 = 76
        _ = v19
    }
    fmt.Println(v18, v18)
    v11 = v11
    fmt.Println(v3, v11, v4)
    for fi20 := 0; fi20 < 5; fi20++ {
        v4 = (v10 & (v18 * (v4 * 411)))
        _ = fi20
    }
    _ = len(v2)
    if ((-1 < -9223372036854775808) || true) {
        var v21 uint16 = (v3 >> ((2) & 63))
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
