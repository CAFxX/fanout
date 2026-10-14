package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{820, 153, 65535, 1}
    var v3 uint16 = 42
    var v4 uint16 = (v3 - (v3 * (v3 - (v3 >> ((42) & 63)))))
    v4 = (^(v4))
    if false {
        v5 := GenS1{F0: 149, F1: uint32(v4), F2: -1}
        v6 := v1[0]
        _ = v5
        _ = v6
    } else {
        _ = len(v2)
        var v7 uint = 468
        _ = v7
    }
    v4 += v3
    v3 = v3
    if (!true) {
        for fi8 := 0; fi8 < 9; fi8++ {
            v4 = (v4 ^ v3)
            var v9 int = (fi8 - (fi8 & fi8))
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 int32 = 2147483647
            _ = fi10
            _ = v11
        }
    } else {
        var v12 uint = 1
        _ = v12
    }
    v13 := v1[-961]
    v14 := GenS2{F0: 42, F1: (v13 / ((v13) | 1))}
    fmt.Println(v4, v4, v3)
    fmt.Println(v3, v13, v13)
    var v15 float64 = 1.5
    var v16 uint32 = v13
    switch ("\n\t" + "a") {
    case "foo bar":
        v16 = (^(v13))
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 uint32 = 0
        _ = fi17
        _ = v18
    }
    var v19 string = ""
    if false {
        var v20 int = -245
        for fi21 := 0; fi21 < 8; fi21++ {
            v22 := GenS2{F0: 213, F1: v16}
            _ = fi21
            _ = v22
        }
        v13 += (v13 << ((4294967295) & 63))
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
