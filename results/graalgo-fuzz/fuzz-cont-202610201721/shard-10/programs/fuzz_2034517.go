package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int64; F2 string }
type GenS4 struct { F0 uint16; F1 uint32; F2 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64; F1 int }
type GenS10 struct { F0 string; F1 uint }

func main() {
    v1 := []uint32{1, 961}
    var v2 uint32 = 0
    var v3 uint64 = 0
    var v4 int16 = 32767
    if true {
        var v5 uint8 = 36
        var v6 float32 = 0.0
        var v7 int32 = -2147483648
        _ = v5
        _ = v6
        _ = v7
    } else {
        _ = len(v1)
    }
    var v8 uint8 = 0
    switch -9223372036854775808 {
    case 853:
        v4 += v4
    case 523:
        v9 := GenS1{F0: 42, F1: (false || ((float64(1.5) / 3.14159) < 3.14159))}
        _ = v9
    case -2147483648:
        v10 := GenS1{F0: -1, F1: ((^(v2)) != 994)}
        _ = v10
    default:
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                if true {
                    if false {
                        v4 = (v4 & v4)
                        if ((384 < 2147483647) && true) {
                            var v13 string = "hello"
                            v14 := GenS1{F0: -109, F1: false}
                            if (!(!(int64(v3) != 1))) {
                                fmt.Println(fi11, v3, v4)
                            }
                            _ = len(v1)
                            _ = v13
                            _ = v14
                        } else {
                            v4 = v4
                            if false {
                                var v15 uint8 = v8
                                if (v15 <= 244) {
                                    fmt.Println(v3, v8)
                                }
                                var v16 string = "foo bar"
                                v16 = "foo bar"
                                var v17 uint16 = 1
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            var v18 uint64 = v3
                            if false {
                                v19 := GenS2{F0: "", F1: uint8(v3)}
                                var v20 string = ("foo bar" + ("a" + "a"))
                                if len(v1) > 0 {
                                    v1[0] = (v2 | 42)
                                }
                                if true {
                                    fmt.Println(v2)
                                    fmt.Println(fi12, v4, v4)
                                }
                                _ = v19
                                _ = v20
                            } else {
                                switch (fi11 >> ((0) & 63)) {
                                case 0:
                                    var v21 float64 = -2.25
                                    _ = v21
                                case -1:
                                    fmt.Println(v2, fi11)
                                default:
                                    v22 := GenS2{F0: "", F1: (v8 >> ((255) & 63))}
                                    _ = v22
                                }
                                v23 := GenS1{F0: -791, F1: (!false)}
                                _ = v23
                            }
                            _ = v18
                        }
                    } else {
                        if (323 > -832) {
                            v8 -= (v8 | 1)
                            v24 := GenS1{F0: 42, F1: true}
                            fmt.Println(fi12, v3)
                            var v25 uint64 = uint64(v2)
                            _ = v24
                            _ = v25
                        }
                        v2 += (^(v2))
                    }
                    v1 = append(v1, v2)
                    var v26 uint8 = (v8 >> ((uint(v3)) & 63))
                    _ = v26
                }
                _ = fi12
            }
            v4 = (v4 << ((905) & 63))
            var v27 uint32 = 2147483648
            _ = fi11
            _ = v27
        }
    }
    var v28 bool = true
    v4 += (v4 ^ v4)
    var v29 int64 = 9223372036854775807
    v30 := GenS2{F0: ((("foo bar" + "a") + ("hello" + "foo bar")) + ""), F1: v8}
    v8 -= (^(v8))
    v4 = -1
    v8 -= 3
    for fi31 := 0; fi31 < 9; fi31++ {
        if (v29 == v29) {
            for fi32 := 0; fi32 < 14; fi32++ {
                v8 += uint8(v3)
                var v33 uint16 = 945
                if (!(4294967295 <= fi31)) {
                    fmt.Println(v3, v28, v3)
                } else {
                    if false {
                        if (v3 >= (v3 >> ((460) & 63))) {
                            var v34 float32 = float32(((((0.1 + ((1e10 - 0.1) * -2.25)) - 3.14159) / -2.25) / (0.1 - ((float64(0.1) / ((3.14159 - 1e10) * float64(1.5))) / (1.5 + -2.25)))))
                            if v28 {
                                v35 := GenS3{F0: 255, F1: v29, F2: ("\n\t" + "a")}
                                _ = v35
                            }
                            if false {
                                var v36 uint64 = uint64(v8)
                                v37 := GenS2{F0: ("hello" + ""), F1: uint8(v2)}
                                if len(v1) > 0 {
                                    v1[0] = (^(v2))
                                }
                                _ = v36
                                _ = v37
                            }
                            v4 = 1000
                            fmt.Println(v29)
                            v38 := GenS3{F0: (^(v8)), F1: v29, F2: "a"}
                            if (!v28) {
                                fmt.Println(v29)
                                if true {
                                    fmt.Println(v4, v34, v2)
                                    fmt.Println(v4)
                                }
                                switch fi32 {
                                case 1:
                                    v39 := GenS2{F0: "a", F1: (v8 * v8)}
                                    _ = v39
                                case 689:
                                    fmt.Println(v33)
                                }
                            } else {
                                fmt.Println(fi31, fi32, v2)
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = v2
                                    }
                                    switch (^(fi32)) {
                                    case -925:
                                        if (uint8(v2) > v8) {
                                            var v40 int64 = -1
                                            v33 -= 0
                                            _ = v40
                                        }
                                    default:
                                        fmt.Println(v29, v29, fi31)
                                        if len(v1) > 0 {
                                            v1[0] = v2
                                        }
                                        var v41 uint8 = v8
                                        _ = v41
                                    }
                                    switch (fi31 & (fi32 * fi31)) {
                                    case -9223372036854775808:
                                        v42 := GenS4{F0: 255, F1: (v2 + v2), F2: ""}
                                        v43 := GenS5{F0: 0}
                                        _ = v42
                                        _ = v43
                                    default:
                                        var v44 uint64 = (v3 * uint64(v33))
                                        _ = v44
                                    }
                                } else {
                                    v8 = 175
                                }
                            }
                            _ = v34
                            _ = v38
                        } else {
                            v2 -= v2
                        }
                        var v45 int8 = -1
                        var v46 uint = 18446744073709551615
                        _ = v45
                        _ = v46
                    } else {
                        v1 = append(v1, 995)
                        var v47 int64 = 42
                        var v48 int32 = -1
                        _ = v47
                        _ = v48
                    }
                }
                if v28 {
                    if ((^(v33)) > 2) {
                        fmt.Println(v3, fi31, v28)
                    }
                    var v49 uint8 = 0
                    _ = v49
                } else {
                    if (!(v8 >= uint8(v33))) {
                        switch ("a" + "") {
                        case "hello":
                            var v50 int16 = v4
                            _ = v50
                        default:
                            if (!false) {
                                if ((1 != 318) && ((false || (false && v28)) && false)) {
                                    var v51 uint64 = v3
                                    var v52 uint32 = (v2 & (v2 + v2))
                                    _ = v51
                                    _ = v52
                                }
                                if (!true) {
                                    if len(v1) > 0 {
                                        v1[0] = v2
                                    }
                                    _ = len(v1)
                                    if ((!v28) && (v29 == v29)) {
                                        if false {
                                            if v28 {
                                                var v53 int32 = 802
                                                v29 = 9223372036854775807
                                                v1 = append(v1, v2)
                                                _ = v53
                                            }
                                        }
                                    }
                                } else {
                                    v4 -= (v4 / (((v4 % ((v4) | 1))) | 1))
                                    if ((v33 & (v33 / ((v33) | 1))) > (v33 << ((2) & 63))) {
                                        fmt.Println(v8)
                                    } else {
                                        if v28 {
                                            v54 := GenS2{F0: "foo bar", F1: 42}
                                            _ = v54
                                        } else {
                                            v4 = (v4 & (v4 << ((2) & 63)))
                                            v3 = v3
                                            switch (fi31 - -762) {
                                            case 32767:
                                                var v55 int32 = 2147483647
                                                _ = v55
                                            case 4294967295:
                                                var v56 uint = 65535
                                                _ = v56
                                            default:
                                                if len(v1) > 0 {
                                                    v1[0] = v2
                                                }
                                            }
                                            switch (fi32 >> ((302) & 63)) {
                                            case 127:
                                                if len(v1) > 0 {
                                                    v1[0] = v2
                                                }
                                            case -1:
                                                var v57 int8 = -1
                                                _ = v57
                                            case 4294967295:
                                                var v58 int32 = 161
                                                if v28 {
                                                    var v59 int8 = 100
                                                    fmt.Println(fi31, v8, v8)
                                                    var v60 int16 = (v4 | v4)
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    v61 := GenS6{F0: v2, F1: 100}
                                                    fmt.Println(v2, v8)
                                                    _ = v61
                                                }
                                                v58 = v58
                                                fmt.Println(v3, v8, v8)
                                                v58 -= 2147483647
                                                if (!(v28 || (v3 == (v3 - v3)))) {
                                                    fmt.Println(v3, v8)
                                                    v62 := GenS7{F0: (-(fi32))}
                                                    var v63 float64 = ((0.1 + 1e10) / 0.1)
                                                    _ = v62
                                                    _ = v63
                                                } else {
                                                    v64 := GenS2{F0: ("a" + "\n\t"), F1: v8}
                                                    if v28 {
                                                        if v28 {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v65 := GenS6{F0: v2, F1: -128}
                                                            v66 := GenS5{F0: 42}
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        if true {
                                                            v29 += int64(v33)
                                                            var v67 uint = 2
                                                            _ = v67
                                                        }
                                                    }
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if (v4 != 446) {
                                                            if (v28 || v28) {
                                                                v68 := GenS5{F0: -125}
                                                                _ = v68
                                                            } else {
                                                                v69 := GenS5{F0: -1}
                                                                switch ("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))) {
                                                                case "a":
                                                                    v3 += (v3 << ((4294967295) & 63))
                                                                case "":
                                                                    v33 = v33
                                                                case "\n\t":
                                                                    var v70 int32 = (v58 >> ((0) & 63))
                                                                    _ = v70
                                                                default:
                                                                    if (true || ((^(v58)) > 1)) {
                                                                        v71 := GenS8{F0: 18446744073709551615}
                                                                        _ = v71
                                                                    } else {
                                                                        if v28 {
                                                                            v28 = ((v33 | v33) > (v33 - v33))
                                                                            fmt.Println(v29, v2, fi32)
                                                                            var v72 uint64 = 42
                                                                            var v73 uint32 = (v2 << ((583) & 63))
                                                                            fmt.Println(v2, fi31, v8)
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                        v58 = -4
                                                                        v8 = (v8 & uint8(v33))
                                                                    }
                                                                }
                                                                v74 := GenS9{F0: (v29 >> ((uint(v8)) & 63)), F1: (fi31 - int(v3))}
                                                                _ = v69
                                                                _ = v74
                                                            }
                                                        } else {
                                                            v75 := GenS2{F0: "\n\t", F1: 1}
                                                            _ = v75
                                                        }
                                                    }
                                                    var v76 uint = 42
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v64
                                                    _ = v76
                                                }
                                                fmt.Println(v58)
                                                _ = v58
                                            default:
                                                v77 := GenS1{F0: -761, F1: (false && (v2 == (v2 * (v2 + v2))))}
                                                _ = v77
                                            }
                                        }
                                        v4 = v4
                                    }
                                    v8 = (v8 & v8)
                                }
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = v2
                        }
                    } else {
                        var v78 int8 = 1
                        _ = v78
                    }
                    var v79 uint64 = 4294967295
                    _ = v79
                }
                fmt.Println(v33)
                _ = fi32
                _ = v33
            }
            v1 = append(v1, (v2 * 4294967295))
        } else {
            for fi80 := 0; fi80 < 3; fi80++ {
                v81 := GenS2{F0: ("" + ("" + "\n\t")), F1: 231}
                if v28 {
                    if v28 {
                        var v82 int16 = (v4 >> ((4294967295) & 63))
                        fmt.Println(v29)
                        _ = v82
                    }
                } else {
                    v8 = 42
                }
                _ = fi80
                _ = v81
            }
            switch (fi31 * fi31) {
            case 72:
                for fi83 := 0; fi83 < 10; fi83++ {
                    var v84 int8 = 100
                    v85 := GenS10{F0: "hello", F1: 2}
                    _ = fi83
                    _ = v84
                    _ = v85
                }
            case -945:
                v28 = false
            }
        }
        var v86 int64 = v29
        var v87 int32 = 2147483647
        _ = fi31
        _ = v86
        _ = v87
    }
    v1 = append(v1, uint32(v8))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v8
}
