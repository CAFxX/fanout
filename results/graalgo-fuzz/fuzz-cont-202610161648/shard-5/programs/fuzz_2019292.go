package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int; F2 uint64 }
type GenS4 struct { F0 int; F1 int; F2 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 int16; F2 int32 }
type GenS7 struct { F0 uint16; F1 int8; F2 int16 }
type GenS8 struct { F0 string; F1 uint16 }
type GenS9 struct { F0 uint16; F1 int32; F2 int }
type GenS10 struct { F0 string; F1 int64 }
func genFn0() int64 {
    var v5 uint64 = 42
    var v6 bool = ((0 >= 1) && true)
    v7 := []int{-398, 82}
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1() string {
    v8 := GenS1{F0: 116}
    var v9 uint8 = 0
    v9 -= v9
    if true {
        var v10 uint = 18446744073709551615
        v11 := GenS1{F0: 32768}
        _ = v10
        _ = v11
    }
    _ = v8
    _ = v9
    return ("foo bar" + "")
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]bool)
    var v3 string = ((v1 + v1) + ("a" + v1))
    var v4 uint64 = 4294967295
    if false {
        if (false || true) {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 int = (^(fi12))
                v1 = ((v3 + (v3 + v3)) + v1)
                switch -764 {
                case -163:
                    var v14 bool = false
                    _ = v14
                case -2147483648:
                    var v15 uint8 = 2
                    _ = v15
                case -9223372036854775808:
                    var v16 float32 = (float32(-2.25) + -2.25)
                    _ = v16
                }
                _ = fi12
                _ = v13
            }
            switch 412 {
            case -248:
                for fi17 := 0; fi17 < 16; fi17++ {
                    delete(v2, 2147483648)
                    v18 := []int{fi17, 906}
                    _ = fi17
                    _ = v18
                }
            case 32768:
                v19 := []int{-529, 128}
                _ = v19
            }
            fmt.Println(v3, v4)
        } else {
            v20 := GenS2{F0: 32767}
            fmt.Println(v4, v3, v3)
            _ = v20
        }
        v21 := GenS3{F0: -2147483648, F1: -761, F2: (^(v4))}
        _ = len(v2)
        _ = v21
    } else {
        v22 := []int{}
        fmt.Println(v4)
        _ = v22
    }
    switch 9223372036854775807 {
    case -1:
        for fi23 := 0; fi23 < 8; fi23++ {
            if ((fi23 / ((46) | 1)) > (fi23 * fi23)) {
                v4 -= 42
                var v24 int8 = 127
                _ = v24
            } else {
                if false {
                    v25 := GenS4{F0: (fi23 << ((173) & 63)), F1: fi23, F2: (fi23 + (fi23 ^ (fi23 % ((fi23) | 1))))}
                    var v26 uint32 = 272
                    _ = v25
                    _ = v26
                } else {
                    switch "a" {
                    case "":
                        fmt.Println(v4, v1)
                    default:
                        v27 := v2[2147483648]
                        _ = v27
                    }
                }
                fmt.Println(fi23, v1, v1)
                if false {
                    var v28 int8 = 0
                    delete(v2, 2147483648)
                    for fi29 := 0; fi29 < 1; fi29++ {
                        fmt.Println(fi29)
                        _ = fi29
                    }
                    _ = v28
                } else {
                    var v30 int8 = -61
                    if (32767 != 1000) {
                        v2[4294967295] = true
                    } else {
                        v1 = ""
                        switch v3 {
                        case "":
                            v31 := GenS5{F0: v3}
                            _ = v31
                        case "a":
                            var v32 uint = 18446744073709551615
                            _ = v32
                        case "\n\t":
                            v33 := []int{}
                            _ = v33
                        default:
                            v34 := []int{-2147483648, (fi23 * fi23), 32767, fi23}
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 14; fi35++ {
                            switch fi35 {
                            case 4294967295:
                                switch fi35 {
                                case 2147483647:
                                    var v36 uint64 = (v4 * (v4 | 18446744073709551615))
                                    _ = v36
                                case 128:
                                    if false {
                                        v3 = v1
                                        delete(v2, 4294967295)
                                        v37 := GenS4{F0: fi23, F1: fi35, F2: fi35}
                                        _ = v37
                                    }
                                case -248:
                                    if (true && false) {
                                        var v38 int = (fi35 << ((148) & 63))
                                        v1 = ""
                                        v39 := []int{v38, (fi35 >> ((1) & 63)), fi35, -1}
                                        _ = v38
                                        _ = v39
                                    } else {
                                        var v40 string = "a"
                                        switch 883 {
                                        case 42:
                                            switch 665 {
                                            case 128:
                                                fmt.Println(v3, v1)
                                            default:
                                                if true {
                                                    switch v1 {
                                                    case "foo bar":
                                                        if true {
                                                            if false {
                                                                var v41 bool = (v30 < v30)
                                                                _ = v41
                                                            } else {
                                                                var v42 float32 = (3.14159 * 3.14159)
                                                                v1 = "\n\t"
                                                                _ = v42
                                                            }
                                                        } else {
                                                            var v43 uint = 1
                                                            var v44 string = (v3 + v3)
                                                            var v45 int64 = -30
                                                            _ = v43
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                    }
                                                    var v46 int32 = -1
                                                    v46 = v46
                                                    _ = v46
                                                }
                                            }
                                        case 127:
                                            if (!true) {
                                                var v47 int = fi35
                                                if true {
                                                    var v48 uint8 = 127
                                                    v47 -= (v47 & (fi23 | 0))
                                                    _ = v48
                                                }
                                                _ = v47
                                            } else {
                                                switch 0 {
                                                case 461:
                                                    var v49 uint8 = 127
                                                    _ = v49
                                                default:
                                                    v1 = ("a" + (((v1 + "") + v3) + v1))
                                                }
                                            }
                                        case 4294967295:
                                            if false {
                                                var v50 uint8 = 128
                                                _ = v50
                                            } else {
                                                v51 := GenS6{F0: 969, F1: 493, F2: 1}
                                                _ = v51
                                            }
                                        }
                                        _ = v40
                                    }
                                default:
                                    v52 := []int{(fi35 | fi35), 32767}
                                    _ = v52
                                }
                            case 2147483647:
                                if (((1e10 / 3.14159) == 0.1) && (true && false)) {
                                    if (v3 > v1) {
                                        var v53 bool = (!false)
                                        v54 := GenS7{F0: 2, F1: 100, F2: 42}
                                        switch 9223372036854775807 {
                                        case -478:
                                            v55 := GenS6{F0: 65535, F1: 0, F2: 42}
                                            _ = v55
                                        case 127:
                                            fmt.Println(fi23, v30)
                                        case -2147483648:
                                            fmt.Println(fi35)
                                        }
                                        _ = v53
                                        _ = v54
                                    } else {
                                        v30 = (v30 & -1)
                                        var v56 float64 = 1e10
                                        _ = v56
                                    }
                                }
                            }
                            _ = fi35
                        }
                    }
                    if false {
                        for fi57 := 0; fi57 < 10; fi57++ {
                            if true {
                                var v58 uint32 = 1
                                _ = v58
                            }
                            _ = fi57
                        }
                        var v59 uint64 = (v4 * 2)
                        v60 := GenS4{F0: fi23, F1: (-(fi23)), F2: (fi23 / (((fi23 + (fi23 * (fi23 ^ (fi23 << ((42) & 63)))))) | 1))}
                        _ = v59
                        _ = v60
                    } else {
                        if (255 <= 0) {
                            var v61 uint8 = 2
                            fmt.Println(v3, fi23, fi23)
                            v62 := []int{fi23, (fi23 % ((fi23) | 1)), fi23}
                            _ = v61
                            _ = v62
                        }
                        v30 -= 42
                    }
                    _ = v30
                }
            }
            v63 := GenS7{F0: 65535, F1: -1, F2: 0}
            _ = fi23
            _ = v63
        }
    case 9223372036854775807:
        if false {
            if (!(-1 >= -1)) {
                switch (v3 + "") {
                case "hello":
                    var v64 uint32 = 2
                    _ = v64
                default:
                    for fi65 := 0; fi65 < 13; fi65++ {
                        v66 := GenS4{F0: 32768, F1: 712, F2: fi65}
                        _ = fi65
                        _ = v66
                    }
                }
            } else {
                for fi67 := 0; fi67 < 5; fi67++ {
                    v68 := GenS7{F0: 1, F1: 42, F2: -786}
                    v1 = "a"
                    _ = fi67
                    _ = v68
                }
                v3 = v1
            }
        } else {
            if false {
                if (float32(((-2.25 + float64(0.1)) + 0.0)) <= float32(-2.25)) {
                    for fi69 := 0; fi69 < 17; fi69++ {
                        v4 += (v4 << ((4294967295) & 63))
                        _ = fi69
                    }
                    for fi70 := 0; fi70 < 14; fi70++ {
                        v4 = 0
                        var v71 uint32 = 1
                        _ = fi70
                        _ = v71
                    }
                }
                var v72 uint = 255
                _ = v72
            }
            v73 := GenS5{F0: v3}
            v74 := GenS8{F0: v3, F1: 0}
            _ = v73
            _ = v74
        }
    case 42:
        v3 = v1
    default:
        v1 = "\n\t"
    }
    for fi75 := 0; fi75 < 1; fi75++ {
        switch v1 {
        case "\n\t":
            v4 = (v4 % ((v4) | 1))
        case "foo bar":
            v4 = 1
        default:
            v1 = v1
        }
        fmt.Println(v4, v1)
        _ = fi75
    }
    var v76 int64 = 396
    switch 718 {
    case 2147483647:
        var v77 uint8 = 0
        _ = v77
    case -159:
        v4 = v4
    default:
        v76 = (^(v76))
    }
    var v78 uint = uint(v4)
    var v79 string = ("foo bar" + (v1 + (("" + ("hello" + v3)) + v3)))
    v2[2] = (false && true)
    for fi80 := 0; fi80 < 4; fi80++ {
        v81 := GenS9{F0: 42, F1: -2147483648, F2: (fi80 / ((9223372036854775807) | 1))}
        v82 := GenS10{F0: v79, F1: -1}
        _ = fi80
        _ = v81
        _ = v82
    }
    v79 = (v1 + "foo bar")
    for fi83 := 0; fi83 < 1; fi83++ {
        var v84 uint = v78
        _ = fi83
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v76
    _ = v78
    _ = v79
}
