package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint; F2 uint64 }
type GenS4 struct { F0 int16; F1 string; F2 uint }
type GenS5 struct { F0 int; F1 int; F2 int32 }
type GenS6 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint32, p1 int64) uint16 {
    v7 := GenS1{F0: 944, F1: (p0 << ((4294967295) & 63)), F2: -50}
    p1 += -1
    var v8 bool = (true || true)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return uint16(p0)
}
func genFn1(p0 uint32) int16 {
    v9 := GenS1{F0: 0, F1: 65535, F2: -930}
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS2{F0: 1000, F1: 127}
        v12 := GenS3{F0: 42, F1: 845, F2: 542}
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 int = -1
    v3 := []string{"foo bar", "a"}
    var v4 float32 = 1.5
    v5 := []uint32{2147483648}
    var v6 int32 = 2147483647
    v6 = (v6 >> ((42) & 63))
    var v13 string = ""
    _ = len(v5)
    var v14 bool = false
    v15 := GenS2{F0: -552, F1: 210}
    var v16 int64 = 0
    var v17 uint16 = 255
    v14 = (!v14)
    v16 -= (v16 + v16)
    if (!true) {
        switch v13 {
        case "foo bar":
            for fi18 := 0; fi18 < 11; fi18++ {
                v16 -= v16
                var v19 uint32 = 4294967295
                _ = fi18
                _ = v19
            }
        default:
            var v20 float64 = (0.1 * (3.14159 + 0.0))
            _ = v20
        }
    } else {
        var v21 float32 = (float32((float64((float32(-2.25) / float32(float64((((-2.25 - 0.1) * -2.25) - v4))))) + (1e10 / -2.25))) / v4)
        fmt.Println(v6, v17)
        delete(v1, ("" + v13))
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        if (100 < 0) {
            v6 -= (v6 - v6)
        } else {
            v23 := GenS1{F0: 40, F1: 4294967295, F2: genFn1(400)}
            var v24 int8 = 0
            for fi25 := 0; fi25 < 4; fi25++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v17, v2)
                v26 := GenS4{F0: genFn1(341), F1: v13, F2: 255}
                _ = fi25
                _ = v26
            }
            _ = v23
            _ = v24
        }
        v13 = ("hello" + v13)
        switch fi22 {
        case 127:
            var v27 uint8 = 127
            _ = v27
        default:
            if ((!(-2.25 == -2.25)) || (!v14)) {
                v28 := GenS5{F0: fi22, F1: (fi22 - (v2 >> ((1) & 63))), F2: v6}
                v2 = (fi22 << ((695) & 63))
                _ = v28
            }
        }
        _ = fi22
    }
    v29 := v1["a"]
    switch v2 {
    case 1:
        v30 := GenS6{F0: v16, F1: 0}
        _ = v30
    case 208:
        v29 += v2
    case -8:
        var v31 uint8 = 0
        _ = v31
    }
    v6 -= 0
    v16 -= v16
    v29 += 42
    v32 := v1[v13]
    v33 := GenS1{F0: 0, F1: 86, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
