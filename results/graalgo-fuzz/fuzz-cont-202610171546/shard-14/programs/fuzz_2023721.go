package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 int; F2 uint8 }
type GenS7 struct { F0 int8; F1 int16; F2 bool }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 bool; F1 bool; F2 uint16 }
type GenS12 struct { F0 int8; F1 uint }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint8; F1 int16 }
type GenS15 struct { F0 int8; F1 uint16; F2 uint }
type GenS16 struct { F0 uint32; F1 uint8; F2 uint8 }
func genFn0(p0 uint8, p1 int16, p2 bool) int {
    var v5 bool = false
    fmt.Println(p1, p1, p1)
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint8 = p0
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 2147483648}
    var v9 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    _ = v9
    return 507
}

func main() {
    var v1 int16 = 58
    var v2 int8 = 114
    var v3 uint = 865
    var v4 int = 42
    v4 += v4
    var v10 uint32 = 42
    if (!true) {
        v4 -= genFn0(128, 573, true)
        switch "" {
        case "":
            var v11 float32 = 3.14159
            _ = v11
        case "foo bar":
            v12 := GenS2{F0: false}
            _ = v12
        case "hello":
            var v13 uint8 = 2
            _ = v13
        }
    }
    var v14 int8 = (v2 << ((v3) & 63))
    v15 := []int{v4, genFn0(70, 1000, ("\n\t" <= (("hello" + ("hello" + "")) + "\n\t"))), v4, (v4 ^ v4)}
    for fi16 := 0; fi16 < 12; fi16++ {
        v15 = append(v15, v4)
        var v17 float64 = (3.14159 - 0.1)
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 int64 = -212
            var v21 int16 = (v1 >> (((^(v3))) & 63))
            v3 -= 1
            _ = fi19
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 uint64 = 0
            v4 -= (fi18 | (fi18 | genFn0(0, (v1 - (-(v1))), true)))
            if len(v15) > 0 {
                v15[0] = int(v2)
            }
            _ = fi22
            _ = v23
        }
        if (-2.25 <= 1.5) {
            v3 -= v3
            fmt.Println(v3)
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 int8 = -21
                if true {
                    if (2 < 198) {
                        switch ("foo bar" + ("hello" + "foo bar")) {
                        case "foo bar":
                            if (!((1e10 + 0.0) != (3.14159 * (1e10 / 0.1)))) {
                                var v26 int16 = v1
                                fmt.Println(v26, v25, v1)
                                _ = v26
                            }
                        }
                        if (!false) {
                            fmt.Println(v25, v10)
                            var v27 uint8 = 2
                            v14 = (v14 - int8(v10))
                            fmt.Println(v4, v14)
                            var v28 int64 = -1
                            _ = v27
                            _ = v28
                        }
                        v14 -= (v14 << ((914) & 63))
                    } else {
                        fmt.Println(v10)
                        var v29 uint8 = 55
                        _ = v29
                    }
                } else {
                    v30 := GenS1{F0: (v10 | v10)}
                    switch v4 {
                    case 9223372036854775807:
                        if ((1e10 <= (-2.25 / 1.0)) || (!(((v1 << ((0) & 63)) <= 1000) || (false && false)))) {
                            v4 += (fi18 - (v4 - 949))
                            v31 := GenS3{F0: (v3 << ((1) & 63)), F1: 0}
                            switch "" {
                            case "foo bar":
                                var v32 float64 = (1e10 + 3.14159)
                                _ = v32
                            case "a":
                                v3 = 698
                            default:
                                if ((1 > uint64(v3)) && true) {
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                    v33 := GenS1{F0: (^(v10))}
                                    _ = v33
                                } else {
                                    fmt.Println(v14)
                                }
                            }
                            _ = v31
                        } else {
                            v15 = append(v15, genFn0(255, v1, (true || (false || ((!true) && (!true))))))
                            v15 = append(v15, fi18)
                        }
                    default:
                        if true {
                            if (-1 <= (-(v1))) {
                                fmt.Println(fi24)
                                v34 := GenS3{F0: v3, F1: 336}
                                _ = v34
                            } else {
                                var v35 uint8 = uint8(v10)
                                v3 += 65535
                                _ = v35
                            }
                        } else {
                            _ = len(v15)
                        }
                    }
                    switch (-(fi18)) {
                    case 42:
                        v2 += (v14 / ((v25) | 1))
                    case 438:
                        v36 := GenS1{F0: (^(v10))}
                        _ = v36
                    }
                    v10 = (^(v10))
                    switch (("foo bar" + "a") + (("foo bar" + "hello") + ("" + "hello"))) {
                    case "":
                        v37 := GenS1{F0: (v10 << ((v3) & 63))}
                        _ = v37
                    case "a":
                        if (false && (v4 > -128)) {
                            var v38 int16 = v1
                            if true {
                                if true {
                                    v10 = 486
                                    if (866 < 32767) {
                                        fmt.Println(v3)
                                        v10 += v10
                                    } else {
                                        v39 := GenS4{F0: 595}
                                        fmt.Println(v1, v2, v38)
                                        _ = v39
                                    }
                                }
                            } else {
                                v1 -= v38
                                v40 := GenS5{F0: v10}
                                _ = v40
                            }
                            var v41 uint8 = 255
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            _ = v38
                            _ = v41
                        } else {
                            v1 += (v1 ^ -32768)
                            v42 := GenS6{F0: (!true), F1: genFn0(0, (^(v1)), (false || true)), F2: 2}
                            switch fi18 {
                            case 2147483647:
                                var v43 int = fi24
                                _ = v43
                            case -2147483648:
                                if (!true) {
                                    if (42 <= 255) {
                                        switch "foo bar" {
                                        case "a":
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(v3, v1)
                                            case "\n\t":
                                                var v44 int16 = (v1 << (((v3 | v3)) & 63))
                                                _ = v44
                                            case "hello":
                                                var v45 int32 = 379
                                                _ = v45
                                            }
                                        case "hello":
                                            v1 = 208
                                        }
                                    }
                                    v15 = append(v15, (v4 + (v4 + (fi24 % ((int(v1)) | 1)))))
                                } else {
                                    if (1 == 0) {
                                        switch (-(fi18)) {
                                        case -886:
                                            v46 := GenS7{F0: (v2 | (^(v2))), F1: -369, F2: true}
                                            _ = v46
                                        case 701:
                                            if len(v15) > 0 {
                                                v15[0] = genFn0(42, 32767, (-1 < v1))
                                            }
                                        default:
                                            fmt.Println(v1)
                                        }
                                        v1 += (v1 & v1)
                                    }
                                }
                            default:
                                var v47 uint64 = 4294967295
                                _ = v47
                            }
                            _ = v42
                        }
                    case "hello":
                        _ = len(v15)
                    default:
                        v1 = (v1 | v1)
                    }
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        v15 = append(v15, fi18)
                    case "\n\t":
                        if true {
                            var v48 string = "\n\t"
                            fmt.Println(v25, v3, v14)
                            if ("hello" <= "") {
                                switch (v48 + "foo bar") {
                                case "hello":
                                    v48 = ((v48 + ((v48 + ("hello" + (v48 + v48))) + "")) + (v48 + ("foo bar" + v48)))
                                    var v49 int64 = 42
                                    _ = v49
                                case "a":
                                    var v50 int32 = -2147483648
                                    _ = v50
                                case "\n\t":
                                    v51 := GenS2{F0: false}
                                    _ = v51
                                default:
                                    _ = len(v15)
                                }
                            } else {
                                fmt.Println(v4, v2)
                                fmt.Println(v10, v3)
                            }
                            fmt.Println(v2, v14, v48)
                            _ = v48
                        } else {
                            if true {
                                var v52 float32 = 1.5
                                if false {
                                    if ((-9223372036854775808 >= 78) && ((!false) && false)) {
                                        if len(v15) > 0 {
                                            v15[0] = genFn0(134, v1, ((-783 < 243) || false))
                                        }
                                    } else {
                                        if ("hello" <= ("" + "hello")) {
                                            var v53 uint = (v3 >> (((v3 / ((v3) | 1))) & 63))
                                            _ = v53
                                        }
                                        v54 := GenS8{F0: 0}
                                        _ = v54
                                    }
                                    if false {
                                        v3 = 0
                                        v55 := GenS9{F0: 0}
                                        _ = v55
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        if false {
                                            if (((("" + (("" + "foo bar") + "hello")) + "\n\t") == "foo bar") && (32767 == 38)) {
                                                v56 := GenS10{F0: -1, F1: (v10 / (((v10 % (((v10 - uint32(v3))) | 1))) | 1))}
                                                _ = v56
                                            } else {
                                                switch "" {
                                                case "a":
                                                    var v57 float32 = 3.14159
                                                    _ = v57
                                                case "foo bar":
                                                    if len(v15) >= 2 {
                                                        v15 = v15[1:len(v15)]
                                                    }
                                                case "":
                                                    v58 := GenS11{F0: true, F1: (v10 >= (v10 & (v10 | 4294967295))), F2: 108}
                                                    _ = v58
                                                }
                                            }
                                            var v59 int = v4
                                            _ = v59
                                        } else {
                                            if len(v15) > 0 {
                                                v15[0] = v4
                                            }
                                            fmt.Println(fi18, v52)
                                        }
                                    case "a":
                                        v60 := GenS12{F0: v25, F1: (v3 * v3)}
                                        _ = v60
                                    default:
                                        switch -948 {
                                        case 269:
                                            var v61 bool = true
                                            _ = v61
                                        }
                                    }
                                }
                                _ = v52
                            }
                            var v62 float32 = 0.1
                            var v63 int16 = v1
                            _ = v62
                            _ = v63
                        }
                    }
                    _ = v30
                }
                v14 = (v2 % ((v2) | 1))
                _ = fi24
                _ = v25
            }
        }
        _ = fi18
    }
    fmt.Println(v4, v10)
    var v64 uint16 = 42
    var v65 int32 = 0
    fmt.Println(v4, v2, v14)
    if true {
        _ = len(v15)
    } else {
        switch (v4 & (v4 + v4)) {
        case 128:
            var v66 bool = (true && false)
            _ = v66
        case -973:
            var v67 uint32 = uint32(v64)
            _ = v67
        }
        var v68 int64 = 111
        if true {
            fmt.Println(v3, v4, v68)
        } else {
            for fi69 := 0; fi69 < 18; fi69++ {
                if (!("hello" <= "\n\t")) {
                    v15 = append(v15, v4)
                    v65 = (v65 & v65)
                }
                var v70 uint = (v3 / ((v3) | 1))
                switch (fi69 & (fi69 + (^(fi69)))) {
                case -2147483648:
                    var v71 int64 = (v68 ^ v68)
                    _ = v71
                }
                _ = fi69
                _ = v70
            }
        }
        _ = v68
    }
    v72 := GenS12{F0: v14, F1: (v3 >> (((v3 | 18446744073709551615)) & 63))}
    for fi73 := 0; fi73 < 11; fi73++ {
        v65 = v65
        v10 += v10
        _ = fi73
    }
    for fi74 := 0; fi74 < 18; fi74++ {
        if (v1 >= int16(v3)) {
            if false {
                v64 -= (v64 << (((v3 * v3)) & 63))
            }
            for fi75 := 0; fi75 < 18; fi75++ {
                v76 := GenS8{F0: (-(v65))}
                v15 = append(v15, (fi74 | (v4 / ((fi75) | 1))))
                v1 = (^(v1))
                _ = fi75
                _ = v76
            }
        }
        v14 = v2
        var v77 int8 = int8(v64)
        _ = fi74
        _ = v77
    }
    var v78 uint32 = (v10 >> (((v3 + (^(v3)))) & 63))
    for fi79 := 0; fi79 < 7; fi79++ {
        _ = len(v15)
        _ = fi79
    }
    if ((v3 & uint(v64)) == v3) {
        for fi80 := 0; fi80 < 19; fi80++ {
            switch ("a" + "foo bar") {
            case "\n\t":
                v81 := GenS13{F0: (^(v65))}
                _ = v81
            case "":
                for fi82 := 0; fi82 < 15; fi82++ {
                    var v83 int16 = v1
                    v64 = v64
                    if false {
                        v78 = v78
                        if true {
                            if true {
                                v84 := GenS2{F0: (1 >= v64)}
                                v64 += v64
                                _ = len(v15)
                                _ = v84
                            }
                            var v85 uint8 = 128
                            _ = v85
                        } else {
                            if ((v64 | v64) <= (v64 * 0)) {
                                v86 := GenS14{F0: 168, F1: -32768}
                                var v87 float64 = float64((((3.14159 / 1.5) + 1e10) + 0.1))
                                v88 := GenS15{F0: (v14 * -1), F1: v64, F2: v3}
                                _ = v86
                                _ = v87
                                _ = v88
                            }
                        }
                        v89 := GenS16{F0: (^(v78)), F1: 255, F2: 127}
                        _ = v89
                    }
                    v14 += 1
                    _ = fi82
                    _ = v83
                }
            case "foo bar":
                v10 = (v10 << ((uint(v64)) & 63))
            }
            _ = fi80
        }
        v14 += (v2 ^ int8(v10))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
    _ = v72
    _ = v78
}
