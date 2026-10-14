package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16; F1 uint32; F2 int8 }
type GenS3 struct { F0 int32; F1 bool; F2 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 uint16; F1 string; F2 bool }
type GenS8 struct { F0 string; F1 int64; F2 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 int64 }
type GenS11 struct { F0 int; F1 uint64; F2 int32 }
func genFn0(p0 int, p1 int16) string {
    var v5 uint = 18446744073709551615
    switch p0 {
    case 730:
        switch p0 {
        case 2147483647:
            p1 = p1
        case -1:
            p1 += 152
        default:
            for fi6 := 0; fi6 < 17; fi6++ {
                fmt.Println(p0, p1, v5)
                var v7 int = p0
                _ = fi6
                _ = v7
            }
        }
    case -921:
        v8 := []int{p0}
        _ = v8
    case 9223372036854775807:
        var v9 uint16 = 42
        _ = v9
    default:
        var v10 uint32 = 2147483648
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        p0 = fi11
        _ = fi11
    }
    v12 := GenS1{F0: 718, F1: (v5 * v5)}
    var v13 uint32 = 411
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v5
    return ""
}
func genFn1(p0 bool, p1 int64, p2 uint32) uint8 {
    v14 := GenS2{F0: -1, F1: 133, F2: 127}
    var v15 uint64 = 538
    v16 := GenS2{F0: 32767, F1: (p2 - 487), F2: 0}
    v15 = (v15 - (v15 >> ((65535) & 63)))
    p1 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return 162
}

func main() {
    v1 := []int{32768, 852, 2147483647, 0}
    v2 := make(map[int]int)
    var v3 uint8 = 2
    var v4 float32 = 3.14159
    switch "\n\t" {
    case "":
        var v17 int16 = -578
        _ = v17
    default:
        switch 32768 {
        case 42:
            for fi18 := 0; fi18 < 5; fi18++ {
                if (-1 <= 1) {
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v1 = append(v1, 42)
                        _ = fi19
                    }
                    v20 := GenS3{F0: 2147483647, F1: false, F2: true}
                    _ = v20
                }
                var v21 int8 = 127
                _ = fi18
                _ = v21
            }
        case 994:
            var v22 uint16 = 0
            _ = v22
        }
    }
    v23 := GenS2{F0: 0, F1: 659, F2: 46}
    var v24 int64 = -9223372036854775808
    for fi25 := 0; fi25 < 12; fi25++ {
        v4 = v4
        for fi26 := 0; fi26 < 10; fi26++ {
            switch "" {
            case "hello":
                fmt.Println(fi26, fi26, fi26)
            case "":
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v24, fi25, fi26)
                case "foo bar":
                    v4 += v4
                case "":
                    if false {
                        v27 := GenS1{F0: -2147483648, F1: 42}
                        v28 := GenS4{F0: 127}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch "hello" {
                        case "a":
                            fmt.Println(fi25)
                        case "foo bar":
                            if false {
                                v29 := GenS4{F0: v3}
                                _ = v29
                            }
                            v30 := GenS2{F0: 0, F1: 2147483648, F2: 0}
                            var v31 uint16 = 32767
                            _ = v30
                            _ = v31
                        case "":
                            var v32 int16 = 0
                            _ = v32
                        }
                        var v33 bool = (!(false && (!false)))
                        _ = v27
                        _ = v28
                        _ = v33
                    } else {
                        switch "" {
                        case "":
                            var v34 bool = (!(!true))
                            _ = v34
                        default:
                            v24 = (^(v24))
                        }
                        var v35 int = fi26
                        _ = v35
                    }
                default:
                    switch "hello" {
                    case "a":
                        fmt.Println(fi25, fi26, fi25)
                    case "hello":
                        v3 = (v3 & (v3 ^ 128))
                    case "\n\t":
                        v24 = (-(v24))
                    default:
                        v3 -= (v3 & v3)
                    }
                    v36 := GenS4{F0: (v3 / (((^(v3))) | 1))}
                    v24 -= (v24 >> ((22) & 63))
                    var v37 uint8 = 1
                    _ = v36
                    _ = v37
                }
            default:
                switch (("\n\t" + "hello") + "") {
                case "":
                    v3 = v3
                }
            }
            v3 = (^(v3))
            _ = fi26
        }
        _ = fi25
    }
    v4 += float32((3.14159 / -2.25))
    switch "hello" {
    case "foo bar":
        if (false && (955 <= 42)) {
            for fi38 := 0; fi38 < 17; fi38++ {
                switch fi38 {
                case 2147483647:
                    v24 = 195
                case -221:
                    if false {
                        v4 = ((0.0 / ((float32(((((0.0 * 1.5) * 3.14159) / float64(((v4 * (3.14159 * v4)) - float32(0.1)))) + (((-2.25 + 3.14159) - (float64(v4) + 1.5)) - ((1e10 - (3.14159 * 1e10)) - (0.0 / 3.14159))))) - 1e10) + ((float32(float64(float32((3.14159 + 1e10)))) + 3.14159) * ((v4 - v4) + v4)))) * (v4 - 1.5))
                        var v39 uint = 4294967295
                        _ = v39
                    } else {
                        v24 -= 368
                        for fi40 := 0; fi40 < 10; fi40++ {
                            v41 := v2[(fi40 % (((fi40 | fi40)) | 1))]
                            v42 := GenS4{F0: v3}
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        v3 = (^(v3))
                    }
                }
                v24 -= (v24 / ((v24) | 1))
                v3 = (v3 << ((0) & 63))
                _ = fi38
            }
            for fi43 := 0; fi43 < 7; fi43++ {
                if (0 > 18446744073709551615) {
                    v44 := v2[int(v3)]
                    var v45 string = "hello"
                    if (v4 >= -2.25) {
                        for fi46 := 0; fi46 < 15; fi46++ {
                            v44 -= fi43
                            _ = fi46
                        }
                        for fi47 := 0; fi47 < 4; fi47++ {
                            v48 := GenS4{F0: (^(v3))}
                            _ = fi47
                            _ = v48
                        }
                        if ((1 > 42) || (!(!(true || (-1 == -2147483648))))) {
                            if (!true) {
                                switch 32767 {
                                case -2147483648:
                                    if (true && true) {
                                        fmt.Println(v3, v45)
                                        if false {
                                            if false {
                                                var v49 bool = (754 == 65535)
                                                var v50 float32 = (v4 / 1.0)
                                                _ = v49
                                                _ = v50
                                            }
                                        }
                                        fmt.Println(fi43, v3, v45)
                                    }
                                }
                                var v51 int8 = 37
                                _ = v51
                            } else {
                                var v52 bool = (319 >= 42)
                                _ = v52
                            }
                            v1 = append(v1, v44)
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            delete(v2, (fi43 | (v44 + (^(fi43)))))
                        }
                    } else {
                        v53 := GenS5{F0: 1}
                        var v54 uint = 2
                        _ = v53
                        _ = v54
                    }
                    _ = v44
                    _ = v45
                }
                for fi55 := 0; fi55 < 4; fi55++ {
                    var v56 uint16 = 541
                    var v57 uint32 = 775
                    v58 := GenS2{F0: -718, F1: (v57 >> ((1) & 63)), F2: 127}
                    switch "foo bar" {
                    case "hello":
                        v59 := GenS6{F0: 42, F1: 158}
                        _ = v59
                    }
                    v1 = append(v1, 32768)
                    _ = fi55
                    _ = v56
                    _ = v57
                    _ = v58
                }
                switch (fi43 * fi43) {
                case 4294967295:
                    delete(v2, (^(fi43)))
                case 32768:
                    var v60 int64 = 9223372036854775807
                    _ = v60
                case 128:
                    v61 := GenS7{F0: uint16(v3), F1: "", F2: (true && (0 < 0))}
                    _ = v61
                }
                _ = fi43
            }
        }
    case "a":
        switch "hello" {
        case "\n\t":
            switch 4294967295 {
            case 260:
                if (v3 <= (^(v3))) {
                    v4 -= v4
                }
            case -1:
                if false {
                    v3 -= v3
                }
            case -9223372036854775808:
                if false {
                    fmt.Println(v4, v3, v4)
                    v62 := GenS7{F0: 942, F1: "foo bar", F2: (289 == 1)}
                    _ = v62
                }
            }
        }
    case "hello":
        v63 := GenS8{F0: ("" + (("\n\t" + "") + "")), F1: v24, F2: true}
        _ = v63
    }
    var v64 uint32 = 83
    for fi65 := 0; fi65 < 8; fi65++ {
        v66 := GenS9{F0: v64}
        if (0 <= -128) {
            var v67 uint8 = 2
            _ = v67
        } else {
            fmt.Println(v24)
        }
        var v68 float64 = (0.1 + -2.25)
        _ = fi65
        _ = v66
        _ = v68
    }
    v69 := v2[-1]
    fmt.Println(v24, v24)
    for fi70 := 0; fi70 < 11; fi70++ {
        var v71 int8 = -110
        if ((v3 - 127) == v3) {
            fmt.Println(v71)
            for fi72 := 0; fi72 < 10; fi72++ {
                v64 = v64
                _ = fi72
            }
            if true {
                if true {
                    fmt.Println(v24)
                    v69 = fi70
                    var v73 uint = 255
                    _ = v73
                } else {
                    for fi74 := 0; fi74 < 20; fi74++ {
                        var v75 uint64 = 0
                        fmt.Println(fi70)
                        _ = fi74
                        _ = v75
                    }
                    for fi76 := 0; fi76 < 18; fi76++ {
                        v64 = 42
                        switch (fi76 << ((955) & 63)) {
                        case -2147483648:
                            var v77 int16 = 32767
                            _ = v77
                        }
                        switch fi76 {
                        case 1:
                            v78 := GenS10{F0: 42, F1: (v24 & v24)}
                            _ = v78
                        }
                        v69 -= 32767
                        _ = fi76
                    }
                    v79 := GenS8{F0: ("foo bar" + "hello"), F1: (v24 ^ (v24 ^ (v24 / ((v24) | 1)))), F2: ((126 > 859) && false)}
                    _ = v79
                }
            } else {
                v80 := GenS4{F0: genFn1(("\n\t" >= ("foo bar" + "")), (v24 % ((v24) | 1)), v64)}
                for fi81 := 0; fi81 < 19; fi81++ {
                    v71 = 0
                    fmt.Println(v3)
                    _ = fi81
                }
                v1 = append(v1, (-(fi70)))
                _ = v80
            }
        }
        _ = fi70
        _ = v71
    }
    for fi82 := 0; fi82 < 17; fi82++ {
        v83 := GenS11{F0: (v69 % ((-1) | 1)), F1: 2, F2: 2147483647}
        fmt.Println(v3, fi82, v3)
        _ = fi82
        _ = v83
    }
    var v84 uint16 = 42
    var v85 int8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    fmt.Println("v69:", v69)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v64
    _ = v69
    _ = v84
    _ = v85
}
