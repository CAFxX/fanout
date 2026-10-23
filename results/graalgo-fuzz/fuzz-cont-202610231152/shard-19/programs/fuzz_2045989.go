package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 int8; F2 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64; F1 int8; F2 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int64; F1 uint16; F2 int16 }
type GenS11 struct { F0 int16; F1 uint }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int64; F1 bool; F2 int32 }
type GenS16 struct { F0 int16; F1 int16 }
func genFn0() string {
    if true {
        var v7 bool = (!false)
        _ = v7
    } else {
        var v8 uint = 42
        v8 = (v8 >> (((v8 - v8)) & 63))
        _ = v8
    }
    var v9 int8 = 45
    _ = v9
    return ("foo bar" + "\n\t")
}

func main() {
    var v1 int = -574
    var v2 uint = 889
    v3 := []string{"hello"}
    var v4 uint8 = 127
    var v5 uint = v2
    v6 := make(map[int]string)
    v6[32768] = ("a" + "\n\t")
    var v10 uint64 = 900
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint32 = 4294967295
        switch (fi11 - (fi11 * (fi11 & v1))) {
        case 128:
            if false {
                v1 += (fi11 + 127)
                if true {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        switch v1 {
                        case 9223372036854775807:
                            v14 := GenS1{F0: 0}
                            fmt.Println(v4, v1)
                            _ = v14
                        case 127:
                            v3 = append(v3, "\n\t")
                        default:
                            var v15 uint32 = (v12 / ((v12) | 1))
                            _ = v15
                        }
                        if false {
                            v16 := GenS2{F0: 0, F1: v10, F2: fi11}
                            if false {
                                v3 = append(v3, "foo bar")
                                if ((("hello" + ("hello" + ("" + ("hello" + "foo bar")))) + "hello") <= ("a" + "a")) {
                                    switch fi13 {
                                    case -1:
                                        if false {
                                            var v17 bool = false
                                            v1 += v1
                                            var v18 uint64 = (^(v10))
                                            _ = v17
                                            _ = v18
                                        } else {
                                            fmt.Println(v1, v12)
                                        }
                                    case 128:
                                        if ((true && (((v1 - (fi11 | (^(v1)))) != (fi13 ^ 128)) && false)) || (1 != 2)) {
                                            if (v5 > v2) {
                                                var v19 int16 = 1
                                                var v20 uint16 = 42
                                                _ = v19
                                                _ = v20
                                            } else {
                                                v21 := GenS3{F0: 957}
                                                _ = v21
                                            }
                                            v22 := GenS4{F0: true, F1: 100, F2: 776}
                                            _ = v22
                                        } else {
                                            var v23 bool = (v12 > (v12 << (((v2 / ((2) | 1))) & 63)))
                                            v1 += (v1 & v1)
                                            fmt.Println(fi11, v23)
                                            _ = v23
                                        }
                                    default:
                                        fmt.Println(fi11, fi11)
                                    }
                                    if true {
                                        v10 = v10
                                        var v24 int32 = -256
                                        _ = v24
                                    }
                                    fmt.Println(v1, fi13)
                                    if (!(!((("a" + (((("\n\t" + "a") + "\n\t") + "hello") + ((("\n\t" + "") + ("\n\t" + "\n\t")) + "\n\t"))) > ("a" + "\n\t")) || (false && (true && true))))) {
                                        var v25 int32 = -2147483648
                                        _ = v25
                                    }
                                } else {
                                    v1 -= v1
                                    v3 = append(v3, "")
                                }
                                _ = len(v3)
                                switch (v1 + (v1 - 248)) {
                                case 0:
                                    fmt.Println(fi13, v12, v5)
                                case -1:
                                    switch ("\n\t" + ("\n\t" + "hello")) {
                                    case "a":
                                        v5 += 4294967295
                                    default:
                                        if false {
                                            v26 := GenS5{F0: -1}
                                            _ = v26
                                        } else {
                                            if ("\n\t" <= "") {
                                                if true {
                                                    var v27 uint = (v2 | (v5 & v5))
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v27
                                                } else {
                                                    v5 = v2
                                                    v28 := GenS5{F0: -37}
                                                    v29 := GenS6{F0: v10}
                                                    _ = v28
                                                    _ = v29
                                                }
                                            } else {
                                                var v30 int16 = 307
                                                v31 := v6[-1]
                                                _ = v30
                                                _ = v31
                                            }
                                            if (false && (-2.25 == 1.5)) {
                                                var v32 bool = (((0.1 / 1.0) != 3.14159) || true)
                                                var v33 int16 = -1
                                                _ = v32
                                                _ = v33
                                            } else {
                                                v4 = v4
                                                switch "\n\t" {
                                                case "":
                                                    v34 := v6[(v1 % ((543) | 1))]
                                                    _ = v34
                                                case "\n\t":
                                                    if ((v4 << (((v5 * (^(v5)))) & 63)) == v4) {
                                                        var v35 float64 = 1e10
                                                        if ((!false) || ((v1 * fi13) <= (fi13 & 1))) {
                                                            v5 -= v5
                                                            v10 -= 2
                                                        } else {
                                                            v36 := GenS7{F0: -193}
                                                            var v37 uint = v2
                                                            _ = len(v6)
                                                            v12 += v12
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        fmt.Println(v35, v2)
                                                        _ = v35
                                                    } else {
                                                        var v38 int32 = 1
                                                        var v39 int64 = 1
                                                        var v40 uint8 = (^(v4))
                                                        _ = v38
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                }
                                            }
                                            if true {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v41 float64 = 1.5
                                                _ = v41
                                            } else {
                                                v2 -= 42
                                                if (42 == uint16(v5)) {
                                                    v10 = v10
                                                    var v42 bool = true
                                                    var v43 int = (v1 | (-(fi13)))
                                                    if len(v3) > 0 {
                                                        v3[0] = (("foo bar" + "foo bar") + ("foo bar" + ("foo bar" + ("foo bar" + "hello"))))
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = "foo bar"
                                                    }
                                                    _ = v42
                                                    _ = v43
                                                }
                                            }
                                        }
                                    }
                                }
                                fmt.Println(v5)
                            } else {
                                if (!(1e10 == ((3.14159 / -2.25) + 1e10))) {
                                    v44 := GenS8{F0: 42, F1: -100, F2: 2}
                                    v45 := GenS4{F0: false, F1: -1, F2: -2147483648}
                                    _ = v44
                                    _ = v45
                                } else {
                                    if (!((("foo bar" + "\n\t") + "foo bar") <= "")) {
                                        v46 := GenS6{F0: (^(v10))}
                                        var v47 int16 = 136
                                        var v48 bool = true
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    } else {
                                        delete(v6, (fi11 >> ((v2) & 63)))
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v49 uint32 = (v12 & (v12 * 42))
                                            _ = v49
                                        case "":
                                            v50 := GenS6{F0: (v10 % ((v10) | 1))}
                                            _ = v50
                                        case "a":
                                            if (false || (!false)) {
                                                v10 -= (^(v10))
                                                switch ("foo bar" + ("hello" + ((("\n\t" + "hello") + ("a" + "foo bar")) + "\n\t"))) {
                                                case "hello":
                                                    v12 -= (^(v12))
                                                default:
                                                    v10 += v10
                                                }
                                                switch fi11 {
                                                case 128:
                                                    v12 = v12
                                                case -754:
                                                    fmt.Println(v5)
                                                default:
                                                    fmt.Println(v5, v1)
                                                    v51 := GenS2{F0: uint16(v4), F1: (^(v10)), F2: fi13}
                                                    var v52 float64 = 1e10
                                                    _ = v51
                                                    _ = v52
                                                }
                                            } else {
                                                var v53 uint8 = 1
                                                v1 += (fi11 & -1)
                                                v2 = (v2 * (v2 & (v2 / ((v2) | 1))))
                                                _ = v53
                                            }
                                        }
                                        switch (fi11 + 9223372036854775807) {
                                        case -905:
                                            var v54 uint = (^(v2))
                                            _ = v54
                                        default:
                                            v55 := GenS9{F0: (v12 << (((v5 | 255)) & 63))}
                                            _ = v55
                                        }
                                    }
                                    fmt.Println(v4, v1, fi13)
                                }
                                v5 = v5
                            }
                            _ = v16
                        }
                        _ = fi13
                    }
                    for fi56 := 0; fi56 < 14; fi56++ {
                        v12 += v12
                        _ = fi56
                    }
                    v1 -= int(v12)
                }
                var v57 int = v1
                _ = v57
            } else {
                var v58 bool = true
                delete(v6, (v1 - 128))
                switch -2147483648 {
                case 127:
                    var v59 uint16 = 0
                    _ = v59
                case 128:
                    switch v1 {
                    case 578:
                        fmt.Println(v5, v4)
                    case 2147483647:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 4294967295:
                        switch v1 {
                        case 32767:
                            if ((255 != (v2 >> ((v2) & 63))) && v58) {
                                v60 := GenS10{F0: 1, F1: 2, F2: 0}
                                v58 = ((v4 - v4) != v4)
                                _ = v60
                            } else {
                                if (-66 >= -83) {
                                    _ = len(v6)
                                }
                                for fi61 := 0; fi61 < 12; fi61++ {
                                    if (v58 || false) {
                                        var v62 uint32 = 2147483647
                                        v63 := GenS2{F0: 130, F1: (v10 & v10), F2: (fi61 / (((fi61 | (fi11 + (fi61 / ((fi61) | 1))))) | 1))}
                                        v12 -= 180
                                        switch ("\n\t" + "foo bar") {
                                        case "foo bar":
                                            v64 := GenS11{F0: -73, F1: (v5 | uint(v10))}
                                            _ = v64
                                        }
                                        v65 := v6[fi61]
                                        fmt.Println(v12, v4, fi11)
                                        _ = v62
                                        _ = v63
                                        _ = v65
                                    }
                                    var v66 string = "\n\t"
                                    _ = fi61
                                    _ = v66
                                }
                            }
                        }
                    default:
                        var v67 uint = (v2 + (^(v5)))
                        _ = v67
                    }
                }
                _ = v58
            }
        case 761:
            v68 := GenS12{F0: 9223372036854775807}
            _ = v68
        default:
            if (("hello" > "foo bar") && (1 < 241)) {
                for fi69 := 0; fi69 < 10; fi69++ {
                    v70 := GenS13{F0: false}
                    var v71 int16 = 32767
                    _ = fi69
                    _ = v70
                    _ = v71
                }
                for fi72 := 0; fi72 < 13; fi72++ {
                    _ = len(v3)
                    fmt.Println(fi72, v2)
                    _ = fi72
                }
                _ = len(v6)
            }
        }
        for fi73 := 0; fi73 < 10; fi73++ {
            v6[fi11] = ""
            var v74 int = v1
            v2 += 1
            _ = fi73
            _ = v74
        }
        _ = fi11
        _ = v12
    }
    v1 += int(v4)
    v75 := GenS14{F0: (^(v10))}
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    switch (v1 & 782) {
    case 273:
        v2 -= (v2 / (((v5 * v5)) | 1))
    case 1:
        fmt.Println(v10, v1)
    case 2147483647:
        if (true || (false && (307 <= 440))) {
            switch "hello" {
            case "\n\t":
                var v76 bool = (-2.25 <= -2.25)
                _ = v76
            case "":
                if true {
                    if true {
                        v5 -= (v2 / (((v5 ^ v2)) | 1))
                        fmt.Println(v10)
                        for fi77 := 0; fi77 < 2; fi77++ {
                            var v78 int16 = 32767
                            fmt.Println(v2, v2)
                            _ = fi77
                            _ = v78
                        }
                    }
                }
            default:
                if false {
                    _ = len(v3)
                    v5 += v2
                    for fi79 := 0; fi79 < 13; fi79++ {
                        var v80 uint64 = v10
                        v10 += (^(v80))
                        var v81 uint = 2
                        _ = fi79
                        _ = v80
                        _ = v81
                    }
                } else {
                    switch ("" + "foo bar") {
                    case "hello":
                        v5 -= 4294967295
                    default:
                        switch (-(v1)) {
                        case 32767:
                            var v82 int32 = 2147483647
                            _ = v82
                        case 9223372036854775807:
                            v3 = append(v3, (("\n\t" + "hello") + ("a" + "a")))
                        case 0:
                            delete(v6, v1)
                        }
                    }
                }
            }
            if (((42 <= 32767) || true) || false) {
                for fi83 := 0; fi83 < 3; fi83++ {
                    v84 := GenS15{F0: 1, F1: (-2147483648 <= (-(fi83))), F2: 0}
                    _ = fi83
                    _ = v84
                }
                for fi85 := 0; fi85 < 15; fi85++ {
                    v2 -= v5
                    _ = fi85
                }
                v86 := GenS13{F0: (v1 != v1)}
                _ = v86
            }
            v6[v1] = "hello"
        }
    default:
        v4 = 127
    }
    switch (v1 + 32767) {
    case 32767:
        if false {
            fmt.Println(v5, v5, v1)
            var v87 float32 = (float32(float64(((1.5 + (3.14159 + 3.14159)) - 3.14159))) * 1e10)
            delete(v6, int(v10))
            _ = v87
        }
    default:
        if len(v3) > 0 {
            v3[0] = ""
        }
    }
    var v88 int32 = 1
    var v89 uint32 = 1
    for fi90 := 0; fi90 < 1; fi90++ {
        var v91 uint8 = (^(v4))
        v92 := GenS16{F0: int16(v4), F1: int16(v4)}
        _ = fi90
        _ = v91
        _ = v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v88
    _ = v89
}
