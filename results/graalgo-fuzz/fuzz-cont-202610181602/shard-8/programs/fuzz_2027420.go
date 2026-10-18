package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int16; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 int; F2 int64 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32; F1 int64; F2 int }
type GenS10 struct { F0 int8; F1 uint; F2 uint }
type GenS11 struct { F0 uint; F1 string }
func genFn0(p0 uint8, p1 int32, p2 uint16) int16 {
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(fi5, p2, p2)
        _ = fi5
    }
    fmt.Println(p1, p0, p2)
    v6 := GenS1{F0: -128}
    var v7 bool = ((true || (100 < 0)) || false)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int8, p1 uint8) uint64 {
    v8 := GenS1{F0: p0}
    p1 -= (p1 - p1)
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    v1 := []uint32{2147483647, 42, 2}
    var v2 uint = 2
    v3 := []uint32{2147483647, 42, 0}
    v4 := make(map[int]string)
    switch "foo bar" {
    case "a":
        if (true && false) {
            var v9 int32 = 42
            fmt.Println(v9, v2)
            v9 -= (v9 - v9)
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 5; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                if (v2 != v2) {
                    switch "hello" {
                    case "a":
                        if false {
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                            if true {
                                v2 += v2
                            } else {
                                v2 += v2
                            }
                            v12 := GenS2{F0: 255, F1: genFn0(1, -2147483648, 42), F2: "foo bar"}
                            _ = v12
                        } else {
                            v13 := GenS3{F0: 40}
                            fmt.Println(fi10)
                            if (false || (false && false)) {
                                v14 := GenS2{F0: 42, F1: 32767, F2: "a"}
                                v2 -= (v2 << ((v2) & 63))
                                var v15 string = "\n\t"
                                var v16 int64 = -9223372036854775808
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                            switch ("foo bar" + ((("" + "") + "foo bar") + "\n\t")) {
                            case "foo bar":
                                v2 -= (v2 / ((680) | 1))
                            case "hello":
                                var v17 int64 = 381
                                _ = v17
                            case "a":
                                v18 := GenS1{F0: 100}
                                _ = v18
                            default:
                                if ("\n\t" > "foo bar") {
                                    v2 = (v2 & 735)
                                    if false {
                                        var v19 uint8 = 255
                                        _ = v19
                                    } else {
                                        fmt.Println(v2, fi11, fi11)
                                        if (!false) {
                                            switch 2147483647 {
                                            case 4294967295:
                                                v2 += 246
                                            case 0:
                                                var v20 float64 = 1.5
                                                _ = v20
                                            default:
                                                v2 = 42
                                                v21 := GenS1{F0: 127}
                                                var v22 float64 = 1e10
                                                _ = v21
                                                _ = v22
                                            }
                                            var v23 int = (fi11 << (((v2 + 0)) & 63))
                                            _ = v23
                                        }
                                        v24 := v4[(fi10 % (((fi11 & (fi10 - fi11))) | 1))]
                                        _ = v24
                                    }
                                } else {
                                    if true {
                                        if true {
                                            switch fi11 {
                                            case 32767:
                                                v4[(fi11 - 2147483647)] = "foo bar"
                                            }
                                        } else {
                                            if true {
                                                v2 += 178
                                                var v25 uint8 = 2
                                                var v26 bool = false
                                                _ = v25
                                                _ = v26
                                            } else {
                                                if (false || false) {
                                                    v27 := GenS4{F0: fi10}
                                                    _ = v27
                                                }
                                                v28 := v4[128]
                                                _ = v28
                                            }
                                            if ((1.5 - (0.1 - 3.14159)) != 0.1) {
                                                v29 := GenS2{F0: 159, F1: 0, F2: "hello"}
                                                _ = len(v4)
                                                _ = v29
                                            } else {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v30 := GenS4{F0: -9223372036854775808}
                                                    if false {
                                                        _ = len(v3)
                                                        var v31 int32 = 224
                                                        _ = v31
                                                    }
                                                    _ = v30
                                                default:
                                                    var v32 string = ""
                                                    _ = v32
                                                }
                                                v33 := GenS5{F0: 0, F1: fi10, F2: -9223372036854775808}
                                                v4[(fi11 % ((2147483647) | 1))] = ("a" + "\n\t")
                                                _ = v33
                                            }
                                        }
                                        switch (fi10 - 437) {
                                        case -366:
                                            switch (fi11 ^ fi10) {
                                            case 2147483647:
                                                _ = len(v1)
                                            default:
                                                if (1e10 < 0.1) {
                                                    v34 := GenS6{F0: -1, F1: 255}
                                                    if (!false) {
                                                        var v35 uint32 = 42
                                                        if false {
                                                            var v36 uint16 = 1
                                                            var v37 uint = 261
                                                            _ = len(v4)
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            v38 := GenS4{F0: (fi10 << ((v2) & 63))}
                                                            _ = v38
                                                        }
                                                        _ = v35
                                                    } else {
                                                        var v39 int64 = -387
                                                        v1 = append(v1, 42)
                                                        var v40 uint64 = 0
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                    var v41 int64 = 773
                                                    _ = v34
                                                    _ = v41
                                                }
                                            }
                                        default:
                                            var v42 uint64 = 0
                                            _ = v42
                                        }
                                        var v43 uint32 = 0
                                        var v44 bool = false
                                        var v45 string = "a"
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                }
                            }
                            _ = v13
                        }
                    }
                    fmt.Println(fi11)
                }
                v46 := GenS3{F0: -128}
                v2 -= 1
                v47 := GenS2{F0: 255, F1: int16(fi11), F2: ("hello" + "hello")}
                v1 = append(v1, 2147483648)
                _ = fi11
                _ = v46
                _ = v47
            }
            var v48 int32 = -1
            switch ("" + "foo bar") {
            case "":
                v49 := GenS7{F0: 33, F1: 0}
                _ = v49
            case "hello":
                if (!true) {
                    var v50 uint32 = 930
                    _ = v50
                } else {
                    for fi51 := 0; fi51 < 1; fi51++ {
                        switch ("foo bar" + (("a" + ("foo bar" + "\n\t")) + ("" + "foo bar"))) {
                        case "":
                            delete(v4, (fi10 << ((v2) & 63)))
                        default:
                            var v52 int16 = 32767
                            _ = v52
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi51
                    }
                    var v53 uint32 = 1
                    for fi54 := 0; fi54 < 2; fi54++ {
                        switch fi10 {
                        case 417:
                            var v55 int8 = 14
                            _ = v55
                        case -9223372036854775808:
                            if true {
                                var v56 int = -9223372036854775808
                                _ = v56
                            } else {
                                var v57 uint8 = 1
                                if ((true && true) && (0 < 772)) {
                                    switch (fi54 ^ -2147483648) {
                                    case 4294967295:
                                        if len(v3) > 0 {
                                            v3[0] = 4294967295
                                        }
                                    case 128:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                                _ = v57
                            }
                        case 9223372036854775807:
                            v58 := GenS8{F0: 65535}
                            _ = v58
                        default:
                            var v59 uint8 = 233
                            _ = v59
                        }
                        switch (fi10 & 0) {
                        case 4294967295:
                            v53 += v53
                        case -778:
                            v60 := GenS9{F0: 2147483647, F1: -338, F2: -952}
                            _ = v60
                        case 127:
                            if false {
                                v2 = (v2 + v2)
                            } else {
                                v61 := GenS10{F0: 100, F1: (v2 | 1), F2: (v2 / (((^(v2))) | 1))}
                                _ = len(v3)
                                var v62 uint16 = 255
                                _ = v61
                                _ = v62
                            }
                        default:
                            v2 = (v2 << ((v2) & 63))
                        }
                        if (false && true) {
                            v63 := GenS4{F0: fi10}
                            switch (fi54 >> (((^(v2))) & 63)) {
                            case 127:
                                if (-1 >= int16(v2)) {
                                    if true {
                                        v2 -= 186
                                        var v64 string = ""
                                        if (!(uint8(v2) != 223)) {
                                            if false {
                                                v65 := GenS11{F0: v2, F1: (v64 + (v64 + (v64 + v64)))}
                                                _ = len(v3)
                                                fmt.Println(fi10, v64)
                                                _ = v65
                                            } else {
                                                v66 := GenS1{F0: -59}
                                                _ = v66
                                            }
                                            v2 -= 729
                                        }
                                        _ = v64
                                    } else {
                                        v67 := GenS10{F0: -120, F1: v2, F2: v2}
                                        _ = v67
                                    }
                                    delete(v4, fi54)
                                    var v68 float32 = -2.25
                                    _ = v68
                                } else {
                                    if false {
                                        var v69 bool = false
                                        if v69 {
                                            v70 := v4[-9223372036854775808]
                                            var v71 uint = 0
                                            switch (fi54 / ((fi54) | 1)) {
                                            case -1:
                                                var v72 float32 = 1e10
                                                _ = v72
                                            case 32768:
                                                var v73 uint16 = 42
                                                _ = v73
                                            }
                                            _ = v70
                                            _ = v71
                                        } else {
                                            var v74 uint64 = 42
                                            _ = v74
                                        }
                                        var v75 string = ""
                                        _ = v69
                                        _ = v75
                                    } else {
                                        var v76 uint32 = (v53 * v53)
                                        _ = v76
                                    }
                                    v77 := GenS4{F0: fi10}
                                    _ = v77
                                }
                            case 694:
                                v48 -= v48
                            }
                            _ = v63
                        } else {
                            var v78 int32 = (-(v48))
                            _ = v78
                        }
                        _ = fi54
                    }
                    _ = v53
                }
            }
            _ = fi10
            _ = v48
        }
    }
    if (!(0 < 0)) {
        v79 := GenS9{F0: 788, F1: -9223372036854775808, F2: 466}
        _ = v79
    } else {
        var v80 uint16 = 865
        _ = v80
    }
    _ = len(v1)
    if ((!(!(!true))) && (!false)) {
        if ("foo bar" >= "a") {
            v2 = v2
            v2 -= (v2 * 2)
            _ = len(v4)
        }
    }
    v2 = v2
    var v81 float32 = 3.14159
    var v82 uint32 = 0
    v2 -= v2
    var v83 float32 = v81
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v81
    _ = v82
    _ = v83
}
