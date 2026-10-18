package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint8; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int64; F2 int16 }
type GenS5 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 uint; F1 int8; F2 uint32 }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 9; fi6++ {
        switch "a" {
        case "\n\t":
            switch fi6 {
            case 547:
                v7 := GenS1{F0: 42, F1: 2, F2: 137}
                _ = v7
            case 4294967295:
                var v8 float32 = 0.1
                _ = v8
            case 0:
                v9 := []int{}
                _ = v9
            default:
                var v10 float32 = 3.14159
                _ = v10
            }
        case "foo bar":
            var v11 uint32 = 900
            _ = v11
        case "":
            v12 := GenS2{F0: 9223372036854775807, F1: 0, F2: "foo bar"}
            _ = v12
        default:
            switch "a" {
            case "foo bar":
                var v13 uint32 = 42
                _ = v13
            case "hello":
                v14 := GenS3{F0: -128}
                _ = v14
            case "":
                fmt.Println(fi6)
            default:
                v15 := GenS4{F0: 2, F1: -360, F2: -810}
                _ = v15
            }
        }
        v16 := []int{}
        v16 = append(v16, 32768)
        _ = fi6
        _ = v16
    }
    var v17 int = -660
    for fi18 := 0; fi18 < 18; fi18++ {
        if (!(!false)) {
            v19 := GenS5{F0: 42, F1: 89, F2: 659}
            fmt.Println(v17, v17)
            if (1 > 0) {
                v17 = v17
                for fi20 := 0; fi20 < 7; fi20++ {
                    v21 := []int{}
                    fmt.Println(fi20)
                    v17 = 32768
                    _ = fi20
                    _ = v21
                }
                v22 := GenS6{F0: -1}
                _ = v22
            }
            _ = v19
        } else {
            switch (("\n\t" + "foo bar") + "a") {
            case "a":
                var v23 uint8 = 69
                _ = v23
            case "":
                for fi24 := 0; fi24 < 5; fi24++ {
                    v25 := []int{fi24}
                    if len(v25) > 0 {
                        v25[0] = 127
                    }
                    var v26 int64 = 42
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            for fi27 := 0; fi27 < 12; fi27++ {
                v17 = (fi18 + fi18)
                v17 = (fi18 / (((fi27 / (((fi27 >> ((18446744073709551615) & 63))) | 1))) | 1))
                v28 := GenS3{F0: -128}
                _ = fi27
                _ = v28
            }
            var v29 int8 = -128
            _ = v29
        }
        _ = fi18
    }
    v30 := []int{v17}
    for fi31 := 0; fi31 < 15; fi31++ {
        for fi32 := 0; fi32 < 13; fi32++ {
            v30 = append(v30, (v17 << ((4294967295) & 63)))
            var v33 int64 = -615
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    _ = v17
    _ = v30
    return 1
}
func genFn1() uint64 {
    var v34 bool = false
    v35 := GenS7{F0: ("" + ("hello" + "\n\t"))}
    _ = v34
    _ = v35
    return 496
}

func main() {
    var v1 int64 = -546
    var v2 int8 = 100
    var v3 int16 = -317
    var v4 bool = (false || (42 > (v1 % ((-786) | 1))))
    v5 := make(map[int]string)
    switch "\n\t" {
    case "foo bar":
        for fi36 := 0; fi36 < 4; fi36++ {
            for fi37 := 0; fi37 < 19; fi37++ {
                switch "hello" {
                case "a":
                    var v38 uint16 = 942
                    _ = v38
                case "hello":
                    var v39 float32 = 1e10
                    _ = v39
                default:
                    v4 = (!v4)
                }
                v40 := []int{-300, fi36, 0}
                _ = fi37
                _ = v40
            }
            v4 = false
            var v41 float32 = 1.5
            _ = fi36
            _ = v41
        }
    case "hello":
        switch 9223372036854775807 {
        case -9223372036854775808:
            v42 := v5[878]
            _ = v42
        }
    case "":
        v43 := GenS5{F0: -66, F1: 1, F2: 822}
        _ = v43
    default:
        switch 128 {
        case 32767:
            if true {
                v5[32767] = "foo bar"
                v44 := []int{127, 911, 9223372036854775807}
                _ = v44
            }
        case -89:
            if (!v4) {
                fmt.Println(v1, v4, v3)
                v5[1] = ("hello" + ("" + "hello"))
                if (2 != 32767) {
                    if v4 {
                        v45 := []int{21, 128, 42}
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 15; fi46++ {
                        if (0.0 != (1e10 * 1e10)) {
                            v4 = v4
                            v3 += -1
                        } else {
                            v3 -= 0
                        }
                        _ = fi46
                    }
                    var v47 uint16 = 42
                    _ = v47
                } else {
                    for fi48 := 0; fi48 < 8; fi48++ {
                        v5[205] = ("foo bar" + "")
                        _ = fi48
                    }
                }
            } else {
                fmt.Println(v1)
                delete(v5, 127)
            }
        default:
            var v49 float64 = (1e10 + (0.1 - (((3.14159 * 0.0) - 1.5) + (float64((((1e10 + -2.25) / 1.0) / (0.1 + ((1e10 - 1.5) * (float32((((-2.25 * -2.25) + 0.1) / 0.1)) - 1e10))))) / 1.5))))
            _ = v49
        }
    }
    switch "hello" {
    case "\n\t":
        v1 += 545
    case "foo bar":
        if true {
            switch 442 {
            case 128:
                for fi50 := 0; fi50 < 17; fi50++ {
                    v51 := []int{fi50, fi50}
                    v52 := GenS7{F0: "a"}
                    _ = len(v51)
                    _ = fi50
                    _ = v51
                    _ = v52
                }
            case 4294967295:
                v4 = v4
            }
            switch -9223372036854775808 {
            case 42:
                fmt.Println(v1, v4)
            default:
                v4 = (v4 && ((3.14159 / 1.5) >= -2.25))
            }
        } else {
            fmt.Println(v2, v1, v1)
        }
    case "a":
        if v4 {
            v1 = 9223372036854775807
            fmt.Println(v1, v3)
            delete(v5, 1)
        } else {
            v1 += v1
            v1 = (v1 << ((18446744073709551615) & 63))
            switch "\n\t" {
            case "hello":
                switch "a" {
                case "":
                    v53 := GenS8{F0: 619, F1: (v1 & v1)}
                    _ = v53
                case "a":
                    if (!true) {
                        delete(v5, -2147483648)
                        fmt.Println(v2, v3, v3)
                    } else {
                        for fi54 := 0; fi54 < 13; fi54++ {
                            for fi55 := 0; fi55 < 10; fi55++ {
                                v56 := []int{-65}
                                _ = fi55
                                _ = v56
                            }
                            if v4 {
                                var v57 int = 128
                                _ = v57
                            }
                            switch "hello" {
                            case "a":
                                v5[(fi54 | -1)] = "\n\t"
                            case "":
                                fmt.Println(fi54, v3)
                            case "\n\t":
                                switch (fi54 % (((fi54 + fi54)) | 1)) {
                                case 954:
                                    for fi58 := 0; fi58 < 18; fi58++ {
                                        fmt.Println(v3)
                                        if false {
                                            if (1 < 92) {
                                                v59 := v5[fi58]
                                                _ = v59
                                            }
                                            var v60 uint64 = 42
                                            v61 := []int{}
                                            _ = v60
                                            _ = v61
                                        } else {
                                            _ = len(v5)
                                            v1 += 1
                                            if v4 {
                                                v4 = v4
                                                v4 = (v4 && ((v4 && (v1 <= v1)) && (fi54 < int(v3))))
                                                v62 := GenS6{F0: -1}
                                                v63 := []int{}
                                                v63 = append(v63, fi54)
                                                v64 := GenS1{F0: 80, F1: genFn1(), F2: 682}
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                            } else {
                                                v5[67] = "\n\t"
                                                var v65 uint8 = 216
                                                var v66 string = "foo bar"
                                                var v67 int = 605
                                                v5[(v67 / (((fi54 >> ((255) & 63))) | 1))] = v66
                                                _ = v65
                                                _ = v66
                                                _ = v67
                                            }
                                        }
                                        if v4 {
                                            var v68 float32 = -2.25
                                            v69 := GenS9{F0: 1, F1: v2, F2: 2147483647}
                                            _ = v68
                                            _ = v69
                                        }
                                        _ = fi58
                                    }
                                case 42:
                                    switch -2147483648 {
                                    case 4294967295:
                                        switch (fi54 >> ((18446744073709551615) & 63)) {
                                        case 127:
                                            var v70 int32 = 1
                                            _ = v70
                                        default:
                                            v3 += (v3 - v3)
                                        }
                                    }
                                default:
                                    var v71 uint8 = 1
                                    _ = v71
                                }
                            }
                            _ = fi54
                        }
                        v1 = (v1 & v1)
                        var v72 float32 = 0.0
                        _ = v72
                    }
                case "hello":
                    fmt.Println(v4, v3, v3)
                default:
                    switch 0 {
                    case 1:
                        var v73 uint64 = 18446744073709551615
                        _ = v73
                    case 72:
                        v2 = v2
                    case -2147483648:
                        v74 := []int{4294967295, 32767, 4294967295, -1}
                        _ = v74
                    }
                }
            case "\n\t":
                if (("\n\t" + "") == "hello") {
                    var v75 float64 = 1e10
                    _ = v75
                }
            case "a":
                v76 := GenS4{F0: 2147483647, F1: -1, F2: v3}
                _ = v76
            }
        }
    }
    switch 32768 {
    case 72:
        fmt.Println(v1)
    case -152:
        for fi77 := 0; fi77 < 4; fi77++ {
            var v78 string = "a"
            v78 = v78
            _ = fi77
            _ = v78
        }
    case 719:
        switch 42 {
        case 0:
            var v79 uint64 = genFn1()
            _ = v79
        }
    default:
        switch 9223372036854775807 {
        case 127:
            fmt.Println(v4)
        case 9223372036854775807:
            v80 := []int{-827, 4294967295, 328}
            _ = v80
        }
    }
    var v81 uint16 = 1
    var v82 uint64 = 2
    var v83 int64 = v1
    var v84 int32 = 162
    var v85 int16 = (v3 << ((uint(v82)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v82
    _ = v83
    _ = v84
    _ = v85
}
