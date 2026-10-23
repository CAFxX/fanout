package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 int32 }
func genFn0() int {
    var v7 int16 = 42
    var v8 int32 = -2147483648
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int16 = (v7 + -1)
        for fi11 := 0; fi11 < 19; fi11++ {
            v8 = v8
            var v12 bool = false
            fmt.Println(v7, fi9, fi11)
            v7 = v10
            _ = fi11
            _ = v12
        }
        switch fi9 {
        case -9223372036854775808:
            v13 := []int{(fi9 - 4294967295), 756}
            _ = v13
        default:
            var v14 uint64 = 42
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    v8 = -2147483648
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 string = (("hello" + ("" + "a")) + "foo bar")
    var v2 uint64 = 222
    var v3 float64 = (0.1 - 0.1)
    var v4 int16 = -913
    var v5 int16 = 42
    var v6 int16 = (v4 % (((v4 << ((1) & 63))) | 1))
    switch 358 {
    case 128:
        var v15 string = ("hello" + "\n\t")
        _ = v15
    case 127:
        var v16 int8 = -128
        _ = v16
    default:
        v17 := GenS1{F0: v2}
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 float32 = (1e10 * 0.1)
            v5 += -425
            var v21 int32 = -1
            _ = fi19
            _ = v20
            _ = v21
        }
        if true {
            fmt.Println(v1, fi18)
            var v22 float64 = v3
            var v23 uint32 = 0
            _ = v22
            _ = v23
        } else {
            if false {
                fmt.Println(v4, v3)
                v6 -= v6
            } else {
                v6 = v6
                v24 := GenS1{F0: (v2 ^ (v2 | (v2 / ((18446744073709551615) | 1))))}
                _ = v24
            }
        }
        v1 = (v1 + (v1 + "hello"))
        _ = fi18
    }
    if false {
        if true {
            for fi25 := 0; fi25 < 4; fi25++ {
                for fi26 := 0; fi26 < 14; fi26++ {
                    v27 := []int{}
                    v28 := GenS2{F0: 721, F1: v5, F2: 0}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                switch "foo bar" {
                case "foo bar":
                    var v29 string = v1
                    _ = v29
                case "a":
                    switch ((v1 + v1) + "") {
                    case "foo bar":
                        var v30 int64 = 9223372036854775807
                        _ = v30
                    default:
                        switch fi25 {
                        case 4294967295:
                            if ((!false) || (1 != 65535)) {
                                for fi31 := 0; fi31 < 17; fi31++ {
                                    v32 := []int{-499}
                                    v5 = 451
                                    _ = fi31
                                    _ = v32
                                }
                            } else {
                                v1 = v1
                                v33 := GenS1{F0: (^(v2))}
                                v6 = (^(v6))
                                _ = v33
                            }
                        case 365:
                            v34 := GenS3{F0: 255, F1: 42}
                            _ = v34
                        case 0:
                            switch (fi25 & genFn0()) {
                            case -110:
                                var v35 int64 = 883
                                _ = v35
                            default:
                                var v36 uint = 4294967295
                                _ = v36
                            }
                        }
                    }
                case "":
                    switch (v1 + v1) {
                    case "":
                        var v37 int32 = 0
                        _ = v37
                    }
                default:
                    var v38 uint8 = 1
                    _ = v38
                }
                _ = fi25
            }
        } else {
            if (!(975 <= 0)) {
                if true {
                    var v39 float32 = (3.14159 - 3.14159)
                    if true {
                        var v40 uint8 = 189
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v39)
                        case "":
                            if true {
                                v41 := GenS4{F0: genFn0(), F1: 813}
                                _ = v41
                            } else {
                                switch 32767 {
                                case 42:
                                    fmt.Println(v39)
                                case -9223372036854775808:
                                    if ((!(507 >= 42)) && true) {
                                        v42 := GenS3{F0: (v40 >> ((804) & 63)), F1: -118}
                                        v43 := GenS5{F0: v40, F1: 2147483647}
                                        _ = v42
                                        _ = v43
                                    }
                                case -2147483648:
                                    v44 := GenS4{F0: 42, F1: 42}
                                    _ = v44
                                }
                            }
                        case "a":
                            v40 += (v40 % ((v40) | 1))
                        default:
                            var v45 int = -9223372036854775808
                            _ = v45
                        }
                        if false {
                            v46 := []int{}
                            _ = v46
                        }
                        _ = v40
                    } else {
                        v4 -= -32768
                        v47 := []int{genFn0()}
                        var v48 uint64 = v2
                        _ = v47
                        _ = v48
                    }
                    fmt.Println(v4, v1, v1)
                    _ = v39
                }
            }
            var v49 int16 = 953
            var v50 uint = 255
            _ = v49
            _ = v50
        }
        if true {
            v3 = 1.5
            fmt.Println(v6, v6)
        }
        if (!(216 < 17)) {
            if true {
                fmt.Println(v6, v5, v3)
                var v51 float32 = -2.25
                _ = v51
            }
            for fi52 := 0; fi52 < 12; fi52++ {
                var v53 bool = true
                _ = fi52
                _ = v53
            }
        } else {
            var v54 int16 = -1
            _ = v54
        }
    }
    switch genFn0() {
    case 918:
        if (42 == 92) {
            v55 := []int{9223372036854775807, -176, -445}
            _ = v55
        }
    default:
        if (true || true) {
            var v56 int64 = 1
            var v57 uint16 = 255
            _ = v56
            _ = v57
        } else {
            for fi58 := 0; fi58 < 1; fi58++ {
                v59 := GenS6{F0: true, F1: 65535}
                v60 := GenS7{F0: (v2 >= v2), F1: 42}
                v61 := []int{127, (fi58 & (-(fi58))), fi58, (fi58 / ((int(v2)) | 1))}
                _ = fi58
                _ = v59
                _ = v60
                _ = v61
            }
            if (!true) {
                if (false || false) {
                    v4 += v4
                    v4 -= 1
                } else {
                    v1 = ("foo bar" + (((v1 + "foo bar") + ((v1 + v1) + v1)) + ((((v1 + (("a" + (v1 + "foo bar")) + "")) + v1) + (v1 + v1)) + v1)))
                    for fi62 := 0; fi62 < 19; fi62++ {
                        v6 = 463
                        _ = fi62
                    }
                }
                switch 32768 {
                case -412:
                    v63 := GenS8{F0: 0}
                    _ = v63
                default:
                    v64 := GenS9{F0: v1, F1: 0}
                    _ = v64
                }
            } else {
                v65 := GenS6{F0: true, F1: 43}
                fmt.Println(v4, v4, v5)
                _ = v65
            }
        }
    }
    if true {
        var v66 bool = (-1 == -1)
        var v67 uint32 = 2
        _ = v66
        _ = v67
    }
    var v68 int64 = 607
    if (2 <= 22) {
        var v69 string = v1
        v70 := GenS3{F0: 128, F1: -22}
        _ = v69
        _ = v70
    }
    if (((v68 != 736) || ((-728 < 0) && ("foo bar" >= ""))) || (false && true)) {
        var v71 int32 = -864
        _ = v71
    } else {
        if true {
            switch v1 {
            case "\n\t":
                v72 := []int{}
                _ = v72
            case "a":
                var v73 float32 = 3.14159
                _ = v73
            }
            for fi74 := 0; fi74 < 11; fi74++ {
                for fi75 := 0; fi75 < 1; fi75++ {
                    var v76 uint8 = 42
                    v77 := []int{}
                    _ = fi75
                    _ = v76
                    _ = v77
                }
                _ = fi74
            }
        } else {
            var v78 int32 = 2147483647
            v79 := GenS3{F0: 253, F1: 1}
            _ = v78
            _ = v79
        }
    }
    v80 := GenS9{F0: (v1 + v1), F1: 0}
    v68 -= v68
    v81 := []int{-1, 409}
    v2 = v2
    v81 = append(v81, -2147483648)
    var v82 int8 = 97
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v81:", len(v81), v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v80
    _ = v81
    _ = v82
}
