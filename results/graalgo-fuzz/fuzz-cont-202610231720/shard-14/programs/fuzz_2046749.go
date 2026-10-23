package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 uint64; F1 int16 }
type GenS11 struct { F0 int8; F1 uint64; F2 int16 }
type GenS12 struct { F0 int16; F1 uint64; F2 int64 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 string; F1 int64 }
func genFn0(p0 int8) int16 {
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := GenS1{F0: 0, F1: 127}
        v6 := []int{int(p0), (fi4 * (-(fi4)))}
        for fi7 := 0; fi7 < 8; fi7++ {
            p0 -= p0
            _ = fi7
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    var v8 bool = (0 <= 42)
    var v9 float64 = 3.14159
    p0 = 0
    _ = p0
    _ = v8
    _ = v9
    return -626
}

func main() {
    v1 := []int{32767, 4294967295, 336}
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    v10 := GenS1{F0: 127, F1: 29}
    var v11 uint = 938
    var v12 string = (("" + "hello") + "a")
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS2{F0: (true && (!(2147483647 > 113))), F1: -672}
        if false {
            v15 := v3[2147483648]
            _ = v15
        }
        v16 := GenS3{F0: 15, F1: -2147483648}
        _ = fi13
        _ = v14
        _ = v16
    }
    if false {
        switch "hello" {
        case "":
            for fi17 := 0; fi17 < 18; fi17++ {
                for fi18 := 0; fi18 < 12; fi18++ {
                    switch "foo bar" {
                    case "foo bar":
                        var v19 int = fi18
                        _ = v19
                    case "":
                        switch v12 {
                        case "a":
                            v20 := GenS2{F0: false, F1: fi18}
                            _ = v20
                        case "\n\t":
                            v21 := v2[2147483648]
                            v12 = v12
                            _ = v21
                        }
                    }
                    fmt.Println(v11, fi18, v11)
                    _ = fi18
                }
                v11 += v11
                _ = fi17
            }
        default:
            for fi22 := 0; fi22 < 11; fi22++ {
                v1 = append(v1, fi22)
                _ = fi22
            }
        }
    }
    v12 = (v12 + ("\n\t" + ((v12 + "a") + v12)))
    fmt.Println(v12)
    switch 2147483647 {
    case -37:
        switch -2147483648 {
        case 394:
            var v23 uint64 = 540
            _ = v23
        default:
            if (false && (-32768 <= 1)) {
                switch v12 {
                case "foo bar":
                    var v24 int16 = genFn0(100)
                    _ = v24
                case "hello":
                    _ = len(v3)
                default:
                    var v25 uint32 = 0
                    _ = v25
                }
                for fi26 := 0; fi26 < 14; fi26++ {
                    switch 133 {
                    case -430:
                        v12 = (v12 + v12)
                    case -898:
                        v11 = 18446744073709551615
                    case -9223372036854775808:
                        v27 := GenS2{F0: (42 == 103), F1: (fi26 + 834)}
                        _ = v27
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    var v28 uint8 = 128
                    switch v12 {
                    case "foo bar":
                        v28 -= (v28 >> (((v11 % ((v11) | 1))) & 63))
                    }
                    _ = fi26
                    _ = v28
                }
                var v29 uint64 = 1
                _ = v29
            }
        }
    }
    var v30 int64 = 183
    v11 = 511
    for fi31 := 0; fi31 < 9; fi31++ {
        if true {
            for fi32 := 0; fi32 < 14; fi32++ {
                v3[2147483647] = (v12 + (v12 + ""))
                v30 += v30
                switch (((v12 + ((v12 + "") + "a")) + "hello") + v12) {
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "hello":
                    switch ("foo bar" + ("foo bar" + v12)) {
                    case "":
                        if false {
                            if false {
                                switch fi31 {
                                case 1:
                                    switch fi32 {
                                    case -691:
                                        fmt.Println(v30)
                                    case 128:
                                        v33 := GenS4{F0: 100, F1: true}
                                        _ = v33
                                    }
                                case 32767:
                                    v3[65535] = (v12 + v12)
                                }
                            } else {
                                v34 := GenS5{F0: 4294967295}
                                var v35 string = ""
                                v30 -= 377
                                _ = v34
                                _ = v35
                            }
                            if true {
                                switch ("" + "hello") {
                                case "":
                                    var v36 float32 = 1.5
                                    _ = v36
                                case "a":
                                    v37 := GenS6{F0: 42, F1: -138}
                                    _ = v37
                                case "\n\t":
                                    var v38 float32 = 3.14159
                                    _ = v38
                                default:
                                    v12 = (v12 + (v12 + (v12 + "foo bar")))
                                }
                            } else {
                                v30 = (v30 << ((42) & 63))
                            }
                        }
                    case "\n\t":
                        if true {
                            var v39 int32 = -300
                            v30 += (v30 ^ 1)
                            _ = v39
                        } else {
                            v30 = -9223372036854775808
                            if len(v1) > 0 {
                                v1[0] = (fi31 + (fi32 << ((v11) & 63)))
                            }
                            v3[42] = (v12 + "")
                        }
                    case "hello":
                        if (42 == 76) {
                            if (v11 <= v11) {
                                fmt.Println(fi32)
                            }
                            switch v12 {
                            case "":
                                if (!(true || ((42 >= 2147483647) || (32767 < 909)))) {
                                    switch (fi32 & (fi32 << ((v11) & 63))) {
                                    case 42:
                                        _ = len(v3)
                                        var v40 uint16 = 565
                                        var v41 string = "foo bar"
                                        _ = v40
                                        _ = v41
                                    case -511:
                                        v42 := GenS7{F0: 10, F1: 733}
                                        _ = v42
                                    case -663:
                                        if (225 >= uint64(v11)) {
                                            v12 = v12
                                            v12 = "\n\t"
                                            v12 = ((v12 + "foo bar") + (v12 + (((v12 + "\n\t") + v12) + "\n\t")))
                                            v43 := GenS8{F0: 32767}
                                            _ = v43
                                        } else {
                                            if (true || false) {
                                                if (-2.25 > (3.14159 / 1e10)) {
                                                    v12 = (v12 + v12)
                                                } else {
                                                    var v44 float32 = (1e10 * float32(((float64(3.14159) * 3.14159) - (float64((1.5 + 0.1)) * 0.1))))
                                                    if (-315 <= v30) {
                                                        if (670 != 605) {
                                                            v11 += v11
                                                            if len(v1) > 0 {
                                                                v1[0] = fi31
                                                            }
                                                            if (2 < 255) {
                                                                fmt.Println(v44, fi32, fi32)
                                                                if true {
                                                                    var v45 bool = true
                                                                    v44 -= (float32((((3.14159 * 0.1) / float64(v44)) - 3.14159)) / v44)
                                                                    delete(v2, 43)
                                                                    v46 := GenS4{F0: 0, F1: (!(1e10 <= 3.14159))}
                                                                    _ = v45
                                                                    _ = v46
                                                                } else {
                                                                    v47 := GenS4{F0: 127, F1: ((4294967295 != (-(fi32))) && true)}
                                                                    fmt.Println(v12, fi31)
                                                                    v44 = 1.5
                                                                    _ = v47
                                                                }
                                                            } else {
                                                                if ((0 != (v11 << (((v11 | (v11 ^ v11))) & 63))) && false) {
                                                                    v44 += v44
                                                                    v12 = ("a" + "\n\t")
                                                                } else {
                                                                    var v48 uint = 255
                                                                    var v49 uint32 = 65535
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                                v11 = v11
                                                                v12 = "\n\t"
                                                            }
                                                        } else {
                                                            var v50 uint32 = 2147483647
                                                            v51 := GenS2{F0: (false || (false && false)), F1: (fi31 | (fi32 + (fi32 >> (((v11 % ((v11) | 1))) & 63))))}
                                                            v44 = ((v44 - (v44 - (((float32(1.5) - v44) / v44) - v44))) * 3.14159)
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    } else {
                                                        var v52 string = (("" + v12) + v12)
                                                        _ = v52
                                                    }
                                                    v53 := GenS9{F0: 42, F1: 255}
                                                    _ = v44
                                                    _ = v53
                                                }
                                            } else {
                                                if (true || false) {
                                                    v12 = v12
                                                    delete(v3, 65535)
                                                    var v54 string = (("foo bar" + "\n\t") + "a")
                                                    _ = len(v1)
                                                    _ = v54
                                                } else {
                                                    if (115 == -12) {
                                                        v12 = "foo bar"
                                                    } else {
                                                        fmt.Println(fi32, v11, v30)
                                                    }
                                                    var v55 bool = false
                                                    _ = v55
                                                }
                                                v11 = (v11 | 4294967295)
                                                if false {
                                                    fmt.Println(v12)
                                                    v56 := GenS9{F0: 42, F1: 65535}
                                                    v30 = v30
                                                    _ = v56
                                                } else {
                                                    switch 4294967295 {
                                                    case 977:
                                                        switch (v12 + ("a" + "\n\t")) {
                                                        case "hello":
                                                            fmt.Println(v30)
                                                        case "":
                                                            delete(v3, 65535)
                                                        case "foo bar":
                                                            var v57 string = "\n\t"
                                                            _ = v57
                                                        }
                                                    }
                                                    var v58 bool = (false && true)
                                                    switch -1 {
                                                    case -753:
                                                        switch 4294967295 {
                                                        case 42:
                                                            var v59 uint = v11
                                                            _ = v59
                                                        case -312:
                                                            var v60 uint8 = 117
                                                            _ = v60
                                                        }
                                                    case 560:
                                                        _ = len(v2)
                                                    }
                                                    _ = v58
                                                }
                                                v11 -= 471
                                                v12 = (v12 + "hello")
                                                v3[993] = "\n\t"
                                            }
                                        }
                                    }
                                    v61 := GenS10{F0: 4294967295, F1: -421}
                                    var v62 uint64 = 18446744073709551615
                                    var v63 float64 = 0.1
                                    v11 += (v11 | 876)
                                    fmt.Println(v12, fi31, v11)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                }
                            case "a":
                                var v64 uint64 = 0
                                _ = v64
                            case "foo bar":
                                if false {
                                    var v65 int16 = 1
                                    var v66 uint64 = 18446744073709551615
                                    _ = v65
                                    _ = v66
                                }
                            }
                            fmt.Println(v30)
                        }
                    default:
                        v11 = (v11 + (v11 & v11))
                    }
                case "\n\t":
                    if (0 != 42) {
                        v67 := GenS11{F0: 42, F1: 201, F2: 42}
                        v68 := v3[0]
                        var v69 uint16 = 60
                        _ = v67
                        _ = v68
                        _ = v69
                    } else {
                        var v70 bool = (true || false)
                        _ = len(v2)
                        v71 := GenS12{F0: int16(fi31), F1: 0, F2: -9223372036854775808}
                        _ = v70
                        _ = v71
                    }
                default:
                    var v72 string = v12
                    _ = v72
                }
                _ = fi32
            }
            for fi73 := 0; fi73 < 9; fi73++ {
                v12 = "\n\t"
                if true {
                    if (255 >= 128) {
                        switch v12 {
                        case "":
                            v30 -= v30
                        case "a":
                            var v74 bool = false
                            _ = v74
                        case "hello":
                            if false {
                                v30 += v30
                            } else {
                                v11 += v11
                            }
                        }
                        var v75 bool = ((v11 - v11) != 13)
                        _ = v75
                    }
                }
                v30 = v30
                if true {
                    var v76 float64 = (3.14159 * (0.0 * 1e10))
                    _ = v76
                }
                _ = fi73
            }
        } else {
            var v77 float64 = -2.25
            for fi78 := 0; fi78 < 14; fi78++ {
                if (42 <= -1) {
                    var v79 uint8 = 127
                    if true {
                        v80 := GenS13{F0: v12}
                        _ = v80
                    }
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            v12 = ""
                            v81 := GenS14{F0: ((100 < 127) || false)}
                            _ = v81
                        } else {
                            var v82 int16 = int16(fi78)
                            fmt.Println(v11, v30)
                            v82 = v82
                            _ = v82
                        }
                        var v83 int32 = -2147483648
                        _ = v83
                    } else {
                        if true {
                            var v84 uint64 = 287
                            _ = v84
                        }
                        var v85 int32 = 0
                        v86 := GenS10{F0: 0, F1: 42}
                        _ = v85
                        _ = v86
                    }
                    _ = v79
                }
                var v87 uint64 = 42
                if false {
                    var v88 int32 = 2147483647
                    _ = v88
                }
                switch (-(fi31)) {
                case -17:
                    fmt.Println(v12)
                case -2147483648:
                    fmt.Println(v87, v11)
                }
                _ = fi78
                _ = v87
            }
            _ = v77
        }
        var v89 int16 = 1
        _ = fi31
        _ = v89
    }
    v12 = v12
    v11 += v11
    var v90 int64 = (v30 | v30)
    v30 = v90
    v91 := GenS15{F0: v12, F1: v90}
    v30 = v90
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v90
    _ = v91
}
