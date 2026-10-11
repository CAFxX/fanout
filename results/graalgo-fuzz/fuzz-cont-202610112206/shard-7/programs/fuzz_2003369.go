package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int64, p2 int) int64 {
    var v5 bool = true
    var v6 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p1 + (p1 << ((4294967295) & 63)))
}
func genFn1() int8 {
    var v7 float64 = (0.1 - (3.14159 * 0.1))
    for fi8 := 0; fi8 < 16; fi8++ {
        switch "\n\t" {
        case "hello":
            fmt.Println(fi8)
        default:
            v7 = v7
        }
        _ = fi8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int32 = 0
        var v11 uint64 = 180
        fmt.Println(v7, fi9)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v7
    return -128
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 0.0
    v3 := make(map[string]int)
    var v4 string = "\n\t"
    fmt.Println(v4)
    var v12 float32 = float32(3.14159)
    var v13 uint32 = 42
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 int16 = -1
        _ = fi14
        _ = v15
    }
    fmt.Println(v4)
    if (2 < 1) {
        v16 := []int{744, 4294967295}
        if (!(0.0 < 0.1)) {
            for fi17 := 0; fi17 < 4; fi17++ {
                for fi18 := 0; fi18 < 10; fi18++ {
                    v12 -= 0.1
                    switch v4 {
                    case "a":
                        v19 := GenS1{F0: 42}
                        _ = v19
                    default:
                        v20 := GenS2{F0: 153, F1: -1}
                        _ = v20
                    }
                    v21 := v1[fi18]
                    _ = fi18
                    _ = v21
                }
                _ = fi17
            }
            var v22 int = 42
            _ = v22
        } else {
            if (-59 == 32767) {
                v23 := v1[128]
                v13 = (v13 % (((^(v13))) | 1))
                v16 = append(v16, v23)
                _ = v23
            } else {
                var v24 uint16 = 255
                _ = v24
            }
            if (false && true) {
                v3[(v4 + v4)] = 388
                v25 := GenS2{F0: 696, F1: 0}
                _ = v25
            } else {
                switch 274 {
                case 472:
                    v16 = append(v16, 9223372036854775807)
                default:
                    v26 := GenS3{F0: v13, F1: -117, F2: 2}
                    _ = v26
                }
            }
            var v27 int32 = 1
            _ = v27
        }
        _ = v16
    }
    if (!(1 != -60)) {
        v28 := GenS1{F0: 1}
        if (true && (false || (true || ((-2.25 - (float64(float32(0.0)) / 1e10)) < 1e10)))) {
            delete(v3, v4)
            switch 127 {
            case 127:
                var v29 int64 = -982
                _ = v29
            default:
                if (2 < 42) {
                    v12 -= v2
                    switch v4 {
                    case "\n\t":
                        delete(v3, (v4 + (v4 + v4)))
                    case "":
                        v30 := GenS2{F0: 554, F1: 42}
                        _ = v30
                    case "foo bar":
                        var v31 bool = true
                        _ = v31
                    }
                }
            }
            v4 = v4
        }
        _ = v28
    } else {
        v32 := GenS1{F0: 1}
        v33 := GenS4{F0: 0}
        _ = v32
        _ = v33
    }
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
