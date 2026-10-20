package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint64; F1 uint8; F2 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 int64 }
func genFn0() string {
    if (!true) {
        v7 := GenS1{F0: 2147483647}
        _ = v7
    }
    if true {
        v8 := GenS1{F0: -3}
        for fi9 := 0; fi9 < 15; fi9++ {
            switch "" {
            case "a":
                var v10 int32 = 112
                _ = v10
            case "foo bar":
                var v11 string = "\n\t"
                _ = v11
            case "hello":
                var v12 int8 = -128
                _ = v12
            }
            switch 1 {
            case 32767:
                var v13 int16 = 1
                _ = v13
            case -642:
                var v14 int8 = 21
                _ = v14
            case -2147483648:
                v15 := []int{(fi9 + (-(fi9))), (fi9 >> ((4294967295) & 63)), (fi9 * (fi9 >> ((4294967295) & 63))), fi9}
                _ = v15
            }
            _ = fi9
        }
        _ = v8
    }
    var v16 string = "a"
    var v17 int = -9223372036854775808
    var v18 bool = (2147483647 >= 28)
    _ = v16
    _ = v17
    _ = v18
    return ((("hello" + v16) + (v16 + "foo bar")) + v16)
}
func genFn1() uint32 {
    var v19 int32 = 998
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            if false {
                var v22 int = (fi21 | (fi20 + fi21))
                if true {
                    if (-2.25 < 3.14159) {
                        if true {
                            switch (fi20 - (fi21 + v22)) {
                            case 32768:
                                v23 := []int{(v22 ^ fi20)}
                                _ = v23
                            default:
                                if (true && (4294967295 < 921)) {
                                    var v24 uint32 = 1
                                    var v25 float32 = 0.0
                                    var v26 uint = 2
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                } else {
                                    var v27 uint16 = 32767
                                    switch fi21 {
                                    case 128:
                                        v22 -= (fi20 % (((v22 / ((127) | 1))) | 1))
                                    case 42:
                                        if true {
                                            var v28 int = v22
                                            if (!true) {
                                                v22 = 0
                                                v27 += (v27 / ((65535) | 1))
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v29 := GenS2{F0: 18446744073709551615, F1: (32767 <= 32767)}
                                                    _ = v29
                                                default:
                                                    v30 := []int{v28, fi20, fi21}
                                                    _ = v30
                                                }
                                                v31 := []int{(fi20 - (fi21 << ((736) & 63)))}
                                                _ = v31
                                            } else {
                                                if true {
                                                    if true {
                                                        fmt.Println(fi20, fi21)
                                                        v27 += 65535
                                                    }
                                                }
                                                var v32 string = ""
                                                _ = v32
                                            }
                                            _ = v28
                                        }
                                    case 9223372036854775807:
                                        var v33 int8 = -6
                                        _ = v33
                                    }
                                    if false {
                                        v34 := GenS1{F0: (v19 & (v19 - v19))}
                                        if true {
                                            fmt.Println(v19, fi20)
                                            v27 -= v27
                                        }
                                        var v35 uint8 = 168
                                        _ = v34
                                        _ = v35
                                    }
                                    if false {
                                        v36 := GenS3{F0: 911, F1: 127, F2: -1}
                                        if true {
                                            var v37 int64 = -469
                                            _ = v37
                                        } else {
                                            fmt.Println(v19, v27, fi21)
                                        }
                                        _ = v36
                                    } else {
                                        var v38 uint64 = 2
                                        if true {
                                            v22 = (fi21 << ((524) & 63))
                                            fmt.Println(v38)
                                            var v39 float64 = -2.25
                                            _ = v39
                                        }
                                        var v40 uint = 42
                                        _ = v38
                                        _ = v40
                                    }
                                    var v41 int64 = 864
                                    v27 -= v27
                                    v41 = v41
                                    _ = v27
                                    _ = v41
                                }
                            }
                        }
                        v42 := []int{(fi20 & fi20), fi21}
                        switch 0 {
                        case 42:
                            if (0 <= -122) {
                                v43 := GenS1{F0: 2147483647}
                                v42 = append(v42, 128)
                                v19 = v19
                                var v44 uint64 = 1
                                _ = len(v42)
                                _ = v43
                                _ = v44
                            }
                        default:
                            v45 := GenS3{F0: 4294967295, F1: 127, F2: fi21}
                            _ = v45
                        }
                        _ = v42
                    }
                    fmt.Println(v22, v19)
                    v46 := []int{52, (v22 + 9223372036854775807), v22, (v22 + 984)}
                    var v47 uint = 4294967295
                    if false {
                        v22 += fi20
                        fmt.Println(v22, v22, fi21)
                    }
                    _ = v46
                    _ = v47
                } else {
                    v48 := GenS4{F0: 0}
                    v49 := []int{1, -644, (fi21 & v22), -2147483648}
                    v22 = -2147483648
                    v22 += 0
                    v49 = append(v49, fi20)
                    switch "" {
                    case "":
                        v22 -= fi20
                    default:
                        var v50 uint8 = 128
                        _ = v50
                    }
                    _ = v48
                    _ = v49
                }
                _ = v22
            }
            var v51 bool = true
            _ = fi21
            _ = v51
        }
        v52 := []int{}
        v19 += (v19 | v19)
        _ = fi20
        _ = v52
    }
    v53 := []int{9223372036854775807, 1}
    if (127 <= 127) {
        var v54 float64 = (0.0 * 0.0)
        for fi55 := 0; fi55 < 13; fi55++ {
            v19 -= (v19 / (((v19 - (v19 % ((0) | 1)))) | 1))
            v19 += int32(fi55)
            var v56 float64 = v54
            _ = fi55
            _ = v56
        }
        _ = v54
    }
    _ = v19
    _ = v53
    return 65535
}

func main() {
    var v1 bool = (42 > 4294967295)
    var v2 int = -1
    v3 := make(map[uint32]string)
    var v4 uint16 = 0
    var v5 string = "a"
    var v6 int = v2
    if (1e10 >= 1e10) {
        var v57 int32 = 2147483647
        switch v5 {
        case "\n\t":
            var v58 bool = false
            _ = v58
        }
        switch (v2 - v6) {
        case 1:
            for fi59 := 0; fi59 < 20; fi59++ {
                if false {
                    v1 = ((v1 && false) || v1)
                    if v1 {
                        if (!(601 == 4294967295)) {
                            for fi60 := 0; fi60 < 14; fi60++ {
                                v57 -= v57
                                _ = fi60
                            }
                        } else {
                            v61 := GenS5{F0: -466, F1: v5}
                            for fi62 := 0; fi62 < 18; fi62++ {
                                v63 := v3[0]
                                v1 = v1
                                switch ((((v63 + v63) + (v5 + "a")) + "a") + "hello") {
                                case "hello":
                                    var v64 uint16 = 544
                                    _ = v64
                                case "":
                                    var v65 int64 = -1
                                    _ = v65
                                }
                                switch v5 {
                                case "a":
                                    switch ("a" + (("" + v5) + v5)) {
                                    case "foo bar":
                                        v66 := v3[0]
                                        _ = v66
                                    case "hello":
                                        var v67 uint = 4294967295
                                        _ = v67
                                    }
                                case "foo bar":
                                    v6 = v2
                                case "":
                                    var v68 int8 = -128
                                    _ = v68
                                }
                                v69 := []int{(fi59 % ((fi59) | 1)), -1, -1, 4294967295}
                                _ = fi62
                                _ = v63
                                _ = v69
                            }
                            _ = v61
                        }
                    }
                } else {
                    var v70 uint32 = 65535
                    _ = v70
                }
                _ = fi59
            }
        case -2147483648:
            if v1 {
                var v71 int8 = 48
                v57 += v57
                _ = v71
            } else {
                v5 = v5
                v72 := GenS1{F0: (v57 << ((1) & 63))}
                _ = v72
            }
        case -1:
            v73 := GenS4{F0: 614}
            _ = v73
        default:
            v1 = ((!(18446744073709551615 <= 381)) || false)
        }
        _ = v57
    } else {
        delete(v3, 65535)
    }
    if (-1 >= -203) {
        for fi74 := 0; fi74 < 7; fi74++ {
            fmt.Println(fi74, v4, v5)
            var v75 int32 = -1
            _ = fi74
            _ = v75
        }
        v76 := GenS2{F0: 1, F1: (1 != uint32(v4))}
        _ = v76
    }
    v1 = true
    v77 := []int{(v6 - (-(v2))), (^(v6)), (v6 << ((255) & 63))}
    v78 := v3[0]
    if v1 {
        if (-2147483648 == -1) {
            v79 := GenS6{F0: 9223372036854775807}
            if (-32768 > 666) {
                v80 := GenS2{F0: 2, F1: (!true)}
                v81 := v3[865]
                _ = v80
                _ = v81
            }
            v5 = "hello"
            _ = v79
        }
        for fi82 := 0; fi82 < 2; fi82++ {
            fmt.Println(v78, v4, v1)
            _ = len(v77)
            _ = fi82
        }
        for fi83 := 0; fi83 < 4; fi83++ {
            v78 = v5
            v84 := GenS3{F0: 641, F1: 188, F2: (^(v2))}
            switch ((v78 + v78) + v78) {
            case "":
                v85 := v3[1]
                _ = v85
            }
            _ = fi83
            _ = v84
        }
    }
    var v86 uint64 = 42
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v78:", v78)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v78
    _ = v86
}
