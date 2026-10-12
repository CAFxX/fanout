package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }

func main() {
    var v1 int16 = 32767
    var v2 int64 = -1
    v3 := []int{-2147483648, 898, 42, 128}
    var v4 float32 = -2.25
    v5 := make(map[int]bool)
    v6 := make(map[int]uint32)
    switch (((("hello" + "foo bar") + ("a" + "hello")) + (("" + "foo bar") + ("a" + ""))) + "foo bar") {
    case "a":
        fmt.Println(v4, v1)
    }
    if ((-128 >= -1) || ((1.5 == float32(1.5)) || ((false || (true || false)) || false))) {
        if len(v3) > 0 {
            v3[0] = 1
        }
    }
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v1 -= (v1 % (((v1 ^ (-(v1)))) | 1))
            var v9 int = 9223372036854775807
            _ = fi8
            _ = v9
        }
        var v10 int8 = 1
        _ = fi7
        _ = v10
    }
    v2 -= 9223372036854775807
    var v11 int16 = (v1 ^ v1)
    for fi12 := 0; fi12 < 3; fi12++ {
        v4 = 3.14159
        v13 := GenS1{F0: 4294967295}
        var v14 int8 = -94
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 1}
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 bool = (!false)
        var v18 int64 = v2
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v1, v2)
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS1{F0: 4294967295}
        _ = fi19
        _ = v20
    }
    var v21 uint32 = 78
    switch 577 {
    case 42:
        v4 = 3.14159
    case 32768:
        fmt.Println(v21, v21, v2)
    case 0:
        if (int64(v21) > v2) {
            var v22 uint64 = 2
            v23 := GenS2{F0: 659}
            _ = v22
            _ = v23
        } else {
            v1 += (v11 >> ((4294967295) & 63))
        }
    }
    v2 = (v2 | v2)
    delete(v5, 498)
    var v24 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
