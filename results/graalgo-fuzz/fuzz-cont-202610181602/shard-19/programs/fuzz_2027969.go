package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
func genFn0(p0 uint, p1 uint16, p2 uint16) int64 {
    v6 := []int{32767, -1, 1, 9223372036854775807}
    for fi7 := 0; fi7 < 3; fi7++ {
        v6 = append(v6, fi7)
        var v8 int64 = 625
        if false {
            _ = len(v6)
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 bool = (true && (((1.5 + 1e10) - (0.1 / 0.1)) < (float32(1e10) - ((1.5 + 3.14159) * (float32(0.0) * 0.0)))))
    v2 := make(map[uint32]string)
    var v3 int64 = -1
    var v4 uint16 = 65535
    var v5 int16 = 1
    fmt.Println(v5, v4)
    v3 = (^(v3))
    if v1 {
        v9 := GenS1{F0: (^(v3)), F1: 0}
        v3 += (^(v3))
        v5 = (v5 / (((v5 / (((v5 >> ((954) & 63))) | 1))) | 1))
        _ = v9
    } else {
        if v1 {
            v3 = (v3 - (v3 * genFn0(255, v4, (^(v4)))))
            if false {
                var v10 int16 = (v5 >> ((540) & 63))
                if (1 == 0) {
                    var v11 int64 = (-(v3))
                    var v12 uint8 = 2
                    _ = v11
                    _ = v12
                }
                _ = v10
            } else {
                for fi13 := 0; fi13 < 12; fi13++ {
                    if (18446744073709551615 <= 0) {
                        var v14 uint8 = 248
                        _ = v14
                    }
                    _ = fi13
                }
                var v15 int = -547
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            switch fi16 {
            case 501:
                var v17 float32 = 1e10
                _ = v17
            case -95:
                switch ("" + "hello") {
                case "foo bar":
                    fmt.Println(v1)
                case "":
                    v18 := GenS2{F0: v1, F1: 572, F2: 255}
                    _ = v18
                default:
                    v3 = 1
                }
            case 4294967295:
                v19 := v2[185]
                _ = v19
            }
            if (1 < 255) {
                switch "foo bar" {
                case "\n\t":
                    var v20 uint16 = v4
                    _ = v20
                }
                if (-9223372036854775808 < 9223372036854775807) {
                    v21 := []int{(^(fi16)), fi16, (fi16 % ((32768) | 1)), fi16}
                    var v22 float64 = 1e10
                    fmt.Println(v22, v22)
                    _ = v21
                    _ = v22
                } else {
                    var v23 int16 = v5
                    switch int(v4) {
                    case 262:
                        v24 := GenS2{F0: false, F1: 32767, F2: 4294967295}
                        _ = v24
                    case 2147483647:
                        v1 = true
                    default:
                        v25 := GenS1{F0: (v3 & 42), F1: 1}
                        _ = v25
                    }
                    v4 = (^(v4))
                    _ = v23
                }
                for fi26 := 0; fi26 < 17; fi26++ {
                    var v27 float64 = (((float64((-2.25 * 1e10)) + float64((float32(float64(float32((1e10 / 1.5)))) / 1e10))) / 1.0) - 1.5)
                    _ = fi26
                    _ = v27
                }
            }
            _ = fi16
        }
    }
    var v28 uint16 = (v4 + (v4 << ((65535) & 63)))
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 int32 = 0
        _ = fi29
        _ = v30
    }
    var v31 int = 753
    fmt.Println(v3)
    var v32 float32 = 3.14159
    v33 := []int{-1, v31}
    var v34 int16 = (v5 % (((-(v5))) | 1))
    if false {
        var v35 uint32 = 2147483647
        v32 += v32
        _ = v35
    } else {
        v5 = (v5 % (((v34 / ((v5) | 1))) | 1))
        v32 += v32
        _ = len(v33)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
