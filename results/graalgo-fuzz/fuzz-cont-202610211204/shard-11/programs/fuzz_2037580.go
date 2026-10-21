package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int32; F2 uint16 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int32, p1 int8, p2 int) int32 {
    var v6 int16 = 565
    var v7 int8 = 19
    if true {
        v8 := GenS1{F0: 9223372036854775807, F1: (p0 >> ((42) & 63))}
        v9 := []int{(p2 >> ((42) & 63)), p2}
        var v10 int = 4294967295
        _ = v10
        _ = v8
        _ = v9
    } else {
        if false {
            fmt.Println(p0, p0)
            for fi11 := 0; fi11 < 9; fi11++ {
                v6 = v6
                v6 -= v6
                _ = fi11
            }
            v6 -= (-(v6))
        }
        v12 := []int{(p2 ^ p2), p2, p2, p2}
        _ = v12
    }
    p2 += p2
    var v13 int16 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 float32 = 3.14159
    v3 := make(map[uint32]uint32)
    var v4 int8 = 100
    var v5 uint = v1
    fmt.Println(v5, v5)
    v1 -= 18446744073709551615
    fmt.Println(v4, v2)
    if (715 > 1) {
        v3[1] = 0
    } else {
        switch "a" {
        case "":
            v3[42] = 653
        default:
            var v14 uint16 = 42
            _ = v14
        }
        v5 -= v1
    }
    fmt.Println(v4)
    var v15 bool = false
    var v16 int32 = genFn0(genFn0(2147483647, v4, 0), (-(v4)), 2147483647)
    if (true && v15) {
        v17 := GenS2{F0: (float32(float64((float32((1e10 - 0.1)) - v2))) < float32(-2.25))}
        if (((0.1 - 1e10) * 1.5) == (0.0 / (-2.25 / 1.5))) {
            v15 = (!((97 >= 484) && (42 < 255)))
            v18 := GenS1{F0: 1, F1: v16}
            if (1e10 == ((-2.25 * 0.0) * 1e10)) {
                for fi19 := 0; fi19 < 2; fi19++ {
                    if false {
                        var v20 uint32 = 1
                        var v21 string = "hello"
                        _ = v20
                        _ = v21
                    } else {
                        v16 += (^(v16))
                        var v22 float64 = ((1.5 * 0.0) / float64(v2))
                        v16 = (v16 ^ int32(v4))
                        _ = v22
                    }
                    _ = fi19
                }
                switch "" {
                case "foo bar":
                    fmt.Println(v16, v16, v15)
                case "a":
                    switch "" {
                    case "":
                        v23 := GenS2{F0: (uint8(v5) <= 1)}
                        _ = v23
                    case "hello":
                        v4 = 64
                    case "\n\t":
                        v16 -= v16
                    default:
                        v24 := []int{-9223372036854775808, 0, -117, 464}
                        _ = v24
                    }
                default:
                    var v25 string = "\n\t"
                    _ = v25
                }
            } else {
                var v26 int = 32767
                _ = v26
            }
            _ = v18
        } else {
            for fi27 := 0; fi27 < 10; fi27++ {
                v5 = v1
                _ = fi27
            }
        }
        _ = v17
    }
    if false {
        v2 += v2
        switch ("" + "a") {
        case "\n\t":
            v28 := GenS1{F0: 1, F1: v16}
            _ = v28
        case "hello":
            fmt.Println(v2)
        }
    } else {
        fmt.Println(v2, v16)
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 int64 = -275
            for fi31 := 0; fi31 < 8; fi31++ {
                v1 = (v1 / ((v5) | 1))
                if true {
                    var v32 uint16 = 2
                    _ = v32
                } else {
                    if (508 < 1) {
                        v33 := GenS3{F0: v30, F1: v16, F2: 32767}
                        v15 = false
                        v2 = float32((((0.0 * 3.14159) / (1.5 * (0.0 + (float64(v2) * 1.5)))) / ((3.14159 / ((1e10 * 1.5) / (3.14159 / (1.5 + 1.5)))) / 1.0)))
                        _ = v33
                    } else {
                        var v34 string = "\n\t"
                        switch (fi29 | fi31) {
                        case 0:
                            v35 := GenS4{F0: genFn0((v16 - (v16 ^ (v16 & genFn0((v16 | v16), -128, fi31)))), (v4 << ((v1) & 63)), (-(fi31)))}
                            _ = v35
                        case -475:
                            v5 += (v1 + (v1 - v5))
                        case 9223372036854775807:
                            fmt.Println(v34, v34, fi31)
                        default:
                            v36 := []int{128, 663}
                            _ = v36
                        }
                        delete(v3, 354)
                        _ = v34
                    }
                    fmt.Println(fi29)
                }
                _ = fi31
            }
            var v37 int16 = int16(v4)
            _ = fi29
            _ = v30
            _ = v37
        }
        var v38 int8 = v4
        _ = v38
    }
    switch ("foo bar" + "a") {
    case "foo bar":
        v4 = -50
    }
    v2 += v2
    switch 127 {
    case -1:
        v39 := GenS4{F0: (v16 >> (((v1 + v5)) & 63))}
        _ = v39
    default:
        var v40 float64 = 1e10
        _ = v40
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        for fi42 := 0; fi42 < 19; fi42++ {
            var v43 uint16 = 58
            var v44 int32 = genFn0((v16 * (v16 * (v16 ^ (v16 >> ((v1) & 63))))), (^(v4)), fi41)
            var v45 float64 = 1e10
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
        if (-1 <= (fi41 | (fi41 / ((fi41) | 1)))) {
            var v46 string = "foo bar"
            _ = v46
        }
        _ = fi41
    }
    var v47 float64 = 1e10
    for fi48 := 0; fi48 < 1; fi48++ {
        var v49 int8 = v4
        _ = fi48
        _ = v49
    }
    v16 += 0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
