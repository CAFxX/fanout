package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int8; F1 int8; F2 int32 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS5 struct { F0 string; F1 int64; F2 uint8 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int, p1 uint32, p2 int16) uint32 {
    p0 = p0
    v5 := []int{(-(p0)), p0, 32768, 42}
    var v6 bool = (2147483647 <= -1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p1 >> ((4294967295) & 63))
}
func genFn1(p0 uint16) uint32 {
    v7 := GenS1{F0: 128, F1: 127}
    for fi8 := 0; fi8 < 5; fi8++ {
        switch fi8 {
        case 42:
            switch (("foo bar" + ("foo bar" + "")) + "hello") {
            case "":
                var v9 int64 = -1
                _ = v9
            case "a":
                var v10 uint32 = 4294967295
                _ = v10
            case "\n\t":
                p0 = 2
            default:
                v11 := []int{fi8}
                _ = v11
            }
        case -590:
            v12 := GenS2{F0: -1, F1: -18, F2: 48}
            _ = v12
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := []int{fi13, -2147483648}
            v15 := GenS3{F0: 255, F1: (^(fi13))}
            p0 = 255
            switch (("foo bar" + ("hello" + "\n\t")) + "foo bar") {
            case "hello":
                if (1.5 <= (float64(0.1) * float64(3.14159))) {
                    p0 = (p0 * (^(p0)))
                }
                p0 = (p0 ^ (p0 | p0))
            case "\n\t":
                var v16 int = fi13
                _ = v16
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi8
    }
    var v17 int16 = 32767
    _ = p0
    _ = v17
    _ = v7
    return 65535
}

func main() {
    var v1 int64 = 42
    var v2 string = ("\n\t" + "\n\t")
    v3 := []int{188}
    var v4 float64 = -2.25
    for fi18 := 0; fi18 < 14; fi18++ {
        switch v2 {
        case "foo bar":
            v19 := GenS2{F0: 0, F1: 100, F2: 868}
            _ = v19
        case "hello":
            if (!false) {
                fmt.Println(v2, v1)
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case "\n\t":
            var v20 float64 = (1e10 / v4)
            _ = v20
        default:
            v21 := GenS2{F0: 127, F1: 127, F2: 42}
            _ = v21
        }
        _ = fi18
    }
    if (false || (965 != 4294967295)) {
        switch (v2 + v2) {
        case "foo bar":
            if false {
                switch -9223372036854775808 {
                case 431:
                    for fi22 := 0; fi22 < 13; fi22++ {
                        fmt.Println(v1, v2, v2)
                        _ = fi22
                    }
                case -2147483648:
                    v4 -= v4
                case -851:
                    for fi23 := 0; fi23 < 13; fi23++ {
                        switch fi23 {
                        case -1:
                            if len(v3) > 0 {
                                v3[0] = fi23
                            }
                        case 127:
                            v4 -= v4
                        case 426:
                            v3 = append(v3, (fi23 + (-(fi23))))
                        default:
                            var v24 bool = false
                            _ = v24
                        }
                        if true {
                            if (0 == (v1 % ((v1) | 1))) {
                                var v25 uint = 4294967295
                                var v26 bool = false
                                v4 += v4
                                _ = v25
                                _ = v26
                            }
                            if false {
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = fi23
                                    }
                                }
                                var v27 bool = (4294967295 >= 4294967295)
                                _ = v27
                            }
                        } else {
                            v28 := GenS4{F0: 389, F1: 245, F2: 18446744073709551615}
                            v1 -= 1
                            for fi29 := 0; fi29 < 7; fi29++ {
                                v30 := GenS2{F0: 42, F1: -128, F2: -698}
                                _ = fi29
                                _ = v30
                            }
                            _ = v28
                        }
                        _ = fi23
                    }
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                for fi31 := 0; fi31 < 12; fi31++ {
                    v4 = v4
                    for fi32 := 0; fi32 < 6; fi32++ {
                        fmt.Println(v2)
                        _ = fi32
                    }
                    v33 := GenS4{F0: 1, F1: 1, F2: 42}
                    _ = fi31
                    _ = v33
                }
                switch v2 {
                case "":
                    var v34 int = 652
                    _ = v34
                case "hello":
                    if false {
                        var v35 bool = (1000 != 1)
                        _ = v35
                    }
                case "a":
                    v4 = v4
                }
            } else {
                switch ("" + "\n\t") {
                case "":
                    switch -9223372036854775808 {
                    case 128:
                        var v36 int = -2147483648
                        _ = v36
                    case 127:
                        fmt.Println(v2, v1, v2)
                    case 2147483647:
                        v37 := GenS1{F0: 2, F1: 938}
                        _ = v37
                    default:
                        for fi38 := 0; fi38 < 1; fi38++ {
                            v1 -= 0
                            _ = fi38
                        }
                    }
                }
            }
        case "hello":
            var v39 int16 = -415
            _ = v39
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v3)
    }
    if (1 != -695) {
        if false {
            v40 := GenS1{F0: 182, F1: 492}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v40
        } else {
            for fi41 := 0; fi41 < 16; fi41++ {
                v2 = v2
                if (false && (true && true)) {
                    if (!(!false)) {
                        switch fi41 {
                        case 203:
                            var v42 uint8 = 255
                            _ = v42
                        }
                        if false {
                            var v43 int64 = 42
                            _ = v43
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(fi41, v2, v2)
                            var v44 int64 = v1
                            _ = v44
                        }
                        var v45 int64 = 679
                        _ = v45
                    }
                    fmt.Println(v2)
                    v46 := GenS5{F0: ((v2 + (v2 + "foo bar")) + ""), F1: v1, F2: 151}
                    _ = v46
                } else {
                    v47 := GenS3{F0: 42, F1: fi41}
                    if ((false && true) && false) {
                        var v48 uint = 18446744073709551615
                        _ = v48
                    }
                    v2 = v2
                    _ = v47
                }
                _ = fi41
            }
        }
    }
    for fi49 := 0; fi49 < 19; fi49++ {
        if true {
            v50 := GenS6{F0: v1}
            for fi51 := 0; fi51 < 2; fi51++ {
                fmt.Println(v2, fi49, fi49)
                _ = fi51
            }
            _ = v50
        } else {
            v1 = v1
            v52 := GenS5{F0: v2, F1: 9223372036854775807, F2: 1}
            _ = v52
        }
        if true {
            if false {
                if (v1 != 149) {
                    if len(v3) > 0 {
                        v3[0] = (fi49 / (((fi49 >> ((0) & 63))) | 1))
                    }
                }
                var v53 uint16 = 255
                _ = v53
            } else {
                switch "" {
                case "":
                    for fi54 := 0; fi54 < 7; fi54++ {
                        switch "foo bar" {
                        case "foo bar":
                            var v55 string = v2
                            _ = v55
                        default:
                            switch fi49 {
                            case 0:
                                fmt.Println(v4, fi49)
                            case -249:
                                v4 += 3.14159
                            }
                        }
                        _ = fi54
                    }
                case "hello":
                    if (!(100 < -1)) {
                        for fi56 := 0; fi56 < 2; fi56++ {
                            v3 = append(v3, 4294967295)
                            _ = fi56
                        }
                        v1 = (^(v1))
                    }
                default:
                    if (255 >= 546) {
                        var v57 float64 = 3.14159
                        v1 = v1
                        _ = v57
                    } else {
                        v58 := GenS2{F0: 42, F1: 42, F2: int32(v1)}
                        v4 = v4
                        _ = v58
                    }
                }
            }
        } else {
            var v59 int8 = 100
            if true {
                if false {
                    var v60 uint16 = 1
                    fmt.Println(v4)
                    for fi61 := 0; fi61 < 14; fi61++ {
                        var v62 int32 = -1
                        v2 = "foo bar"
                        _ = fi61
                        _ = v62
                    }
                    _ = v60
                }
                var v63 uint64 = 2
                _ = v63
            }
            _ = v59
        }
        _ = fi49
    }
    var v64 int8 = 100
    v3 = append(v3, -9223372036854775808)
    if len(v3) > 0 {
        v3[0] = 32767
    }
    v65 := GenS6{F0: (-(v1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
}
