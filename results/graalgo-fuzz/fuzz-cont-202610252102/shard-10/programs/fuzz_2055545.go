package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 int32; F1 bool; F2 uint8 }
type GenS7 struct { F0 uint8 }
func genFn0() int {
    v7 := GenS1{F0: 0, F1: 2}
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            if (!(!false)) {
                var v10 uint8 = 255
                var v11 int8 = -128
                _ = v10
                _ = v11
            } else {
                var v12 int16 = 1000
                switch (fi9 / ((int(v12)) | 1)) {
                case 9223372036854775807:
                    v13 := GenS2{F0: 42, F1: 0, F2: 238}
                    _ = v13
                case -9223372036854775808:
                    var v14 float64 = 3.14159
                    _ = v14
                }
                var v15 int32 = 555
                _ = v12
                _ = v15
            }
            switch fi8 {
            case 32768:
                var v16 int16 = 0
                _ = v16
            case 32767:
                fmt.Println(fi8)
            default:
                fmt.Println(fi9)
            }
            _ = fi9
        }
        _ = fi8
    }
    var v17 string = "a"
    var v18 float64 = 0.0
    _ = v17
    _ = v18
    _ = v7
    return -1
}
func genFn1() int64 {
    if (!(1 > -1)) {
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(fi19)
            var v20 int64 = 9223372036854775807
            v21 := []int{fi19, 2147483647, -151, (fi19 & fi19)}
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        if false {
            v22 := GenS2{F0: 0, F1: 127, F2: 255}
            var v23 int8 = 100
            _ = v22
            _ = v23
        }
        var v24 int32 = 0
        _ = v24
    }
    v25 := []int{128, -2147483648}
    _ = v25
    return 1
}

func main() {
    var v1 uint16 = 255
    var v2 int64 = 415
    var v3 uint64 = 179
    v4 := make(map[uint32]uint32)
    var v5 int8 = 0
    var v6 int64 = 1
    v2 = (^(v2))
    v6 -= -9223372036854775808
    for fi26 := 0; fi26 < 15; fi26++ {
        var v27 uint64 = v3
        if true {
            v5 -= v5
        }
        v27 = v3
        _ = fi26
        _ = v27
    }
    var v28 float32 = float32(3.14159)
    switch int(v3) {
    case -320:
        v1 = 32767
    case 32768:
        fmt.Println(v6, v5, v6)
    case 9223372036854775807:
        v29 := []int{101, -2147483648, 128, 637}
        _ = v29
    default:
        v28 -= (v28 + v28)
    }
    var v30 int = 0
    v31 := []int{127}
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 uint64 = v3
        var v34 int = genFn0()
        var v35 uint8 = 117
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    switch v30 {
    case -386:
        var v36 uint8 = 1
        _ = v36
    case 1:
        fmt.Println(v3, v30, v1)
    default:
        var v37 string = ("hello" + ("" + "\n\t"))
        _ = v37
    }
    var v38 float64 = 0.1
    v39 := GenS3{F0: v3, F1: 65535, F2: -2147483648}
    v40 := v4[155]
    for fi41 := 0; fi41 < 20; fi41++ {
        if true {
            v28 -= (v28 + (v28 / v28))
            switch "\n\t" {
            case "foo bar":
                if (!false) {
                    switch 735 {
                    case 127:
                        v42 := GenS2{F0: (v6 / ((v2) | 1)), F1: 127, F2: 1}
                        _ = v42
                    }
                    if (9223372036854775807 < -589) {
                        v4[(v40 * 2147483647)] = 0
                        fmt.Println(v1, v2)
                    }
                    v30 += fi41
                } else {
                    switch "a" {
                    case "a":
                        var v43 int64 = genFn1()
                        _ = v43
                    case "foo bar":
                        if false {
                            v44 := GenS4{F0: (v1 + (v1 * v1)), F1: v1, F2: -480}
                            var v45 string = "foo bar"
                            _ = v44
                            _ = v45
                        }
                    default:
                        if false {
                            for fi46 := 0; fi46 < 5; fi46++ {
                                var v47 string = "foo bar"
                                _ = fi46
                                _ = v47
                            }
                        }
                    }
                    switch 128 {
                    case -107:
                        var v48 uint32 = v40
                        _ = v48
                    case -653:
                        var v49 int16 = -102
                        _ = v49
                    default:
                        v2 += 508
                    }
                    for fi50 := 0; fi50 < 6; fi50++ {
                        fmt.Println(fi50, v30, fi50)
                        var v51 uint8 = 127
                        _ = fi50
                        _ = v51
                    }
                }
            default:
                if true {
                    switch (-(fi41)) {
                    case 2147483647:
                        switch 2147483647 {
                        case 4294967295:
                            fmt.Println(v30)
                        default:
                            v52 := GenS3{F0: (^(v3)), F1: 42, F2: 42}
                            _ = v52
                        }
                    default:
                        if (500 == 42) {
                            for fi53 := 0; fi53 < 12; fi53++ {
                                var v54 uint64 = (v3 * (v3 & (v3 % ((42) | 1))))
                                switch ("hello" + "a") {
                                case "":
                                    var v55 uint32 = (v40 & (v40 ^ v40))
                                    _ = v55
                                default:
                                    if false {
                                        var v56 bool = true
                                        _ = v56
                                    }
                                }
                                var v57 bool = (((^(v3)) != (v54 - v3)) && false)
                                _ = fi53
                                _ = v54
                                _ = v57
                            }
                            for fi58 := 0; fi58 < 5; fi58++ {
                                v59 := GenS5{F0: (v40 * (v40 | 42)), F1: int32(fi41)}
                                var v60 int8 = v5
                                fmt.Println(fi58)
                                _ = fi58
                                _ = v59
                                _ = v60
                            }
                        }
                    }
                } else {
                    var v61 uint8 = 127
                    for fi62 := 0; fi62 < 7; fi62++ {
                        v5 -= v5
                        _ = fi62
                    }
                    fmt.Println(v1)
                    _ = v61
                }
            }
        } else {
            delete(v4, 2147483647)
        }
        _ = fi41
    }
    v6 -= v6
    for fi63 := 0; fi63 < 7; fi63++ {
        v31 = append(v31, v30)
        _ = fi63
    }
    switch genFn0() {
    case 4294967295:
        if (!true) {
            v31 = append(v31, genFn0())
            switch "a" {
            case "hello":
                var v64 int16 = 1
                _ = v64
            case "":
                switch "hello" {
                case "a":
                    v65 := GenS5{F0: (v40 % (((v40 - v40)) | 1)), F1: -1}
                    _ = v65
                default:
                    for fi66 := 0; fi66 < 2; fi66++ {
                        v5 -= v5
                        if true {
                            fmt.Println(v40, v3)
                            v40 = 585
                            var v67 int16 = -32768
                            _ = v67
                        }
                        _ = fi66
                    }
                }
            default:
                if (!true) {
                    fmt.Println(v28, v6, v38)
                    delete(v4, 65535)
                    if ((v30 * (v30 ^ genFn0())) != (^(v30))) {
                        v68 := GenS6{F0: 2147483647, F1: true, F2: 0}
                        switch 32768 {
                        case 1:
                            switch "" {
                            case "":
                                for fi69 := 0; fi69 < 16; fi69++ {
                                    v31 = append(v31, fi69)
                                    _ = fi69
                                }
                            default:
                                v70 := GenS5{F0: v40, F1: -32}
                                _ = v70
                            }
                        case 2147483647:
                            v71 := GenS7{F0: 34}
                            _ = v71
                        case -107:
                            v72 := GenS6{F0: 2147483647, F1: true, F2: 1}
                            _ = v72
                        }
                        v4[v40] = (v40 + v40)
                        _ = v68
                    }
                }
            }
        } else {
            switch "foo bar" {
            case "hello":
                v73 := v4[(^(v40))]
                _ = v73
            case "foo bar":
                var v74 float64 = v38
                _ = v74
            case "":
                var v75 int64 = int64(v3)
                _ = v75
            }
            var v76 string = "foo bar"
            v2 = v6
            _ = v76
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
