package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int16; F2 int64 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16; F1 uint }
type GenS9 struct { F0 uint; F1 string }
type GenS10 struct { F0 int16; F1 uint8 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int64; F1 int32 }
type GenS13 struct { F0 string; F1 uint16; F2 int16 }
func genFn0(p0 int16) int16 {
    v7 := []int{-368, -1, -459, 0}
    fmt.Println(p0)
    v8 := GenS1{F0: 42, F1: false, F2: 782}
    _ = p0
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 uint64, p1 uint64, p2 string) bool {
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: 0, F1: (false && (false && true)), F2: 65535}
        _ = fi9
        _ = v10
    }
    switch 32767 {
    case 4294967295:
        p2 = p2
    case 128:
        if (-1 < -9223372036854775808) {
            v11 := GenS2{F0: 0, F1: 94}
            if (true || false) {
                if (65535 < 65535) {
                    v12 := []int{127, -2147483648}
                    var v13 int = 2147483647
                    _ = v12
                    _ = v13
                } else {
                    v14 := GenS1{F0: 1, F1: false, F2: 65535}
                    switch p2 {
                    case "\n\t":
                        var v15 float64 = (0.1 - (3.14159 + -2.25))
                        _ = v15
                    }
                    _ = v14
                }
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 uint16 = 0
                    _ = fi16
                    _ = v17
                }
            } else {
                var v18 float64 = -2.25
                var v19 uint64 = (p1 | p1)
                _ = v18
                _ = v19
            }
            if (((p2 + p2) + p2) == p2) {
                fmt.Println(p2, p0)
            }
            _ = v11
        } else {
            var v20 bool = true
            _ = v20
        }
    }
    v21 := GenS3{F0: 901}
    v22 := []int{4294967295, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return (-1 > 9223372036854775807)
}

func main() {
    var v1 int16 = 1000
    var v2 bool = true
    var v3 string = "\n\t"
    var v4 uint32 = 1
    v5 := []int{}
    var v6 bool = (1.5 != float64(1.5))
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 bool = false
        for fi25 := 0; fi25 < 6; fi25++ {
            var v26 uint16 = 860
            if true {
                var v27 uint = 65535
                _ = v27
            } else {
                var v28 bool = v6
                v29 := GenS4{F0: v1, F1: (^(v1)), F2: -57}
                _ = v28
                _ = v29
            }
            if (1.5 == (-2.25 + (float64((0.0 * (float32((((0.0 + (0.1 + (1.5 / (1e10 / 0.1)))) * (1e10 + ((((float64(float32(-2.25)) + 3.14159) / 0.1) + 3.14159) / 1e10))) - float64((0.1 * ((float32(-2.25) - ((1e10 + -2.25) / 1e10)) / float32(3.14159)))))) * 1e10))) / -2.25))) {
                var v30 bool = true
                v1 += (v1 << ((887) & 63))
                if v24 {
                    if (v24 && true) {
                        fmt.Println(v26)
                    }
                    v4 = uint32(v26)
                    v31 := GenS5{F0: 472, F1: int(v4)}
                    var v32 uint8 = 127
                    if len(v5) > 0 {
                        v5[0] = fi23
                    }
                    _ = v31
                    _ = v32
                }
                if len(v5) > 0 {
                    v5[0] = fi23
                }
                switch (v3 + "foo bar") {
                case "foo bar":
                    v33 := GenS6{F0: 42, F1: 0}
                    _ = v33
                case "hello":
                    v4 -= v4
                }
                _ = v30
            } else {
                if v6 {
                    var v34 int32 = 42
                    if ((716 != 1) || false) {
                        _ = len(v5)
                    }
                    var v35 uint = 73
                    _ = v34
                    _ = v35
                } else {
                    fmt.Println(v2, fi23, v26)
                }
            }
            switch 42 {
            case -9223372036854775808:
                v6 = v2
            case 2147483647:
                var v36 int = 32768
                _ = v36
            default:
                if false {
                    if v6 {
                        switch fi25 {
                        case -9223372036854775808:
                            v37 := GenS5{F0: 42, F1: (fi25 << ((42) & 63))}
                            _ = v37
                        case 2147483647:
                            switch (((v3 + v3) + (v3 + v3)) + v3) {
                            case "":
                                var v38 uint16 = (v26 >> ((255) & 63))
                                _ = v38
                            case "a":
                                v3 = "hello"
                            default:
                                if len(v5) > 0 {
                                    v5[0] = 9223372036854775807
                                }
                            }
                            if (v3 <= "") {
                                if false {
                                    fmt.Println(v24)
                                } else {
                                    v39 := GenS6{F0: 1, F1: (v4 << ((65535) & 63))}
                                    _ = v39
                                }
                            } else {
                                v40 := GenS1{F0: 2, F1: (!v24), F2: v26}
                                v5 = append(v5, fi23)
                                fmt.Println(fi23, fi23, v24)
                                _ = v40
                            }
                            fmt.Println(v2, fi25, v2)
                        default:
                            var v41 bool = (v4 < v4)
                            _ = v41
                        }
                        fmt.Println(v2, v24)
                    } else {
                        switch fi23 {
                        case 9223372036854775807:
                            v3 = (v3 + v3)
                        case 672:
                            var v42 int = fi25
                            _ = v42
                        }
                        fmt.Println(fi23, v24)
                        var v43 int8 = 100
                        _ = v43
                    }
                    if (((((-2.25 / float32(1e10)) / ((1e10 - float32(float64(((((((((0.1 * (1.5 - 1e10)) + 1e10) - 1.5) + (float32(0.0) - 1e10)) * -2.25) * 3.14159) * (3.14159 - -2.25)) / float32((-2.25 - (((1e10 / -2.25) / -2.25) - ((1e10 / 3.14159) / 1.0)))))))) - float32((1.5 + ((3.14159 * 1.5) / -2.25))))) / (0.1 * 1e10)) + 0.0) < 0.1) {
                        var v44 int8 = -128
                        if ((!v2) && v24) {
                            var v45 bool = (28 > 255)
                            var v46 uint64 = 1
                            _ = v45
                            _ = v46
                        }
                        _ = v44
                    }
                }
            }
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    for fi47 := 0; fi47 < 15; fi47++ {
        if v2 {
            v5 = append(v5, (fi47 + (fi47 | fi47)))
            _ = len(v5)
            v1 = (v1 ^ (v1 >> ((0) & 63)))
        } else {
            switch ((v3 + "foo bar") + v3) {
            case "foo bar":
                fmt.Println(v3)
            case "hello":
                v48 := GenS7{F0: 2147483647}
                _ = v48
            case "\n\t":
                if (true || (true || false)) {
                    v4 = v4
                    v2 = false
                }
            default:
                var v49 int = fi47
                _ = v49
            }
            fmt.Println(v3, v1)
            v50 := GenS8{F0: 364, F1: 42}
            _ = v50
        }
        switch v3 {
        case "hello":
            fmt.Println(v4, v2)
        case "\n\t":
            fmt.Println(fi47, v2, v6)
        case "foo bar":
            if (fi47 <= -9223372036854775808) {
                fmt.Println(v4, v6)
                if len(v5) > 0 {
                    v5[0] = (fi47 | (fi47 & (fi47 << ((65535) & 63))))
                }
            }
        }
        if len(v5) > 0 {
            v5[0] = fi47
        }
        _ = fi47
    }
    var v51 int16 = (v1 + genFn0(v1))
    _ = len(v5)
    for fi52 := 0; fi52 < 12; fi52++ {
        fmt.Println(v6, v3, v6)
        _ = fi52
    }
    switch v3 {
    case "hello":
        var v53 uint = 162
        _ = v53
    case "\n\t":
        _ = len(v5)
    case "":
        for fi54 := 0; fi54 < 9; fi54++ {
            var v55 int = 4294967295
            v56 := GenS9{F0: 65535, F1: "foo bar"}
            fmt.Println(v51, v2, v2)
            _ = fi54
            _ = v55
            _ = v56
        }
    default:
        for fi57 := 0; fi57 < 12; fi57++ {
            if v6 {
                v3 = (("" + "") + "\n\t")
            }
            if (!v6) {
                var v58 uint = 0
                _ = v58
            } else {
                fmt.Println(v6, v6)
                v59 := GenS5{F0: 2, F1: fi57}
                switch 1 {
                case -825:
                    for fi60 := 0; fi60 < 17; fi60++ {
                        v4 = v4
                        _ = fi60
                    }
                case 128:
                    var v61 int16 = 1000
                    _ = v61
                default:
                    var v62 float32 = float32(float64(float32((-2.25 / 0.1))))
                    _ = v62
                }
                _ = v59
            }
            var v63 int32 = -2147483648
            _ = fi57
            _ = v63
        }
    }
    var v64 int16 = genFn0((v51 >> ((2) & 63)))
    v64 = -1
    var v65 int8 = 100
    v66 := GenS7{F0: 42}
    if v6 {
        v6 = v2
    } else {
        var v67 int64 = 888
        v5 = append(v5, -2147483648)
        v67 -= (v67 / (((v67 ^ (v67 - 42))) | 1))
        _ = v67
    }
    if (1 > 857) {
        if (!true) {
            v6 = (!(!v6))
        } else {
            v68 := GenS1{F0: 128, F1: false, F2: 1}
            for fi69 := 0; fi69 < 1; fi69++ {
                v6 = (!v6)
                _ = fi69
            }
            var v70 uint64 = 42
            _ = v68
            _ = v70
        }
    } else {
        if ((false && (0 >= 32767)) || v6) {
            v71 := GenS10{F0: (v1 ^ (v51 - v51)), F1: 245}
            if true {
                switch ((v3 + "a") + v3) {
                case "foo bar":
                    switch v3 {
                    case "hello":
                        _ = len(v5)
                    case "a":
                        var v72 int64 = -9223372036854775808
                        _ = v72
                    default:
                        for fi73 := 0; fi73 < 3; fi73++ {
                            v74 := GenS10{F0: -1, F1: 255}
                            _ = fi73
                            _ = v74
                        }
                    }
                case "a":
                    switch "a" {
                    case "foo bar":
                        v75 := GenS11{F0: 2}
                        _ = v75
                    default:
                        if (false || (!(-119 > -62))) {
                            v76 := GenS12{F0: -142, F1: -965}
                            _ = v76
                        }
                    }
                default:
                    var v77 uint32 = (v4 * v4)
                    _ = v77
                }
                var v78 float32 = -2.25
                _ = v78
            } else {
                v79 := GenS6{F0: 2, F1: v4}
                _ = v79
            }
            _ = v71
        }
    }
    v80 := GenS9{F0: 239, F1: v3}
    if v2 {
        v81 := GenS13{F0: "a", F1: 0, F2: (v51 & v51)}
        _ = v81
    }
    if len(v5) > 0 {
        v5[0] = 32768
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v64
    _ = v65
    _ = v66
    _ = v80
}
