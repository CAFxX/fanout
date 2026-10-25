package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 bool; F2 uint }
func genFn0(p0 int32, p1 int64, p2 uint32) uint8 {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int64 = p1
        if (false && true) {
            if (1000 != 32767) {
                v8 -= -9223372036854775808
            }
        } else {
            v9 := []int{}
            _ = v9
        }
        fmt.Println(v8, p0, p2)
        _ = fi7
        _ = v8
    }
    p2 -= (p2 & (p2 >> ((0) & 63)))
    switch 436 {
    case -9223372036854775808:
        fmt.Println(p0)
    case 128:
        var v10 int32 = p0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 9
}
func genFn1() int16 {
    switch 128 {
    case -9223372036854775808:
        var v11 uint64 = 4294967295
        _ = v11
    }
    switch 502 {
    case -2147483648:
        v12 := []int{603, 128}
        _ = v12
    default:
        var v13 uint64 = 0
        _ = v13
    }
    var v14 int32 = -2147483648
    _ = v14
    return 0
}

func main() {
    var v1 int16 = -32768
    var v2 float64 = -2.25
    var v3 int64 = 42
    v4 := make(map[string]bool)
    var v5 int16 = v1
    var v6 int8 = -128
    var v15 uint32 = 2147483648
    if true {
        var v16 float64 = 3.14159
        if true {
            var v17 float64 = v16
            switch -580 {
            case 4294967295:
                v18 := GenS1{F0: false}
                _ = v18
            case 9223372036854775807:
                var v19 uint = 42
                _ = v19
            case -9223372036854775808:
                v20 := GenS1{F0: false}
                _ = v20
            default:
                v15 += v15
            }
            _ = v17
        } else {
            v21 := []int{32768, int(v5), 4294967295}
            switch 4294967295 {
            case -103:
                fmt.Println(v15)
            case 42:
                switch int(v5) {
                case -2147483648:
                    fmt.Println(v2, v1)
                case -518:
                    v22 := GenS1{F0: ((143 >= -753) && true)}
                    _ = v22
                case 300:
                    v23 := GenS1{F0: (false || true)}
                    _ = v23
                default:
                    v2 += v16
                }
            default:
                for fi24 := 0; fi24 < 8; fi24++ {
                    v25 := GenS2{F0: 703, F1: 423}
                    _ = fi24
                    _ = v25
                }
            }
            if len(v21) > 0 {
                v21[0] = -145
            }
            _ = v21
        }
        v4[(((("a" + "\n\t") + "") + "a") + "")] = true
        _ = v16
    }
    if (!true) {
        v26 := GenS3{F0: 2, F1: false}
        if true {
            var v27 float64 = (0.1 / v2)
            if true {
                v28 := GenS2{F0: 160, F1: v3}
                _ = v28
            } else {
                fmt.Println(v1)
            }
            var v29 int8 = 127
            _ = v27
            _ = v29
        }
        v15 = 641
        _ = v26
    } else {
        v1 = v5
        var v30 uint = 4294967295
        _ = v30
    }
    var v31 bool = (42 < 255)
    if v31 {
        fmt.Println(v1, v15)
    }
    v32 := GenS4{F0: 18446744073709551615}
    v33 := []int{-250, 316, 9223372036854775807}
    var v34 int64 = (v3 ^ 0)
    if (((3.14159 < ((1e10 * -2.25) / 1.0)) || (255 == 255)) && (1 < 0)) {
        if (0.1 >= v2) {
            for fi35 := 0; fi35 < 17; fi35++ {
                if false {
                    v36 := GenS1{F0: true}
                    _ = v36
                }
                var v37 uint16 = 255
                v38 := GenS5{F0: (v34 / ((v3) | 1)), F1: v31, F2: 335}
                _ = fi35
                _ = v37
                _ = v38
            }
            v39 := v4[""]
            _ = v39
        } else {
            _ = len(v33)
        }
    } else {
        if v31 {
            if ((!(224 <= 1)) && v31) {
                v34 = 1
            }
            for fi40 := 0; fi40 < 7; fi40++ {
                delete(v4, ("\n\t" + "a"))
                _ = fi40
            }
        } else {
            v31 = v31
            v33 = append(v33, -2147483648)
        }
        if false {
            var v41 int64 = (v34 / ((v3) | 1))
            if (!v31) {
                v42 := GenS1{F0: v31}
                _ = v42
            } else {
                var v43 float64 = (((float64(-2.25) - (1e10 - 0.0)) * 1e10) - ((((v2 * v2) / (v2 + (float64(0.0) / 0.1))) + 0.1) / (((v2 - v2) + float64(((1e10 * 0.0) * 1e10))) * (v2 + (v2 + v2)))))
                v44 := v4["foo bar"]
                _ = v43
                _ = v44
            }
            _ = v41
        }
        v34 += (v3 - v34)
    }
    v6 = (v6 % ((v6) | 1))
    v45 := GenS3{F0: 634, F1: (2 < 0)}
    v1 += -262
    var v46 float64 = float64(0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
