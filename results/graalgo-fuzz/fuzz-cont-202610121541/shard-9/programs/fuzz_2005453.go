package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int8; F2 bool }
func genFn0(p0 int8) uint16 {
    var v6 bool = (((0 > 42) || (!false)) && (!false))
    if (((!true) || true) || v6) {
        var v7 int32 = 1
        v8 := GenS1{F0: -119}
        v9 := GenS1{F0: 0}
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: -627}
    fmt.Println(v6, v6)
    v11 := GenS1{F0: 1}
    _ = p0
    _ = v10
    _ = v11
    _ = v6
    return 545
}
func genFn1(p0 int32) bool {
    p0 = (p0 - p0)
    var v12 int8 = 100
    _ = p0
    _ = v12
    return (v12 <= v12)
}

func main() {
    var v1 uint32 = 162
    var v2 int16 = -1
    v3 := make(map[uint32]int)
    var v4 float64 = ((1.5 - 1e10) * 0.0)
    var v5 int16 = 32767
    var v13 uint64 = 44
    var v14 uint16 = genFn0(52)
    var v15 int8 = 100
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int = (fi16 % ((fi16) | 1))
        v18 := GenS2{F0: "\n\t", F1: -2147483648, F2: 0}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := v3[v1]
    fmt.Println(v1, v13, v1)
    var v20 float32 = -2.25
    v21 := []int{(v19 | v19), (v19 & v19), (v19 - (v19 * v19)), (v19 / ((42) | 1))}
    var v22 int32 = -684
    var v23 float32 = float32((float64(v20) / v4))
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 int = (v19 | (fi24 & v19))
        for fi26 := 0; fi26 < 15; fi26++ {
            v27 := GenS2{F0: ("" + "\n\t"), F1: (v22 * (v22 << ((2) & 63))), F2: 1}
            _ = fi26
            _ = v27
        }
        var v28 uint = 323
        _ = fi24
        _ = v25
        _ = v28
    }
    var v29 uint8 = 142
    var v30 uint64 = v13
    switch "\n\t" {
    case "":
        v31 := GenS2{F0: "foo bar", F1: (^(v22)), F2: (v29 ^ 0)}
        _ = v31
    case "hello":
        v5 = int16(v19)
    default:
        if false {
            for fi32 := 0; fi32 < 8; fi32++ {
                for fi33 := 0; fi33 < 2; fi33++ {
                    fmt.Println(v20, v23, v19)
                    v5 -= (v5 & -350)
                    _ = len(v21)
                    _ = fi33
                }
                v34 := GenS2{F0: "", F1: v22, F2: v29}
                var v35 uint = 2
                _ = fi32
                _ = v34
                _ = v35
            }
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
        } else {
            if (v15 >= v15) {
                v36 := GenS1{F0: -32768}
                _ = v36
            } else {
                for fi37 := 0; fi37 < 15; fi37++ {
                    var v38 int = 20
                    switch v38 {
                    case 4294967295:
                        var v39 uint8 = 151
                        _ = v39
                    case -9223372036854775808:
                        var v40 int64 = 1
                        _ = v40
                    case 597:
                        fmt.Println(v2)
                    default:
                        v41 := GenS3{F0: v14, F1: 1, F2: (false && false)}
                        _ = v41
                    }
                    _ = fi37
                    _ = v38
                }
                v42 := GenS2{F0: ("foo bar" + (("" + "hello") + "foo bar")), F1: v22, F2: v29}
                _ = v42
            }
            v43 := GenS1{F0: v2}
            fmt.Println(v13, v20, v15)
            _ = v43
        }
    }
    fmt.Println(v30, v22)
    var v44 string = ("a" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v5
}
