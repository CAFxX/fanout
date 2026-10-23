package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int32; F1 int32; F2 int }
type GenS5 struct { F0 uint8 }
func genFn0() int16 {
    v7 := GenS1{F0: -1}
    var v8 uint8 = 97
    v9 := []int{32768, 128}
    if false {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint16 = 255
            _ = fi10
            _ = v11
        }
    } else {
        if false {
            var v12 uint32 = 448
            v12 += v12
            _ = v12
        } else {
            fmt.Println(v8)
            switch "foo bar" {
            case "":
                v13 := GenS2{F0: v8, F1: 42, F2: 1}
                _ = v13
            case "\n\t":
                v14 := GenS2{F0: 42, F1: 1, F2: 1}
                _ = v14
            }
            var v15 int32 = 491
            _ = v15
        }
        v8 -= (v8 << ((703) & 63))
    }
    _ = v7
    _ = v8
    _ = v9
    return 925
}
func genFn1(p0 int32, p1 string) uint16 {
    for fi16 := 0; fi16 < 17; fi16++ {
        switch -833 {
        case 4294967295:
            v17 := GenS3{F0: (!false), F1: 2}
            _ = v17
        default:
            if (float64(float32((float64(0.0) + 1e10))) > -2.25) {
                v18 := []int{(fi16 % (((fi16 / ((fi16) | 1))) | 1)), fi16, 32767}
                fmt.Println(fi16, p1, fi16)
                _ = v18
            }
        }
        switch p1 {
        case "\n\t":
            var v19 float32 = ((0.1 / float32(0.1)) - float32(0.1))
            _ = v19
        case "foo bar":
            switch "\n\t" {
            case "hello":
                p1 = "foo bar"
            case "":
                p1 = "\n\t"
            default:
                p1 = (p1 + p1)
            }
        default:
            v20 := GenS4{F0: p0, F1: p0, F2: fi16}
            _ = v20
        }
        var v21 int16 = 1000
        _ = fi16
        _ = v21
    }
    v22 := []int{-434, 1, 377, 1}
    fmt.Println(p1, p1)
    if len(v22) > 0 {
        v22[0] = 0
    }
    v23 := GenS3{F0: (3.14159 >= (0.1 + float32(0.1))), F1: 46}
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 2
}

func main() {
    v1 := []uint32{42, 65535, 65535, 709}
    var v2 int8 = 127
    var v3 uint = 1
    v4 := []int{32768}
    var v5 int32 = 0
    v6 := []int{4294967295}
    for fi24 := 0; fi24 < 18; fi24++ {
        switch (fi24 + fi24) {
        case 42:
            var v25 float64 = (0.0 - 0.0)
            _ = v25
        case -1:
            if (!true) {
                v5 -= 0
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v26 uint64 = 18446744073709551615
                if true {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        var v28 uint = v3
                        _ = fi27
                        _ = v28
                    }
                }
                _ = v26
            }
        case 32768:
            v5 -= (-(v5))
        default:
            switch 32767 {
            case -9223372036854775808:
                v2 = (v2 | (v2 & 127))
            case -396:
                var v29 string = "foo bar"
                _ = v29
            case 0:
                if false {
                    v3 = v3
                    v2 += (v2 >> ((174) & 63))
                    switch (-(fi24)) {
                    case -2147483648:
                        switch "hello" {
                        case "hello":
                            v5 = v5
                        case "":
                            var v30 uint64 = 306
                            _ = v30
                        case "a":
                            var v31 uint8 = 70
                            _ = v31
                        default:
                            var v32 uint = 2
                            _ = v32
                        }
                    default:
                        v2 = v2
                    }
                }
            default:
                for fi33 := 0; fi33 < 14; fi33++ {
                    switch "hello" {
                    case "a":
                        v34 := GenS1{F0: (v2 & v2)}
                        _ = v34
                    case "\n\t":
                        var v35 uint8 = 2
                        _ = v35
                    case "hello":
                        v5 += (-(v5))
                        var v36 float32 = float32((0.0 * 0.1))
                        switch "" {
                        case "":
                            var v37 bool = false
                            _ = v37
                        default:
                            v6 = append(v6, -9223372036854775808)
                        }
                        v36 = 1.5
                        _ = v36
                    }
                    v38 := GenS4{F0: 2147483647, F1: v5, F2: (fi24 | 9223372036854775807)}
                    v4 = append(v4, fi24)
                    _ = fi33
                    _ = v38
                }
            }
        }
        fmt.Println(fi24)
        for fi39 := 0; fi39 < 3; fi39++ {
            v6 = append(v6, (fi39 << ((600) & 63)))
            v40 := GenS4{F0: (v5 - v5), F1: (v5 ^ -668), F2: fi39}
            fmt.Println(fi39)
            _ = fi39
            _ = v40
        }
        _ = fi24
    }
    v2 += (v2 | 127)
    var v41 uint16 = 761
    for fi42 := 0; fi42 < 1; fi42++ {
        var v43 int64 = -9223372036854775808
        if (!(v2 == 35)) {
            v44 := GenS3{F0: (!((v2 < (v2 + 127)) || ((v2 * 100) != (v2 - (v2 / ((101) | 1)))))), F1: 2}
            v4 = append(v4, (^(fi42)))
            _ = v44
        } else {
            var v45 uint16 = (v41 | (v41 | 255))
            switch fi42 {
            case -609:
                v41 += 255
            case 42:
                v46 := GenS5{F0: 1}
                _ = v46
            default:
                v6 = append(v6, (fi42 << ((v3) & 63)))
            }
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    var v47 int = 2147483647
    fmt.Println(v47)
    var v48 int64 = 1
    var v49 int8 = int8(v5)
    if (true && (586 > 532)) {
        v49 = (v49 - int8(v47))
        v2 += v49
        v47 = v47
    }
    if false {
        if true {
            fmt.Println(v3, v3, v3)
            v47 -= 988
        }
        var v50 bool = (42 != 951)
        _ = v50
    }
    var v51 float64 = -2.25
    for fi52 := 0; fi52 < 8; fi52++ {
        _ = len(v1)
        _ = fi52
    }
    if len(v6) > 0 {
        v6[0] = (v47 * (v47 << ((v3) & 63)))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v51
    _ = v6
}
