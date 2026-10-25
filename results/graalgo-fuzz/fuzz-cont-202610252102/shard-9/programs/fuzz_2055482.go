package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int; F2 uint32 }

func main() {
    v1 := []uint32{315, 0, 2147483647}
    var v2 int8 = -128
    v3 := make(map[string]int)
    v4 := []uint32{724, 42, 65535, 4294967295}
    var v5 int8 = (-(v2))
    fmt.Println(v2)
    if true {
        v2 = (v5 - 1)
        v2 -= v5
        var v6 uint16 = 2
        _ = v6
    } else {
        v2 -= v2
        var v7 float64 = float64(0.1)
        if true {
            var v8 bool = true
            var v9 int8 = v2
            _ = v8
            _ = v9
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 int8 = v2
                v11 = v5
                _ = fi10
                _ = v11
            }
            v7 = (v7 - v7)
        }
        _ = v7
    }
    fmt.Println(v2, v5)
    v5 -= v2
    v5 += v2
    switch "a" {
    case "hello":
        if ((0.0 < (0.1 * (0.0 * (3.14159 / (1e10 + ((float64(float32((-2.25 - 0.0))) * 1e10) + float64(1e10))))))) && true) {
            var v12 float32 = (-2.25 - 1e10)
            var v13 int16 = 0
            _ = v12
            _ = v13
        }
    }
    v5 -= -128
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := GenS1{F0: 255, F1: 184}
        v16 := GenS2{F0: 18446744073709551615, F1: 42, F2: 906}
        v17 := GenS2{F0: 4294967295, F1: (fi14 >> ((1) & 63)), F2: 4294967295}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 float64 = -2.25
    v19 := GenS2{F0: 2, F1: 778, F2: 2147483648}
    switch 117 {
    case -29:
        for fi20 := 0; fi20 < 7; fi20++ {
            v5 = (v2 << ((436) & 63))
            _ = fi20
        }
    case -197:
        var v21 int32 = -397
        _ = v21
    case 203:
        v5 = 42
    }
    if ((0.0 * (float32(1e10) + 3.14159)) >= (3.14159 / 1e10)) {
        switch 128 {
        case -455:
            var v22 string = ("\n\t" + ("hello" + ((("" + "\n\t") + "") + "hello")))
            _ = v22
        }
        _ = len(v3)
    } else {
        for fi23 := 0; fi23 < 14; fi23++ {
            v4 = append(v4, 65535)
            _ = len(v4)
            _ = fi23
        }
        v2 += v5
        fmt.Println(v18)
    }
    if true {
        var v24 int32 = 830
        v25 := GenS1{F0: 173, F1: 255}
        _ = v24
        _ = v25
    }
    var v26 uint8 = 255
    var v27 int = 42
    v27 = (v27 + (^(v27)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
