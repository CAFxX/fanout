package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int }
type GenS5 struct { F0 int8; F1 int; F2 int }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 uint; F1 uint; F2 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint64; F1 uint64 }
type GenS11 struct { F0 int; F1 uint32; F2 uint8 }
type GenS12 struct { F0 bool }

func main() {
    var v1 float64 = (3.14159 + 0.1)
    var v2 string = (((("" + ("" + ("hello" + "foo bar"))) + "a") + "foo bar") + "\n\t")
    v3 := make(map[string]int)
    var v4 int = 42
    v5 := make(map[int]int)
    var v6 uint16 = 2
    fmt.Println(v4, v4, v6)
    v6 = v6
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            if (-1 < 0) {
                switch v2 {
                case "":
                    var v9 float64 = 0.0
                    _ = v9
                case "a":
                    if false {
                        fmt.Println(v4)
                    }
                case "foo bar":
                    v10 := []int{(fi7 >> ((4294967295) & 63)), v4, fi7}
                    _ = v10
                }
                switch (v4 - (fi8 - -861)) {
                case 42:
                    var v11 uint64 = 0
                    _ = v11
                case -1:
                    v12 := []int{}
                    _ = v12
                case 16:
                    v13 := []int{v4, (v4 - v4)}
                    _ = v13
                default:
                    if (-274 > 417) {
                        var v14 string = v2
                        switch "\n\t" {
                        case "\n\t":
                            if (((0.0 * (0.0 / (3.14159 / ((0.1 / 1.5) + 0.1)))) >= 0.1) && (-1 == 42)) {
                                _ = len(v5)
                            }
                        case "":
                            fmt.Println(v6, v4)
                        case "a":
                            var v15 float64 = -2.25
                            _ = v15
                        }
                        if false {
                            switch v4 {
                            case 1:
                                v16 := GenS1{F0: (255 != 0), F1: 133}
                                _ = v16
                            case -17:
                                v6 = 32767
                            case 32767:
                                switch "" {
                                case "":
                                    v17 := GenS1{F0: false, F1: 1}
                                    _ = v17
                                case "hello":
                                    if (true && (!true)) {
                                        v18 := GenS2{F0: 682}
                                        v2 = "hello"
                                        v14 = v2
                                        v19 := []int{-975, (^(fi8)), -604, fi8}
                                        var v20 int16 = 42
                                        _ = len(v5)
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    }
                                case "\n\t":
                                    if false {
                                        switch fi8 {
                                        case -836:
                                            v21 := []int{}
                                            if false {
                                                var v22 uint8 = 0
                                                fmt.Println(v14, fi8)
                                                _ = v22
                                            } else {
                                                if true {
                                                    v2 = v2
                                                    v23 := GenS2{F0: 125}
                                                    var v24 int8 = 1
                                                    _ = v23
                                                    _ = v24
                                                } else {
                                                    if len(v21) > 0 {
                                                        v21[0] = (v4 >> ((1) & 63))
                                                    }
                                                    if (4294967295 == 1) {
                                                        if true {
                                                            v25 := GenS3{F0: 42, F1: 2, F2: 192}
                                                            var v26 uint64 = 1
                                                            if (v14 < (v14 + v14)) {
                                                                if true {
                                                                    v4 = -668
                                                                    var v27 uint8 = 1
                                                                    _ = v27
                                                                } else {
                                                                    v28 := GenS4{F0: 42, F1: v26, F2: (v4 / ((-9223372036854775808) | 1))}
                                                                    v6 = v6
                                                                    _ = v28
                                                                }
                                                                if (false || ((709 > 284) || (-9223372036854775808 <= 42))) {
                                                                    fmt.Println(v6, v4, v6)
                                                                    delete(v5, fi8)
                                                                }
                                                            }
                                                            var v29 uint16 = 160
                                                            _ = v25
                                                            _ = v26
                                                            _ = v29
                                                        }
                                                        var v30 uint16 = 255
                                                        var v31 uint32 = 2
                                                        _ = v30
                                                        _ = v31
                                                    } else {
                                                        fmt.Println(v4, v4, v4)
                                                        v32 := GenS5{F0: 100, F1: (v4 >> ((164) & 63)), F2: (v4 - (fi8 ^ 32767))}
                                                        _ = v32
                                                    }
                                                }
                                                if true {
                                                    v6 = v6
                                                    v6 -= v6
                                                    _ = len(v21)
                                                } else {
                                                    if false {
                                                        v21 = append(v21, (^(v4)))
                                                        fmt.Println(v2)
                                                        if (738 <= 2) {
                                                            v2 = "hello"
                                                            fmt.Println(fi7, v4)
                                                            switch (v2 + ((((v14 + (v2 + "a")) + "\n\t") + v14) + "a")) {
                                                            case "hello":
                                                                v33 := GenS6{F0: "\n\t", F1: 247}
                                                                _ = v33
                                                            }
                                                        }
                                                    } else {
                                                        var v34 uint16 = (v6 % ((344) | 1))
                                                        var v35 bool = false
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                    fmt.Println(v14, v2)
                                                    v14 = ("hello" + v2)
                                                }
                                            }
                                            var v36 int = fi7
                                            if true {
                                                switch 24 {
                                                case -1:
                                                    v1 = ((v1 * (0.0 / 1e10)) * ((-2.25 / 1.5) / 1e10))
                                                case 0:
                                                    fmt.Println(fi7, v6, v2)
                                                case 42:
                                                    v36 -= (fi8 / ((fi8) | 1))
                                                }
                                            }
                                            _ = v21
                                            _ = v36
                                        default:
                                            v37 := []int{fi8, 42}
                                            _ = v37
                                        }
                                        v38 := GenS7{F0: -32768, F1: 255}
                                        if (-9223372036854775808 >= fi8) {
                                            var v39 int8 = 66
                                            v40 := GenS8{F0: 2, F1: 65535, F2: v6}
                                            var v41 int8 = v39
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        } else {
                                            var v42 int8 = -11
                                            if false {
                                                v43 := GenS3{F0: 923, F1: 2, F2: 0}
                                                v44 := GenS3{F0: int64(v42), F1: uint32(v6), F2: 62}
                                                var v45 uint32 = 2147483648
                                                _ = v43
                                                _ = v44
                                                _ = v45
                                            } else {
                                                var v46 uint16 = 65535
                                                fmt.Println(fi7, v14)
                                                _ = v46
                                            }
                                            v47 := GenS3{F0: 684, F1: 846, F2: 0}
                                            v48 := GenS5{F0: v42, F1: 9223372036854775807, F2: (v4 >> ((487) & 63))}
                                            _ = v42
                                            _ = v47
                                            _ = v48
                                        }
                                        _ = v38
                                    } else {
                                        fmt.Println(fi8, v1)
                                        if true {
                                            _ = len(v3)
                                        }
                                    }
                                    v4 = 127
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        var v49 int = (fi8 | fi7)
                                        _ = v49
                                    case "\n\t":
                                        fmt.Println(fi8, v6)
                                    case "":
                                        fmt.Println(v6, v2, v14)
                                    }
                                default:
                                    var v50 uint = 550
                                    _ = v50
                                }
                            }
                        } else {
                            if true {
                                switch "\n\t" {
                                case "a":
                                    if (!(0 <= -128)) {
                                        var v51 uint16 = v6
                                        if (uint64(v6) != 0) {
                                            v52 := GenS9{F0: 65535}
                                            v53 := GenS4{F0: 780, F1: 4294967295, F2: v4}
                                            v14 = v2
                                            _ = v52
                                            _ = v53
                                        } else {
                                            v54 := []int{(^(v4)), (fi8 * 32768)}
                                            switch fi8 {
                                            case -72:
                                                var v55 int8 = -56
                                                _ = v55
                                            case 32768:
                                                if true {
                                                    if (!(true || true)) {
                                                        var v56 bool = false
                                                        var v57 int16 = -32768
                                                        var v58 uint8 = 2
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        v51 -= v6
                                                        if true {
                                                            if false {
                                                                var v59 float64 = 3.14159
                                                                v2 = v14
                                                                v60 := GenS10{F0: 1, F1: 18446744073709551615}
                                                                _ = v59
                                                                _ = v60
                                                            } else {
                                                                switch -609 {
                                                                case 914:
                                                                    v61 := GenS11{F0: (fi7 % (((fi8 - fi7)) | 1)), F1: 65535, F2: 1}
                                                                    _ = v61
                                                                }
                                                            }
                                                        }
                                                        v3[(v2 + v2)] = (fi7 % ((fi8) | 1))
                                                    }
                                                    fmt.Println(v14, v2)
                                                } else {
                                                    var v62 uint32 = 2147483648
                                                    _ = v62
                                                }
                                            case 127:
                                                v63 := GenS10{F0: 18446744073709551615, F1: 18446744073709551615}
                                                _ = v63
                                            }
                                            _ = v54
                                        }
                                        _ = v51
                                    }
                                case "hello":
                                    v64 := GenS4{F0: 952, F1: 211, F2: (^(fi8))}
                                    _ = v64
                                case "\n\t":
                                    v2 = v2
                                default:
                                    v65 := GenS5{F0: 127, F1: fi8, F2: fi7}
                                    _ = v65
                                }
                                if ((v14 + v14) == ("hello" + (("hello" + v14) + "a"))) {
                                    v66 := GenS12{F0: false}
                                    switch (v14 + v2) {
                                    case "foo bar":
                                        v2 = "hello"
                                    default:
                                        var v67 int16 = 0
                                        _ = v67
                                    }
                                    _ = v66
                                } else {
                                    v6 -= (v6 % ((32767) | 1))
                                    v14 = "\n\t"
                                    fmt.Println(v2)
                                }
                                if (2147483647 > 42) {
                                    if true {
                                        switch 173 {
                                        case -2147483648:
                                            switch fi8 {
                                            case 32768:
                                                v2 = "foo bar"
                                            }
                                        case 1:
                                            var v68 float32 = 3.14159
                                            _ = v68
                                        case -9223372036854775808:
                                            v69 := GenS7{F0: -515, F1: 4}
                                            _ = v69
                                        }
                                        var v70 int16 = 42
                                        fmt.Println(fi7, v6)
                                        _ = v70
                                    }
                                    v6 = 32767
                                    v1 -= float64(float32(0.0))
                                }
                            }
                            fmt.Println(v4, fi7)
                            var v71 bool = true
                            _ = v71
                        }
                        var v72 uint64 = 42
                        var v73 uint8 = 161
                        v74 := GenS5{F0: 94, F1: (fi8 * fi7), F2: int(v73)}
                        _ = v14
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                }
            }
            v75 := []int{127}
            _ = fi8
            _ = v75
        }
        for fi76 := 0; fi76 < 5; fi76++ {
            v2 = v2
            fmt.Println(v4)
            _ = fi76
        }
        _ = fi7
    }
    if false {
        v1 = (0.1 * (v1 / 0.1))
        var v77 string = v2
        var v78 uint64 = 4294967295
        _ = v77
        _ = v78
    } else {
        v1 = (v1 * float64(float32(-2.25)))
    }
    for fi79 := 0; fi79 < 12; fi79++ {
        for fi80 := 0; fi80 < 13; fi80++ {
            if true {
                var v81 uint = 813
                _ = v81
            }
            v2 = (v2 + v2)
            _ = fi80
        }
        _ = fi79
    }
    v2 = ("foo bar" + (v2 + (v2 + v2)))
    v82 := []int{v4, (-(v4))}
    var v83 uint16 = v6
    for fi84 := 0; fi84 < 5; fi84++ {
        var v85 string = "foo bar"
        _ = fi84
        _ = v85
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v82:", len(v82), v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v82
    _ = v83
}
