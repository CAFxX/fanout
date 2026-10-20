package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int8, p2 uint8) uint32 {
    v4 := GenS1{F0: 1, F1: 127, F2: 9223372036854775807}
    switch "" {
    case "foo bar":
        var v5 uint32 = 2147483648
        _ = v5
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        p0 -= p0
        v7 := []int{-2147483648, fi6, (-(fi6)), -9223372036854775808}
        var v8 int32 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2147483648
}
func genFn1(p0 int64, p1 uint32) uint {
    if false {
        if true {
            v9 := []int{2147483647}
            switch "" {
            case "":
                p0 = (p0 / (((p0 / ((42) | 1))) | 1))
            case "foo bar":
                var v10 int32 = 882
                _ = v10
            case "hello":
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := GenS1{F0: 0, F1: 127, F2: 1}
                    p1 = (p1 + genFn0(uint8(p1), -1, 96))
                    _ = len(v9)
                    _ = fi11
                    _ = v12
                }
            default:
                for fi13 := 0; fi13 < 4; fi13++ {
                    if len(v9) > 0 {
                        v9[0] = fi13
                    }
                    _ = fi13
                }
            }
            v14 := GenS1{F0: 255, F1: -969, F2: (-(p0))}
            _ = v14
            _ = v9
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := []int{(fi15 / ((fi15) | 1))}
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS1{F0: 173, F1: -208, F2: (p0 ^ 42)}
    if (int(p1) != 0) {
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 uint = 1
            _ = fi18
            _ = v19
        }
        p0 = p0
    }
    if true {
        v20 := GenS1{F0: 967, F1: 32767, F2: p0}
        for fi21 := 0; fi21 < 19; fi21++ {
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 int64 = 9223372036854775807
                v24 := GenS1{F0: 29, F1: fi22, F2: 9223372036854775807}
                v23 = 1
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v17
    return uint(p1)
}

func main() {
    v1 := []int{32768, 32767}
    var v2 uint64 = 18446744073709551615
    var v3 float64 = 1e10
    if true {
        v1 = append(v1, 528)
    } else {
        fmt.Println(v3, v2)
        v3 = v3
    }
    _ = len(v1)
    v25 := GenS2{F0: ((("" + "foo bar") + (("\n\t" + ((("" + ("hello" + ("" + "\n\t"))) + (("\n\t" + ("hello" + "\n\t")) + "\n\t")) + ("hello" + "a"))) + "hello")) + "\n\t")}
    v2 = 46
    if true {
        for fi26 := 0; fi26 < 15; fi26++ {
            if false {
                switch "" {
                case "foo bar":
                    var v27 string = "foo bar"
                    _ = v27
                case "hello":
                    v3 = (-2.25 / 1.5)
                case "a":
                    var v28 float64 = (((0.0 + (v3 + (v3 / (v3 + v3)))) + 0.0) * float64(-2.25))
                    _ = v28
                default:
                    var v29 float32 = -2.25
                    _ = v29
                }
            } else {
                v30 := GenS2{F0: "foo bar"}
                var v31 int8 = 42
                _ = v30
                _ = v31
            }
            v2 = v2
            switch 0 {
            case 42:
                if false {
                    v32 := GenS2{F0: "foo bar"}
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        v33 := GenS3{F0: 1, F1: (fi26 << ((992) & 63))}
                        _ = v33
                    default:
                        if true {
                            if (-1 <= 127) {
                                switch ("foo bar" + "\n\t") {
                                case "hello":
                                    v3 -= (v3 / 1e10)
                                case "foo bar":
                                    fmt.Println(v2)
                                case "\n\t":
                                    var v34 uint64 = v2
                                    _ = v34
                                }
                            }
                            var v35 uint32 = 0
                            fmt.Println(v35)
                            _ = v35
                        } else {
                            fmt.Println(v2, fi26)
                            v36 := GenS2{F0: ("\n\t" + ("hello" + "a"))}
                            _ = v36
                        }
                    }
                    for fi37 := 0; fi37 < 8; fi37++ {
                        v38 := GenS3{F0: 128, F1: 127}
                        var v39 int16 = 32767
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v3 = (v3 / (1e10 * (1e10 + (0.1 / 1.5))))
                            v2 += (v2 * (v2 / ((v2) | 1)))
                        }
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    _ = v32
                }
            case -2147483648:
                if (3.14159 >= 0.1) {
                    v1 = append(v1, (fi26 / ((fi26) | 1)))
                    _ = len(v1)
                    v3 += ((((v3 - v3) + (((v3 + (float64((0.1 + (float32((3.14159 / (float64((-2.25 + 0.0)) + v3))) * float32(((float64(-2.25) * (0.1 - (v3 + v3))) - ((v3 / v3) / 1e10)))))) + v3)) * ((v3 / v3) + v3)) - 0.0)) + (0.1 + (0.1 / ((1e10 / v3) / v3)))) * v3)
                } else {
                    switch fi26 {
                    case 128:
                        v3 = 0.0
                    case 421:
                        for fi40 := 0; fi40 < 1; fi40++ {
                            if ((-711 < 1) || ((1e10 + 3.14159) <= 0.0)) {
                                v41 := GenS3{F0: 42, F1: fi40}
                                _ = v41
                            } else {
                                switch ("a" + "hello") {
                                case "hello":
                                    if false {
                                        v1 = append(v1, -9223372036854775808)
                                        var v42 uint16 = 2
                                        var v43 int16 = 1
                                        if ((554 == 951) && false) {
                                            v2 += (^(v2))
                                            var v44 uint = 1
                                            _ = v44
                                        }
                                        var v45 bool = false
                                        _ = v42
                                        _ = v43
                                        _ = v45
                                    } else {
                                        if ((65535 >= 2147483647) || (0 != 32767)) {
                                            var v46 int = fi26
                                            fmt.Println(v46, fi40, fi26)
                                            switch "hello" {
                                            case "\n\t":
                                                var v47 int8 = 111
                                                _ = v47
                                            case "a":
                                                var v48 uint32 = 337
                                                _ = v48
                                            case "":
                                                switch fi26 {
                                                case -9223372036854775808:
                                                    if true {
                                                        v49 := GenS4{F0: -20}
                                                        v50 := GenS5{F0: 2147483647}
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        var v51 uint8 = uint8(v2)
                                                        var v52 int = fi40
                                                        if false {
                                                            v52 += (-(v46))
                                                            if (!(-793 != -644)) {
                                                                v53 := GenS5{F0: 178}
                                                                var v54 int64 = -610
                                                                _ = v53
                                                                _ = v54
                                                            } else {
                                                                var v55 float32 = ((float32(float64(0.0)) * 0.0) * 1e10)
                                                                v55 = (v55 - 1e10)
                                                                _ = v55
                                                            }
                                                            _ = len(v1)
                                                        }
                                                        switch 4294967295 {
                                                        case 818:
                                                            v3 -= v3
                                                        default:
                                                            if false {
                                                                v2 = (v2 >> ((65535) & 63))
                                                                v56 := GenS5{F0: genFn0((v51 - 128), -1, 128)}
                                                                var v57 float32 = (3.14159 + (0.0 - float32((v3 - 0.0))))
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                case 32768:
                                                    if (false && false) {
                                                        v1 = append(v1, -830)
                                                    }
                                                case -640:
                                                    v58 := GenS2{F0: "hello"}
                                                    _ = v58
                                                }
                                            default:
                                                if (!(true || true)) {
                                                    _ = len(v1)
                                                    v46 = (fi40 & 4294967295)
                                                    v46 = (fi26 + fi40)
                                                }
                                            }
                                            _ = v46
                                        } else {
                                            fmt.Println(fi26)
                                            var v59 uint32 = 42
                                            _ = v59
                                        }
                                        v3 = v3
                                        v60 := GenS6{F0: (((true && (42 < 42)) || false) && (-715 <= -238))}
                                        _ = v60
                                    }
                                case "":
                                    var v61 int8 = 100
                                    switch (("hello" + "\n\t") + "hello") {
                                    case "foo bar":
                                        var v62 uint16 = 255
                                        _ = v62
                                    case "a":
                                        switch "hello" {
                                        case "hello":
                                            var v63 int16 = 1000
                                            _ = v63
                                        case "\n\t":
                                            var v64 int8 = 84
                                            _ = v64
                                        case "":
                                            _ = len(v1)
                                        default:
                                            var v65 uint8 = 255
                                            _ = v65
                                        }
                                    default:
                                        if true {
                                            if true {
                                                fmt.Println(v2, v61)
                                                var v66 int64 = 912
                                                v3 = (float64(0.1) + v3)
                                                _ = v66
                                            }
                                            v61 -= (-(v61))
                                            var v67 int8 = (v61 >> ((genFn1(9223372036854775807, 2)) & 63))
                                            _ = v67
                                        }
                                    }
                                    _ = v61
                                default:
                                    var v68 uint8 = 2
                                    _ = v68
                                }
                            }
                            _ = fi40
                        }
                    case 583:
                        if false {
                            if false {
                                v3 -= -2.25
                            } else {
                                v2 = v2
                                v69 := GenS4{F0: 127}
                                _ = v69
                            }
                            fmt.Println(fi26)
                        } else {
                            if (!(128 > 164)) {
                                v70 := GenS6{F0: false}
                                v71 := GenS7{F0: 0}
                                _ = len(v1)
                                _ = v70
                                _ = v71
                            }
                            v2 += v2
                        }
                    }
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi26
        }
        var v72 bool = (float64(-2.25) != 0.0)
        _ = v72
    } else {
        switch 32768 {
        case -99:
            _ = len(v1)
        case 128:
            fmt.Println(v2, v2)
        default:
            for fi73 := 0; fi73 < 9; fi73++ {
                var v74 string = ("\n\t" + "a")
                _ = fi73
                _ = v74
            }
        }
        for fi75 := 0; fi75 < 9; fi75++ {
            v2 = v2
            fmt.Println(v2, fi75)
            _ = fi75
        }
        fmt.Println(v2, v3)
    }
    fmt.Println(v3)
    switch (("a" + "") + "foo bar") {
    case "hello":
        var v76 float64 = v3
        _ = v76
    case "":
        var v77 uint8 = 0
        _ = v77
    default:
        fmt.Println(v3, v2)
    }
    var v78 float32 = 1e10
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v78
}
