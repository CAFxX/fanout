package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 bool }
func genFn0(p0 string, p1 int64, p2 bool) int16 {
    if (!p2) {
        var v4 uint16 = 42
        p1 = 930
        _ = v4
    } else {
        v5 := []int{32767, 42}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(p1)
        _ = v5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 42, F1: 42, F2: 18446744073709551615}
            switch 42 {
            case 471:
                p2 = p2
            case -2147483648:
                fmt.Println(fi6, fi6)
            }
            _ = fi7
            _ = v8
        }
        var v9 int8 = -128
        fmt.Println(p2, p0, fi6)
        _ = fi6
        _ = v9
    }
    fmt.Println(p0)
    var v10 int = 36
    for fi11 := 0; fi11 < 19; fi11++ {
        switch "\n\t" {
        case "foo bar":
            if (!p2) {
                var v12 uint16 = 65535
                var v13 float32 = 1e10
                var v14 int16 = -32768
                _ = v12
                _ = v13
                _ = v14
            } else {
                p2 = (!((0.0 + (0.0 + (((float64(float32(0.0)) + 0.0) + ((1.5 + 1e10) + 3.14159)) / (((3.14159 - -2.25) / (-2.25 * 3.14159)) + 1e10)))) > ((3.14159 - 0.0) + 0.0)))
            }
        case "a":
            if p2 {
                p1 -= p1
                var v15 int16 = -103
                _ = v15
            } else {
                p0 = "foo bar"
                var v16 float64 = (0.0 * (3.14159 - 1.5))
                v17 := GenS2{F0: 2}
                _ = v16
                _ = v17
            }
        default:
            fmt.Println(p1, p1)
        }
        var v18 uint = 4294967295
        _ = fi11
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}
func genFn1(p0 int8, p1 int64, p2 uint) int64 {
    if true {
        v19 := []int{42}
        var v20 int8 = 1
        _ = v19
        _ = v20
    } else {
        switch "foo bar" {
        case "":
            p0 = (p0 ^ p0)
        }
    }
    if true {
        v21 := GenS3{F0: ("" + "foo bar"), F1: 924}
        _ = v21
    }
    if false {
        if false {
            for fi22 := 0; fi22 < 7; fi22++ {
                p0 = (p0 * 24)
                _ = fi22
            }
            var v23 uint16 = 241
            switch -321 {
            case 42:
                switch 1 {
                case -742:
                    v24 := []int{9223372036854775807, 128, 838}
                    _ = v24
                case 9223372036854775807:
                    var v25 uint8 = 255
                    _ = v25
                case -1:
                    for fi26 := 0; fi26 < 10; fi26++ {
                        v27 := []int{}
                        if true {
                            for fi28 := 0; fi28 < 14; fi28++ {
                                var v29 uint32 = 0
                                _ = fi28
                                _ = v29
                            }
                            switch 9223372036854775807 {
                            case 42:
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    var v31 int8 = p0
                                    _ = fi30
                                    _ = v31
                                }
                            default:
                                v32 := GenS2{F0: 141}
                                _ = v32
                            }
                            switch "a" {
                            case "a":
                                v33 := GenS4{F0: 0, F1: 2147483647}
                                _ = v33
                            case "\n\t":
                                for fi34 := 0; fi34 < 14; fi34++ {
                                    switch fi26 {
                                    case -9223372036854775808:
                                        v35 := GenS1{F0: 255, F1: -2147483648, F2: 42}
                                        _ = v35
                                    }
                                    _ = fi34
                                }
                            case "hello":
                                switch "" {
                                case "a":
                                    switch "" {
                                    case "a":
                                        p1 += p1
                                    case "\n\t":
                                        var v36 uint32 = 2147483648
                                        _ = v36
                                    }
                                case "\n\t":
                                    p0 -= p0
                                case "hello":
                                    for fi37 := 0; fi37 < 1; fi37++ {
                                        if true {
                                            p0 += p0
                                        } else {
                                            p0 = 127
                                        }
                                        var v38 uint64 = 828
                                        _ = fi37
                                        _ = v38
                                    }
                                default:
                                    p1 = 811
                                }
                            default:
                                p0 = (p0 - p0)
                            }
                        }
                        _ = fi26
                        _ = v27
                    }
                default:
                    switch 0 {
                    case 4294967295:
                        switch "hello" {
                        case "\n\t":
                            var v39 float32 = (0.0 + 0.1)
                            _ = v39
                        case "":
                            if (false || (32767 <= 575)) {
                                var v40 uint = (^(p2))
                                var v41 uint64 = 42
                                var v42 uint64 = v41
                                _ = v40
                                _ = v41
                                _ = v42
                            } else {
                                var v43 bool = true
                                fmt.Println(v23, p2)
                                var v44 int = 32767
                                _ = v43
                                _ = v44
                            }
                        case "hello":
                            var v45 int16 = genFn0("", (p1 | (p1 << ((p2) & 63))), (0 < 1))
                            _ = v45
                        }
                    case 2147483647:
                        fmt.Println(p0, v23, p1)
                    }
                }
            case 1:
                switch "hello" {
                case "a":
                    var v46 uint32 = 954
                    _ = v46
                case "foo bar":
                    var v47 float64 = 0.0
                    _ = v47
                case "\n\t":
                    p0 = p0
                }
            default:
                v48 := []int{2147483647}
                _ = v48
            }
            _ = v23
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return int64(p0)
}

func main() {
    var v1 int16 = -766
    var v2 uint64 = 42
    var v3 string = ("hello" + ("\n\t" + ("hello" + ("" + "\n\t"))))
    for fi49 := 0; fi49 < 12; fi49++ {
        v50 := []int{}
        fmt.Println(v1, v1)
        _ = fi49
        _ = v50
    }
    v2 = 18446744073709551615
    var v51 int64 = -1
    var v52 uint16 = 42
    v53 := GenS2{F0: 1}
    for fi54 := 0; fi54 < 19; fi54++ {
        switch 209 {
        case 256:
            var v55 uint64 = v2
            _ = v55
        case 32767:
            if (3.14159 < ((float64(3.14159) / 3.14159) * 0.0)) {
                fmt.Println(v52, v51)
            }
        }
        v56 := []int{fi54, fi54}
        for fi57 := 0; fi57 < 13; fi57++ {
            v56 = append(v56, (fi54 << ((446) & 63)))
            switch fi57 {
            case -1:
                fmt.Println(v52, fi57)
            default:
                if ("a" >= ((((v3 + ((v3 + v3) + (v3 + "hello"))) + v3) + v3) + ("" + ((v3 + v3) + ((v3 + (v3 + v3)) + v3))))) {
                    v56 = append(v56, 42)
                }
            }
            if len(v56) >= 2 {
                v56 = v56[1:len(v56)]
            }
            var v58 float64 = 0.0
            _ = fi57
            _ = v58
        }
        _ = fi54
        _ = v56
    }
    v51 = genFn1(100, v51, 952)
    var v59 int32 = 1
    var v60 uint = uint(v52)
    var v61 int8 = 108
    v62 := GenS5{F0: v60}
    var v63 bool = true
    switch (v3 + v3) {
    case "\n\t":
        v59 -= (v59 + v59)
    case "":
        if v63 {
            var v64 string = ""
            fmt.Println(v61)
            fmt.Println(v60, v64)
            _ = v64
        } else {
            switch "hello" {
            case "hello":
                for fi65 := 0; fi65 < 17; fi65++ {
                    switch int(v52) {
                    case 9223372036854775807:
                        if v63 {
                            v66 := []int{4294967295, (fi65 - fi65)}
                            v60 += uint(v52)
                            _ = v66
                        }
                    case 195:
                        v67 := GenS4{F0: (v59 << ((v60) & 63)), F1: 401}
                        _ = v67
                    }
                    v68 := GenS6{F0: v51, F1: (!(!(!v63)))}
                    v60 += 65535
                    _ = fi65
                    _ = v68
                }
            }
            v59 -= (^(v59))
            if (!(!v63)) {
                fmt.Println(v63, v2)
                v61 = (v61 >> ((2) & 63))
                fmt.Println(v1)
            } else {
                v69 := []int{int(v51), 805}
                _ = v69
            }
        }
    }
    var v70 uint32 = 1
    if ((false && false) || v63) {
        for fi71 := 0; fi71 < 9; fi71++ {
            v72 := []int{}
            v59 = v59
            if (285 >= 2) {
                if len(v72) >= 2 {
                    v72 = v72[1:len(v72)]
                }
                v61 = int8(v52)
                if v63 {
                    fmt.Println(v59, v1)
                } else {
                    var v73 uint8 = 255
                    v61 = (v61 ^ int8(v60))
                    if v63 {
                        v74 := GenS3{F0: (v3 + (("foo bar" + v3) + "")), F1: v70}
                        _ = len(v72)
                        _ = v74
                    } else {
                        v59 = int32(v2)
                        if (0.1 <= (1e10 * 3.14159)) {
                            if len(v72) > 0 {
                                v72[0] = (^(fi71))
                            }
                        }
                    }
                    _ = v73
                }
            } else {
                v59 += 0
                var v75 uint64 = (v2 >> ((710) & 63))
                _ = v75
            }
            _ = fi71
            _ = v72
        }
        fmt.Println(v51, v63, v61)
        fmt.Println(v52, v63, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v53
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v70
}
