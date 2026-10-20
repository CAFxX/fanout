package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint8; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 uint32; F1 uint64; F2 string }
func genFn0(p0 uint64, p1 int32, p2 uint64) int {
    fmt.Println(p2, p0, p1)
    var v5 uint16 = 61
    p2 -= (p2 / ((p0) | 1))
    var v6 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint32) string {
    v7 := GenS1{F0: "foo bar", F1: p0}
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            if false {
                var v10 int32 = -2147483648
                switch "" {
                case "a":
                    fmt.Println(v10, v10)
                    p0 += p0
                case "\n\t":
                    var v11 int16 = 42
                    _ = v11
                case "hello":
                    v12 := []int{fi8, (fi9 % ((genFn0(912, (v10 + 181), 18446744073709551615)) | 1)), (-(fi9)), fi9}
                    _ = v12
                }
                if (4294967295 < p0) {
                    p0 = 17
                    v13 := GenS2{F0: p0, F1: 255}
                    _ = v13
                } else {
                    v14 := GenS1{F0: ("foo bar" + "foo bar"), F1: 42}
                    var v15 int32 = (^(v10))
                    _ = v14
                    _ = v15
                }
                _ = v10
            } else {
                var v16 uint32 = 165
                _ = v16
            }
            p0 += (p0 ^ p0)
            var v17 string = ("a" + "foo bar")
            if true {
                v18 := GenS2{F0: p0, F1: 16}
                _ = v18
            } else {
                p0 += p0
            }
            var v19 int = (fi8 % (((fi9 ^ (-(fi8)))) | 1))
            v20 := GenS3{F0: p0}
            _ = fi9
            _ = v17
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: "hello", F1: p0}
        _ = fi8
        _ = v21
    }
    p0 = 704
    switch 0 {
    case -1:
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 uint16 = 101
            switch fi22 {
            case -1:
                v24 := GenS4{F0: 2147483647, F1: 127, F2: "foo bar"}
                _ = v24
            case 32768:
                v25 := []int{171, (fi22 % (((fi22 & genFn0(42, -1, 0))) | 1)), (fi22 & fi22), genFn0(0, 1, 0)}
                _ = v25
            default:
                for fi26 := 0; fi26 < 7; fi26++ {
                    v27 := []int{}
                    _ = fi26
                    _ = v27
                }
            }
            fmt.Println(fi22, v23, p0)
            _ = fi22
            _ = v23
        }
    }
    _ = p0
    _ = v7
    return "a"
}

func main() {
    var v1 uint8 = 0
    var v2 uint16 = 221
    var v3 int32 = 42
    var v4 uint16 = (v2 - (v2 >> ((4294967295) & 63)))
    v28 := []int{9223372036854775807, 127, 42}
    v2 = v4
    _ = len(v28)
    switch 128 {
    case 4294967295:
        v4 += (^(v2))
    case 32767:
        if false {
            v2 -= (^(v2))
            if true {
                if true {
                    v29 := GenS5{F0: -32768}
                    if ((false && (float32((0.0 / 1.0)) <= ((1e10 + ((0.1 + ((0.0 + (1e10 / 0.1)) + 3.14159)) * (0.1 * ((0.1 - 1.5) / 1.5)))) / 0.1))) || false) {
                        var v30 int16 = 1
                        var v31 int16 = (v30 >> ((255) & 63))
                        v3 += v3
                        _ = v30
                        _ = v31
                    } else {
                        if true {
                            v1 = 127
                            var v32 int = genFn0(169, v3, 451)
                            _ = v32
                        } else {
                            for fi33 := 0; fi33 < 20; fi33++ {
                                if (42 > 42) {
                                    if (2147483647 != 1) {
                                        if (282 >= 65535) {
                                            var v34 uint32 = 4294967295
                                            _ = v34
                                        }
                                    } else {
                                        if false {
                                            v28 = append(v28, (fi33 | (^(fi33))))
                                        } else {
                                            var v35 float64 = (0.0 - 0.0)
                                            v28 = append(v28, (fi33 + (fi33 << ((2) & 63))))
                                            _ = len(v28)
                                            _ = v35
                                        }
                                        switch (^(fi33)) {
                                        case 987:
                                            for fi36 := 0; fi36 < 7; fi36++ {
                                                fmt.Println(v4, v4)
                                                v3 = v3
                                                v1 = v1
                                                _ = fi36
                                            }
                                        default:
                                            for fi37 := 0; fi37 < 11; fi37++ {
                                                if false {
                                                    var v38 string = ""
                                                    switch (fi37 & (fi37 / (((fi37 + fi37)) | 1))) {
                                                    case 128:
                                                        v1 -= v1
                                                    default:
                                                        v39 := GenS5{F0: 32767}
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            _ = len(v28)
                                                        case "hello":
                                                            fmt.Println(fi33, v38, fi37)
                                                        case "a":
                                                            fmt.Println(v38, fi33, fi33)
                                                        default:
                                                            var v40 int8 = -119
                                                            _ = v40
                                                        }
                                                        var v41 int = 5
                                                        _ = v39
                                                        _ = v41
                                                    }
                                                    var v42 int = (fi33 << ((42) & 63))
                                                    _ = v38
                                                    _ = v42
                                                }
                                                fmt.Println(v4)
                                                _ = fi37
                                            }
                                        }
                                        for fi43 := 0; fi43 < 7; fi43++ {
                                            var v44 int32 = (v3 >> ((42) & 63))
                                            v1 += 255
                                            v3 += v44
                                            _ = fi43
                                            _ = v44
                                        }
                                    }
                                } else {
                                    if false {
                                        v1 = (v1 | v1)
                                        v2 += 42
                                        for fi45 := 0; fi45 < 20; fi45++ {
                                            v2 = 42
                                            _ = fi45
                                        }
                                    }
                                    var v46 string = "\n\t"
                                    switch (fi33 & fi33) {
                                    case -9223372036854775808:
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                    default:
                                        var v47 int8 = -128
                                        _ = v47
                                    }
                                    _ = v46
                                }
                                var v48 uint32 = 18
                                _ = fi33
                                _ = v48
                            }
                            var v49 uint = 18446744073709551615
                            v2 += 32767
                            _ = v49
                        }
                        v50 := GenS3{F0: 1}
                        if ((v2 | (^(v2))) > v4) {
                            if false {
                                for fi51 := 0; fi51 < 18; fi51++ {
                                    var v52 int16 = 825
                                    _ = fi51
                                    _ = v52
                                }
                                v3 += -440
                                switch "" {
                                case "":
                                    var v53 bool = true
                                    _ = v53
                                case "foo bar":
                                    fmt.Println(v3)
                                case "a":
                                    var v54 int16 = int16(v1)
                                    _ = v54
                                }
                            }
                            for fi55 := 0; fi55 < 16; fi55++ {
                                var v56 uint = 65535
                                if (1.5 < -2.25) {
                                    if (false || (9223372036854775807 > 484)) {
                                        var v57 uint32 = 2
                                        for fi58 := 0; fi58 < 6; fi58++ {
                                            if ((v1 | uint8(v56)) >= 0) {
                                                switch fi58 {
                                                case -682:
                                                    var v59 uint32 = (v57 | v57)
                                                    _ = v59
                                                case 563:
                                                    v56 = v56
                                                default:
                                                    v3 = (v3 & (v3 % ((v3) | 1)))
                                                }
                                            } else {
                                                if false {
                                                    var v60 bool = false
                                                    var v61 int8 = 42
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    var v62 float64 = (0.0 + (float64((float32((0.0 * 1.5)) + 1e10)) * ((((3.14159 / 1.0) - (-2.25 * ((1.5 / 3.14159) - 1.5))) + ((0.0 / 3.14159) - 0.0)) + 0.0)))
                                                    _ = v62
                                                }
                                                var v63 int = fi55
                                                _ = v63
                                            }
                                            v56 -= uint(v1)
                                            v64 := GenS6{F0: -578, F1: ((v57 * (v57 - 2147483648)) <= v57)}
                                            _ = fi58
                                            _ = v64
                                        }
                                        var v65 int64 = -9223372036854775808
                                        _ = v57
                                        _ = v65
                                    } else {
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                    }
                                    if (!false) {
                                        if true {
                                            var v66 int32 = v3
                                            _ = v66
                                        }
                                        fmt.Println(v4)
                                        v56 = 4294967295
                                    }
                                }
                                _ = fi55
                                _ = v56
                            }
                        } else {
                            switch "hello" {
                            case "":
                                var v67 string = ""
                                _ = v67
                            case "foo bar":
                                if ((float32(0.0) / 1.0) == 3.14159) {
                                    _ = len(v28)
                                } else {
                                    v68 := GenS4{F0: 65535, F1: 76, F2: "hello"}
                                    _ = v68
                                }
                            }
                        }
                        _ = v50
                    }
                    _ = v29
                }
            } else {
                var v69 uint16 = 2
                fmt.Println(v3)
                switch "a" {
                case "":
                    v3 += -1
                case "foo bar":
                    var v70 int16 = -32768
                    _ = v70
                }
                _ = v69
            }
        }
    case 127:
        v71 := GenS4{F0: 853, F1: (v1 << ((4294967295) & 63)), F2: ("" + ("a" + "hello"))}
        _ = v71
    }
    var v72 int64 = 588
    switch -2147483648 {
    case 9223372036854775807:
        fmt.Println(v3, v3, v3)
    case -214:
        v73 := GenS2{F0: 2147483648, F1: v1}
        _ = v73
    }
    for fi74 := 0; fi74 < 19; fi74++ {
        for fi75 := 0; fi75 < 3; fi75++ {
            v2 -= v2
            var v76 float64 = 0.0
            v28 = append(v28, genFn0(4294967295, v3, 42))
            var v77 uint16 = v2
            _ = fi75
            _ = v76
            _ = v77
        }
        var v78 bool = true
        v79 := GenS7{F0: 2, F1: 4294967295, F2: "\n\t"}
        _ = fi74
        _ = v78
        _ = v79
    }
    var v80 int32 = 0
    for fi81 := 0; fi81 < 10; fi81++ {
        for fi82 := 0; fi82 < 11; fi82++ {
            _ = len(v28)
            _ = fi82
        }
        _ = fi81
    }
    fmt.Println(v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v72:", v72)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v72
    _ = v80
}
