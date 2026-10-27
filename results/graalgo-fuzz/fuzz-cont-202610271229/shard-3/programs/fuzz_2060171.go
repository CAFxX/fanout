package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 int32; F1 int64; F2 int32 }
func genFn0(p0 uint32) int64 {
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint32 = p0
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if (!(true || true)) {
        var v9 uint16 = 161
        switch "" {
        case "":
            v10 := []int{}
            _ = v10
        case "foo bar":
            switch -1 {
            case 556:
                var v11 float32 = ((1e10 * 1e10) + 3.14159)
                _ = v11
            }
        case "a":
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := GenS1{F0: 570}
                p0 = p0
                _ = fi12
                _ = v13
            }
        default:
            var v14 int8 = 116
            _ = v14
        }
        _ = v9
    } else {
        if (-415 == -879) {
            if (!(!false)) {
                fmt.Println(p0)
            }
            v15 := GenS2{F0: 1, F1: 206}
            _ = v15
        } else {
            var v16 uint64 = 42
            for fi17 := 0; fi17 < 15; fi17++ {
                fmt.Println(p0, p0, v16)
                var v18 float64 = (-2.25 / 0.1)
                switch (fi17 << ((1) & 63)) {
                case -9223372036854775808:
                    v19 := []int{32767, -2147483648, fi17}
                    _ = v19
                case -713:
                    var v20 int32 = 0
                    _ = v20
                case 2147483647:
                    fmt.Println(v18)
                }
                _ = fi17
                _ = v18
            }
            switch 32768 {
            case 32767:
                switch "foo bar" {
                case "":
                    v21 := []int{1, 128, -860}
                    _ = v21
                case "a":
                    for fi22 := 0; fi22 < 12; fi22++ {
                        p0 = p0
                        fmt.Println(v16, fi22, p0)
                        if true {
                            for fi23 := 0; fi23 < 1; fi23++ {
                                if (!(false && false)) {
                                    fmt.Println(p0)
                                    v16 = (v16 | v16)
                                }
                                v16 += (v16 & (v16 % ((v16) | 1)))
                                _ = fi23
                            }
                        }
                        _ = fi22
                    }
                }
            }
            _ = v16
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            switch 1 {
            case -1:
                var v25 int = 128
                _ = v25
            case -353:
                for fi26 := 0; fi26 < 20; fi26++ {
                    p0 = p0
                    _ = fi26
                }
            default:
                var v27 int64 = int64(p0)
                _ = v27
            }
            var v28 string = "foo bar"
            if (fi24 == (fi24 * fi24)) {
                switch (v28 + (("a" + v28) + "\n\t")) {
                case "a":
                    var v29 string = "a"
                    _ = v29
                }
                for fi30 := 0; fi30 < 1; fi30++ {
                    v31 := []int{(^(fi30)), fi30, (fi24 | fi30)}
                    if (true || true) {
                        v32 := GenS1{F0: 255}
                        _ = v32
                    } else {
                        if (3.14159 == 1.5) {
                            var v33 uint64 = 18446744073709551615
                            _ = v33
                        }
                    }
                    _ = fi30
                    _ = v31
                }
                fmt.Println(p0, v28, v28)
            }
            _ = fi24
            _ = v28
        }
    }
    _ = p0
    return 914
}

func main() {
    var v1 uint64 = 2
    var v2 int = 128
    v3 := make(map[string]bool)
    var v4 bool = true
    var v5 string = ("\n\t" + ((("\n\t" + "") + "") + "a"))
    v1 = v1
    if ((^(v1)) < 0) {
        for fi34 := 0; fi34 < 13; fi34++ {
            _ = len(v3)
            if (!(true || v4)) {
                v35 := []int{}
                for fi36 := 0; fi36 < 12; fi36++ {
                    var v37 int64 = 0
                    _ = fi36
                    _ = v37
                }
                _ = v35
            }
            v38 := []int{}
            _ = fi34
            _ = v38
        }
        fmt.Println(v1, v2, v5)
    } else {
        var v39 string = v5
        _ = v39
    }
    v5 = (v5 + "\n\t")
    if v4 {
        v40 := GenS3{F0: 29}
        _ = v40
    } else {
        v2 += 1
        if v4 {
            var v41 float32 = 3.14159
            v42 := GenS4{F0: "hello", F1: 1}
            v2 -= (v2 >> ((uint(v1)) & 63))
            _ = v41
            _ = v42
        } else {
            v5 = ((v5 + "") + ("foo bar" + ""))
            v2 = v2
            if (false && (0 <= uint8(v1))) {
                for fi43 := 0; fi43 < 19; fi43++ {
                    if (42 <= 417) {
                        _ = len(v3)
                        if v4 {
                            if v4 {
                                v44 := GenS5{F0: -56, F1: 1}
                                _ = v44
                            } else {
                                v4 = (!(((100 > 42) && v4) || (!v4)))
                                if (42 > 2147483648) {
                                    fmt.Println(v2, fi43)
                                    v1 = v1
                                }
                            }
                            var v45 uint = 255
                            _ = v45
                        } else {
                            if true {
                                if (1 >= 1) {
                                    v46 := []int{(fi43 / (((fi43 ^ v2)) | 1)), 42}
                                    switch v2 {
                                    case 32768:
                                        v5 = (((((("" + v5) + v5) + v5) + v5) + "foo bar") + (v5 + v5))
                                    case 1:
                                        if true {
                                            if (-2.25 <= ((((((-2.25 + (-2.25 + -2.25)) - 3.14159) + 1e10) / (1.5 - 3.14159)) + 3.14159) - (3.14159 + -2.25))) {
                                                v47 := GenS4{F0: v5, F1: -1}
                                                for fi48 := 0; fi48 < 4; fi48++ {
                                                    _ = len(v3)
                                                    switch -2147483648 {
                                                    case 100:
                                                        switch fi48 {
                                                        case -9223372036854775808:
                                                            v46 = append(v46, (v2 + (fi43 << ((618) & 63))))
                                                            var v49 uint32 = 1
                                                            _ = v49
                                                        case -997:
                                                            if v4 {
                                                                v50 := GenS4{F0: "a", F1: 1}
                                                                var v51 int16 = 1
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        default:
                                                            fmt.Println(v1)
                                                        }
                                                        v52 := GenS6{F0: 128}
                                                        var v53 int32 = 1
                                                        var v54 uint32 = 65535
                                                        _ = v52
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    _ = fi48
                                                }
                                                _ = v47
                                            }
                                            var v55 int8 = 42
                                            _ = v55
                                        } else {
                                            var v56 float32 = (1.5 / -2.25)
                                            v4 = (-65 != 42)
                                            _ = v56
                                        }
                                    default:
                                        var v57 int = 127
                                        _ = v57
                                    }
                                    _ = v46
                                } else {
                                    v3["hello"] = v4
                                    var v58 int32 = 0
                                    switch 127 {
                                    case 2147483647:
                                        for fi59 := 0; fi59 < 14; fi59++ {
                                            switch (v5 + (("\n\t" + "") + (((v5 + "foo bar") + ("foo bar" + (("foo bar" + v5) + v5))) + ((v5 + "") + "")))) {
                                            case "":
                                                if v4 {
                                                    v60 := GenS5{F0: 127, F1: 1}
                                                    _ = v60
                                                }
                                            }
                                            fmt.Println(v2, v58)
                                            v58 -= v58
                                            _ = fi59
                                        }
                                    case 127:
                                        if true {
                                            var v61 bool = v4
                                            _ = v61
                                        }
                                    default:
                                        var v62 uint32 = 4294967295
                                        _ = v62
                                    }
                                    _ = v58
                                }
                                v5 = v5
                            }
                            var v63 int8 = 127
                            _ = v63
                        }
                    }
                    var v64 bool = ((!v4) || true)
                    for fi65 := 0; fi65 < 9; fi65++ {
                        var v66 int8 = -1
                        var v67 bool = true
                        if (-9223372036854775808 > -9223372036854775808) {
                            v2 -= 32767
                            v5 = v5
                        }
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                    _ = fi43
                    _ = v64
                }
            }
        }
    }
    v2 += (-(v2))
    v68 := GenS1{F0: 375}
    var v69 float64 = (1e10 - 1.5)
    for fi70 := 0; fi70 < 3; fi70++ {
        var v71 uint64 = (v1 * 4294967295)
        v72 := GenS2{F0: 18446744073709551615, F1: 127}
        for fi73 := 0; fi73 < 14; fi73++ {
            var v74 bool = (v4 || v4)
            v71 += v71
            _ = fi73
            _ = v74
        }
        _ = fi70
        _ = v71
        _ = v72
    }
    if v4 {
        for fi75 := 0; fi75 < 20; fi75++ {
            var v76 float64 = v69
            _ = fi75
            _ = v76
        }
        var v77 int8 = 0
        v77 = (v77 >> ((395) & 63))
        _ = v77
    } else {
        v78 := GenS6{F0: 255}
        for fi79 := 0; fi79 < 1; fi79++ {
            for fi80 := 0; fi80 < 8; fi80++ {
                v4 = (float64(((3.14159 / (0.0 + float32((float64(1e10) / (v69 * ((v69 * (v69 + -2.25)) / v69)))))) - 3.14159)) < ((v69 * float64(0.0)) - (v69 * (0.1 + (v69 - (3.14159 / v69))))))
                v81 := GenS7{F0: v1, F1: 64}
                var v82 uint32 = 226
                _ = fi80
                _ = v81
                _ = v82
            }
            var v83 uint8 = 2
            v84 := GenS4{F0: "\n\t", F1: 1000}
            _ = fi79
            _ = v83
            _ = v84
        }
        var v85 uint32 = 1
        _ = v78
        _ = v85
    }
    for fi86 := 0; fi86 < 4; fi86++ {
        v87 := []int{(fi86 << ((595) & 63)), (v2 ^ (fi86 - fi86))}
        if false {
            var v88 uint16 = 1
            if v4 {
                _ = len(v87)
            } else {
                var v89 int32 = 0
                v1 = 4294967295
                _ = v89
            }
            _ = v88
        } else {
            v90 := v3[(v5 + ((("" + ((("\n\t" + "foo bar") + "\n\t") + (v5 + "\n\t"))) + "hello") + "a"))]
            for fi91 := 0; fi91 < 15; fi91++ {
                v1 -= (^(v1))
                v92 := GenS8{F0: 629, F1: 1, F2: -2147483648}
                if ((v4 && v90) && v4) {
                    var v93 uint64 = v1
                    var v94 string = v5
                    _ = v93
                    _ = v94
                }
                _ = fi91
                _ = v92
            }
            _ = v90
        }
        _ = fi86
        _ = v87
    }
    var v95 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v69:", v69)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
    _ = v95
}
