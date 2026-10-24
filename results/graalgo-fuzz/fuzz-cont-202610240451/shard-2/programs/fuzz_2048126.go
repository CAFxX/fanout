package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int; F2 string }
func genFn0() string {
    if (9 == 0) {
        var v7 uint64 = 0
        for fi8 := 0; fi8 < 9; fi8++ {
            v7 = (v7 | 4294967295)
            for fi9 := 0; fi9 < 20; fi9++ {
                v7 = v7
                fmt.Println(v7, fi9)
                var v10 int64 = 0
                _ = fi9
                _ = v10
            }
            var v11 uint64 = 18446744073709551615
            _ = fi8
            _ = v11
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            if false {
                var v13 float64 = (float64(3.14159) * 1e10)
                _ = v13
            } else {
                v14 := []int{}
                var v15 int8 = 1
                _ = v14
                _ = v15
            }
            switch "hello" {
            case "":
                v7 = v7
            case "a":
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := []int{542, fi12}
                    var v18 int32 = 400
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            _ = fi12
        }
        _ = v7
    }
    var v19 uint32 = 679
    _ = v19
    return "a"
}

func main() {
    var v1 uint = 91
    v2 := []uint32{65535}
    var v3 float32 = 1e10
    var v4 int8 = 22
    v5 := []uint32{0, 0}
    var v6 float64 = ((1.5 / 1.5) * (float64((0.1 * v3)) * 1.5))
    v20 := GenS1{F0: uint8(v1), F1: 1}
    fmt.Println(v6)
    switch "hello" {
    case "foo bar":
        for fi21 := 0; fi21 < 2; fi21++ {
            if len(v5) > 0 {
                v5[0] = uint32(v1)
            }
            _ = fi21
        }
    default:
        var v22 string = ((("\n\t" + ("\n\t" + "hello")) + ("a" + "\n\t")) + ("hello" + "hello"))
        _ = v22
    }
    v4 = v4
    var v23 bool = ((42 < 255) && true)
    if v23 {
        if v23 {
            var v24 uint64 = 1
            v3 += float32(v6)
            _ = v24
        } else {
            fmt.Println(v3)
            fmt.Println(v1)
        }
        fmt.Println(v4, v4)
    } else {
        if ((v1 ^ 439) < v1) {
            if len(v2) > 0 {
                v2[0] = 65535
            }
            fmt.Println(v1, v4)
            v4 += (v4 + v4)
        } else {
            var v25 string = "\n\t"
            var v26 string = ("hello" + v25)
            fmt.Println(v25)
            _ = v25
            _ = v26
        }
    }
    switch "hello" {
    case "hello":
        v1 += v1
    case "a":
        if (!v23) {
            var v27 bool = v23
            v1 -= v1
            var v28 int = 80
            _ = v27
            _ = v28
        } else {
            var v29 uint = v1
            for fi30 := 0; fi30 < 9; fi30++ {
                switch ("a" + "\n\t") {
                case "\n\t":
                    if false {
                        var v31 uint8 = 128
                        _ = v31
                    }
                default:
                    var v32 int16 = 42
                    _ = v32
                }
                var v33 uint8 = 255
                _ = fi30
                _ = v33
            }
            _ = v29
        }
    case "":
        if v23 {
            fmt.Println(v3)
        } else {
            if (((!false) || v23) && false) {
                v1 = v1
                fmt.Println(v1, v6)
            }
            v3 -= (float32(0.1) - v3)
            if (994 > 2147483647) {
                var v34 uint = (v1 % ((v1) | 1))
                _ = v34
            } else {
                if v23 {
                    var v35 float64 = ((float64(float32(v6)) + 0.1) - v6)
                    _ = v35
                }
                if (!((-9223372036854775808 > 9223372036854775807) && (2 == 779))) {
                    var v36 uint = (v1 - 65535)
                    if v23 {
                        v37 := GenS1{F0: 0, F1: 751}
                        for fi38 := 0; fi38 < 16; fi38++ {
                            v36 = v36
                            _ = fi38
                        }
                        _ = v37
                    } else {
                        v36 -= v36
                        var v39 int64 = -436
                        _ = v39
                    }
                    if (0 == -1) {
                        fmt.Println(v23, v6)
                        var v40 bool = v23
                        if len(v5) > 0 {
                            v5[0] = 4294967295
                        }
                        _ = v40
                    } else {
                        for fi41 := 0; fi41 < 6; fi41++ {
                            v3 -= (v3 + ((v3 / 1.0) + ((3.14159 * ((0.1 / v3) / 1.5)) * (-2.25 - float32(3.14159)))))
                            if false {
                                v42 := GenS2{F0: 2147483647, F1: 463}
                                if false {
                                    v36 = (^(v1))
                                    if true {
                                        v43 := GenS1{F0: 128, F1: 42}
                                        var v44 uint64 = 2
                                        _ = v43
                                        _ = v44
                                    } else {
                                        var v45 uint = (^(v1))
                                        v4 = 1
                                        _ = v45
                                    }
                                }
                                _ = v42
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            _ = fi41
                        }
                        for fi46 := 0; fi46 < 3; fi46++ {
                            fmt.Println(v23, v3)
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = 2147483647
                                }
                                var v47 int64 = 0
                                _ = v47
                            } else {
                                if true {
                                    for fi48 := 0; fi48 < 17; fi48++ {
                                        if (int32(v4) > 2147483647) {
                                            fmt.Println(v1)
                                            v49 := GenS1{F0: 127, F1: 42}
                                            fmt.Println(v36)
                                            _ = v49
                                        }
                                        _ = fi48
                                    }
                                    v36 = 42
                                }
                                for fi50 := 0; fi50 < 17; fi50++ {
                                    if true {
                                        _ = len(v2)
                                        var v51 uint = v1
                                        _ = v51
                                    }
                                    var v52 float32 = (-2.25 / -2.25)
                                    _ = fi50
                                    _ = v52
                                }
                            }
                            _ = fi46
                        }
                        if ((v3 * v3) != v3) {
                            var v53 uint32 = 2147483647
                            v54 := GenS1{F0: 2, F1: 340}
                            _ = v53
                            _ = v54
                        } else {
                            var v55 int16 = -32768
                            _ = len(v2)
                            v56 := GenS3{F0: (v55 | (-(v55))), F1: ((255 == 255) || v23)}
                            _ = v55
                            _ = v56
                        }
                    }
                    _ = v36
                } else {
                    switch 127 {
                    case 322:
                        v3 -= -2.25
                    case 857:
                        if (32767 >= 789) {
                            if ((float64(v3) / 1.5) >= v6) {
                                switch "foo bar" {
                                case "\n\t":
                                    if (!v23) {
                                        if true {
                                            v6 -= v6
                                            v3 = 0.1
                                            for fi57 := 0; fi57 < 8; fi57++ {
                                                v58 := GenS1{F0: 104, F1: 157}
                                                v6 += (-2.25 + v6)
                                                v2 = append(v2, 65535)
                                                _ = fi57
                                                _ = v58
                                            }
                                        } else {
                                            v3 = ((v3 * (float32((v6 - float64(1e10))) + float32(1.5))) / (float32(v6) + (v3 + (v3 + (((v3 / v3) * v3) + -2.25)))))
                                            var v59 uint = (v1 << (((^(v1))) & 63))
                                            _ = v59
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 42
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            for fi60 := 0; fi60 < 8; fi60++ {
                                                v61 := GenS3{F0: -998, F1: (2 >= 167)}
                                                _ = fi60
                                                _ = v61
                                            }
                                        }
                                    }
                                case "hello":
                                    for fi62 := 0; fi62 < 15; fi62++ {
                                        v1 -= v1
                                        var v63 uint8 = 127
                                        _ = fi62
                                        _ = v63
                                    }
                                case "a":
                                    v2 = append(v2, 42)
                                }
                                switch 2147483647 {
                                case 767:
                                    v64 := GenS1{F0: 128, F1: -266}
                                    _ = v64
                                }
                                switch 2147483647 {
                                case 273:
                                    v23 = v23
                                case 2147483647:
                                    fmt.Println(v4, v1)
                                case 128:
                                    v65 := GenS1{F0: 1, F1: 38}
                                    _ = v65
                                }
                            } else {
                                if (0 <= -2147483648) {
                                    switch 32768 {
                                    case -511:
                                        fmt.Println(v6)
                                    default:
                                        var v66 uint = v1
                                        _ = v66
                                    }
                                    v3 += ((v3 * (v3 / v3)) / v3)
                                }
                                for fi67 := 0; fi67 < 12; fi67++ {
                                    _ = len(v2)
                                    if false {
                                        switch "hello" {
                                        case "\n\t":
                                            switch (("foo bar" + ("hello" + "foo bar")) + ("hello" + "hello")) {
                                            case "hello":
                                                v6 = ((v6 / ((v6 - (v6 * v6)) * (v6 / ((((float64(1e10) + ((3.14159 - 1.5) * v6)) / 1.5) / 1.0) / -2.25)))) - v6)
                                            }
                                        default:
                                            if false {
                                                for fi68 := 0; fi68 < 14; fi68++ {
                                                    if v23 {
                                                        var v69 int64 = 0
                                                        v3 = (((float32(0.0) * 3.14159) - v3) - v3)
                                                        _ = v69
                                                    }
                                                    v70 := GenS4{F0: 42, F1: 0}
                                                    fmt.Println(v4)
                                                    v3 -= v3
                                                    _ = fi68
                                                    _ = v70
                                                }
                                            }
                                        }
                                        v23 = (!(2 > 255))
                                    }
                                    _ = fi67
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            var v71 int16 = 1000
                            switch "foo bar" {
                            case "\n\t":
                                for fi72 := 0; fi72 < 3; fi72++ {
                                    v23 = v23
                                    v4 = v4
                                    _ = fi72
                                }
                            }
                            _ = v71
                        }
                    case 1:
                        var v73 uint8 = 255
                        _ = v73
                    }
                }
            }
        }
    }
    fmt.Println(v3, v6)
    v74 := GenS5{F0: 127, F1: 127, F2: ((("hello" + "foo bar") + (("hello" + "foo bar") + (("" + "a") + ""))) + "hello")}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
}
