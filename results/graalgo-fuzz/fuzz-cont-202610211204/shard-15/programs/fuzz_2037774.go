package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 bool; F1 int32; F2 uint32 }

func main() {
    v1 := []string{}
    var v2 uint16 = 170
    var v3 uint32 = 65535
    var v4 int32 = 0
    v5 := make(map[int]bool)
    v6 := GenS1{F0: (v2 + v2), F1: -297}
    var v7 bool = (!false)
    fmt.Println(v7)
    var v8 uint32 = v3
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 float64 = 3.14159
        var v11 uint16 = 2
        _ = fi9
        _ = v10
        _ = v11
    }
    if ((^(v4)) < (v4 - (v4 * 707))) {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 int8 = -128
            fmt.Println(v4)
            _ = fi12
            _ = v13
        }
        v14 := v5[562]
        _ = v14
    }
    var v15 bool = false
    _ = len(v1)
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 float64 = 0.0
        v18 := GenS1{F0: (v2 | v2), F1: fi16}
        _ = fi16
        _ = v17
        _ = v18
    }
    if (("a" + "foo bar") < "foo bar") {
        v19 := GenS1{F0: (v2 << ((1) & 63)), F1: 0}
        for fi20 := 0; fi20 < 1; fi20++ {
            fmt.Println(v7, v3)
            _ = fi20
        }
        _ = v19
    } else {
        v21 := GenS2{F0: (135 > 0), F1: v4, F2: (v3 >> ((255) & 63))}
        _ = v21
    }
    v22 := GenS1{F0: 542, F1: 4294967295}
    for fi23 := 0; fi23 < 14; fi23++ {
        if (((-979 < int64(fi23)) && true) || (!v15)) {
            var v24 float32 = 0.1
            if (((127 <= 42) && (float32(0.0) < ((v24 * 1e10) - float32(3.14159)))) && ((v24 * float32(0.0)) == 1.5)) {
                v3 = v3
            }
            _ = v24
        } else {
            if ((-32768 < 32767) || true) {
                var v25 uint = 2
                v8 = (v8 | v3)
                _ = v25
            } else {
                v15 = v15
                fmt.Println(v4, v4, v15)
                v7 = (119 >= 255)
            }
        }
        v26 := v5[fi23]
        _ = fi23
        _ = v26
    }
    v3 += (v3 << ((500) & 63))
    var v27 uint16 = v2
    v1 = append(v1, ("a" + "\n\t"))
    v27 = (v27 * 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
