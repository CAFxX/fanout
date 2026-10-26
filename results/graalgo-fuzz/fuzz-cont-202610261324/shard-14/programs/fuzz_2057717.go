package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int8) bool {
    p0 = (p0 % ((p0) | 1))
    if ((("hello" + "foo bar") + "") > (("" + "a") + "")) {
        var v5 uint = 781
        v6 := []int{22, 127}
        switch 32768 {
        case -226:
            v7 := GenS1{F0: "", F1: 2, F2: 1}
            _ = v7
        case -9223372036854775808:
            v5 -= (v5 ^ v5)
        case 2147483647:
            _ = len(v6)
        default:
            v5 += 121
        }
        _ = v5
        _ = v6
    } else {
        if (255 >= 65535) {
            var v8 uint16 = 65535
            v8 = v8
            _ = v8
        }
        switch "a" {
        case "a":
            v9 := GenS1{F0: ("a" + ""), F1: 2147483647, F2: -670}
            _ = v9
        default:
            p0 = (p0 / ((p0) | 1))
        }
    }
    v10 := []int{-2147483648, 128, -9223372036854775808, 127}
    var v11 int = -67
    _ = p0
    _ = v10
    _ = v11
    return true
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ("hello" + "foo bar")
    var v3 uint64 = 4294967295
    var v4 int32 = 0
    v3 += v3
    for fi12 := 0; fi12 < 2; fi12++ {
        if true {
            v13 := GenS1{F0: v2, F1: 2147483648, F2: 42}
            _ = v13
        }
        _ = fi12
    }
    v2 = v2
    v3 -= (v3 << ((65535) & 63))
    v4 -= v4
    v14 := GenS2{F0: 100}
    for fi15 := 0; fi15 < 13; fi15++ {
        fmt.Println(v4)
        v16 := GenS2{F0: 127}
        _ = fi15
        _ = v16
    }
    v4 -= (v4 - (v4 * v4))
    _ = len(v1)
    v17 := []int{-644, -29}
    fmt.Println(v3, v3, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
