package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := []int{-2147483648, 128, 1, 32768}
    var v2 uint8 = 42
    var v3 int = 4294967295
    var v4 uint16 = 255
    v2 += v2
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            v3 += v3
            switch "a" {
            case "a":
                var v7 int32 = -2147483648
                _ = v7
            case "":
                v4 -= (^(v4))
            default:
                var v8 bool = (!false)
                var v9 int64 = 104
                var v10 int16 = -651
                _ = v10
                _ = v8
                _ = v9
            }
            v4 = (^(v4))
            _ = fi6
        }
        _ = fi5
    }
    var v11 int8 = 0
    if true {
        var v12 string = "hello"
        _ = v12
    } else {
        v13 := GenS1{F0: 0, F1: -1}
        for fi14 := 0; fi14 < 4; fi14++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi14
        }
        if (!true) {
            var v15 float64 = 3.14159
            _ = v15
        } else {
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS2{F0: 65535, F1: 42}
                _ = fi16
                _ = v17
            }
        }
        _ = v13
    }
    var v18 uint16 = v4
    v18 = (v4 % ((42) | 1))
    if (42 <= 208) {
        v19 := GenS2{F0: 4294967295, F1: 2}
        var v20 float64 = (-2.25 - float64(1e10))
        _ = v19
        _ = v20
    }
    v3 += (v3 % (((v3 % ((v3) | 1))) | 1))
    if (!(!((v2 >> ((1) & 63)) == (v2 & v2)))) {
        v21 := GenS2{F0: 2147483648, F1: 2}
        var v22 int32 = 0
        _ = len(v1)
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 1; fi23++ {
            fmt.Println(v4, v4)
            _ = fi23
        }
        v24 := GenS1{F0: 4294967295, F1: -32768}
        _ = v24
    }
    var v25 uint = 65535
    var v26 int = (v3 >> ((0) & 63))
    v4 += 0
    v27 := GenS1{F0: 42, F1: 0}
    if true {
        var v28 int64 = 0
        _ = v28
    } else {
        fmt.Println(v2, v25)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v26 += v3
    }
    v29 := GenS2{F0: 0, F1: uint32(v25)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
}
