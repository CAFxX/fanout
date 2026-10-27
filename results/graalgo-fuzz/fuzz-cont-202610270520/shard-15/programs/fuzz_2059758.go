package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int16; F1 uint32; F2 bool }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint64; F1 uint; F2 bool }
type GenS5 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int; F1 uint; F2 uint8 }
type GenS7 struct { F0 int; F1 int8; F2 uint32 }
func genFn0(p0 uint32, p1 bool) int8 {
    var v7 uint8 = 1
    var v8 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -60
}
func genFn1() uint64 {
    var v9 float64 = ((1e10 + (1.5 - 0.0)) * (-2.25 / float64(((3.14159 / 1.5) + 1e10))))
    v10 := []int{9223372036854775807, 9223372036854775807}
    for fi11 := 0; fi11 < 3; fi11++ {
        v9 += float64(-2.25)
        _ = fi11
    }
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := []string{"a", "a"}
    v2 := []string{"", "\n\t"}
    var v3 uint16 = 1
    var v4 string = ("" + ("\n\t" + "hello"))
    var v5 uint16 = (v3 ^ (v3 | 1))
    var v6 uint8 = 173
    var v12 int32 = 1
    _ = len(v1)
    if len(v2) > 0 {
        v2[0] = v4
    }
    v5 = 958
    var v13 int16 = 32767
    fmt.Println(v3)
    if (!false) {
        v12 = int32(v6)
        v14 := GenS1{F0: 42, F1: -9223372036854775808}
        for fi15 := 0; fi15 < 3; fi15++ {
            switch fi15 {
            case 32767:
                for fi16 := 0; fi16 < 16; fi16++ {
                    v17 := GenS2{F0: (v13 | (v13 + (v13 & (v13 | (v13 / (((v13 * 263)) | 1)))))), F1: 807, F2: (v12 == (^(v12)))}
                    v18 := GenS3{F0: (v13 << ((42) & 63)), F1: 42}
                    switch v4 {
                    case "a":
                        v13 += v13
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            case 32768:
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 int64 = 0
                    var v21 uint64 = 42
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            case -985:
                var v22 int32 = v12
                _ = v22
            default:
                switch v4 {
                case "\n\t":
                    v23 := GenS4{F0: 977, F1: 2, F2: true}
                    _ = v23
                }
            }
            for fi24 := 0; fi24 < 4; fi24++ {
                v25 := GenS5{F0: (v13 / ((v13) | 1)), F1: (v5 << ((2) & 63)), F2: 0}
                v1 = append(v1, (v4 + (("\n\t" + (v4 + "a")) + ("\n\t" + (v4 + (v4 + v4))))))
                _ = fi24
                _ = v25
            }
            _ = fi15
        }
        _ = v14
    } else {
        var v26 uint = 1
        _ = v26
    }
    switch -2147483648 {
    case 32767:
        if false {
            _ = len(v1)
            switch 128 {
            case 128:
                v1 = append(v1, v4)
            case 1:
                fmt.Println(v5, v4)
            }
            v27 := GenS6{F0: 1, F1: 2, F2: v6}
            _ = v27
        } else {
            var v28 int8 = -1
            _ = v28
        }
    default:
        switch -179 {
        case 1:
            var v29 int16 = v13
            _ = v29
        case 127:
            switch v4 {
            case "":
                var v30 uint16 = 519
                _ = v30
            default:
                v31 := GenS7{F0: -2147483648, F1: 100, F2: 2147483648}
                _ = v31
            }
        case 32768:
            fmt.Println(v3)
        default:
            _ = len(v1)
        }
    }
    v4 = v4
    var v32 int = -9223372036854775808
    v32 -= (v32 | (^(v32)))
    fmt.Println(v5, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
