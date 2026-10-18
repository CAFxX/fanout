package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int64 }
type GenS2 struct { F0 bool; F1 int; F2 int }
type GenS3 struct { F0 int; F1 int8; F2 uint64 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 uint8; F1 uint8 }
type GenS9 struct { F0 int16; F1 uint16 }
type GenS10 struct { F0 uint }

func main() {
    var v1 int16 = -890
    var v2 uint = 168
    var v3 uint16 = 191
    for fi4 := 0; fi4 < 7; fi4++ {
        if false {
            v5 := []int{-1, (fi4 + 9223372036854775807), fi4}
            _ = v5
        } else {
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 uint8 = 144
                v1 = v1
                _ = fi6
                _ = v7
            }
        }
        _ = fi4
    }
    v1 = v1
    v3 -= 253
    var v8 int = 45
    v9 := []int{}
    v3 = v3
    v2 = 255
    for fi10 := 0; fi10 < 19; fi10++ {
        switch v8 {
        case 128:
            fmt.Println(v3)
        case 4294967295:
            v3 = v3
        case -693:
            if true {
                for fi11 := 0; fi11 < 6; fi11++ {
                    _ = len(v9)
                    fmt.Println(fi10)
                    fmt.Println(fi10, v3, fi11)
                    _ = fi11
                }
            } else {
                for fi12 := 0; fi12 < 3; fi12++ {
                    if len(v9) > 0 {
                        v9[0] = -237
                    }
                    v8 -= fi12
                    v13 := GenS1{F0: 2147483647, F1: 1, F2: 1}
                    var v14 uint = 4294967295
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
        default:
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 uint64 = 0
                if (-1 < -929) {
                    v9 = append(v9, int(v3))
                    if true {
                        if len(v9) > 0 {
                            v9[0] = (fi10 / ((42) | 1))
                        }
                    }
                }
                var v17 string = "\n\t"
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        var v18 float64 = 3.14159
        if ((false && true) && false) {
            var v19 int = 2147483647
            _ = v19
        } else {
            if (!false) {
                v9 = append(v9, (^(v8)))
                v20 := GenS1{F0: 1, F1: 2, F2: 0}
                _ = v20
            }
            if true {
                v8 -= (-(fi10))
            } else {
                v1 = (v1 / ((v1) | 1))
                _ = len(v9)
            }
        }
        _ = fi10
        _ = v18
    }
    switch ("a" + "") {
    case "a":
        v2 = 42
    case "foo bar":
        fmt.Println(v8)
    }
    if false {
        switch "\n\t" {
        case "":
            var v21 uint32 = 628
            _ = v21
        case "\n\t":
            fmt.Println(v1)
        case "hello":
            if (42 != 966) {
                v1 += (v1 | v1)
                if len(v9) > 0 {
                    v9[0] = (v8 & v8)
                }
                var v22 float64 = float64((0.0 / 3.14159))
                _ = v22
            } else {
                for fi23 := 0; fi23 < 6; fi23++ {
                    if true {
                        fmt.Println(v3, v1, v3)
                        v2 = v2
                    } else {
                        if (!false) {
                            v24 := GenS2{F0: ((v8 - fi23) == (v8 - fi23)), F1: -9223372036854775808, F2: (v8 | (v8 - (fi23 + (-(fi23)))))}
                            fmt.Println(v2)
                            v1 = v1
                            _ = v24
                        }
                    }
                    switch "\n\t" {
                    case "hello":
                        v1 = v1
                    case "\n\t":
                        if len(v9) > 0 {
                            v9[0] = (fi23 - fi23)
                        }
                    case "a":
                        v25 := GenS1{F0: 0, F1: 2, F2: 9223372036854775807}
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 1; fi26++ {
                        switch ("\n\t" + "") {
                        case "foo bar":
                            var v27 string = ""
                            _ = v27
                        case "a":
                            var v28 int16 = v1
                            _ = v28
                        case "\n\t":
                            v2 += 987
                        default:
                            var v29 uint = (^(v2))
                            _ = v29
                        }
                        _ = fi26
                    }
                    _ = fi23
                }
                if true {
                    v30 := GenS3{F0: (v8 - v8), F1: -12, F2: 18446744073709551615}
                    switch "foo bar" {
                    case "a":
                        for fi31 := 0; fi31 < 13; fi31++ {
                            v32 := GenS4{F0: (v1 / ((v1) | 1)), F1: "\n\t"}
                            var v33 int16 = v1
                            fmt.Println(v3, v3, fi31)
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    case "\n\t":
                        v1 = v1
                    default:
                        v2 = 42
                    }
                    var v34 uint16 = (v3 << (((v2 * (v2 & 2))) & 63))
                    _ = v30
                    _ = v34
                }
                var v35 int = (v8 | 311)
                _ = v35
            }
        default:
            _ = len(v9)
        }
    } else {
        if false {
            if (0 < 42) {
                v3 += uint16(v2)
                for fi36 := 0; fi36 < 3; fi36++ {
                    var v37 int64 = 143
                    _ = fi36
                    _ = v37
                }
            } else {
                v2 += (v2 + 0)
                if true {
                    switch v8 {
                    case -277:
                        var v38 int16 = 268
                        _ = v38
                    }
                    v9 = append(v9, -1)
                } else {
                    v9 = append(v9, v8)
                }
                if true {
                    for fi39 := 0; fi39 < 6; fi39++ {
                        v3 -= (^(v3))
                        if true {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        } else {
                            var v40 int16 = (^(v1))
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 5; fi41++ {
                            switch int(v1) {
                            case 0:
                                v8 = 105
                            case -1:
                                switch (fi41 * (fi39 ^ -9223372036854775808)) {
                                case 0:
                                    var v42 int = (fi41 + (v8 % ((570) | 1)))
                                    _ = v42
                                case 128:
                                    v43 := GenS4{F0: 968, F1: "foo bar"}
                                    _ = v43
                                default:
                                    v3 -= (v3 % (((v3 / ((v3) | 1))) | 1))
                                }
                            default:
                                v8 = (fi39 ^ fi39)
                            }
                            _ = fi41
                        }
                        _ = fi39
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v44 uint8 = 1
                    _ = v44
                } else {
                    v3 += v3
                    v2 -= (v2 ^ v2)
                }
            }
            v8 = v8
        } else {
            v8 -= (^(v8))
        }
        v45 := GenS5{F0: v2}
        if (1 != -47) {
            var v46 uint8 = 128
            for fi47 := 0; fi47 < 1; fi47++ {
                if true {
                    var v48 int32 = -307
                    if false {
                        if (false && (true || true)) {
                            if true {
                                _ = len(v9)
                                var v49 float32 = 1e10
                                fmt.Println(fi47, v1)
                                _ = v49
                            } else {
                                v3 = v3
                                if true {
                                    fmt.Println(v1, v46)
                                    if true {
                                        var v50 int = (v8 - int(v1))
                                        _ = v50
                                    }
                                    for fi51 := 0; fi51 < 14; fi51++ {
                                        if (421 > 2147483647) {
                                            v52 := GenS6{F0: "a", F1: (v48 << ((v2) & 63))}
                                            _ = len(v9)
                                            _ = v52
                                        } else {
                                            if (uint8(v3) == (^(v46))) {
                                                v2 -= (^(v2))
                                                switch "foo bar" {
                                                case "":
                                                    if (!(true || true)) {
                                                        var v53 float32 = 0.1
                                                        _ = v53
                                                    } else {
                                                        v1 -= v1
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                        _ = len(v9)
                                                        v48 = v48
                                                        v48 -= (v48 | v48)
                                                    }
                                                case "hello":
                                                    v8 += fi51
                                                case "foo bar":
                                                    v1 -= v1
                                                default:
                                                    var v54 int16 = (v1 + 32767)
                                                    if true {
                                                        switch "foo bar" {
                                                        case "":
                                                            v2 = (v2 << ((uint(v3)) & 63))
                                                        case "a":
                                                            v2 = (v2 / (((v2 ^ 317)) | 1))
                                                        case "hello":
                                                            var v55 uint16 = (v3 % ((32767) | 1))
                                                            _ = v55
                                                        }
                                                    }
                                                    fmt.Println(v3, v3)
                                                    _ = v54
                                                }
                                            } else {
                                                var v56 float64 = -2.25
                                                _ = v56
                                            }
                                        }
                                        if len(v9) > 0 {
                                            v9[0] = fi47
                                        }
                                        _ = fi51
                                    }
                                } else {
                                    var v57 uint16 = 42
                                    var v58 int16 = v1
                                    var v59 uint8 = v46
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                if ((!true) && true) {
                                    if (!false) {
                                        v60 := GenS7{F0: false, F1: (v48 % (((^(v48))) | 1))}
                                        for fi61 := 0; fi61 < 5; fi61++ {
                                            if true {
                                                var v62 uint = uint(v46)
                                                if (-1 != 539) {
                                                    switch (("\n\t" + ("foo bar" + "foo bar")) + "a") {
                                                    case "\n\t":
                                                        v63 := GenS4{F0: -1, F1: ("foo bar" + "foo bar")}
                                                        _ = v63
                                                    case "foo bar":
                                                        switch "" {
                                                        case "a":
                                                            switch "foo bar" {
                                                            case "hello":
                                                                var v64 int32 = (^(v48))
                                                                _ = v64
                                                            case "a":
                                                                var v65 uint64 = 1
                                                                _ = v65
                                                            }
                                                        }
                                                    default:
                                                        v66 := GenS8{F0: v46, F1: (v46 << ((v62) & 63))}
                                                        _ = v66
                                                    }
                                                    switch int(v2) {
                                                    case 1:
                                                        var v67 uint64 = 0
                                                        _ = v67
                                                    default:
                                                        _ = len(v9)
                                                    }
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                }
                                                _ = v62
                                            }
                                            _ = fi61
                                        }
                                        _ = v60
                                    } else {
                                        v68 := GenS9{F0: v1, F1: v3}
                                        v69 := GenS8{F0: 36, F1: v46}
                                        _ = v68
                                        _ = v69
                                    }
                                    fmt.Println(v46)
                                    var v70 int64 = -1
                                    _ = v70
                                } else {
                                    var v71 uint32 = 65535
                                    _ = v71
                                }
                            }
                            v1 = (v1 >> (((v2 / (((v2 % (((v2 * v2)) | 1))) | 1))) & 63))
                            if true {
                                v2 += v2
                                _ = len(v9)
                                var v72 uint8 = 22
                                _ = v72
                            }
                        } else {
                            if (false || false) {
                                var v73 uint8 = (^(v46))
                                _ = v73
                            }
                            v2 += v2
                            var v74 bool = false
                            _ = v74
                        }
                    } else {
                        v8 -= fi47
                    }
                    for fi75 := 0; fi75 < 8; fi75++ {
                        if len(v9) > 0 {
                            v9[0] = v8
                        }
                        _ = fi75
                    }
                    _ = v48
                } else {
                    v46 -= (v46 - uint8(v3))
                }
                _ = fi47
            }
            _ = v46
        } else {
            if len(v9) > 0 {
                v9[0] = 32767
            }
            v1 = v1
        }
        _ = v45
    }
    v76 := GenS10{F0: (^(v2))}
    var v77 bool = false
    v8 += (v8 - (v8 - 128))
    if len(v9) > 0 {
        v9[0] = 4294967295
    }
    v78 := GenS5{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v77:", v77)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v76
    _ = v77
    _ = v78
    _ = v8
    _ = v9
}
