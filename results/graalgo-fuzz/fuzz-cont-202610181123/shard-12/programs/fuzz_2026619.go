package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int64; F2 bool }

func main() {
    var v1 float64 = float64(-2.25)
    var v2 int8 = -82
    v3 := []uint32{}
    var v4 uint = 42
    var v5 bool = (1 >= -32768)
    v6 := make(map[int]string)
    _ = len(v6)
    if (!(!true)) {
        if v5 {
            delete(v6, -1)
        }
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int = fi7
        v2 = v2
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: -707}
    v10 := GenS2{F0: 0, F1: -1, F2: (false || (-976 > -1))}
    v11 := GenS1{F0: 198}
    v1 = (float64(float32(0.1)) + 0.1)
    var v12 uint = v4
    var v13 int16 = 32767
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS2{F0: -639, F1: -1, F2: (fi14 >= -9223372036854775808)}
        v13 -= (v13 | (v13 & int16(v2)))
        var v16 float32 = 1.5
        _ = fi14
        _ = v15
        _ = v16
    }
    if v5 {
        var v17 uint32 = 249
        _ = v17
    } else {
        var v18 uint16 = 810
        switch ("\n\t" + ("\n\t" + ("hello" + "\n\t"))) {
        case "":
            var v19 uint16 = v18
            _ = v19
        case "foo bar":
            delete(v6, -9223372036854775808)
        default:
            v20 := GenS1{F0: 0}
            _ = v20
        }
        _ = v18
    }
    if (0.1 == 1e10) {
        if (51 != 2) {
            if (v5 && v5) {
                for fi21 := 0; fi21 < 10; fi21++ {
                    if v5 {
                        fmt.Println(v5)
                    }
                    _ = fi21
                }
                v12 = 944
            }
            fmt.Println(v4, v13, v5)
            v12 = v12
        }
        var v22 int32 = 0
        _ = v22
    } else {
        v2 = v2
    }
    v13 = (v13 >> ((2) & 63))
    var v23 int16 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
