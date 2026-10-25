package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 int32; F2 uint16 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint16, p2 int8) int {
    var v5 string = (("\n\t" + "") + "")
    switch -1 {
    case 9223372036854775807:
        if (true && (float64(0.0) >= float64(0.1))) {
            v6 := []int{9223372036854775807, 127, 9223372036854775807, -770}
            if (!true) {
                v5 = "foo bar"
            }
            _ = v6
        } else {
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 int64 = p0
                for fi9 := 0; fi9 < 4; fi9++ {
                    p1 = (p1 << ((18446744073709551615) & 63))
                    var v10 float32 = 1e10
                    _ = fi9
                    _ = v10
                }
                var v11 int8 = p2
                _ = fi7
                _ = v11
                _ = v8
            }
            var v12 int16 = 1000
            _ = v12
        }
    case 2147483647:
        if true {
            if true {
                switch (v5 + v5) {
                case "foo bar":
                    if ((0 < 182) && false) {
                        v13 := []int{2147483647, -842, 726}
                        _ = v13
                    } else {
                        if ((p2 >> ((18446744073709551615) & 63)) == 127) {
                            p1 = p1
                        }
                        v14 := GenS1{F0: p0}
                        _ = v14
                    }
                case "\n\t":
                    var v15 uint = 255
                    _ = v15
                case "a":
                    var v16 uint8 = 2
                    _ = v16
                }
                if (!(0.0 < ((float64(3.14159) * -2.25) + ((float64((0.0 - 0.0)) + 0.1) - float64(0.1))))) {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        if (!(p2 != (-(p2)))) {
                            var v18 float64 = (3.14159 - 3.14159)
                            v19 := []int{(-(fi17))}
                            v19 = append(v19, (fi17 * 127))
                            _ = v18
                            _ = v19
                        }
                        _ = fi17
                    }
                    var v20 uint = 1
                    _ = v20
                }
            }
        } else {
            p0 -= (p0 >> ((2) & 63))
            if (!(0 >= 141)) {
                if (42 > (p0 * (p0 - (p0 / (((p0 | 0)) | 1))))) {
                    v5 = "a"
                    switch v5 {
                    case "\n\t":
                        p0 = p0
                    case "hello":
                        v21 := []int{4294967295, -1}
                        _ = v21
                    default:
                        var v22 int = -1
                        _ = v22
                    }
                } else {
                    p0 -= (p0 ^ p0)
                    for fi23 := 0; fi23 < 9; fi23++ {
                        if (255 > p1) {
                            switch "\n\t" {
                            case "\n\t":
                                v24 := []int{}
                                _ = v24
                            case "foo bar":
                                v25 := GenS2{F0: 127}
                                _ = v25
                            case "":
                                var v26 float32 = (float32((-2.25 * (((-2.25 + ((float64(0.0) - float64(0.1)) - -2.25)) * (1.5 * 1e10)) - 0.1))) * (1e10 + 0.1))
                                _ = v26
                            }
                            v27 := GenS1{F0: (-(p0))}
                            v28 := GenS3{F0: (1e10 == 1.5)}
                            _ = v27
                            _ = v28
                        }
                        if (-479 != (p0 | -9223372036854775808)) {
                            var v29 uint64 = 42
                            _ = v29
                        }
                        _ = fi23
                    }
                    var v30 bool = true
                    _ = v30
                }
            } else {
                var v31 uint = 761
                p2 = (p2 | (p2 ^ p2))
                fmt.Println(p0, p0, p1)
                _ = v31
            }
        }
    case 0:
        v32 := GenS3{F0: (!(true && ((true || (!(true && true))) || true)))}
        _ = v32
    }
    var v33 uint = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    _ = v5
    return 941
}

func main() {
    v1 := []uint32{2147483648, 42, 1, 2147483647}
    v2 := make(map[string]uint32)
    var v3 int32 = 43
    v4 := make(map[string]bool)
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        fmt.Println(fi34, v3)
        if ("a" >= "") {
            fmt.Println(fi34)
        } else {
            switch (("\n\t" + "foo bar") + "\n\t") {
            case "hello":
                delete(v4, ("" + "a"))
            case "\n\t":
                v35 := GenS3{F0: ((fi34 | 2147483647) >= (fi34 ^ (fi34 >> ((42) & 63))))}
                _ = v35
            default:
                if false {
                    var v36 int64 = -9223372036854775808
                    _ = v36
                } else {
                    v37 := GenS3{F0: true}
                    for fi38 := 0; fi38 < 12; fi38++ {
                        fmt.Println(v3, v3, fi38)
                        v1 = append(v1, 42)
                        var v39 uint64 = 18446744073709551615
                        _ = fi38
                        _ = v39
                    }
                    _ = v37
                }
            }
        }
        _ = fi34
    }
    v40 := GenS4{F0: v3, F1: 1, F2: 0}
    v3 = v3
    switch "a" {
    case "foo bar":
        for fi41 := 0; fi41 < 18; fi41++ {
            if true {
                var v42 int64 = -1
                _ = v42
            } else {
                switch -570 {
                case 352:
                    v43 := GenS2{F0: 30}
                    _ = v43
                case 128:
                    v3 -= 42
                case 211:
                    v3 += (v3 + v3)
                }
            }
            fmt.Println(v3, fi41)
            _ = fi41
        }
    default:
        for fi44 := 0; fi44 < 10; fi44++ {
            for fi45 := 0; fi45 < 3; fi45++ {
                v3 += 42
                fmt.Println(fi44, v3)
                v3 = v3
                _ = fi45
            }
            _ = fi44
        }
    }
    v46 := GenS5{F0: -603, F1: 30}
    v47 := GenS1{F0: -1}
    for fi48 := 0; fi48 < 13; fi48++ {
        switch genFn0(1, 0, 99) {
        case 394:
            _ = len(v1)
        case 127:
            if (782 == 0) {
                v49 := GenS3{F0: (false || true)}
                v50 := v4["\n\t"]
                v50 = true
                _ = v49
                _ = v50
            }
        }
        v3 += (v3 << ((313) & 63))
        if false {
            v3 -= (-(v3))
            var v51 float64 = 0.1
            _ = v51
        }
        _ = fi48
    }
    v3 -= v3
    for fi52 := 0; fi52 < 19; fi52++ {
        v3 += v3
        for fi53 := 0; fi53 < 4; fi53++ {
            var v54 string = "foo bar"
            switch genFn0(-1, 32767, 127) {
            case 4294967295:
                switch v54 {
                case "":
                    var v55 float32 = float32(1e10)
                    _ = v55
                case "\n\t":
                    if (((float32(float64((0.0 - (1.5 + 0.1)))) / -2.25) / 0.1) < (0.1 + 1e10)) {
                        v56 := GenS2{F0: 0}
                        v3 = 42
                        fmt.Println(fi53)
                        _ = v56
                    }
                case "hello":
                    var v57 int32 = (v3 & (^(v3)))
                    _ = v57
                default:
                    var v58 uint = 255
                    _ = v58
                }
            case 205:
                v59 := GenS1{F0: 0}
                _ = v59
            case 2147483647:
                v60 := GenS6{F0: -20}
                _ = v60
            default:
                v61 := GenS7{F0: 42}
                _ = v61
            }
            _ = fi53
            _ = v54
        }
        v62 := GenS4{F0: v3, F1: v3, F2: 2}
        _ = fi52
        _ = v62
    }
    v3 = (^(v3))
    for fi63 := 0; fi63 < 18; fi63++ {
        if true {
            if true {
                _ = len(v4)
                _ = len(v4)
                switch (fi63 * fi63) {
                case -585:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 9223372036854775807:
                    fmt.Println(v3)
                case -9223372036854775808:
                    v3 += (v3 << ((42) & 63))
                default:
                    v3 -= v3
                }
            } else {
                for fi64 := 0; fi64 < 14; fi64++ {
                    var v65 int32 = 129
                    _ = fi64
                    _ = v65
                }
            }
        }
        var v66 int = fi63
        if false {
            switch v66 {
            case 0:
                if true {
                    if false {
                        fmt.Println(v3, v3)
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                }
            case 128:
                var v67 uint64 = 1
                _ = v67
            case 2147483647:
                var v68 float32 = 0.1
                _ = v68
            default:
                v1 = append(v1, 142)
            }
        }
        _ = fi63
        _ = v66
    }
    var v69 int8 = 42
    var v70 int64 = 1
    switch "\n\t" {
    case "":
        v69 -= (v69 << ((420) & 63))
    case "a":
        var v71 uint16 = 0
        _ = v71
    case "foo bar":
        v69 = -1
    default:
        v69 += 0
    }
    switch "" {
    case "\n\t":
        v72 := GenS8{F0: 1}
        _ = v72
    }
    var v73 string = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v46
    _ = v47
    _ = v69
    _ = v70
    _ = v73
}
