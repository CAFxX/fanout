package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint; F2 int32 }
type GenS3 struct { F0 int8; F1 uint8; F2 string }
type GenS4 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 string; F2 uint }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 bool; F1 string }
type GenS10 struct { F0 uint16; F1 uint32 }
type GenS11 struct { F0 uint; F1 int }
type GenS12 struct { F0 uint32; F1 int16 }

func main() {
    var v1 float64 = (3.14159 / 3.14159)
    var v2 int64 = 42
    v3 := make(map[string]bool)
    var v4 float32 = (1e10 - -2.25)
    var v5 float32 = v4
    var v6 float32 = float32((v1 * float64(v5)))
    fmt.Println(v5, v2, v2)
    var v7 bool = false
    switch "\n\t" {
    case "a":
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := []int{(fi8 | (fi8 / ((fi8) | 1)))}
            _ = fi8
            _ = v9
        }
    case "foo bar":
        switch 2147483647 {
        case 4294967295:
            var v10 int32 = 42
            _ = v10
        case 42:
            fmt.Println(v2)
        }
    }
    var v11 bool = v7
    if (-344 == v2) {
        var v12 int64 = (v2 + v2)
        _ = v12
    } else {
        fmt.Println(v6, v11, v1)
        v4 = 1.5
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 float64 = 1.5
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint64 = 0
        v2 -= v2
        if (false && ((127 <= 2) && (float32(1.5) >= v5))) {
            for fi17 := 0; fi17 < 10; fi17++ {
                if v11 {
                    fmt.Println(v7, fi15, v2)
                    delete(v3, "hello")
                    switch fi15 {
                    case -9223372036854775808:
                        var v18 uint8 = 0
                        _ = v18
                    }
                } else {
                    var v19 float64 = v1
                    v4 = 0.1
                    if v11 {
                        v5 = (v6 * v4)
                    } else {
                        fmt.Println(v5, v6, fi17)
                        var v20 float32 = v4
                        if (-1 >= 0) {
                            v21 := []int{(fi15 ^ fi17), (fi17 + 32767), 4294967295, (fi17 % ((600) | 1))}
                            v22 := GenS1{F0: 888}
                            _ = v21
                            _ = v22
                        } else {
                            switch "foo bar" {
                            case "hello":
                                if (v16 <= 18446744073709551615) {
                                    fmt.Println(fi17, v6, v1)
                                }
                            case "\n\t":
                                switch (-(fi17)) {
                                case 435:
                                    v23 := GenS2{F0: 44, F1: 4294967295, F2: -1}
                                    _ = v23
                                case 9223372036854775807:
                                    v24 := GenS3{F0: 42, F1: 127, F2: (((("foo bar" + "hello") + "foo bar") + "\n\t") + "\n\t")}
                                    _ = v24
                                case 742:
                                    switch ("" + ("hello" + "")) {
                                    case "a":
                                        var v25 uint32 = 65535
                                        _ = v25
                                    case "\n\t":
                                        v26 := []int{-1, -9223372036854775808}
                                        _ = v26
                                    case "":
                                        if (128 == 255) {
                                            v2 = v2
                                        } else {
                                            v27 := GenS4{F0: 1, F1: v16, F2: 127}
                                            _ = v27
                                        }
                                    default:
                                        v3[""] = (685 > 2147483647)
                                    }
                                }
                            case "foo bar":
                                switch "foo bar" {
                                case "\n\t":
                                    if ((^(v16)) < v16) {
                                        v5 += (float32(float64(1.5)) * v6)
                                        v6 = v20
                                        if (1 == 430) {
                                            if v7 {
                                                v19 -= float64((v20 - (v20 / v6)))
                                            } else {
                                                switch fi15 {
                                                case -9223372036854775808:
                                                    v28 := GenS4{F0: -378, F1: 18446744073709551615, F2: 0}
                                                    _ = v28
                                                case -716:
                                                    var v29 uint8 = 128
                                                    _ = v29
                                                default:
                                                    var v30 uint16 = 1
                                                    _ = v30
                                                }
                                                var v31 string = ("" + "")
                                                _ = v31
                                            }
                                            v32 := []int{fi15, fi17}
                                            v6 = -2.25
                                            v7 = (!(v7 || false))
                                            _ = v32
                                        } else {
                                            v33 := GenS5{F0: v2}
                                            _ = v33
                                        }
                                    }
                                case "foo bar":
                                    v34 := GenS5{F0: v2}
                                    _ = v34
                                case "hello":
                                    if false {
                                        switch int(v16) {
                                        case 0:
                                            var v35 float64 = (float64(v20) - (v19 / v1))
                                            _ = v35
                                        default:
                                            var v36 uint64 = (v16 ^ 308)
                                            _ = v36
                                        }
                                        if v7 {
                                            var v37 string = "a"
                                            switch -108 {
                                            case 32767:
                                                if ((0 > 305) && v7) {
                                                    switch "hello" {
                                                    case "hello":
                                                        v38 := GenS2{F0: 0, F1: 18446744073709551615, F2: 42}
                                                        _ = v38
                                                    case "\n\t":
                                                        v2 = (v2 ^ v2)
                                                        if v7 {
                                                            if false {
                                                                switch (fi15 + fi17) {
                                                                case 32767:
                                                                    switch v37 {
                                                                    case "a":
                                                                        v39 := GenS6{F0: 42, F1: "a", F2: 875}
                                                                        _ = v39
                                                                    }
                                                                    fmt.Println(v6, fi15, v4)
                                                                case -958:
                                                                    fmt.Println(v11)
                                                                case 9223372036854775807:
                                                                    switch v37 {
                                                                    case "\n\t":
                                                                        if false {
                                                                            var v40 bool = (true || v7)
                                                                            switch fi15 {
                                                                            case 971:
                                                                                v41 := GenS7{F0: fi15}
                                                                                _ = v41
                                                                            default:
                                                                                var v42 uint16 = 42
                                                                                _ = v42
                                                                            }
                                                                            fmt.Println(v5, v20)
                                                                            var v43 uint = 2
                                                                            _ = v40
                                                                            _ = v43
                                                                        } else {
                                                                            v44 := v3[("foo bar" + v37)]
                                                                            var v45 int32 = -931
                                                                            _ = len(v3)
                                                                            _ = v44
                                                                            _ = v45
                                                                        }
                                                                    case "a":
                                                                        var v46 int64 = v2
                                                                        _ = v46
                                                                    default:
                                                                        v47 := GenS8{F0: (-(fi15)), F1: "\n\t"}
                                                                        _ = v47
                                                                    }
                                                                default:
                                                                    fmt.Println(fi15, v37, fi15)
                                                                }
                                                            }
                                                        }
                                                        v1 = ((v19 / v19) + ((v19 / v19) / v1))
                                                    case "a":
                                                        v19 = v1
                                                        v1 += v1
                                                        var v48 string = "\n\t"
                                                        v6 = (-2.25 - -2.25)
                                                        _ = v48
                                                    default:
                                                        fmt.Println(v4)
                                                    }
                                                    v7 = (2 < 189)
                                                }
                                            case 128:
                                                if (303 > 2) {
                                                    fmt.Println(v6)
                                                    v7 = false
                                                    v49 := GenS4{F0: 1, F1: v16, F2: 128}
                                                    var v50 float32 = (((v20 * (v5 - float32(v1))) + v4) / 3.14159)
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    var v51 bool = v7
                                                    var v52 uint16 = 236
                                                    delete(v3, v37)
                                                    _ = v51
                                                    _ = v52
                                                }
                                            case 4294967295:
                                                v53 := []int{1, (fi17 / ((969) | 1)), 4294967295, 956}
                                                _ = v53
                                            default:
                                                fmt.Println(v6)
                                            }
                                            _ = v37
                                        }
                                        if v11 {
                                            if (9223372036854775807 == (v2 & (-(v2)))) {
                                                var v54 int16 = -524
                                                switch "foo bar" {
                                                case "a":
                                                    v55 := GenS9{F0: (v7 || false), F1: ""}
                                                    _ = v55
                                                case "":
                                                    v56 := GenS3{F0: 32, F1: 133, F2: "hello"}
                                                    _ = v56
                                                case "foo bar":
                                                    delete(v3, ("" + ("foo bar" + ("hello" + "foo bar"))))
                                                }
                                                var v57 int64 = 559
                                                _ = v54
                                                _ = v57
                                            } else {
                                                v58 := []int{(fi17 / (((fi17 << ((uint(v16)) & 63))) | 1)), fi17, fi15, (^(fi17))}
                                                var v59 float64 = v1
                                                v2 -= -9223372036854775808
                                                _ = v58
                                                _ = v59
                                            }
                                            fmt.Println(v5, v6, fi15)
                                            if (36 >= 0) {
                                                v4 = (v4 + (v4 - v5))
                                                var v60 int8 = 0
                                                var v61 int64 = v2
                                                _ = v60
                                                _ = v61
                                            }
                                        }
                                    } else {
                                        fmt.Println(v2, v20, v11)
                                        switch fi17 {
                                        case -2147483648:
                                            var v62 float32 = (-2.25 * v20)
                                            _ = v62
                                        case 636:
                                            v4 += ((v6 - 0.0) - ((v4 * float32((1e10 / (3.14159 - 0.0)))) * v6))
                                        default:
                                            if true {
                                                var v63 int16 = -1
                                                _ = v63
                                            } else {
                                                v11 = (!false)
                                                var v64 uint32 = 1
                                                var v65 uint = 18446744073709551615
                                                v3["foo bar"] = v11
                                                _ = v64
                                                _ = v65
                                            }
                                            var v66 string = "foo bar"
                                            if false {
                                                if true {
                                                    v1 = (v19 / (((v1 * (v1 * v1)) + (v19 - v19)) / (v19 * v1)))
                                                }
                                                if (1000 == -32768) {
                                                    if v11 {
                                                        v67 := GenS10{F0: 2, F1: 974}
                                                        _ = v67
                                                    } else {
                                                        if v11 {
                                                            var v68 float64 = (0.1 / v1)
                                                            _ = v68
                                                        }
                                                    }
                                                    var v69 uint32 = 1
                                                    if ((true || ((v66 != v66) || v7)) || (!(v19 >= -2.25))) {
                                                        switch 1 {
                                                        case 4294967295:
                                                            var v70 string = "foo bar"
                                                            _ = v70
                                                        }
                                                        v19 -= v1
                                                        v7 = v11
                                                    } else {
                                                        if v11 {
                                                            var v71 int = fi17
                                                            _ = v71
                                                        }
                                                    }
                                                    _ = v69
                                                } else {
                                                    v72 := GenS11{F0: 504, F1: (fi15 * (fi15 - -1))}
                                                    v3[((v66 + "\n\t") + "")] = false
                                                    _ = v72
                                                }
                                            } else {
                                                var v73 int = (fi17 + fi17)
                                                fmt.Println(fi15)
                                                v1 = 1e10
                                                v74 := []int{v73}
                                                _ = v73
                                                _ = v74
                                            }
                                            if (218 != 150) {
                                                fmt.Println(v11)
                                                var v75 float32 = v4
                                                v2 = 682
                                                _ = v75
                                            } else {
                                                fmt.Println(v19)
                                                v11 = ((v66 >= "\n\t") && true)
                                            }
                                            _ = v66
                                        }
                                    }
                                default:
                                    var v76 int16 = 0
                                    _ = v76
                                }
                            default:
                                v77 := []int{(fi15 ^ (fi15 / ((127) | 1))), (fi15 * (fi15 + (fi15 - -1)))}
                                _ = v77
                            }
                        }
                        v3["hello"] = ((42 <= 0) && true)
                        _ = v20
                    }
                    _ = v19
                }
                _ = fi17
            }
            var v78 float32 = v6
            _ = v78
        }
        _ = fi15
        _ = v16
    }
    v79 := GenS12{F0: 4294967295, F1: 138}
    var v80 uint64 = 163
    var v81 float64 = v1
    var v82 float64 = (3.14159 * ((1e10 * (-2.25 / ((v1 * v81) * v81))) / ((v1 - (0.0 * v81)) + v1)))
    v80 -= v80
    var v83 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v79
    _ = v80
    _ = v81
    _ = v82
    _ = v83
}
