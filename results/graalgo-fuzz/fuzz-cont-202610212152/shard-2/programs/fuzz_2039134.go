package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int64; F1 int16; F2 string }
type GenS3 struct { F0 uint16; F1 uint; F2 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 uint8; F2 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int32; F1 bool; F2 uint8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8; F1 string }
func genFn0() int {
    if false {
        var v4 string = "foo bar"
        if false {
            v4 = v4
        }
        _ = v4
    } else {
        var v5 string = "hello"
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int32 = 2147483647
            v5 = v5
            var v8 uint16 = 0
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    }
    var v9 uint8 = 42
    if (1 >= 65535) {
        switch ("hello" + ("foo bar" + ((("\n\t" + "\n\t") + "foo bar") + ("hello" + "a")))) {
        case "a":
            fmt.Println(v9)
        case "foo bar":
            v9 -= 128
        case "\n\t":
            v9 += v9
        default:
            v9 = 125
        }
        v10 := []int{1}
        _ = v10
    }
    if (42 > -128) {
        var v11 bool = (1 >= 1)
        for fi12 := 0; fi12 < 1; fi12++ {
            v11 = (-9223372036854775808 > 0)
            if (v11 || v11) {
                for fi13 := 0; fi13 < 18; fi13++ {
                    v14 := GenS1{F0: 0, F1: 32767}
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v9)
            } else {
                v9 = (v9 << ((42) & 63))
                v15 := GenS1{F0: -128, F1: 4294967295}
                _ = v15
            }
            _ = fi12
        }
        for fi16 := 0; fi16 < 6; fi16++ {
            if ((-1 <= -177) || (500 != 345)) {
                switch 2147483647 {
                case 446:
                    var v17 uint32 = 128
                    _ = v17
                }
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := []int{}
                    v20 := GenS2{F0: 690, F1: 1, F2: "foo bar"}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v21 := []int{fi16, (fi16 - fi16), (fi16 ^ (^(fi16)))}
                _ = v21
            }
            for fi22 := 0; fi22 < 5; fi22++ {
                if (298 >= 601) {
                    if v11 {
                        v11 = (-128 >= 26)
                        switch "foo bar" {
                        case "a":
                            if (v11 || v11) {
                                v23 := GenS1{F0: 0, F1: 0}
                                fmt.Println(fi16)
                                _ = v23
                            }
                            v9 -= (v9 * v9)
                        case "foo bar":
                            v24 := []int{fi22, fi16, (fi16 - (fi22 << ((1) & 63))), fi16}
                            _ = v24
                        default:
                            v25 := []int{1, fi16, (^(fi16))}
                            _ = v25
                        }
                        if true {
                            v9 -= v9
                        } else {
                            v11 = v11
                            v11 = v11
                            v26 := []int{(fi22 ^ (-(fi22))), -1, fi22, fi16}
                            _ = v26
                        }
                    }
                    var v27 uint = 255
                    _ = v27
                }
                if false {
                    var v28 int64 = 1
                    var v29 float32 = -2.25
                    v30 := GenS3{F0: 65535, F1: 1, F2: -1}
                    _ = v28
                    _ = v29
                    _ = v30
                }
                if v11 {
                    v11 = false
                    var v31 uint64 = 42
                    _ = v31
                } else {
                    v32 := GenS4{F0: 32767}
                    _ = v32
                }
                _ = fi22
            }
            _ = fi16
        }
        _ = v11
    } else {
        fmt.Println(v9)
    }
    v33 := GenS1{F0: 42, F1: -208}
    _ = v33
    _ = v9
    return 638
}
func genFn1(p0 int, p1 int64, p2 int32) int8 {
    var v34 int64 = (p1 & (p1 + -434))
    var v35 int = 848
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    _ = v35
    return 37
}

func main() {
    v1 := []uint32{132, 42, 1}
    var v2 bool = false
    v3 := []string{"foo bar", "a", "foo bar"}
    for fi36 := 0; fi36 < 17; fi36++ {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = fi36
    }
    if v2 {
        var v37 string = "hello"
        var v38 bool = (0 == 196)
        _ = v37
        _ = v38
    } else {
        v39 := GenS5{F0: 0}
        v3 = append(v3, ("foo bar" + "hello"))
        _ = v39
    }
    v2 = (0.0 < 3.14159)
    v2 = (!(v2 || (v2 || v2)))
    var v40 bool = (false && true)
    if (2 <= 1) {
        switch 179 {
        case 1:
            switch 32768 {
            case 351:
                v2 = false
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        case 9223372036854775807:
            switch ("\n\t" + ("foo bar" + ("foo bar" + "a"))) {
            case "foo bar":
                for fi41 := 0; fi41 < 4; fi41++ {
                    var v42 uint16 = 111
                    var v43 uint16 = (^(v42))
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            case "a":
                switch "a" {
                case "hello":
                    switch genFn0() {
                    case 128:
                        if (-483 <= -2147483648) {
                            var v44 string = "hello"
                            fmt.Println(v2)
                            _ = v44
                        }
                    case 867:
                        if true {
                            switch 42 {
                            case 9223372036854775807:
                                if len(v3) > 0 {
                                    v3[0] = "foo bar"
                                }
                            case 541:
                                fmt.Println(v2, v2)
                            case -347:
                                for fi45 := 0; fi45 < 10; fi45++ {
                                    var v46 int8 = 71
                                    for fi47 := 0; fi47 < 7; fi47++ {
                                        var v48 uint32 = 4294967295
                                        v46 += 42
                                        _ = fi47
                                        _ = v48
                                    }
                                    switch (^(fi45)) {
                                    case -2147483648:
                                        if len(v1) > 0 {
                                            v1[0] = 966
                                        }
                                    case 127:
                                        if (v40 && false) {
                                            var v49 string = ""
                                            fmt.Println(fi45, v40, v49)
                                            _ = v49
                                        }
                                    }
                                    _ = fi45
                                    _ = v46
                                }
                            }
                            var v50 uint64 = 18446744073709551615
                            _ = v50
                        }
                    }
                default:
                    var v51 int32 = 2147483647
                    _ = v51
                }
            }
        case -9223372036854775808:
            switch 695 {
            case 974:
                for fi52 := 0; fi52 < 16; fi52++ {
                    v53 := GenS4{F0: 0}
                    if ("\n\t" < "\n\t") {
                        v40 = (1.5 != ((float64(((1e10 / 3.14159) + ((1.5 / float32((1e10 + 1.5))) + (float32(3.14159) + 1.5)))) * -2.25) + 0.0))
                        v54 := GenS3{F0: 296, F1: 2, F2: 42}
                        _ = v54
                    }
                    v40 = (-1 != 2147483647)
                    _ = fi52
                    _ = v53
                }
            case -9223372036854775808:
                var v55 string = ("\n\t" + "hello")
                _ = v55
            case 42:
                for fi56 := 0; fi56 < 17; fi56++ {
                    if v2 {
                        v57 := GenS6{F0: 2, F1: 2, F2: (255 < 1)}
                        if true {
                            var v58 int = 0
                            var v59 bool = v40
                            for fi60 := 0; fi60 < 13; fi60++ {
                                v61 := GenS5{F0: 56}
                                _ = fi60
                                _ = v61
                            }
                            _ = v58
                            _ = v59
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                _ = len(v3)
                            case "a":
                                if v2 {
                                    switch genFn0() {
                                    case 42:
                                        fmt.Println(v2, fi56, fi56)
                                    case 32767:
                                        var v62 int8 = -62
                                        _ = v62
                                    case 2147483647:
                                        if (!((2 == 255) && (!v40))) {
                                            var v63 string = ""
                                            if (false && (((1e10 - 1.5) + (float64(0.0) * -2.25)) >= ((1e10 + (1e10 * 1.5)) + 1e10))) {
                                                fmt.Println(v2)
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                for fi64 := 0; fi64 < 7; fi64++ {
                                                    if len(v1) > 0 {
                                                        v1[0] = 1
                                                    }
                                                    _ = fi64
                                                }
                                            }
                                            _ = v63
                                        } else {
                                            v1 = append(v1, 2147483647)
                                            var v65 uint8 = 178
                                            _ = v65
                                        }
                                    }
                                    v40 = (163 >= 2)
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v40 = ((65535 != 1) || false)
                                    _ = len(v1)
                                }
                            }
                        }
                        _ = v57
                    }
                    _ = fi56
                }
            default:
                switch -1 {
                case 360:
                    switch -954 {
                    case 9223372036854775807:
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                    }
                case 127:
                    v2 = v40
                case 32768:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    if (("hello" + "hello") <= (("foo bar" + "\n\t") + "a")) {
                        _ = len(v1)
                        v40 = false
                        switch 9223372036854775807 {
                        case 9223372036854775807:
                            fmt.Println(v2, v40)
                        default:
                            fmt.Println(v2)
                        }
                    }
                }
            }
        }
        v40 = (v2 && (65535 == 2))
        fmt.Println(v40)
    }
    for fi66 := 0; fi66 < 10; fi66++ {
        v67 := GenS4{F0: -32768}
        v68 := GenS7{F0: 505}
        if v2 {
            if v2 {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch fi66 {
                case -1:
                    var v69 int = (fi66 ^ fi66)
                    _ = v69
                case 0:
                    v2 = (2 > 32767)
                case -2147483648:
                    v3 = append(v3, "a")
                }
                if v40 {
                    v40 = v2
                } else {
                    v40 = v40
                    if false {
                        v2 = (v2 && (!("a" > "")))
                    }
                }
            }
            for fi70 := 0; fi70 < 11; fi70++ {
                var v71 float64 = float64(0.1)
                var v72 uint = 438
                switch "" {
                case "\n\t":
                    _ = len(v1)
                case "a":
                    v2 = (v2 && ((v2 && (32767 <= -176)) && v2))
                default:
                    v2 = (!("foo bar" > (("foo bar" + ("hello" + "hello")) + "\n\t")))
                }
                var v73 bool = (((32767 != 1) || (((0.0 / 1.5) > 3.14159) && v40)) || (260 < 2))
                _ = fi70
                _ = v71
                _ = v72
                _ = v73
            }
            v3 = append(v3, ("" + "\n\t"))
        }
        _ = fi66
        _ = v67
        _ = v68
    }
    v2 = v40
    for fi74 := 0; fi74 < 1; fi74++ {
        for fi75 := 0; fi75 < 3; fi75++ {
            var v76 uint16 = 42
            v2 = (!((^(v76)) != 255))
            _ = len(v3)
            _ = fi75
            _ = v76
        }
        var v77 float64 = 1e10
        _ = fi74
        _ = v77
    }
    fmt.Println(v2)
    v40 = (65535 >= 177)
    for fi78 := 0; fi78 < 4; fi78++ {
        fmt.Println(v40)
        for fi79 := 0; fi79 < 8; fi79++ {
            if v2 {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v80 := GenS8{F0: 42, F1: v40, F2: 42}
                _ = v80
            } else {
                var v81 uint8 = 140
                if (507 < 1) {
                    v2 = v40
                    _ = len(v1)
                    if v2 {
                        var v82 int64 = -9223372036854775808
                        switch -305 {
                        case -332:
                            switch "foo bar" {
                            case "hello":
                                var v83 float64 = 3.14159
                                _ = v83
                            case "":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                var v84 uint8 = (v81 << ((0) & 63))
                                _ = v84
                            }
                        case -1:
                            if v40 {
                                v3 = append(v3, "\n\t")
                                v1 = append(v1, 862)
                                if len(v3) > 0 {
                                    v3[0] = "foo bar"
                                }
                            }
                        }
                        _ = v82
                    }
                }
                _ = v81
            }
            _ = fi79
        }
        _ = fi78
    }
    if (!(!true)) {
        for fi85 := 0; fi85 < 14; fi85++ {
            v2 = v2
            if v2 {
                for fi86 := 0; fi86 < 17; fi86++ {
                    v40 = v40
                    var v87 uint16 = 0
                    v40 = ((1 != -1) || v2)
                    _ = fi86
                    _ = v87
                }
                v40 = (!("" < ""))
                if true {
                    var v88 uint8 = 1
                    _ = v88
                } else {
                    if v40 {
                        if (!((fi85 % ((fi85) | 1)) < (fi85 & fi85))) {
                            for fi89 := 0; fi89 < 10; fi89++ {
                                fmt.Println(v40, fi89)
                                v40 = (1 != 42)
                                _ = fi89
                            }
                        }
                    }
                }
            } else {
                for fi90 := 0; fi90 < 7; fi90++ {
                    v91 := GenS2{F0: 0, F1: -182, F2: "hello"}
                    if v2 {
                        v40 = (0 < 0)
                        v92 := GenS9{F0: 1}
                        _ = v92
                    }
                    var v93 float32 = 3.14159
                    _ = fi90
                    _ = v91
                    _ = v93
                }
                for fi94 := 0; fi94 < 7; fi94++ {
                    var v95 int8 = -128
                    _ = len(v1)
                    _ = fi94
                    _ = v95
                }
            }
            _ = fi85
        }
        var v96 float32 = float32(1e10)
        _ = v96
    }
    if (4294967295 >= 1) {
        v2 = (false && v2)
        var v97 uint = 42
        v98 := GenS6{F0: 72, F1: 150, F2: v2}
        _ = v97
        _ = v98
    } else {
        v99 := GenS10{F0: 1, F1: (("foo bar" + "hello") + ("hello" + ""))}
        fmt.Println(v2, v40)
        _ = v99
    }
    _ = len(v3)
    switch 127 {
    case 128:
        v40 = v2
    case 0:
        _ = len(v3)
    }
    for fi100 := 0; fi100 < 13; fi100++ {
        v101 := GenS6{F0: 255, F1: 61, F2: (357 != 444)}
        _ = fi100
        _ = v101
    }
    for fi102 := 0; fi102 < 12; fi102++ {
        var v103 uint = 504
        _ = fi102
        _ = v103
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
}
