package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool }
func genFn0(p0 uint16, p1 uint16, p2 int8) int16 {
    p2 += (p2 & (p2 / (((p2 | 107)) | 1)))
    var v4 int32 = 1
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: p2}
        for fi7 := 0; fi7 < 13; fi7++ {
            p2 = (^(p2))
            _ = fi7
        }
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 float64 = (0.0 * float64(float32(0.1)))
            if (false && (!true)) {
                v4 += v4
            }
            var v10 string = "hello"
            p1 = (p0 + (p0 - (p1 % ((p1) | 1))))
            switch fi5 {
            case 4294967295:
                p1 += (p1 * (p1 ^ (p1 % ((2) | 1))))
            case 32768:
                fmt.Println(p0, v10)
            case -818:
                var v11 uint32 = 2147483648
                _ = v11
            default:
                var v12 float32 = 1.5
                _ = v12
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    var v13 uint64 = 673
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v4
    return int16(v4)
}
func genFn1() uint {
    var v14 int16 = 365
    var v15 int16 = v14
    _ = v14
    _ = v15
    return 2
}

func main() {
    v1 := []uint32{2147483648, 2, 42}
    v2 := make(map[uint32]string)
    var v3 int32 = -2147483648
    for fi16 := 0; fi16 < 13; fi16++ {
        v3 -= -120
        for fi17 := 0; fi17 < 14; fi17++ {
            v3 -= v3
            if ("" >= "a") {
                v3 = (v3 / ((v3) | 1))
                if false {
                    v18 := GenS1{F0: -1}
                    if (4294967295 > 4294967295) {
                        var v19 uint16 = 609
                        fmt.Println(fi16, v3)
                        v20 := GenS1{F0: 46}
                        _ = v19
                        _ = v20
                    } else {
                        v21 := GenS1{F0: 42}
                        _ = v21
                    }
                    _ = v18
                }
                if (1.5 > (1e10 / 1e10)) {
                    fmt.Println(fi16, fi17)
                    v22 := GenS1{F0: -1}
                    _ = v22
                } else {
                    v3 = 2147483647
                }
            }
            if (!true) {
                if true {
                    v23 := GenS2{F0: 280}
                    v3 += (v3 + v3)
                    _ = v23
                }
                v3 = v3
                var v24 uint16 = 42
                _ = v24
            } else {
                v25 := GenS3{F0: 127, F1: 0}
                v26 := GenS1{F0: -95}
                _ = v25
                _ = v26
            }
            _ = fi17
        }
        _ = fi16
    }
    if (true && ((0.1 >= 0.0) && true)) {
        var v27 float64 = 3.14159
        _ = v27
    }
    v28 := GenS4{F0: 530}
    var v29 int = 9223372036854775807
    v30 := GenS3{F0: 127, F1: 992}
    if ((-655 == 407) && true) {
        v3 = (v3 >> ((0) & 63))
        _ = len(v2)
    } else {
        if true {
            var v31 int8 = 127
            fmt.Println(v3)
            for fi32 := 0; fi32 < 11; fi32++ {
                fmt.Println(fi32)
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                var v33 int8 = (v31 + (-(v31)))
                _ = fi32
                _ = v33
            }
            _ = v31
        }
        v29 = 4294967295
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        var v35 uint16 = 65535
        switch "foo bar" {
        case "foo bar":
            v36 := GenS5{F0: 93}
            _ = v36
        }
        v1 = append(v1, 478)
        _ = fi34
        _ = v35
    }
    var v37 uint32 = 65535
    v1 = append(v1, v37)
    var v38 uint16 = 65535
    v37 += 745
    v1 = append(v1, (v37 << ((65535) & 63)))
    for fi39 := 0; fi39 < 7; fi39++ {
        var v40 float64 = 1.5
        var v41 uint64 = 376
        for fi42 := 0; fi42 < 15; fi42++ {
            if true {
                switch "foo bar" {
                case "\n\t":
                    switch "" {
                    case "hello":
                        v3 = -1
                    case "a":
                        var v43 int8 = -1
                        v40 = 0.1
                        _ = v43
                    case "":
                        v44 := GenS6{F0: (v3 / ((-2147483648) | 1))}
                        _ = v44
                    }
                case "hello":
                    fmt.Println(v38, v38)
                case "foo bar":
                    var v45 uint64 = 0
                    _ = v45
                }
                var v46 float64 = 1.5
                _ = v46
            }
            switch fi42 {
            case 32768:
                if true {
                    var v47 string = "a"
                    v29 = (fi42 + (-(fi39)))
                    var v48 int64 = 42
                    _ = v47
                    _ = v48
                } else {
                    var v49 uint16 = 32767
                    _ = v49
                }
            case 128:
                v50 := GenS7{F0: true}
                _ = v50
            case 651:
                v51 := GenS6{F0: (v3 << ((150) & 63))}
                _ = v51
            default:
                fmt.Println(fi39)
            }
            _ = fi42
        }
        _ = fi39
        _ = v40
        _ = v41
    }
    v37 += (v37 + 65535)
    v29 += 837
    var v52 uint8 = 6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v52
}
