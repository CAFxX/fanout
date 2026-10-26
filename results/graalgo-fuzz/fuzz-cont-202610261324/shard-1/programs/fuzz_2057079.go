package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int64; F1 string; F2 uint }
type GenS6 struct { F0 int32; F1 uint32 }
type GenS7 struct { F0 string; F1 int64; F2 int32 }
type GenS8 struct { F0 int; F1 int32 }
type GenS9 struct { F0 int64; F1 int16; F2 int16 }

func main() {
    var v1 int = 128
    v2 := make(map[int]uint32)
    var v3 uint = 255
    var v4 float64 = 0.0
    var v5 int = v1
    var v6 int16 = -771
    v7 := GenS1{F0: 18446744073709551615, F1: true}
    var v8 string = "a"
    switch 4294967295 {
    case -505:
        for fi9 := 0; fi9 < 18; fi9++ {
            if true {
                if false {
                    var v10 int = fi9
                    _ = v10
                }
            } else {
                for fi11 := 0; fi11 < 20; fi11++ {
                    v6 = v6
                    v12 := []int{(-(v1)), (^(fi9))}
                    _ = fi11
                    _ = v12
                }
                if false {
                    var v13 float64 = 3.14159
                    for fi14 := 0; fi14 < 6; fi14++ {
                        if ((65535 < uint32(v3)) && true) {
                            switch ((v8 + (((("\n\t" + v8) + "") + v8) + "")) + v8) {
                            case "hello":
                                var v15 int = 822
                                _ = v15
                            default:
                                fmt.Println(fi14, v3)
                            }
                            fmt.Println(v4, v4, v6)
                            _ = len(v2)
                        } else {
                            if false {
                                v13 += v13
                            } else {
                                if (v5 < v1) {
                                    var v16 int16 = v6
                                    switch ("hello" + "\n\t") {
                                    case "":
                                        var v17 uint16 = 2
                                        _ = v17
                                    case "foo bar":
                                        fmt.Println(v1, v16)
                                    }
                                    _ = v16
                                } else {
                                    var v18 bool = (!(false || (-9223372036854775808 >= fi14)))
                                    _ = v18
                                }
                            }
                        }
                        v13 = ((v13 * (v4 + (1e10 - float64(float32(0.1))))) / (v4 - 1.5))
                        v19 := GenS2{F0: 100}
                        var v20 uint = v3
                        _ = fi14
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 6; fi21++ {
                        v22 := []int{fi9, (v5 + -1), fi9, v1}
                        v23 := v2[fi21]
                        var v24 float64 = ((0.1 * (float64(((float32(1.5) / 1.5) + (0.1 + float32((v4 * -2.25))))) * v13)) * ((float64((0.1 / 1.5)) + v4) * v4))
                        switch (fi21 | (fi9 >> ((41) & 63))) {
                        case 641:
                            switch v8 {
                            case "\n\t":
                                _ = len(v22)
                            case "":
                                if true {
                                    switch v8 {
                                    case "hello":
                                        if true {
                                            if (!true) {
                                                if len(v22) > 0 {
                                                    v22[0] = 609
                                                }
                                                v25 := GenS2{F0: 100}
                                                v1 -= fi21
                                                _ = v25
                                            }
                                            var v26 uint64 = 739
                                            v26 = 1
                                            _ = v26
                                        }
                                    case "a":
                                        _ = len(v22)
                                    case "":
                                        v23 = (v23 / ((2) | 1))
                                    default:
                                        v27 := GenS3{F0: 255, F1: (v6 + v6)}
                                        fmt.Println(v23, v1)
                                        _ = v27
                                    }
                                    _ = len(v2)
                                }
                            }
                        case -9223372036854775808:
                            v8 = "hello"
                        case -1:
                            v24 = ((0.1 + ((float64((float32(0.1) * (0.1 / 1.5))) - (((v24 / (v13 * v13)) - (3.14159 * v4)) + ((v13 + (v4 + 0.0)) + v24))) - v4)) + -2.25)
                        }
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = v13
                } else {
                    for fi28 := 0; fi28 < 16; fi28++ {
                        fmt.Println(fi28, v1, v3)
                        _ = fi28
                    }
                }
                if true {
                    if true {
                        var v29 uint16 = 65535
                        _ = v29
                    }
                } else {
                    v30 := GenS4{F0: "hello", F1: 78}
                    _ = v30
                }
            }
            if false {
                v31 := []int{(fi9 % ((v5) | 1)), (v5 | 9223372036854775807)}
                _ = v31
            } else {
                var v32 float32 = 1.5
                for fi33 := 0; fi33 < 15; fi33++ {
                    v8 = v8
                    _ = fi33
                }
                var v34 int64 = 9223372036854775807
                _ = v32
                _ = v34
            }
            _ = fi9
        }
    case 1:
        switch "" {
        case "\n\t":
            var v35 float64 = float64(1.5)
            _ = v35
        case "hello":
            for fi36 := 0; fi36 < 10; fi36++ {
                v37 := GenS3{F0: 14, F1: (v6 ^ v6)}
                v3 -= (v3 ^ v3)
                _ = fi36
                _ = v37
            }
        default:
            for fi38 := 0; fi38 < 3; fi38++ {
                v39 := []int{(fi38 - 128), fi38}
                var v40 uint16 = 32767
                v41 := GenS1{F0: v3, F1: (int(v3) >= int(v6))}
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
        }
    case 512:
        if true {
            for fi42 := 0; fi42 < 16; fi42++ {
                v4 = v4
                _ = fi42
            }
            var v43 uint64 = 42
            fmt.Println(v4, v3, v3)
            _ = v43
        }
    default:
        v44 := GenS5{F0: -762, F1: ((v8 + "\n\t") + v8), F2: 18446744073709551615}
        _ = v44
    }
    var v45 int16 = (v6 / (((v6 << ((v3) & 63))) | 1))
    var v46 float32 = float32((v4 - ((v4 * ((v4 - v4) * float64((1e10 * 1e10)))) / 1e10)))
    var v47 uint32 = 0
    var v48 float64 = v4
    v49 := GenS4{F0: ("hello" + ("foo bar" + "\n\t")), F1: 42}
    v3 = (v3 / (((v3 / ((4294967295) | 1))) | 1))
    var v50 int8 = 42
    if (!false) {
        for fi51 := 0; fi51 < 6; fi51++ {
            var v52 uint = (v3 | v3)
            v48 = v4
            var v53 uint64 = uint64(v47)
            _ = fi51
            _ = v52
            _ = v53
        }
        if false {
            v1 = int(v47)
        } else {
            if false {
                var v54 int16 = 638
                var v55 uint64 = 4294967295
                fmt.Println(v46)
                _ = v54
                _ = v55
            }
        }
        fmt.Println(v8)
    }
    v56 := GenS6{F0: 2147483647, F1: v47}
    v57 := GenS6{F0: 1, F1: v47}
    if true {
        v48 = (float64(v46) / 1.5)
        var v58 float64 = 3.14159
        if (!(!(v46 >= v46))) {
            v3 -= 36
        } else {
            v59 := GenS7{F0: v8, F1: 738, F2: 0}
            _ = v59
        }
        _ = v58
    }
    var v60 int32 = 2147483647
    if false {
        for fi61 := 0; fi61 < 9; fi61++ {
            if false {
                for fi62 := 0; fi62 < 10; fi62++ {
                    fmt.Println(v46)
                    _ = fi62
                }
            }
            _ = fi61
        }
        _ = len(v2)
    } else {
        if false {
            if false {
                fmt.Println(v48, v50, v8)
                var v63 bool = true
                _ = v63
            }
            v3 = v3
            v64 := GenS8{F0: (-(v5)), F1: v60}
            _ = v64
        }
        if true {
            var v65 int8 = 1
            fmt.Println(v6, v50)
            var v66 int8 = v50
            _ = v65
            _ = v66
        }
    }
    if true {
        v67 := GenS6{F0: -2147483648, F1: v47}
        _ = v67
    } else {
        v5 = 380
        for fi68 := 0; fi68 < 11; fi68++ {
            for fi69 := 0; fi69 < 6; fi69++ {
                fmt.Println(fi68, v48)
                _ = fi69
            }
            if ((v6 & -32768) < (v45 - (^(v45)))) {
                if false {
                    if (true && (0 < int16(v1))) {
                        switch v5 {
                        case 2147483647:
                            var v70 float64 = (1.5 * v48)
                            _ = v70
                        case 1:
                            v6 = 1000
                        case -421:
                            v48 += v4
                        }
                        var v71 int = v1
                        var v72 int16 = v45
                        _ = v71
                        _ = v72
                    } else {
                        if (!(false && (v3 <= v3))) {
                            switch v8 {
                            case "foo bar":
                                v73 := GenS3{F0: 32767, F1: (v6 / ((v6) | 1))}
                                _ = v73
                            case "hello":
                                switch -9223372036854775808 {
                                case 4294967295:
                                    v47 = (v47 - (v47 >> ((310) & 63)))
                                default:
                                    var v74 uint32 = (v47 | 42)
                                    _ = v74
                                }
                            }
                            var v75 uint16 = 0
                            for fi76 := 0; fi76 < 17; fi76++ {
                                v77 := GenS3{F0: v75, F1: v45}
                                switch (v5 ^ (fi76 % ((fi76) | 1))) {
                                case 42:
                                    v60 -= (v60 % ((v60) | 1))
                                case 127:
                                    v78 := GenS6{F0: v60, F1: (v47 + v47)}
                                    _ = v78
                                default:
                                    var v79 float32 = v46
                                    fmt.Println(v47, v48)
                                    _ = v79
                                }
                                _ = fi76
                                _ = v77
                            }
                            _ = v75
                        } else {
                            if (true || true) {
                                if false {
                                    var v80 float32 = ((float32(-2.25) / (1e10 + v46)) / v46)
                                    v81 := GenS2{F0: 1}
                                    v82 := []int{9223372036854775807, (v1 % (((-(v1))) | 1)), 128}
                                    _ = v80
                                    _ = v81
                                    _ = v82
                                }
                            } else {
                                var v83 uint64 = 19
                                switch "" {
                                case "\n\t":
                                    v46 -= 1e10
                                default:
                                    v84 := GenS9{F0: 9223372036854775807, F1: (v6 ^ (v45 + (v6 | -32768))), F2: (v45 ^ (v6 / (((v45 << ((v3) & 63))) | 1)))}
                                    _ = v84
                                }
                                _ = v83
                            }
                        }
                    }
                    v85 := []int{-499, 688, (v5 - (fi68 | (v5 + v1))), (v1 & 0)}
                    _ = v85
                } else {
                    if false {
                        v48 -= v48
                        for fi86 := 0; fi86 < 3; fi86++ {
                            var v87 int16 = (-(v6))
                            v88 := v2[(v5 | 233)]
                            _ = fi86
                            _ = v87
                            _ = v88
                        }
                    }
                    v89 := GenS8{F0: v5, F1: -1}
                    var v90 uint = v3
                    _ = v89
                    _ = v90
                }
                switch (v5 ^ v1) {
                case 546:
                    fmt.Println(v3)
                default:
                    fmt.Println(v60, v4)
                }
                var v91 uint8 = 133
                _ = v91
            } else {
                if (57 <= 767) {
                    var v92 uint16 = 65535
                    var v93 int16 = v6
                    _ = v92
                    _ = v93
                }
                for fi94 := 0; fi94 < 10; fi94++ {
                    var v95 int16 = 0
                    switch 22 {
                    case 127:
                        v96 := GenS7{F0: "hello", F1: 9223372036854775807, F2: v60}
                        if true {
                            var v97 int = (-(fi68))
                            v1 = (v5 - (^(v5)))
                            _ = v97
                        }
                        var v98 uint32 = 65535
                        _ = v96
                        _ = v98
                    default:
                        var v99 uint = 2
                        _ = v99
                    }
                    var v100 uint64 = 2
                    if true {
                        fmt.Println(v100)
                        v101 := GenS7{F0: "foo bar", F1: 9223372036854775807, F2: v60}
                        v46 = (v46 * (1.5 / v46))
                        if (432 >= 392) {
                            var v102 float64 = v4
                            _ = v102
                        }
                        _ = v101
                    } else {
                        var v103 uint64 = 893
                        _ = v103
                    }
                    v2[973] = (^(v47))
                    _ = fi94
                    _ = v100
                    _ = v95
                }
                v45 = v6
            }
            v104 := GenS7{F0: (((((v8 + v8) + ((v8 + "\n\t") + "")) + v8) + v8) + v8), F1: 1, F2: -18}
            _ = fi68
            _ = v104
        }
        fmt.Println(v1, v8)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v56
    _ = v57
    _ = v6
    _ = v60
    _ = v7
    _ = v8
}
