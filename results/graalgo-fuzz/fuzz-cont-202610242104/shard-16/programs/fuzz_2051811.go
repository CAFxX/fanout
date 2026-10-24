package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 int16, p1 uint64, p2 int) uint32 {
    switch p2 {
    case 2147483647:
        switch 42 {
        case 32768:
            var v5 uint16 = 269
            _ = v5
        case 42:
            v6 := GenS1{F0: -2147483648, F1: 65535}
            _ = v6
        }
    case -618:
        fmt.Println(p0, p1)
    case 128:
        p2 = (p2 << ((912) & 63))
    }
    if false {
        v7 := GenS1{F0: -387, F1: 0}
        v8 := GenS2{F0: 86, F1: -1}
        _ = v7
        _ = v8
    }
    var v9 int8 = 1
    switch (p2 / ((p2) | 1)) {
    case 2147483647:
        var v10 int = (p2 ^ (p2 >> ((42) & 63)))
        _ = v10
    case -1:
        var v11 float64 = 3.14159
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        p2 = (p2 | int(v9))
        var v13 string = "hello"
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2
}
func genFn1(p0 string, p1 int64) uint16 {
    p0 = p0
    p1 = (p1 - p1)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 string = ""
    v2 := make(map[string]string)
    v3 := []string{}
    var v4 int = 9223372036854775807
    v4 = (v4 + (v4 & (v4 & (v4 % (((v4 % ((1) | 1))) | 1)))))
    switch (v4 % ((v4) | 1)) {
    case 1:
        for fi14 := 0; fi14 < 20; fi14++ {
            _ = len(v3)
            _ = fi14
        }
    case 393:
        v4 = 32768
    }
    v3 = append(v3, "a")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = 32768
    fmt.Println(v4, v4)
    var v15 int32 = 882
    var v16 int64 = 92
    if false {
        fmt.Println(v15)
    } else {
        v17 := GenS2{F0: 227, F1: -906}
        var v18 int32 = 124
        _ = v17
        _ = v18
    }
    _ = len(v2)
    var v19 int32 = v15
    v20 := v2[(((v1 + "hello") + (((v1 + v1) + v1) + v1)) + v1)]
    switch (^(v4)) {
    case 2147483647:
        switch v4 {
        case 4294967295:
            var v21 bool = false
            _ = v21
        case -680:
            fmt.Println(v16)
        default:
            v15 -= (^(v19))
        }
    case 611:
        var v22 uint16 = 678
        _ = v22
    default:
        v23 := GenS1{F0: 42, F1: genFn0(42, 1, (v4 * (^(v4))))}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
