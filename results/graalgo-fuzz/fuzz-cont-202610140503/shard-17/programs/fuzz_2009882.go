package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int16; F2 int8 }
type GenS3 struct { F0 uint; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 uint; F2 int64 }
type GenS7 struct { F0 int64; F1 uint16; F2 int32 }

func main() {
    var v1 bool = false
    v2 := make(map[string]bool)
    var v3 uint32 = 4294967295
    var v4 uint16 = 1
    if (!(42 <= -1)) {
        v4 = v4
        var v5 uint = 255
        _ = v5
    }
    if true {
        v4 = v4
        v6 := GenS1{F0: (^(v3))}
        _ = v6
    }
    v7 := []int{127, 127, 32767, 128}
    fmt.Println(v1, v4, v1)
    if false {
        switch "a" {
        case "":
            if (42 == -1) {
                v1 = (v1 || (false || (!false)))
                v8 := GenS1{F0: 2}
                switch "hello" {
                case "a":
                    switch 0 {
                    case 128:
                        switch -1 {
                        case 704:
                            fmt.Println(v4)
                        default:
                            if (-502 >= 2147483647) {
                                v4 += 2
                                if false {
                                    v1 = (168 < 9223372036854775807)
                                    for fi9 := 0; fi9 < 15; fi9++ {
                                        if false {
                                            fmt.Println(fi9)
                                            v10 := GenS1{F0: 1}
                                            v4 += 361
                                            _ = v10
                                        } else {
                                            v7 = append(v7, fi9)
                                        }
                                        var v11 int = (fi9 >> ((42) & 63))
                                        if true {
                                            v3 += v3
                                            v12 := GenS1{F0: v3}
                                            for fi13 := 0; fi13 < 10; fi13++ {
                                                fmt.Println(v3)
                                                v11 += fi13
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                                v7 = append(v7, (-(v11)))
                                                _ = fi13
                                            }
                                            _ = v12
                                        }
                                        _ = fi9
                                        _ = v11
                                    }
                                }
                                v4 += 844
                            } else {
                                for fi14 := 0; fi14 < 20; fi14++ {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    _ = fi14
                                }
                                v7 = append(v7, 128)
                            }
                        }
                    case 0:
                        if (2147483647 < 0) {
                            for fi15 := 0; fi15 < 7; fi15++ {
                                if len(v7) > 0 {
                                    v7[0] = (fi15 & fi15)
                                }
                                v16 := GenS1{F0: (v3 - (^(v3)))}
                                _ = len(v7)
                                _ = fi15
                                _ = v16
                            }
                        } else {
                            for fi17 := 0; fi17 < 4; fi17++ {
                                var v18 uint16 = v4
                                if ((v1 || v1) || (-357 > 0)) {
                                    var v19 bool = v1
                                    _ = v19
                                } else {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                }
                                _ = fi17
                                _ = v18
                            }
                        }
                    case 92:
                        if (false && (v1 && v1)) {
                            switch 42 {
                            case 42:
                                for fi20 := 0; fi20 < 15; fi20++ {
                                    if v1 {
                                        switch ("hello" + "foo bar") {
                                        case "\n\t":
                                            v21 := GenS2{F0: fi20, F1: -356, F2: 0}
                                            _ = v21
                                        default:
                                            switch "" {
                                            case "\n\t":
                                                if v1 {
                                                    if v1 {
                                                        v22 := GenS3{F0: 4294967295, F1: (v4 + v4), F2: int64(v3)}
                                                        _ = v22
                                                    }
                                                } else {
                                                    switch int(v3) {
                                                    case 9223372036854775807:
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v1 = (((("foo bar" + ("hello" + "\n\t")) + "foo bar") + (("" + ("\n\t" + (("a" + (("\n\t" + "") + "foo bar")) + ""))) + "hello")) > ("\n\t" + "a"))
                                                        case "\n\t":
                                                            for fi23 := 0; fi23 < 9; fi23++ {
                                                                var v24 bool = false
                                                                if (v3 < v3) {
                                                                    switch (("hello" + "") + "a") {
                                                                    case "a":
                                                                        if false {
                                                                            var v25 uint8 = 18
                                                                            if v24 {
                                                                                switch 0 {
                                                                                case -2147483648:
                                                                                    var v26 float32 = ((-2.25 / 1e10) + 0.1)
                                                                                    _ = v26
                                                                                case 42:
                                                                                    switch fi23 {
                                                                                    case 32768:
                                                                                        v27 := GenS4{F0: 2147483647, F1: -271}
                                                                                        _ = v27
                                                                                    case -163:
                                                                                        switch fi20 {
                                                                                        case 32767:
                                                                                            v4 = v4
                                                                                        case 32768:
                                                                                            if v24 {
                                                                                                var v28 uint16 = v4
                                                                                                _ = v28
                                                                                            }
                                                                                        case 128:
                                                                                            var v29 bool = (v3 == (v3 * (v3 - (v3 % ((v3) | 1)))))
                                                                                            _ = v29
                                                                                        default:
                                                                                            fmt.Println(v1, v24)
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v25)
                                                                                    }
                                                                                }
                                                                                var v30 float64 = 1.5
                                                                                _ = v30
                                                                            }
                                                                            switch (fi20 / ((fi20) | 1)) {
                                                                            case 95:
                                                                                var v31 uint64 = 4294967295
                                                                                _ = v31
                                                                            case -868:
                                                                                var v32 uint64 = 42
                                                                                _ = v32
                                                                            }
                                                                            _ = v25
                                                                        }
                                                                    case "foo bar":
                                                                        if ("\n\t" != ((("\n\t" + "a") + ("foo bar" + "foo bar")) + "")) {
                                                                            var v33 uint = 65535
                                                                            _ = v33
                                                                        } else {
                                                                            var v34 string = "a"
                                                                            var v35 int16 = int16(v4)
                                                                            var v36 int = (fi23 << ((4294967295) & 63))
                                                                            _ = v34
                                                                            _ = v35
                                                                            _ = v36
                                                                        }
                                                                    case "\n\t":
                                                                        delete(v2, ("a" + "\n\t"))
                                                                    default:
                                                                        if len(v7) >= 2 {
                                                                            v7 = v7[1:len(v7)]
                                                                        }
                                                                        fmt.Println(fi20, v3)
                                                                    }
                                                                    v1 = v1
                                                                    v4 = v4
                                                                } else {
                                                                    var v37 uint = 274
                                                                    v38 := GenS3{F0: 921, F1: (v4 & v4), F2: 1}
                                                                    v4 = (v4 - v4)
                                                                    _ = v37
                                                                    _ = v38
                                                                }
                                                                v39 := v2["a"]
                                                                _ = fi23
                                                                _ = v24
                                                                _ = v39
                                                            }
                                                        case "foo bar":
                                                            fmt.Println(v3, fi20)
                                                        default:
                                                            v3 += 585
                                                        }
                                                    case -164:
                                                        if true {
                                                            if v1 {
                                                                var v40 int = int(v4)
                                                                _ = v40
                                                            } else {
                                                                v4 -= v4
                                                                v1 = true
                                                                var v41 uint8 = uint8(v3)
                                                                _ = v41
                                                            }
                                                        } else {
                                                            v4 += (v4 / ((v4) | 1))
                                                            var v42 float32 = 1.5
                                                            fmt.Println(v42, v1)
                                                            _ = v42
                                                        }
                                                    case -2147483648:
                                                        var v43 float32 = 1e10
                                                        _ = v43
                                                    default:
                                                        if true {
                                                            v3 = (^(v3))
                                                            v7 = append(v7, fi20)
                                                        }
                                                    }
                                                    var v44 uint8 = 127
                                                    _ = v44
                                                }
                                            case "hello":
                                                if (4294967295 < 18446744073709551615) {
                                                    if (v1 && true) {
                                                        fmt.Println(v4, v1, fi20)
                                                        if (v1 || (!v1)) {
                                                            if true {
                                                                _ = len(v2)
                                                            }
                                                            if ((fi20 / ((9223372036854775807) | 1)) < (fi20 & fi20)) {
                                                                for fi45 := 0; fi45 < 12; fi45++ {
                                                                    var v46 uint32 = uint32(v4)
                                                                    _ = fi45
                                                                    _ = v46
                                                                }
                                                                if ("hello" > (("hello" + "") + "")) {
                                                                    if (1.5 <= 1e10) {
                                                                        v47 := GenS2{F0: fi20, F1: -711, F2: 1}
                                                                        var v48 uint16 = v4
                                                                        _ = v47
                                                                        _ = v48
                                                                    } else {
                                                                        v1 = v1
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            if false {
                                                                                v4 -= v4
                                                                                v49 := GenS5{F0: v4}
                                                                                _ = v49
                                                                            } else {
                                                                                v50 := GenS4{F0: (v3 - v3), F1: -1}
                                                                                _ = v50
                                                                            }
                                                                        case "hello":
                                                                            for fi51 := 0; fi51 < 16; fi51++ {
                                                                                v52 := GenS1{F0: v3}
                                                                                _ = fi51
                                                                                _ = v52
                                                                            }
                                                                        }
                                                                    }
                                                                    if v1 {
                                                                        _ = len(v2)
                                                                        fmt.Println(v3, fi20, v3)
                                                                        for fi53 := 0; fi53 < 7; fi53++ {
                                                                            v54 := GenS6{F0: 397, F1: 210, F2: 339}
                                                                            fmt.Println(v4, fi53, v3)
                                                                            _ = fi53
                                                                            _ = v54
                                                                        }
                                                                    }
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        fmt.Println(v1)
                                                                    case "a":
                                                                        var v55 int = (fi20 + 32767)
                                                                        _ = v55
                                                                    case "hello":
                                                                        v4 += v4
                                                                    default:
                                                                        for fi56 := 0; fi56 < 10; fi56++ {
                                                                            v57 := GenS1{F0: (v3 << ((2) & 63))}
                                                                            var v58 uint = 2
                                                                            v59 := GenS4{F0: (v3 % ((v3) | 1)), F1: 1}
                                                                            _ = fi56
                                                                            _ = v57
                                                                            _ = v58
                                                                            _ = v59
                                                                        }
                                                                    }
                                                                } else {
                                                                    fmt.Println(v4)
                                                                }
                                                            } else {
                                                                v4 += v4
                                                            }
                                                        } else {
                                                            fmt.Println(fi20, v4, v4)
                                                            var v60 int64 = -179
                                                            _ = v60
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                v1 = v1
                                            default:
                                                fmt.Println(fi20, v3, fi20)
                                            }
                                        }
                                        var v61 uint32 = v3
                                        _ = v61
                                    }
                                    _ = fi20
                                }
                            }
                            v2["hello"] = true
                        }
                    }
                }
                _ = v8
            } else {
                var v62 uint64 = 1
                _ = v62
            }
        case "a":
            delete(v2, "a")
        case "\n\t":
            for fi63 := 0; fi63 < 8; fi63++ {
                if false {
                    if false {
                        if v1 {
                            v64 := v2["hello"]
                            if false {
                                var v65 uint = 588
                                _ = v65
                            } else {
                                var v66 int8 = 0
                                fmt.Println(v3, v4)
                                _ = v66
                            }
                            for fi67 := 0; fi67 < 14; fi67++ {
                                v1 = (684 < 2)
                                var v68 int32 = 2147483647
                                var v69 uint8 = 2
                                v68 = (^(v68))
                                if v1 {
                                    var v70 int64 = -595
                                    _ = v70
                                }
                                switch "a" {
                                case "\n\t":
                                    if v64 {
                                        if len(v7) > 0 {
                                            v7[0] = (-(fi63))
                                        }
                                        fmt.Println(v68, v4)
                                    }
                                default:
                                    _ = len(v7)
                                }
                                var v71 uint64 = 2
                                var v72 uint16 = (v4 % ((255) | 1))
                                fmt.Println(v71, v71)
                                _ = fi67
                                _ = v68
                                _ = v69
                                _ = v71
                                _ = v72
                            }
                            _ = v64
                        }
                    } else {
                        for fi73 := 0; fi73 < 20; fi73++ {
                            if v1 {
                                switch (fi73 >> ((2) & 63)) {
                                case 1:
                                    var v74 float32 = 0.0
                                    _ = v74
                                case 0:
                                    delete(v2, "foo bar")
                                default:
                                    v75 := GenS1{F0: (v3 | (v3 ^ v3))}
                                    _ = v75
                                }
                                v4 = 1
                                if v1 {
                                    if true {
                                        fmt.Println(v4)
                                    } else {
                                        var v76 float64 = 0.0
                                        var v77 float64 = v76
                                        _ = v76
                                        _ = v77
                                    }
                                } else {
                                    v78 := GenS4{F0: v3, F1: 1000}
                                    _ = v78
                                }
                                if (2147483647 <= 192) {
                                    var v79 int32 = -2147483648
                                    if v1 {
                                        var v80 bool = false
                                        _ = len(v7)
                                        _ = v80
                                    }
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    _ = v79
                                }
                            } else {
                                v81 := GenS2{F0: fi73, F1: -997, F2: -128}
                                v82 := GenS6{F0: (v3 * v3), F1: 2, F2: -121}
                                _ = v81
                                _ = v82
                            }
                            var v83 uint32 = (v3 % ((v3) | 1))
                            _ = fi73
                            _ = v83
                        }
                    }
                } else {
                    var v84 int8 = -99
                    fmt.Println(v4)
                    _ = v84
                }
                _ = fi63
            }
        default:
            if v1 {
                v85 := GenS7{F0: 218, F1: v4, F2: -192}
                for fi86 := 0; fi86 < 4; fi86++ {
                    v2[("a" + "a")] = (v3 == (v3 | v3))
                    _ = fi86
                }
                var v87 float64 = (((-2.25 + 1.5) - 1.5) * ((3.14159 - float64(0.0)) * 1e10))
                _ = v85
                _ = v87
            }
        }
    } else {
        v1 = (60 <= 126)
    }
    var v88 uint8 = 2
    v88 = 18
    fmt.Println(v3, v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v88
}
