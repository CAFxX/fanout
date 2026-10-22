package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint; F2 uint32 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint32, p1 int8) int16 {
    var v5 uint = 0
    switch 127 {
    case -1:
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := []int{}
            if (32767 <= 755) {
                v8 := GenS1{F0: 2, F1: (((p1 & (p1 ^ -34)) == -67) || true)}
                v5 += 802
                _ = v8
            }
            _ = fi6
            _ = v7
        }
    default:
        for fi9 := 0; fi9 < 4; fi9++ {
            p0 -= p0
            v10 := []int{4294967295}
            switch "" {
            case "hello":
                switch "foo bar" {
                case "hello":
                    p0 += (p0 * 4294967295)
                case "":
                    v5 -= v5
                case "a":
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                default:
                    fmt.Println(fi9)
                }
            case "foo bar":
                var v11 uint32 = (p0 * p0)
                _ = v11
            case "":
                for fi12 := 0; fi12 < 14; fi12++ {
                    if false {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        var v13 float32 = 3.14159
                        switch (fi12 * -137) {
                        case 4294967295:
                            switch "a" {
                            case "\n\t":
                                switch fi12 {
                                case 684:
                                    var v14 int8 = int8(v5)
                                    _ = v14
                                }
                            case "a":
                                switch ("" + "hello") {
                                case "foo bar":
                                    p0 = (p0 + p0)
                                }
                            case "hello":
                                v10 = append(v10, (fi12 >> ((4294967295) & 63)))
                            }
                        case 32767:
                            var v15 bool = false
                            _ = v15
                        default:
                            fmt.Println(p0, fi9, v5)
                        }
                        _ = v13
                    }
                    v16 := GenS2{F0: 42}
                    p1 = (^(p1))
                    _ = fi12
                    _ = v16
                }
            default:
                v17 := GenS3{F0: fi9}
                _ = v17
            }
            _ = fi9
            _ = v10
        }
    }
    if true {
        switch "\n\t" {
        case "\n\t":
            var v18 int64 = -9223372036854775808
            _ = v18
        case "foo bar":
            p0 -= 4294967295
        case "hello":
            v19 := GenS3{F0: 128}
            _ = v19
        }
    }
    switch "\n\t" {
    case "foo bar":
        if (85 == 42) {
            v20 := GenS2{F0: (v5 << ((v5) & 63))}
            _ = v20
        }
    case "hello":
        for fi21 := 0; fi21 < 6; fi21++ {
            switch 128 {
            case 1:
                v22 := GenS2{F0: 4294967295}
                _ = v22
            case 972:
                var v23 int8 = (p1 | p1)
                _ = v23
            case 862:
                for fi24 := 0; fi24 < 12; fi24++ {
                    p0 -= (p0 ^ p0)
                    switch ("hello" + "a") {
                    case "hello":
                        v25 := []int{127, (fi21 << ((v5) & 63)), 899, fi24}
                        _ = v25
                    case "a":
                        p0 += 181
                    }
                    var v26 uint8 = 127
                    _ = fi24
                    _ = v26
                }
            default:
                fmt.Println(p1, v5)
            }
            if true {
                v27 := []int{-9223372036854775808, (-(fi21)), fi21}
                v5 += v5
                _ = v27
            } else {
                if (((float64(1.5) + float64((((1.5 / (-2.25 + (-2.25 / 1.5))) - 0.1) / -2.25))) - 1e10) < -2.25) {
                    var v28 int = 9223372036854775807
                    p0 -= p0
                    var v29 int32 = int32(v5)
                    _ = v28
                    _ = v29
                } else {
                    v5 = v5
                    v5 += (v5 * 611)
                    v30 := []int{(-(fi21)), fi21, (^(fi21)), fi21}
                    _ = v30
                }
                switch fi21 {
                case 32768:
                    if true {
                        var v31 float32 = (1.5 / (float32((float64(float32(0.0)) / (3.14159 - (1e10 * (1.5 + (-2.25 + 0.1)))))) - 1e10))
                        v31 = ((v31 - (((1e10 - 3.14159) - v31) - (float32(1e10) / v31))) - 1.5)
                        _ = v31
                    } else {
                        switch "hello" {
                        case "\n\t":
                            v32 := GenS3{F0: (^(fi21))}
                            _ = v32
                        default:
                            v33 := []int{}
                            _ = v33
                        }
                    }
                case 740:
                    p0 -= (p0 ^ (p0 / ((p0) | 1)))
                default:
                    v34 := []int{128, fi21}
                    _ = v34
                }
            }
            var v35 float32 = 3.14159
            _ = fi21
            _ = v35
        }
    case "a":
        fmt.Println(p1, v5)
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    v1 := []int{}
    v2 := []int{-661}
    var v3 int64 = 1
    var v4 uint32 = 71
    v3 += -9223372036854775808
    v36 := GenS3{F0: 0}
    v37 := GenS4{F0: -22, F1: 42, F2: 2147483648}
    fmt.Println(v3)
    var v38 float64 = (1e10 + (1.5 * (0.0 - float64(0.1))))
    var v39 float32 = -2.25
    var v40 int32 = 952
    for fi41 := 0; fi41 < 18; fi41++ {
        if false {
            var v42 int16 = -1
            v43 := GenS2{F0: 556}
            var v44 int32 = (v40 / ((-2147483648) | 1))
            _ = v42
            _ = v43
            _ = v44
        } else {
            switch "foo bar" {
            case "a":
                var v45 string = (("hello" + "foo bar") + "\n\t")
                _ = v45
            case "hello":
                if true {
                    v46 := GenS5{F0: false}
                    v39 -= (1.5 - ((1e10 - (float32(3.14159) / v39)) + float32(v38)))
                    v38 += float64((v39 / -2.25))
                    _ = v46
                } else {
                    var v47 float64 = (((((0.0 - (v38 + v38)) - float64(-2.25)) / v38) / v38) / float64(v39))
                    _ = v47
                }
            case "":
                v2 = append(v2, 32768)
            default:
                for fi48 := 0; fi48 < 12; fi48++ {
                    v3 = v3
                    _ = fi48
                }
            }
            var v49 int16 = 32767
            _ = v49
        }
        _ = fi41
    }
    for fi50 := 0; fi50 < 20; fi50++ {
        var v51 float64 = v38
        switch fi50 {
        case -438:
            v39 = (1.5 * 1e10)
        }
        _ = fi50
        _ = v51
    }
    if len(v2) > 0 {
        v2[0] = 32768
    }
    if false {
        for fi52 := 0; fi52 < 9; fi52++ {
            var v53 int = (fi52 & 42)
            _ = fi52
            _ = v53
        }
    }
    v38 = (-2.25 + v38)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
