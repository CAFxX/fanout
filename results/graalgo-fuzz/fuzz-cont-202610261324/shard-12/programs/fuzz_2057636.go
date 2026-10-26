package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 int }
type GenS4 struct { F0 int; F1 int; F2 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint32; F2 uint8 }
func genFn0() uint16 {
    if true {
        switch (("hello" + "foo bar") + "hello") {
        case "foo bar":
            var v6 int8 = 127
            _ = v6
        case "hello":
            var v7 uint8 = 128
            _ = v7
        case "":
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 bool = false
                var v10 uint32 = 948
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        v11 := GenS1{F0: 5}
        _ = v11
    }
    var v12 int16 = 32767
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: (fi13 & fi13)}
        var v15 uint32 = 2
        _ = fi13
        _ = v14
        _ = v15
    }
    switch "\n\t" {
    case "hello":
        switch "\n\t" {
        case "hello":
            v16 := GenS2{F0: 2, F1: 13}
            _ = v16
        case "\n\t":
            switch "\n\t" {
            case "a":
                fmt.Println(v12)
            case "":
                v17 := []int{}
                _ = v17
            case "\n\t":
                v18 := GenS3{F0: 1, F1: 42, F2: 9223372036854775807}
                _ = v18
            }
        case "":
            var v19 uint = 42
            _ = v19
        default:
            for fi20 := 0; fi20 < 3; fi20++ {
                switch 4294967295 {
                case 32768:
                    v12 += v12
                }
                fmt.Println(v12, fi20)
                var v21 string = "hello"
                _ = fi20
                _ = v21
            }
        }
    }
    var v22 uint64 = 799
    _ = v12
    _ = v22
    return 162
}
func genFn1() int16 {
    var v23 bool = (255 <= 0)
    v24 := GenS3{F0: genFn0(), F1: 1, F2: 962}
    var v25 float32 = 0.1
    _ = v23
    _ = v24
    _ = v25
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 18446744073709551615
    v3 := []int{127}
    var v4 uint16 = 793
    var v5 uint16 = (^(v4))
    if true {
        var v26 int8 = 42
        _ = v26
    }
    switch ("a" + ("\n\t" + "hello")) {
    case "a":
        if ((128 >= 127) && false) {
            v27 := GenS4{F0: 4294967295, F1: 192, F2: 2}
            _ = v27
        }
    case "":
        delete(v1, ("" + "hello"))
    default:
        v2 -= 1
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        switch ("hello" + ("foo bar" + "")) {
        case "hello":
            v4 = (v4 % (((v4 ^ v5)) | 1))
        }
        switch "a" {
        case "":
            var v29 uint64 = 0
            _ = v29
        case "hello":
            for fi30 := 0; fi30 < 16; fi30++ {
                fmt.Println(fi28)
                var v31 int16 = 0
                _ = fi30
                _ = v31
            }
        case "a":
            if (-420 == 0) {
                fmt.Println(v2, v4)
                v5 = (v4 ^ 65535)
            } else {
                for fi32 := 0; fi32 < 18; fi32++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch "" {
                    case "foo bar":
                        v4 = (v4 >> ((773) & 63))
                    }
                    _ = fi32
                }
            }
        default:
            v4 = genFn0()
        }
        _ = fi28
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        if ((1.5 != 3.14159) || false) {
            var v34 uint64 = 1
            _ = v34
        } else {
            if false {
                if false {
                    v5 = genFn0()
                } else {
                    switch "hello" {
                    case "":
                        fmt.Println(v4)
                    }
                }
            } else {
                if false {
                    if (true || (42 == -9223372036854775808)) {
                        v3 = append(v3, fi33)
                        if false {
                            _ = len(v1)
                        } else {
                            var v35 int32 = -2147483648
                            var v36 int64 = int64(v5)
                            _ = v35
                            _ = v36
                        }
                        var v37 int = fi33
                        _ = v37
                    } else {
                        v2 = (v2 % ((1) | 1))
                    }
                } else {
                    v5 += v5
                    var v38 bool = true
                    v39 := GenS2{F0: 76, F1: 2}
                    _ = v38
                    _ = v39
                }
            }
            switch (("" + ("" + ((("a" + (((((("hello" + "foo bar") + (("\n\t" + "a") + ("" + "a"))) + "\n\t") + ("a" + "a")) + (("a" + "\n\t") + ("" + ""))) + ("" + "a"))) + ("a" + (("hello" + "") + "a"))) + (("hello" + (("" + "\n\t") + ("" + ("a" + (("foo bar" + (("foo bar" + ("hello" + "\n\t")) + "")) + "foo bar"))))) + "a")))) + "\n\t") {
            case "":
                for fi40 := 0; fi40 < 5; fi40++ {
                    if true {
                        var v41 int32 = 2147483647
                        if ((v4 << ((1) & 63)) > v4) {
                            switch fi40 {
                            case -9223372036854775808:
                                fmt.Println(v5, fi33, v41)
                            case 32768:
                                var v42 int8 = -19
                                _ = v42
                            }
                            var v43 int16 = -32768
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v43
                        } else {
                            v41 -= 42
                            v5 += 6
                        }
                        _ = v41
                    } else {
                        v44 := GenS2{F0: 160, F1: 65535}
                        v45 := GenS2{F0: 0, F1: 65535}
                        _ = v44
                        _ = v45
                    }
                    if (false || false) {
                        fmt.Println(v5, v2)
                    }
                    _ = fi40
                }
            case "foo bar":
                fmt.Println(v2)
            }
            if false {
                v4 = (v4 | v4)
            }
        }
        var v46 uint = 0
        for fi47 := 0; fi47 < 6; fi47++ {
            v5 = (v4 >> ((v46) & 63))
            v48 := GenS2{F0: 152, F1: (v46 % ((v46) | 1))}
            _ = fi47
            _ = v48
        }
        _ = fi33
        _ = v46
    }
    if (127 > 127) {
        if true {
            for fi49 := 0; fi49 < 1; fi49++ {
                if false {
                    var v50 uint32 = 2147483647
                    v4 -= (v4 >> ((65535) & 63))
                    _ = v50
                } else {
                    var v51 int32 = -1
                    v5 -= v4
                    _ = v51
                }
                var v52 int32 = 0
                _ = fi49
                _ = v52
            }
        }
        if true {
            v4 = (v4 % (((v4 ^ (v4 + (v4 & (v4 & (v4 | v5)))))) | 1))
            for fi53 := 0; fi53 < 11; fi53++ {
                v1["foo bar"] = 574
                v54 := GenS3{F0: 1, F1: 914, F2: fi53}
                _ = fi53
                _ = v54
            }
            fmt.Println(v2, v5, v4)
        } else {
            if true {
                for fi55 := 0; fi55 < 19; fi55++ {
                    for fi56 := 0; fi56 < 7; fi56++ {
                        v5 = (v4 + (v4 | 1))
                        _ = fi56
                    }
                    v57 := GenS2{F0: 127, F1: 18446744073709551615}
                    _ = fi55
                    _ = v57
                }
                var v58 uint8 = 42
                _ = v58
            } else {
                v4 += v4
                if false {
                    v1["hello"] = 1
                } else {
                    _ = len(v3)
                }
            }
        }
    }
    if (false && (!(255 <= 0))) {
        for fi59 := 0; fi59 < 3; fi59++ {
            switch "\n\t" {
            case "hello":
                if (0 != 1) {
                    for fi60 := 0; fi60 < 18; fi60++ {
                        if (1 < 18446744073709551615) {
                            switch fi60 {
                            case -796:
                                if true {
                                    if (42 == 2) {
                                        var v61 uint16 = (v4 + v5)
                                        _ = v61
                                    } else {
                                        if true {
                                            _ = len(v1)
                                            switch 104 {
                                            case -564:
                                                var v62 int8 = 0
                                                _ = v62
                                            case 9223372036854775807:
                                                v2 -= 0
                                            default:
                                                var v63 float64 = ((((((3.14159 / ((0.1 / 0.1) + 0.0)) + 1e10) - -2.25) * 3.14159) + (1.5 * 3.14159)) * 0.1)
                                                _ = v63
                                            }
                                        }
                                        v4 = 42
                                    }
                                    fmt.Println(v2, v2)
                                    v4 = v4
                                    switch "a" {
                                    case "a":
                                        _ = len(v3)
                                        if (true || true) {
                                            _ = len(v1)
                                        }
                                    case "foo bar":
                                        v4 += genFn0()
                                    default:
                                        var v64 int = (fi59 << ((1) & 63))
                                        _ = v64
                                    }
                                } else {
                                    v65 := GenS5{F0: -590}
                                    _ = v65
                                }
                            default:
                                v2 = (v2 / ((1) | 1))
                            }
                            v1[("foo bar" + ("foo bar" + (("\n\t" + (("hello" + "foo bar") + "")) + "\n\t")))] = 65535
                            var v66 int8 = -127
                            var v67 int32 = 2147483647
                            _ = v66
                            _ = v67
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v68 int64 = 42
                        _ = fi60
                        _ = v68
                    }
                }
            case "foo bar":
                v5 = (v4 >> ((1) & 63))
            }
            _ = fi59
        }
        v2 -= (v2 * v2)
        fmt.Println(v5, v5)
    } else {
        for fi69 := 0; fi69 < 20; fi69++ {
            for fi70 := 0; fi70 < 7; fi70++ {
                v71 := v1[(("a" + "\n\t") + "a")]
                fmt.Println(v2, v71)
                _ = fi70
                _ = v71
            }
            switch (fi69 - (fi69 | (fi69 & (fi69 + (fi69 - 9223372036854775807))))) {
            case 2147483647:
                v72 := GenS5{F0: 317}
                _ = v72
            case 42:
                for fi73 := 0; fi73 < 17; fi73++ {
                    v74 := GenS4{F0: (fi69 | fi69), F1: (fi69 + (fi73 - fi73)), F2: 14}
                    v4 += (v5 & 65535)
                    _ = fi73
                    _ = v74
                }
            case -1:
                if false {
                    v75 := GenS3{F0: (v5 + 42), F1: -2147483648, F2: 9223372036854775807}
                    v5 = v5
                    _ = v75
                }
            default:
                v76 := GenS6{F0: 42, F1: 4294967295, F2: 127}
                _ = v76
            }
            _ = fi69
        }
        var v77 int64 = 0
        _ = v77
    }
    fmt.Println(v4, v5)
    v78 := GenS5{F0: 0}
    fmt.Println(v4, v2, v5)
    var v79 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v78
    _ = v79
}
