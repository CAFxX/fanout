package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int16; F1 int8; F2 uint64 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint32 }

func main() {
    var v1 string = (("" + "hello") + "a")
    v2 := []int{-575, 32768, 139}
    v3 := []string{(v1 + "\n\t"), "", v1}
    var v4 int8 = -58
    v5 := GenS1{F0: 42, F1: (-1 != 1)}
    if false {
        switch v1 {
        case "\n\t":
            var v6 uint64 = 0
            _ = v6
        }
        v4 -= (v4 & v4)
    } else {
        if (0 > 42) {
            var v7 string = "hello"
            var v8 int16 = 1
            _ = v7
            _ = v8
        } else {
            fmt.Println(v4, v1)
            v4 += (v4 | v4)
        }
        if (468 == 2) {
            v9 := GenS1{F0: 127, F1: ((v4 >> ((243) & 63)) <= (v4 - (v4 & (v4 * 1))))}
            switch -9223372036854775808 {
            case 127:
                v10 := GenS2{F0: -32768, F1: 0, F2: 42}
                _ = v10
            case -2147483648:
                var v11 uint8 = 127
                _ = v11
            case -289:
                _ = len(v3)
            default:
                v12 := GenS3{F0: 1000, F1: 127, F2: 0}
                _ = v12
            }
            fmt.Println(v4)
            _ = v9
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v3)
    switch v1 {
    case "a":
        switch v1 {
        case "a":
            var v13 int16 = 1
            _ = v13
        case "foo bar":
            v14 := GenS1{F0: 1, F1: ((v4 & v4) >= v4)}
            _ = v14
        case "\n\t":
            v3 = append(v3, (v1 + v1))
        }
    case "foo bar":
        v1 = "\n\t"
    case "hello":
        v4 = v4
    default:
        v15 := GenS4{F0: v1, F1: -9223372036854775808}
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            _ = len(v2)
            var v18 uint64 = 106
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v1 = v1
    if (3.14159 <= 3.14159) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        switch "foo bar" {
        case "hello":
            fmt.Println(v4)
        default:
            switch -299 {
            case 997:
                for fi19 := 0; fi19 < 5; fi19++ {
                    fmt.Println(fi19)
                    fmt.Println(fi19, v1)
                    for fi20 := 0; fi20 < 11; fi20++ {
                        fmt.Println(v1)
                        switch 127 {
                        case 32768:
                            v21 := GenS2{F0: -24, F1: 193, F2: 4294967295}
                            _ = v21
                        case -780:
                            var v22 int8 = v4
                            _ = v22
                        case 9223372036854775807:
                            var v23 uint = 255
                            _ = v23
                        default:
                            _ = len(v2)
                        }
                        v24 := GenS1{F0: (fi20 + (fi19 / ((fi20) | 1))), F1: false}
                        _ = fi20
                        _ = v24
                    }
                    _ = fi19
                }
            default:
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 int = (fi25 >> ((42) & 63))
                    if true {
                        var v27 string = v1
                        var v28 uint32 = 642
                        _ = v27
                        _ = v28
                    }
                    v26 = (fi25 ^ (v26 % (((fi25 % ((v26) | 1))) | 1)))
                    _ = fi25
                    _ = v26
                }
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        if false {
            v3 = append(v3, "a")
            v4 -= v4
            if ((0.1 / 0.1) <= (((1e10 - 1e10) - -2.25) / 1.0)) {
                v4 -= v4
                fmt.Println(v4, v4)
            } else {
                switch 32767 {
                case 42:
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                case 876:
                    switch -136 {
                    case 4294967295:
                        var v29 int8 = (v4 + v4)
                        _ = v29
                    case -9223372036854775808:
                        switch -1 {
                        case 1:
                            v30 := GenS4{F0: v1, F1: 381}
                            _ = v30
                        default:
                            v2 = append(v2, 32767)
                        }
                    case 785:
                        var v31 float64 = -2.25
                        _ = v31
                    }
                case 128:
                    var v32 float64 = (0.1 * 3.14159)
                    _ = v32
                default:
                    if (-276 == -768) {
                        var v33 bool = (!(-245 != -1))
                        _ = v33
                    } else {
                        var v34 bool = true
                        _ = len(v2)
                        _ = v34
                    }
                }
                if (-32768 > 1) {
                    v35 := GenS5{F0: 1}
                    switch ((v1 + v1) + ("a" + "foo bar")) {
                    case "hello":
                        switch 127 {
                        case 903:
                            v4 += (v4 >> ((2) & 63))
                        case 0:
                            if len(v2) > 0 {
                                v2[0] = 32768
                            }
                        default:
                            for fi36 := 0; fi36 < 2; fi36++ {
                                v4 = (v4 + v4)
                                v37 := GenS5{F0: -32768}
                                if (32768 <= fi36) {
                                    var v38 int64 = 9223372036854775807
                                    var v39 float64 = float64((1e10 * 0.1))
                                    _ = v38
                                    _ = v39
                                } else {
                                    _ = len(v3)
                                    v40 := GenS6{F0: -624, F1: false, F2: 4294967295}
                                    v41 := GenS5{F0: -165}
                                    _ = v40
                                    _ = v41
                                }
                                _ = fi36
                                _ = v37
                            }
                        }
                    case "":
                        var v42 float32 = (((0.0 + 0.0) * 1e10) * (1.5 - 0.1))
                        _ = v42
                    default:
                        for fi43 := 0; fi43 < 13; fi43++ {
                            fmt.Println(v4, fi43, fi43)
                            for fi44 := 0; fi44 < 16; fi44++ {
                                var v45 string = v1
                                fmt.Println(v4, v1, fi43)
                                _ = fi44
                                _ = v45
                            }
                            _ = fi43
                        }
                    }
                    var v46 bool = (-2147483648 > -2147483648)
                    _ = v35
                    _ = v46
                } else {
                    for fi47 := 0; fi47 < 14; fi47++ {
                        v1 = "a"
                        v48 := GenS7{F0: 128, F1: 4294967295, F2: 1}
                        v4 = (^(v4))
                        _ = fi47
                        _ = v48
                    }
                    fmt.Println(v4, v4)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = len(v2)
            }
        }
        _ = len(v3)
    } else {
        v4 += 0
        var v49 int32 = 1
        _ = v49
    }
    switch v1 {
    case "\n\t":
        var v50 bool = false
        _ = v50
    case "":
        if (2147483647 <= 1) {
            var v51 int = -341
            v51 -= v51
            for fi52 := 0; fi52 < 5; fi52++ {
                fmt.Println(v1, v4, v1)
                _ = fi52
            }
            _ = v51
        }
    case "a":
        var v53 int64 = 1
        _ = v53
    default:
        var v54 uint32 = 2147483647
        _ = v54
    }
    v4 -= v4
    var v55 int64 = -691
    for fi56 := 0; fi56 < 4; fi56++ {
        if (!(v1 == (("foo bar" + ("" + v1)) + v1))) {
            fmt.Println(v4)
            var v57 int = fi56
            _ = v57
        }
        fmt.Println(v4)
        v3 = append(v3, v1)
        _ = fi56
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
}
