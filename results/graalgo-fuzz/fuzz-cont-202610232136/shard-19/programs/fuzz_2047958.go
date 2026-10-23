package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint8; F2 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int }

func main() {
    var v1 uint16 = 0
    var v2 uint8 = uint8(v1)
    var v3 float64 = float64(float32(0.0))
    v4 := []uint32{802, 2147483648}
    v4 = append(v4, 793)
    switch "\n\t" {
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        var v5 float64 = (float64(0.1) - v3)
        _ = v5
    }
    if ((42 != -1) || true) {
        if false {
            v1 += v1
            fmt.Println(v3, v1)
        } else {
            for fi6 := 0; fi6 < 18; fi6++ {
                v2 += v2
                _ = fi6
            }
            if (true || true) {
                if (309 == -2147483648) {
                    v3 = v3
                    switch 9223372036854775807 {
                    case 42:
                        var v7 float32 = (1.5 * (1.5 + 0.0))
                        _ = v7
                    case 32768:
                        if (0 < -578) {
                            var v8 float64 = v3
                            _ = v8
                        } else {
                            if (false || (32767 > -67)) {
                                v3 += v3
                                fmt.Println(v2, v3)
                                v3 = 0.0
                            } else {
                                v4 = append(v4, 0)
                                var v9 int8 = 42
                                _ = v9
                            }
                            var v10 int64 = -9223372036854775808
                            _ = v10
                        }
                    case -352:
                        v11 := GenS1{F0: (-1 != 100)}
                        _ = v11
                    default:
                        var v12 uint8 = (^(v2))
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 19; fi13++ {
                        var v14 int32 = 2147483647
                        _ = fi13
                        _ = v14
                    }
                }
            } else {
                if (255 >= 823) {
                    var v15 float32 = 0.0
                    for fi16 := 0; fi16 < 15; fi16++ {
                        switch ("foo bar" + "\n\t") {
                        case "hello":
                            v3 -= v3
                        case "foo bar":
                            var v17 int8 = -128
                            _ = v17
                        case "\n\t":
                            fmt.Println(v2, fi16)
                        default:
                            if (!(0 <= 558)) {
                                v2 -= (v2 & uint8(v1))
                            } else {
                                if false {
                                    v18 := GenS2{F0: fi16}
                                    if false {
                                        v3 = v3
                                    } else {
                                        for fi19 := 0; fi19 < 4; fi19++ {
                                            v3 += (v3 / v3)
                                            switch 42 {
                                            case 42:
                                                switch fi19 {
                                                case 2147483647:
                                                    if false {
                                                        if (!true) {
                                                            var v20 uint8 = v2
                                                            _ = v20
                                                        }
                                                        var v21 int16 = 42
                                                        if false {
                                                            v2 += 128
                                                        } else {
                                                            fmt.Println(v3, fi19, fi19)
                                                            v22 := GenS1{F0: (false || true)}
                                                            _ = v22
                                                        }
                                                        _ = v21
                                                    }
                                                case 32768:
                                                    v1 = (v1 >> ((0) & 63))
                                                }
                                            }
                                            switch ("foo bar" + (("\n\t" + "a") + "foo bar")) {
                                            case "a":
                                                var v23 uint = 921
                                                _ = v23
                                            case "":
                                                fmt.Println(v15)
                                            case "\n\t":
                                                if false {
                                                    v3 = float64(v15)
                                                    fmt.Println(v3)
                                                    v3 = float64(float32(1.5))
                                                    var v24 uint64 = 2
                                                    if (330 > 18446744073709551615) {
                                                        if false {
                                                            var v25 float32 = 1e10
                                                            v26 := GenS1{F0: ((!false) || ((fi16 | -2147483648) < (^(fi19))))}
                                                            _ = v25
                                                            _ = v26
                                                        } else {
                                                            if ("foo bar" >= ("a" + "a")) {
                                                                var v27 string = ((("\n\t" + "a") + "foo bar") + "\n\t")
                                                                _ = v27
                                                            } else {
                                                                var v28 uint8 = v2
                                                                if true {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        if true {
                                                                            switch 1 {
                                                                            case 9223372036854775807:
                                                                                v15 = 3.14159
                                                                            case 127:
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = 42
                                                                                }
                                                                            case 2147483647:
                                                                                fmt.Println(v15)
                                                                            default:
                                                                                v4 = append(v4, 1)
                                                                            }
                                                                            v3 += v3
                                                                            fmt.Println(v1)
                                                                        }
                                                                    case "\n\t":
                                                                        v29 := GenS3{F0: 15, F1: v2, F2: 410}
                                                                        _ = v29
                                                                    }
                                                                }
                                                                if (!((v1 + v1) > v1)) {
                                                                    var v30 int16 = 100
                                                                    v15 += v15
                                                                    v28 += (v2 ^ v2)
                                                                    _ = v30
                                                                } else {
                                                                    fmt.Println(fi19, v3, v1)
                                                                    var v31 int8 = 1
                                                                    _ = v31
                                                                }
                                                                _ = v28
                                                            }
                                                            if (!false) {
                                                                fmt.Println(v2)
                                                                if true {
                                                                    switch ("\n\t" + ("hello" + "a")) {
                                                                    case "hello":
                                                                        var v32 string = "\n\t"
                                                                        switch 42 {
                                                                        case 128:
                                                                            v33 := GenS3{F0: 100, F1: 244, F2: 1}
                                                                            _ = v33
                                                                        default:
                                                                            var v34 bool = (!false)
                                                                            _ = v34
                                                                        }
                                                                        var v35 int16 = 0
                                                                        _ = v32
                                                                        _ = v35
                                                                    default:
                                                                        v24 = v24
                                                                    }
                                                                    v36 := GenS4{F0: -9223372036854775808}
                                                                    _ = v36
                                                                } else {
                                                                    _ = len(v4)
                                                                    v37 := GenS3{F0: 1, F1: v2, F2: -1}
                                                                    if true {
                                                                        v38 := GenS4{F0: -1}
                                                                        _ = len(v4)
                                                                        if true {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = 200
                                                                            }
                                                                            var v39 uint = 0
                                                                            fmt.Println(v3, v39)
                                                                            _ = v39
                                                                        }
                                                                        _ = v38
                                                                    }
                                                                    v1 += 42
                                                                    _ = v37
                                                                }
                                                                if false {
                                                                    var v40 float64 = float64(v15)
                                                                    if false {
                                                                        var v41 int8 = 0
                                                                        fmt.Println(v1, v3)
                                                                        v41 -= (v41 % (((v41 | (v41 - (v41 / (((v41 >> ((65535) & 63))) | 1))))) | 1))
                                                                        _ = v41
                                                                    } else {
                                                                        fmt.Println(v24, v15)
                                                                    }
                                                                    _ = v40
                                                                } else {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 2
                                                                    }
                                                                    v42 := GenS5{F0: 1}
                                                                    if true {
                                                                        var v43 uint = 2
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 249
                                                                        }
                                                                        switch -468 {
                                                                        case -9223372036854775808:
                                                                            var v44 uint = 42
                                                                            _ = v44
                                                                        default:
                                                                            if ((-128 == 0) || (v2 != (^(v2)))) {
                                                                                v15 -= (v15 * ((v15 + v15) * v15))
                                                                                switch ("\n\t" + "\n\t") {
                                                                                case "\n\t":
                                                                                    v1 = v1
                                                                                case "foo bar":
                                                                                    _ = len(v4)
                                                                                }
                                                                                fmt.Println(v24)
                                                                            } else {
                                                                                v15 = 0.1
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = 2147483648
                                                                                }
                                                                            }
                                                                        }
                                                                        v45 := GenS6{F0: 0}
                                                                        v24 += (v24 | (v24 & 4294967295))
                                                                        _ = v43
                                                                        _ = v45
                                                                    }
                                                                    var v46 string = "a"
                                                                    _ = v42
                                                                    _ = v46
                                                                }
                                                            }
                                                        }
                                                        v1 -= v1
                                                    } else {
                                                        if false {
                                                            var v47 int64 = 922
                                                            _ = len(v4)
                                                            var v48 int = 13
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                    }
                                                    _ = v24
                                                }
                                            }
                                            _ = fi19
                                        }
                                    }
                                    _ = v18
                                }
                            }
                        }
                        var v49 int8 = -67
                        _ = fi16
                        _ = v49
                    }
                    _ = v15
                }
                switch ("" + ("hello" + ((("a" + ("a" + ("" + "hello"))) + ((("foo bar" + ("" + "\n\t")) + "") + (("" + "foo bar") + ""))) + ""))) {
                case "foo bar":
                    v50 := GenS3{F0: 127, F1: (v2 >> ((18446744073709551615) & 63)), F2: -596}
                    _ = v50
                case "\n\t":
                    if true {
                        var v51 float32 = 1e10
                        for fi52 := 0; fi52 < 13; fi52++ {
                            if len(v4) > 0 {
                                v4[0] = 1
                            }
                            v51 = 1.5
                            v53 := GenS3{F0: 100, F1: v2, F2: -9223372036854775808}
                            _ = fi52
                            _ = v53
                        }
                        var v54 bool = false
                        _ = v51
                        _ = v54
                    } else {
                        var v55 string = "a"
                        var v56 string = v55
                        _ = v55
                        _ = v56
                    }
                case "":
                    v2 += (v2 << ((18446744073709551615) & 63))
                }
                var v57 uint32 = 4294967295
                _ = v57
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    fmt.Println(v3)
    var v58 bool = false
    v2 = uint8(v1)
    if (0 <= v1) {
        v2 = v2
    } else {
        var v59 uint32 = 113
        var v60 uint16 = (^(v1))
        _ = v59
        _ = v60
    }
    if v58 {
        var v61 int = 220
        if (2147483647 <= -2147483648) {
            switch "foo bar" {
            case "foo bar":
                if ((v58 || (!v58)) && ((!(!((true && v58) || (-9223372036854775808 >= 650)))) && false)) {
                    fmt.Println(v3, v61, v3)
                    switch ("a" + "a") {
                    case "":
                        v61 += (v61 >> ((803) & 63))
                    case "\n\t":
                        if ((v58 || (true && v58)) && v58) {
                            for fi62 := 0; fi62 < 11; fi62++ {
                                v58 = ((v2 ^ (v2 + 255)) == v2)
                                var v63 float64 = v3
                                _ = fi62
                                _ = v63
                            }
                            var v64 int8 = 42
                            _ = v64
                        } else {
                            v3 -= 0.1
                            v65 := GenS2{F0: -1}
                            _ = v65
                        }
                    case "a":
                        fmt.Println(v1, v58)
                    default:
                        fmt.Println(v3)
                    }
                }
            }
            for fi66 := 0; fi66 < 9; fi66++ {
                _ = len(v4)
                _ = fi66
            }
        }
        _ = v61
    } else {
        v3 = 0.0
        if ((!(v58 && (!false))) || (538 != 4294967295)) {
            if v58 {
                v3 -= (1.5 / 1e10)
                var v67 uint16 = uint16(v2)
                _ = v67
            }
        } else {
            v4 = append(v4, 1)
            if false {
                var v68 uint64 = 0
                v4 = append(v4, 4294967295)
                _ = v68
            }
            for fi69 := 0; fi69 < 2; fi69++ {
                v3 -= float64(0.0)
                var v70 float64 = v3
                if len(v4) > 0 {
                    v4[0] = 136
                }
                _ = fi69
                _ = v70
            }
        }
        v71 := GenS2{F0: 9223372036854775807}
        _ = v71
    }
    var v72 int16 = -936
    fmt.Println(v2)
    if v58 {
        var v73 int = -9223372036854775808
        _ = v73
    } else {
        if true {
            fmt.Println(v3, v3)
            _ = len(v4)
            if v58 {
                v58 = v58
                for fi74 := 0; fi74 < 19; fi74++ {
                    v1 -= (v1 >> ((42) & 63))
                    v75 := GenS7{F0: fi74}
                    _ = fi74
                    _ = v75
                }
                switch ("\n\t" + "hello") {
                case "":
                    var v76 uint = 42
                    _ = v76
                case "foo bar":
                    if len(v4) > 0 {
                        v4[0] = 299
                    }
                }
            }
        }
        var v77 float32 = float32(1.5)
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v58:", v58)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v72
}
