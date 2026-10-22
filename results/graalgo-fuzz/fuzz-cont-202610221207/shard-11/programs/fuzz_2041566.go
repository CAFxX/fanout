package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int; F1 string; F2 bool }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 uint32; F1 int }
func genFn0(p0 int16, p1 bool, p2 string) int32 {
    v6 := GenS1{F0: 63, F1: 213, F2: 539}
    switch p2 {
    case "hello":
        p1 = (0 <= 37)
    default:
        for fi7 := 0; fi7 < 17; fi7++ {
            if ((576 == -1) && (!(((p1 && (true || (18446744073709551615 <= 0))) && p1) || true))) {
                if (42 <= 1) {
                    fmt.Println(fi7, fi7)
                    v8 := []int{(-(fi7)), fi7, (fi7 * (fi7 << ((586) & 63))), 32767}
                    for fi9 := 0; fi9 < 17; fi9++ {
                        var v10 int = fi9
                        fmt.Println(p0, v10)
                        _ = fi9
                        _ = v10
                    }
                    _ = v8
                } else {
                    p2 = "\n\t"
                    v11 := GenS1{F0: 1, F1: 128, F2: 42}
                    v12 := GenS1{F0: -128, F1: 225, F2: 2}
                    _ = v11
                    _ = v12
                }
            }
            _ = fi7
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1(p0 int8, p1 uint) uint8 {
    v13 := []int{-9223372036854775808, 1, 32767, -9223372036854775808}
    fmt.Println(p1, p0)
    switch "" {
    case "hello":
        var v14 uint8 = 127
        _ = v14
    case "":
        v15 := GenS1{F0: p0, F1: 255, F2: 549}
        _ = v15
    case "foo bar":
        switch "foo bar" {
        case "":
            for fi16 := 0; fi16 < 13; fi16++ {
                p1 += 18446744073709551615
                fmt.Println(p1, fi16, p1)
                p0 = p0
                _ = fi16
            }
        case "hello":
            if (0 > 127) {
                switch (("\n\t" + "a") + "\n\t") {
                case "a":
                    switch "\n\t" {
                    case "hello":
                        for fi17 := 0; fi17 < 6; fi17++ {
                            v18 := GenS2{F0: (^(p1))}
                            var v19 uint64 = 0
                            switch fi17 {
                            case 32767:
                                fmt.Println(p1)
                            case 42:
                                var v20 int64 = 0
                                _ = v20
                            }
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    case "a":
                        switch ("" + "foo bar") {
                        case "foo bar":
                            var v21 int16 = 32767
                            _ = v21
                        case "a":
                            for fi22 := 0; fi22 < 10; fi22++ {
                                for fi23 := 0; fi23 < 5; fi23++ {
                                    var v24 int = 127
                                    _ = fi23
                                    _ = v24
                                }
                                var v25 int8 = -106
                                _ = fi22
                                _ = v25
                            }
                        default:
                            var v26 float32 = 0.1
                            _ = v26
                        }
                    default:
                        switch -9223372036854775808 {
                        case 1:
                            fmt.Println(p0, p1)
                        case 127:
                            p0 -= (p0 << ((p1) & 63))
                        case -352:
                            if (!true) {
                                v27 := GenS1{F0: p0, F1: 112, F2: 1}
                                if true {
                                    var v28 uint16 = 244
                                    v28 = (v28 & (v28 - (v28 & 65535)))
                                    _ = v28
                                }
                                _ = v27
                            } else {
                                fmt.Println(p1, p1)
                                for fi29 := 0; fi29 < 1; fi29++ {
                                    var v30 int16 = 42
                                    var v31 uint = (^(p1))
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            }
                        default:
                            v32 := GenS3{F0: 42, F1: 42}
                            _ = v32
                        }
                    }
                case "foo bar":
                    var v33 string = ("hello" + "\n\t")
                    _ = v33
                case "":
                    p0 += int8(p1)
                }
                v34 := GenS3{F0: 0, F1: 85}
                _ = v34
            }
        case "a":
            if true {
                switch ("a" + ("" + (("" + "foo bar") + "a"))) {
                case "":
                    if true {
                        switch "foo bar" {
                        case "a":
                            if len(v13) > 0 {
                                v13[0] = -1
                            }
                        }
                        v35 := GenS1{F0: p0, F1: 178, F2: 89}
                        for fi36 := 0; fi36 < 19; fi36++ {
                            p0 = p0
                            _ = fi36
                        }
                        _ = v35
                    } else {
                        p1 -= (p1 % ((p1) | 1))
                    }
                }
                v37 := GenS1{F0: 0, F1: 68, F2: 418}
                var v38 int64 = -9223372036854775808
                _ = v37
                _ = v38
            } else {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            }
        }
    default:
        var v39 uint = (^(p1))
        _ = v39
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 bool = (p0 <= p0)
        p1 = p1
        _ = fi40
        _ = v41
    }
    _ = p0
    _ = p1
    _ = v13
    return 54
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 uint16 = 1
    var v4 float64 = ((3.14159 - 3.14159) + (0.0 / (((-2.25 - (0.0 / 3.14159)) - (-2.25 - 0.0)) + 3.14159)))
    v5 := []string{("" + "foo bar"), ("foo bar" + ""), "hello"}
    switch 127 {
    case 9223372036854775807:
        v1[-2147483648] = 32768
    case 717:
        var v42 uint16 = v3
        _ = v42
    }
    v3 = v3
    if (-228 == -239) {
        for fi43 := 0; fi43 < 17; fi43++ {
            var v44 uint = 42
            _ = len(v2)
            _ = fi43
            _ = v44
        }
        if true {
            var v45 int8 = 100
            if len(v5) > 0 {
                v5[0] = ""
            }
            _ = v45
        } else {
            for fi46 := 0; fi46 < 16; fi46++ {
                for fi47 := 0; fi47 < 14; fi47++ {
                    v3 = (v3 | (v3 >> ((2) & 63)))
                    v4 += 0.0
                    fmt.Println(v4)
                    _ = fi47
                }
                v48 := GenS1{F0: -86, F1: 0, F2: 0}
                switch (fi46 >> ((2) & 63)) {
                case 287:
                    v49 := GenS4{F0: (fi46 & fi46), F1: (("" + ((("hello" + "hello") + "") + "a")) + ("foo bar" + "foo bar")), F2: (!false)}
                    _ = v49
                case 127:
                    if (650 != 993) {
                        switch 0 {
                        case 128:
                            v4 = (float64(3.14159) - v4)
                        case 127:
                            delete(v1, (fi46 ^ fi46))
                        case 1:
                            for fi50 := 0; fi50 < 12; fi50++ {
                                fmt.Println(v3, v3)
                                v3 = 2
                                _ = fi50
                            }
                        }
                    } else {
                        switch fi46 {
                        case 128:
                            v3 += (v3 / (((v3 | 65535)) | 1))
                        case -2147483648:
                            for fi51 := 0; fi51 < 6; fi51++ {
                                fmt.Println(fi51, fi46)
                                v52 := GenS1{F0: -128, F1: 148, F2: 1}
                                _ = fi51
                                _ = v52
                            }
                        default:
                            if (0.1 != v4) {
                                v3 = 181
                            }
                        }
                        if (v4 < v4) {
                            v4 = v4
                            for fi53 := 0; fi53 < 14; fi53++ {
                                switch ((("a" + "\n\t") + "foo bar") + ("foo bar" + "foo bar")) {
                                case "":
                                    v54 := GenS5{F0: "a", F1: 510}
                                    _ = v54
                                case "a":
                                    var v55 uint32 = 2147483647
                                    _ = v55
                                case "\n\t":
                                    v56 := GenS6{F0: -604, F1: 588}
                                    v3 -= v3
                                    v57 := GenS4{F0: fi46, F1: ("\n\t" + (("" + ("foo bar" + "foo bar")) + "foo bar")), F2: (!false)}
                                    v58 := GenS7{F0: (^(v3)), F1: 65535}
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                default:
                                    var v59 uint64 = 0
                                    _ = v59
                                }
                                v60 := GenS3{F0: -1, F1: 1}
                                _ = fi53
                                _ = v60
                            }
                            var v61 uint64 = 1
                            _ = v61
                        } else {
                            for fi62 := 0; fi62 < 15; fi62++ {
                                var v63 int8 = -1
                                var v64 uint8 = 42
                                _ = fi62
                                _ = v63
                                _ = v64
                            }
                        }
                        for fi65 := 0; fi65 < 1; fi65++ {
                            fmt.Println(fi46, v3)
                            v66 := GenS4{F0: (fi65 * fi65), F1: "foo bar", F2: false}
                            v4 = 1.5
                            _ = fi65
                            _ = v66
                        }
                    }
                default:
                    if true {
                        var v67 uint64 = 18446744073709551615
                        v5 = append(v5, ("foo bar" + "foo bar"))
                        _ = v67
                    } else {
                        v3 += (v3 >> ((2) & 63))
                    }
                }
                _ = fi46
                _ = v48
            }
        }
    } else {
        v68 := GenS8{F0: 972, F1: 153}
        _ = v68
    }
    v4 += 1e10
    if (18446744073709551615 > 486) {
        v1[32768] = 338
        v69 := GenS7{F0: 255, F1: v3}
        _ = v69
    }
    for fi70 := 0; fi70 < 17; fi70++ {
        v4 -= ((v4 + (3.14159 + v4)) * 1e10)
        _ = fi70
    }
    if false {
        v4 -= ((1e10 * (1e10 - v4)) * v4)
    }
    var v71 bool = (32767 == 552)
    var v72 int32 = 2147483647
    var v73 float64 = (v4 - ((v4 - float64(3.14159)) / (v4 * v4)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v72
    _ = v73
}
