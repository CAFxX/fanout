package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 int32; F2 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 int64; F1 string }
type GenS10 struct { F0 int32; F1 int16; F2 uint }
func genFn0(p0 int16, p1 uint8, p2 uint64) int {
    var v7 int8 = 1
    var v8 uint16 = 0
    v9 := []int{-891, -1, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -206
}

func main() {
    var v1 uint8 = 255
    var v2 bool = (true && true)
    v3 := make(map[int]uint32)
    var v4 bool = v2
    v5 := make(map[int]string)
    var v6 float32 = 1e10
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: 169, F1: -9223372036854775808, F2: 1}
        if (0 > -2147483648) {
            v12 := GenS1{F0: 218, F1: 42, F2: 0}
            if v2 {
                if true {
                    v13 := GenS1{F0: (v1 >> ((255) & 63)), F1: -9223372036854775808, F2: 945}
                    var v14 int32 = -2147483648
                    v4 = v4
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 19; fi15++ {
                    v4 = (255 < 510)
                    var v16 float32 = v6
                    _ = fi15
                    _ = v16
                }
                v17 := GenS1{F0: 79, F1: 119, F2: 155}
                _ = v17
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                fmt.Println(v4)
                fmt.Println(fi10, fi18, v4)
                _ = fi18
            }
            _ = v12
        } else {
            v19 := []int{fi10}
            v3[(fi10 | fi10)] = 0
            _ = v19
        }
        _ = fi10
        _ = v11
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        if v2 {
            if ((^(fi20)) == (^(fi20))) {
                if v2 {
                    if true {
                        var v21 float64 = (1.5 * 3.14159)
                        _ = v21
                    } else {
                        v22 := GenS1{F0: 0, F1: 1, F2: 1}
                        switch ("a" + (((("\n\t" + ("a" + "a")) + "hello") + "hello") + "\n\t")) {
                        case "a":
                            fmt.Println(v2, v6, v2)
                        case "":
                            v23 := GenS1{F0: (v1 ^ (^(v1))), F1: -841, F2: 99}
                            _ = v23
                        case "hello":
                            for fi24 := 0; fi24 < 16; fi24++ {
                                v25 := []int{}
                                v6 -= float32(-2.25)
                                _ = fi24
                                _ = v25
                            }
                        }
                        var v26 int16 = -1
                        _ = v22
                        _ = v26
                    }
                    if false {
                        for fi27 := 0; fi27 < 5; fi27++ {
                            var v28 uint32 = 65535
                            var v29 int = 32768
                            fmt.Println(v1, v4)
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        switch genFn0(-1, 35, 18446744073709551615) {
                        case -2147483648:
                            var v30 int8 = -105
                            _ = v30
                        }
                    } else {
                        v31 := GenS2{F0: 2, F1: 4294967295, F2: 0}
                        v32 := []int{}
                        if (true && v4) {
                            v6 += 0.0
                            fmt.Println(v4)
                        }
                        _ = v31
                        _ = v32
                    }
                }
                v33 := []int{fi20, fi20, fi20, fi20}
                _ = v33
            } else {
                if ((!(((v1 + v1) <= (v1 * 1)) && (947 != 255))) && (((("a" + "foo bar") + (("hello" + "a") + ("a" + ("foo bar" + "")))) + "foo bar") < "\n\t")) {
                    for fi34 := 0; fi34 < 7; fi34++ {
                        var v35 float64 = 1e10
                        var v36 uint16 = 65535
                        v5[fi20] = "hello"
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    var v37 uint8 = (v1 - (v1 + (v1 / ((v1) | 1))))
                    switch (fi20 / (((fi20 & fi20)) | 1)) {
                    case 922:
                        v38 := GenS3{F0: 1, F1: "hello"}
                        _ = v38
                    case 42:
                        v37 = (^(v37))
                    default:
                        var v39 int64 = -1
                        _ = v39
                    }
                    _ = v37
                } else {
                    var v40 int = 42
                    v41 := GenS4{F0: (0 >= 255)}
                    v1 = 216
                    _ = v40
                    _ = v41
                }
                v42 := []int{(fi20 << ((2) & 63))}
                _ = v42
            }
            switch ("foo bar" + "\n\t") {
            case "a":
                var v43 uint = 718
                _ = v43
            case "hello":
                v4 = true
            case "\n\t":
                var v44 float64 = -2.25
                _ = v44
            }
            v45 := []int{(fi20 - (^(fi20)))}
            _ = v45
        }
        v46 := GenS5{F0: 4294967295}
        _ = fi20
        _ = v46
    }
    if false {
        v47 := GenS3{F0: 42, F1: "foo bar"}
        v1 -= v1
        _ = v47
    }
    switch ("hello" + "a") {
    case "hello":
        v4 = ((false && (196 != 42)) && v4)
    case "a":
        if false {
            v2 = (!v4)
            v48 := GenS6{F0: int64(v1), F1: 0, F2: 9223372036854775807}
            switch -1 {
            case 0:
                var v49 uint16 = 65535
                _ = v49
            default:
                var v50 uint = 0
                _ = v50
            }
            _ = v48
        }
    case "foo bar":
        v6 = (0.0 * 1.5)
    default:
        if v2 {
            switch -2147483648 {
            case -9223372036854775808:
                switch "foo bar" {
                case "\n\t":
                    for fi51 := 0; fi51 < 3; fi51++ {
                        var v52 int64 = 659
                        var v53 float64 = 1e10
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                case "hello":
                    fmt.Println(v4)
                case "foo bar":
                    if v4 {
                        v1 = 204
                        for fi54 := 0; fi54 < 10; fi54++ {
                            for fi55 := 0; fi55 < 3; fi55++ {
                                v56 := v5[genFn0(1, (v1 & v1), 355)]
                                _ = fi55
                                _ = v56
                            }
                            fmt.Println(v2, v1)
                            _ = fi54
                        }
                    } else {
                        if (!v4) {
                            switch 2147483647 {
                            case 843:
                                if ((-1 >= 0) && (v2 && (46 == 100))) {
                                    v57 := []int{32768, 32768, 4294967295}
                                    _ = len(v5)
                                    v1 = 127
                                    _ = v57
                                }
                            case -2147483648:
                                fmt.Println(v2, v2)
                            case 2147483647:
                                var v58 string = "hello"
                                _ = v58
                            default:
                                switch 2147483647 {
                                case 9223372036854775807:
                                    var v59 int16 = 0
                                    _ = v59
                                case 42:
                                    if (true && (float64(-2.25) < 1.5)) {
                                        v60 := GenS4{F0: (42 < 18446744073709551615)}
                                        if v4 {
                                            for fi61 := 0; fi61 < 10; fi61++ {
                                                v5[0] = ("" + "a")
                                                if (v1 >= 2) {
                                                    v62 := GenS7{F0: 42}
                                                    if v2 {
                                                        var v63 uint = 255
                                                        if true {
                                                            if (-32768 != -1) {
                                                                v64 := []int{}
                                                                v65 := v3[fi61]
                                                                _ = v64
                                                                _ = v65
                                                            } else {
                                                                if (265 != 1) {
                                                                    var v66 int16 = 32767
                                                                    _ = v66
                                                                } else {
                                                                    var v67 int16 = 32767
                                                                    _ = v67
                                                                }
                                                            }
                                                            var v68 int16 = -32768
                                                            v1 -= (^(v1))
                                                            _ = v68
                                                        }
                                                        delete(v5, (fi61 - genFn0(406, v1, 42)))
                                                        _ = v63
                                                    } else {
                                                        if (0 == 255) {
                                                            var v69 float64 = 1e10
                                                            fmt.Println(v2)
                                                            _ = v69
                                                        }
                                                        v6 += ((v6 / float32(1.5)) + (((v6 + 0.1) * (float32(-2.25) * float32((3.14159 - 3.14159)))) - (v6 + v6)))
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            var v70 uint16 = 32767
                                                            _ = v70
                                                        case "a":
                                                            v71 := GenS1{F0: (^(v1)), F1: 9223372036854775807, F2: -1}
                                                            _ = v71
                                                        case "":
                                                            if v4 {
                                                                v1 = v1
                                                            } else {
                                                                v72 := []int{fi61, genFn0(-1, (^(v1)), 4294967295), (fi61 << ((318) & 63)), (fi61 * (fi61 + (fi61 | fi61)))}
                                                                v3[fi61] = 799
                                                                delete(v3, (fi61 << ((708) & 63)))
                                                                _ = v72
                                                            }
                                                        default:
                                                            v73 := []int{fi61, fi61}
                                                            _ = v73
                                                        }
                                                    }
                                                    v74 := GenS8{F0: (fi61 + fi61), F1: -687}
                                                    _ = v62
                                                    _ = v74
                                                }
                                                for fi75 := 0; fi75 < 1; fi75++ {
                                                    fmt.Println(v6, fi61)
                                                    switch ((("hello" + "hello") + "") + "\n\t") {
                                                    case "foo bar":
                                                        var v76 uint16 = 777
                                                        _ = v76
                                                    default:
                                                        v4 = v4
                                                    }
                                                    v2 = (0 >= 0)
                                                    _ = fi75
                                                }
                                                _ = fi61
                                            }
                                            v77 := []int{}
                                            v1 = v1
                                            _ = v77
                                        } else {
                                            for fi78 := 0; fi78 < 5; fi78++ {
                                                for fi79 := 0; fi79 < 8; fi79++ {
                                                    var v80 int32 = -641
                                                    switch "" {
                                                    case "hello":
                                                        var v81 bool = (18446744073709551615 <= 696)
                                                        _ = v81
                                                    case "\n\t":
                                                        v1 = (v1 ^ (v1 << ((1) & 63)))
                                                    case "foo bar":
                                                        fmt.Println(v2, fi78, fi78)
                                                    default:
                                                        v82 := GenS8{F0: (-(fi79)), F1: 1000}
                                                        _ = v82
                                                    }
                                                    fmt.Println(fi79, fi79, v1)
                                                    v6 -= -2.25
                                                    if (!true) {
                                                        v83 := GenS4{F0: v4}
                                                        _ = v83
                                                    }
                                                    _ = fi79
                                                    _ = v80
                                                }
                                                _ = fi78
                                            }
                                            for fi84 := 0; fi84 < 3; fi84++ {
                                                v85 := GenS7{F0: 506}
                                                _ = fi84
                                                _ = v85
                                            }
                                        }
                                        v2 = (v1 >= 128)
                                        _ = v60
                                    }
                                case 759:
                                    for fi86 := 0; fi86 < 2; fi86++ {
                                        v87 := []int{(fi86 << ((93) & 63))}
                                        var v88 uint64 = 1
                                        _ = fi86
                                        _ = v87
                                        _ = v88
                                    }
                                }
                            }
                            v89 := []int{}
                            _ = v89
                        } else {
                            v4 = v2
                            v2 = (1.5 > (1e10 + (3.14159 * (((1e10 - (((-2.25 - 3.14159) - 0.1) + ((3.14159 * -2.25) + 0.1))) / 1.0) / 1.0))))
                        }
                        v90 := []int{-160, 9223372036854775807, 4294967295}
                        v3[-535] = 42
                        _ = v90
                    }
                }
            default:
                switch ("a" + "foo bar") {
                case "foo bar":
                    var v91 bool = (42 >= -9223372036854775808)
                    _ = v91
                default:
                    v92 := GenS8{F0: -608, F1: 1000}
                    _ = v92
                }
            }
            v93 := GenS3{F0: 168, F1: "hello"}
            fmt.Println(v2, v4, v1)
            _ = v93
        } else {
            var v94 string = ""
            switch ("hello" + v94) {
            case "":
                fmt.Println(v4, v94)
            case "hello":
                var v95 float64 = 3.14159
                _ = v95
            }
            _ = v94
        }
    }
    if true {
        v2 = (0 > 672)
        var v96 int16 = -32768
        _ = v96
    }
    v6 = ((1.5 - (-2.25 - float32(1.5))) / v6)
    switch genFn0(-421, (v1 << ((255) & 63)), 187) {
    case -302:
        v2 = v4
    case 9223372036854775807:
        v97 := []int{genFn0(711, (v1 / (((v1 & 122)) | 1)), 18446744073709551615), 0, 968}
        _ = v97
    case 127:
        fmt.Println(v1, v2, v6)
    default:
        v98 := GenS9{F0: 42, F1: ("\n\t" + "")}
        _ = v98
    }
    v99 := GenS1{F0: (^(v1)), F1: -9223372036854775808, F2: -32768}
    switch ("hello" + "") {
    case "\n\t":
        var v100 uint8 = v1
        _ = v100
    case "":
        v101 := GenS10{F0: 1, F1: 752, F2: 0}
        _ = v101
    default:
        v102 := GenS3{F0: uint16(v1), F1: (("foo bar" + "\n\t") + ("" + "\n\t"))}
        _ = v102
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
