package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int16; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 210
    v3 := make(map[uint32]string)
    if (2 < 29) {
        var v4 float64 = -2.25
        fmt.Println(v4, v2)
        _ = v4
    }
    switch ("hello" + "hello") {
    case "":
        fmt.Println(v2)
    case "hello":
        if false {
            v5 := []int{127, 2147483647, 2147483647, 9223372036854775807}
            v6 := GenS1{F0: 1, F1: true, F2: -2147483648}
            _ = v5
            _ = v6
        }
    case "a":
        v2 = v2
    }
    var v7 int = 42
    if (!true) {
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(v7, v2, v7)
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := []int{fi8, fi9}
                _ = fi9
                _ = v10
            }
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 int8 = 1
                _ = fi11
                _ = v12
            }
            _ = fi8
        }
        var v13 float64 = float64((float32(0.0) * (0.0 + 1e10)))
        _ = v13
    }
    v14 := GenS1{F0: 42, F1: ((true || true) && true), F2: (v7 / (((^(v7))) | 1))}
    switch "\n\t" {
    case "hello":
        var v15 int64 = -9223372036854775808
        _ = v15
    case "a":
        if false {
            if (!false) {
                v16 := []int{}
                var v17 int32 = -263
                switch ("a" + "") {
                case "hello":
                    var v18 float64 = 3.14159
                    _ = v18
                case "\n\t":
                    _ = len(v1)
                case "":
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if ((0 == 392) && (42 == 32767)) {
                            for fi20 := 0; fi20 < 6; fi20++ {
                                var v21 float32 = 3.14159
                                _ = fi20
                                _ = v21
                            }
                            fmt.Println(v17, v7, v7)
                        } else {
                            v17 = (-(v17))
                            v17 = v17
                            v2 = 1
                        }
                        var v22 int32 = v17
                        _ = fi19
                        _ = v22
                    }
                }
                _ = v16
                _ = v17
            }
            v23 := []int{}
            _ = v23
        }
    case "":
        var v24 uint16 = 2
        _ = v24
    }
    v7 -= 0
    var v25 int64 = 0
    if false {
        if (v25 > -9223372036854775808) {
            switch "foo bar" {
            case "":
                v2 = (v2 & v2)
            default:
                switch v7 {
                case -9223372036854775808:
                    fmt.Println(v25, v25, v25)
                default:
                    fmt.Println(v25)
                }
            }
            for fi26 := 0; fi26 < 7; fi26++ {
                v2 -= v2
                _ = fi26
            }
        }
        var v27 int32 = -747
        _ = v27
    } else {
        for fi28 := 0; fi28 < 14; fi28++ {
            switch "a" {
            case "a":
                if false {
                    switch (("foo bar" + "") + "\n\t") {
                    case "foo bar":
                        var v29 uint64 = v2
                        _ = v29
                    case "":
                        if (("hello" + ((("\n\t" + ("a" + ("foo bar" + "a"))) + "foo bar") + "a")) >= "\n\t") {
                            v30 := GenS1{F0: 2, F1: false, F2: (v7 + (v7 >> ((2) & 63)))}
                            for fi31 := 0; fi31 < 11; fi31++ {
                                if (false || false) {
                                    if (float32(1e10) < (((0.1 / float32(1.5)) - (float32(0.0) - (float32((3.14159 - 1.5)) * 1.5))) + float32((1e10 / (-2.25 / 1.5))))) {
                                        fmt.Println(v2)
                                    } else {
                                        var v32 uint64 = (v2 % ((v2) | 1))
                                        _ = v32
                                    }
                                } else {
                                    v7 = (fi28 * fi31)
                                    if (!false) {
                                        v3[65535] = ("foo bar" + (("" + "foo bar") + "a"))
                                        if false {
                                            if (255 >= 0) {
                                                v25 += -1
                                                var v33 bool = (("" <= "\n\t") && true)
                                                switch ("\n\t" + "foo bar") {
                                                case "foo bar":
                                                    _ = len(v1)
                                                case "\n\t":
                                                    fmt.Println(v2, fi31)
                                                case "a":
                                                    var v34 string = ((("foo bar" + ("\n\t" + "")) + "\n\t") + "foo bar")
                                                    _ = v34
                                                default:
                                                    switch fi28 {
                                                    case 42:
                                                        v7 = (v7 - fi31)
                                                    case -2147483648:
                                                        var v35 int = fi28
                                                        _ = v35
                                                    }
                                                }
                                                v1[42] = v7
                                                _ = v33
                                            } else {
                                                if (1 < 889) {
                                                    if (42 >= 496) {
                                                        fmt.Println(fi31, v25)
                                                        var v36 int32 = 42
                                                        _ = v36
                                                    }
                                                    var v37 int8 = 100
                                                    _ = v37
                                                }
                                            }
                                        } else {
                                            var v38 int8 = 0
                                            _ = v38
                                        }
                                        v1[65535] = -1
                                    } else {
                                        var v39 uint8 = 127
                                        _ = v39
                                    }
                                }
                                _ = fi31
                            }
                            _ = v30
                        } else {
                            v7 = (-(fi28))
                        }
                    case "a":
                        var v40 uint = 89
                        _ = v40
                    default:
                        if false {
                            v25 -= (v25 - 9223372036854775807)
                            var v41 int8 = int8(v25)
                            _ = v41
                        } else {
                            fmt.Println(v7, v7, v7)
                            v25 = v25
                            var v42 int8 = -128
                            _ = v42
                        }
                    }
                    for fi43 := 0; fi43 < 15; fi43++ {
                        if false {
                            v44 := GenS1{F0: 42, F1: true, F2: (v7 << ((398) & 63))}
                            _ = v44
                        } else {
                            v2 += 42
                            if true {
                                var v45 int16 = 1000
                                var v46 string = ""
                                var v47 float64 = -2.25
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                        }
                        _ = fi43
                    }
                }
            case "\n\t":
                v48 := []int{}
                _ = v48
            case "":
                fmt.Println(v7, v7, v7)
            default:
                var v49 float64 = 0.1
                _ = v49
            }
            var v50 int32 = 212
            _ = fi28
            _ = v50
        }
        var v51 bool = (!("" >= ""))
        v25 += (v25 - v25)
        _ = v51
    }
    if true {
        v52 := GenS2{F0: 351, F1: ((("\n\t" + (("hello" + "") + "foo bar")) + ((("hello" + "foo bar") + "hello") + "hello")) + ("" + "\n\t"))}
        fmt.Println(v7)
        _ = v52
    }
    var v53 int16 = -596
    v7 += v7
    if true {
        v25 -= (-(v25))
    } else {
        if true {
            switch "\n\t" {
            case "hello":
                if (899 >= 2) {
                    if false {
                        var v54 int16 = v53
                        _ = v54
                    }
                } else {
                    v2 = (v2 & 42)
                    if (false || false) {
                        var v55 uint64 = (v2 & v2)
                        var v56 int16 = (v53 >> ((2) & 63))
                        for fi57 := 0; fi57 < 2; fi57++ {
                            var v58 int = fi57
                            _ = fi57
                            _ = v58
                        }
                        _ = v55
                        _ = v56
                    } else {
                        v2 = (^(v2))
                        for fi59 := 0; fi59 < 15; fi59++ {
                            var v60 uint64 = 4294967295
                            switch "hello" {
                            case "foo bar":
                                var v61 uint64 = v2
                                _ = v61
                            case "\n\t":
                                v62 := GenS3{F0: (^(v53)), F1: -2147483648}
                                _ = v62
                            }
                            _ = fi59
                            _ = v60
                        }
                    }
                }
            default:
                _ = len(v1)
            }
            v7 -= v7
            var v63 int8 = 1
            _ = v63
        }
    }
    v53 = (v53 ^ v53)
    if (float64(((((0.0 * 3.14159) / (float32(3.14159) + 1e10)) - 0.1) / ((-2.25 * float32(((-2.25 + (-2.25 / 1.0)) / 1.0))) - 0.1))) > 3.14159) {
        v64 := []int{1, (v7 - (v7 % ((v7) | 1)))}
        switch (v7 & (v7 + (v7 * (v7 << ((65535) & 63))))) {
        case -242:
            var v65 bool = false
            _ = v65
        case -1:
            v66 := GenS3{F0: 987, F1: -573}
            _ = v66
        case 0:
            fmt.Println(v7, v2)
        }
        _ = v64
    }
    v67 := GenS3{F0: v53, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v53:", v53)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v53
    _ = v67
    _ = v7
}
