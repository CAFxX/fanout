package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool; F1 int64; F2 int8 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint }
type GenS4 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS7 struct { F0 uint8; F1 int; F2 int16 }
func genFn0(p0 uint16, p1 int8, p2 int8) int32 {
    p2 += (p1 + int8(p0))
    v5 := GenS1{F0: 975, F1: false, F2: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 213
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    for fi6 := 0; fi6 < 15; fi6++ {
        switch "\n\t" {
        case "a":
            v7 := []int{}
            _ = v7
        case "foo bar":
            switch (-(fi6)) {
            case -1:
                if ((1 == 1) && (false || (0 >= 42))) {
                    v2 = (v2 >> ((355) & 63))
                    v2 = -42
                } else {
                    v2 = (-(v2))
                }
            }
        case "":
            v8 := GenS2{F0: true, F1: 42, F2: 110}
            _ = v8
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS2{F0: (9223372036854775807 <= -9223372036854775808), F1: -9223372036854775808, F2: 0}
        var v11 string = ("a" + "foo bar")
        if (!((!(42 == -32768)) || (((0 > 526) && true) || false))) {
            v2 += genFn0(251, 127, 100)
            if true {
                v2 = v2
                if true {
                    v12 := []int{}
                    if false {
                        v2 = genFn0(684, 100, 27)
                        if (1 <= 255) {
                            v13 := GenS3{F0: -128, F1: 127, F2: 4294967295}
                            if len(v12) > 0 {
                                v12[0] = fi9
                            }
                            _ = v13
                        }
                        var v14 uint32 = 4294967295
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 15; fi15++ {
                        if false {
                            if len(v12) > 0 {
                                v12[0] = fi15
                            }
                            var v16 uint = 4294967295
                            _ = v16
                        }
                        var v17 int32 = (-(v2))
                        _ = fi15
                        _ = v17
                    }
                    _ = v12
                }
                v18 := GenS1{F0: 726, F1: true, F2: 118}
                _ = v18
            } else {
                v11 = v11
                var v19 uint8 = 120
                _ = v19
            }
        } else {
            var v20 int32 = -2147483648
            _ = v20
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v21 uint32 = 882
    switch 127 {
    case 0:
        v22 := GenS4{F0: 1, F1: 42, F2: 255}
        _ = v22
    case 32767:
        for fi23 := 0; fi23 < 9; fi23++ {
            v24 := GenS5{F0: 1, F1: 1, F2: 493}
            _ = fi23
            _ = v24
        }
    default:
        var v25 bool = (96 >= 127)
        _ = v25
    }
    if ((v21 << ((1) & 63)) < 850) {
        switch 32768 {
        case 32767:
            for fi26 := 0; fi26 < 9; fi26++ {
                for fi27 := 0; fi27 < 11; fi27++ {
                    if (((88 > 100) && (false && ((v2 | (v2 - (v2 + (v2 | (^(v2)))))) >= (^(v2))))) && false) {
                        var v28 int32 = v2
                        if (1 > 1) {
                            v29 := GenS2{F0: ((1.5 - ((1.5 * 3.14159) + -2.25)) < 1.5), F1: 915, F2: -1}
                            var v30 string = "a"
                            _ = v29
                            _ = v30
                        } else {
                            var v31 uint32 = v21
                            switch (fi27 >> ((2) & 63)) {
                            case 9223372036854775807:
                                var v32 string = "hello"
                                _ = v32
                            case 0:
                                v33 := []int{int(v31), -2147483648}
                                _ = v33
                            }
                            v34 := GenS2{F0: ((int(v21) == fi26) || (1 >= uint16(v21))), F1: -471, F2: 127}
                            _ = v31
                            _ = v34
                        }
                        v2 -= genFn0(2, 0, 114)
                        v21 += 361
                        fmt.Println(fi27, v28, fi26)
                        if false {
                            fmt.Println(v21, fi27, fi26)
                            if false {
                                if true {
                                    var v35 uint64 = 114
                                    _ = v35
                                }
                                v36 := GenS2{F0: (127 > -81), F1: 0, F2: 53}
                                _ = v36
                            }
                            v2 = v2
                            v37 := GenS4{F0: 267, F1: 2, F2: 32767}
                            v4[(^(v21))] = ""
                            _ = v37
                        } else {
                            v38 := GenS3{F0: 112, F1: 42, F2: 3}
                            _ = v38
                        }
                        v39 := []int{(fi27 << ((234) & 63)), fi26, (fi27 % ((fi27) | 1)), fi27}
                        v21 = v21
                        _ = v28
                        _ = v39
                    }
                    v21 = v21
                    v21 = (v21 - v21)
                    _ = fi27
                }
                v40 := GenS6{F0: 255, F1: 42, F2: 4294967295}
                _ = fi26
                _ = v40
            }
        case 0:
            if (1 > v2) {
                for fi41 := 0; fi41 < 2; fi41++ {
                    var v42 uint = 255
                    _ = fi41
                    _ = v42
                }
                switch 32767 {
                case 32768:
                    var v43 float64 = 1.5
                    _ = v43
                case 551:
                    v44 := GenS1{F0: (^(v21)), F1: false, F2: 1}
                    _ = v44
                case 902:
                    var v45 uint16 = 255
                    _ = v45
                }
                v21 += v21
            }
        case 671:
            switch ("foo bar" + "hello") {
            case "a":
                var v46 float64 = (3.14159 + 0.0)
                _ = v46
            case "\n\t":
                fmt.Println(v2, v2)
            }
        default:
            var v47 uint16 = 65535
            _ = v47
        }
    }
    fmt.Println(v21)
    var v48 string = ((("foo bar" + ("hello" + "a")) + "") + "a")
    var v49 float64 = float64(1e10)
    if ((128 < 152) && (7 == v21)) {
        var v50 uint64 = 1
        _ = len(v1)
        _ = v50
    }
    v49 += v49
    var v51 int8 = -7
    switch int(v2) {
    case -971:
        var v52 int64 = 1
        _ = v52
    case 4294967295:
        v21 = (^(v21))
    case 213:
        var v53 uint16 = 32767
        _ = v53
    }
    v54 := GenS4{F0: 42, F1: 0, F2: 32767}
    if true {
        _ = len(v1)
        v55 := GenS3{F0: 57, F1: 0, F2: 49}
        _ = v55
    }
    for fi56 := 0; fi56 < 3; fi56++ {
        v57 := GenS7{F0: uint8(v21), F1: fi56, F2: -570}
        if ((v49 + 3.14159) < (1.5 - v49)) {
            v58 := GenS6{F0: 316, F1: 255, F2: 4294967295}
            _ = v58
        } else {
            for fi59 := 0; fi59 < 18; fi59++ {
                v51 = 59
                _ = fi59
            }
        }
        var v60 int8 = v51
        _ = fi56
        _ = v57
        _ = v60
    }
    v61 := GenS2{F0: false, F1: 0, F2: (^(v51))}
    switch -9223372036854775808 {
    case 572:
        if ("" >= v48) {
            if (!(9223372036854775807 < -902)) {
                v62 := GenS6{F0: 0, F1: 255, F2: 42}
                var v63 int64 = -1
                _ = v62
                _ = v63
            } else {
                for fi64 := 0; fi64 < 3; fi64++ {
                    v21 -= (v21 * (^(v21)))
                    v21 += v21
                    var v65 int8 = (v51 & 127)
                    _ = fi64
                    _ = v65
                }
            }
        } else {
            switch -277 {
            case 32768:
                fmt.Println(v49, v48, v48)
            case 2147483647:
                var v66 string = v48
                _ = v66
            case 4294967295:
                v49 = ((v49 * float64((0.0 - -2.25))) / -2.25)
            default:
                var v67 uint64 = 4294967295
                _ = v67
            }
            _ = len(v4)
            v2 = (v2 * 2147483647)
        }
    case -2147483648:
        v48 = v48
    case 127:
        var v68 uint32 = 561
        _ = v68
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v51
    _ = v54
    _ = v61
}
