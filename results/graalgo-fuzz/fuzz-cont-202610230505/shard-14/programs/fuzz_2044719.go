package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int; F1 uint16; F2 int64 }
func genFn0(p0 uint64, p1 uint64) uint8 {
    if (true || false) {
        p1 += p0
    } else {
        v6 := GenS1{F0: -38, F1: -891, F2: 4294967295}
        var v7 float64 = 0.0
        var v8 int8 = 90
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 617, F1: 624}
    switch 127 {
    case -611:
        v10 := GenS1{F0: 0, F1: 42, F2: 2147483647}
        _ = v10
    case -668:
        fmt.Println(p1, p1)
    case 32768:
        var v11 bool = (1.5 > 1.5)
        _ = v11
    default:
        var v12 int8 = 1
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        switch (fi13 - (fi13 & (fi13 ^ 128))) {
        case -9223372036854775808:
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 float32 = (0.0 * -2.25)
                if true {
                    v16 := GenS2{F0: 0, F1: 2}
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    v1 := []uint32{65535, 2}
    var v2 int8 = 42
    var v3 float32 = 1.5
    var v4 int16 = -972
    v5 := []int{127, int(v2), 0, 127}
    var v17 uint32 = 2
    var v18 float64 = (3.14159 + 0.1)
    for fi19 := 0; fi19 < 17; fi19++ {
        fmt.Println(v2, v17, v4)
        _ = fi19
    }
    switch ("foo bar" + "a") {
    case "foo bar":
        v20 := GenS3{F0: "", F1: (v17 << ((1) & 63))}
        _ = v20
    case "":
        _ = len(v1)
    default:
        if true {
            v21 := GenS1{F0: 42, F1: v4, F2: (v17 & (v17 / ((v17) | 1)))}
            _ = v21
        } else {
            _ = len(v5)
            v4 -= (v4 & (-(v4)))
        }
    }
    switch "\n\t" {
    case "a":
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 float32 = -2.25
            var v24 int64 = 1
            _ = fi22
            _ = v23
            _ = v24
        }
    case "\n\t":
        if false {
            if true {
                if (155 >= 216) {
                    fmt.Println(v18)
                    v4 = (v4 / (((v4 * (v4 & v4))) | 1))
                }
            }
        }
    case "":
        if false {
            if (false || false) {
                v4 = (v4 << ((255) & 63))
                fmt.Println(v18, v3, v3)
                var v25 int32 = 1
                _ = v25
            } else {
                v3 = 1e10
                var v26 uint64 = 308
                _ = v26
            }
        } else {
            var v27 uint16 = 65535
            v27 = (^(v27))
            if true {
                var v28 string = "a"
                v29 := GenS4{F0: -1, F1: (v27 ^ (v27 >> ((18446744073709551615) & 63))), F2: 9223372036854775807}
                _ = v28
                _ = v29
            }
            _ = v27
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 int = fi30
        _ = fi30
        _ = v31
    }
    var v32 uint64 = 870
    v18 -= 1e10
    v18 += (v18 + float64(v3))
    v18 = v18
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
