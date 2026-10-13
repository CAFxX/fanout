package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16; F1 uint64 }

func main() {
    v1 := []uint32{65535, 42}
    var v2 string = ("a" + "foo bar")
    var v3 float32 = 0.0
    var v4 uint = 42
    v4 += v4
    v5 := GenS1{F0: v2}
    if false {
        if (!(!true)) {
            for fi6 := 0; fi6 < 2; fi6++ {
                if (!(!(999 < 0))) {
                    v2 = v2
                } else {
                    v7 := GenS2{F0: 0, F1: 867}
                    _ = v7
                }
                if (1 <= 9223372036854775807) {
                    v8 := GenS1{F0: (v2 + "hello")}
                    _ = v8
                } else {
                    v3 = (1.5 / v3)
                    if false {
                        var v9 string = (("hello" + v2) + ("foo bar" + v2))
                        v2 = ((v9 + (v2 + (((v2 + (((v2 + v2) + v9) + v2)) + "") + v2))) + (("hello" + "") + (("hello" + v9) + v2)))
                        _ = v9
                    }
                }
                var v10 string = v2
                _ = fi6
                _ = v10
            }
        }
        fmt.Println(v4, v3, v3)
    } else {
        for fi11 := 0; fi11 < 16; fi11++ {
            _ = len(v1)
            var v12 float64 = -2.25
            _ = fi11
            _ = v12
        }
        _ = len(v1)
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            v3 = (((v3 - v3) - 1e10) / v3)
            v4 = v4
            if len(v1) > 0 {
                v1[0] = 474
            }
            v15 := GenS2{F0: 65535, F1: 1}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    var v16 int16 = -1
    var v17 int16 = v16
    v17 = 1
    v18 := GenS3{F0: 1, F1: 321}
    v16 = -167
    if (!false) {
        var v19 string = v2
        if ((!true) && false) {
            switch ("hello" + "a") {
            case "":
                if (false || false) {
                    for fi20 := 0; fi20 < 10; fi20++ {
                        var v21 int8 = 42
                        fmt.Println(v2)
                        switch fi20 {
                        case 4294967295:
                            for fi22 := 0; fi22 < 16; fi22++ {
                                if (1 > 152) {
                                    v21 += v21
                                    fmt.Println(fi22, fi20)
                                    _ = len(v1)
                                } else {
                                    v23 := GenS4{F0: 305, F1: 360}
                                    v24 := GenS3{F0: 9223372036854775807, F1: 2147483647}
                                    var v25 int64 = 1
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                v26 := GenS5{F0: 18, F1: 255}
                                v3 += ((v3 / (0.1 + (v3 * float32(-2.25)))) + ((v3 + 1e10) - v3))
                                switch v2 {
                                case "a":
                                    if (514 > uint64(v4)) {
                                        _ = len(v1)
                                    }
                                case "hello":
                                    switch "a" {
                                    case "a":
                                        switch fi22 {
                                        case 32768:
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(v16, v3)
                                            case "foo bar":
                                                v17 -= int16(v21)
                                            case "a":
                                                switch (fi22 ^ fi22) {
                                                case -2147483648:
                                                    var v27 float64 = 0.1
                                                    _ = v27
                                                default:
                                                    var v28 int64 = 9223372036854775807
                                                    _ = v28
                                                }
                                            }
                                        case 128:
                                            v1 = append(v1, 754)
                                        case -904:
                                            switch fi22 {
                                            case 4294967295:
                                                var v29 uint16 = 65535
                                                _ = v29
                                            case 32767:
                                                switch (fi20 >> (((^(v4))) & 63)) {
                                                case 219:
                                                    switch (fi22 + fi22) {
                                                    case 1:
                                                        v1 = append(v1, 656)
                                                    default:
                                                        v16 += v16
                                                    }
                                                case 128:
                                                    v30 := GenS6{F0: (v19 + v2)}
                                                    v31 := GenS7{F0: 100, F1: 2147483648}
                                                    _ = v30
                                                    _ = v31
                                                case 123:
                                                    v17 = (v17 | v16)
                                                default:
                                                    if (v19 == v19) {
                                                        v21 += (v21 % ((v21) | 1))
                                                        if false {
                                                            v32 := GenS5{F0: v21, F1: 1}
                                                            v3 += v3
                                                            _ = v32
                                                        }
                                                    } else {
                                                        v16 += int16(fi22)
                                                        fmt.Println(v3, fi22, fi22)
                                                        v33 := GenS8{F0: (fi22 * fi22)}
                                                        fmt.Println(v17)
                                                        _ = v33
                                                    }
                                                }
                                            case -1:
                                                fmt.Println(v19)
                                            default:
                                                var v34 int64 = 0
                                                _ = v34
                                            }
                                        default:
                                            if (!true) {
                                                var v35 uint64 = 18446744073709551615
                                                v21 -= 92
                                                var v36 int32 = -1
                                                _ = v35
                                                _ = v36
                                            } else {
                                                _ = len(v1)
                                                v2 = v2
                                                v2 = v19
                                                v37 := GenS1{F0: ""}
                                                var v38 bool = (!(32767 >= 127))
                                                _ = v37
                                                _ = v38
                                            }
                                        }
                                    case "":
                                        switch (fi22 ^ -992) {
                                        case 9223372036854775807:
                                            v17 = v17
                                        case -920:
                                            v39 := GenS9{F0: (365 > 2)}
                                            _ = v39
                                        }
                                    case "foo bar":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        var v40 uint16 = 255
                                        _ = v40
                                    }
                                    v17 = v17
                                case "":
                                    fmt.Println(fi22)
                                default:
                                    if (false || false) {
                                        fmt.Println(v4)
                                        switch fi22 {
                                        case 32768:
                                            var v41 uint = (^(v4))
                                            _ = v41
                                        case 4294967295:
                                            var v42 int64 = -9223372036854775808
                                            var v43 float32 = (v3 / (-2.25 * (v3 * v3)))
                                            var v44 uint8 = 1
                                            fmt.Println(v19)
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        case -2147483648:
                                            _ = len(v1)
                                        }
                                    } else {
                                        if (!((255 < 128) || false)) {
                                            var v45 string = v19
                                            var v46 float64 = float64(1e10)
                                            _ = v45
                                            _ = v46
                                        } else {
                                            v19 = "hello"
                                            v19 = ""
                                        }
                                    }
                                }
                                _ = fi22
                                _ = v26
                            }
                        case 32768:
                            v2 = (((v2 + v19) + v2) + v19)
                        case 32767:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            switch v2 {
                            case "":
                                for fi47 := 0; fi47 < 19; fi47++ {
                                    var v48 int = fi47
                                    var v49 int8 = (v21 * 100)
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            default:
                                v16 -= 1
                            }
                        }
                        _ = fi20
                        _ = v21
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v50 := GenS10{F0: v17, F1: 2}
                    _ = v50
                } else {
                    v3 = v3
                }
            case "foo bar":
                var v51 bool = (v17 <= v17)
                _ = v51
            case "a":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                if true {
                    var v52 uint8 = 212
                    var v53 uint = v4
                    var v54 string = "hello"
                    _ = v52
                    _ = v53
                    _ = v54
                } else {
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                }
            }
        } else {
            fmt.Println(v16, v4, v3)
            fmt.Println(v19, v16)
        }
        _ = v19
    } else {
        if false {
            v1 = append(v1, 2)
        }
        v55 := GenS10{F0: (v17 + (v17 * v17)), F1: 2}
        _ = v55
    }
    v4 -= 4294967295
    for fi56 := 0; fi56 < 11; fi56++ {
        var v57 float64 = (((float64(v3) - -2.25) / (1e10 * 1e10)) + 0.1)
        v4 = v4
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi56
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
