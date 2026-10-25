package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int64 }
type GenS2 struct { F0 uint; F1 int64; F2 string }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int; F1 uint32 }
func genFn0(p0 string, p1 uint16, p2 uint32) bool {
    fmt.Println(p1, p2)
    var v7 uint64 = 862
    v8 := []int{32767, 9223372036854775807, 584, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ((^(v7)) >= 4294967295)
}
func genFn1(p0 uint16) uint64 {
    for fi9 := 0; fi9 < 10; fi9++ {
        if false {
            v10 := []int{}
            if false {
                var v11 uint8 = 159
                if true {
                    var v12 uint16 = (p0 | p0)
                    _ = v12
                }
                _ = v11
            } else {
                var v13 bool = (!false)
                _ = v13
            }
            v10 = append(v10, (fi9 ^ fi9))
            _ = v10
        }
        _ = fi9
    }
    fmt.Println(p0)
    _ = p0
    return 137
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{648, 1, 439, 2147483647}
    v3 := make(map[int]int)
    var v4 float64 = ((((1e10 - 0.0) + float64(0.1)) * float64(((0.1 / float32(0.1)) - 1.5))) + float64((((((float32(1.5) - 1e10) - float32(0.1)) * float32(float64(((0.0 * ((1e10 + (0.0 / 1e10)) / 1.0)) / (1.5 - 0.1))))) / 1.0) * 0.0)))
    var v5 int8 = 66
    v6 := make(map[string]uint32)
    if (!true) {
        v14 := GenS1{F0: 65535, F1: -32768, F2: -9223372036854775808}
        for fi15 := 0; fi15 < 7; fi15++ {
            v5 += (v5 % ((v5) | 1))
            var v16 uint = 1
            v17 := GenS1{F0: 2147483648, F1: -689, F2: 42}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    }
    var v18 float64 = float64((1.5 - -2.25))
    switch 1 {
    case 32767:
        switch 1 {
        case 1:
            var v19 uint64 = 1
            _ = v19
        case -201:
            v5 = (-(v5))
        case 127:
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 int16 = 1000
                v4 = v4
                switch ("hello" + ("hello" + "foo bar")) {
                case "foo bar":
                    for fi22 := 0; fi22 < 1; fi22++ {
                        var v23 uint64 = 2
                        _ = fi22
                        _ = v23
                    }
                case "\n\t":
                    if false {
                        v4 += float64((((3.14159 * float32((float64(0.1) * 3.14159))) + (1e10 * 1e10)) + 1.5))
                        var v24 int32 = -1
                        v2 = append(v2, 1)
                        _ = v24
                    }
                }
                _ = fi20
                _ = v21
            }
        default:
            var v25 int64 = 9223372036854775807
            _ = v25
        }
    default:
        v26 := GenS2{F0: 18446744073709551615, F1: -1, F2: "a"}
        _ = v26
    }
    switch ("hello" + "foo bar") {
    case "\n\t":
        v4 = v18
    default:
        fmt.Println(v18)
    }
    switch int(v5) {
    case -101:
        var v27 uint64 = 2
        _ = v27
    case -1:
        var v28 uint8 = 76
        _ = v28
    case 32768:
        if (-32768 == 0) {
            var v29 string = "hello"
            if false {
                v30 := GenS2{F0: 65535, F1: -1, F2: v29}
                v31 := GenS3{F0: v5, F1: 1}
                switch v29 {
                case "\n\t":
                    var v32 int64 = 352
                    _ = v32
                }
                _ = v30
                _ = v31
            } else {
                var v33 float32 = float32((1e10 / (float64(3.14159) / (v18 - 0.0))))
                for fi34 := 0; fi34 < 17; fi34++ {
                    var v35 uint64 = 2
                    if ((-1 == -128) || (!true)) {
                        var v36 int64 = 44
                        _ = v36
                    } else {
                        v18 = -2.25
                        v33 -= v33
                        var v37 float64 = ((0.0 - 1.5) / (float64(v33) - 0.1))
                        _ = v37
                    }
                    switch (fi34 ^ (fi34 - fi34)) {
                    case 896:
                        if ((fi34 * fi34) == 102) {
                            v38 := GenS2{F0: 0, F1: 9223372036854775807, F2: v29}
                            for fi39 := 0; fi39 < 12; fi39++ {
                                v40 := GenS4{F0: (fi34 / (((fi34 / ((128) | 1))) | 1)), F1: 2147483647}
                                switch (-(fi39)) {
                                case -448:
                                    switch ("foo bar" + "hello") {
                                    case "foo bar":
                                        v5 += 127
                                    default:
                                        if true {
                                            v4 -= float64(((v33 + float32(v18)) - v33))
                                            v29 = (v29 + v29)
                                            var v41 int = fi34
                                            _ = v41
                                        }
                                    }
                                case 430:
                                    v18 = ((v4 + v4) - (v4 - v4))
                                }
                                _ = fi39
                                _ = v40
                            }
                            fmt.Println(v33, v33)
                            _ = v38
                        }
                    case 127:
                        var v42 int = (-(fi34))
                        _ = v42
                    default:
                        v18 = v4
                    }
                    _ = fi34
                    _ = v35
                }
                _ = v33
            }
            v2 = append(v2, 0)
            _ = v29
        } else {
            var v43 uint8 = 0
            if (!((false && false) || false)) {
                for fi44 := 0; fi44 < 15; fi44++ {
                    var v45 uint16 = 65535
                    var v46 int16 = 1000
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                v47 := GenS4{F0: 128, F1: 2}
                _ = v47
            }
            _ = v43
        }
    default:
        fmt.Println(v18)
    }
    v18 = v18
    var v48 int64 = 42
    fmt.Println(v48, v48)
    var v49 int32 = 374
    for fi50 := 0; fi50 < 7; fi50++ {
        v18 -= (float64(((-2.25 * (0.0 / 1e10)) + 0.0)) * (v4 + (v18 * v18)))
        _ = fi50
    }
    var v51 uint16 = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v51
    _ = v6
}
