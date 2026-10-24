package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 bool; F1 uint; F2 int16 }
type GenS3 struct { F0 int8; F1 int16; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 int64; F1 uint }

func main() {
    var v1 float64 = 1e10
    var v2 int64 = 42
    var v3 uint16 = 255
    var v4 uint = 0
    v2 = (v2 & v2)
    v4 = v4
    var v5 int16 = 158
    v5 += -180
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            if (!(v4 < 333)) {
                if false {
                    var v8 uint16 = v3
                    var v9 int32 = 1
                    _ = v8
                    _ = v9
                }
                v3 = 861
            } else {
                var v10 int64 = (v2 + v2)
                var v11 float64 = ((v1 + float64(float32((v1 + 1e10)))) - v1)
                v3 -= (v3 / ((uint16(v4)) | 1))
                _ = v10
                _ = v11
            }
            _ = fi7
        }
        if (0 != 65) {
            v2 = 282
            var v12 uint16 = (v3 >> ((v4) & 63))
            _ = v12
        } else {
            v13 := []int{(fi6 + (-(fi6))), fi6}
            if false {
                var v14 int = fi6
                switch ("\n\t" + "hello") {
                case "hello":
                    var v15 int16 = (-(v5))
                    _ = v15
                }
                _ = v14
            } else {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 string = "\n\t"
                    if ((1.5 == (v1 - 1e10)) && (int16(v3) < -32768)) {
                        if false {
                            var v18 int64 = int64(v5)
                            _ = v18
                        }
                    } else {
                        v13 = append(v13, (fi16 << ((v4) & 63)))
                    }
                    if ((v4 & v4) > v4) {
                        var v19 uint = v4
                        var v20 uint = v4
                        _ = v19
                        _ = v20
                    } else {
                        var v21 int32 = -1
                        v4 = v4
                        var v22 int64 = 0
                        _ = v21
                        _ = v22
                    }
                    _ = fi16
                    _ = v17
                }
                fmt.Println(v3, v4)
            }
            _ = v13
        }
        var v23 string = "hello"
        _ = fi6
        _ = v23
    }
    var v24 int = -2147483648
    switch (v24 + (v24 >> (((v4 / ((1) | 1))) & 63))) {
    case -2147483648:
        var v25 int32 = -456
        _ = v25
    case 42:
        for fi26 := 0; fi26 < 14; fi26++ {
            v27 := GenS1{F0: "hello", F1: (v5 >> ((1) & 63))}
            fmt.Println(v5, v2)
            _ = fi26
            _ = v27
        }
    }
    v28 := []int{v24, v24}
    if false {
        v3 = v3
        v29 := GenS2{F0: true, F1: (v4 ^ 346), F2: v5}
        _ = len(v28)
        _ = v29
    }
    v30 := GenS2{F0: false, F1: (^(v4)), F2: v5}
    var v31 uint64 = 42
    var v32 int32 = -762
    if true {
        var v33 uint64 = v31
        switch "hello" {
        case "\n\t":
            var v34 uint = v4
            _ = v34
        case "":
            if true {
                var v35 uint16 = v3
                _ = v35
            } else {
                if true {
                    for fi36 := 0; fi36 < 18; fi36++ {
                        for fi37 := 0; fi37 < 15; fi37++ {
                            var v38 int64 = v2
                            fmt.Println(v4, v38)
                            if ((232 <= 385) || true) {
                                fmt.Println(v2, fi36, v24)
                                v24 = (fi37 ^ -1)
                                var v39 int8 = 1
                                if true {
                                    var v40 int64 = 42
                                    _ = v40
                                }
                                switch "\n\t" {
                                case "a":
                                    if (uint64(v4) <= 4294967295) {
                                        _ = len(v28)
                                        var v41 float64 = v1
                                        _ = v41
                                    }
                                case "\n\t":
                                    v28 = append(v28, v24)
                                default:
                                    fmt.Println(fi37, v3, v4)
                                }
                                v42 := GenS2{F0: ((v3 ^ v3) == v3), F1: 18446744073709551615, F2: v5}
                                fmt.Println(v39, v32)
                                _ = v39
                                _ = v42
                            } else {
                                if true {
                                    v4 += (v4 ^ 255)
                                    v1 = (v1 + v1)
                                } else {
                                    if len(v28) > 0 {
                                        v28[0] = fi36
                                    }
                                    var v43 uint8 = 42
                                    _ = v43
                                }
                                v44 := GenS1{F0: "foo bar", F1: -32768}
                                v45 := GenS3{F0: 100, F1: -32768, F2: 65535}
                                _ = v44
                                _ = v45
                            }
                            _ = fi37
                            _ = v38
                        }
                        v46 := GenS3{F0: -108, F1: (v5 << (((v4 - (^(v4)))) & 63)), F2: (v4 >> ((v4) & 63))}
                        if false {
                            v5 = (v5 + (-(v5)))
                            for fi47 := 0; fi47 < 20; fi47++ {
                                var v48 int8 = 85
                                if len(v28) > 0 {
                                    v28[0] = (fi47 / ((68) | 1))
                                }
                                var v49 int32 = 42
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                        } else {
                            var v50 bool = (true || (true || ((("hello" < "hello") && (32767 == (^(v3)))) && true)))
                            fmt.Println(v3, v50)
                            _ = v50
                        }
                        _ = fi36
                        _ = v46
                    }
                    if true {
                        if (v4 > v4) {
                            var v51 float32 = 3.14159
                            v2 = (v2 / (((v2 & v2)) | 1))
                            switch 2147483647 {
                            case -172:
                                switch "\n\t" {
                                case "\n\t":
                                    v32 -= -1
                                default:
                                    if len(v28) >= 2 {
                                        v28 = v28[1:len(v28)]
                                    }
                                }
                            case 127:
                                switch "hello" {
                                case "\n\t":
                                    v32 = 42
                                case "":
                                    v33 = 18446744073709551615
                                }
                            case 9223372036854775807:
                                v5 = v5
                            default:
                                var v52 int32 = v32
                                _ = v52
                            }
                            _ = v51
                        } else {
                            for fi53 := 0; fi53 < 2; fi53++ {
                                var v54 uint64 = v33
                                _ = fi53
                                _ = v54
                            }
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                        }
                        switch "a" {
                        case "":
                            for fi55 := 0; fi55 < 7; fi55++ {
                                v2 -= v2
                                _ = fi55
                            }
                        case "hello":
                            v4 = (v4 ^ 356)
                        default:
                            var v56 float32 = (1.5 + (3.14159 + 3.14159))
                            _ = v56
                        }
                    } else {
                        for fi57 := 0; fi57 < 8; fi57++ {
                            fmt.Println(v32, fi57)
                            v5 = v5
                            v58 := GenS1{F0: "a", F1: v5}
                            _ = fi57
                            _ = v58
                        }
                        v59 := GenS4{F0: v3}
                        var v60 int = int(v4)
                        _ = v59
                        _ = v60
                    }
                    if false {
                        v28 = append(v28, v24)
                    } else {
                        var v61 float32 = -2.25
                        _ = v61
                    }
                } else {
                    if (417 >= (v2 | (-(v2)))) {
                        v62 := GenS4{F0: v3}
                        var v63 int64 = (v2 * v2)
                        v64 := GenS5{F0: (-(v32))}
                        _ = v62
                        _ = v63
                        _ = v64
                    } else {
                        fmt.Println(v2, v32, v4)
                        if ((v2 - (v2 ^ (v2 << (((v4 << ((v4) & 63))) & 63)))) != v2) {
                            var v65 int = v24
                            v66 := GenS6{F0: 42}
                            _ = v65
                            _ = v66
                        }
                        if true {
                            if len(v28) > 0 {
                                v28[0] = 127
                            }
                            v24 -= (v24 >> (((v4 % ((v4) | 1))) & 63))
                        }
                    }
                    fmt.Println(v33, v32, v33)
                    if len(v28) > 0 {
                        v28[0] = v24
                    }
                }
            }
        case "a":
            var v67 int64 = (v2 / ((v2) | 1))
            _ = v67
        }
        _ = v33
    }
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    v4 += 265
    if false {
        fmt.Println(v5, v24)
        v68 := GenS3{F0: 42, F1: (v5 - (v5 * 1000)), F2: v4}
        if true {
            if (!true) {
                switch "" {
                case "hello":
                    fmt.Println(v31)
                case "foo bar":
                    var v69 uint8 = 214
                    _ = v69
                }
                var v70 uint64 = (v31 + 0)
                v5 = (v5 + v5)
                _ = v70
            }
            fmt.Println(v1, v4)
        }
        _ = v68
    } else {
        var v71 int8 = 42
        switch ("foo bar" + ("" + "foo bar")) {
        case "foo bar":
            var v72 uint32 = 1
            _ = v72
        case "a":
            v71 += v71
        case "\n\t":
            for fi73 := 0; fi73 < 9; fi73++ {
                _ = len(v28)
                var v74 int16 = (v5 ^ int16(fi73))
                _ = fi73
                _ = v74
            }
        default:
            for fi75 := 0; fi75 < 7; fi75++ {
                v4 -= v4
                var v76 bool = true
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                _ = fi75
                _ = v76
            }
        }
        v77 := GenS7{F0: 387, F1: v71}
        _ = v71
        _ = v77
    }
    var v78 float64 = (v1 / (v1 / v1))
    switch -596 {
    case 32767:
        if (true || false) {
            v79 := GenS2{F0: true, F1: (^(v4)), F2: (v5 / ((v5) | 1))}
            if true {
                if (42 > (v5 << (((v4 % ((v4) | 1))) & 63))) {
                    switch v24 {
                    case 128:
                        var v80 uint = 0
                        _ = v80
                    case -816:
                        v81 := GenS8{F0: 38, F1: ((^(v3)) != (v3 * 37))}
                        _ = v81
                    default:
                        v3 = v3
                    }
                    var v82 float32 = float32(3.14159)
                    _ = v82
                }
            }
            if true {
                if true {
                    var v83 uint = v4
                    v4 = 536
                    _ = v83
                } else {
                    var v84 string = ""
                    var v85 int64 = v2
                    v78 = v1
                    _ = v84
                    _ = v85
                }
                for fi86 := 0; fi86 < 4; fi86++ {
                    var v87 bool = false
                    for fi88 := 0; fi88 < 5; fi88++ {
                        fmt.Println(v32, v32, v3)
                        var v89 int16 = (v5 + (^(v5)))
                        v90 := GenS9{F0: int64(fi88), F1: (v4 >> ((v4) & 63))}
                        _ = fi88
                        _ = v89
                        _ = v90
                    }
                    if false {
                        switch fi86 {
                        case 9223372036854775807:
                            for fi91 := 0; fi91 < 6; fi91++ {
                                fmt.Println(v31, v87)
                                _ = fi91
                            }
                        case 0:
                            var v92 bool = v87
                            _ = v92
                        case -2147483648:
                            v5 = (v5 - v5)
                        default:
                            var v93 uint64 = v31
                            _ = v93
                        }
                        v28 = append(v28, -995)
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                    }
                    _ = fi86
                    _ = v87
                }
                fmt.Println(v5, v3)
            } else {
                v32 = (v32 | v32)
            }
            _ = v79
        }
    case 0:
        var v94 string = "foo bar"
        _ = v94
    case -2147483648:
        v1 = v1
    default:
        v24 += 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v78
}
