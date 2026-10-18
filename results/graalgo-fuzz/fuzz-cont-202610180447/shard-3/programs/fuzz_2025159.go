package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 string; F1 uint16; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 int8; F1 int8 }

func main() {
    var v1 bool = false
    var v2 uint64 = 4294967295
    var v3 bool = (!false)
    var v4 uint8 = 1
    var v5 float64 = (0.1 - 1.5)
    v5 -= float64(-2.25)
    switch "foo bar" {
    case "a":
        v1 = (!((!false) || v1))
    }
    if false {
        v1 = (605 != 4294967295)
        if (((2 == 4294967295) && v3) && v1) {
            fmt.Println(v5)
            v1 = v3
        }
        v6 := GenS1{F0: v4, F1: 65535, F2: 2147483647}
        _ = v6
    } else {
        v1 = false
    }
    v2 = (^(v2))
    v4 -= v4
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 string = ((((("a" + "a") + "hello") + "a") + (("\n\t" + "foo bar") + ((("foo bar" + "\n\t") + "a") + "hello"))) + "a")
        v9 := []int{fi7, (fi7 - fi7), (fi7 | (^(fi7))), fi7}
        _ = len(v9)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 bool = true
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 int16 = 96
            _ = fi12
            _ = v13
        }
        if ((((255 <= 447) && (v2 != (v2 & 2))) && (780 != -1)) && v3) {
            var v14 float64 = v5
            v15 := GenS1{F0: (^(v4)), F1: 32767, F2: 353}
            _ = v14
            _ = v15
        } else {
            if false {
                var v16 uint = 18446744073709551615
                v5 = (v5 / (v5 - v5))
                for fi17 := 0; fi17 < 14; fi17++ {
                    if true {
                        v18 := []int{(^(fi17)), -2147483648, fi11}
                        v18 = append(v18, 663)
                        var v19 uint = 18446744073709551615
                        _ = v18
                        _ = v19
                    } else {
                        v3 = true
                        v4 = (v4 & v4)
                        if true {
                            v2 = (v2 * (v2 >> (((v16 << ((v16) & 63))) & 63)))
                            switch "hello" {
                            case "hello":
                                if (4294967295 <= 2147483648) {
                                    var v20 int64 = -1
                                    v21 := GenS2{F0: uint32(v16), F1: 566}
                                    _ = v20
                                    _ = v21
                                }
                                if false {
                                    if v3 {
                                        if ("foo bar" >= "hello") {
                                            v22 := []int{fi11, fi17, (fi11 | fi11), 42}
                                            v22 = append(v22, fi17)
                                            _ = v22
                                        } else {
                                            v23 := []int{fi11, fi11, fi11}
                                            fmt.Println(fi11, v2, v5)
                                            _ = v23
                                        }
                                    }
                                }
                                switch 447 {
                                case 32768:
                                    v3 = true
                                default:
                                    if v1 {
                                        var v24 int8 = -1
                                        fmt.Println(fi11)
                                        _ = v24
                                    }
                                }
                            case "foo bar":
                                var v25 int = 9223372036854775807
                                _ = v25
                            default:
                                fmt.Println(v16, v10, v3)
                            }
                        }
                    }
                    switch "hello" {
                    case "":
                        var v26 int8 = 127
                        _ = v26
                    case "hello":
                        if v3 {
                            v27 := []int{(fi17 | fi11)}
                            if false {
                                v10 = (!true)
                                if (v2 >= (v2 >> ((v16) & 63))) {
                                    v3 = v3
                                }
                                if false {
                                    var v28 string = ("foo bar" + ("foo bar" + "\n\t"))
                                    v29 := GenS1{F0: (v4 * 1), F1: 2, F2: 4294967295}
                                    var v30 int16 = 356
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 bool = v3
                                    switch (^(fi17)) {
                                    case 538:
                                        var v32 int32 = -2147483648
                                        _ = v32
                                    case 938:
                                        if len(v27) >= 2 {
                                            v27 = v27[1:len(v27)]
                                        }
                                    default:
                                        var v33 uint16 = 32767
                                        _ = v33
                                    }
                                    v4 = (v4 + (v4 >> (((v16 * v16)) & 63)))
                                    _ = v31
                                }
                                var v34 int32 = -389
                                var v35 bool = v3
                                var v36 bool = true
                                if v3 {
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                    var v37 uint32 = 202
                                    if (127 != -73) {
                                        var v38 uint32 = (v37 >> (((v16 ^ (v16 | (v16 * (v16 / ((v16) | 1)))))) & 63))
                                        _ = v38
                                    } else {
                                        var v39 int64 = 9223372036854775807
                                        if v3 {
                                            v36 = v36
                                            v10 = (((^(v39)) > (v39 / ((v39) | 1))) && (v3 || true))
                                        }
                                        _ = v39
                                    }
                                    _ = v37
                                }
                                if (v4 < (v4 + v4)) {
                                    switch "a" {
                                    case "foo bar":
                                        v40 := GenS3{F0: ("foo bar" + "hello"), F1: 42, F2: v16}
                                        _ = v40
                                    case "hello":
                                        v36 = (!(!(42 != 243)))
                                    case "":
                                        if (v4 < 2) {
                                            _ = len(v27)
                                            v27 = append(v27, 42)
                                            fmt.Println(v4)
                                        } else {
                                            v41 := GenS3{F0: ("a" + ""), F1: 32767, F2: (^(v16))}
                                            var v42 uint64 = (^(v2))
                                            v36 = false
                                            _ = v41
                                            _ = v42
                                        }
                                    }
                                    v43 := GenS4{F0: (fi17 / ((2147483647) | 1))}
                                    var v44 uint16 = uint16(v16)
                                    _ = v43
                                    _ = v44
                                } else {
                                    var v45 int32 = 1
                                    var v46 bool = false
                                    _ = v45
                                    _ = v46
                                }
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            _ = v27
                        } else {
                            var v47 float32 = 1.5
                            var v48 int64 = 16
                            _ = v47
                            _ = v48
                        }
                    default:
                        if v3 {
                            v49 := []int{}
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                if ((3.14159 / (0.0 * v5)) <= -2.25) {
                                    v4 = 162
                                    v3 = true
                                } else {
                                    var v50 int16 = 32767
                                    if ((541 != 42) && (1.5 < 1.5)) {
                                        var v51 bool = (fi11 <= 42)
                                        if len(v49) > 0 {
                                            v49[0] = -810
                                        }
                                        var v52 uint16 = 65535
                                        _ = v51
                                        _ = v52
                                    }
                                    fmt.Println(fi17)
                                    _ = v50
                                }
                            default:
                                var v53 uint16 = 1
                                _ = v53
                            }
                            _ = v49
                        }
                    }
                    _ = fi17
                }
                _ = v16
            } else {
                v10 = (-1 == 56)
                v3 = true
            }
            v4 -= v4
            for fi54 := 0; fi54 < 4; fi54++ {
                if (false && true) {
                    switch fi11 {
                    case 32767:
                        var v55 uint = 18446744073709551615
                        _ = v55
                    case 2147483647:
                        if v3 {
                            v56 := []int{(^(fi54)), (fi11 + (fi11 % (((fi54 / ((fi54) | 1))) | 1))), -24}
                            v56 = append(v56, (fi54 | fi11))
                            v57 := GenS5{F0: 255}
                            _ = v56
                            _ = v57
                        } else {
                            if true {
                                v58 := GenS1{F0: (v4 + v4), F1: 42, F2: 2147483647}
                                _ = v58
                            }
                            if ((fi54 + fi11) == (fi54 >> ((1) & 63))) {
                                switch ("" + ("a" + (("a" + "hello") + ("hello" + "hello")))) {
                                case "a":
                                    v59 := []int{}
                                    _ = v59
                                default:
                                    var v60 int8 = -1
                                    _ = v60
                                }
                            } else {
                                if (v1 && true) {
                                    fmt.Println(v3)
                                } else {
                                    switch (fi54 - 4294967295) {
                                    case -232:
                                        v2 += (v2 ^ 42)
                                    case -555:
                                        var v61 uint = 18446744073709551615
                                        _ = v61
                                    default:
                                        fmt.Println(fi54, fi11)
                                    }
                                    v62 := []int{(^(fi54)), (fi54 << ((878) & 63)), (fi11 - 1)}
                                    var v63 uint32 = 624
                                    v1 = (v10 || (843 > -868))
                                    if v10 {
                                        if (0 >= -917) {
                                            var v64 bool = v10
                                            v2 = 2
                                            switch (fi11 | fi54) {
                                            case 820:
                                                v65 := GenS2{F0: v63, F1: 1}
                                                _ = v65
                                            case 184:
                                                v10 = false
                                            }
                                            v3 = (!(v64 || (!((0.0 * 1e10) > 3.14159))))
                                            _ = v64
                                        }
                                        var v66 bool = (v3 || (!(v1 || (v10 || false))))
                                        var v67 uint64 = v2
                                        _ = v66
                                        _ = v67
                                    } else {
                                        v68 := GenS2{F0: uint32(v4), F1: 32767}
                                        v69 := GenS5{F0: 2}
                                        fmt.Println(v5, fi54, v5)
                                        _ = v68
                                        _ = v69
                                    }
                                    var v70 uint = 18446744073709551615
                                    _ = v62
                                    _ = v63
                                    _ = v70
                                }
                            }
                        }
                    default:
                        if (966 == 2) {
                            var v71 bool = v1
                            _ = v71
                        }
                    }
                    v72 := []int{}
                    var v73 int8 = 42
                    _ = v72
                    _ = v73
                } else {
                    v74 := GenS3{F0: "hello", F1: 0, F2: 42}
                    _ = v74
                }
                switch 242 {
                case -236:
                    if v10 {
                        if v3 {
                            var v75 bool = v3
                            v76 := []int{fi54}
                            _ = v75
                            _ = v76
                        } else {
                            fmt.Println(v5)
                            if true {
                                var v77 int = (fi54 + fi11)
                                v78 := []int{(fi54 + fi11), fi54, (fi11 + 0)}
                                v4 -= (^(v4))
                                _ = v77
                                _ = v78
                            }
                        }
                        switch fi11 {
                        case 0:
                            if (!(0 >= -1)) {
                                var v79 int8 = 1
                                v2 = v2
                                fmt.Println(v2, v10)
                                _ = v79
                            }
                            v80 := GenS5{F0: 516}
                            if false {
                                var v81 int64 = 1
                                var v82 int16 = 136
                                _ = v81
                                _ = v82
                            }
                            _ = v80
                        default:
                            var v83 int64 = 9223372036854775807
                            _ = v83
                        }
                    }
                default:
                    switch fi54 {
                    case 1:
                        if (66 == 502) {
                            var v84 string = "foo bar"
                            switch v84 {
                            case "foo bar":
                                var v85 uint8 = v4
                                _ = v85
                            case "a":
                                if v3 {
                                    v86 := []int{-516, 4294967295, int(v2)}
                                    var v87 int32 = 928
                                    if v1 {
                                        fmt.Println(v3, v84)
                                        if v3 {
                                            v2 += v2
                                        } else {
                                            fmt.Println(v10, v2, v4)
                                        }
                                        switch fi11 {
                                        case 4294967295:
                                            switch fi54 {
                                            case 128:
                                                var v88 uint64 = (v2 % ((v2) | 1))
                                                _ = v88
                                            }
                                        }
                                        if v10 {
                                            v1 = true
                                            v89 := GenS6{F0: (v2 - v2), F1: (v3 || false), F2: 1}
                                            v90 := GenS2{F0: 2, F1: 982}
                                            _ = v89
                                            _ = v90
                                        }
                                    } else {
                                        if true {
                                            v87 -= int32(v4)
                                            var v91 float64 = v5
                                            v5 = v91
                                            _ = v91
                                        } else {
                                            v92 := GenS7{F0: (v84 + v84), F1: 2147483647}
                                            _ = v92
                                        }
                                    }
                                    var v93 string = (((v84 + v84) + v84) + "foo bar")
                                    _ = v86
                                    _ = v87
                                    _ = v93
                                }
                                v94 := GenS2{F0: 65535, F1: 255}
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v95 float64 = ((v5 / (v5 + 3.14159)) + 0.0)
                                        v96 := []int{}
                                        switch v84 {
                                        case "foo bar":
                                            v10 = (-2147483648 > -2147483648)
                                        case "\n\t":
                                            var v97 float32 = (0.1 * 1e10)
                                            _ = v97
                                        }
                                        v96 = append(v96, 133)
                                        v96 = append(v96, (-(fi54)))
                                        if len(v96) >= 2 {
                                            v96 = v96[1:len(v96)]
                                        }
                                        _ = v95
                                        _ = v96
                                    case "":
                                        fmt.Println(v10, v1)
                                    }
                                    var v98 uint = 2
                                    var v99 uint = (v98 | (v98 * v98))
                                    v84 = v84
                                    v100 := GenS2{F0: 65535, F1: 32767}
                                    if (!((42 != -32768) || (0.1 == 1e10))) {
                                        v99 -= (v99 ^ 888)
                                    }
                                    _ = v100
                                    _ = v98
                                    _ = v99
                                }
                                var v101 string = (v84 + v84)
                                v3 = true
                                _ = v101
                                _ = v94
                            }
                            _ = v84
                        } else {
                            v10 = (701 < 4294967295)
                            if ((!(v5 <= ((float64(0.1) * ((((float64(((0.0 + 1e10) + 1e10)) + v5) * v5) * 0.1) + 3.14159)) + ((v5 - v5) / (float64(float32((v5 * (float64(float32(v5)) + v5)))) / v5))))) || false) {
                                if v1 {
                                    switch ("hello" + "hello") {
                                    case "a":
                                        v102 := []int{(fi54 << ((0) & 63))}
                                        _ = v102
                                    case "hello":
                                        var v103 uint8 = (v4 ^ v4)
                                        _ = v103
                                    default:
                                        v2 = 4294967295
                                    }
                                    fmt.Println(v1, v2, v5)
                                    if (false || false) {
                                        v3 = false
                                    } else {
                                        var v104 float64 = (((0.0 + (float64(float32(((v5 / v5) / 1.5))) - 3.14159)) * ((((float64(((0.1 + (-2.25 - 1e10)) + 1e10)) + (float64(1.5) - v5)) + 0.1) + (v5 - 0.1)) * v5)) * v5)
                                        var v105 float32 = 0.0
                                        v5 -= v5
                                        _ = v104
                                        _ = v105
                                    }
                                }
                                var v106 int8 = 6
                                _ = v106
                            } else {
                                var v107 uint16 = 0
                                v2 = v2
                                _ = v107
                            }
                            v108 := []int{(fi11 | (-(fi54))), 9223372036854775807, -9223372036854775808}
                            _ = v108
                        }
                    case 9223372036854775807:
                        var v109 float64 = -2.25
                        _ = v109
                    case 42:
                        v110 := GenS8{F0: 9223372036854775807, F1: v4}
                        _ = v110
                    }
                }
                v111 := GenS9{F0: int8(v2), F1: 100}
                _ = fi54
                _ = v111
            }
        }
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
