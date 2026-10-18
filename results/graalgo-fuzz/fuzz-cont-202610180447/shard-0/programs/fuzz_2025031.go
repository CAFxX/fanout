package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 int16; F1 int32 }
func genFn0() uint16 {
    v7 := []int{}
    var v8 uint64 = 2
    fmt.Println(v8)
    if true {
        fmt.Println(v8)
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        if false {
            if (("hello" < ("a" + ("" + ""))) || true) {
                var v10 int = (fi9 ^ fi9)
                v7 = append(v7, (v10 * (-(fi9))))
                _ = v10
            }
            v11 := GenS1{F0: 2, F1: 0}
            _ = v11
        }
        _ = fi9
    }
    _ = v7
    _ = v8
    return 533
}

func main() {
    var v1 int = 0
    var v2 uint64 = 0
    var v3 uint = 896
    var v4 int16 = 0
    var v5 int = (^(v1))
    var v6 int64 = 460
    fmt.Println(v4, v3, v4)
    fmt.Println(v6)
    var v12 uint8 = 1
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 int8 = 0
        var v15 int8 = (v14 - (v14 | v14))
        _ = fi13
        _ = v14
        _ = v15
    }
    switch v5 {
    case 701:
        var v16 int = 42
        _ = v16
    }
    var v17 int32 = 42
    fmt.Println(v5, v3)
    switch "" {
    case "foo bar":
        if false {
            var v18 int = (v1 * v1)
            var v19 float64 = 0.0
            _ = v18
            _ = v19
        } else {
            v6 = int64(v4)
            v4 = v4
        }
    case "hello":
        v20 := GenS2{F0: (v6 % ((v6) | 1)), F1: (v3 << (((v3 ^ (v3 + v3))) & 63))}
        _ = v20
    default:
        v21 := []int{v1}
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v6 = v6
        _ = fi22
    }
    fmt.Println(v4, v1, v1)
    if (v17 > (v17 << ((v3) & 63))) {
        v23 := GenS1{F0: 237, F1: v17}
        v5 -= 2147483647
        _ = v23
    }
    fmt.Println(v6, v1)
    var v24 uint = (v3 << ((255) & 63))
    switch "\n\t" {
    case "a":
        switch "foo bar" {
        case "foo bar":
            if true {
                for fi25 := 0; fi25 < 3; fi25++ {
                    switch "hello" {
                    case "":
                        var v26 bool = ((0.0 - 0.1) > -2.25)
                        _ = v26
                    case "a":
                        v27 := []int{}
                        _ = v27
                    case "\n\t":
                        var v28 uint64 = 2
                        _ = v28
                    default:
                        v29 := []int{}
                        _ = v29
                    }
                    v30 := GenS2{F0: int64(v4), F1: v24}
                    v6 += v6
                    _ = fi25
                    _ = v30
                }
                v1 = (v1 ^ (v1 * (v1 - (v5 | v1))))
            } else {
                fmt.Println(v24)
                if false {
                    v4 -= (v4 | 1)
                    var v31 uint8 = 255
                    v32 := GenS1{F0: 0, F1: int32(v3)}
                    _ = v31
                    _ = v32
                }
                v1 -= v5
            }
        case "\n\t":
            if true {
                fmt.Println(v12, v4, v5)
                v17 = (v17 + v17)
                var v33 uint64 = v2
                _ = v33
            } else {
                var v34 bool = true
                v35 := GenS3{F0: 0, F1: 93}
                _ = v34
                _ = v35
            }
        }
    case "foo bar":
        if (!(false || (("" + "hello") <= "\n\t"))) {
            fmt.Println(v24, v4, v24)
            v36 := []int{}
            _ = v36
        } else {
            for fi37 := 0; fi37 < 12; fi37++ {
                v12 -= v12
                switch "foo bar" {
                case "hello":
                    for fi38 := 0; fi38 < 11; fi38++ {
                        v17 -= v17
                        _ = fi38
                    }
                case "":
                    for fi39 := 0; fi39 < 16; fi39++ {
                        switch (v5 >> (((v3 + (v24 - (v24 & v3)))) & 63)) {
                        case -515:
                            var v40 int = fi37
                            _ = v40
                        case -9223372036854775808:
                            v41 := []int{fi37}
                            _ = v41
                        case 737:
                            v42 := GenS4{F0: v4, F1: (42 > 1)}
                            _ = v42
                        }
                        switch "hello" {
                        case "foo bar":
                            v43 := GenS2{F0: int64(v2), F1: (v24 * (v24 + (v24 & v3)))}
                            _ = v43
                        case "a":
                            v24 = (v24 & 131)
                        case "\n\t":
                            fmt.Println(v12)
                        }
                        fmt.Println(v5, v3, v12)
                        _ = fi39
                    }
                case "foo bar":
                    if ("" == "foo bar") {
                        v44 := []int{fi37, (v5 % (((-(v1))) | 1)), (v5 | (v5 * 127))}
                        if true {
                            var v45 float64 = 1.5
                            switch "a" {
                            case "\n\t":
                                v12 -= 127
                            }
                            _ = v45
                        } else {
                            switch v1 {
                            case -9223372036854775808:
                                var v46 int64 = 9223372036854775807
                                _ = v46
                            case -2147483648:
                                var v47 uint32 = 0
                                _ = v47
                            default:
                                switch "foo bar" {
                                case "a":
                                    if (!false) {
                                        v12 -= v12
                                        var v48 bool = true
                                        _ = v48
                                    }
                                case "foo bar":
                                    v6 = (v6 | int64(v12))
                                case "":
                                    var v49 int16 = (v4 - (v4 / ((v4) | 1)))
                                    _ = v49
                                }
                            }
                        }
                        if false {
                            if len(v44) >= 2 {
                                v44 = v44[1:len(v44)]
                            }
                            if (!true) {
                                var v50 string = "a"
                                var v51 uint32 = uint32(v2)
                                var v52 float32 = (((0.0 / 1e10) + ((3.14159 + ((0.1 + ((((0.1 + float32(3.14159)) / ((float32(1e10) / 1.0) / 0.1)) / 3.14159) / 1.0)) + ((0.0 + 1e10) - (float32(0.0) - 1.5)))) / 0.1)) / -2.25)
                                _ = v50
                                _ = v51
                                _ = v52
                            } else {
                                if len(v44) > 0 {
                                    v44[0] = -538
                                }
                            }
                            if true {
                                switch (("a" + "foo bar") + "") {
                                case "\n\t":
                                    for fi53 := 0; fi53 < 20; fi53++ {
                                        if false {
                                            v12 -= (v12 * 128)
                                            v3 = v3
                                        }
                                        v1 += v1
                                        _ = fi53
                                    }
                                case "a":
                                    var v54 float32 = 1.5
                                    _ = v54
                                default:
                                    for fi55 := 0; fi55 < 6; fi55++ {
                                        v4 += (-(v4))
                                        var v56 float64 = 0.0
                                        fmt.Println(v1, v56, v5)
                                        _ = fi55
                                        _ = v56
                                    }
                                }
                                for fi57 := 0; fi57 < 4; fi57++ {
                                    v4 -= v4
                                    if (2147483647 != 1) {
                                        v58 := GenS2{F0: 1, F1: v3}
                                        fmt.Println(v2, v6, v6)
                                        _ = v58
                                    } else {
                                        var v59 bool = true
                                        fmt.Println(fi57, v59)
                                        _ = v59
                                    }
                                    v60 := GenS4{F0: 1, F1: (65535 != 255)}
                                    _ = fi57
                                    _ = v60
                                }
                                v61 := GenS3{F0: 255, F1: 2}
                                _ = v61
                            }
                        } else {
                            var v62 uint64 = v2
                            var v63 uint = 2
                            fmt.Println(v1, v1, v3)
                            _ = v62
                            _ = v63
                        }
                        _ = v44
                    } else {
                        if true {
                            for fi64 := 0; fi64 < 20; fi64++ {
                                v65 := GenS4{F0: 0, F1: false}
                                fmt.Println(v17, v17)
                                v66 := GenS5{F0: (-(v4)), F1: v17}
                                _ = fi64
                                _ = v65
                                _ = v66
                            }
                            var v67 string = ""
                            v12 -= (v12 / ((v12) | 1))
                            _ = v67
                        }
                        if ("a" != ("hello" + "\n\t")) {
                            fmt.Println(fi37, fi37)
                            v68 := []int{4294967295, 9223372036854775807}
                            _ = v68
                        } else {
                            v69 := []int{v1, 1, (fi37 | (-(v1)))}
                            _ = v69
                        }
                        fmt.Println(v4)
                    }
                default:
                    var v70 uint64 = (v2 + (^(v2)))
                    _ = v70
                }
                _ = fi37
            }
        }
    default:
        switch int(v4) {
        case -1:
            if ((("foo bar" + ("hello" + "foo bar")) + (("a" + "a") + "a")) > "") {
                if ((v2 << ((65535) & 63)) != v2) {
                    v24 += (v24 - v3)
                    v71 := GenS3{F0: v24, F1: 2}
                    _ = v71
                }
            }
        default:
            var v72 uint64 = (^(v2))
            _ = v72
        }
    }
    v2 -= 2
    v4 = (v4 % (((v4 >> (((v24 + (v3 / (((v24 - v24)) | 1)))) & 63))) | 1))
    var v73 uint8 = 200
    v74 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    fmt.Println("slicelen v74:", len(v74), v74)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
}
