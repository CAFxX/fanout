package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 int16 }
func genFn0(p0 int64, p1 uint, p2 uint16) string {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int64 = (p0 - int64(fi6))
        v8 := GenS1{F0: 6, F1: 430, F2: -357}
        v7 = v7
        _ = fi6
        _ = v7
        _ = v8
    }
    p2 = p2
    var v9 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "a"
}

func main() {
    var v1 uint64 = 42
    v2 := []int{2147483647, 42, -41, -9223372036854775808}
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    v5 := []uint32{2147483648, 2147483647}
    for fi10 := 0; fi10 < 11; fi10++ {
        fmt.Println(v1)
        _ = fi10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(v1, v1)
        if (255 >= 231) {
            v12 := GenS2{F0: "", F1: 65535}
            _ = v12
        }
        _ = fi11
    }
    switch "" {
    case "hello":
        var v13 int16 = 42
        _ = v13
    case "":
        var v14 uint32 = 1
        _ = v14
    }
    v15 := GenS1{F0: 2, F1: 120, F2: 0}
    var v16 float32 = (float32(-2.25) + 0.0)
    var v17 int8 = 0
    fmt.Println(v1)
    if (127 > 254) {
        var v18 int32 = 2147483647
        var v19 float64 = 0.1
        _ = v18
        _ = v19
    }
    v5 = append(v5, 0)
    var v20 int64 = -9223372036854775808
    fmt.Println(v16, v1)
    v16 = v16
    if (float32(((float64(v16) * 3.14159) * -2.25)) >= ((((3.14159 + v16) - (float32((0.1 * 0.0)) + (1e10 / (v16 / 3.14159)))) * 3.14159) * (float32(-2.25) / (v16 - (((v16 / 0.1) + ((float32(-2.25) / 1.0) * ((float32(1e10) - (1e10 / float32(0.1))) / v16))) / 1e10))))) {
        var v21 int = 97
        fmt.Println(v17)
        _ = v21
    }
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    if true {
        var v22 uint = 732
        v20 += (v20 % ((v20) | 1))
        for fi23 := 0; fi23 < 15; fi23++ {
            switch ("" + "hello") {
            case "":
                var v24 uint16 = 65535
                _ = v24
            case "a":
                if (!true) {
                    if ((true && false) || (116 > 42)) {
                        var v25 uint = v22
                        _ = v25
                    } else {
                        fmt.Println(v20, v17, v17)
                        v16 = 0.1
                        var v26 bool = (32767 > 9223372036854775807)
                        _ = v26
                    }
                } else {
                    _ = len(v2)
                    if false {
                        if len(v5) > 0 {
                            v5[0] = 2147483648
                        }
                        for fi27 := 0; fi27 < 9; fi27++ {
                            var v28 uint16 = 2
                            _ = fi27
                            _ = v28
                        }
                    } else {
                        var v29 uint8 = 127
                        v16 = -2.25
                        _ = v29
                    }
                    if (false && true) {
                        switch ("" + "foo bar") {
                        case "hello":
                            var v30 int64 = v20
                            _ = v30
                        case "":
                            v16 = ((float32(3.14159) - float32(0.1)) * float32(1.5))
                        }
                        v31 := GenS3{F0: (v22 + v22), F1: 2147483647}
                        var v32 bool = true
                        _ = v31
                        _ = v32
                    } else {
                        switch "hello" {
                        case "hello":
                            if (("hello" + "") < ("" + "a")) {
                                v1 = (^(v1))
                                if false {
                                    v33 := v4[(-(fi23))]
                                    _ = v33
                                } else {
                                    var v34 int64 = (v20 + -9223372036854775808)
                                    v16 = v16
                                    v16 = v16
                                    _ = v34
                                }
                            } else {
                                for fi35 := 0; fi35 < 8; fi35++ {
                                    switch fi23 {
                                    case 911:
                                        var v36 uint = (v22 | v22)
                                        _ = v36
                                    case 2147483647:
                                        v22 = (v22 >> ((18446744073709551615) & 63))
                                    case 4294967295:
                                        v1 = (v1 * v1)
                                    default:
                                        var v37 int16 = -32768
                                        switch "\n\t" {
                                        case "foo bar":
                                            v3[(fi35 + 565)] = 2147483647
                                        case "a":
                                            if ((true || (0 <= v22)) && (1 >= 147)) {
                                                if (!false) {
                                                    v17 -= (^(v17))
                                                    v38 := GenS1{F0: 127, F1: (v37 << (((v22 ^ 74)) & 63)), F2: (v37 & (v37 | v37))}
                                                    var v39 int64 = -9223372036854775808
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v17 -= 70
                                                }
                                                v17 = v17
                                                v40 := GenS4{F0: v20, F1: v1, F2: 42}
                                                _ = v40
                                            }
                                        default:
                                            v41 := GenS5{F0: "hello"}
                                            _ = v41
                                        }
                                        v22 += v22
                                        v42 := GenS2{F0: "a", F1: 2}
                                        var v43 int32 = 0
                                        v43 = int32(v17)
                                        var v44 uint64 = (v1 + 0)
                                        v45 := GenS6{F0: v17, F1: v37}
                                        v46 := GenS5{F0: ""}
                                        _ = v37
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    }
                                    var v47 uint16 = 0
                                    switch fi35 {
                                    case 1:
                                        v16 += ((1.5 + 0.1) / 1.0)
                                        if (!(!(!((!((v1 * v1) != (^(v1)))) && false)))) {
                                            v4[fi35] = 2
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                if len(v5) > 0 {
                                                    v5[0] = 796
                                                }
                                            default:
                                                if len(v5) > 0 {
                                                    v5[0] = 0
                                                }
                                            }
                                            var v48 int8 = v17
                                            v16 -= (v16 - v16)
                                            switch "\n\t" {
                                            case "hello":
                                                var v49 uint32 = 943
                                                _ = v49
                                            }
                                            var v50 uint16 = (v47 + (v47 + (v47 << ((v22) & 63))))
                                            _ = v48
                                            _ = v50
                                        }
                                        v1 = (v1 * (^(v1)))
                                    case 227:
                                        var v51 uint64 = 42
                                        _ = v51
                                    default:
                                        if true {
                                            switch "" {
                                            case "a":
                                                var v52 int = fi23
                                                _ = v52
                                            case "\n\t":
                                                var v53 uint32 = 42
                                                var v54 float64 = -2.25
                                                _ = v53
                                                _ = v54
                                            default:
                                                switch ("a" + ("\n\t" + ("\n\t" + "hello"))) {
                                                case "a":
                                                    fmt.Println(fi23, fi23)
                                                case "foo bar":
                                                    var v55 uint8 = 1
                                                    _ = v55
                                                case "hello":
                                                    var v56 uint = v22
                                                    _ = v56
                                                }
                                            }
                                            v57 := v3[0]
                                            var v58 uint32 = (^(v57))
                                            _ = v57
                                            _ = v58
                                        }
                                    }
                                    _ = fi35
                                    _ = v47
                                }
                            }
                        case "foo bar":
                            for fi59 := 0; fi59 < 5; fi59++ {
                                delete(v3, fi59)
                                var v60 int64 = v20
                                v60 = (v20 & v20)
                                v22 = 222
                                _ = fi59
                                _ = v60
                            }
                        default:
                            var v61 string = "\n\t"
                            _ = v61
                        }
                    }
                }
            case "hello":
                v62 := GenS3{F0: v22, F1: 2147483647}
                _ = v62
            }
            _ = fi23
        }
        _ = v22
    } else {
        for fi63 := 0; fi63 < 20; fi63++ {
            v1 += (v1 - (v1 << ((42) & 63)))
            _ = fi63
        }
    }
    for fi64 := 0; fi64 < 8; fi64++ {
        fmt.Println(v1, v16)
        v1 = v1
        _ = fi64
    }
    v16 = v16
    var v65 float32 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v65
}
