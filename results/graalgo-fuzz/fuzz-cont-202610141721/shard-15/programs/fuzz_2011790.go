package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint16; F1 int16; F2 int32 }
type GenS8 struct { F0 string }

func main() {
    v1 := make(map[string]string)
    var v2 int8 = -128
    v3 := make(map[uint32]int)
    var v4 uint8 = 0
    var v5 int = 4294967295
    var v6 uint = 255
    var v7 uint64 = 4294967295
    v5 += v5
    if false {
        switch "" {
        case "hello":
            v7 = (v7 / ((v7) | 1))
        case "foo bar":
            var v8 string = ""
            _ = v8
        default:
            if (((("\n\t" + "\n\t") + ("a" + (("\n\t" + "foo bar") + "\n\t"))) + "a") <= "") {
                v9 := GenS1{F0: 79, F1: 51}
                _ = v9
            }
        }
        switch "" {
        case "hello":
            v5 = v5
        }
        var v10 int8 = v2
        _ = v10
    } else {
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS1{F0: 1, F1: (v2 ^ (-(v2)))}
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 1; fi14++ {
                if false {
                    v7 = (v7 | (^(v7)))
                    switch ("\n\t" + ("foo bar" + ("a" + ""))) {
                    case "a":
                        var v15 int = fi14
                        _ = v15
                    }
                } else {
                    var v16 float64 = 1.5
                    _ = v16
                }
                switch "a" {
                case "hello":
                    switch "foo bar" {
                    case "hello":
                        var v17 uint32 = uint32(v4)
                        _ = v17
                    case "a":
                        fmt.Println(v7, v4)
                    case "\n\t":
                        fmt.Println(v4, v6)
                    }
                case "":
                    delete(v3, 1)
                default:
                    fmt.Println(v6, v6)
                }
                v18 := []int{fi11, (^(fi14))}
                _ = fi14
                _ = v18
            }
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 float64 = 0.1
                fmt.Println(fi11, v7, v5)
                var v21 int = 288
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi11
        }
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 uint = (v6 % (((^(v6))) | 1))
            var v24 uint = v6
            for fi25 := 0; fi25 < 3; fi25++ {
                if true {
                    v26 := []int{fi25, v5, v5}
                    v6 = v24
                    _ = v26
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        v27 := []int{}
                        _ = v27
                    case "":
                        v28 := GenS2{F0: uint8(v7)}
                        _ = v28
                    default:
                        delete(v3, 1)
                    }
                    v29 := GenS1{F0: -43, F1: 0}
                    _ = v29
                }
                if false {
                    v30 := v1[("foo bar" + ((((("a" + ("foo bar" + "")) + "a") + (("\n\t" + "hello") + "foo bar")) + ("a" + (("a" + "") + ""))) + (("" + "") + "a")))]
                    if true {
                        switch 128 {
                        case 4294967295:
                            var v31 bool = (121 == 1)
                            _ = v31
                        default:
                            v32 := GenS3{F0: 1, F1: (true && (v30 <= v30)), F2: (^(v7))}
                            var v33 uint = 65535
                            var v34 uint8 = v4
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    } else {
                        fmt.Println(v7, v30, v30)
                    }
                    var v35 uint64 = uint64(v4)
                    _ = v30
                    _ = v35
                } else {
                    if true {
                        if false {
                            v6 = uint(v7)
                        } else {
                            var v36 int8 = v2
                            v37 := GenS4{F0: 0}
                            _ = v36
                            _ = v37
                        }
                        var v38 int64 = 0
                        var v39 bool = (66 > 680)
                        _ = v38
                        _ = v39
                    }
                    fmt.Println(v7)
                    if ((("\n\t" + "foo bar") + "foo bar") != "") {
                        var v40 uint8 = v4
                        if (v2 > v2) {
                            _ = len(v1)
                            v41 := GenS2{F0: (v40 * (^(v4)))}
                            var v42 float32 = 1e10
                            _ = v41
                            _ = v42
                        }
                        var v43 int32 = -52
                        if false {
                            if ((1e10 / (3.14159 + 0.0)) != 1.5) {
                                switch (("" + (("foo bar" + "hello") + "hello")) + "foo bar") {
                                case "a":
                                    v44 := GenS1{F0: (v2 | 1), F1: (v2 + (v2 % (((v2 << ((1) & 63))) | 1)))}
                                    _ = v44
                                case "\n\t":
                                    v7 += v7
                                case "foo bar":
                                    switch (fi22 >> ((v6) & 63)) {
                                    case 4294967295:
                                        if true {
                                            v45 := GenS3{F0: 32767, F1: false, F2: v7}
                                            v46 := []int{(v5 + 32768)}
                                            if ((!false) || (42 >= 1)) {
                                                var v47 uint32 = 192
                                                switch (v5 & (^(v5))) {
                                                case -2147483648:
                                                    if len(v46) > 0 {
                                                        v46[0] = (fi25 >> ((uint(v40)) & 63))
                                                    }
                                                }
                                                _ = len(v46)
                                                _ = v47
                                            } else {
                                                v43 += v43
                                            }
                                            _ = v45
                                            _ = v46
                                        } else {
                                            var v48 uint = (v23 << ((v6) & 63))
                                            var v49 uint32 = 19
                                            _ = v48
                                            _ = v49
                                        }
                                    case -322:
                                        v50 := GenS2{F0: v40}
                                        _ = v50
                                    }
                                }
                                v43 -= 1
                                v51 := GenS5{F0: 0, F1: 755, F2: 79}
                                _ = v51
                            } else {
                                if true {
                                    var v52 uint32 = 358
                                    _ = v52
                                }
                                v53 := GenS5{F0: 1, F1: v7, F2: (^(v4))}
                                _ = v53
                            }
                        } else {
                            v54 := GenS6{F0: (-(v43)), F1: v2}
                            _ = v54
                        }
                        _ = v40
                        _ = v43
                    } else {
                        var v55 uint8 = v4
                        if true {
                            if false {
                                var v56 string = "hello"
                                v2 = 127
                                fmt.Println(v4)
                                _ = v56
                            }
                        } else {
                            var v57 uint8 = v4
                            _ = v57
                        }
                        v58 := GenS6{F0: 0, F1: (v2 & 1)}
                        fmt.Println(v2, v6)
                        var v59 bool = (2147483647 <= 42)
                        _ = v55
                        _ = v58
                        _ = v59
                    }
                }
                switch (^(v5)) {
                case -826:
                    v2 += v2
                }
                v60 := GenS4{F0: 65535}
                var v61 uint8 = (^(v4))
                _ = fi25
                _ = v60
                _ = v61
            }
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    v62 := GenS6{F0: -1, F1: v2}
    switch (v5 ^ 216) {
    case 279:
        switch (v5 | 996) {
        case 9223372036854775807:
            if false {
                for fi63 := 0; fi63 < 16; fi63++ {
                    v64 := GenS4{F0: 2147483648}
                    _ = fi63
                    _ = v64
                }
                v65 := GenS7{F0: 164, F1: 32767, F2: -2147483648}
                _ = v65
            }
        default:
            switch "hello" {
            case "hello":
                v66 := GenS8{F0: "a"}
                _ = v66
            case "":
                var v67 uint16 = 1
                _ = v67
            default:
                if false {
                    v6 += 18446744073709551615
                    for fi68 := 0; fi68 < 18; fi68++ {
                        var v69 uint64 = 18446744073709551615
                        for fi70 := 0; fi70 < 7; fi70++ {
                            var v71 string = "\n\t"
                            _ = fi70
                            _ = v71
                        }
                        _ = fi68
                        _ = v69
                    }
                }
            }
        }
    case -1:
        fmt.Println(v5)
    default:
        v1[("\n\t" + "hello")] = ("\n\t" + ("" + "foo bar"))
    }
    if (-5 >= v2) {
        switch "a" {
        case "foo bar":
            delete(v3, 1)
        case "\n\t":
            switch "" {
            case "\n\t":
                v72 := []int{(-(v5)), (v5 | (v5 << ((v6) & 63))), (v5 - -226), (^(v5))}
                _ = v72
            case "":
                if (v5 > (v5 << (((v6 / (((v6 - v6)) | 1))) & 63))) {
                    v7 = (v7 + v7)
                }
            case "foo bar":
                if (1 >= 678) {
                    if ((v2 > 39) || ((v5 % ((v5) | 1)) >= (v5 & (v5 >> (((^(v6))) & 63))))) {
                        if false {
                            var v73 float64 = float64(float32(-2.25))
                            var v74 int64 = 9223372036854775807
                            _ = v73
                            _ = v74
                        }
                        if (782 > -1) {
                            delete(v3, 0)
                            switch ("foo bar" + ("foo bar" + "hello")) {
                            case "a":
                                switch (v5 >> ((v6) & 63)) {
                                case 9223372036854775807:
                                    for fi75 := 0; fi75 < 17; fi75++ {
                                        var v76 int = (v5 ^ 0)
                                        _ = len(v1)
                                        _ = fi75
                                        _ = v76
                                    }
                                case 200:
                                    v4 -= v4
                                case 897:
                                    for fi77 := 0; fi77 < 2; fi77++ {
                                        var v78 uint8 = 128
                                        switch "hello" {
                                        case "hello":
                                            switch "" {
                                            case "a":
                                                fmt.Println(fi77, v78, fi77)
                                            case "":
                                                var v79 uint8 = (v4 - v78)
                                                _ = v79
                                            case "\n\t":
                                                v5 -= fi77
                                            default:
                                                v80 := GenS7{F0: 123, F1: int16(v6), F2: 0}
                                                _ = v80
                                            }
                                        case "a":
                                            var v81 float64 = 1e10
                                            _ = v81
                                        default:
                                            switch "hello" {
                                            case "hello":
                                                if (false || (!false)) {
                                                    var v82 uint32 = 2147483647
                                                    _ = v82
                                                }
                                            case "":
                                                v83 := []int{v5, (v5 - fi77), 32767}
                                                _ = v83
                                            case "foo bar":
                                                var v84 int8 = 127
                                                _ = v84
                                            default:
                                                for fi85 := 0; fi85 < 13; fi85++ {
                                                    v86 := []int{127, 4294967295}
                                                    var v87 uint64 = (^(v7))
                                                    v6 -= 117
                                                    v88 := GenS7{F0: 255, F1: 74, F2: -141}
                                                    if (828 <= 9223372036854775807) {
                                                        v89 := GenS4{F0: 0}
                                                        v86 = append(v86, v5)
                                                        _ = v89
                                                    }
                                                    _ = fi85
                                                    _ = v86
                                                    _ = v87
                                                    _ = v88
                                                }
                                            }
                                        }
                                        _ = fi77
                                        _ = v78
                                    }
                                }
                            default:
                                var v90 string = ("\n\t" + "")
                                _ = v90
                            }
                            var v91 uint16 = 65535
                            _ = v91
                        } else {
                            for fi92 := 0; fi92 < 19; fi92++ {
                                fmt.Println(v2, v6)
                                _ = fi92
                            }
                            v4 = (^(v4))
                        }
                        v93 := GenS4{F0: 2147483647}
                        _ = v93
                    }
                    fmt.Println(v6, v6, v7)
                } else {
                    v94 := []int{v5, (v5 / ((v5) | 1)), v5}
                    _ = v94
                }
            }
        }
    }
    var v95 uint32 = 626
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v7
    _ = v95
}
