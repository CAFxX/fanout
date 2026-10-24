package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS8 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int16, p1 uint, p2 uint) int64 {
    v4 := []int{0, 32768, 127}
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 uint = 4294967295
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1() int8 {
    v7 := []int{}
    switch (("a" + ("\n\t" + "")) + "foo bar") {
    case "hello":
        if true {
            switch 4294967295 {
            case 1:
                _ = len(v7)
            case 614:
                _ = len(v7)
            case 42:
                _ = len(v7)
            default:
                for fi8 := 0; fi8 < 18; fi8++ {
                    for fi9 := 0; fi9 < 5; fi9++ {
                        var v10 string = "hello"
                        if (true && (42 > 277)) {
                            switch 739 {
                            case -1:
                                v10 = v10
                            case 854:
                                var v11 bool = (1000 > -331)
                                _ = v11
                            default:
                                if len(v7) > 0 {
                                    v7[0] = fi9
                                }
                            }
                        } else {
                            var v12 string = v10
                            if false {
                                v13 := GenS1{F0: 42}
                                fmt.Println(fi9, v10, v12)
                                v12 = "foo bar"
                                _ = v13
                            } else {
                                var v14 float64 = float64(1.5)
                                v15 := GenS2{F0: 255}
                                _ = v14
                                _ = v15
                            }
                            v10 = (((v10 + "a") + (v12 + v10)) + "a")
                            _ = v12
                        }
                        var v16 bool = (!(true && (v10 >= v10)))
                        _ = fi9
                        _ = v10
                        _ = v16
                    }
                    var v17 string = ""
                    v17 = ("hello" + "\n\t")
                    _ = fi8
                    _ = v17
                }
            }
        } else {
            var v18 uint32 = 2147483648
            _ = v18
        }
    default:
        switch "" {
        case "":
            var v19 float64 = 0.1
            _ = v19
        case "\n\t":
            var v20 int = 0
            _ = v20
        case "foo bar":
            var v21 uint32 = 65535
            _ = v21
        }
    }
    if (1 >= 42) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        if true {
            switch 602 {
            case 1:
                for fi22 := 0; fi22 < 15; fi22++ {
                    v23 := GenS3{F0: 0, F1: 0, F2: 44}
                    var v24 string = "a"
                    v24 = ""
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            case -658:
                var v25 uint32 = 2147483647
                _ = v25
            default:
                switch "a" {
                case "\n\t":
                    v26 := GenS4{F0: 716}
                    _ = v26
                case "":
                    switch ("a" + "a") {
                    case "hello":
                        _ = len(v7)
                    case "foo bar":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    default:
                        var v27 string = ""
                        _ = v27
                    }
                default:
                    _ = len(v7)
                }
            }
        }
        if (1 != 100) {
            v28 := GenS4{F0: 2147483647}
            for fi29 := 0; fi29 < 8; fi29++ {
                v7 = append(v7, fi29)
                _ = fi29
            }
            _ = v28
        } else {
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 int64 = -795
                _ = fi30
                _ = v31
            }
            var v32 bool = false
            v32 = (v32 && v32)
            _ = v32
        }
        for fi33 := 0; fi33 < 6; fi33++ {
            var v34 uint64 = 1
            switch "a" {
            case "foo bar":
                v35 := GenS5{F0: (fi33 << ((1) & 63)), F1: true}
                _ = v35
            case "\n\t":
                if len(v7) > 0 {
                    v7[0] = (-(fi33))
                }
            case "hello":
                v34 += v34
            default:
                for fi36 := 0; fi36 < 11; fi36++ {
                    var v37 int64 = genFn0(0, 255, 361)
                    v34 += (v34 % (((v34 + v34)) | 1))
                    _ = fi36
                    _ = v37
                }
            }
            _ = fi33
            _ = v34
        }
    }
    v7 = append(v7, 9223372036854775807)
    _ = v7
    return 42
}

func main() {
    var v1 float32 = 3.14159
    v2 := []int{-536, 986, 128}
    var v3 float32 = v1
    switch ("a" + "foo bar") {
    case "a":
        var v38 int64 = 978
        _ = v38
    case "foo bar":
        fmt.Println(v1)
    }
    v39 := GenS6{F0: 2147483647}
    if false {
        _ = len(v2)
        switch 128 {
        case -2147483648:
            if (104 < 64) {
                var v40 uint = 4294967295
                for fi41 := 0; fi41 < 13; fi41++ {
                    _ = len(v2)
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        if ((fi41 + (fi41 >> (((v40 % (((v40 - v40)) | 1))) & 63))) > -607) {
                            v1 = float32(float64((float32(0.1) * (v1 - v3))))
                        }
                    case "foo bar":
                        var v42 uint8 = 42
                        _ = v42
                    default:
                        for fi43 := 0; fi43 < 14; fi43++ {
                            v3 -= ((v3 / v3) - 3.14159)
                            var v44 int = (fi41 - -1)
                            fmt.Println(v3)
                            v40 = 65535
                            v45 := GenS6{F0: 485}
                            if false {
                                if (false && false) {
                                    v3 = v3
                                    _ = len(v2)
                                    v46 := GenS2{F0: 994}
                                    _ = v46
                                }
                                var v47 float32 = (v3 / v1)
                                _ = v47
                            }
                            var v48 int16 = -805
                            _ = fi43
                            _ = v44
                            _ = v45
                            _ = v48
                        }
                    }
                    var v49 uint32 = 264
                    _ = fi41
                    _ = v49
                }
                for fi50 := 0; fi50 < 11; fi50++ {
                    if len(v2) > 0 {
                        v2[0] = fi50
                    }
                    _ = fi50
                }
                _ = v40
            }
        case 91:
            if true {
                if (0 < 1) {
                    if (1 >= 7) {
                        switch "foo bar" {
                        case "\n\t":
                            v3 -= v1
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v1, v1)
                    } else {
                        if ((253 >= 0) || (-1 <= -1)) {
                            for fi51 := 0; fi51 < 9; fi51++ {
                                var v52 int32 = 1
                                for fi53 := 0; fi53 < 1; fi53++ {
                                    v54 := GenS5{F0: fi51, F1: (!true)}
                                    _ = fi53
                                    _ = v54
                                }
                                v55 := GenS7{F0: 703, F1: 69, F2: genFn1()}
                                _ = fi51
                                _ = v52
                                _ = v55
                            }
                            if (373 >= 1) {
                                v1 = v3
                            }
                        }
                        fmt.Println(v3)
                    }
                }
                fmt.Println(v3, v1)
                switch 324 {
                case 357:
                    v3 += v1
                case 724:
                    switch "foo bar" {
                    case "a":
                        var v56 float64 = 3.14159
                        _ = v56
                    case "\n\t":
                        for fi57 := 0; fi57 < 9; fi57++ {
                            for fi58 := 0; fi58 < 20; fi58++ {
                                if (0 != 0) {
                                    fmt.Println(v3)
                                    v3 = 0.0
                                }
                                var v59 float64 = ((1.5 - 1e10) * 0.0)
                                _ = fi58
                                _ = v59
                            }
                            var v60 int8 = genFn1()
                            if false {
                                _ = len(v2)
                            }
                            _ = fi57
                            _ = v60
                        }
                    }
                case -741:
                    for fi61 := 0; fi61 < 6; fi61++ {
                        if false {
                            var v62 int16 = 357
                            _ = v62
                        }
                        var v63 int32 = 378
                        v63 += 0
                        _ = fi61
                        _ = v63
                    }
                }
            } else {
                var v64 uint32 = 0
                _ = v64
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    var v65 uint16 = 0
    var v66 int8 = 100
    switch -372 {
    case -711:
        v65 = 255
    case -932:
        var v67 uint = 750
        _ = v67
    default:
        switch "\n\t" {
        case "\n\t":
            var v68 float64 = (1e10 * 3.14159)
            _ = v68
        }
    }
    v65 = (v65 - v65)
    v66 -= v66
    if (9223372036854775807 >= 1) {
        v69 := GenS8{F0: 65535, F1: (^(v65))}
        var v70 int = 42
        _ = v69
        _ = v70
    }
    for fi71 := 0; fi71 < 12; fi71++ {
        fmt.Println(v66, fi71)
        _ = fi71
    }
    v1 = v1
    fmt.Println(v65, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v65
    _ = v66
}
