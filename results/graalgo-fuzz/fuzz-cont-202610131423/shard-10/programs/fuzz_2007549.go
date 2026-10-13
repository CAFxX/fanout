package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 int; F2 int8 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 int32; F1 int32 }

func main() {
    var v1 bool = true
    v2 := []string{}
    var v3 int = 32768
    if ((0 < -63) && v1) {
        fmt.Println(v1)
        v4 := GenS1{F0: 987, F1: 0, F2: 2147483647}
        _ = v4
    }
    v1 = (("\n\t" <= "\n\t") || false)
    var v5 uint = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6 := GenS1{F0: 1, F1: 0, F2: 4294967295}
    v1 = (-1 <= 0)
    var v7 int = (^(v3))
    if ((v5 % ((42) | 1)) <= (v5 % (((^(v5))) | 1))) {
        v7 += (^(v7))
        fmt.Println(v7, v1, v1)
        switch v3 {
        case 2147483647:
            var v8 uint8 = 2
            _ = v8
        case 4294967295:
            for fi9 := 0; fi9 < 12; fi9++ {
                for fi10 := 0; fi10 < 4; fi10++ {
                    var v11 int16 = -958
                    v1 = false
                    if false {
                        v12 := GenS2{F0: "\n\t", F1: -212}
                        if v1 {
                            var v13 uint16 = 8
                            v13 -= (v13 * uint16(v5))
                            var v14 int64 = -9223372036854775808
                            v15 := GenS3{F0: -163}
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                        if ("" != "\n\t") {
                            v16 := GenS4{F0: false, F1: (fi9 & v7), F2: int8(v11)}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v3 = fi9
                            _ = v16
                        }
                        _ = v12
                    }
                    switch "\n\t" {
                    case "\n\t":
                        switch (v3 % ((v3) | 1)) {
                        case -528:
                            switch ("" + "foo bar") {
                            case "":
                                var v17 int8 = 42
                                v18 := GenS4{F0: (uint32(v5) != 42), F1: v3, F2: v17}
                                v17 = v17
                                _ = v17
                                _ = v18
                            }
                        }
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
        }
    }
    fmt.Println(v7, v3)
    switch "\n\t" {
    case "hello":
        if (-471 == 375) {
            var v19 uint16 = 0
            var v20 uint8 = 128
            _ = v19
            _ = v20
        }
    default:
        v5 -= 42
    }
    if len(v2) > 0 {
        v2[0] = ("foo bar" + "a")
    }
    if v1 {
        if len(v2) > 0 {
            v2[0] = "a"
        }
    } else {
        v1 = v1
    }
    v21 := GenS5{F0: 456, F1: -1}
    if (42 != -818) {
        v22 := GenS4{F0: (-458 > -1), F1: (v7 + (^(v3))), F2: -40}
        v7 = -91
        if v1 {
            var v23 int32 = 761
            v5 += 4294967295
            fmt.Println(v1, v5, v3)
            _ = v23
        } else {
            v24 := GenS4{F0: true, F1: (v7 & (v7 | v3)), F2: -105}
            _ = v24
        }
        _ = v22
    } else {
        v7 += (-(v3))
        v25 := GenS2{F0: "\n\t", F1: 42}
        for fi26 := 0; fi26 < 17; fi26++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v27 int32 = 42
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    switch v7 {
    case 128:
        v28 := GenS4{F0: (32767 > 2), F1: -861, F2: -121}
        _ = v28
    case -2147483648:
        var v29 uint32 = 2147483648
        _ = v29
    case 127:
        if (4294967295 != 2147483648) {
            fmt.Println(v3, v7, v1)
            v7 = (v7 * v3)
            _ = len(v2)
        } else {
            v2 = append(v2, "\n\t")
        }
    default:
        v7 -= v3
    }
    switch v3 {
    case -2147483648:
        var v30 uint32 = uint32(v5)
        _ = v30
    case 42:
        v7 = (v3 | (^(v3)))
    case -970:
        v5 += v5
    default:
        if v1 {
            _ = len(v2)
            v31 := GenS5{F0: int64(v7), F1: 0}
            _ = v31
        } else {
            if v1 {
                var v32 float32 = float32(0.0)
                _ = v32
            } else {
                if ("a" != "\n\t") {
                    var v33 bool = (!((0 > 32767) || v1))
                    _ = v33
                } else {
                    for fi34 := 0; fi34 < 16; fi34++ {
                        v35 := GenS4{F0: false, F1: -1, F2: 100}
                        for fi36 := 0; fi36 < 11; fi36++ {
                            v7 -= (fi36 ^ (fi36 & int(v5)))
                            v37 := GenS5{F0: -1, F1: 9223372036854775807}
                            fmt.Println(fi36)
                            _ = fi36
                            _ = v37
                        }
                        var v38 bool = (127 >= -1)
                        _ = fi34
                        _ = v35
                        _ = v38
                    }
                    v7 = -236
                    v39 := GenS6{F0: 1, F1: 128}
                    _ = v39
                }
                if (2 < 0) {
                    for fi40 := 0; fi40 < 8; fi40++ {
                        if false {
                            v41 := GenS7{F0: fi40, F1: fi40}
                            if ((!v1) || ((!v1) || (!((-902 > -1) && (!(int16(v5) == -1)))))) {
                                var v42 uint16 = 882
                                _ = v42
                            } else {
                                fmt.Println(v1)
                                switch (fi40 & (v7 % ((-2147483648) | 1))) {
                                case 32767:
                                    var v43 float64 = float64((-2.25 * 1.5))
                                    _ = v43
                                case 128:
                                    var v44 bool = ((false && false) || (!(!(!(!(!false))))))
                                    _ = v44
                                case 902:
                                    _ = len(v2)
                                }
                            }
                            var v45 int8 = -106
                            _ = v41
                            _ = v45
                        }
                        var v46 uint32 = 657
                        _ = fi40
                        _ = v46
                    }
                    v47 := GenS6{F0: 1, F1: 1}
                    var v48 uint16 = 783
                    _ = v47
                    _ = v48
                } else {
                    if ((v7 != 486) && v1) {
                        v7 = v7
                    } else {
                        v3 += v7
                        var v49 float64 = ((1.5 * (0.1 * ((0.0 - 0.1) * ((((0.0 / 3.14159) / (1e10 + 0.1)) * 0.0) / 1.0)))) / 1.5)
                        _ = v49
                    }
                    if true {
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                        v50 := GenS2{F0: ("" + "\n\t"), F1: -1}
                        _ = v50
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            for fi51 := 0; fi51 < 19; fi51++ {
                v3 -= v7
                _ = fi51
            }
            v52 := GenS8{F0: -2147483648, F1: -2147483648}
            _ = v52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
