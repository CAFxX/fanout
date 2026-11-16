package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 bool; F1 uint32 }

func main() {
    v1 := []int{}
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 uint64 = 42
    var v5 uint32 = 4294967295
    var v6 uint64 = (v4 & (v4 ^ (v4 << ((255) & 63))))
    var v7 uint16 = 880
    v8 := GenS1{F0: true}
    var v9 uint64 = (v6 | 18446744073709551615)
    var v10 int32 = 0
    v7 = (^(v7))
    v11 := GenS1{F0: (-128 < 127)}
    for fi12 := 0; fi12 < 20; fi12++ {
        v9 -= (v9 + 1)
        for fi13 := 0; fi13 < 9; fi13++ {
            switch "hello" {
            case "":
                var v14 float32 = 1.5
                _ = v14
            case "hello":
                v5 -= (v5 | v5)
            case "foo bar":
                if (v5 > 4294967295) {
                    var v15 uint64 = (^(v6))
                    switch "hello" {
                    case "\n\t":
                        delete(v2, "")
                    case "foo bar":
                        v16 := GenS1{F0: (v15 >= uint64(v5))}
                        _ = v16
                    case "hello":
                        var v17 int64 = 0
                        fmt.Println(v17, v4, v17)
                        _ = v17
                    default:
                        switch fi12 {
                        case 0:
                            v15 = 4294967295
                        case 34:
                            v10 -= (v10 >> ((18446744073709551615) & 63))
                        case 2147483647:
                            var v18 string = "foo bar"
                            _ = v18
                        default:
                            switch (("\n\t" + (("\n\t" + ("foo bar" + ("\n\t" + ("foo bar" + "\n\t")))) + "foo bar")) + ("foo bar" + "\n\t")) {
                            case "a":
                                if false {
                                    v19 := v3["hello"]
                                    v20 := GenS2{F0: v19}
                                    _ = v19
                                    _ = v20
                                } else {
                                    v21 := GenS3{F0: -731}
                                    if true {
                                        fmt.Println(v4)
                                    }
                                    var v22 uint = 241
                                    _ = v21
                                    _ = v22
                                }
                            case "\n\t":
                                v4 = (v9 ^ v15)
                            }
                        }
                    }
                    var v23 int8 = int8(v7)
                    _ = v15
                    _ = v23
                } else {
                    v9 += v9
                }
            default:
                v9 -= v9
            }
            var v24 uint8 = 19
            _ = fi13
            _ = v24
        }
        _ = fi12
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        fmt.Println(fi25, v6, fi25)
        _ = fi25
    }
    delete(v2, "")
    if false {
        if true {
            v4 = v6
            if true {
                if true {
                    var v26 int8 = 1
                    _ = v26
                } else {
                    if (true && false) {
                        v9 += v4
                        v6 += v6
                    } else {
                        if true {
                            v27 := GenS3{F0: -622}
                            var v28 uint64 = v6
                            _ = v27
                            _ = v28
                        }
                    }
                }
                switch 9223372036854775807 {
                case -897:
                    for fi29 := 0; fi29 < 7; fi29++ {
                        var v30 uint8 = 128
                        for fi31 := 0; fi31 < 20; fi31++ {
                            fmt.Println(fi31, fi29)
                            var v32 float32 = -2.25
                            if true {
                                v5 = v5
                                var v33 string = "hello"
                                v33 = v33
                                v6 += 0
                                _ = v33
                            }
                            _ = fi31
                            _ = v32
                        }
                        if (v7 == uint16(v6)) {
                            for fi34 := 0; fi34 < 16; fi34++ {
                                v3[(((("a" + "\n\t") + "hello") + "hello") + "\n\t")] = ("foo bar" + "hello")
                                v30 = v30
                                v35 := GenS1{F0: true}
                                _ = fi34
                                _ = v35
                            }
                            for fi36 := 0; fi36 < 11; fi36++ {
                                v4 = 42
                                _ = fi36
                            }
                        } else {
                            for fi37 := 0; fi37 < 19; fi37++ {
                                v30 = uint8(v7)
                                fmt.Println(v5, v9, v4)
                                switch (fi37 & fi37) {
                                case -696:
                                    var v38 uint64 = uint64(v5)
                                    _ = v38
                                default:
                                    v39 := GenS1{F0: ((!(1.5 == ((((1e10 * -2.25) / ((-2.25 * 3.14159) - 1e10)) + 0.1) / (float64(-2.25) - 1e10)))) && (!(false || true)))}
                                    _ = v39
                                }
                                _ = fi37
                            }
                        }
                        _ = fi29
                        _ = v30
                    }
                case -2147483648:
                    v40 := GenS4{F0: -9223372036854775808, F1: v9}
                    _ = v40
                default:
                    var v41 bool = false
                    _ = v41
                }
            }
            var v42 uint = 4294967295
            _ = v42
        }
        switch ("a" + (("a" + "hello") + "foo bar")) {
        case "\n\t":
            switch "hello" {
            case "hello":
                if (!true) {
                    if len(v1) > 0 {
                        v1[0] = -2147483648
                    }
                    if (!(254 != 176)) {
                        v6 += v9
                        for fi43 := 0; fi43 < 4; fi43++ {
                            v5 += (^(v5))
                            v2[""] = ""
                            var v44 int8 = -71
                            _ = fi43
                            _ = v44
                        }
                        fmt.Println(v5, v4)
                    } else {
                        if (!false) {
                            v45 := GenS4{F0: 0, F1: v6}
                            _ = v45
                        } else {
                            switch "foo bar" {
                            case "":
                                for fi46 := 0; fi46 < 10; fi46++ {
                                    var v47 uint8 = 255
                                    v48 := GenS3{F0: (fi46 * (fi46 / (((fi46 >> ((18446744073709551615) & 63))) | 1)))}
                                    switch ("hello" + "\n\t") {
                                    case "hello":
                                        var v49 bool = (100 < -21)
                                        _ = v49
                                    case "":
                                        var v50 uint = 18446744073709551615
                                        _ = v50
                                    }
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                }
                            }
                        }
                        if true {
                            var v51 uint64 = (v9 + 361)
                            var v52 float32 = 3.14159
                            v52 += (0.1 + 1.5)
                            _ = v51
                            _ = v52
                        }
                        if true {
                            var v53 float64 = float64(float32(1e10))
                            _ = v53
                        } else {
                            switch 4294967295 {
                            case -1:
                                if len(v1) > 0 {
                                    v1[0] = -2147483648
                                }
                            case 370:
                                if (false || false) {
                                    for fi54 := 0; fi54 < 15; fi54++ {
                                        v55 := GenS3{F0: (fi54 | -9223372036854775808)}
                                        v56 := GenS3{F0: (fi54 & (fi54 * fi54))}
                                        _ = fi54
                                        _ = v55
                                        _ = v56
                                    }
                                } else {
                                    switch -9223372036854775808 {
                                    case 4294967295:
                                        var v57 float32 = 0.1
                                        _ = v57
                                    }
                                    fmt.Println(v7, v4)
                                    v7 = v7
                                }
                            default:
                                v5 += (v5 & (v5 % ((uint32(v4)) | 1)))
                            }
                        }
                    }
                    switch 0 {
                    case 352:
                        for fi58 := 0; fi58 < 18; fi58++ {
                            if true {
                                v4 = (v9 | 2)
                                v7 = 255
                                var v59 int64 = -9223372036854775808
                                _ = v59
                            } else {
                                _ = len(v2)
                                v6 = (v6 + v6)
                            }
                            var v60 uint = uint(v6)
                            _ = fi58
                            _ = v60
                        }
                    case 128:
                        switch -1 {
                        case 4294967295:
                            for fi61 := 0; fi61 < 1; fi61++ {
                                v5 = v5
                                v62 := GenS4{F0: 42, F1: 0}
                                _ = fi61
                                _ = v62
                            }
                        }
                    case 0:
                        v63 := GenS5{F0: 42, F1: 1}
                        _ = v63
                    }
                } else {
                    delete(v3, "hello")
                }
            case "":
                for fi64 := 0; fi64 < 9; fi64++ {
                    switch "hello" {
                    case "":
                        if len(v1) > 0 {
                            v1[0] = (fi64 << ((744) & 63))
                        }
                    default:
                        if false {
                            switch ("a" + "hello") {
                            case "a":
                                if true {
                                    var v65 float32 = (float32((0.0 + 0.0)) * 1e10)
                                    _ = v65
                                } else {
                                    switch "hello" {
                                    case "":
                                        if (uint32(v9) <= (^(v5))) {
                                            fmt.Println(v6, v7, v10)
                                            if (true && ((0 > -618) || (!true))) {
                                                var v66 int8 = -90
                                                _ = v66
                                            } else {
                                                for fi67 := 0; fi67 < 12; fi67++ {
                                                    var v68 uint16 = v7
                                                    fmt.Println(v10, v6, v68)
                                                    fmt.Println(fi67, v10)
                                                    var v69 uint64 = (v6 * v4)
                                                    fmt.Println(fi64)
                                                    _ = fi67
                                                    _ = v68
                                                    _ = v69
                                                }
                                                var v70 int32 = 299
                                                _ = v70
                                            }
                                            v9 += (^(v6))
                                        } else {
                                            switch ("foo bar" + "\n\t") {
                                            case "hello":
                                                v6 = 42
                                            case "a":
                                                v7 = v7
                                            case "foo bar":
                                                v71 := GenS6{F0: (-29 < -61), F1: (v5 ^ (v5 << ((162) & 63)))}
                                                _ = v71
                                            }
                                        }
                                    }
                                    if (32767 <= -761) {
                                        v9 += (v4 >> ((1) & 63))
                                        fmt.Println(v6, v10)
                                    }
                                }
                            case "foo bar":
                                fmt.Println(v4, fi64)
                            case "":
                                switch "hello" {
                                case "hello":
                                    fmt.Println(fi64, fi64, v7)
                                case "":
                                    var v72 float32 = -2.25
                                    _ = v72
                                default:
                                    fmt.Println(v6, v9)
                                }
                            }
                        }
                    }
                    _ = fi64
                }
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            _ = len(v3)
        }
        for fi73 := 0; fi73 < 17; fi73++ {
            v6 = (v6 + v9)
            v10 = (-(v10))
            switch "a" {
            case "hello":
                v74 := GenS4{F0: 42, F1: (v6 - v6)}
                _ = v74
            default:
                if (0.0 < -2.25) {
                    var v75 int32 = v10
                    var v76 int16 = 1
                    var v77 uint32 = (v5 & (v5 - 176))
                    _ = v75
                    _ = v76
                    _ = v77
                }
            }
            _ = fi73
        }
    } else {
        var v78 uint16 = 2
        _ = v78
    }
    v1 = append(v1, -9223372036854775808)
    for fi79 := 0; fi79 < 10; fi79++ {
        var v80 float64 = (((0.1 - (1e10 - ((-2.25 / (0.1 * (float64((0.0 + 3.14159)) / 1.0))) / (-2.25 + float64(0.1))))) + 3.14159) / (-2.25 - 1e10))
        _ = fi79
        _ = v80
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
