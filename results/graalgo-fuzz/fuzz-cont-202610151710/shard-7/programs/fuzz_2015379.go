package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 bool; F2 bool }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16 }
func genFn0() uint8 {
    var v7 int8 = 0
    v8 := GenS1{F0: "", F1: 476, F2: 84}
    fmt.Println(v7)
    v7 += v7
    var v9 uint32 = 120
    _ = v7
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 uint8, p1 uint8) bool {
    p1 = (p1 & 0)
    fmt.Println(p1, p1)
    v10 := GenS1{F0: "", F1: 1, F2: 4294967295}
    _ = p0
    _ = p1
    _ = v10
    return false
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    v3 := []uint32{0, 2147483647}
    v4 := make(map[string]int)
    v5 := []int{}
    var v6 int64 = -805
    if true {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int16 = 131
            _ = fi11
            _ = v12
        }
    } else {
        var v13 bool = true
        _ = v13
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            v6 = (v6 / ((v6) | 1))
            v16 := GenS1{F0: "", F1: 18446744073709551615, F2: 0}
            _ = fi15
            _ = v16
        }
        switch "foo bar" {
        case "a":
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := GenS1{F0: "\n\t", F1: 86, F2: 689}
                _ = fi17
                _ = v18
            }
        case "\n\t":
            fmt.Println(fi14, v6)
        case "foo bar":
            var v19 bool = (true && (!true))
            _ = v19
        }
        v6 += (v6 + v6)
        _ = fi14
    }
    v4["\n\t"] = 4294967295
    var v20 uint = 18446744073709551615
    switch -982 {
    case 9223372036854775807:
        for fi21 := 0; fi21 < 6; fi21++ {
            v6 += -9223372036854775808
            v22 := GenS2{F0: 255, F1: true, F2: (131 >= 76)}
            var v23 float32 = (0.1 * -2.25)
            _ = fi21
            _ = v22
            _ = v23
        }
    case 32767:
        if ((((-1 > -128) || false) && (false || true)) || (0 <= v6)) {
            if (!false) {
                v24 := GenS2{F0: 42, F1: (0.0 > 0.0), F2: false}
                fmt.Println(v20, v20)
                v25 := GenS3{F0: 378, F1: 1}
                _ = v24
                _ = v25
            } else {
                if (!true) {
                    var v26 int32 = 42
                    var v27 int = 0
                    _ = v26
                    _ = v27
                } else {
                    var v28 bool = (false || false)
                    var v29 bool = true
                    _ = v28
                    _ = v29
                }
                switch 2147483647 {
                case 2147483647:
                    v30 := GenS1{F0: ("hello" + ("a" + "a")), F1: 2, F2: 0}
                    _ = v30
                case -955:
                    v6 = (v6 % (((-(v6))) | 1))
                case -193:
                    var v31 uint64 = 2
                    _ = v31
                default:
                    fmt.Println(v20, v20)
                }
            }
            v6 = v6
        }
    default:
        v3 = append(v3, 0)
    }
    v32 := GenS4{F0: 42}
    for fi33 := 0; fi33 < 11; fi33++ {
        v6 += v6
        if false {
            if true {
                v34 := GenS5{F0: 2}
                if (((0.1 / 1.0) / 1.5) == (float64(0.1) * 0.1)) {
                    var v35 int64 = v6
                    _ = v35
                }
                var v36 bool = ((true || false) || true)
                _ = v34
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 9; fi37++ {
            v20 -= (v20 % (((v20 ^ v20)) | 1))
            var v38 uint32 = 42
            _ = fi37
            _ = v38
        }
        _ = fi33
    }
    v20 = (v20 | (v20 << (((^(v20))) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
