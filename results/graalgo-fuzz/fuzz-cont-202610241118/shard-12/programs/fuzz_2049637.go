package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int64; F2 uint16 }
type GenS4 struct { F0 int; F1 bool; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 int64; F2 uint16 }
type GenS7 struct { F0 bool; F1 string; F2 uint8 }
type GenS8 struct { F0 uint8; F1 int64 }
type GenS9 struct { F0 uint64; F1 int16 }
type GenS10 struct { F0 string; F1 int }
type GenS11 struct { F0 uint32; F1 uint64; F2 uint }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int16; F1 uint8 }
type GenS14 struct { F0 int64; F1 int8; F2 uint32 }
func genFn0(p0 uint, p1 bool, p2 int64) int {
    p0 = (p0 + p0)
    if (65535 <= 0) {
        var v4 int = 521
        var v5 uint8 = 0
        _ = v4
        _ = v5
    } else {
        if false {
            var v6 int = -1
            v7 := GenS1{F0: 1000}
            fmt.Println(p1)
            _ = v6
            _ = v7
        } else {
            var v8 uint8 = 255
            _ = v8
        }
        v9 := []int{42, 9223372036854775807}
        v10 := GenS2{F0: 173}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS2{F0: 2}
        switch "foo bar" {
        case "hello":
            p0 = (^(p0))
        case "foo bar":
            for fi13 := 0; fi13 < 5; fi13++ {
                fmt.Println(p2)
                v14 := []int{}
                if len(v14) > 0 {
                    v14[0] = (fi13 - fi13)
                }
                _ = fi13
                _ = v14
            }
        case "\n\t":
            p1 = p1
        }
        _ = fi11
        _ = v12
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            fmt.Println(fi15, p0, fi16)
            _ = fi16
        }
        var v17 float64 = (((3.14159 / 1.0) - (3.14159 / 1e10)) + 3.14159)
        _ = fi15
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    v1 := []uint32{657, 894, 0, 4294967295}
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    v1 = append(v1, 23)
    for fi18 := 0; fi18 < 9; fi18++ {
        v19 := GenS3{F0: false, F1: 1, F2: 1}
        switch "" {
        case "":
            var v20 bool = true
            _ = v20
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = 65535
            }
        default:
            fmt.Println(fi18)
        }
        _ = fi18
        _ = v19
    }
    _ = len(v3)
    _ = len(v1)
    v1 = append(v1, 0)
    if true {
        v21 := GenS4{F0: 9223372036854775807, F1: false, F2: 1}
        _ = v21
    } else {
        var v22 uint32 = 2147483647
        if true {
            if true {
                switch 128 {
                case 128:
                    v22 = 2147483648
                case 32767:
                    var v23 int32 = 784
                    _ = v23
                case 2147483647:
                    v22 -= 901
                }
                v3[1] = ("hello" + "foo bar")
                if (!false) {
                    v24 := GenS5{F0: 195}
                    var v25 int32 = 1
                    var v26 uint16 = 243
                    _ = v24
                    _ = v25
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 9; fi27++ {
                        for fi28 := 0; fi28 < 2; fi28++ {
                            v29 := GenS3{F0: (!(12 != 42)), F1: 42, F2: 133}
                            _ = fi28
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 13; fi30++ {
                            var v31 uint = 265
                            if true {
                                if (18446744073709551615 != 361) {
                                    fmt.Println(v31, fi27)
                                    switch ("foo bar" + "a") {
                                    case "a":
                                        var v32 int = (fi30 | genFn0(v31, true, -9223372036854775808))
                                        _ = v32
                                    case "foo bar":
                                        v31 -= (v31 % (((v31 % ((v31) | 1))) | 1))
                                    case "":
                                        v2[("\n\t" + "foo bar")] = fi27
                                    default:
                                        var v33 float64 = ((float64(3.14159) * float64(float32(-2.25))) / 3.14159)
                                        _ = v33
                                    }
                                    v34 := GenS6{F0: (v31 > 588), F1: 0, F2: 1}
                                    _ = v34
                                }
                                if len(v1) > 0 {
                                    v1[0] = (v22 >> ((v31) & 63))
                                }
                                switch ("\n\t" + ("hello" + "")) {
                                case "\n\t":
                                    v35 := GenS5{F0: 79}
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(fi30, fi27, v31)
                                    }
                                    var v36 int32 = -1
                                    v22 += (^(v22))
                                    v22 = v22
                                    _ = v35
                                    _ = v36
                                case "hello":
                                    switch (fi27 >> ((v31) & 63)) {
                                    case -625:
                                        v37 := GenS7{F0: (!(v31 != (v31 + 315))), F1: ("" + "\n\t"), F2: 127}
                                        _ = v37
                                    case 0:
                                        v1 = append(v1, v22)
                                    }
                                default:
                                    if true {
                                        v22 = 65535
                                        if true {
                                            v31 += v31
                                            var v38 int32 = -2147483648
                                            if (true && false) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v39 uint = 2
                                                v40 := GenS8{F0: 255, F1: int64(v39)}
                                                if (fi27 > fi30) {
                                                    v22 = (v22 ^ v22)
                                                    switch (^(fi27)) {
                                                    case 42:
                                                        v41 := v3[(v22 / ((v22) | 1))]
                                                        _ = v41
                                                    }
                                                }
                                                _ = v39
                                                _ = v40
                                            } else {
                                                switch 2147483647 {
                                                case 107:
                                                    v42 := GenS9{F0: 2, F1: -419}
                                                    _ = v42
                                                case 128:
                                                    v31 = v31
                                                default:
                                                    var v43 int8 = 100
                                                    _ = v43
                                                }
                                                if true {
                                                    var v44 int8 = 127
                                                    _ = v44
                                                } else {
                                                    switch (fi27 ^ (^(fi30))) {
                                                    case -432:
                                                        var v45 float64 = (-2.25 - 0.0)
                                                        _ = v45
                                                    case 68:
                                                        v22 = 630
                                                    default:
                                                        delete(v2, ("hello" + "\n\t"))
                                                    }
                                                    v31 += v31
                                                    if false {
                                                        v46 := GenS4{F0: 42, F1: (1 <= 572), F2: -686}
                                                        _ = v46
                                                    } else {
                                                        var v47 int64 = -787
                                                        _ = v47
                                                    }
                                                }
                                                switch ("foo bar" + "foo bar") {
                                                case "a":
                                                    fmt.Println(fi27)
                                                case "\n\t":
                                                    delete(v2, "")
                                                default:
                                                    v48 := GenS8{F0: 42, F1: -9223372036854775808}
                                                    _ = v48
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v49 int8 = 18
                                                _ = v49
                                            }
                                            _ = v38
                                        } else {
                                            var v50 bool = (-882 != 893)
                                            fmt.Println(fi30)
                                            _ = v50
                                        }
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        v51 := GenS9{F0: 42, F1: -1}
                                        _ = v51
                                    }
                                }
                            } else {
                                fmt.Println(fi30, fi27)
                                if true {
                                    _ = len(v1)
                                }
                                var v52 float32 = 0.0
                                _ = v52
                            }
                            var v53 float32 = -2.25
                            _ = fi30
                            _ = v31
                            _ = v53
                        }
                        v22 -= (v22 ^ (v22 & (v22 * (v22 >> ((42) & 63)))))
                        _ = fi27
                    }
                }
            } else {
                v22 = v22
            }
            fmt.Println(v22)
        } else {
            if len(v1) > 0 {
                v1[0] = (v22 & (v22 % (((v22 + (^(v22)))) | 1)))
            }
            for fi54 := 0; fi54 < 2; fi54++ {
                var v55 int16 = 1000
                if (false && (1 >= 1)) {
                    v22 += v22
                    if true {
                        var v56 int64 = 1
                        _ = v56
                    }
                    if false {
                        switch "" {
                        case "\n\t":
                            delete(v3, v22)
                        case "a":
                            for fi57 := 0; fi57 < 6; fi57++ {
                                if (0 >= 0) {
                                    if false {
                                        v55 = (v55 / ((v55) | 1))
                                        var v58 string = ((("a" + "\n\t") + ("hello" + "foo bar")) + ("\n\t" + "\n\t"))
                                        _ = v58
                                    }
                                    v22 = 4294967295
                                }
                                _ = fi57
                            }
                        case "":
                            fmt.Println(v55, fi54, fi54)
                        default:
                            fmt.Println(v22)
                        }
                        _ = len(v3)
                    }
                } else {
                    fmt.Println(v55, v22)
                    v1 = append(v1, (v22 | 347))
                    switch genFn0(18446744073709551615, false, -403) {
                    case -9223372036854775808:
                        switch (fi54 << ((18446744073709551615) & 63)) {
                        case 9223372036854775807:
                            if (true && (true || false)) {
                                v22 -= v22
                                for fi59 := 0; fi59 < 9; fi59++ {
                                    var v60 int8 = -128
                                    var v61 int32 = -1
                                    _ = fi59
                                    _ = v60
                                    _ = v61
                                }
                            }
                        case 42:
                            v22 -= v22
                        default:
                            v62 := GenS10{F0: "", F1: fi54}
                            _ = v62
                        }
                    case -561:
                        for fi63 := 0; fi63 < 8; fi63++ {
                            var v64 int8 = 15
                            _ = fi63
                            _ = v64
                        }
                    case -365:
                        v22 = v22
                    }
                }
                _ = fi54
                _ = v55
            }
        }
        _ = v22
    }
    var v65 uint8 = 0
    v65 = v65
    fmt.Println(v65)
    if true {
        v66 := GenS11{F0: 2147483648, F1: 511, F2: 720}
        if (false || (!false)) {
            _ = len(v1)
            switch genFn0(0, false, 9223372036854775807) {
            case 32767:
                v67 := GenS12{F0: 92}
                _ = v67
            case 4294967295:
                if false {
                    v68 := GenS13{F0: 1000, F1: (v65 >> ((65535) & 63))}
                    v69 := GenS3{F0: (("\n\t" + "\n\t") != "\n\t"), F1: 0, F2: 42}
                    _ = v68
                    _ = v69
                }
            default:
                switch 32768 {
                case -426:
                    v3[2] = ""
                default:
                    for fi70 := 0; fi70 < 15; fi70++ {
                        var v71 int = 42
                        var v72 float32 = 3.14159
                        _ = fi70
                        _ = v71
                        _ = v72
                    }
                }
            }
            var v73 uint32 = 701
            _ = v73
        } else {
            for fi74 := 0; fi74 < 6; fi74++ {
                switch fi74 {
                case 128:
                    switch "a" {
                    case "\n\t":
                        v65 -= v65
                    case "":
                        v65 += (v65 * v65)
                    default:
                        if true {
                            v65 += 178
                            delete(v3, 42)
                            v75 := GenS14{F0: 0, F1: 1, F2: 2147483647}
                            _ = v75
                        }
                    }
                case 32767:
                    v65 = (v65 + (v65 / ((v65) | 1)))
                case 908:
                    for fi76 := 0; fi76 < 5; fi76++ {
                        v77 := v2["a"]
                        var v78 string = "a"
                        _ = fi76
                        _ = v77
                        _ = v78
                    }
                default:
                    if len(v1) > 0 {
                        v1[0] = 784
                    }
                }
                _ = fi74
            }
        }
        for fi79 := 0; fi79 < 18; fi79++ {
            v65 += 128
            v65 -= (v65 >> ((753) & 63))
            _ = fi79
        }
        _ = v66
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
}
