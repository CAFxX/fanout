package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint32; F1 int16; F2 uint }
func genFn0() uint {
    var v4 uint8 = 59
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: int8(fi5), F1: 42, F2: 911}
        v7 := GenS1{F0: 0, F1: 2147483647, F2: -9223372036854775808}
        var v8 uint64 = 1
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 0}
    _ = v4
    _ = v9
    return 42
}
func genFn1() int {
    var v10 uint32 = 0
    switch -755 {
    case -2147483648:
        var v11 int16 = 32767
        _ = v11
    case 42:
        v10 = 2147483648
    case 128:
        var v12 int16 = 371
        _ = v12
    default:
        v13 := GenS3{F0: 1}
        _ = v13
    }
    v14 := []int{-2147483648}
    if (-2.25 == 0.1) {
        v15 := GenS4{F0: 42, F1: 190, F2: -32768}
        _ = v15
    }
    _ = v10
    _ = v14
    return 47
}

func main() {
    var v1 string = "\n\t"
    var v2 float64 = 1e10
    var v3 uint32 = 391
    if (32768 != 1) {
        v2 -= v2
    } else {
        if true {
            v16 := []int{1, genFn1(), 9223372036854775807}
            v16 = append(v16, genFn1())
            _ = v16
        } else {
            if true {
                switch v1 {
                case "a":
                    var v17 float64 = 3.14159
                    _ = v17
                case "":
                    if ((9223372036854775807 != 9223372036854775807) || (!(127 <= 1))) {
                        v1 = v1
                        v18 := GenS1{F0: 127, F1: (v3 / ((v3) | 1)), F2: 9223372036854775807}
                        fmt.Println(v1, v1, v1)
                        _ = v18
                    } else {
                        v19 := []int{853, -9223372036854775808, int(v3)}
                        for fi20 := 0; fi20 < 13; fi20++ {
                            if (v1 <= "foo bar") {
                                _ = len(v19)
                                if (!(!false)) {
                                    var v21 int8 = 42
                                    _ = v21
                                }
                            } else {
                                fmt.Println(v2, v2)
                                var v22 int16 = 1
                                if true {
                                    if false {
                                        fmt.Println(v22, v22, v3)
                                        v23 := GenS4{F0: 74, F1: 2, F2: v22}
                                        _ = v23
                                    } else {
                                        v3 += 381
                                        for fi24 := 0; fi24 < 13; fi24++ {
                                            v2 = v2
                                            var v25 uint64 = 2
                                            v2 += (float64(1e10) + (0.1 / 3.14159))
                                            v25 -= v25
                                            var v26 uint = 2
                                            v27 := GenS1{F0: 100, F1: v3, F2: -1}
                                            v3 -= 7
                                            var v28 float32 = (1.5 + -2.25)
                                            _ = len(v19)
                                            _ = fi24
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                    }
                                    var v29 string = (("a" + "") + "\n\t")
                                    switch fi20 {
                                    case 0:
                                        var v30 int = fi20
                                        _ = v30
                                    case -469:
                                        var v31 int64 = 9223372036854775807
                                        _ = v31
                                    case -2147483648:
                                        for fi32 := 0; fi32 < 11; fi32++ {
                                            var v33 float32 = ((1e10 + -2.25) + (float32(3.14159) - 1e10))
                                            v29 = "\n\t"
                                            var v34 uint = 65535
                                            _ = fi32
                                            _ = v33
                                            _ = v34
                                        }
                                    default:
                                        v35 := GenS3{F0: 1}
                                        _ = v35
                                    }
                                    _ = v29
                                } else {
                                    switch (fi20 / ((fi20) | 1)) {
                                    case 354:
                                        var v36 uint = 878
                                        _ = v36
                                    case 128:
                                        var v37 string = v1
                                        _ = v37
                                    case -494:
                                        v1 = v1
                                    }
                                    if (0 != -1) {
                                        if false {
                                            if (!(genFn1() >= (fi20 / (((fi20 | (fi20 + fi20))) | 1)))) {
                                                var v38 bool = (!false)
                                                for fi39 := 0; fi39 < 10; fi39++ {
                                                    if (!(!v38)) {
                                                        if (1.5 >= 1.5) {
                                                            v19 = append(v19, fi39)
                                                            switch 128 {
                                                            case -687:
                                                                v19 = append(v19, (^(fi39)))
                                                            case 594:
                                                                v3 -= (v3 | v3)
                                                                var v40 bool = v38
                                                                _ = v40
                                                            }
                                                        }
                                                        var v41 string = v1
                                                        var v42 int16 = v22
                                                        v22 -= 42
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    _ = fi39
                                                }
                                                for fi43 := 0; fi43 < 10; fi43++ {
                                                    v22 = (v22 << ((509) & 63))
                                                    if (!(1 >= 44)) {
                                                        var v44 int64 = 1
                                                        fmt.Println(v3, fi43)
                                                        if (42 > 0) {
                                                            v45 := GenS4{F0: 244, F1: 255, F2: v22}
                                                            var v46 int = (fi43 >> ((513) & 63))
                                                            _ = len(v19)
                                                            _ = v45
                                                            _ = v46
                                                        } else {
                                                            v38 = (!v38)
                                                            v3 = (v3 - 2147483647)
                                                        }
                                                        _ = v44
                                                    } else {
                                                        fmt.Println(v1, v22, v38)
                                                        v47 := GenS3{F0: 128}
                                                        _ = v47
                                                    }
                                                    v1 = "hello"
                                                    _ = fi43
                                                }
                                                _ = v38
                                            } else {
                                                var v48 uint = 682
                                                v48 -= v48
                                                for fi49 := 0; fi49 < 13; fi49++ {
                                                    v50 := GenS2{F0: -1}
                                                    var v51 int8 = 42
                                                    if (v51 < 100) {
                                                        v22 = (v22 << (((v48 ^ v48)) & 63))
                                                    }
                                                    _ = fi49
                                                    _ = v50
                                                    _ = v51
                                                }
                                                _ = v48
                                            }
                                            var v52 uint16 = 204
                                            _ = v52
                                        } else {
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                            if true {
                                                v22 += v22
                                                if (401 <= 2) {
                                                    _ = len(v19)
                                                } else {
                                                    fmt.Println(v2, v1, v22)
                                                    switch (fi20 | 32767) {
                                                    case 2147483647:
                                                        for fi53 := 0; fi53 < 15; fi53++ {
                                                            if true {
                                                                v22 = v22
                                                            }
                                                            var v54 uint8 = 128
                                                            _ = fi53
                                                            _ = v54
                                                        }
                                                    case -1:
                                                        fmt.Println(fi20, v22)
                                                    }
                                                }
                                            }
                                        }
                                        if true {
                                            fmt.Println(v1, v1, v3)
                                        }
                                        v55 := GenS2{F0: -128}
                                        _ = v55
                                    } else {
                                        v2 += v2
                                    }
                                }
                                _ = v22
                            }
                            _ = fi20
                        }
                        _ = v19
                    }
                default:
                    v56 := []int{816, -9223372036854775808, 550}
                    _ = v56
                }
                if (((-32768 > 1000) || (v3 < v3)) && false) {
                    v57 := []int{}
                    if false {
                        v57 = append(v57, 128)
                        if true {
                            if true {
                                var v58 int = genFn1()
                                _ = v58
                            }
                        } else {
                            var v59 uint32 = (v3 - (v3 ^ (v3 / (((v3 / ((v3) | 1))) | 1))))
                            var v60 string = (v1 + v1)
                            fmt.Println(v2)
                            _ = v59
                            _ = v60
                        }
                        fmt.Println(v3, v1)
                    }
                    _ = v57
                }
                var v61 int32 = 587
                _ = v61
            }
        }
    }
    v3 += v3
    v62 := []int{4294967295}
    var v63 int = 875
    v1 = (v1 + ("" + v1))
    var v64 int64 = 221
    var v65 string = (v1 + v1)
    v63 += (v63 >> ((1) & 63))
    var v66 uint32 = 816
    v67 := GenS3{F0: 246}
    v68 := GenS1{F0: 1, F1: 2147483647, F2: 460}
    v2 = (v2 - 1e10)
    v63 += (v63 | v63)
    v69 := GenS5{F0: (("foo bar" + "foo bar") + (("" + v65) + v1)), F1: false}
    fmt.Println(v1)
    if (!(((float32((v2 * v2)) / float32(v2)) * 0.0) == (0.0 - float32(0.1)))) {
        v1 = "\n\t"
        var v70 int8 = -27
        _ = v70
    } else {
        var v71 uint64 = 1
        if (!(!false)) {
            v65 = (((v65 + "a") + "") + "foo bar")
            _ = len(v62)
        } else {
            for fi72 := 0; fi72 < 1; fi72++ {
                v73 := GenS6{F0: v66, F1: 32767, F2: 195}
                _ = fi72
                _ = v73
            }
        }
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v62
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v69
}
