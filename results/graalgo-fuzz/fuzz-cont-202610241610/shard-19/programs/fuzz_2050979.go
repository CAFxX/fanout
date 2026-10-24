package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint }
func genFn0() int8 {
    var v4 float64 = (1e10 + (1.5 / 1.5))
    var v5 uint8 = 226
    var v6 float64 = 1.5
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v4, v6, v4)
        v8 := GenS1{F0: 84}
        _ = fi7
        _ = v8
    }
    v4 -= (float64(0.1) * v6)
    _ = v4
    _ = v5
    _ = v6
    return -128
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 uint16 = 0
    var v9 uint16 = (v3 % ((801) | 1))
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(fi10, v3, v9)
        v11 := GenS2{F0: true, F1: 42}
        if (!false) {
            var v12 int16 = -1
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    v3 = v3
    v1["a"] = -485
    fmt.Println(v9, v9)
    var v13 int32 = 157
    v13 -= (v13 >> ((42) & 63))
    switch 32767 {
    case -2147483648:
        var v14 int16 = 894
        _ = v14
    }
    switch ("hello" + ("hello" + "a")) {
    case "a":
        fmt.Println(v13, v3, v13)
    case "\n\t":
        var v15 string = (("foo bar" + ("a" + "a")) + "hello")
        _ = v15
    case "foo bar":
        var v16 int16 = 42
        _ = v16
    default:
        v17 := GenS1{F0: 91}
        _ = v17
    }
    _ = len(v2)
    if false {
        if (0.1 < -2.25) {
            v18 := []int{}
            var v19 uint64 = 12
            var v20 float64 = 0.0
            _ = v18
            _ = v19
            _ = v20
        } else {
            v9 = 1
            var v21 int8 = 0
            _ = v21
        }
        for fi22 := 0; fi22 < 16; fi22++ {
            v13 -= v13
            _ = fi22
        }
        var v23 string = (("" + "hello") + ("hello" + ("foo bar" + ("a" + "\n\t"))))
        _ = v23
    } else {
        v24 := GenS1{F0: 197}
        _ = v24
    }
    v2[65535] = (((65535 != 65535) || (1 > -337)) && true)
    var v25 int = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v9
}
