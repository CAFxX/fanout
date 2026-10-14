package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 uint16; F1 int64; F2 uint64 }
func genFn0(p0 string) int32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        switch fi7 {
        case -2147483648:
            var v8 int = (fi7 & -1)
            _ = v8
        case 880:
            var v9 int = (-(fi7))
            _ = v9
        case 32767:
            v10 := GenS1{F0: 1, F1: -2147483648, F2: p0}
            _ = v10
        }
        _ = fi7
    }
    p0 = ("foo bar" + (((p0 + p0) + p0) + ""))
    p0 = ""
    v11 := GenS1{F0: 2, F1: -2147483648, F2: ("foo bar" + (p0 + p0))}
    v12 := []int{}
    _ = p0
    _ = v11
    _ = v12
    return 0
}
func genFn1() int64 {
    var v13 string = ("hello" + ("foo bar" + ""))
    for fi14 := 0; fi14 < 20; fi14++ {
        if (true && false) {
            for fi15 := 0; fi15 < 20; fi15++ {
                v16 := GenS2{F0: -2147483648, F1: 4294967295}
                _ = fi15
                _ = v16
            }
            switch (fi14 ^ fi14) {
            case 798:
                for fi17 := 0; fi17 < 7; fi17++ {
                    v18 := GenS3{F0: 690, F1: 255, F2: 1}
                    _ = fi17
                    _ = v18
                }
            default:
                v13 = v13
            }
            v19 := GenS4{F0: (v13 + "")}
            _ = v19
        }
        v20 := GenS5{F0: 2}
        _ = fi14
        _ = v20
    }
    _ = v13
    return -9223372036854775808
}

func main() {
    var v1 int16 = 32767
    v2 := []string{"a"}
    var v3 int32 = 707
    var v4 uint32 = 1
    v5 := []uint32{2, v4, 489, 730}
    var v6 float32 = (3.14159 * -2.25)
    for fi21 := 0; fi21 < 12; fi21++ {
        fmt.Println(fi21, v1)
        _ = fi21
    }
    fmt.Println(v3)
    if (127 != 519) {
        v22 := GenS5{F0: 65535}
        fmt.Println(v4, v1, v4)
        _ = v22
    }
    v4 += 1
    var v23 float64 = -2.25
    for fi24 := 0; fi24 < 4; fi24++ {
        v4 = (v4 % (((v4 ^ (^(v4)))) | 1))
        v1 += (v1 ^ (v1 & (v1 - v1)))
        switch "a" {
        case "foo bar":
            v25 := GenS1{F0: 42, F1: (v3 / (((v3 - 1)) | 1)), F2: ("" + ("hello" + ""))}
            _ = v25
        default:
            var v26 uint8 = 42
            _ = v26
        }
        _ = fi24
    }
    var v27 uint64 = 18446744073709551615
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 bool = true
        v30 := GenS6{F0: uint16(v4), F1: 0}
        _ = fi28
        _ = v29
        _ = v30
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 = append(v2, ("a" + "foo bar"))
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 float32 = (v6 - v6)
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        for fi34 := 0; fi34 < 14; fi34++ {
            v1 -= v1
            var v35 int16 = (^(v1))
            switch "hello" {
            case "":
                v36 := GenS7{F0: 245, F1: -9223372036854775808, F2: (^(v27))}
                _ = v36
            }
            _ = fi34
            _ = v35
        }
        _ = len(v2)
        _ = fi33
    }
    if true {
        switch ("" + "hello") {
        case "a":
            v37 := GenS2{F0: v3, F1: 2}
            _ = v37
        default:
            var v38 uint32 = v4
            _ = v38
        }
        for fi39 := 0; fi39 < 9; fi39++ {
            for fi40 := 0; fi40 < 4; fi40++ {
                v4 = 65535
                if ((^(v27)) <= v27) {
                    v27 -= (v27 % ((0) | 1))
                    var v41 uint8 = 0
                    _ = v41
                }
                switch ("" + "") {
                case "\n\t":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi40
            }
            v4 = v4
            var v42 int8 = 1
            _ = fi39
            _ = v42
        }
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v1, v3)
        case "\n\t":
            v43 := GenS1{F0: 583, F1: v3, F2: "a"}
            _ = v43
        default:
            for fi44 := 0; fi44 < 19; fi44++ {
                var v45 float64 = ((((v23 / v23) + (((-2.25 - v23) * v23) - v23)) / 1.0) / v23)
                _ = fi44
                _ = v45
            }
        }
    }
    if false {
        v4 += 795
        v1 += (v1 ^ (v1 / (((v1 + v1)) | 1)))
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        v23 = v23
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
