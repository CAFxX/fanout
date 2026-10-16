package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 bool }
func genFn0() int {
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 string = v7
            _ = fi9
            _ = v10
        }
        switch v7 {
        case "a":
            v11 := []int{}
            _ = v11
        default:
            var v12 int8 = -128
            _ = v12
        }
        var v13 float32 = (((1e10 - 3.14159) * 1.5) + (3.14159 / (float32(1e10) + 0.1)))
        _ = fi8
        _ = v13
    }
    _ = v7
    return 127
}
func genFn1(p0 int16) uint8 {
    if true {
        for fi14 := 0; fi14 < 5; fi14++ {
            p0 = p0
            if (!false) {
                var v15 uint32 = 2147483647
                for fi16 := 0; fi16 < 7; fi16++ {
                    if (637 < 42) {
                        var v17 uint32 = 42
                        fmt.Println(fi14, p0)
                        var v18 int32 = -2147483648
                        _ = v17
                        _ = v18
                    }
                    _ = fi16
                }
                _ = v15
            }
            _ = fi14
        }
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 int8 = -1
            _ = fi19
            _ = v20
        }
        if (1.5 < float64((3.14159 - 1.5))) {
            var v21 uint = 255
            v21 += 456
            v21 = (v21 | (v21 | (^(v21))))
            _ = v21
        } else {
            p0 += (p0 ^ 32767)
            fmt.Println(p0)
        }
    } else {
        v22 := []int{2147483647}
        _ = v22
    }
    p0 = (-(p0))
    if (false || (-1 > -128)) {
        p0 -= p0
    } else {
        var v23 int32 = -1
        _ = v23
    }
    p0 = (p0 / ((p0) | 1))
    var v24 bool = (!false)
    _ = p0
    _ = v24
    return 2
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    v4 := []string{("" + (("a" + "a") + "foo bar")), ("a" + ("" + "hello"))}
    var v5 int32 = 2147483647
    var v6 int8 = 0
    if v1 {
        v5 = (v5 ^ (v5 | v5))
    }
    if v1 {
        if v1 {
            v25 := GenS1{F0: 42, F1: -9223372036854775808, F2: 188}
            _ = v25
        } else {
            v6 -= v6
            var v26 bool = v1
            v27 := GenS1{F0: 2, F1: genFn0(), F2: 261}
            _ = v26
            _ = v27
        }
    }
    if v1 {
        if (v5 > (v5 ^ -2147483648)) {
            fmt.Println(v6)
            var v28 uint64 = 2
            v29 := GenS1{F0: 65535, F1: 150, F2: -1}
            _ = v28
            _ = v29
        }
        var v30 uint32 = 925
        var v31 int16 = -478
        _ = v30
        _ = v31
    }
    v1 = v1
    v6 -= (^(v6))
    for fi32 := 0; fi32 < 19; fi32++ {
        switch "" {
        case "hello":
            if v1 {
                for fi33 := 0; fi33 < 7; fi33++ {
                    switch fi32 {
                    case 128:
                        var v34 bool = v1
                        _ = v34
                    case -2147483648:
                        v35 := GenS2{F0: 2147483647, F1: v6, F2: v1}
                        _ = v35
                    case -118:
                        var v36 string = ""
                        _ = v36
                    default:
                        var v37 int16 = 1000
                        _ = v37
                    }
                    _ = fi33
                }
                v38 := GenS1{F0: 245, F1: 0, F2: -1}
                if ((2 <= 0) && false) {
                    fmt.Println(fi32, fi32, v6)
                    for fi39 := 0; fi39 < 15; fi39++ {
                        fmt.Println(v1, fi32, fi32)
                        if ((v6 >> ((255) & 63)) != v6) {
                            v1 = v1
                        }
                        _ = fi39
                    }
                } else {
                    v5 -= -2147483648
                    fmt.Println(fi32)
                }
                _ = v38
            } else {
                v40 := GenS2{F0: -572, F1: v6, F2: v1}
                v1 = (!(-489 <= 42))
                for fi41 := 0; fi41 < 11; fi41++ {
                    var v42 float32 = 3.14159
                    v43 := GenS2{F0: int32(v6), F1: -128, F2: (v1 && v1)}
                    v6 = 127
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = v40
            }
        default:
            if v1 {
                delete(v3, 2147483648)
            } else {
                v4 = append(v4, ("hello" + "hello"))
                v5 = (v5 + (v5 | (^(v5))))
            }
        }
        for fi44 := 0; fi44 < 16; fi44++ {
            fmt.Println(v6)
            _ = fi44
        }
        _ = fi32
    }
    switch 4294967295 {
    case -593:
        if ((v5 - 2147483647) <= v5) {
            for fi45 := 0; fi45 < 12; fi45++ {
                var v46 int64 = -1
                _ = fi45
                _ = v46
            }
        } else {
            v47 := GenS1{F0: 548, F1: -645, F2: 277}
            switch 32767 {
            case 32768:
                fmt.Println(v1)
            case 127:
                switch "a" {
                case "foo bar":
                    v1 = v1
                default:
                    var v48 float32 = 3.14159
                    _ = v48
                }
            }
            v1 = (false || (v1 || (743 < -1)))
            _ = v47
        }
    case 4294967295:
        var v49 int64 = 559
        _ = v49
    case 127:
        fmt.Println(v6, v6)
    }
    var v50 int8 = (v6 & (v6 - v6))
    fmt.Println(v1, v50)
    v51 := GenS2{F0: v5, F1: 20, F2: false}
    v6 = v50
    fmt.Println(v1, v1, v50)
    fmt.Println(v6, v1, v50)
    var v52 string = "foo bar"
    v52 = "foo bar"
    if (526 == 42) {
        _ = len(v2)
        fmt.Println(v6)
    }
    var v53 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
