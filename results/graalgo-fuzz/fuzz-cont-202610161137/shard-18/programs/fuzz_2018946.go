package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int32; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int; F2 uint }
func genFn0(p0 int32, p1 uint32, p2 bool) uint64 {
    var v4 int = 42
    v4 = 9223372036854775807
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := []int{(fi5 + (fi5 & fi5)), -2147483648}
        if (0 > -32768) {
            v7 := GenS1{F0: 0}
            var v8 uint8 = 42
            _ = v7
            _ = v8
        }
        _ = len(v6)
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 153
}
func genFn1(p0 string, p1 uint32) bool {
    var v9 int8 = -58
    v10 := GenS1{F0: 738}
    if (p0 <= p0) {
        v11 := GenS2{F0: 9223372036854775807, F1: 42, F2: 1}
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := []int{}
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        fmt.Println(p1)
        if (((p1 >= 42) || (false && false)) || (831 < 42)) {
            if (278 != 2) {
                switch p0 {
                case "a":
                    var v14 uint32 = (p1 + p1)
                    _ = v14
                case "hello":
                    p0 = (p0 + "a")
                case "\n\t":
                    for fi15 := 0; fi15 < 5; fi15++ {
                        var v16 uint16 = 32767
                        _ = fi15
                        _ = v16
                    }
                default:
                    var v17 uint8 = 146
                    _ = v17
                }
                if ((1e10 != 1.5) || (int32(p1) != -2147483648)) {
                    if false {
                        var v18 uint64 = 4294967295
                        _ = v18
                    }
                }
            } else {
                v19 := []int{}
                switch 830 {
                case -499:
                    var v20 float64 = 0.0
                    _ = v20
                }
                v21 := GenS3{F0: 0}
                _ = v19
                _ = v21
            }
            switch p0 {
            case "\n\t":
                v9 -= v9
            case "foo bar":
                var v22 int32 = 296
                _ = v22
            case "hello":
                if true {
                    v23 := GenS2{F0: 9223372036854775807, F1: 42, F2: 42}
                    v24 := []int{}
                    _ = v23
                    _ = v24
                } else {
                    v25 := []int{449}
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                    _ = v25
                }
            default:
                fmt.Println(p1, v9, p0)
            }
            v26 := []int{127, -2147483648, 282}
            _ = v26
        } else {
            v9 = v9
            v9 = 100
        }
    }
    v27 := GenS2{F0: 0, F1: 1, F2: 2}
    _ = p0
    _ = p1
    _ = v10
    _ = v27
    _ = v9
    return true
}

func main() {
    var v1 int16 = 0
    v2 := []uint32{892, 490, 42, 549}
    var v3 uint64 = 962
    var v28 float32 = ((1e10 + 0.1) - 0.0)
    switch "" {
    case "hello":
        var v29 int16 = v1
        _ = v29
    default:
        for fi30 := 0; fi30 < 8; fi30++ {
            fmt.Println(v3, fi30, v3)
            _ = fi30
        }
    }
    v3 -= (v3 ^ (^(v3)))
    var v31 int16 = (v1 / ((v1) | 1))
    if len(v2) > 0 {
        v2[0] = 506
    }
    _ = len(v2)
    var v32 uint = 1
    var v33 uint8 = 127
    v34 := GenS3{F0: v33}
    var v35 uint32 = 472
    var v36 uint32 = (v35 << (((^(v32))) & 63))
    v37 := GenS4{F0: -59, F1: 42, F2: (^(v32))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
