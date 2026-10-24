package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint8; F1 uint64; F2 string }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int64; F1 uint32; F2 uint }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 uint8; F1 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 string; F1 uint32; F2 uint64 }
type GenS13 struct { F0 int32; F1 int32; F2 bool }
type GenS14 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS15 struct { F0 int; F1 bool; F2 int16 }
type GenS16 struct { F0 int8 }
func genFn0(p0 string, p1 int32) string {
    var v6 uint16 = 152
    v6 = (^(v6))
    var v7 float32 = 1.5
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: -1, F1: 186}
            _ = fi9
            _ = v10
        }
        if (!false) {
            var v11 int32 = (-(p1))
            if (false || (!(2 <= 0))) {
                var v12 uint8 = 127
                _ = v12
            } else {
                for fi13 := 0; fi13 < 16; fi13++ {
                    if (false && true) {
                        switch p0 {
                        case "":
                            if (fi13 >= 9223372036854775807) {
                                var v14 uint64 = 18446744073709551615
                                var v15 bool = (p0 <= ("\n\t" + ((p0 + p0) + p0)))
                                v16 := []int{int(v11)}
                                if v15 {
                                    var v17 bool = v15
                                    v6 = v6
                                    _ = v17
                                } else {
                                    p0 = p0
                                    fmt.Println(fi8, p1)
                                    p1 = p1
                                    switch ((p0 + ("hello" + ((((p0 + ((p0 + p0) + ("\n\t" + "a"))) + ((p0 + p0) + p0)) + "\n\t") + ((("hello" + "\n\t") + "") + "\n\t")))) + p0) {
                                    case "hello":
                                        var v18 int64 = -1
                                        if true {
                                            v19 := GenS2{F0: 255, F1: v15}
                                            _ = v19
                                        } else {
                                            v7 += (v7 * (float32((3.14159 / float64(3.14159))) - (v7 - v7)))
                                        }
                                        if true {
                                            v20 := GenS3{F0: 255, F1: (v14 / ((v14) | 1)), F2: ("a" + "\n\t")}
                                            v11 -= (p1 & (p1 + 345))
                                            switch 128 {
                                            case 128:
                                                _ = len(v16)
                                            }
                                            var v21 int64 = v18
                                            var v22 uint = 185
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        }
                                        _ = v18
                                    case "foo bar":
                                        v23 := GenS4{F0: 2, F1: true}
                                        _ = v23
                                    case "\n\t":
                                        if (1 > -601) {
                                            v16 = append(v16, fi8)
                                        }
                                    }
                                    _ = len(v16)
                                }
                                _ = v14
                                _ = v15
                                _ = v16
                            } else {
                                var v24 int32 = (p1 << ((72) & 63))
                                if true {
                                    v6 = 2
                                } else {
                                    switch (fi8 & fi8) {
                                    case 42:
                                        v25 := GenS4{F0: 281, F1: (-128 > 0)}
                                        _ = v25
                                    case -9223372036854775808:
                                        if ((p1 << ((496) & 63)) != 849) {
                                            var v26 string = p0
                                            if true {
                                                v27 := []int{(fi13 ^ fi8), 0, 675, fi13}
                                                fmt.Println(v6, p1)
                                                _ = v27
                                            }
                                            if ((v7 * v7) >= ((-2.25 - float32(3.14159)) / v7)) {
                                                v28 := GenS3{F0: 128, F1: 42, F2: p0}
                                                v26 = v26
                                                _ = v28
                                            }
                                            _ = v26
                                        }
                                    default:
                                        if ((fi13 * -367) > fi8) {
                                            if false {
                                                p0 = (p0 + (p0 + p0))
                                                v7 = v7
                                                fmt.Println(p0, v6, fi13)
                                            } else {
                                                p1 = (v11 >> ((65535) & 63))
                                            }
                                            p1 = (v11 - (p1 - -90))
                                        } else {
                                            fmt.Println(fi13)
                                            v29 := GenS5{F0: -875, F1: 90, F2: 772}
                                            _ = v29
                                        }
                                    }
                                    v24 -= (p1 >> ((18446744073709551615) & 63))
                                }
                                v30 := GenS1{F0: -820, F1: 18446744073709551615}
                                _ = v24
                                _ = v30
                            }
                        case "foo bar":
                            v31 := []int{}
                            fmt.Println(fi8)
                            var v32 int32 = (p1 + p1)
                            v7 += v7
                            _ = v31
                            _ = v32
                        default:
                            var v33 int = (fi8 >> ((199) & 63))
                            _ = v33
                        }
                        var v34 bool = true
                        if v34 {
                            fmt.Println(v11, p0, v6)
                            var v35 int = fi8
                            var v36 uint8 = 2
                            _ = v35
                            _ = v36
                        }
                        _ = v34
                    }
                    v37 := GenS2{F0: 65535, F1: true}
                    if (0 < 2) {
                        v11 += v11
                    }
                    _ = fi13
                    _ = v37
                }
            }
            switch 32767 {
            case -686:
                v38 := []int{fi8, fi8}
                _ = v38
            case -2147483648:
                v39 := []int{(fi8 & (^(fi8))), (fi8 ^ (fi8 | (fi8 << ((65535) & 63)))), (^(fi8))}
                _ = v39
            }
            _ = v11
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 uint16) bool {
    v40 := GenS6{F0: -969, F1: 2147483648}
    var v41 uint32 = 2147483647
    _ = p0
    _ = v40
    _ = v41
    return (!(true && ((p0 - 1) >= p0)))
}

func main() {
    var v1 int32 = -2147483648
    var v2 float64 = 3.14159
    var v3 uint64 = 4294967295
    v4 := make(map[string]uint32)
    var v5 int32 = -190
    v1 += v5
    for fi42 := 0; fi42 < 2; fi42++ {
        var v43 uint16 = 47
        _ = fi42
        _ = v43
    }
    var v44 int8 = 126
    var v45 int8 = -13
    if (!true) {
        if true {
            for fi46 := 0; fi46 < 14; fi46++ {
                switch "foo bar" {
                case "":
                    v44 += (v45 * v45)
                case "a":
                    var v47 uint = 789
                    _ = v47
                case "\n\t":
                    switch "hello" {
                    case "\n\t":
                        v48 := GenS7{F0: (fi46 * (fi46 << ((18446744073709551615) & 63)))}
                        _ = v48
                    case "foo bar":
                        v3 -= 18446744073709551615
                    }
                }
                _ = fi46
            }
        }
        for fi49 := 0; fi49 < 16; fi49++ {
            var v50 uint8 = 57
            v51 := []int{}
            var v52 uint16 = 65535
            _ = fi49
            _ = v50
            _ = v51
            _ = v52
        }
    } else {
        for fi53 := 0; fi53 < 14; fi53++ {
            v2 += (v2 + ((v2 * ((v2 / v2) / v2)) * 1.5))
            for fi54 := 0; fi54 < 12; fi54++ {
                var v55 int = 42
                v56 := GenS8{F0: 42, F1: (v3 & v3)}
                if false {
                    switch "a" {
                    case "":
                        v57 := GenS9{F0: 1, F1: 42}
                        _ = v57
                    case "foo bar":
                        v3 -= (v3 - (v3 ^ (v3 - v3)))
                    case "a":
                        v55 = v55
                    }
                    v2 += v2
                }
                v45 = (v45 % ((v44) | 1))
                var v58 int64 = 120
                _ = fi54
                _ = v55
                _ = v56
                _ = v58
            }
            switch "\n\t" {
            case "a":
                switch "a" {
                case "hello":
                    v59 := GenS5{F0: 9223372036854775807, F1: 65535, F2: 0}
                    _ = v59
                case "foo bar":
                    v44 -= (v44 * 1)
                case "\n\t":
                    var v60 int8 = (v44 | (v44 % ((v45) | 1)))
                    _ = v60
                }
            case "":
                for fi61 := 0; fi61 < 1; fi61++ {
                    fmt.Println(v45, v1, v44)
                    fmt.Println(v2)
                    v62 := v4["foo bar"]
                    _ = fi61
                    _ = v62
                }
            default:
                switch "" {
                case "hello":
                    var v63 float32 = -2.25
                    _ = v63
                case "foo bar":
                    fmt.Println(v2)
                default:
                    var v64 uint8 = 88
                    _ = v64
                }
            }
            _ = fi53
        }
    }
    v65 := []int{9223372036854775807}
    for fi66 := 0; fi66 < 14; fi66++ {
        if false {
            for fi67 := 0; fi67 < 3; fi67++ {
                v68 := GenS10{F0: 4294967295}
                if (581 > 0) {
                    v2 = v2
                    v5 += 61
                    _ = len(v65)
                    var v69 int = fi66
                    _ = v69
                } else {
                    if false {
                        v70 := GenS11{F0: "hello"}
                        v71 := GenS12{F0: (("\n\t" + ("" + ("hello" + "foo bar"))) + ("" + "")), F1: 1, F2: v3}
                        v72 := GenS10{F0: 360}
                        _ = v70
                        _ = v71
                        _ = v72
                    } else {
                        v73 := v4["a"]
                        fmt.Println(v44)
                        _ = v73
                    }
                    var v74 int64 = -9223372036854775808
                    if ((255 > 18446744073709551615) || ((v2 * v2) < v2)) {
                        v44 -= (^(v44))
                    }
                    _ = v74
                }
                switch (fi67 / (((fi66 % (((fi67 | (fi66 * int(v45)))) | 1))) | 1)) {
                case 127:
                    var v75 bool = true
                    _ = v75
                case -704:
                    fmt.Println(v3)
                    v76 := GenS5{F0: 42, F1: 1, F2: 18446744073709551615}
                    if len(v65) > 0 {
                        v65[0] = fi67
                    }
                    _ = v76
                case 656:
                    var v77 float64 = v2
                    _ = v77
                }
                _ = fi67
                _ = v68
            }
            if len(v65) > 0 {
                v65[0] = (^(fi66))
            }
        }
        if (("hello" + "hello") > (("" + "foo bar") + "a")) {
            if (556 < 0) {
                v44 = v44
                if false {
                    v44 = v44
                    var v78 float32 = 0.1
                    _ = v78
                } else {
                    switch (("hello" + "a") + ("foo bar" + "hello")) {
                    case "\n\t":
                        v65 = append(v65, fi66)
                    case "foo bar":
                        _ = len(v4)
                    case "hello":
                        var v79 int8 = 100
                        _ = v79
                    default:
                        var v80 uint64 = v3
                        _ = v80
                    }
                    for fi81 := 0; fi81 < 20; fi81++ {
                        fmt.Println(v1, fi66, v5)
                        v82 := GenS13{F0: (-(v1)), F1: v5, F2: false}
                        var v83 int32 = (v1 >> ((860) & 63))
                        _ = fi81
                        _ = v82
                        _ = v83
                    }
                    switch ("\n\t" + "hello") {
                    case "":
                        if true {
                            switch "" {
                            case "hello":
                                v84 := GenS9{F0: 1, F1: (^(v3))}
                                _ = v84
                            case "":
                                delete(v4, "hello")
                            case "a":
                                v85 := GenS6{F0: 0, F1: uint32(v3)}
                                _ = v85
                            }
                        }
                    case "\n\t":
                        v86 := GenS2{F0: 255, F1: false}
                        _ = v86
                    default:
                        for fi87 := 0; fi87 < 15; fi87++ {
                            switch ("\n\t" + "foo bar") {
                            case "foo bar":
                                if (false && false) {
                                    v44 = (v45 + (^(v45)))
                                    var v88 int16 = -641
                                    _ = v88
                                } else {
                                    if true {
                                        _ = len(v65)
                                        var v89 int16 = 6
                                        _ = v89
                                    } else {
                                        v65 = append(v65, fi87)
                                    }
                                    v90 := GenS14{F0: -1, F1: 4294967295, F2: 42}
                                    fmt.Println(v45, fi87)
                                    _ = v90
                                }
                            case "hello":
                                var v91 uint8 = 255
                                _ = v91
                            default:
                                v92 := GenS15{F0: (fi87 / (((fi87 | fi87)) | 1)), F1: (!(!false)), F2: int16(v44)}
                                _ = v92
                            }
                            v93 := GenS15{F0: fi87, F1: (0.0 != (1.5 + ((float32((v2 * 0.1)) - 1e10) * 0.1))), F2: 735}
                            v44 = (v45 >> ((691) & 63))
                            _ = fi87
                            _ = v93
                        }
                    }
                }
                for fi94 := 0; fi94 < 4; fi94++ {
                    fmt.Println(v45, v44)
                    var v95 uint8 = 2
                    v96 := v4[("foo bar" + ("\n\t" + "hello"))]
                    v96 = (v96 + (v96 & 42))
                    _ = fi94
                    _ = v95
                    _ = v96
                }
            }
        } else {
            v97 := GenS3{F0: uint8(v3), F1: v3, F2: "a"}
            switch ("" + (("foo bar" + "a") + "hello")) {
            case "":
                var v98 float64 = 1.5
                _ = v98
            }
            if len(v65) >= 2 {
                v65 = v65[1:len(v65)]
            }
            _ = v97
        }
        _ = fi66
    }
    for fi99 := 0; fi99 < 2; fi99++ {
        v2 = (float64(1e10) * ((float64(0.1) / 1e10) * float64(-2.25)))
        _ = len(v4)
        _ = fi99
    }
    for fi100 := 0; fi100 < 20; fi100++ {
        switch fi100 {
        case 32767:
            switch "" {
            case "a":
                _ = len(v65)
            }
        case 1:
            for fi101 := 0; fi101 < 19; fi101++ {
                var v102 string = ("foo bar" + "hello")
                _ = len(v4)
                if false {
                    v45 = v44
                } else {
                    if len(v65) >= 2 {
                        v65 = v65[1:len(v65)]
                    }
                }
                _ = fi101
                _ = v102
            }
        case 0:
            switch 0 {
            case -2147483648:
                v4[""] = 500
            case 2147483647:
                var v103 int16 = -1
                _ = v103
            case 32767:
                if ((-(fi100)) != (fi100 & (-(fi100)))) {
                    v1 = (v1 << ((2) & 63))
                    _ = len(v65)
                }
            default:
                v4["foo bar"] = 413
            }
        }
        var v104 int8 = v44
        switch "" {
        case "foo bar":
            if true {
                v105 := v4[("a" + "")]
                _ = v105
            }
        }
        _ = fi100
        _ = v104
    }
    fmt.Println(v45, v2)
    v106 := GenS15{F0: 4294967295, F1: (314 == 4294967295), F2: -32768}
    if (0 != 1) {
        for fi107 := 0; fi107 < 9; fi107++ {
            v108 := GenS9{F0: 255, F1: 4294967295}
            _ = len(v65)
            _ = fi107
            _ = v108
        }
        fmt.Println(v45)
    } else {
        switch "foo bar" {
        case "foo bar":
            v109 := GenS16{F0: -90}
            _ = v109
        default:
            if ((-(v5)) <= v1) {
                var v110 uint8 = 128
                _ = v110
            } else {
                v111 := GenS12{F0: "a", F1: 2, F2: (v3 + (v3 | (v3 & v3)))}
                fmt.Println(v5, v1, v44)
                v2 += (float64(0.0) + -2.25)
                _ = v111
            }
        }
    }
    v45 = v45
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v65:", len(v65), v65)
    _ = v1
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v65
}
