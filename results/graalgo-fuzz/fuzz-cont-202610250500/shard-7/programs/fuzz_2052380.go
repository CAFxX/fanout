package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int8 }
func genFn0(p0 string, p1 uint, p2 bool) int16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            if (!(p2 && ((^(p1)) > (p1 * p1)))) {
                p2 = p2
                p0 = p0
            } else {
                v9 := GenS1{F0: p1, F1: 0, F2: -1}
                if (255 >= 463) {
                    var v10 uint16 = 1
                    p2 = (42 >= 1)
                    v10 = (v10 | (v10 * (v10 ^ 473)))
                    _ = v10
                }
                if (!p2) {
                    var v11 int8 = 42
                    if (p0 <= p0) {
                        var v12 uint64 = 2
                        switch "foo bar" {
                        case "hello":
                            v13 := GenS2{F0: 861, F1: v11}
                            _ = v13
                        default:
                            var v14 uint = (p1 * (p1 | p1))
                            _ = v14
                        }
                        _ = v12
                    } else {
                        var v15 bool = false
                        _ = v15
                    }
                    p2 = p2
                    _ = v11
                } else {
                    p2 = p2
                }
                switch p0 {
                case "\n\t":
                    switch -298 {
                    case 9223372036854775807:
                        if false {
                            fmt.Println(p1, p2, fi8)
                        }
                    case 128:
                        p0 = "foo bar"
                        if p2 {
                            var v16 bool = p2
                            p2 = v16
                            _ = v16
                        }
                    case -2147483648:
                        var v17 float64 = (((3.14159 + 1.5) - ((-2.25 - float64(-2.25)) * (1.5 / (1.5 / (1.5 - float64((-2.25 / 1.0))))))) / float64(((((1e10 * (-2.25 * 0.1)) / float32(((((0.1 / (-2.25 / (3.14159 - 0.1))) * (float64((1.5 * (float32((float64(1e10) / 3.14159)) - 3.14159))) - 0.0)) + 3.14159) - 0.0))) + (0.0 / (3.14159 / (((1e10 / 1.0) + -2.25) * (1.5 * 1e10))))) + 0.1)))
                        _ = v17
                    }
                case "hello":
                    p0 = ("hello" + "")
                }
                p0 = (p0 + "a")
                _ = v9
            }
            switch fi8 {
            case -1:
                fmt.Println(fi8, p2)
            }
            v18 := GenS2{F0: 42, F1: 96}
            var v19 uint = p1
            p2 = p2
            fmt.Println(fi7)
            fmt.Println(fi8, v19, fi7)
            _ = fi8
            _ = v18
            _ = v19
        }
        switch fi7 {
        case 9223372036854775807:
            if (0 < 1) {
                p2 = p2
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 uint8 = 1
                    _ = fi20
                    _ = v21
                }
                var v22 uint16 = 363
                _ = v22
            }
        default:
            p0 = p0
        }
        fmt.Println(p1, fi7, p2)
        _ = fi7
    }
    if (0.0 > 0.1) {
        v23 := []int{306}
        v23 = append(v23, 127)
        _ = v23
    } else {
        fmt.Println(p2, p0)
        var v24 int8 = -106
        var v25 uint16 = 2
        _ = v24
        _ = v25
    }
    if p2 {
        switch 32768 {
        case 127:
            for fi26 := 0; fi26 < 3; fi26++ {
                if ((fi26 + (fi26 ^ (fi26 >> ((p1) & 63)))) == fi26) {
                    var v27 int8 = -1
                    _ = v27
                }
                _ = fi26
            }
        default:
            p0 = p0
        }
        v28 := GenS2{F0: 7, F1: 121}
        _ = v28
    }
    var v29 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = (1.5 - (0.0 - ((1e10 + 0.1) + 0.0)))
    var v3 int64 = 689
    var v4 uint16 = 1
    v5 := []string{"", "hello", ("a" + "a")}
    v6 := make(map[uint32]uint32)
    v6[2] = 88
    var v30 uint64 = 0
    v31 := GenS3{F0: v4, F1: uint16(v30), F2: -86}
    var v32 int32 = -574
    _ = len(v5)
    v30 = v30
    var v33 uint16 = v4
    if true {
        var v34 int8 = 0
        v35 := GenS3{F0: v33, F1: (v33 | v4), F2: (v34 & 1)}
        fmt.Println(v30, v4)
        _ = v34
        _ = v35
    } else {
        v33 -= v33
        var v36 int = -1
        _ = v36
    }
    v4 -= 142
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
