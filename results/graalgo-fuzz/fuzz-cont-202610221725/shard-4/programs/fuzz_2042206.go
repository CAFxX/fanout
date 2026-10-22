package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int8; F1 int; F2 int8 }
type GenS7 struct { F0 int32; F1 int32; F2 int8 }
type GenS8 struct { F0 uint32; F1 uint }

func main() {
    var v1 uint16 = 255
    var v2 uint = 18446744073709551615
    v3 := []int{4294967295}
    v4 := []uint32{}
    var v5 float32 = ((0.1 + (3.14159 / 0.1)) + float32(-2.25))
    var v6 uint16 = 2
    if false {
        v2 -= 76
    } else {
        if true {
            v6 = (^(v6))
            for fi7 := 0; fi7 < 19; fi7++ {
                for fi8 := 0; fi8 < 15; fi8++ {
                    v6 += (^(v1))
                    fmt.Println(v2)
                    if true {
                        var v9 uint32 = 963
                        if (int64(v2) > -760) {
                            v10 := GenS1{F0: 42}
                            v2 = v2
                            _ = v10
                        }
                        _ = v9
                    } else {
                        var v11 bool = (1 <= 42)
                        _ = v11
                    }
                    _ = fi8
                }
                _ = fi7
            }
            v6 = (v6 + v6)
        } else {
            if (-609 != -587) {
                var v12 int8 = -128
                v13 := GenS1{F0: 127}
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS1{F0: 154}
                    _ = fi14
                    _ = v15
                }
                _ = v12
                _ = v13
            }
            v6 = (v6 / ((255) | 1))
        }
        var v16 float32 = v5
        v6 = (v1 | v6)
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        v1 += v1
        switch fi17 {
        case 635:
            if (250 > 2147483647) {
                for fi18 := 0; fi18 < 13; fi18++ {
                    fmt.Println(v5, v5, v6)
                    if false {
                        v6 += v1
                    } else {
                        var v19 uint64 = 546
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if true {
                            if len(v4) > 0 {
                                v4[0] = 4294967295
                            }
                            v6 -= v6
                            fmt.Println(v6)
                        }
                        _ = v19
                    }
                    v2 = 792
                    _ = fi18
                }
            } else {
                var v20 float32 = float32(3.14159)
                var v21 string = "hello"
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
                _ = v20
                _ = v21
            }
        default:
            var v22 int16 = 32767
            _ = v22
        }
        _ = fi17
    }
    var v23 int16 = 796
    v3 = append(v3, 32767)
    if false {
        v24 := GenS2{F0: 0, F1: 1}
        if (true && (int64(v2) != 42)) {
            for fi25 := 0; fi25 < 7; fi25++ {
                v26 := GenS2{F0: 0, F1: 18446744073709551615}
                if (!(757 != 42)) {
                    var v27 int = fi25
                    switch 42 {
                    case 2147483647:
                        for fi28 := 0; fi28 < 19; fi28++ {
                            switch -186 {
                            case 9223372036854775807:
                                v29 := GenS3{F0: v2}
                                if true {
                                    if false {
                                        var v30 uint8 = 42
                                        _ = v30
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        var v31 float32 = 3.14159
                                        _ = v31
                                    case "":
                                        v32 := GenS4{F0: 127, F1: 1}
                                        var v33 uint8 = 255
                                        if true {
                                            var v34 uint = (v2 * 255)
                                            v35 := GenS2{F0: 567, F1: 543}
                                            var v36 int8 = 127
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                        }
                                        fmt.Println(v6, v1)
                                        if false {
                                            v1 += (v6 | 255)
                                        }
                                        v27 = fi25
                                        var v37 int = 1
                                        v37 = 4294967295
                                        _ = v32
                                        _ = v33
                                        _ = v37
                                    }
                                    if (!(false || (true || (uint32(v2) != 2)))) {
                                        fmt.Println(v5, fi25)
                                        v38 := GenS1{F0: 171}
                                        _ = v38
                                    } else {
                                        var v39 uint16 = 239
                                        _ = v39
                                    }
                                    v40 := GenS3{F0: (v2 - (v2 ^ v2))}
                                    _ = v40
                                }
                                if true {
                                    v41 := GenS1{F0: 1}
                                    _ = v41
                                } else {
                                    if true {
                                        _ = len(v3)
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        fmt.Println(v5, v6)
                                        if len(v4) > 0 {
                                            v4[0] = 2147483647
                                        }
                                    }
                                    var v42 uint64 = 2
                                    if false {
                                        var v43 string = ""
                                        _ = v43
                                    } else {
                                        v4 = append(v4, uint32(v1))
                                        v4 = append(v4, 4294967295)
                                        fmt.Println(v42, v27, v23)
                                        switch fi25 {
                                        case 42:
                                            v44 := GenS5{F0: 42, F1: (v2 | 65535)}
                                            _ = v44
                                        }
                                        v27 -= (v27 - 979)
                                    }
                                    _ = v42
                                }
                                v23 += (v23 * -1)
                                fmt.Println(v23, fi28, fi28)
                                _ = v29
                            case -606:
                                fmt.Println(v6)
                            case 471:
                                switch "\n\t" {
                                case "a":
                                    v5 += v5
                                case "foo bar":
                                    v6 -= 0
                                case "":
                                    v5 += (float32(0.0) * 1.5)
                                default:
                                    var v45 int = 720
                                    if true {
                                        var v46 int8 = int8(fi25)
                                        _ = v46
                                    }
                                    if (4294967295 <= 346) {
                                        v27 -= (fi28 / ((v27) | 1))
                                    } else {
                                        _ = len(v3)
                                    }
                                    _ = v45
                                }
                            }
                            _ = fi28
                        }
                    }
                    _ = v27
                } else {
                    for fi47 := 0; fi47 < 6; fi47++ {
                        v5 = v5
                        _ = fi47
                    }
                }
                var v48 int8 = -128
                _ = fi25
                _ = v26
                _ = v48
            }
        } else {
            fmt.Println(v1, v1, v2)
            var v49 int64 = -1
            if false {
                switch -647 {
                case 32767:
                    switch "foo bar" {
                    case "hello":
                        var v50 uint = (v2 / (((^(v2))) | 1))
                        _ = v50
                    case "\n\t":
                        fmt.Println(v1)
                    case "foo bar":
                        switch -9223372036854775808 {
                        case -64:
                            for fi51 := 0; fi51 < 4; fi51++ {
                                fmt.Println(v49)
                                for fi52 := 0; fi52 < 15; fi52++ {
                                    if len(v4) > 0 {
                                        v4[0] = 495
                                    }
                                    v6 += (^(v6))
                                    var v53 int = fi51
                                    _ = fi52
                                    _ = v53
                                }
                                fmt.Println(v5, v23)
                                _ = fi51
                            }
                        default:
                            v54 := GenS6{F0: -34, F1: -319, F2: 42}
                            _ = v54
                        }
                    default:
                        v1 += (v6 & (v1 % ((v6) | 1)))
                    }
                case -739:
                    if (("a" + "") < ("" + "foo bar")) {
                        var v55 uint8 = 127
                        var v56 bool = ((!(!false)) && true)
                        _ = v55
                        _ = v56
                    } else {
                        if ((-794 >= -9223372036854775808) && false) {
                            v6 = v6
                        } else {
                            var v57 uint = 2
                            _ = v57
                        }
                        switch "a" {
                        case "\n\t":
                            if false {
                                var v58 uint8 = 0
                                _ = v58
                            } else {
                                v2 = v2
                                var v59 float32 = (1e10 * float32(3.14159))
                                var v60 uint = v2
                                _ = v59
                                _ = v60
                            }
                        case "a":
                            var v61 uint32 = 1
                            _ = v61
                        case "hello":
                            var v62 bool = (1.5 == float64(v5))
                            _ = v62
                        default:
                            switch 21 {
                            case 32768:
                                var v63 uint8 = 234
                                _ = v63
                            case 42:
                                if false {
                                    var v64 int8 = -113
                                    _ = v64
                                }
                            default:
                                v6 -= v1
                            }
                        }
                    }
                default:
                    if ((true || ((true && (((float32(0.0) / v5) - v5) < v5)) && (2 <= 663))) && false) {
                        var v65 uint64 = 1
                        for fi66 := 0; fi66 < 19; fi66++ {
                            v2 -= uint(v65)
                            _ = fi66
                        }
                        _ = v65
                    } else {
                        var v67 uint8 = 255
                        _ = v67
                    }
                }
            }
            _ = v49
        }
        v68 := GenS7{F0: 496, F1: 2147483647, F2: 1}
        _ = v24
        _ = v68
    } else {
        for fi69 := 0; fi69 < 4; fi69++ {
            var v70 bool = (1.5 >= (float64(1.5) - (-2.25 - 0.0)))
            _ = fi69
            _ = v70
        }
    }
    var v71 uint64 = 530
    v72 := GenS8{F0: 4294967295, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v72
}
