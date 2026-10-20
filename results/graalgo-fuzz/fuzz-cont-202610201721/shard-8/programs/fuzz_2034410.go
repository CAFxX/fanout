package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int16; F1 int16; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 int64; F2 bool }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int; F1 string; F2 int }
type GenS9 struct { F0 int8; F1 int; F2 int8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 string; F1 int; F2 int8 }
func genFn0(p0 bool) int64 {
    fmt.Println(p0)
    p0 = ((1 <= -9223372036854775808) && (!p0))
    v5 := []int{0}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 0
    var v3 float64 = 0.0
    var v4 bool = (!(false || false))
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint32 = 2
        v8 := GenS1{F0: v7, F1: 4294967295}
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v3, v2, v3)
    v9 := GenS2{F0: -187, F1: 0, F2: 542}
    fmt.Println(v4, v3)
    for fi10 := 0; fi10 < 17; fi10++ {
        if v4 {
            if false {
                v3 = (3.14159 - v3)
                var v11 int8 = -37
                switch fi10 {
                case -1:
                    v12 := GenS2{F0: -32768, F1: 32767, F2: fi10}
                    _ = v12
                default:
                    var v13 bool = (((("foo bar" + "\n\t") + "hello") + "foo bar") > "hello")
                    _ = v13
                }
                _ = v11
            } else {
                v4 = ((int8(fi10) != 1) || v4)
            }
            fmt.Println(fi10)
        } else {
            v3 = ((v3 + (((((v3 + ((v3 + 3.14159) - v3)) / v3) + v3) - v3) / 1e10)) - float64(0.0))
            v14 := GenS2{F0: 0, F1: 1000, F2: -209}
            v15 := []int{-686, 537, 32768}
            _ = v14
            _ = v15
        }
        _ = fi10
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        if v4 {
            v4 = false
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 uint16 = 1
                v18 = (v18 % ((v18) | 1))
                if (!(((v2 << ((2) & 63)) >= v2) && true)) {
                    v19 := GenS3{F0: -390}
                    _ = v19
                } else {
                    switch 32767 {
                    case 1:
                        var v20 string = ""
                        _ = v20
                    case 4294967295:
                        if v4 {
                            v21 := []int{fi16, -948, (-(fi17)), fi17}
                            _ = v21
                        }
                    case 128:
                        fmt.Println(v4)
                    }
                    var v22 float64 = ((((v3 * (v3 + v3)) + (0.1 / v3)) / 1.5) - 1e10)
                    _ = v22
                }
                _ = fi17
                _ = v18
            }
        } else {
            for fi23 := 0; fi23 < 3; fi23++ {
                v2 -= 42
                v4 = false
                switch ("hello" + "a") {
                case "hello":
                    v2 -= v2
                    v24 := []int{532, fi16}
                    _ = v24
                case "\n\t":
                    v25 := []int{fi23, fi23, (fi16 & fi23), fi23}
                    _ = v25
                case "":
                    fmt.Println(v2)
                default:
                    switch (fi16 << ((442) & 63)) {
                    case -9223372036854775808:
                        if v4 {
                            if (v3 > v3) {
                                v2 = 42
                                v26 := GenS1{F0: 334, F1: (v2 * v2)}
                                if (v4 || (2 == 65535)) {
                                    var v27 uint64 = v2
                                    fmt.Println(v2)
                                    v4 = (true || (-9223372036854775808 > 9223372036854775807))
                                    _ = v27
                                }
                                _ = v26
                            } else {
                                var v28 uint8 = 1
                                _ = v28
                            }
                            v1[4294967295] = ("foo bar" + "foo bar")
                            var v29 uint = 255
                            v30 := GenS1{F0: 2, F1: v2}
                            _ = v29
                            _ = v30
                        }
                    default:
                        v3 -= v3
                    }
                }
                _ = fi23
            }
            if v4 {
                v31 := GenS1{F0: 1, F1: 18446744073709551615}
                switch 523 {
                case 4294967295:
                    for fi32 := 0; fi32 < 14; fi32++ {
                        var v33 int32 = 2147483647
                        v3 -= (v3 / v3)
                        v33 += v33
                        _ = fi32
                        _ = v33
                    }
                }
                _ = v31
            } else {
                for fi34 := 0; fi34 < 13; fi34++ {
                    v35 := GenS2{F0: 0, F1: -32768, F2: (fi16 + fi34)}
                    v3 -= 1e10
                    _ = fi34
                    _ = v35
                }
                v3 -= v3
            }
        }
        v36 := GenS4{F0: v4, F1: -480, F2: (true && (v4 && (!true)))}
        v37 := GenS4{F0: (!false), F1: 878, F2: v4}
        _ = fi16
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        var v39 int32 = 42
        v40 := []int{}
        _ = fi38
        _ = v39
        _ = v40
    }
    switch 1 {
    case 433:
        v41 := GenS2{F0: 124, F1: -1, F2: -471}
        _ = v41
    case 125:
        delete(v1, 32767)
    default:
        fmt.Println(v2, v4, v4)
    }
    var v42 bool = (uint(v2) >= 554)
    var v43 uint16 = 65535
    v44 := GenS5{F0: -1, F1: 741}
    switch "" {
    case "":
        switch 2147483647 {
        case -1:
            for fi45 := 0; fi45 < 8; fi45++ {
                if ((fi45 << ((1) & 63)) != (fi45 ^ (fi45 % ((4294967295) | 1)))) {
                    for fi46 := 0; fi46 < 2; fi46++ {
                        if false {
                            v42 = true
                        } else {
                            v42 = ("" == "foo bar")
                        }
                        _ = fi46
                    }
                    v4 = (!(v42 && true))
                    v4 = v4
                } else {
                    v47 := []int{155, (-(fi45)), (fi45 | (fi45 + 32768)), int(v43)}
                    var v48 int8 = 22
                    _ = v47
                    _ = v48
                }
                _ = fi45
            }
        case 1:
            fmt.Println(v42, v43, v42)
        default:
            if (3.14159 > 1.5) {
                if (false && (1 != -2147483648)) {
                    var v49 uint8 = 2
                    _ = v49
                } else {
                    fmt.Println(v42)
                    if (!(!(v43 <= (v43 >> ((1) & 63))))) {
                        v43 -= v43
                    } else {
                        v43 = v43
                        if (-361 >= 42) {
                            v50 := GenS6{F0: 753, F1: 9223372036854775807}
                            if (0.0 >= 0.0) {
                                v51 := GenS6{F0: -1, F1: -1}
                                var v52 bool = (982 < 42)
                                v42 = v52
                                _ = v51
                                _ = v52
                            } else {
                                fmt.Println(v42, v43)
                                v53 := []int{}
                                _ = v53
                            }
                            _ = v50
                        } else {
                            v54 := []int{-2147483648, 22}
                            v42 = v42
                            for fi55 := 0; fi55 < 12; fi55++ {
                                if v4 {
                                    fmt.Println(v3)
                                    var v56 int16 = 1
                                    var v57 uint32 = 42
                                    _ = v56
                                    _ = v57
                                }
                                var v58 uint = 18446744073709551615
                                switch "foo bar" {
                                case "hello":
                                    var v59 int = (fi55 >> (((v58 >> (((^(v58))) & 63))) & 63))
                                    _ = v59
                                case "foo bar":
                                    if ("foo bar" <= "a") {
                                        v42 = false
                                        fmt.Println(v42, v4)
                                        for fi60 := 0; fi60 < 6; fi60++ {
                                            var v61 int64 = 9223372036854775807
                                            _ = fi60
                                            _ = v61
                                        }
                                    }
                                case "":
                                    if ((0.0 + -2.25) != 1.5) {
                                        fmt.Println(v3, v4, v58)
                                        var v62 int32 = 0
                                        _ = v62
                                    }
                                }
                                _ = fi55
                                _ = v58
                            }
                            _ = v54
                        }
                    }
                    fmt.Println(v43, v4, v43)
                }
                if v4 {
                    for fi63 := 0; fi63 < 20; fi63++ {
                        if false {
                            switch fi63 {
                            case 32768:
                                var v64 string = ("" + "foo bar")
                                _ = v64
                            case 2147483647:
                                v42 = v4
                            case 128:
                                for fi65 := 0; fi65 < 5; fi65++ {
                                    switch "" {
                                    case "a":
                                        if (!v4) {
                                            switch ((((("\n\t" + "a") + "hello") + "\n\t") + "hello") + "foo bar") {
                                            case "":
                                                v66 := GenS1{F0: 412, F1: (v2 | (v2 - 42))}
                                                _ = v66
                                            case "\n\t":
                                                var v67 uint64 = (^(v2))
                                                _ = v67
                                            case "foo bar":
                                                var v68 int16 = -1
                                                _ = v68
                                            default:
                                                v69 := []int{(fi65 ^ fi65), 1}
                                                _ = v69
                                            }
                                        } else {
                                            v70 := []int{fi63}
                                            _ = v70
                                        }
                                    case "\n\t":
                                        var v71 string = "foo bar"
                                        _ = v71
                                    case "hello":
                                        v4 = (4294967295 != 909)
                                    default:
                                        var v72 int16 = -32768
                                        _ = v72
                                    }
                                    fmt.Println(fi65, fi63, v2)
                                    _ = fi65
                                }
                            }
                            var v73 int64 = 41
                            v73 -= v73
                            _ = v73
                        }
                        var v74 uint8 = 2
                        _ = fi63
                        _ = v74
                    }
                    v75 := []int{-785, 358, -660, 42}
                    switch "hello" {
                    case "\n\t":
                        if (127 >= 128) {
                            var v76 float32 = ((float32(0.1) - -2.25) * 3.14159)
                            _ = v76
                        } else {
                            for fi77 := 0; fi77 < 18; fi77++ {
                                if true {
                                    v78 := GenS7{F0: fi77}
                                    fmt.Println(v2, v3, fi77)
                                    v79 := GenS8{F0: fi77, F1: ("foo bar" + (("foo bar" + "a") + ("" + (("\n\t" + "\n\t") + "")))), F2: (fi77 / ((9223372036854775807) | 1))}
                                    _ = v78
                                    _ = v79
                                }
                                var v80 int8 = 127
                                var v81 int = fi77
                                _ = fi77
                                _ = v80
                                _ = v81
                            }
                        }
                    default:
                        if ((32767 < 461) && true) {
                            var v82 int16 = -1
                            fmt.Println(v3, v43, v4)
                            switch -736 {
                            case -887:
                                v83 := v1[2147483647]
                                _ = v83
                            case 127:
                                v84 := GenS8{F0: -315, F1: "foo bar", F2: 1}
                                _ = v84
                            }
                            _ = v82
                        } else {
                            if false {
                                for fi85 := 0; fi85 < 10; fi85++ {
                                    fmt.Println(v43)
                                    fmt.Println(v3, v2)
                                    _ = fi85
                                }
                            }
                            v86 := GenS2{F0: 0, F1: 42, F2: 946}
                            v42 = (v2 <= (^(v2)))
                            _ = v86
                        }
                    }
                    _ = v75
                }
                v4 = (32767 == 191)
            } else {
                v42 = (!(65535 == 42))
                switch -320 {
                case 1:
                    v87 := GenS9{F0: -1, F1: 57, F2: 0}
                    _ = v87
                case 128:
                    v2 = 42
                default:
                    fmt.Println(v43, v3, v43)
                }
                v4 = (false || (v4 && v4))
            }
        }
    case "\n\t":
        v88 := []int{2147483647}
        _ = v88
    default:
        v42 = (("" + "foo bar") != "foo bar")
    }
    v89 := GenS10{F0: false}
    var v90 string = ("foo bar" + "foo bar")
    var v91 int64 = -922
    for fi92 := 0; fi92 < 17; fi92++ {
        v93 := []int{}
        v94 := GenS11{F0: -1}
        v42 = v42
        _ = fi92
        _ = v93
        _ = v94
    }
    if false {
        fmt.Println(v90, v90)
        for fi95 := 0; fi95 < 5; fi95++ {
            for fi96 := 0; fi96 < 16; fi96++ {
                var v97 bool = (true || v4)
                _ = fi96
                _ = v97
            }
            var v98 int32 = -710
            for fi99 := 0; fi99 < 15; fi99++ {
                v100 := []int{}
                _ = fi99
                _ = v100
            }
            _ = fi95
            _ = v98
        }
        v101 := GenS12{F0: v90, F1: -2147483648, F2: 1}
        _ = v101
    }
    var v102 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v89
    _ = v9
    _ = v90
    _ = v91
}
