package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 int8; F2 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8; F1 bool }
func genFn0(p0 string, p1 bool) int64 {
    var v4 uint16 = 1
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 bool = false
        if (!false) {
            p1 = true
            var v7 float64 = 0.0
            for fi8 := 0; fi8 < 20; fi8++ {
                fmt.Println(v4, fi5, v6)
                _ = fi8
            }
            _ = v7
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            if (true || ((1.5 * 0.1) < -2.25)) {
                if (583 > -9223372036854775808) {
                    if (1.5 != 3.14159) {
                        fmt.Println(p0)
                        v10 := []int{fi5, fi9}
                        if (float32(-2.25) <= (-2.25 + 0.0)) {
                            p0 = "hello"
                            fmt.Println(p0, p0, v6)
                            v6 = (p1 && (42 >= 0))
                            v10 = append(v10, fi9)
                        } else {
                            var v11 uint = 18446744073709551615
                            v4 -= v4
                            _ = v11
                        }
                        _ = v10
                    }
                    if (!(2 != 127)) {
                        switch "foo bar" {
                        case "a":
                            v12 := GenS1{F0: 42}
                            _ = v12
                        case "hello":
                            v13 := GenS2{F0: p0, F1: -626}
                            _ = v13
                        }
                    }
                    if p1 {
                        v14 := GenS2{F0: ((p0 + ("hello" + p0)) + ("hello" + p0)), F1: 681}
                        v15 := []int{fi9}
                        v16 := GenS1{F0: 2}
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                } else {
                    if p1 {
                        var v17 bool = (657 > v4)
                        var v18 int32 = -1
                        _ = v17
                        _ = v18
                    } else {
                        v19 := []int{127, (fi5 << ((2) & 63)), -9223372036854775808}
                        if (!true) {
                            var v20 bool = (!(((!(!v6)) || (!(((810 == 327) && v6) && p1))) || false))
                            var v21 uint16 = v4
                            if (!v6) {
                                v22 := GenS1{F0: 4294967295}
                                v23 := GenS3{F0: -94}
                                switch p0 {
                                case "hello":
                                    v24 := GenS1{F0: 18446744073709551615}
                                    _ = v24
                                case "":
                                    switch 128 {
                                    case 32768:
                                        if v6 {
                                            var v25 int8 = 1
                                            _ = v25
                                        } else {
                                            var v26 float32 = 0.1
                                            _ = v26
                                        }
                                    default:
                                        var v27 uint = 1
                                        _ = v27
                                    }
                                case "foo bar":
                                    if v20 {
                                        if len(v19) > 0 {
                                            v19[0] = 1
                                        }
                                    } else {
                                        var v28 int = (fi5 + 9223372036854775807)
                                        if len(v19) > 0 {
                                            v19[0] = v28
                                        }
                                        _ = v28
                                    }
                                }
                                _ = v22
                                _ = v23
                            }
                            var v29 uint32 = 1
                            var v30 int16 = 1000
                            if (17 == 0) {
                                var v31 bool = (970 != v29)
                                var v32 uint16 = 255
                                v29 = 2147483647
                                var v33 int64 = 0
                                v33 = v33
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            v4 = (v4 >> ((4294967295) & 63))
                            _ = v20
                            _ = v21
                            _ = v29
                            _ = v30
                        } else {
                            var v34 uint32 = 4294967295
                            var v35 string = p0
                            var v36 uint8 = 128
                            fmt.Println(v6)
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        _ = v19
                    }
                    var v37 float32 = (0.0 + -2.25)
                    var v38 uint32 = 65535
                    _ = v37
                    _ = v38
                }
                v39 := GenS1{F0: 18446744073709551615}
                fmt.Println(p1)
                v40 := []int{}
                _ = v39
                _ = v40
            } else {
                v41 := GenS2{F0: p0, F1: 485}
                var v42 uint32 = 0
                v4 += (v4 | (v4 % ((593) | 1)))
                _ = v41
                _ = v42
            }
            v43 := []int{32767, (fi5 & (fi9 | 42)), (fi9 & -662)}
            _ = fi9
            _ = v43
        }
        _ = fi5
        _ = v6
    }
    var v44 float32 = ((((1.5 + 0.1) + -2.25) * 1.5) + (3.14159 / 0.1))
    _ = p0
    _ = p1
    _ = v4
    _ = v44
    return -801
}
func genFn1() int64 {
    v45 := GenS2{F0: "foo bar", F1: 32767}
    for fi46 := 0; fi46 < 12; fi46++ {
        var v47 uint8 = 2
        fmt.Println(v47)
        _ = fi46
        _ = v47
    }
    var v48 int8 = 81
    v48 = v48
    _ = v45
    _ = v48
    return 42
}

func main() {
    var v1 string = ("hello" + "hello")
    var v2 uint16 = 1
    var v3 string = v1
    v2 = 2
    if (516 >= -1) {
        for fi49 := 0; fi49 < 13; fi49++ {
            v1 = v3
            var v50 bool = (true && (0 > -1))
            _ = fi49
            _ = v50
        }
        v51 := []int{-721, 4294967295, 0, 112}
        if false {
            if true {
                var v52 float32 = 3.14159
                v2 = (^(v2))
                fmt.Println(v2, v1)
                _ = v52
            } else {
                if (42 > -171) {
                    var v53 string = (v3 + ((("foo bar" + ((("" + "") + "hello") + (("a" + v1) + (v1 + (v3 + (v1 + (((v1 + v1) + (v1 + v1)) + v1))))))) + (((v3 + (v1 + "\n\t")) + (((v3 + "\n\t") + (v3 + "")) + (v3 + ""))) + ("a" + ("" + ("" + v3))))) + (("foo bar" + v3) + v1)))
                    _ = v53
                } else {
                    if len(v51) >= 2 {
                        v51 = v51[1:len(v51)]
                    }
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v1, v1, v2)
                    }
                }
                if (189 > 108) {
                    v2 -= v2
                    fmt.Println(v2, v2)
                    if (763 != 787) {
                        for fi54 := 0; fi54 < 13; fi54++ {
                            switch ((v3 + "") + (v1 + ("" + (v1 + v1)))) {
                            case "a":
                                var v55 uint8 = 165
                                _ = v55
                            case "foo bar":
                                for fi56 := 0; fi56 < 11; fi56++ {
                                    var v57 uint32 = 0
                                    v58 := GenS2{F0: ("\n\t" + (v3 + "hello")), F1: 1}
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                }
                            }
                            if len(v51) >= 2 {
                                v51 = v51[1:len(v51)]
                            }
                            var v59 int16 = -32768
                            _ = fi54
                            _ = v59
                        }
                        if (128 != 127) {
                            if len(v51) > 0 {
                                v51[0] = 229
                            }
                            if true {
                                var v60 float32 = float32(0.1)
                                var v61 uint = 65535
                                _ = v60
                                _ = v61
                            } else {
                                fmt.Println(v3, v2, v2)
                                v51 = append(v51, -2147483648)
                            }
                            switch v3 {
                            case "":
                                var v62 bool = (!true)
                                _ = v62
                            case "hello":
                                v63 := GenS3{F0: 127}
                                _ = v63
                            case "foo bar":
                                v64 := GenS2{F0: "a", F1: 0}
                                _ = v64
                            default:
                                v65 := GenS2{F0: ((v1 + (v1 + (v1 + v3))) + v1), F1: 42}
                                _ = v65
                            }
                        }
                    } else {
                        _ = len(v51)
                        v3 = v3
                    }
                } else {
                    v3 = v1
                    fmt.Println(v1)
                    v66 := GenS2{F0: (v1 + (v1 + v1)), F1: 702}
                    _ = v66
                }
                var v67 uint8 = 234
                _ = v67
            }
        }
        _ = v51
    } else {
        v68 := GenS4{F0: 42, F1: -23, F2: "hello"}
        var v69 string = v3
        v70 := GenS5{F0: 1}
        _ = v68
        _ = v69
        _ = v70
    }
    v1 = ((("" + (v3 + (v3 + ((v3 + "\n\t") + (v1 + (v3 + v1)))))) + "hello") + v3)
    v71 := []int{24, 642, 257}
    var v72 bool = false
    for fi73 := 0; fi73 < 17; fi73++ {
        fmt.Println(v3, fi73, v1)
        v74 := GenS1{F0: 255}
        _ = fi73
        _ = v74
    }
    var v75 uint = 26
    var v76 uint8 = 135
    for fi77 := 0; fi77 < 18; fi77++ {
        var v78 string = ("foo bar" + (v1 + (v1 + (v1 + v3))))
        _ = fi77
        _ = v78
    }
    switch 527 {
    case 2147483647:
        if len(v71) >= 2 {
            v71 = v71[1:len(v71)]
        }
    case 4294967295:
        if v72 {
            if len(v71) > 0 {
                v71[0] = 710
            }
        }
    case 680:
        var v79 uint8 = 222
        _ = v79
    }
    v80 := GenS4{F0: 42, F1: 42, F2: v1}
    v81 := GenS6{F0: (v76 << ((843) & 63)), F1: (-546 >= 9223372036854775807)}
    if false {
        var v82 uint = 255
        if (true || v72) {
            fmt.Println(v75)
            var v83 uint64 = 4294967295
            _ = v83
        } else {
            v76 -= (v76 | 127)
        }
        var v84 uint32 = 0
        _ = v82
        _ = v84
    } else {
        v76 = 230
    }
    fmt.Println(v1, v3, v3)
    for fi85 := 0; fi85 < 5; fi85++ {
        v72 = (-450 > -909)
        if (!false) {
            v72 = (!(!((false && (false || (972 >= 2147483647))) || false)))
        }
        _ = fi85
    }
    fmt.Println(v76, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v71:", len(v71), v71)
    fmt.Println("v72:", v72)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v71
    _ = v72
    _ = v75
    _ = v76
    _ = v80
    _ = v81
}
