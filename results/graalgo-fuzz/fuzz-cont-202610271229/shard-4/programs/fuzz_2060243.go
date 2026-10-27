package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int32; F1 int; F2 string }
func genFn0(p0 int16, p1 uint32) uint64 {
    v7 := GenS1{F0: -127}
    v8 := GenS2{F0: false, F1: false}
    switch -842 {
    case 42:
        var v9 int16 = p0
        _ = v9
    case -619:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 289
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 0
    v3 := make(map[uint32]uint32)
    v4 := []int{0, -112, int(v2), -9223372036854775808}
    var v5 int16 = -593
    var v6 uint = 255
    v10 := GenS3{F0: 0}
    var v11 int16 = v5
    v6 = (v6 << ((18446744073709551615) & 63))
    var v12 float64 = (float64(0.0) / 1.0)
    v13 := GenS4{F0: 721, F1: -1}
    for fi14 := 0; fi14 < 3; fi14++ {
        fmt.Println(v11, v12, fi14)
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 float64 = (v12 / (0.0 * v12))
            v5 += (v11 ^ v2)
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            fmt.Println(fi14, v12)
            v18 := GenS2{F0: (1 < 177), F1: false}
            _ = len(v3)
            _ = fi17
            _ = v18
        }
        _ = fi14
    }
    if false {
        if (!(!false)) {
            if len(v4) > 0 {
                v4[0] = 128
            }
            v11 -= v2
        } else {
            v4 = append(v4, 9223372036854775807)
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 uint8 = 42
                var v21 uint64 = genFn0((^(v11)), 0)
                _ = fi19
                _ = v20
                _ = v21
            }
            switch ("hello" + ("" + "\n\t")) {
            case "\n\t":
                var v22 uint = (v6 << ((2) & 63))
                _ = v22
            case "hello":
                var v23 int16 = (-(v11))
                _ = v23
            }
        }
    }
    v24 := GenS2{F0: (0.1 <= v12), F1: (1e10 < 3.14159)}
    v25 := GenS5{F0: (-(v11)), F1: -1}
    v5 += (v5 | v11)
    if len(v4) > 0 {
        v4[0] = 128
    }
    v26 := GenS6{F0: 42, F1: 1, F2: ""}
    var v27 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
