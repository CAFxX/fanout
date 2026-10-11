package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int32; F2 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 int64; F1 uint32 }
func genFn0() int32 {
    var v4 int8 = 122
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int8 = v4
        fmt.Println(fi5, v4)
        v6 += (v4 | 42)
        _ = fi5
        _ = v6
    }
    _ = v4
    return -1
}
func genFn1() string {
    if ("hello" <= "a") {
        switch "" {
        case "foo bar":
            if (!true) {
                v7 := []int{-9223372036854775808, -2147483648, 1}
                var v8 uint = 1
                var v9 int16 = -1
                _ = v7
                _ = v8
                _ = v9
            }
        default:
            v10 := GenS1{F0: 146}
            _ = v10
        }
    } else {
        if true {
            v11 := []int{}
            _ = v11
        }
        var v12 bool = false
        _ = v12
    }
    var v13 float64 = 0.0
    _ = v13
    return ("\n\t" + "hello")
}

func main() {
    var v1 uint = 0
    var v2 uint = v1
    v3 := []uint32{}
    if len(v3) > 0 {
        v3[0] = 528
    }
    var v14 int16 = 516
    fmt.Println(v2, v2)
    v14 += v14
    if ("a" > "foo bar") {
        var v15 uint16 = 32767
        for fi16 := 0; fi16 < 16; fi16++ {
            switch (("a" + "") + "foo bar") {
            case "":
                if true {
                    if false {
                        v15 = 255
                        if ((fi16 * fi16) != int(v1)) {
                            v17 := GenS1{F0: 71}
                            for fi18 := 0; fi18 < 4; fi18++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v19 uint64 = 2
                                _ = fi18
                                _ = v19
                            }
                            _ = v17
                        } else {
                            if true {
                                for fi20 := 0; fi20 < 12; fi20++ {
                                    switch (fi16 << (((^(v1))) & 63)) {
                                    case 237:
                                        switch "foo bar" {
                                        case "hello":
                                            switch 9223372036854775807 {
                                            case 0:
                                                var v21 bool = false
                                                _ = v21
                                            case -2147483648:
                                                fmt.Println(v2, fi20, v14)
                                            }
                                        }
                                        var v22 uint = (v1 + v2)
                                        _ = v22
                                    case 1:
                                        var v23 uint = 255
                                        _ = v23
                                    default:
                                        v24 := GenS1{F0: 42}
                                        _ = v24
                                    }
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(v2, fi20, v15)
                                    default:
                                        var v25 int = (-(fi16))
                                        _ = v25
                                    }
                                    var v26 bool = true
                                    if len(v3) > 0 {
                                        v3[0] = 2147483648
                                    }
                                    v2 -= v1
                                    _ = fi20
                                    _ = v26
                                }
                                v14 += (^(v14))
                            }
                            v27 := GenS2{F0: 1, F1: 2147483647, F2: 2147483647}
                            v15 = (v15 / (((v15 ^ uint16(v1))) | 1))
                            _ = v27
                        }
                        if (3.14159 != 0.0) {
                            for fi28 := 0; fi28 < 15; fi28++ {
                                var v29 int64 = -1
                                if (!(!(255 == (v15 & (v15 / (((v15 ^ (v15 * 65535))) | 1)))))) {
                                    v1 = 2
                                    var v30 int = fi28
                                    _ = v30
                                } else {
                                    v3 = append(v3, 2)
                                    var v31 float64 = 3.14159
                                    _ = v31
                                }
                                _ = fi28
                                _ = v29
                            }
                            var v32 int16 = 0
                            _ = v32
                        } else {
                            switch fi16 {
                            case 664:
                                v2 -= v2
                            case 4294967295:
                                var v33 float64 = 1.5
                                _ = v33
                            default:
                                for fi34 := 0; fi34 < 4; fi34++ {
                                    if true {
                                        v35 := GenS3{F0: 1}
                                        _ = v35
                                    }
                                    _ = fi34
                                }
                            }
                        }
                    } else {
                        for fi36 := 0; fi36 < 2; fi36++ {
                            var v37 int32 = 2147483647
                            if false {
                                fmt.Println(v15, v1, v15)
                                fmt.Println(v37)
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = len(v3)
                            }
                            if len(v3) > 0 {
                                v3[0] = 2
                            }
                            _ = fi36
                            _ = v37
                        }
                    }
                    switch ("" + "\n\t") {
                    case "foo bar":
                        v38 := GenS2{F0: 2, F1: 88, F2: -818}
                        _ = v38
                    }
                } else {
                    fmt.Println(v15, v1, v2)
                    fmt.Println(v2)
                }
            case "hello":
                var v39 int8 = 1
                _ = v39
            default:
                v1 = v1
            }
            v1 += (v2 - (v2 * (v2 / (((v2 << (((v1 << (((v1 * 866)) & 63))) & 63))) | 1))))
            _ = fi16
        }
        var v40 int8 = 0
        _ = v15
        _ = v40
    } else {
        if ((((float32(float64(0.0)) * 1e10) != 0.0) || false) || false) {
            switch ("foo bar" + ("a" + "a")) {
            case "":
                if len(v3) > 0 {
                    v3[0] = 2
                }
            case "foo bar":
                var v41 uint = (^(v2))
                _ = v41
            default:
                for fi42 := 0; fi42 < 1; fi42++ {
                    var v43 int32 = genFn0()
                    if ((v1 << ((255) & 63)) > 42) {
                        v1 = v1
                    }
                    var v44 uint32 = 2147483648
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            }
            if (2147483647 > 2147483647) {
                v3 = append(v3, 65535)
                switch 111 {
                case -927:
                    switch 791 {
                    case 4294967295:
                        v45 := GenS4{F0: 127}
                        _ = v45
                    }
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v46 := GenS1{F0: 128}
                _ = v46
            }
            switch "hello" {
            case "foo bar":
                var v47 int = 128
                _ = v47
            case "":
                for fi48 := 0; fi48 < 20; fi48++ {
                    v49 := GenS1{F0: 128}
                    v14 = -1
                    _ = fi48
                    _ = v49
                }
            case "\n\t":
                var v50 uint32 = 0
                _ = v50
            default:
                v51 := GenS5{F0: 18446744073709551615, F1: 42}
                _ = v51
            }
        }
        for fi52 := 0; fi52 < 3; fi52++ {
            for fi53 := 0; fi53 < 14; fi53++ {
                var v54 uint64 = 4294967295
                _ = len(v3)
                var v55 float64 = 0.0
                _ = fi53
                _ = v54
                _ = v55
            }
            fmt.Println(v2, v1, v14)
            _ = fi52
        }
    }
    if ((^(v1)) <= 296) {
        if false {
            if ((!false) && true) {
                v56 := GenS3{F0: 255}
                if false {
                    var v57 float32 = ((0.0 / 3.14159) * 0.1)
                    _ = v57
                }
                _ = v56
            }
        } else {
            fmt.Println(v14, v1)
            v3 = append(v3, 88)
            var v58 uint64 = 2
            _ = v58
        }
    }
    var v59 int = 9223372036854775807
    if ((^(v1)) <= 1) {
        var v60 uint64 = 4294967295
        _ = v60
    } else {
        if ((1e10 + -2.25) >= -2.25) {
            var v61 uint8 = 255
            if false {
                var v62 int16 = 42
                _ = v62
            } else {
                var v63 uint16 = 0
                v2 = uint(v63)
                if true {
                    for fi64 := 0; fi64 < 4; fi64++ {
                        var v65 int64 = -9223372036854775808
                        var v66 int64 = (v65 + -9223372036854775808)
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                    for fi67 := 0; fi67 < 5; fi67++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi67
                    }
                    var v68 string = "foo bar"
                    _ = v68
                } else {
                    if true {
                        fmt.Println(v63)
                        if false {
                            if true {
                                var v69 uint32 = 1
                                _ = len(v3)
                                fmt.Println(v14)
                                _ = v69
                            } else {
                                if true {
                                    switch (("" + "") + ("foo bar" + ("hello" + (("a" + "") + "\n\t")))) {
                                    case "\n\t":
                                        switch (((("a" + "hello") + "hello") + "hello") + "\n\t") {
                                        case "a":
                                            v70 := GenS6{F0: 0}
                                            _ = v70
                                        case "hello":
                                            var v71 int16 = v14
                                            _ = v71
                                        case "":
                                            var v72 bool = true
                                            _ = v72
                                        }
                                    }
                                    var v73 uint64 = 1
                                    v74 := GenS5{F0: 0, F1: 1}
                                    _ = v73
                                    _ = v74
                                } else {
                                    switch ("\n\t" + "") {
                                    case "":
                                        v75 := GenS7{F0: false, F1: v14}
                                        _ = v75
                                    case "a":
                                        v76 := GenS8{F0: -9223372036854775808, F1: 42}
                                        _ = v76
                                    }
                                }
                                v59 = (v59 * v59)
                            }
                            _ = len(v3)
                            var v77 uint64 = 2
                            _ = v77
                        } else {
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                            v78 := GenS1{F0: 12}
                            _ = v78
                        }
                        v3 = append(v3, 65535)
                    }
                }
                _ = v63
            }
            _ = v61
        }
        v14 -= 1
        var v79 uint16 = 970
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v59
}
