package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint; F1 int64; F2 int16 }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint; F1 uint32 }

func main() {
    var v1 float32 = ((3.14159 - (-2.25 + -2.25)) / 1.0)
    v2 := []int{1, 0, 4294967295}
    var v3 uint32 = 65535
    var v4 int64 = -410
    v5 := GenS1{F0: 65535, F1: 799, F2: 0}
    fmt.Println(v1)
    switch -756 {
    case 1:
        fmt.Println(v4, v4)
    case 2147483647:
        switch ((("foo bar" + "") + "a") + "\n\t") {
        case "\n\t":
            switch 884 {
            case -963:
                for fi6 := 0; fi6 < 7; fi6++ {
                    switch "\n\t" {
                    case "hello":
                        v1 -= v1
                    }
                    switch fi6 {
                    case 462:
                        var v7 uint32 = v3
                        _ = v7
                    }
                    _ = fi6
                }
            case 0:
                _ = len(v2)
            default:
                for fi8 := 0; fi8 < 13; fi8++ {
                    if (0.1 != (v1 - v1)) {
                        if true {
                            v4 = (v4 % ((-9223372036854775808) | 1))
                            fmt.Println(v4)
                            v9 := GenS1{F0: 979, F1: 42, F2: 1}
                            _ = v9
                        } else {
                            _ = len(v2)
                            v2 = append(v2, -1)
                        }
                        v2 = append(v2, (fi8 & (fi8 + (fi8 % ((fi8) | 1)))))
                        for fi10 := 0; fi10 < 2; fi10++ {
                            var v11 string = "a"
                            switch (fi8 + (fi8 & fi8)) {
                            case -9223372036854775808:
                                fmt.Println(v11, v1)
                            default:
                                switch v11 {
                                case "hello":
                                    if (127 > -128) {
                                        v4 -= (v4 & (v4 % (((-(v4))) | 1)))
                                        if (4294967295 <= 1) {
                                            switch fi8 {
                                            case -466:
                                                v1 -= -2.25
                                            case 223:
                                                var v12 uint64 = 18446744073709551615
                                                _ = v12
                                            default:
                                                var v13 int8 = 1
                                                _ = v13
                                            }
                                            var v14 uint16 = 948
                                            var v15 uint64 = 4294967295
                                            switch 4294967295 {
                                            case 9223372036854775807:
                                                var v16 uint16 = (v14 ^ v14)
                                                _ = v16
                                            case -748:
                                                fmt.Println(v15, v1)
                                            case -726:
                                                var v17 uint32 = v3
                                                _ = v17
                                            default:
                                                if true {
                                                    v4 = v4
                                                } else {
                                                    v14 += (^(v14))
                                                    v18 := GenS2{F0: v11, F1: v14}
                                                    fmt.Println(v11)
                                                    _ = v18
                                                }
                                            }
                                            _ = v14
                                            _ = v15
                                        } else {
                                            v4 += v4
                                            fmt.Println(v1, fi10)
                                            v2 = append(v2, 594)
                                        }
                                    } else {
                                        v3 = (^(v3))
                                        v3 = (v3 % (((v3 & v3)) | 1))
                                        if false {
                                            v1 -= v1
                                        }
                                    }
                                }
                                if false {
                                    var v19 uint32 = v3
                                    v3 += (v19 >> ((0) & 63))
                                    v1 = 1.5
                                    fmt.Println(v4)
                                    if (255 <= 127) {
                                        if len(v2) > 0 {
                                            v2[0] = (fi8 << ((394) & 63))
                                        }
                                        v1 += (float32(((0.1 + (1e10 - float64(v1))) + (3.14159 / 3.14159))) * (v1 * float32(1.5)))
                                    } else {
                                        v20 := GenS1{F0: 255, F1: 2147483647, F2: v4}
                                        v4 += 1
                                        v11 = "\n\t"
                                        var v21 uint32 = v3
                                        _ = v20
                                        _ = v21
                                    }
                                    _ = v19
                                }
                            }
                            _ = fi10
                            _ = v11
                        }
                    }
                    for fi22 := 0; fi22 < 10; fi22++ {
                        v23 := GenS3{F0: 42}
                        switch fi8 {
                        case 687:
                            var v24 float32 = v1
                            _ = v24
                        case -2147483648:
                            v25 := GenS4{F0: 42, F1: 42}
                            _ = v25
                        case -315:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            v2 = append(v2, fi8)
                        }
                        _ = fi22
                        _ = v23
                    }
                    _ = fi8
                }
            }
        case "":
            if (!false) {
                var v26 int64 = v4
                v3 = v3
                v27 := GenS4{F0: -1, F1: -128}
                _ = v26
                _ = v27
            }
        case "hello":
            if false {
                v2 = append(v2, -263)
                _ = len(v2)
                fmt.Println(v4)
            }
        default:
            v3 = (v3 / ((v3) | 1))
        }
    default:
        v4 -= (v4 * (^(v4)))
    }
    v28 := GenS5{F0: 388}
    v3 = (v3 >> ((65535) & 63))
    for fi29 := 0; fi29 < 3; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            var v31 uint = 1
            if true {
                fmt.Println(v3, v1, fi29)
                v3 -= 825
            }
            _ = fi30
            _ = v31
        }
        v32 := GenS1{F0: 2, F1: 212, F2: (v4 | v4)}
        v2 = append(v2, 32767)
        _ = fi29
        _ = v32
    }
    if (123 >= 71) {
        var v33 int32 = -1
        _ = v33
    } else {
        var v34 bool = false
        v35 := GenS5{F0: 0}
        if v34 {
            if ((0.0 != (0.1 / (((3.14159 + (3.14159 - 1e10)) - float64((v1 / v1))) * 0.1))) && (1 != 0)) {
                switch 9223372036854775807 {
                case -9223372036854775808:
                    if v34 {
                        v1 = (v1 - ((v1 * v1) * v1))
                        switch "a" {
                        case "\n\t":
                            switch 32768 {
                            case -192:
                                if len(v2) > 0 {
                                    v2[0] = 46
                                }
                            }
                        case "hello":
                            if false {
                                v3 += v3
                            } else {
                                v4 = (v4 - v4)
                            }
                        case "a":
                            fmt.Println(v4, v34)
                        default:
                            var v36 uint32 = v3
                            _ = v36
                        }
                    }
                }
                for fi37 := 0; fi37 < 10; fi37++ {
                    v1 -= (0.1 / v1)
                    var v38 string = ("foo bar" + "hello")
                    fmt.Println(v4, fi37, v38)
                    _ = fi37
                    _ = v38
                }
                var v39 uint16 = 65535
                _ = v39
            }
        } else {
            var v40 uint16 = 151
            if ((false || (((("\n\t" + "foo bar") + ("\n\t" + "foo bar")) + ("foo bar" + "foo bar")) > "\n\t")) || true) {
                var v41 int8 = 1
                if len(v2) > 0 {
                    v2[0] = -830
                }
                _ = v41
            }
            var v42 uint16 = (^(v40))
            _ = v40
            _ = v42
        }
        _ = v34
        _ = v35
    }
    var v43 bool = true
    for fi44 := 0; fi44 < 16; fi44++ {
        var v45 int8 = 127
        fmt.Println(v43, v43)
        v46 := GenS2{F0: (("hello" + "") + ("hello" + "\n\t")), F1: 457}
        _ = fi44
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 4; fi47++ {
        var v48 string = "foo bar"
        v48 = (v48 + v48)
        if ("\n\t" > "hello") {
            for fi49 := 0; fi49 < 12; fi49++ {
                if len(v2) > 0 {
                    v2[0] = fi47
                }
                fmt.Println(v48, v4)
                _ = fi49
            }
        } else {
            fmt.Println(v48, v43, v3)
            fmt.Println(v4, v48)
            v3 += 2147483648
        }
        _ = fi47
        _ = v48
    }
    if (float64(v1) >= 0.1) {
        v4 = v4
        fmt.Println(v1, v4, v3)
    } else {
        var v50 uint16 = 32767
        v1 += (1.5 / v1)
        if len(v2) > 0 {
            v2[0] = -696
        }
        _ = v50
    }
    for fi51 := 0; fi51 < 5; fi51++ {
        v52 := GenS6{F0: 249, F1: 4294967295}
        if v43 {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v53 := GenS7{F0: 65535, F1: v4, F2: 746}
            _ = v53
        }
        var v54 string = ("a" + ("foo bar" + ("foo bar" + ("a" + "a"))))
        _ = fi51
        _ = v52
        _ = v54
    }
    switch 127 {
    case 9223372036854775807:
        var v55 int8 = 127
        _ = v55
    case 525:
        for fi56 := 0; fi56 < 3; fi56++ {
            v57 := GenS8{F0: 603, F1: 2147483647}
            if (!(128 > 0)) {
                var v58 uint64 = 42
                var v59 float64 = 3.14159
                if v43 {
                    v58 = (^(v58))
                } else {
                    v60 := GenS9{F0: -848}
                    v61 := GenS10{F0: 963, F1: 42}
                    _ = v60
                    _ = v61
                }
                _ = v58
                _ = v59
            }
            _ = fi56
            _ = v57
        }
    case 42:
        if ((v4 | (v4 ^ v4)) > (v4 & (v4 ^ v4))) {
            var v62 int8 = 112
            v63 := GenS2{F0: ("hello" + "hello"), F1: 0}
            _ = v62
            _ = v63
        }
    default:
        if ((((float32((3.14159 * (0.0 + 1.5))) / ((1.5 + -2.25) / ((v1 + ((float32(3.14159) - (v1 - v1)) / v1)) / v1))) + float32(3.14159)) < (((((v1 / float32(-2.25)) / -2.25) * v1) - v1) + v1)) || (!((-1 > -1) || (!v43)))) {
            v64 := GenS4{F0: 42, F1: -1}
            for fi65 := 0; fi65 < 8; fi65++ {
                var v66 bool = ((0 >= 42) && false)
                _ = fi65
                _ = v66
            }
            var v67 uint64 = 229
            _ = v64
            _ = v67
        } else {
            fmt.Println(v3)
        }
    }
    for fi68 := 0; fi68 < 17; fi68++ {
        switch "foo bar" {
        case "":
            v4 += (v4 | 23)
        case "\n\t":
            var v69 uint8 = 1
            _ = v69
        case "foo bar":
            fmt.Println(fi68)
        default:
            var v70 string = "foo bar"
            _ = v70
        }
        v2 = append(v2, 128)
        fmt.Println(v4, v3)
        _ = fi68
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
