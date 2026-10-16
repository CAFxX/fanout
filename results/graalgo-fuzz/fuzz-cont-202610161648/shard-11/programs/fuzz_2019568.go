package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int32; F1 int8; F2 uint32 }
func genFn0(p0 int8, p1 uint32) int16 {
    for fi4 := 0; fi4 < 18; fi4++ {
        p1 += (p1 % ((p1) | 1))
        _ = fi4
    }
    v5 := GenS1{F0: 42, F1: (p0 >> ((2) & 63))}
    v6 := GenS1{F0: 1, F1: (p0 + 127)}
    p0 -= p0
    var v7 uint16 = 255
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int64 = 0
    var v2 int16 = 42
    v3 := []int{-123, 423, 2147483647}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 string = "a"
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS1{F0: 536, F1: 13}
        if (!(((fi10 | 32767) == (fi10 << ((649) & 63))) || (1e10 != 1e10))) {
            if ((((1e10 - 3.14159) != (3.14159 + (float64(-2.25) / (1e10 + 1e10)))) && (v1 != -1)) && true) {
                if (65535 >= 65535) {
                    v12 := GenS1{F0: 283, F1: 93}
                    _ = v12
                }
                if len(v3) > 0 {
                    v3[0] = 32767
                }
            }
            v3 = append(v3, 4294967295)
        } else {
            v1 = (v1 + v1)
        }
        if false {
            switch fi10 {
            case 815:
                v1 -= v1
            case 693:
                _ = len(v3)
            case 0:
                var v13 int16 = v2
                _ = v13
            }
            v2 += 42
            var v14 int = (fi10 << ((2) & 63))
            _ = v14
        } else {
            var v15 float64 = float64(1.5)
            var v16 float32 = (1.5 - float32(v15))
            var v17 uint8 = 0
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi10
        _ = v11
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        switch fi18 {
        case 9223372036854775807:
            _ = len(v3)
        case 42:
            v3 = append(v3, fi18)
        default:
            var v19 float64 = float64(0.1)
            _ = v19
        }
        _ = fi18
    }
    v1 -= v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v20 uint64 = 42
    var v21 uint16 = 42
    v20 += (v20 >> ((42) & 63))
    switch 1 {
    case -9223372036854775808:
        v22 := GenS2{F0: 0, F1: 42, F2: 348}
        _ = v22
    case -670:
        var v23 uint8 = 131
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
