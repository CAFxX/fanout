package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 int32 }
func genFn0(p0 bool, p1 string, p2 uint16) uint16 {
    fmt.Println(p2, p1)
    v4 := []int{32767, 648, 9223372036854775807, 128}
    p2 -= (p2 - (p2 + 65535))
    switch 889 {
    case -2147483648:
        v5 := GenS1{F0: 412, F1: p2}
        _ = v5
    case 32768:
        fmt.Println(p1, p1)
    case 235:
        var v6 int64 = -9223372036854775808
        _ = v6
    }
    if (false && p0) {
        p2 = 32767
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v7 := GenS1{F0: 65535, F1: (p2 % ((65535) | 1))}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 255
}
func genFn1(p0 int64, p1 uint) int32 {
    var v8 bool = (false && true)
    p1 -= p1
    _ = p0
    _ = p1
    _ = v8
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 1
    v3 := []uint32{}
    v9 := GenS2{F0: 453, F1: -9223372036854775808, F2: -1}
    var v10 uint = 0
    var v11 int8 = -1
    switch "foo bar" {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    }
    v2 = 4294967295
    fmt.Println(v2, v2)
    v2 = 1
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 int16 = 751
            _ = fi13
            _ = v14
        }
        v2 += v2
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = 476
            }
        }
        _ = fi12
    }
    var v15 int = 1
    v2 = 42
    for fi16 := 0; fi16 < 19; fi16++ {
        fmt.Println(v11)
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
