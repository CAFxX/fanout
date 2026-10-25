package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 bool }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint16 = 65535
    var v2 string = ("" + "a")
    v3 := make(map[string]uint32)
    var v4 uint32 = 4294967295
    v5 := make(map[uint32]bool)
    v1 -= v1
    var v6 uint64 = 37
    if false {
        switch v2 {
        case "":
            v7 := GenS1{F0: -5, F1: 1, F2: true}
            _ = v7
        default:
            v8 := GenS1{F0: 128, F1: 1, F2: false}
            _ = v8
        }
        if ((v1 ^ 0) >= 42) {
            var v9 int8 = -128
            _ = v9
        } else {
            v5[(v4 - (v4 * 2))] = true
            v10 := GenS1{F0: 99, F1: -1, F2: false}
            _ = v10
        }
        var v11 int16 = -670
        _ = v11
    } else {
        v12 := GenS2{F0: int8(v6), F1: ((v2 + ((v2 + ((("hello" + ("foo bar" + ("a" + v2))) + "a") + v2)) + v2)) + v2)}
        fmt.Println(v6, v4)
        fmt.Println(v1, v1, v6)
        _ = v12
    }
    var v13 float64 = 3.14159
    switch v2 {
    case "hello":
        v4 = (v4 | v4)
    case "foo bar":
        for fi14 := 0; fi14 < 17; fi14++ {
            fmt.Println(v13, v1, v2)
            _ = fi14
        }
    default:
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 float64 = v13
            var v17 float64 = (v13 - 1.5)
            for fi18 := 0; fi18 < 10; fi18++ {
                fmt.Println(fi15)
                v1 += (v1 << ((uint(v1)) & 63))
                _ = fi18
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    var v19 uint32 = v4
    var v20 uint = 135
    var v21 int64 = 42
    switch v2 {
    case "foo bar":
        fmt.Println(v13, v4, v2)
    case "":
        var v22 uint16 = v1
        _ = v22
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS2{F0: -21, F1: "a"}
        v19 = 2147483648
        _ = fi23
        _ = v24
    }
    fmt.Println(v19, v4)
    for fi25 := 0; fi25 < 3; fi25++ {
        v6 = v6
        v26 := GenS3{F0: (v20 << ((v20) & 63))}
        _ = fi25
        _ = v26
    }
    var v27 int32 = 0
    v5[(v4 * 65535)] = true
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
