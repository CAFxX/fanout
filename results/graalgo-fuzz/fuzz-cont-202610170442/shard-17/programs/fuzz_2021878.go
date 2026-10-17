package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS6 struct { F0 string; F1 int32; F2 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS10 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS11 struct { F0 int; F1 uint16 }
func genFn0() bool {
    v5 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: 0}
    var v6 int = 32768
    v6 = v6
    _ = v5
    _ = v6
    return (2 >= 2147483648)
}

func main() {
    var v1 int16 = -671
    v2 := make(map[string]int)
    var v3 int = -358
    v4 := make(map[int]uint32)
    v7 := []int{32767, v3, (v3 / (((v3 + v3)) | 1)), (v3 >> ((4294967295) & 63))}
    switch (v3 & v3) {
    case -589:
        var v8 uint64 = 2
        _ = v8
    case 42:
        v3 = v3
    default:
        v9 := GenS1{F0: 1, F1: 0, F2: 0}
        _ = v9
    }
    switch 0 {
    case 1:
        var v10 uint16 = 42
        _ = v10
    case -9223372036854775808:
        v11 := GenS2{F0: 1, F1: ("" + ("\n\t" + "\n\t"))}
        _ = v11
    case -44:
        for fi12 := 0; fi12 < 5; fi12++ {
            v1 += (v1 ^ (-(v1)))
            v3 = 2147483647
            v3 += v3
            _ = fi12
        }
    default:
        v4[v3] = 760
    }
    v3 -= (v3 + v3)
    var v13 uint32 = 930
    var v14 string = (("hello" + (("hello" + "\n\t") + "\n\t")) + "hello")
    v15 := GenS1{F0: 4294967295, F1: 916, F2: 961}
    switch ("foo bar" + v14) {
    case "a":
        v16 := GenS3{F0: (-(v1)), F1: 77}
        _ = v16
    case "\n\t":
        var v17 uint64 = 0
        _ = v17
    case "hello":
        v3 += v3
    default:
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int64 = 9223372036854775807
            var v20 uint32 = 0
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        if true {
            if len(v7) > 0 {
                v7[0] = int(v1)
            }
            v4[(fi21 | (fi21 ^ (v3 / ((v3) | 1))))] = 4294967295
        } else {
            fmt.Println(v13)
            if (255 < 226) {
                var v22 float64 = 0.0
                switch v14 {
                case "foo bar":
                    switch (v3 / (((v3 & fi21)) | 1)) {
                    case 127:
                        var v23 uint8 = 92
                        _ = v23
                    case 65:
                        for fi24 := 0; fi24 < 2; fi24++ {
                            v25 := GenS3{F0: v1, F1: 42}
                            switch -311 {
                            case -623:
                                fmt.Println(v1, v1)
                            case 0:
                                fmt.Println(fi21, v14, fi21)
                            default:
                                v13 = v13
                            }
                            fmt.Println(v22)
                            var v26 int = (v3 | (fi24 ^ (fi21 >> ((1) & 63))))
                            v27 := GenS2{F0: -1, F1: ("a" + (v14 + (v14 + (v14 + v14))))}
                            var v28 int = 42
                            v28 += (fi21 ^ v26)
                            var v29 float64 = (0.0 * v22)
                            var v30 uint = 65535
                            if true {
                                fmt.Println(v29, v29, fi24)
                                switch (v14 + v14) {
                                case "":
                                    if (2 != v30) {
                                        if len(v7) > 0 {
                                            v7[0] = (fi21 % ((v26) | 1))
                                        }
                                        _ = len(v7)
                                        if true {
                                            v31 := GenS1{F0: 42, F1: v30, F2: 1}
                                            _ = v31
                                        }
                                    }
                                case "hello":
                                    switch v14 {
                                    case "hello":
                                        var v32 int64 = 1
                                        _ = v32
                                    }
                                    switch (fi21 - (v28 & 4294967295)) {
                                    case 270:
                                        v33 := GenS4{F0: v1}
                                        _ = v33
                                    default:
                                        if false {
                                            fmt.Println(v14)
                                            switch v28 {
                                            case -700:
                                                var v34 bool = false
                                                _ = v34
                                            default:
                                                var v35 uint32 = (v13 | 285)
                                                _ = v35
                                            }
                                            var v36 int64 = 4
                                            if false {
                                                _ = len(v7)
                                                fmt.Println(fi21)
                                                v37 := GenS4{F0: v1}
                                                _ = v37
                                            } else {
                                                if false {
                                                    v26 -= (fi21 ^ fi21)
                                                    if false {
                                                        fmt.Println(v28)
                                                    } else {
                                                        v28 = fi21
                                                        v38 := GenS2{F0: 2, F1: "foo bar"}
                                                        _ = v38
                                                    }
                                                    switch (v3 / ((v28) | 1)) {
                                                    case 408:
                                                        delete(v4, fi21)
                                                        fmt.Println(fi24, v22)
                                                    }
                                                    if true {
                                                        fmt.Println(v30)
                                                    } else {
                                                        v2[v14] = (v3 % (((-(fi21))) | 1))
                                                        var v39 string = v14
                                                        v40 := GenS4{F0: int16(v13)}
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                } else {
                                                    v3 += fi24
                                                    if false {
                                                        var v41 uint32 = (v13 << (((v30 >> ((v30) & 63))) & 63))
                                                        _ = v41
                                                    }
                                                    v42 := GenS5{F0: 0, F1: (v36 ^ v36), F2: 2}
                                                    _ = v42
                                                }
                                                v29 = v22
                                            }
                                            if (((0.0 - 3.14159) != float32(v22)) && false) {
                                                var v43 int = v3
                                                if ((false || (!((v14 + "\n\t") < v14))) && false) {
                                                    var v44 string = (v14 + (v14 + "a"))
                                                    var v45 uint64 = 0
                                                    if (v29 > 0.0) {
                                                        var v46 uint64 = v45
                                                        var v47 uint = v30
                                                        v29 += (float64(0.1) + (float64(1e10) + (0.1 + (v22 + float64((0.0 - (3.14159 * (1e10 / 0.1))))))))
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    _ = v44
                                                    _ = v45
                                                } else {
                                                    v28 = fi21
                                                    var v48 bool = (v29 == v22)
                                                    _ = v48
                                                }
                                                v36 = 386
                                                delete(v4, fi21)
                                                v49 := GenS3{F0: -1, F1: -80}
                                                _ = v43
                                                _ = v49
                                            } else {
                                                v36 = (v36 * (^(v36)))
                                            }
                                            _ = v36
                                        } else {
                                            v26 -= fi24
                                            fmt.Println(fi21)
                                            var v50 int64 = -1
                                            _ = v50
                                        }
                                    }
                                }
                                v14 = v14
                            } else {
                                if ((!true) && true) {
                                    if ((3.14159 / 1.5) > float32((0.0 + float64((0.0 * (1e10 + 1.5)))))) {
                                        var v51 uint8 = 127
                                        v52 := v2[(v14 + "a")]
                                        v53 := GenS6{F0: (v14 + v14), F1: 0, F2: (v30 & (v30 << (((v30 - v30)) & 63)))}
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    } else {
                                        switch (v14 + ("" + ("foo bar" + v14))) {
                                        case "a":
                                            v54 := v4[9223372036854775807]
                                            _ = v54
                                        case "foo bar":
                                            v22 = v29
                                        case "":
                                            var v55 int = v28
                                            _ = v55
                                        }
                                        delete(v2, (v14 + "foo bar"))
                                    }
                                }
                                v2[((v14 + ("" + "\n\t")) + (("a" + (v14 + "")) + ((v14 + v14) + "\n\t")))] = (v26 + int(v13))
                            }
                            _ = fi24
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    default:
                        v56 := GenS1{F0: 218, F1: 42, F2: 0}
                        _ = v56
                    }
                case "hello":
                    var v57 int32 = 42
                    _ = v57
                case "a":
                    switch (v14 + "\n\t") {
                    case "":
                        if (0 != 1) {
                            v22 += ((float64(0.0) + -2.25) + -2.25)
                            for fi58 := 0; fi58 < 12; fi58++ {
                                _ = len(v7)
                                if (!((v13 - 42) <= v13)) {
                                    var v59 uint8 = 42
                                    _ = v59
                                } else {
                                    switch (v3 * 973) {
                                    case 1:
                                        var v60 uint64 = 86
                                        v13 = 2
                                        _ = v60
                                    case 32767:
                                        fmt.Println(v22)
                                    case 32768:
                                        switch v3 {
                                        case -145:
                                            fmt.Println(v14, v1, fi21)
                                        case 9223372036854775807:
                                            if true {
                                                v61 := GenS5{F0: 141, F1: -9223372036854775808, F2: (v13 & v13)}
                                                v14 = v14
                                                _ = v61
                                            } else {
                                                switch (v3 << ((656) & 63)) {
                                                case 32768:
                                                    var v62 float64 = v22
                                                    _ = v62
                                                }
                                                v1 = 1000
                                                v63 := GenS7{F0: 0}
                                                _ = v63
                                            }
                                        case 4294967295:
                                            var v64 int64 = 1
                                            _ = v64
                                        }
                                    default:
                                        v13 += v13
                                    }
                                    var v65 uint8 = 245
                                    _ = v65
                                }
                                var v66 string = ("" + "")
                                v66 = v14
                                _ = fi58
                                _ = v66
                            }
                            if false {
                                v67 := GenS3{F0: -106, F1: -2147483648}
                                v68 := GenS4{F0: (v1 ^ (v1 % (((v1 << ((18446744073709551615) & 63))) | 1)))}
                                _ = v67
                                _ = v68
                            } else {
                                v69 := GenS8{F0: 36, F1: -32768}
                                var v70 int = (v3 + 832)
                                fmt.Println(fi21)
                                _ = v69
                                _ = v70
                            }
                        }
                    case "hello":
                        switch (v3 ^ 42) {
                        case 686:
                            v71 := v4[fi21]
                            _ = v71
                        case -527:
                            for fi72 := 0; fi72 < 20; fi72++ {
                                var v73 uint16 = uint16(v13)
                                fmt.Println(v13, v3, v1)
                                _ = fi72
                                _ = v73
                            }
                        }
                    default:
                        v14 = (v14 + "a")
                    }
                }
                v4[v3] = (v13 | (v13 - (v13 << ((638) & 63))))
                _ = v22
            }
        }
        for fi74 := 0; fi74 < 9; fi74++ {
            var v75 int16 = (v1 / ((int16(v13)) | 1))
            fmt.Println(v75)
            if false {
                switch "a" {
                case "":
                    if true {
                        switch int(v13) {
                        case 49:
                            if false {
                                v1 -= (^(v75))
                                var v76 bool = true
                                _ = v76
                            } else {
                                if true {
                                    switch -1 {
                                    case 128:
                                        fmt.Println(v3)
                                    case 9223372036854775807:
                                        v14 = v14
                                    }
                                } else {
                                    var v77 float32 = 0.0
                                    if false {
                                        v1 -= 678
                                        v78 := GenS1{F0: 0, F1: 65535, F2: 4294967295}
                                        _ = v78
                                    }
                                    _ = v77
                                }
                                v13 = v13
                            }
                        case 128:
                            v79 := GenS9{F0: 1, F1: 0, F2: v1}
                            _ = v79
                        case 9223372036854775807:
                            switch (fi21 * v3) {
                            case -348:
                                switch "" {
                                case "hello":
                                    if (!(!((-2.25 >= (float64(1e10) * 1e10)) && (!true)))) {
                                        v80 := GenS6{F0: "\n\t", F1: -2147483648, F2: 65535}
                                        _ = v80
                                    }
                                case "foo bar":
                                    fmt.Println(v75)
                                }
                            case -2147483648:
                                var v81 uint16 = 2
                                _ = v81
                            default:
                                var v82 int16 = v75
                                _ = v82
                            }
                        }
                    }
                case "hello":
                    v14 = "\n\t"
                case "\n\t":
                    v83 := v4[(v3 * fi74)]
                    _ = v83
                }
                var v84 int32 = 0
                _ = v84
            }
            _ = fi74
            _ = v75
        }
        v85 := GenS10{F0: 42, F1: (^(v13)), F2: int64(v13)}
        _ = fi21
        _ = v85
    }
    v86 := GenS11{F0: 922, F1: 32767}
    var v87 int32 = 1
    for fi88 := 0; fi88 < 13; fi88++ {
        _ = len(v2)
        var v89 bool = (!(true && ((0 == 2) || true)))
        _ = fi88
        _ = v89
    }
    fmt.Println(v3, v3, v13)
    v90 := GenS5{F0: 1, F1: 42, F2: (v13 << ((799) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v86
    _ = v87
    _ = v90
}
