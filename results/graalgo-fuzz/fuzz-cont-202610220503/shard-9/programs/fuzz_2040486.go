package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 uint }
func genFn0() uint {
    var v6 uint32 = 0
    switch 9223372036854775807 {
    case -1:
        for fi7 := 0; fi7 < 4; fi7++ {
            if true {
                fmt.Println(v6, v6)
                switch 886 {
                case -823:
                    v6 = v6
                case 288:
                    var v8 uint32 = v6
                    _ = v8
                default:
                    v9 := GenS1{F0: -32768, F1: 42}
                    _ = v9
                }
            }
            _ = fi7
        }
    }
    v10 := []int{128, 128}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v11 float32 = 3.14159
    _ = v10
    _ = v11
    _ = v6
    return 171
}
func genFn1() int32 {
    v12 := GenS2{F0: "a", F1: 127}
    v13 := GenS2{F0: "", F1: 127}
    _ = v12
    _ = v13
    return -921
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 33
    var v3 uint8 = 16
    v4 := make(map[int]bool)
    var v5 uint = 2
    v5 = (v5 % ((genFn0()) | 1))
    v14 := GenS1{F0: 1, F1: -1}
    v4[128] = (!false)
    v15 := GenS1{F0: -207, F1: 9223372036854775807}
    if false {
        var v16 float64 = 1.5
        v3 = 128
        v17 := GenS3{F0: 127}
        _ = v16
        _ = v17
    }
    v3 = (v3 + 128)
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(fi18)
        v3 += 255
        _ = fi18
    }
    delete(v4, (v2 << ((v5) & 63)))
    var v19 int32 = -2147483648
    fmt.Println(v2, v5, v2)
    switch (v2 | v2) {
    case 2147483647:
        var v20 int16 = 1
        _ = v20
    case 9223372036854775807:
        var v21 int32 = v19
        _ = v21
    case 125:
        v22 := []int{v2, 9223372036854775807}
        _ = v22
    default:
        v3 = v3
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        delete(v1, v2)
        v24 := []int{}
        var v25 int = (v2 * v2)
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 uint32 = 65535
    v3 = v3
    if ((((-9223372036854775808 > 319) || false) && (-2.25 >= (-2.25 + (((((0.1 - 0.0) + 1.5) - 3.14159) / 1.0) - 3.14159)))) || true) {
        var v27 int16 = 408
        if true {
            switch (v2 ^ (v2 - v2)) {
            case -9223372036854775808:
                switch "" {
                case "foo bar":
                    if (!((((-2.25 / 1.5) * 1.5) / 0.1) <= 0.0)) {
                        var v28 string = "a"
                        var v29 uint64 = 2
                        if false {
                            v3 = 128
                        }
                        _ = v28
                        _ = v29
                    }
                case "\n\t":
                    for fi30 := 0; fi30 < 10; fi30++ {
                        v31 := []int{fi30, (v2 >> ((genFn0()) & 63))}
                        _ = fi30
                        _ = v31
                    }
                }
            default:
                if false {
                    var v32 float32 = 0.1
                    var v33 float64 = -2.25
                    if (!(255 > (v5 / ((v5) | 1)))) {
                        var v34 float64 = v33
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v33, v27)
                        case "foo bar":
                            if true {
                                var v35 int8 = -12
                                if (!true) {
                                    delete(v1, 9223372036854775807)
                                    delete(v1, -195)
                                }
                                var v36 uint8 = v3
                                _ = v35
                                _ = v36
                            }
                        default:
                            var v37 uint32 = uint32(v3)
                            _ = v37
                        }
                        _ = v34
                    } else {
                        v38 := GenS3{F0: -47}
                        _ = v38
                    }
                    _ = v32
                    _ = v33
                } else {
                    if false {
                        var v39 string = "hello"
                        v19 = (v19 + v19)
                        v40 := []int{v2}
                        _ = v39
                        _ = v40
                    } else {
                        var v41 int32 = 1
                        v5 += (^(v5))
                        _ = len(v1)
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 14; fi42++ {
                        switch v2 {
                        case -338:
                            for fi43 := 0; fi43 < 3; fi43++ {
                                v3 -= 1
                                var v44 int8 = 12
                                v45 := GenS4{F0: (^(v3)), F1: (v3 - (v3 * 255))}
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                        case -703:
                            v46 := GenS4{F0: v3, F1: v3}
                            _ = v46
                        case 2147483647:
                            v47 := []int{(fi42 ^ (fi42 * (v2 | -2147483648))), -723, fi42}
                            _ = v47
                        }
                        switch ("" + "foo bar") {
                        case "\n\t":
                            delete(v1, -9223372036854775808)
                        case "foo bar":
                            v5 = v5
                        case "a":
                            _ = len(v4)
                        default:
                            v1[42] = v26
                        }
                        _ = fi42
                    }
                }
            }
        } else {
            v48 := []int{v2, v2, 519}
            _ = v48
        }
        if false {
            v49 := []int{}
            var v50 int32 = v19
            _ = v49
            _ = v50
        } else {
            switch "a" {
            case "foo bar":
                for fi51 := 0; fi51 < 9; fi51++ {
                    var v52 uint64 = uint64(v3)
                    _ = fi51
                    _ = v52
                }
            case "hello":
                v2 = v2
            case "\n\t":
                v53 := []int{127, (v2 - v2)}
                _ = v53
            default:
                if false {
                    var v54 string = "foo bar"
                    v19 = (v19 | (-(v19)))
                    var v55 int16 = (v27 & (-(v27)))
                    _ = v54
                    _ = v55
                } else {
                    var v56 uint8 = (^(v3))
                    if true {
                        v19 = v19
                        fmt.Println(v3)
                        v19 -= 2147483647
                    } else {
                        var v57 int32 = (^(v19))
                        _ = v57
                    }
                    v58 := []int{-2147483648}
                    _ = v56
                    _ = v58
                }
            }
            v59 := GenS2{F0: ("\n\t" + "foo bar"), F1: 1}
            _ = v59
        }
        _ = v27
    } else {
        fmt.Println(v3)
        for fi60 := 0; fi60 < 10; fi60++ {
            for fi61 := 0; fi61 < 9; fi61++ {
                v62 := GenS5{F0: (v5 << ((genFn0()) & 63))}
                if (("hello" + "foo bar") > "\n\t") {
                    v63 := GenS3{F0: 1}
                    _ = v63
                }
                _ = fi61
                _ = v62
            }
            _ = fi60
        }
        var v64 uint8 = v3
        _ = v64
    }
    v2 = (v2 - (v2 * v2))
    var v65 uint32 = v26
    v66 := v4[v2]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
}
