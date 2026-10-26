package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint64; F1 int64; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint8, p1 uint) uint16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(p0)
        switch fi7 {
        case 1:
            fmt.Println(p1, p1)
        case 127:
            p0 -= 255
        }
        fmt.Println(p0, fi7)
        _ = fi7
    }
    if true {
        var v8 bool = true
        _ = v8
    } else {
        p1 = (p1 & (^(p1)))
    }
    var v9 uint = p1
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    var v1 int = -2147483648
    var v2 float32 = 1e10
    v3 := []int{}
    v4 := make(map[int]uint32)
    var v5 uint64 = 716
    var v6 string = "\n\t"
    var v10 uint = 270
    v6 = v6
    v11 := GenS1{F0: "a"}
    v12 := GenS1{F0: v6}
    v2 -= (v2 * (0.0 + ((float32(0.1) - 1.5) + 3.14159)))
    v13 := GenS2{F0: false, F1: 12}
    if (!(4294967295 <= v1)) {
        var v14 uint8 = 2
        var v15 int32 = 0
        v5 = (v5 - v5)
        _ = v14
        _ = v15
    } else {
        v16 := GenS1{F0: v6}
        v17 := GenS2{F0: ((1.5 <= 1e10) || false), F1: v10}
        _ = v16
        _ = v17
    }
    switch v6 {
    case "hello":
        v1 = (v1 / ((v1) | 1))
    case "a":
        switch (v1 << (((^(v10))) & 63)) {
        case 127:
            switch v1 {
            case 0:
                v18 := GenS3{F0: (^(v5)), F1: -795, F2: 785}
                _ = v18
            case 42:
                for fi19 := 0; fi19 < 1; fi19++ {
                    switch 58 {
                    case -9223372036854775808:
                        var v20 int64 = 582
                        _ = v20
                    case -2147483648:
                        var v21 uint = v10
                        _ = v21
                    case 786:
                        _ = len(v4)
                    }
                    _ = fi19
                }
            case -9223372036854775808:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v22 := GenS1{F0: v6}
                _ = v22
            }
        case 128:
            var v23 float64 = (3.14159 * (1.5 / 3.14159))
            _ = v23
        case 9223372036854775807:
            fmt.Println(v10, v5)
        }
    case "":
        if false {
            switch (v6 + (v6 + "foo bar")) {
            case "hello":
                fmt.Println(v6, v2)
            case "foo bar":
                v24 := GenS2{F0: (!(!(192 != 0))), F1: (^(v10))}
                _ = v24
            case "\n\t":
                var v25 uint = v10
                _ = v25
            }
        }
    default:
        if (!true) {
            var v26 int32 = 0
            _ = v26
        } else {
            v6 = ("\n\t" + "")
        }
    }
    v27 := GenS3{F0: (v5 | (v5 * v5)), F1: 0, F2: 1000}
    v28 := v4[127]
    if (0.1 == -2.25) {
        v1 -= 1
    }
    switch v1 {
    case 1:
        if false {
            var v29 int64 = 1
            if (true || (2147483648 < (v28 | v28))) {
                v1 += (v1 ^ (v1 / ((v1) | 1)))
            } else {
                var v30 int64 = v29
                for fi31 := 0; fi31 < 20; fi31++ {
                    v32 := GenS2{F0: (1 != 42), F1: uint(v5)}
                    _ = fi31
                    _ = v32
                }
                var v33 int32 = -2147483648
                _ = v30
                _ = v33
            }
            _ = v29
        }
    case -832:
        for fi34 := 0; fi34 < 7; fi34++ {
            v35 := GenS4{F0: 128, F1: 42, F2: 0}
            v36 := GenS5{F0: v5, F1: false, F2: v10}
            if true {
                if (fi34 == (fi34 / ((fi34) | 1))) {
                    var v37 uint = 4294967295
                    if true {
                        var v38 int32 = int32(v10)
                        v2 = v2
                        var v39 uint32 = v28
                        _ = v38
                        _ = v39
                    }
                    if (996 == 42) {
                        v28 -= v28
                    }
                    _ = v37
                }
            } else {
                v40 := v4[fi34]
                fmt.Println(v6, fi34, fi34)
                var v41 int64 = 0
                _ = v40
                _ = v41
            }
            _ = fi34
            _ = v35
            _ = v36
        }
    case 4294967295:
        if ((v6 + v6) > v6) {
            var v42 bool = false
            switch ("hello" + (v6 + ((v6 + ((v6 + (v6 + "\n\t")) + "a")) + v6))) {
            case "hello":
                for fi43 := 0; fi43 < 1; fi43++ {
                    switch (v1 | -9223372036854775808) {
                    case 257:
                        v42 = (589 >= -1)
                    case 32767:
                        v2 = 1e10
                    case 42:
                        if (fi43 < (-(fi43))) {
                            switch v6 {
                            case "foo bar":
                                if v42 {
                                    fmt.Println(v6)
                                    for fi44 := 0; fi44 < 6; fi44++ {
                                        v10 -= v10
                                        _ = fi44
                                    }
                                } else {
                                    v42 = (2 <= v28)
                                    if false {
                                        v2 = v2
                                        var v45 bool = v42
                                        _ = v45
                                    }
                                    var v46 bool = v42
                                    _ = v46
                                }
                            }
                            var v47 float64 = -2.25
                            _ = v47
                        } else {
                            switch 437 {
                            case 32768:
                                var v48 uint16 = 42
                                _ = v48
                            case 386:
                                var v49 int16 = 1
                                _ = v49
                            case 2147483647:
                                switch (^(fi43)) {
                                case -766:
                                    var v50 float64 = (3.14159 / -2.25)
                                    _ = v50
                                case -72:
                                    if false {
                                        _ = len(v3)
                                        _ = len(v4)
                                        v6 = (v6 + "foo bar")
                                    }
                                }
                            default:
                                if len(v3) > 0 {
                                    v3[0] = -975
                                }
                            }
                        }
                    }
                    v10 += uint(v5)
                    _ = fi43
                }
            case "\n\t":
                for fi51 := 0; fi51 < 15; fi51++ {
                    v42 = true
                    var v52 bool = v42
                    _ = fi51
                    _ = v52
                }
            }
            _ = v42
        }
    }
    switch v6 {
    case "hello":
        v5 -= 42
    default:
        if ((128 != 255) && false) {
            for fi53 := 0; fi53 < 12; fi53++ {
                for fi54 := 0; fi54 < 1; fi54++ {
                    if ((2 > 127) || ((fi53 * (v1 / ((208) | 1))) > fi53)) {
                        var v55 int64 = -714
                        v56 := GenS6{F0: 1}
                        _ = v55
                        _ = v56
                    } else {
                        v4[(fi54 | 979)] = 2147483648
                        var v57 uint64 = v5
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v57
                    }
                    switch fi53 {
                    case 365:
                        if (1.5 > 3.14159) {
                            switch (v6 + v6) {
                            case "hello":
                                fmt.Println(v1)
                            case "":
                                _ = len(v4)
                            default:
                                var v58 uint64 = 1
                                _ = v58
                            }
                            v3 = append(v3, -885)
                            v28 = v28
                        }
                    case 0:
                        v59 := GenS1{F0: v6}
                        _ = v59
                    }
                    if true {
                        v6 = ""
                        if true {
                            v60 := GenS2{F0: true, F1: v10}
                            var v61 int = (v1 ^ -2147483648)
                            var v62 bool = (-1 != 94)
                            _ = v60
                            _ = v61
                            _ = v62
                        }
                    } else {
                        if false {
                            v28 = 0
                            fmt.Println(fi54)
                            delete(v4, v1)
                            v63 := GenS6{F0: 1000}
                            _ = v63
                        } else {
                            if true {
                                v64 := v4[fi53]
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v64
                            }
                        }
                        var v65 uint8 = 217
                        if (!true) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v66 := GenS7{F0: (v65 / (((v65 / (((v65 & 0)) | 1))) | 1))}
                            _ = v66
                        } else {
                            var v67 int32 = -482
                            if (!(!(907 < v10))) {
                                v68 := GenS7{F0: v65}
                                var v69 uint16 = 2
                                _ = v68
                                _ = v69
                            } else {
                                switch (fi53 >> (((^(v10))) & 63)) {
                                case 2147483647:
                                    v70 := GenS8{F0: genFn0(v65, 1), F1: -1}
                                    _ = v70
                                case -638:
                                    v65 = (v65 + (v65 / (((^(v65))) | 1)))
                                }
                            }
                            var v71 float32 = v2
                            _ = v67
                            _ = v71
                        }
                        _ = v65
                    }
                    _ = fi54
                }
                v72 := GenS5{F0: v5, F1: false, F2: v10}
                _ = fi53
                _ = v72
            }
            if ((("hello" + "") + v6) > "") {
                var v73 uint8 = 1
                v4[v1] = v28
                _ = v73
            } else {
                fmt.Println(v2, v2, v6)
                switch "hello" {
                case "":
                    fmt.Println(v2)
                }
            }
        } else {
            var v74 bool = true
            v2 = v2
            if v74 {
                var v75 float32 = v2
                var v76 uint8 = 73
                _ = v75
                _ = v76
            }
            _ = v74
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
