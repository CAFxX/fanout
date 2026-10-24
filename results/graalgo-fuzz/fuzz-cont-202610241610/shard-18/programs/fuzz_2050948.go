package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int64; F1 int64; F2 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 int32; F2 uint64 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 int32; F1 uint }
type GenS8 struct { F0 uint16; F1 uint; F2 int }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = float64(1e10)
    var v3 int64 = 0
    v4 := make(map[string]int)
    v5 := GenS1{F0: 4294967295, F1: 0, F2: 601}
    for fi6 := 0; fi6 < 15; fi6++ {
        if false {
            var v7 bool = false
            var v8 string = "hello"
            v9 := GenS2{F0: 25, F1: 2}
            _ = v7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    if (18446744073709551615 == 42) {
        var v10 int16 = -1
        v10 = -904
        v3 -= v3
        _ = v10
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        switch "hello" {
        case "\n\t":
            fmt.Println(v3, fi11)
        case "a":
            v3 += (v3 * v3)
        case "":
            for fi12 := 0; fi12 < 16; fi12++ {
                v3 += (^(v3))
                _ = fi12
            }
        default:
            fmt.Println(v3, v3)
        }
        var v13 int8 = 120
        _ = fi11
        _ = v13
    }
    if true {
        switch "hello" {
        case "":
            v14 := []int{-255}
            _ = v14
        case "hello":
            var v15 bool = (false && true)
            _ = v15
        }
        v2 = (3.14159 * (1.5 * (v2 - float64(0.1))))
    } else {
        fmt.Println(v3, v3)
        var v16 string = "foo bar"
        var v17 float32 = 0.1
        _ = v16
        _ = v17
    }
    fmt.Println(v3)
    for fi18 := 0; fi18 < 6; fi18++ {
        if false {
            for fi19 := 0; fi19 < 11; fi19++ {
                if (v2 == 1e10) {
                    var v20 int16 = -710
                    _ = v20
                } else {
                    if false {
                        v2 = 1e10
                        var v21 uint = 0
                        _ = v21
                    }
                }
                _ = fi19
            }
        } else {
            var v22 int16 = -213
            _ = v22
        }
        _ = fi18
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 float32 = float32((v2 - ((0.0 - v2) / (v2 / 1.0))))
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: 116, F1: 42, F2: 1}
    var v26 string = (("\n\t" + "hello") + ("" + "\n\t"))
    if ((1 != 42) && (9223372036854775807 <= -2147483648)) {
        if false {
            fmt.Println(v2, v3, v2)
        } else {
            var v27 uint32 = 340
            if true {
                v27 = (v27 & (v27 - v27))
            } else {
                fmt.Println(v2)
                if true {
                    if false {
                        switch v26 {
                        case "\n\t":
                            v2 = (float64((3.14159 - 3.14159)) / v2)
                        case "hello":
                            switch 127 {
                            case 4294967295:
                                var v28 uint32 = v27
                                _ = v28
                            case 42:
                                var v29 int64 = (v3 / (((v3 * v3)) | 1))
                                _ = v29
                            case 245:
                                v30 := v1[383]
                                _ = v30
                            default:
                                for fi31 := 0; fi31 < 5; fi31++ {
                                    if (!true) {
                                        v3 -= -9223372036854775808
                                        v27 = v27
                                    } else {
                                        for fi32 := 0; fi32 < 3; fi32++ {
                                            var v33 uint64 = 0
                                            fmt.Println(fi31, v33)
                                            if false {
                                                if true {
                                                    v27 -= (v27 - (v27 & 2147483648))
                                                    if (9223372036854775807 > 9223372036854775807) {
                                                        var v34 int16 = 699
                                                        v35 := GenS2{F0: 2, F1: uint8(v27)}
                                                        v36 := GenS3{F0: v3, F1: 926, F2: 100}
                                                        switch 2147483647 {
                                                        case 1:
                                                            switch fi32 {
                                                            case 42:
                                                                v27 += (v27 - v27)
                                                            case 32768:
                                                                if (!(!false)) {
                                                                    if ((false || false) && ((^(fi32)) > -2147483648)) {
                                                                        var v37 uint = 0
                                                                        var v38 uint64 = v33
                                                                        v33 += (^(v38))
                                                                        if true {
                                                                            fmt.Println(v38, v37, v3)
                                                                            var v39 int16 = (v34 ^ (v34 >> ((v37) & 63)))
                                                                            if (false || false) {
                                                                                if (true || (false || false)) {
                                                                                    v40 := GenS1{F0: v27, F1: 177, F2: 967}
                                                                                    var v41 uint32 = (v27 & (v27 & (v27 % ((v27) | 1))))
                                                                                    var v42 int16 = (^(v39))
                                                                                    _ = v40
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                }
                                                                                switch (fi32 % ((fi32) | 1)) {
                                                                                case -106:
                                                                                    fmt.Println(fi32, v37)
                                                                                case 2147483647:
                                                                                    v34 -= 42
                                                                                case 860:
                                                                                    v26 = ((("foo bar" + v26) + v26) + ("foo bar" + v26))
                                                                                }
                                                                                var v43 int8 = 30
                                                                                _ = v43
                                                                            } else {
                                                                                v39 = (v39 ^ (-(v39)))
                                                                            }
                                                                            _ = v39
                                                                        } else {
                                                                            v44 := GenS4{F0: "\n\t"}
                                                                            _ = v44
                                                                        }
                                                                        v3 = (v3 * 9223372036854775807)
                                                                        _ = v37
                                                                        _ = v38
                                                                    }
                                                                    v26 = (v26 + v26)
                                                                } else {
                                                                    v3 = (-(v3))
                                                                    fmt.Println(fi32)
                                                                    var v45 uint64 = 0
                                                                    _ = v45
                                                                }
                                                            default:
                                                                v26 = "\n\t"
                                                            }
                                                        case 9223372036854775807:
                                                            var v46 int16 = -32768
                                                            _ = v46
                                                        }
                                                        fmt.Println(v27)
                                                        _ = v34
                                                        _ = v35
                                                        _ = v36
                                                    } else {
                                                        v33 = (v33 - (^(v33)))
                                                        var v47 uint8 = 217
                                                        _ = v47
                                                    }
                                                    var v48 bool = false
                                                    _ = v48
                                                } else {
                                                    if (117 < 42) {
                                                        if (true && (!false)) {
                                                            switch fi32 {
                                                            case 0:
                                                                v26 = "foo bar"
                                                            case 124:
                                                                v33 -= v33
                                                            case 32768:
                                                                v49 := []int{fi31, (fi31 & (fi31 / ((fi31) | 1)))}
                                                                _ = v49
                                                            default:
                                                                var v50 int64 = v3
                                                                _ = v50
                                                            }
                                                            v27 -= 65535
                                                            v33 = v33
                                                            fmt.Println(v33, v26, v26)
                                                        }
                                                    } else {
                                                        var v51 string = v26
                                                        var v52 int64 = (-(v3))
                                                        v53 := GenS5{F0: fi31, F1: -2147483648, F2: 18446744073709551615}
                                                        _ = v51
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    if (138 < -924) {
                                                        fmt.Println(v2, fi32, v27)
                                                        if (!true) {
                                                            v26 = (v26 + v26)
                                                            v54 := GenS4{F0: ((v26 + "foo bar") + v26)}
                                                            _ = v54
                                                        } else {
                                                            v55 := GenS6{F0: 901, F1: -93}
                                                            v56 := GenS1{F0: (v27 - 205), F1: 2147483647, F2: -1}
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    }
                                                    v57 := GenS1{F0: v27, F1: 1, F2: 2147483647}
                                                    _ = v57
                                                }
                                            } else {
                                                v58 := []int{158, fi32, fi31}
                                                _ = v58
                                            }
                                            _ = fi32
                                            _ = v33
                                        }
                                        fmt.Println(v3, v26)
                                    }
                                    fmt.Println(v3, v26)
                                    _ = fi31
                                }
                            }
                        }
                        v59 := GenS2{F0: 1, F1: 1}
                        _ = v59
                    }
                }
            }
            var v60 int = int(v3)
            _ = v27
            _ = v60
        }
        fmt.Println(v2, v26)
    }
    if false {
        v26 = ((v26 + ("\n\t" + ("" + ""))) + v26)
    }
    if (true || (382 == 0)) {
        if false {
            if false {
                switch ("foo bar" + v26) {
                case "foo bar":
                    if ((255 <= 65535) && false) {
                        v61 := GenS7{F0: 42, F1: 528}
                        _ = v61
                    } else {
                        v62 := v4["foo bar"]
                        var v63 string = "foo bar"
                        var v64 int = -2147483648
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                case "hello":
                    v2 = v2
                default:
                    fmt.Println(v26, v3, v3)
                }
            } else {
                delete(v4, "")
                v1[793] = ""
                v65 := GenS8{F0: 353, F1: 272, F2: -472}
                _ = v65
            }
            for fi66 := 0; fi66 < 1; fi66++ {
                for fi67 := 0; fi67 < 7; fi67++ {
                    var v68 int16 = 911
                    delete(v4, "a")
                    _ = fi67
                    _ = v68
                }
                fmt.Println(v3)
                v69 := GenS9{F0: 99}
                _ = fi66
                _ = v69
            }
        }
    }
    v70 := GenS10{F0: 300, F1: 1}
    var v71 int32 = 42
    v72 := []int{-9223372036854775808}
    fmt.Println(v2, v2)
    v73 := GenS1{F0: 785, F1: -260, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v71:", v71)
    fmt.Println("slicelen v72:", len(v72), v72)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v71
    _ = v72
    _ = v73
}
