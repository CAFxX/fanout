package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint; F1 int; F2 uint }
type GenS4 struct { F0 int; F1 uint16; F2 int32 }
type GenS5 struct { F0 int; F1 uint16; F2 bool }
type GenS6 struct { F0 uint8; F1 int; F2 int32 }
type GenS7 struct { F0 int32; F1 uint16 }
func genFn0(p0 string, p1 int32, p2 bool) bool {
    switch 9223372036854775807 {
    case 42:
        fmt.Println(p1, p1)
    case -1:
        fmt.Println(p0)
    }
    v5 := GenS1{F0: 142}
    v6 := GenS1{F0: 128}
    for fi7 := 0; fi7 < 2; fi7++ {
        switch (p0 + "") {
        case "\n\t":
            if false {
                if ((966 > 42) && false) {
                    v8 := GenS2{F0: 32767, F1: 61}
                    var v9 int32 = p1
                    _ = v8
                    _ = v9
                }
                v10 := []int{-106, (fi7 | (-(fi7))), (fi7 % ((fi7) | 1)), (fi7 * (fi7 ^ -1))}
                switch "a" {
                case "":
                    v10 = append(v10, (fi7 - 127))
                default:
                    var v11 string = ((p0 + ((((("\n\t" + "") + (p0 + "hello")) + p0) + "\n\t") + p0)) + (p0 + p0))
                    _ = v11
                }
                _ = v10
            }
        case "hello":
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 float32 = (1e10 + 1e10)
                var v14 int32 = -486
                _ = fi12
                _ = v13
                _ = v14
            }
        case "a":
            switch "\n\t" {
            case "hello":
                if p2 {
                    p2 = p2
                }
            case "foo bar":
                var v15 float32 = ((((((3.14159 * ((0.1 + (3.14159 / -2.25)) - (1.5 / -2.25))) + (0.0 * 0.0)) + -2.25) + 3.14159) + (0.0 + 3.14159)) / 3.14159)
                _ = v15
            case "\n\t":
                p2 = (float64(3.14159) < float64(3.14159))
            default:
                var v16 float32 = 3.14159
                _ = v16
            }
        default:
            for fi17 := 0; fi17 < 18; fi17++ {
                p2 = true
                switch "" {
                case "":
                    var v18 int = fi17
                    _ = v18
                case "foo bar":
                    fmt.Println(fi17, p0)
                case "\n\t":
                    fmt.Println(fi17, p1, p0)
                default:
                    if (!p2) {
                        v19 := []int{fi17, fi17}
                        p0 = (("" + p0) + p0)
                        _ = v19
                    } else {
                        v20 := GenS3{F0: 18446744073709551615, F1: fi7, F2: 2}
                        p1 += -2147483648
                        v21 := GenS1{F0: 255}
                        _ = v20
                        _ = v21
                    }
                }
                _ = fi17
            }
        }
        v22 := GenS2{F0: 1, F1: 1}
        for fi23 := 0; fi23 < 9; fi23++ {
            v24 := GenS4{F0: fi23, F1: 250, F2: p1}
            if (100 >= 88) {
                v25 := GenS4{F0: 1, F1: 2, F2: 42}
                fmt.Println(fi7, fi7)
                switch "" {
                case "":
                    var v26 bool = ((0.0 * 0.0) >= float64(-2.25))
                    _ = v26
                case "hello":
                    p2 = (86 == 42)
                }
                _ = v25
            } else {
                var v27 float32 = 0.0
                _ = v27
            }
            _ = fi23
            _ = v24
        }
        _ = fi7
        _ = v22
    }
    var v28 float64 = (0.1 - (-2.25 + 1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v5
    _ = v6
    return (4294967295 > 255)
}

func main() {
    var v1 int8 = -1
    var v2 int32 = -699
    var v3 uint16 = 2
    v4 := []int{}
    var v29 uint64 = 1
    if len(v4) > 0 {
        v4[0] = 1
    }
    v4 = append(v4, 4294967295)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        for fi31 := 0; fi31 < 15; fi31++ {
            if len(v4) > 0 {
                v4[0] = 855
            }
            _ = fi31
        }
        v32 := GenS1{F0: 208}
        _ = fi30
        _ = v32
    }
    _ = len(v4)
    fmt.Println(v1, v3)
    v4 = append(v4, 122)
    for fi33 := 0; fi33 < 2; fi33++ {
        _ = len(v4)
        if false {
            if (0.0 == (3.14159 / (1.5 - float32(0.0)))) {
                var v34 int64 = -43
                v35 := GenS5{F0: fi33, F1: (v3 + v3), F2: false}
                for fi36 := 0; fi36 < 7; fi36++ {
                    switch (^(fi36)) {
                    case 32767:
                        var v37 uint32 = 1
                        _ = v37
                    case 997:
                        switch (fi36 >> ((0) & 63)) {
                        case -428:
                            v1 += (v1 * v1)
                        case -9223372036854775808:
                            v38 := GenS2{F0: 1000, F1: 1}
                            _ = v38
                        }
                    }
                    _ = fi36
                }
                _ = v34
                _ = v35
            }
            fmt.Println(v2, v1, v2)
        }
        _ = fi33
    }
    fmt.Println(v2, v1)
    for fi39 := 0; fi39 < 9; fi39++ {
        v29 += 622
        if false {
            var v40 uint16 = (v3 | v3)
            var v41 int8 = v1
            _ = v40
            _ = v41
        }
        v42 := GenS6{F0: 127, F1: 32767, F2: (v2 ^ (-(v2)))}
        _ = fi39
        _ = v42
    }
    fmt.Println(v29, v29)
    v2 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v43 := GenS7{F0: v2, F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v43
}
