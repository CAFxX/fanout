package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint16; F1 uint; F2 uint64 }

func main() {
    var v1 uint = 0
    var v2 int16 = -491
    var v3 int16 = v2
    var v4 uint64 = 42
    v5 := make(map[int]uint32)
    var v6 bool = true
    v7 := GenS1{F0: (v1 << ((v1) & 63))}
    var v8 int64 = 0
    v2 += -102
    switch -2147483648 {
    case 0:
        switch -29 {
        case 127:
            _ = len(v5)
        case 2147483647:
            v9 := GenS2{F0: 331}
            _ = v9
        default:
            var v10 uint16 = 255
            _ = v10
        }
    default:
        v2 -= v2
    }
    fmt.Println(v4)
    for fi11 := 0; fi11 < 16; fi11++ {
        _ = len(v5)
        if (32767 > (v2 >> ((v1) & 63))) {
            var v12 float64 = -2.25
            switch -678 {
            case 200:
                var v13 bool = (!("foo bar" <= "a"))
                _ = v13
            case 559:
                var v14 int32 = 1
                _ = v14
            default:
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 float32 = 1.5
                    _ = fi15
                    _ = v16
                }
            }
            _ = len(v5)
            _ = v12
        }
        _ = fi11
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v2 = 1
        v1 += v1
        _ = fi17
    }
    var v18 int = -2147483648
    v3 = v2
    if true {
        var v19 float32 = 0.0
        if false {
            if (645 < v2) {
                for fi20 := 0; fi20 < 19; fi20++ {
                    v1 = v1
                    fmt.Println(v2, v18, fi20)
                    _ = fi20
                }
                for fi21 := 0; fi21 < 7; fi21++ {
                    for fi22 := 0; fi22 < 10; fi22++ {
                        v23 := GenS1{F0: 1}
                        var v24 float32 = v19
                        var v25 int32 = 42
                        v26 := GenS1{F0: (^(v1))}
                        v24 = (v19 / (((v24 / ((v19 * 1.5) - float32(0.0))) * 1.5) / 1.0))
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    if v6 {
                        v8 = (v8 - (v8 >> ((v1) & 63)))
                        for fi27 := 0; fi27 < 7; fi27++ {
                            v4 += (v4 * (v4 << ((v1) & 63)))
                            if false {
                                switch fi21 {
                                case -204:
                                    delete(v5, (v18 % ((fi21) | 1)))
                                case -417:
                                    v28 := v5[fi27]
                                    _ = v28
                                }
                                var v29 int8 = 100
                                _ = v29
                            } else {
                                v8 = 475
                                switch fi21 {
                                case -711:
                                    v3 -= -1
                                default:
                                    v6 = ("a" < "a")
                                }
                                fmt.Println(v3, fi27, v2)
                            }
                            if (v6 || true) {
                                var v30 bool = true
                                if ((^(v8)) != (v8 / (((v8 - v8)) | 1))) {
                                    v30 = false
                                    v31 := GenS3{F0: (^(v2))}
                                    _ = v31
                                }
                                _ = v30
                            }
                            _ = fi27
                        }
                        v32 := GenS3{F0: -311}
                        _ = v32
                    } else {
                        v33 := GenS2{F0: (v8 % ((1) | 1))}
                        switch (v18 ^ v18) {
                        case -334:
                            v34 := v5[fi21]
                            _ = v34
                        case 540:
                            if (2 >= 42) {
                                var v35 float32 = (float32(1e10) + (v19 / v19))
                                _ = v35
                            }
                        default:
                            if (42 <= int8(v1)) {
                                var v36 int8 = 66
                                var v37 uint8 = 1
                                _ = v36
                                _ = v37
                            }
                        }
                        _ = v33
                    }
                    for fi38 := 0; fi38 < 1; fi38++ {
                        v3 -= int16(v1)
                        v39 := GenS1{F0: v1}
                        _ = fi38
                        _ = v39
                    }
                    _ = fi21
                }
            } else {
                if (false && ((v8 << (((v1 >> ((465) & 63))) & 63)) < 1)) {
                    v8 -= v8
                    for fi40 := 0; fi40 < 11; fi40++ {
                        if v6 {
                            v41 := []int{fi40, 0}
                            fmt.Println(v8)
                            _ = v41
                        }
                        var v42 uint = 42
                        _ = fi40
                        _ = v42
                    }
                }
                switch 4294967295 {
                case 42:
                    var v43 uint32 = uint32(v1)
                    _ = v43
                default:
                    var v44 int32 = -247
                    _ = v44
                }
                v19 += v19
            }
            v45 := GenS2{F0: v8}
            switch ("foo bar" + "foo bar") {
            case "hello":
                if true {
                    v46 := GenS4{F0: 685, F1: 233}
                    _ = v46
                }
            case "\n\t":
                fmt.Println(v8, v19, v4)
            }
            _ = v45
        }
        _ = v19
    } else {
        if (65535 == (v1 / ((v1) | 1))) {
            if (!(127 < 153)) {
                switch (v18 & v18) {
                case 1:
                    if ((v4 & (^(v4))) < v4) {
                        v4 = 798
                        for fi47 := 0; fi47 < 2; fi47++ {
                            var v48 float64 = (float64(0.0) * ((0.1 / 1.0) * (0.1 / (0.0 + -2.25))))
                            v49 := []int{v18, fi47}
                            var v50 string = "foo bar"
                            _ = fi47
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                        var v51 int32 = -2147483648
                        _ = v51
                    } else {
                        v18 += (-(v18))
                        var v52 uint64 = (v4 / ((v4) | 1))
                        _ = v52
                    }
                default:
                    v2 = (v2 + v3)
                }
            }
            if (!v6) {
                for fi53 := 0; fi53 < 3; fi53++ {
                    var v54 int = fi53
                    v55 := GenS5{F0: 196}
                    var v56 uint32 = 329
                    _ = fi53
                    _ = v54
                    _ = v55
                    _ = v56
                }
            } else {
                var v57 float64 = 1.5
                switch (v18 >> (((v1 * (v1 | (v1 >> ((255) & 63))))) & 63)) {
                case -9223372036854775808:
                    for fi58 := 0; fi58 < 9; fi58++ {
                        v59 := GenS4{F0: 1, F1: v8}
                        v18 = (v18 - (fi58 * v18))
                        v2 -= v2
                        _ = fi58
                        _ = v59
                    }
                case 9223372036854775807:
                    v8 -= (v8 << (((v1 - 2)) & 63))
                case -224:
                    v3 = (v2 >> ((v1) & 63))
                }
                _ = v57
            }
        }
        switch "hello" {
        case "":
            v3 = v3
        case "hello":
            for fi60 := 0; fi60 < 17; fi60++ {
                v61 := v5[fi60]
                _ = fi60
                _ = v61
            }
        default:
            if true {
                v6 = ((v8 * v8) >= 9223372036854775807)
                if true {
                    for fi62 := 0; fi62 < 1; fi62++ {
                        _ = len(v5)
                        fmt.Println(fi62)
                        v63 := GenS6{F0: 2, F1: ("hello" + ("\n\t" + "a"))}
                        _ = fi62
                        _ = v63
                    }
                    for fi64 := 0; fi64 < 4; fi64++ {
                        v65 := []int{(v18 / ((fi64) | 1))}
                        _ = fi64
                        _ = v65
                    }
                } else {
                    for fi66 := 0; fi66 < 13; fi66++ {
                        v67 := GenS1{F0: v1}
                        var v68 bool = v6
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                }
            }
        }
    }
    switch 907 {
    case 1:
        var v69 string = "\n\t"
        _ = v69
    case 32768:
        var v70 int8 = 127
        _ = v70
    }
    switch ("hello" + "foo bar") {
    case "hello":
        if (((v4 & (v4 >> ((v1) & 63))) <= v4) || false) {
            switch "" {
            case "foo bar":
                v5[1] = 199
            default:
                for fi71 := 0; fi71 < 19; fi71++ {
                    v3 = (v3 / ((v2) | 1))
                    _ = fi71
                }
            }
        }
    }
    v72 := GenS7{F0: 32767, F1: (v1 + v1), F2: v4}
    switch ("hello" + "") {
    case "hello":
        v5[(v18 | -9223372036854775808)] = 2
    default:
        fmt.Println(v1, v8)
    }
    for fi73 := 0; fi73 < 12; fi73++ {
        switch ("\n\t" + "\n\t") {
        case "hello":
            var v74 uint8 = 127
            _ = v74
        case "a":
            var v75 int8 = -51
            _ = v75
        case "foo bar":
            switch "a" {
            case "":
                fmt.Println(v1, fi73, v6)
            case "\n\t":
                v76 := v5[v18]
                _ = v76
            default:
                v3 += v2
            }
        }
        _ = fi73
    }
    fmt.Println(v18, v6)
    var v77 int64 = v8
    var v78 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v72
    _ = v77
    _ = v78
    _ = v8
}
