package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 int64; F2 uint32 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]uint32)
    var v3 int64 = 0
    var v4 uint = 18446744073709551615
    v5 := make(map[uint32]int)
    var v6 bool = (v1 || (v4 != (v4 >> ((v4) & 63))))
    var v7 bool = (v3 != v3)
    var v8 bool = v1
    if v6 {
        var v9 int64 = v3
        if v1 {
            var v10 uint64 = 0
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 int64 = 9223372036854775807
                v13 := []int{341}
                for fi14 := 0; fi14 < 5; fi14++ {
                    v15 := GenS1{F0: 1}
                    fmt.Println(v1, fi14, v1)
                    var v16 int8 = -128
                    if (1 > 2) {
                        var v17 float64 = -2.25
                        _ = v17
                    } else {
                        if (100 > int8(v3)) {
                            switch fi14 {
                            case 32767:
                                switch ("a" + (("" + ("a" + "foo bar")) + "")) {
                                case "hello":
                                    _ = len(v2)
                                case "foo bar":
                                    if v6 {
                                        var v18 uint32 = 1
                                        var v19 int16 = int16(v16)
                                        v20 := GenS1{F0: 786}
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    } else {
                                        var v21 uint32 = 1
                                        var v22 string = ""
                                        _ = v21
                                        _ = v22
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        if (3.14159 >= 0.1) {
                                            v7 = false
                                            if v1 {
                                                switch ((("" + "hello") + "hello") + "") {
                                                case "\n\t":
                                                    if true {
                                                        v8 = (v10 < (^(v10)))
                                                        v23 := GenS1{F0: -1}
                                                        if ("\n\t" >= ((("\n\t" + ("\n\t" + "hello")) + ((("a" + "hello") + "a") + "")) + "foo bar")) {
                                                            var v24 uint8 = 41
                                                            v12 = v12
                                                            _ = v24
                                                        } else {
                                                            switch (("a" + ("hello" + "hello")) + "hello") {
                                                            case "a":
                                                                v9 += (v9 % (((v3 % (((v12 % (((v9 << ((v4) & 63))) | 1))) | 1))) | 1))
                                                            case "":
                                                                _ = len(v5)
                                                            case "\n\t":
                                                                var v25 string = "hello"
                                                                _ = v25
                                                            default:
                                                                var v26 float32 = 1.5
                                                                _ = v26
                                                            }
                                                            v27 := v5[2]
                                                            _ = v27
                                                        }
                                                        _ = v23
                                                    } else {
                                                        if (!(!(-905 >= 866))) {
                                                            v28 := v5[2147483648]
                                                            if true {
                                                                v29 := GenS1{F0: 42}
                                                                var v30 float32 = (1.5 / 3.14159)
                                                                _ = v29
                                                                _ = v30
                                                            } else {
                                                                var v31 string = ""
                                                                _ = v31
                                                            }
                                                            v28 += fi11
                                                            _ = v28
                                                        }
                                                    }
                                                case "foo bar":
                                                    switch 604 {
                                                    case -688:
                                                        switch (^(fi14)) {
                                                        case 42:
                                                            v32 := v2[4294967295]
                                                            _ = v32
                                                        }
                                                    case -1:
                                                        var v33 uint16 = 32767
                                                        _ = v33
                                                    default:
                                                        v10 = uint64(v4)
                                                        var v34 uint64 = v10
                                                        var v35 float64 = float64(0.0)
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                }
                                                fmt.Println(v7)
                                            } else {
                                                var v36 int16 = -1
                                                v8 = v1
                                                _ = v36
                                            }
                                            v37 := GenS1{F0: 545}
                                            if (v10 > (v10 | 698)) {
                                                v1 = (!false)
                                                v1 = v6
                                            }
                                            _ = v37
                                        } else {
                                            if v6 {
                                                if ((!v1) && v6) {
                                                    fmt.Println(v3)
                                                }
                                                fmt.Println(v10)
                                                v6 = v6
                                            } else {
                                                var v38 int16 = -1
                                                v1 = (v10 <= (v10 - (^(v10))))
                                                switch "a" {
                                                case "hello":
                                                    var v39 uint16 = uint16(v10)
                                                    _ = v39
                                                }
                                                _ = v38
                                            }
                                        }
                                    case "foo bar":
                                        fmt.Println(v6, v9)
                                    }
                                    var v40 uint16 = 32767
                                    var v41 int16 = 831
                                    _ = v40
                                    _ = v41
                                }
                                fmt.Println(v3, v4)
                                if ((float32((1e10 + 0.0)) != -2.25) && false) {
                                    var v42 uint32 = 42
                                    if (v16 != -1) {
                                        var v43 uint8 = 1
                                        switch (fi14 * (fi14 % (((fi11 << (((v4 - (^(v4)))) & 63))) | 1))) {
                                        case 1:
                                            v6 = ("\n\t" >= "foo bar")
                                        case 0:
                                            v44 := GenS2{F0: 42, F1: "a"}
                                            v13 = append(v13, int(v10))
                                            _ = v44
                                        case -2147483648:
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                        }
                                        _ = v43
                                    }
                                    _ = v42
                                } else {
                                    switch -811 {
                                    case 184:
                                        switch (fi11 % (((fi11 * 32768)) | 1)) {
                                        case 405:
                                            var v45 int64 = (v3 & (v9 - (v3 | -9223372036854775808)))
                                            _ = v45
                                        case 436:
                                            delete(v2, 65535)
                                        }
                                    case -1:
                                        v46 := GenS3{F0: (^(v10)), F1: 42, F2: v10}
                                        _ = v46
                                    }
                                    fmt.Println(v8, fi14)
                                    switch fi14 {
                                    case 474:
                                        var v47 int64 = v12
                                        switch (fi14 / (((fi14 | fi14)) | 1)) {
                                        case 599:
                                            v10 += (v10 / ((4294967295) | 1))
                                            fmt.Println(v16, v6, fi11)
                                        case 86:
                                            var v48 int64 = (v12 / ((v47) | 1))
                                            _ = v48
                                        case -702:
                                            var v49 bool = v1
                                            _ = v49
                                        }
                                        var v50 int16 = 32767
                                        _ = len(v13)
                                        _ = v47
                                        _ = v50
                                    default:
                                        v3 += (v12 / (((v9 % ((v9) | 1))) | 1))
                                    }
                                }
                            case -192:
                                var v51 int8 = 0
                                _ = v51
                            case 233:
                                var v52 int16 = -71
                                _ = v52
                            }
                            var v53 uint32 = 172
                            _ = v53
                        }
                        v8 = (float64(0.0) <= 3.14159)
                    }
                    fmt.Println(v10, v4)
                    var v54 uint32 = 2147483647
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v54
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        }
        v55 := []int{32767, 32767}
        _ = v55
        _ = v9
    }
    fmt.Println(v6)
    v3 = -239
    for fi56 := 0; fi56 < 6; fi56++ {
        v7 = v7
        switch fi56 {
        case 127:
            v7 = (2 < 42)
        case -823:
            switch "foo bar" {
            case "foo bar":
                v1 = false
            case "hello":
                v57 := []int{(fi56 | (fi56 - fi56)), (fi56 ^ fi56), (fi56 << (((v4 ^ 42)) & 63))}
                _ = v57
            case "a":
                v58 := []int{}
                _ = v58
            default:
                var v59 uint16 = 131
                _ = v59
            }
        }
        _ = fi56
    }
    for fi60 := 0; fi60 < 13; fi60++ {
        v6 = v1
        v4 = (v4 >> ((v4) & 63))
        _ = fi60
    }
    v6 = v8
    v3 += v3
    fmt.Println(v8)
    v61 := GenS4{F0: (v3 + v3), F1: v4}
    switch "" {
    case "":
        var v62 int64 = (v3 + v3)
        _ = v62
    case "foo bar":
        v63 := []int{-518}
        _ = v63
    case "hello":
        v64 := v5[384]
        _ = v64
    }
    fmt.Println(v1)
    var v65 float64 = ((0.0 - ((0.1 * 0.0) * ((-2.25 / 0.1) / 1.0))) + -2.25)
    if (!((127 <= 9223372036854775807) || v6)) {
        var v66 float64 = (0.1 + 1.5)
        v8 = v6
        v67 := []int{9223372036854775807}
        _ = v66
        _ = v67
    } else {
        for fi68 := 0; fi68 < 13; fi68++ {
            v8 = v6
            for fi69 := 0; fi69 < 3; fi69++ {
                v70 := GenS2{F0: 0, F1: ("foo bar" + ((("foo bar" + ("hello" + "foo bar")) + "a") + ("\n\t" + (("" + ("hello" + "")) + ""))))}
                var v71 bool = (!v1)
                var v72 float32 = 1e10
                _ = fi69
                _ = v70
                _ = v71
                _ = v72
            }
            v6 = (-2.25 >= 1e10)
            _ = fi68
        }
        v73 := GenS5{F0: 2147483648, F1: 4294967295, F2: 393}
        for fi74 := 0; fi74 < 17; fi74++ {
            for fi75 := 0; fi75 < 5; fi75++ {
                if false {
                    v4 += 4294967295
                    switch 42 {
                    case -420:
                        if true {
                            var v76 uint64 = 1
                            fmt.Println(v4)
                            switch ("" + (((("foo bar" + ("" + "a")) + (("a" + "foo bar") + "foo bar")) + (("" + "hello") + ((("\n\t" + "hello") + "hello") + "hello"))) + "")) {
                            case "":
                                delete(v2, 2147483647)
                            case "hello":
                                delete(v2, 510)
                            case "foo bar":
                                v65 = v65
                            default:
                                var v77 bool = v1
                                _ = v77
                            }
                            _ = v76
                        } else {
                            var v78 uint64 = 18446744073709551615
                            switch (fi75 + fi74) {
                            case 70:
                                v79 := GenS6{F0: 1}
                                _ = v79
                            case -1:
                                v80 := GenS4{F0: v3, F1: v4}
                                v81 := GenS7{F0: int32(v4), F1: int64(fi74), F2: 950}
                                fmt.Println(v6)
                                _ = v80
                                _ = v81
                            default:
                                v8 = false
                            }
                            _ = v78
                        }
                    default:
                        v6 = false
                    }
                    v82 := GenS4{F0: v3, F1: 2}
                    _ = v82
                } else {
                    v83 := []int{}
                    _ = len(v2)
                    v65 = 3.14159
                    _ = v83
                }
                _ = fi75
            }
            var v84 int16 = -69
            _ = fi74
            _ = v84
        }
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v65
    _ = v7
    _ = v8
}
