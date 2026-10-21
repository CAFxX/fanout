package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 bool; F1 int64; F2 int32 }
func genFn0(p0 bool, p1 int64, p2 string) string {
    var v6 uint16 = 32767
    v6 += 2
    v6 = 255
    if (1.5 >= float64(0.1)) {
        fmt.Println(v6)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1(p0 int64) int64 {
    if true {
        var v7 bool = ((true && true) || false)
        _ = v7
    } else {
        p0 = (p0 & (p0 & -817))
    }
    v8 := []int{0, 32768, 127, 42}
    v9 := GenS1{F0: 1, F1: (127 > 42)}
    if (2 > 34) {
        switch "" {
        case "\n\t":
            for fi10 := 0; fi10 < 16; fi10++ {
                p0 += -857
                for fi11 := 0; fi11 < 5; fi11++ {
                    p0 = p0
                    if true {
                        if true {
                            p0 = (p0 >> ((42) & 63))
                            var v12 int64 = p0
                            _ = v12
                        } else {
                            p0 -= (p0 - (p0 * p0))
                            if (!false) {
                                var v13 int8 = -128
                                _ = v13
                            }
                            fmt.Println(fi11, fi10, fi11)
                        }
                        var v14 int = (fi11 + fi11)
                        var v15 uint64 = 42
                        _ = v14
                        _ = v15
                    } else {
                        var v16 float32 = -2.25
                        _ = v16
                    }
                    _ = fi11
                }
                p0 -= p0
                _ = fi10
            }
        case "a":
            p0 = (p0 ^ 742)
        case "hello":
            var v17 int32 = 1
            _ = v17
        default:
            if false {
                var v18 uint64 = 2
                var v19 int64 = p0
                _ = v18
                _ = v19
            } else {
                p0 += -575
                p0 = (p0 / (((p0 % (((p0 ^ p0)) | 1))) | 1))
            }
        }
        for fi20 := 0; fi20 < 2; fi20++ {
            p0 += 511
            var v21 string = ("\n\t" + "")
            var v22 int8 = -128
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 float64 = float64(3.14159)
        _ = v23
    } else {
        v24 := GenS1{F0: 18446744073709551615, F1: false}
        _ = v24
    }
    _ = p0
    _ = v8
    _ = v9
    return (p0 ^ p0)
}

func main() {
    var v1 int8 = 116
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v4 int8 = 0
    v5 := make(map[int]uint32)
    v4 = (v1 - v4)
    v25 := GenS2{F0: ((0 < 18446744073709551615) || false), F1: 9223372036854775807, F2: -607}
    if (-554 >= genFn1(-487)) {
        var v26 uint16 = 255
        var v27 int64 = genFn1(-9223372036854775808)
        _ = v26
        _ = v27
    } else {
        v28 := v3[9223372036854775807]
        v29 := GenS1{F0: 936, F1: false}
        v30 := GenS1{F0: 2, F1: true}
        _ = v28
        _ = v29
        _ = v30
    }
    var v31 uint32 = 2
    v4 = 127
    v31 -= v31
    v31 = v31
    switch 9223372036854775807 {
    case -87:
        var v32 int32 = 1
        _ = v32
    case 42:
        if false {
            switch "\n\t" {
            case "hello":
                v33 := []int{}
                _ = v33
            case "":
                v34 := GenS1{F0: 435, F1: (false || (true || false))}
                _ = v34
            }
            for fi35 := 0; fi35 < 6; fi35++ {
                _ = len(v2)
                var v36 int64 = 9223372036854775807
                _ = fi35
                _ = v36
            }
            v37 := []int{4294967295}
            _ = v37
        } else {
            v4 += 0
            var v38 bool = true
            _ = v38
        }
    default:
        for fi39 := 0; fi39 < 9; fi39++ {
            var v40 uint = 18446744073709551615
            v4 = v4
            var v41 int16 = 734
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    v1 -= (v1 + (v1 % ((v4) | 1)))
    var v42 uint32 = (^(v31))
    v42 -= v31
    v4 = (v1 + (v1 >> ((2) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v42
    _ = v5
}
