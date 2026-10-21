package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint }
type GenS5 struct { F0 bool; F1 uint16; F2 int16 }

func main() {
    var v1 uint8 = 42
    var v2 uint64 = 545
    var v3 uint = 927
    var v4 bool = true
    v5 := GenS1{F0: 18446744073709551615}
    switch ("hello" + ("foo bar" + "hello")) {
    case "hello":
        var v6 int32 = -2147483648
        _ = v6
    case "a":
        var v7 bool = true
        _ = v7
    case "\n\t":
        v2 = 0
    default:
        v1 = 255
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        v2 -= v2
        _ = fi8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := []int{-9223372036854775808, fi9, -30}
        v10 = append(v10, (fi9 - fi9))
        var v11 float32 = 1e10
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 42, F1: 42, F2: -1}
    if v4 {
        v4 = true
    }
    v2 -= uint64(v1)
    var v13 float32 = ((((-2.25 + 0.0) - 0.0) / 1e10) - (0.0 * 0.0))
    if ((1 > (^(v2))) && (!false)) {
        if (v1 != v1) {
            v14 := GenS3{F0: 2147483647}
            switch "" {
            case "a":
                for fi15 := 0; fi15 < 17; fi15++ {
                    if true {
                        v16 := []int{fi15}
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        var v17 uint32 = 65535
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 20; fi18++ {
                        if false {
                            if v4 {
                                if ((0.0 / float64(v13)) > 3.14159) {
                                    v2 = (v2 + uint64(v1))
                                    v19 := GenS4{F0: 1, F1: 73, F2: v3}
                                    v20 := []int{fi18, fi18}
                                    if v4 {
                                        v21 := GenS2{F0: 0, F1: 0, F2: 681}
                                        _ = v21
                                    }
                                    _ = v19
                                    _ = v20
                                }
                            } else {
                                v3 += (v3 >> (((v3 % ((833) | 1))) & 63))
                                if v4 {
                                    var v22 float64 = float64((v13 - v13))
                                    _ = v22
                                }
                            }
                        }
                        v23 := GenS4{F0: 42, F1: -1, F2: 679}
                        _ = fi18
                        _ = v23
                    }
                    _ = fi15
                }
            case "foo bar":
                if v4 {
                    for fi24 := 0; fi24 < 18; fi24++ {
                        for fi25 := 0; fi25 < 16; fi25++ {
                            v4 = ((1 != 0) && v4)
                            if (-1 >= 2147483647) {
                                var v26 uint8 = (v1 | v1)
                                switch "a" {
                                case "hello":
                                    fmt.Println(v1)
                                default:
                                    v3 += v3
                                    fmt.Println(v1, v26, fi24)
                                    v27 := []int{}
                                    _ = v27
                                }
                                _ = v26
                            }
                            if false {
                                if v4 {
                                    var v28 int = (-(fi25))
                                    fmt.Println(v2, v2)
                                    _ = v28
                                }
                                v29 := GenS5{F0: (!false), F1: 255, F2: 32767}
                                _ = v29
                            }
                            _ = fi25
                        }
                        v4 = v4
                        _ = fi24
                    }
                } else {
                    v30 := GenS1{F0: 0}
                    for fi31 := 0; fi31 < 4; fi31++ {
                        var v32 string = "\n\t"
                        switch (fi31 % ((fi31) | 1)) {
                        case -493:
                            if v4 {
                                v32 = "foo bar"
                                switch "a" {
                                case "hello":
                                    if v4 {
                                        v2 -= (v2 & v2)
                                        fmt.Println(v32, v2)
                                        for fi33 := 0; fi33 < 19; fi33++ {
                                            var v34 float64 = 1e10
                                            _ = fi33
                                            _ = v34
                                        }
                                    }
                                case "a":
                                    v35 := []int{fi31, fi31, (^(fi31)), fi31}
                                    _ = v35
                                case "\n\t":
                                    v36 := []int{127, -766}
                                    _ = v36
                                default:
                                    if true {
                                        for fi37 := 0; fi37 < 11; fi37++ {
                                            var v38 uint64 = v2
                                            var v39 uint8 = 6
                                            v38 = v38
                                            _ = fi37
                                            _ = v38
                                            _ = v39
                                        }
                                        fmt.Println(v13, v2, v1)
                                        if false {
                                            if true {
                                                v40 := GenS2{F0: 74, F1: -9223372036854775808, F2: -9223372036854775808}
                                                v41 := []int{-938, fi31}
                                                _ = v40
                                                _ = v41
                                            } else {
                                                var v42 bool = (0.0 < 1.5)
                                                for fi43 := 0; fi43 < 3; fi43++ {
                                                    var v44 bool = true
                                                    fmt.Println(v3, v2)
                                                    _ = fi43
                                                    _ = v44
                                                }
                                                _ = v42
                                            }
                                            fmt.Println(v13)
                                        } else {
                                            if v4 {
                                                v2 = (v2 << ((18446744073709551615) & 63))
                                                switch (fi31 % ((fi31) | 1)) {
                                                case -924:
                                                    fmt.Println(v2)
                                                case 462:
                                                    if true {
                                                        v45 := GenS5{F0: (v1 != 3), F1: 386, F2: 32767}
                                                        switch v32 {
                                                        case "a":
                                                            var v46 int16 = 42
                                                            _ = v46
                                                        case "\n\t":
                                                            if v4 {
                                                                var v47 float32 = (v13 - v13)
                                                                _ = v47
                                                            }
                                                        default:
                                                            var v48 uint = v3
                                                            _ = v48
                                                        }
                                                        _ = v45
                                                    } else {
                                                        switch v32 {
                                                        case "":
                                                            var v49 int32 = -915
                                                            _ = v49
                                                        case "a":
                                                            switch ((v32 + "") + v32) {
                                                            case "":
                                                                if ((v3 / ((v3) | 1)) < (v3 | (v3 & uint(v1)))) {
                                                                    switch fi31 {
                                                                    case 9223372036854775807:
                                                                        var v50 string = (v32 + "foo bar")
                                                                        _ = v50
                                                                    case 4294967295:
                                                                        v51 := GenS3{F0: 0}
                                                                        _ = v51
                                                                    }
                                                                }
                                                            default:
                                                                if false {
                                                                    fmt.Println(v4, v3)
                                                                    var v52 uint8 = 128
                                                                    _ = v52
                                                                }
                                                            }
                                                        case "hello":
                                                            var v53 uint8 = v1
                                                            _ = v53
                                                        }
                                                    }
                                                }
                                            } else {
                                                var v54 uint64 = 4294967295
                                                v55 := []int{fi31, fi31, fi31}
                                                fmt.Println(v13, v54)
                                                _ = v54
                                                _ = v55
                                            }
                                            switch v32 {
                                            case "hello":
                                                var v56 int64 = -9223372036854775808
                                                _ = v56
                                            case "a":
                                                v32 = ((v32 + v32) + v32)
                                            case "foo bar":
                                                if (!(fi31 != 2147483647)) {
                                                    v57 := []int{-9223372036854775808, (fi31 >> ((206) & 63)), fi31}
                                                    _ = v57
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        _ = fi31
                        _ = v32
                    }
                    var v58 uint32 = 2147483648
                    _ = v30
                    _ = v58
                }
            default:
                v2 = (v2 + v2)
            }
            _ = v14
        }
    } else {
        var v59 uint32 = 636
        _ = v59
    }
    var v60 int64 = 0
    v60 = (v60 >> (((v3 % (((v3 & v3)) | 1))) & 63))
    for fi61 := 0; fi61 < 12; fi61++ {
        var v62 float32 = v13
        v63 := GenS4{F0: 1, F1: -128, F2: (^(v3))}
        var v64 float32 = (3.14159 / (v62 + (v13 + v62)))
        _ = fi61
        _ = v62
        _ = v63
        _ = v64
    }
    var v65 string = ("" + ("" + "a"))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v65
}
