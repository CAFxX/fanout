package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 bool; F1 uint8; F2 bool }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 string; F1 int16; F2 int8 }
type GenS10 struct { F0 string; F1 string }
type GenS11 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS12 struct { F0 int8; F1 string }
type GenS13 struct { F0 int16; F1 int }
type GenS14 struct { F0 uint64; F1 int; F2 uint }
func genFn0(p0 uint32, p1 int32) int32 {
    var v6 string = "\n\t"
    p0 = (p0 ^ (p0 - 1))
    switch 128 {
    case -1:
        v7 := []int{42, 349}
        _ = v7
    case 4294967295:
        var v8 uint64 = 2
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return p1
}
func genFn1(p0 int16) int64 {
    p0 = -32768
    v9 := []int{0, 1, -2147483648, 42}
    var v10 int = -811
    v10 = (-(v10))
    _ = p0
    _ = v10
    _ = v9
    return -516
}

func main() {
    var v1 uint8 = 255
    var v2 bool = (v1 == v1)
    var v3 int8 = 1
    v4 := make(map[string]bool)
    v5 := make(map[int]int)
    if v2 {
        v11 := GenS1{F0: (^(v3))}
        _ = v11
    } else {
        v12 := GenS2{F0: -128, F1: "foo bar"}
        v1 -= 42
        v2 = false
        _ = v12
    }
    delete(v4, (("hello" + "") + ""))
    v13 := GenS3{F0: 1}
    var v14 int64 = 9223372036854775807
    if v2 {
        switch "foo bar" {
        case "foo bar":
            switch ("a" + ("\n\t" + (("a" + "foo bar") + "\n\t"))) {
            case "":
                v1 -= v1
            default:
                fmt.Println(v2, v3, v2)
            }
        case "hello":
            if (float32(0.1) <= 1.5) {
                for fi15 := 0; fi15 < 15; fi15++ {
                    v16 := GenS3{F0: 2147483647}
                    switch "foo bar" {
                    case "foo bar":
                        v2 = v2
                    case "\n\t":
                        fmt.Println(fi15, v1, v2)
                    case "hello":
                        v1 = v1
                    default:
                        var v17 int16 = 765
                        _ = v17
                    }
                    _ = fi15
                    _ = v16
                }
                fmt.Println(v2, v1, v2)
                v3 -= v3
            } else {
                fmt.Println(v1, v1)
                var v18 int = 859
                if v2 {
                    var v19 int16 = -1
                    for fi20 := 0; fi20 < 10; fi20++ {
                        v14 = v14
                        v21 := []int{(fi20 ^ (-(v18))), 515, int(v14), fi20}
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                } else {
                    var v22 int64 = (v14 << ((1) & 63))
                    var v23 float32 = (float32((3.14159 / ((-2.25 - 0.0) * float64(float32((((0.0 - 0.1) + 0.1) + 0.1)))))) / -2.25)
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v1 = (^(v1))
                        _ = fi24
                    }
                    _ = v22
                    _ = v23
                }
                _ = v18
            }
        case "":
            v1 -= v1
        }
        var v25 int64 = v14
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v14)
        case "":
            v25 = v14
        }
        _ = v25
    } else {
        var v26 int32 = 42
        _ = v26
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v14)
    case "\n\t":
        v2 = true
    case "hello":
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 float32 = float32(0.1)
            v1 = v1
            v29 := GenS1{F0: 1}
            _ = fi27
            _ = v28
            _ = v29
        }
    default:
        if v2 {
            fmt.Println(v2, v3, v1)
        } else {
            for fi30 := 0; fi30 < 16; fi30++ {
                if v2 {
                    for fi31 := 0; fi31 < 18; fi31++ {
                        var v32 float64 = (-2.25 / (0.1 + 0.0))
                        switch ("foo bar" + "hello") {
                        case "a":
                            v33 := GenS1{F0: v3}
                            _ = v33
                        case "\n\t":
                            var v34 int16 = -266
                            _ = v34
                        case "":
                            v35 := GenS4{F0: 665, F1: 18446744073709551615}
                            _ = v35
                        default:
                            v14 -= -9223372036854775808
                        }
                        _ = fi31
                        _ = v32
                    }
                    v36 := GenS5{F0: 42, F1: 940}
                    _ = v36
                } else {
                    if ((("\n\t" + ("" + ("foo bar" + "a"))) <= "foo bar") && true) {
                        switch (("" + ("a" + "hello")) + (("" + "hello") + "")) {
                        case "hello":
                            v37 := GenS5{F0: 32767, F1: 0}
                            _ = v37
                        case "a":
                            var v38 uint64 = 1
                            _ = v38
                        case "foo bar":
                            if (42 == 65535) {
                                _ = len(v5)
                            } else {
                                v39 := GenS5{F0: -739, F1: 18446744073709551615}
                                v3 += (v3 >> ((42) & 63))
                                _ = v39
                            }
                        default:
                            var v40 float64 = 0.0
                            _ = v40
                        }
                        var v41 int8 = v3
                        _ = v41
                    } else {
                        switch 128 {
                        case 9223372036854775807:
                            v42 := GenS3{F0: genFn0(59, -791)}
                            _ = v42
                        case 158:
                            if v2 {
                                var v43 string = "foo bar"
                                var v44 bool = (v14 > (^(v14)))
                                switch "foo bar" {
                                case "\n\t":
                                    v3 = v3
                                case "hello":
                                    if v2 {
                                        if (2147483648 <= 2147483648) {
                                            v45 := GenS4{F0: 47, F1: 0}
                                            _ = v45
                                        } else {
                                            for fi46 := 0; fi46 < 11; fi46++ {
                                                var v47 int16 = 0
                                                _ = fi46
                                                _ = v47
                                            }
                                        }
                                        var v48 uint32 = 2
                                        v2 = ((fi30 | (fi30 | 32767)) == (fi30 ^ (fi30 << ((0) & 63))))
                                        _ = v48
                                    } else {
                                        var v49 uint16 = 65535
                                        _ = v49
                                    }
                                default:
                                    fmt.Println(v14, v14)
                                }
                                _ = v43
                                _ = v44
                            }
                        case 800:
                            v50 := []int{(fi30 >> ((995) & 63)), (-(fi30))}
                            _ = v50
                        }
                    }
                    var v51 int16 = 1000
                    _ = v51
                }
                switch (fi30 >> ((0) & 63)) {
                case 32767:
                    if v2 {
                        v52 := v5[2147483647]
                        v1 = 255
                        _ = v52
                    } else {
                        _ = len(v4)
                        fmt.Println(v3, v2)
                    }
                case 127:
                    if v2 {
                        v53 := []int{}
                        for fi54 := 0; fi54 < 13; fi54++ {
                            v55 := GenS2{F0: v3, F1: ("hello" + ("foo bar" + ""))}
                            var v56 int32 = 0
                            var v57 bool = (!(!v2))
                            _ = fi54
                            _ = v55
                            _ = v56
                            _ = v57
                        }
                        _ = v53
                    } else {
                        if false {
                            fmt.Println(v1, v2)
                        }
                    }
                }
                _ = fi30
            }
        }
    }
    v3 = v3
    if (!v2) {
        var v58 bool = v2
        v59 := GenS4{F0: 65535, F1: 266}
        v60 := GenS6{F0: 764, F1: 4294967295}
        _ = v58
        _ = v59
        _ = v60
    } else {
        if v2 {
            fmt.Println(v3, v3, v3)
            for fi61 := 0; fi61 < 14; fi61++ {
                if v2 {
                    v3 = (v3 >> ((712) & 63))
                    var v62 int = 32767
                    _ = v62
                } else {
                    v14 = genFn1(0)
                }
                var v63 string = ""
                v64 := GenS4{F0: 0, F1: 255}
                _ = fi61
                _ = v63
                _ = v64
            }
            var v65 string = "hello"
            _ = v65
        } else {
            if true {
                if (v2 || true) {
                    switch 559 {
                    case 32768:
                        v66 := GenS7{F0: (65535 <= 2), F1: (v1 >> ((2) & 63)), F2: v2}
                        _ = v66
                    case 4294967295:
                        for fi67 := 0; fi67 < 19; fi67++ {
                            for fi68 := 0; fi68 < 12; fi68++ {
                                v69 := GenS1{F0: v3}
                                v70 := GenS8{F0: 32767, F1: (fi68 << ((255) & 63))}
                                var v71 uint = 342
                                _ = fi68
                                _ = v69
                                _ = v70
                                _ = v71
                            }
                            v3 = (v3 << ((18446744073709551615) & 63))
                            for fi72 := 0; fi72 < 6; fi72++ {
                                var v73 int8 = -1
                                v74 := GenS9{F0: "foo bar", F1: -1, F2: 42}
                                v3 = v3
                                _ = fi72
                                _ = v73
                                _ = v74
                            }
                            _ = fi67
                        }
                    }
                    fmt.Println(v1, v3, v14)
                    if v2 {
                        v1 += 42
                        if (2 == 2) {
                            fmt.Println(v14, v3)
                        } else {
                            var v75 string = ((("foo bar" + "foo bar") + "a") + "foo bar")
                            v2 = v2
                            _ = v75
                        }
                    } else {
                        switch -939 {
                        case 1:
                            fmt.Println(v3, v1, v3)
                        default:
                            v1 = v1
                        }
                    }
                }
            } else {
                v76 := GenS10{F0: ("hello" + (("\n\t" + "\n\t") + "")), F1: "foo bar"}
                for fi77 := 0; fi77 < 13; fi77++ {
                    if ((!v2) || false) {
                        if v2 {
                            var v78 int16 = 42
                            var v79 float64 = 0.0
                            _ = v78
                            _ = v79
                        }
                    } else {
                        v80 := []int{fi77}
                        if false {
                            if v2 {
                                v81 := GenS4{F0: 1, F1: 1}
                                v82 := GenS6{F0: 335, F1: 42}
                                delete(v4, "")
                                _ = v81
                                _ = v82
                            }
                        } else {
                            switch fi77 {
                            case 4294967295:
                                var v83 int64 = 469
                                _ = v83
                            case -9223372036854775808:
                                switch 42 {
                                case 943:
                                    for fi84 := 0; fi84 < 5; fi84++ {
                                        v85 := GenS7{F0: v2, F1: 42, F2: (30 >= 2)}
                                        v86 := GenS11{F0: 0, F1: 60, F2: (v1 + 128)}
                                        v87 := GenS12{F0: v3, F1: ""}
                                        _ = fi84
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                    }
                                case 2147483647:
                                    v88 := GenS13{F0: 42, F1: (fi77 + (^(fi77)))}
                                    _ = v88
                                default:
                                    var v89 uint16 = 0
                                    _ = v89
                                }
                            }
                            v3 = v3
                            fmt.Println(v1, v1)
                        }
                        _ = v80
                    }
                    delete(v4, "hello")
                    var v90 float32 = 0.1
                    _ = fi77
                    _ = v90
                }
                _ = v76
            }
            switch "foo bar" {
            case "":
                if v2 {
                    v4[""] = (true && (!(716 <= 0)))
                    var v91 bool = false
                    var v92 float64 = (0.1 * 0.1)
                    _ = v91
                    _ = v92
                } else {
                    var v93 uint8 = (^(v1))
                    _ = v93
                }
            default:
                v1 -= 1
            }
        }
        var v94 int8 = (v3 & 0)
        v3 += 100
        _ = v94
    }
    v95 := GenS14{F0: 0, F1: 9223372036854775807, F2: 555}
    v2 = true
    var v96 float64 = 0.1
    v97 := []int{}
    var v98 float64 = (v96 * v96)
    fmt.Println(v96)
    v1 += (v1 & v1)
    switch "foo bar" {
    case "":
        v2 = v2
    default:
        if (0 <= 2) {
            if ((3.14159 * (3.14159 + 1e10)) >= 1.5) {
                var v99 int8 = (v3 << ((540) & 63))
                var v100 bool = (("" + (("hello" + "\n\t") + "a")) > "foo bar")
                _ = v100
                _ = v99
            }
            v96 += float64((float32(v96) + float32((float64(1.5) + v98))))
            if v2 {
                var v101 bool = v2
                _ = v101
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v96:", v96)
    fmt.Println("slicelen v97:", len(v97), v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v95
    _ = v96
    _ = v97
    _ = v98
}
