package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 uint }
func genFn0(p0 int8, p1 int8) int16 {
    v4 := GenS1{F0: 65535, F1: -1, F2: (("\n\t" + "a") + "a")}
    fmt.Println(p0, p0)
    var v5 uint32 = 2147483648
    v6 := []int{42, -232}
    switch "hello" {
    case "a":
        if false {
            fmt.Println(p1)
            switch (("a" + "") + "a") {
            case "foo bar":
                v7 := GenS2{F0: ("" + ("hello" + ""))}
                _ = v7
            case "":
                switch ((("foo bar" + "\n\t") + ("foo bar" + ("hello" + "\n\t"))) + "\n\t") {
                case "a":
                    for fi8 := 0; fi8 < 1; fi8++ {
                        fmt.Println(p0, v5)
                        var v9 uint = 140
                        var v10 int8 = (p0 % ((int8(v5)) | 1))
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                case "foo bar":
                    var v11 string = (((("foo bar" + ("hello" + "\n\t")) + ("foo bar" + (("" + "hello") + "hello"))) + "") + ("a" + ""))
                    _ = v11
                case "hello":
                    var v12 uint16 = 105
                    _ = v12
                default:
                    var v13 int = 0
                    _ = v13
                }
            }
        } else {
            if (!(1 == 0)) {
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 int16 = 518
                    var v16 uint8 = 1
                    v17 := GenS3{F0: 547}
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                var v18 uint8 = 42
                for fi19 := 0; fi19 < 17; fi19++ {
                    if false {
                        switch fi19 {
                        case 127:
                            for fi20 := 0; fi20 < 11; fi20++ {
                                if false {
                                    v21 := GenS4{F0: -140}
                                    fmt.Println(v5, p1, p1)
                                    _ = v21
                                }
                                _ = fi20
                            }
                        case 493:
                            if (426 < 42) {
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    var v23 uint = 18446744073709551615
                                    if true {
                                        if len(v6) > 0 {
                                            v6[0] = 9223372036854775807
                                        }
                                        if (((v18 % ((v18) | 1)) <= 128) || false) {
                                            var v24 string = "\n\t"
                                            var v25 int16 = 0
                                            _ = v24
                                            _ = v25
                                        }
                                    } else {
                                        v26 := GenS3{F0: 67}
                                        p1 -= p0
                                        var v27 uint = 708
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = fi22
                                    _ = v23
                                }
                                v28 := GenS1{F0: v5, F1: p1, F2: ("hello" + "foo bar")}
                                _ = v28
                            }
                        case -2147483648:
                            if true {
                                var v29 int32 = -443
                                v30 := GenS5{F0: v29, F1: 4294967295}
                                _ = v29
                                _ = v30
                            }
                        }
                        p0 += (p0 >> ((143) & 63))
                        _ = len(v6)
                    }
                    v6 = append(v6, fi19)
                    _ = fi19
                }
                _ = v18
            } else {
                switch ("hello" + "\n\t") {
                case "\n\t":
                    if (2 == 22) {
                        var v31 bool = (!(255 >= 65535))
                        for fi32 := 0; fi32 < 16; fi32++ {
                            v33 := GenS2{F0: "\n\t"}
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v34 uint64 = 4294967295
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        _ = v31
                    }
                case "":
                    if len(v6) > 0 {
                        v6[0] = 0
                    }
                }
                if (-2147483648 == 42) {
                    fmt.Println(p0, v5)
                } else {
                    if false {
                        p1 = p0
                        p1 = 1
                    }
                    var v35 int = 9223372036854775807
                    v35 = v35
                    _ = v35
                }
                v6 = append(v6, 128)
            }
            for fi36 := 0; fi36 < 6; fi36++ {
                if false {
                    var v37 int64 = 81
                    for fi38 := 0; fi38 < 15; fi38++ {
                        v39 := GenS1{F0: (v5 & (^(v5))), F1: (p1 & (p1 / (((p1 * -1)) | 1))), F2: "hello"}
                        fmt.Println(v37)
                        v40 := GenS4{F0: 0}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    _ = v37
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi36
            }
            v41 := GenS2{F0: ("foo bar" + "a")}
            _ = v41
        }
    case "":
        var v42 uint16 = 42
        _ = v42
    default:
        var v43 uint32 = (v5 & (v5 >> ((255) & 63)))
        _ = v43
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -32768
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 65
    var v3 uint = 911
    var v44 uint32 = 514
    v3 -= v3
    v45 := GenS4{F0: 42}
    v46 := GenS5{F0: 2147483647, F1: -119}
    var v47 uint32 = (v44 << (((v3 - (^(v3)))) & 63))
    v1[(v47 + (^(v44)))] = (!(-229 >= -284))
    v1[v47] = false
    var v48 uint8 = v2
    switch "a" {
    case "a":
        switch -799 {
        case 32767:
            if (-1 != 42) {
                v44 += v47
                v49 := []int{0, 219, 127}
                _ = v49
            } else {
                for fi50 := 0; fi50 < 4; fi50++ {
                    for fi51 := 0; fi51 < 10; fi51++ {
                        if ("a" != "foo bar") {
                            fmt.Println(v47, fi50)
                        }
                        v52 := GenS6{F0: v3}
                        v44 -= uint32(v48)
                        _ = fi51
                        _ = v52
                    }
                    if (42 >= 1) {
                        if false {
                            var v53 uint64 = 2
                            _ = v53
                        } else {
                            var v54 uint = v3
                            _ = v54
                        }
                    } else {
                        var v55 int64 = 1
                        v56 := []int{fi50, fi50, 1, 128}
                        _ = v55
                        _ = v56
                    }
                    _ = fi50
                }
            }
        case 1:
            var v57 uint32 = v47
            _ = v57
        }
    case "":
        v58 := GenS3{F0: 1}
        _ = v58
    }
    if (2 > 1) {
        var v59 float64 = float64(0.1)
        _ = v59
    } else {
        fmt.Println(v3, v47, v48)
        if ((!true) && (65535 == (v3 % (((^(v3))) | 1)))) {
            for fi60 := 0; fi60 < 9; fi60++ {
                var v61 string = "a"
                _ = fi60
                _ = v61
            }
            switch 1 {
            case -181:
                v62 := []int{835, -1, 2147483647, -2147483648}
                _ = v62
            case -734:
                if ((347 >= 617) && false) {
                    v47 = v44
                }
            default:
                v2 -= 42
            }
            if false {
                fmt.Println(v48, v3, v48)
            } else {
                delete(v1, v47)
            }
        }
    }
    for fi63 := 0; fi63 < 12; fi63++ {
        v2 = uint8(v3)
        v64 := []int{(fi63 | (^(fi63))), (fi63 / ((9223372036854775807) | 1))}
        _ = fi63
        _ = v64
    }
    v44 -= v47
    for fi65 := 0; fi65 < 14; fi65++ {
        switch (^(fi65)) {
        case 128:
            if true {
                for fi66 := 0; fi66 < 18; fi66++ {
                    switch 780 {
                    case -2147483648:
                        switch fi66 {
                        case -470:
                            v67 := []int{fi65, (^(fi65)), 4294967295, (fi66 - (fi66 ^ (fi66 | fi65)))}
                            _ = v67
                        case 1:
                            v47 = (v47 / ((2147483647) | 1))
                        case 2147483647:
                            v3 = (v3 | (v3 % ((v3) | 1)))
                        }
                        fmt.Println(v48)
                    case -848:
                        delete(v1, (v44 << ((454) & 63)))
                    }
                    if (!false) {
                        if true {
                            v47 = (v47 - (v44 ^ (v44 & v47)))
                            fmt.Println(v48)
                            v44 = v47
                        }
                        if (1.5 == 1e10) {
                            var v68 int16 = genFn0(100, 127)
                            _ = v68
                        }
                    }
                    _ = fi66
                }
            }
        }
        _ = len(v1)
        _ = fi65
    }
    v3 = v3
    for fi69 := 0; fi69 < 8; fi69++ {
        v3 -= (v3 | v3)
        v70 := []int{(fi69 % ((417) | 1)), -313, (-(fi69)), 127}
        _ = fi69
        _ = v70
    }
    for fi71 := 0; fi71 < 17; fi71++ {
        v47 = 0
        _ = fi71
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
