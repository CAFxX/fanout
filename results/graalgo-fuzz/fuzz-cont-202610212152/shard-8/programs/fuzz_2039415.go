package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 int8 }
type GenS3 struct { F0 int64; F1 string; F2 uint32 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 int64; F1 uint8 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]string)
    var v3 string = (("" + "foo bar") + "")
    var v4 uint16 = 255
    v5 := make(map[string]string)
    fmt.Println(v3, v4)
    v3 = v3
    v1 += 0.1
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 bool = (18446744073709551615 >= 0)
        var v8 uint32 = 42
        switch (fi6 >> ((255) & 63)) {
        case 128:
            _ = len(v2)
        case 4294967295:
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 uint16 = (v4 >> ((18446744073709551615) & 63))
                v2[(fi9 / ((128) | 1))] = v3
                _ = fi9
                _ = v10
            }
        default:
            var v11 float64 = 1.5
            _ = v11
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        switch v3 {
        case "\n\t":
            v13 := []int{(fi12 << ((249) & 63))}
            _ = v13
        default:
            for fi14 := 0; fi14 < 14; fi14++ {
                if (251 < 153) {
                    var v15 int32 = 1
                    _ = v15
                }
                fmt.Println(fi12)
                var v16 uint = 1
                _ = fi14
                _ = v16
            }
        }
        var v17 int32 = 971
        v18 := v2[9223372036854775807]
        _ = fi12
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 834}
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS1{F0: 42}
        for fi22 := 0; fi22 < 11; fi22++ {
            v3 = v3
            v1 += (v1 - (1e10 / v1))
            v1 = v1
            _ = fi22
        }
        for fi23 := 0; fi23 < 11; fi23++ {
            v2[(fi20 & -927)] = v3
            switch ("a" + ("a" + "")) {
            case "\n\t":
                fmt.Println(v4, fi23, fi23)
            }
            if true {
                v3 = "\n\t"
                v24 := []int{fi23, fi23, fi23}
                v25 := GenS1{F0: 0}
                _ = v24
                _ = v25
            } else {
                v26 := GenS2{F0: 306, F1: 646, F2: 100}
                var v27 int32 = 0
                switch (fi23 | (fi23 + fi23)) {
                case 2147483647:
                    var v28 int = fi23
                    _ = v28
                case 421:
                    if false {
                        v29 := GenS2{F0: 8, F1: 2147483647, F2: 1}
                        _ = v29
                    }
                default:
                    if false {
                        var v30 int8 = 127
                        fmt.Println(v27, v3)
                        _ = v30
                    }
                }
                _ = v26
                _ = v27
            }
            delete(v5, v3)
            _ = fi23
        }
        _ = fi20
        _ = v21
    }
    switch 4294967295 {
    case -9223372036854775808:
        var v31 float32 = (v1 + v1)
        _ = v31
    case 1:
        fmt.Println(v4)
    case -159:
        v32 := GenS1{F0: 32767}
        _ = v32
    }
    fmt.Println(v1, v1, v1)
    if (v1 < v1) {
        var v33 uint16 = (^(v4))
        if (true && (0 <= 933)) {
            v34 := GenS2{F0: 601, F1: 1, F2: 127}
            switch v3 {
            case "hello":
                v3 = v3
            case "a":
                v35 := GenS1{F0: 42}
                _ = v35
            case "":
                var v36 int = 37
                _ = v36
            default:
                v37 := GenS3{F0: -958, F1: v3, F2: 74}
                _ = v37
            }
            _ = v34
        } else {
            var v38 float32 = 0.0
            _ = v38
        }
        if true {
            v33 -= (v4 % ((35) | 1))
            v39 := v2[128]
            _ = v39
        } else {
            fmt.Println(v33)
        }
        _ = v33
    } else {
        var v40 uint = 51
        fmt.Println(v1, v3, v40)
        _ = v40
    }
    if false {
        var v41 bool = true
        if true {
            v42 := GenS3{F0: 0, F1: "foo bar", F2: 2}
            v43 := GenS1{F0: 248}
            _ = v42
            _ = v43
        } else {
            var v44 string = (v3 + "hello")
            var v45 string = ("foo bar" + "hello")
            if true {
                var v46 bool = (0.0 <= 0.0)
                delete(v5, (v44 + "\n\t"))
                var v47 uint16 = (v4 >> ((0) & 63))
                _ = v46
                _ = v47
            } else {
                var v48 int8 = 0
                _ = v48
            }
            _ = v44
            _ = v45
        }
        _ = v41
    } else {
        v3 = v3
        v4 = 879
    }
    if (!(0 != -128)) {
        var v49 int8 = 100
        if true {
            fmt.Println(v1, v1, v1)
            fmt.Println(v1, v49)
        }
        _ = v49
    } else {
        switch 9223372036854775807 {
        case 2147483647:
            var v50 string = "a"
            _ = v50
        default:
            v51 := []int{-17, 32767, 2147483647}
            _ = v51
        }
        var v52 int = -393
        for fi53 := 0; fi53 < 19; fi53++ {
            var v54 uint16 = (v4 / (((v4 >> ((1) & 63))) | 1))
            if false {
                var v55 uint8 = 1
                switch (v3 + v3) {
                case "\n\t":
                    v3 = v3
                case "hello":
                    v56 := []int{fi53}
                    _ = v56
                case "a":
                    for fi57 := 0; fi57 < 5; fi57++ {
                        if true {
                            v58 := GenS1{F0: int16(v55)}
                            v55 = v55
                            var v59 int = v52
                            _ = v58
                            _ = v59
                        }
                        var v60 bool = (-91 > 0)
                        v61 := v5[("" + v3)]
                        _ = fi57
                        _ = v60
                        _ = v61
                    }
                }
                v52 -= 127
                _ = v55
            } else {
                if (18446744073709551615 <= 72) {
                    v3 = "foo bar"
                } else {
                    v1 -= (1e10 + (1e10 * v1))
                    v62 := GenS1{F0: 32767}
                    _ = v62
                }
            }
            var v63 float32 = 1.5
            _ = fi53
            _ = v54
            _ = v63
        }
        _ = v52
    }
    if (true || (-1 != -1)) {
        var v64 float64 = 1.5
        if (27 == 0) {
            v3 = (v3 + v3)
            var v65 int8 = 100
            _ = v65
        }
        v1 = (v1 + (((v1 + v1) / v1) / v1))
        _ = v64
    }
    var v66 int32 = 2147483647
    var v67 uint8 = 2
    if true {
        if true {
            if ((-2.25 * (3.14159 - 0.0)) < 1.5) {
                fmt.Println(v3, v1, v4)
                v2[32767] = v3
            } else {
                for fi68 := 0; fi68 < 17; fi68++ {
                    var v69 int16 = -32768
                    var v70 float32 = v1
                    _ = fi68
                    _ = v69
                    _ = v70
                }
                var v71 uint = 255
                if true {
                    v71 = v71
                } else {
                    switch 454 {
                    case -9223372036854775808:
                        for fi72 := 0; fi72 < 20; fi72++ {
                            switch (fi72 * (fi72 % (((fi72 * 503)) | 1))) {
                            case 9223372036854775807:
                                if true {
                                    v73 := GenS4{F0: (v66 ^ 331), F1: 42}
                                    _ = v73
                                } else {
                                    if (!false) {
                                        if (true && true) {
                                            v67 -= v67
                                            v74 := []int{int(v67), (fi72 & 1)}
                                            for fi75 := 0; fi75 < 18; fi75++ {
                                                v76 := GenS5{F0: 127}
                                                _ = fi75
                                                _ = v76
                                            }
                                            _ = v74
                                        } else {
                                            for fi77 := 0; fi77 < 20; fi77++ {
                                                var v78 int64 = 0
                                                v79 := GenS6{F0: (v78 & v78), F1: 0}
                                                var v80 uint8 = (v67 + (v67 & (v67 << ((255) & 63))))
                                                _ = fi77
                                                _ = v78
                                                _ = v79
                                                _ = v80
                                            }
                                            var v81 int16 = -594
                                            _ = len(v5)
                                            _ = v81
                                        }
                                        var v82 int8 = -29
                                        _ = v82
                                    } else {
                                        v2[fi72] = v3
                                    }
                                    v83 := GenS6{F0: -9223372036854775808, F1: 0}
                                    _ = v83
                                }
                            case 32768:
                                v84 := GenS7{F0: 9223372036854775807, F1: 2}
                                _ = v84
                            case 32767:
                                var v85 int8 = -128
                                _ = v85
                            default:
                                var v86 float64 = -2.25
                                _ = v86
                            }
                            _ = fi72
                        }
                    case 42:
                        for fi87 := 0; fi87 < 4; fi87++ {
                            fmt.Println(v1, v71, fi87)
                            _ = fi87
                        }
                    default:
                        var v88 string = v3
                        _ = v88
                    }
                }
                _ = v71
            }
            if true {
                v3 = (v3 + ("foo bar" + v3))
                if false {
                    var v89 bool = true
                    _ = v89
                }
                if (18446744073709551615 != 4294967295) {
                    var v90 uint16 = v4
                    _ = v90
                } else {
                    if (42 == -9223372036854775808) {
                        v91 := GenS2{F0: 42, F1: uint32(v67), F2: 100}
                        _ = v91
                    } else {
                        var v92 string = "\n\t"
                        var v93 int8 = -89
                        v3 = ""
                        _ = v92
                        _ = v93
                    }
                }
            }
        }
        for fi94 := 0; fi94 < 10; fi94++ {
            v4 = v4
            var v95 uint8 = v67
            _ = fi94
            _ = v95
        }
        fmt.Println(v4)
    } else {
        v96 := []int{9223372036854775807, 1, 0}
        _ = v96
    }
    var v97 int8 = 70
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v67
    _ = v97
}
