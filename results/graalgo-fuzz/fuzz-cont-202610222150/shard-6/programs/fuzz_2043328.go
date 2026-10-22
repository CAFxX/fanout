package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int, p1 int16) uint {
    p0 += -1
    fmt.Println(p0, p1)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v7
    return 197
}

func main() {
    var v1 uint8 = 127
    var v2 uint = 538
    var v3 int16 = int16(v2)
    var v4 int = 128
    var v5 float32 = (0.0 / 1e10)
    v6 := make(map[string]uint32)
    var v8 uint32 = 231
    v5 = float32(float64(1.5))
    if false {
        var v9 int = 32767
        _ = v9
    }
    var v10 uint8 = v1
    v11 := GenS1{F0: (v1 * (v10 ^ v1)), F1: 1}
    var v12 int64 = 308
    v13 := GenS2{F0: 2147483647, F1: -1, F2: 2}
    for fi14 := 0; fi14 < 13; fi14++ {
        if true {
            fmt.Println(v3, fi14, v8)
            var v15 bool = ((v8 & v8) >= v8)
            switch "a" {
            case "foo bar":
                _ = len(v6)
            case "\n\t":
                v12 -= (-(v12))
            case "a":
                for fi16 := 0; fi16 < 14; fi16++ {
                    if v15 {
                        v17 := v6[""]
                        if (!(477 == (v8 << ((v2) & 63)))) {
                            v18 := GenS3{F0: 4294967295, F1: (v2 << ((612) & 63)), F2: ("foo bar" + "foo bar")}
                            v1 = (^(v10))
                            _ = v18
                        }
                        if v15 {
                            v19 := []int{(v4 - (fi14 / (((fi14 - (fi16 % ((0) | 1)))) | 1))), (v4 * 0), (^(v4)), (^(fi14))}
                            _ = v19
                        }
                        _ = v17
                    }
                    _ = fi16
                }
            }
            _ = v15
        }
        _ = fi14
    }
    _ = len(v6)
    var v20 uint32 = (v8 | (v8 - 2))
    var v21 uint8 = (v1 >> (((v2 / (((v2 - 42)) | 1))) & 63))
    v22 := GenS4{F0: ""}
    switch 42 {
    case 0:
        switch "foo bar" {
        case "foo bar":
            for fi23 := 0; fi23 < 5; fi23++ {
                for fi24 := 0; fi24 < 4; fi24++ {
                    v12 += 12
                    fmt.Println(v20, v20)
                    _ = fi24
                }
                _ = fi23
            }
        }
    case 2147483647:
        var v25 int32 = -102
        _ = v25
    case -2147483648:
        switch v4 {
        case 1:
            v1 -= v1
        case -684:
            fmt.Println(v12, v21, v12)
        default:
            v26 := GenS3{F0: 898, F1: (v2 * 4294967295), F2: "\n\t"}
            _ = v26
        }
    }
    v10 = v1
    v27 := []int{v4, (v4 - v4)}
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    var v28 uint16 = 1
    v29 := GenS5{F0: (v21 ^ (^(v10))), F1: (v28 << (((v2 / ((18446744073709551615) | 1))) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
