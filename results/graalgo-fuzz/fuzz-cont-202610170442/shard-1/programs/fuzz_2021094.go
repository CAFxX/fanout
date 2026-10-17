package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int16; F2 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS8 struct { F0 int16 }
func genFn0() bool {
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{fi5}
        _ = fi5
        _ = v6
    }
    var v7 float64 = (1.5 * -2.25)
    for fi8 := 0; fi8 < 18; fi8++ {
        v7 += v7
        v9 := []int{fi8, -2147483648, 605, fi8}
        _ = fi8
        _ = v9
    }
    _ = v7
    return (1 < 42)
}

func main() {
    var v1 uint8 = 1
    var v2 string = "foo bar"
    v3 := make(map[string]string)
    var v4 bool = (9223372036854775807 != 32768)
    v10 := GenS1{F0: 818}
    var v11 int = -752
    v12 := GenS1{F0: 42}
    for fi13 := 0; fi13 < 13; fi13++ {
        v1 = v1
        _ = fi13
    }
    var v14 uint64 = 2
    var v15 float32 = 0.0
    v16 := []int{v11, (v11 << ((2) & 63))}
    v11 = 32767
    var v17 int8 = 118
    var v18 int32 = 2147483647
    for fi19 := 0; fi19 < 10; fi19++ {
        if (v4 && (v4 || v4)) {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
        if true {
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 float64 = 1.5
                var v22 string = v2
                v23 := GenS1{F0: (v18 >> ((22) & 63))}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            v24 := v3["hello"]
            if len(v16) > 0 {
                v16[0] = fi19
            }
            _ = v24
        }
        _ = fi19
    }
    if (42 == v11) {
        for fi25 := 0; fi25 < 11; fi25++ {
            v1 += v1
            if true {
                switch (^(fi25)) {
                case 127:
                    var v26 int64 = 656
                    _ = v26
                case 1:
                    var v27 int = fi25
                    _ = v27
                case 32767:
                    var v28 int32 = (v18 + (v18 % ((int32(v17)) | 1)))
                    _ = v28
                }
                for fi29 := 0; fi29 < 2; fi29++ {
                    if true {
                        v30 := GenS1{F0: v18}
                        if true {
                            var v31 float32 = v15
                            _ = v31
                        } else {
                            fmt.Println(v15, fi25)
                        }
                        if v4 {
                            var v32 uint8 = (v1 & v1)
                            if (v4 && v4) {
                                _ = len(v16)
                            }
                            v16 = append(v16, fi29)
                            _ = v32
                        } else {
                            var v33 int8 = (v17 + -1)
                            if v4 {
                                var v34 int16 = 464
                                _ = v34
                            } else {
                                v15 = float32(3.14159)
                            }
                            _ = v33
                        }
                        _ = v30
                    }
                    _ = fi29
                }
                if v4 {
                    if ((v17 >> ((18446744073709551615) & 63)) <= 100) {
                        if v4 {
                            v14 = (v14 << ((1) & 63))
                        } else {
                            var v35 string = ("a" + v2)
                            _ = v35
                        }
                    } else {
                        for fi36 := 0; fi36 < 10; fi36++ {
                            switch "\n\t" {
                            case "":
                                switch (v2 + v2) {
                                case "foo bar":
                                    fmt.Println(v14, v14, v14)
                                case "\n\t":
                                    fmt.Println(v2, fi36, v17)
                                case "":
                                    switch v2 {
                                    case "a":
                                        v14 -= 698
                                    case "\n\t":
                                        v18 += -499
                                    case "hello":
                                        switch ("foo bar" + "") {
                                        case "foo bar":
                                            var v37 uint8 = v1
                                            _ = v37
                                        case "":
                                            var v38 uint32 = 2147483647
                                            _ = v38
                                        default:
                                            v39 := GenS1{F0: v18}
                                            _ = v39
                                        }
                                    default:
                                        if ((v2 < v2) && v4) {
                                            fmt.Println(v4, fi25, v15)
                                            var v40 int16 = -132
                                            _ = v40
                                        } else {
                                            fmt.Println(v18)
                                            fmt.Println(v18, v2, v14)
                                            v14 -= v14
                                        }
                                    }
                                }
                            case "foo bar":
                                switch (("" + v2) + v2) {
                                case "":
                                    v41 := GenS1{F0: (v18 + (v18 - (v18 * int32(fi36))))}
                                    _ = v41
                                case "a":
                                    v4 = true
                                }
                            default:
                                switch ("a" + "foo bar") {
                                case "\n\t":
                                    switch fi36 {
                                    case 32768:
                                        _ = len(v16)
                                        var v42 uint8 = v1
                                        _ = v42
                                    case 127:
                                        v43 := GenS1{F0: v18}
                                        _ = v43
                                    case 4294967295:
                                        v3["hello"] = v2
                                    default:
                                        var v44 uint64 = v14
                                        _ = v44
                                    }
                                default:
                                    v4 = (2147483648 <= 979)
                                }
                            }
                            var v45 uint16 = 32767
                            _ = fi36
                            _ = v45
                        }
                        v1 = v1
                        v4 = ((v17 - (v17 ^ v17)) <= -128)
                    }
                    for fi46 := 0; fi46 < 7; fi46++ {
                        if v4 {
                            v47 := GenS1{F0: v18}
                            _ = v47
                        }
                        _ = fi46
                    }
                    var v48 float64 = (1.5 - (float64((3.14159 / (float32((-2.25 - 1.5)) - (v15 * (v15 * 1e10))))) * (1.5 + 3.14159)))
                    _ = v48
                }
            } else {
                switch (fi25 | (-(v11))) {
                case 128:
                    var v49 float32 = ((((3.14159 + (0.0 - (3.14159 + ((v15 * v15) / v15)))) * v15) + float32((-2.25 - 0.1))) + 0.1)
                    _ = v49
                default:
                    _ = len(v16)
                }
            }
            if v4 {
                var v50 float64 = -2.25
                if (!true) {
                    if (false && ((v15 < 1e10) && (v14 >= 1))) {
                        if len(v16) > 0 {
                            v16[0] = 952
                        }
                    }
                } else {
                    v2 = (("a" + "a") + v2)
                    v51 := GenS2{F0: v1, F1: v17}
                    var v52 int32 = (v18 * v18)
                    _ = v51
                    _ = v52
                }
                _ = v50
            }
            _ = fi25
        }
        switch (v11 - (^(v11))) {
        case 1:
            for fi53 := 0; fi53 < 17; fi53++ {
                v18 += -710
                _ = fi53
            }
        default:
            var v54 int32 = 553
            _ = v54
        }
        var v55 float32 = ((float32(0.1) / v15) * 3.14159)
        _ = v55
    } else {
        switch ("foo bar" + ("" + v2)) {
        case "a":
            v56 := GenS3{F0: 4294967295, F1: (v14 + 1), F2: 1}
            _ = v56
        case "foo bar":
            for fi57 := 0; fi57 < 3; fi57++ {
                v15 -= v15
                switch v2 {
                case "":
                    v58 := GenS4{F0: 2}
                    _ = v58
                case "\n\t":
                    switch v2 {
                    case "hello":
                        v14 = 505
                    case "":
                        for fi59 := 0; fi59 < 1; fi59++ {
                            v60 := GenS5{F0: (-(v17)), F1: int16(v17), F2: 1}
                            v18 = int32(v14)
                            _ = fi59
                            _ = v60
                        }
                    case "foo bar":
                        var v61 uint64 = v14
                        _ = v61
                    }
                case "foo bar":
                    if v4 {
                        v1 = (v1 - v1)
                    } else {
                        if ((v14 & 529) == v14) {
                            if (!(!(65535 < 65535))) {
                                v14 = (v14 | v14)
                                v4 = v4
                            }
                            v15 = v15
                            switch 32767 {
                            case 4294967295:
                                delete(v3, v2)
                            default:
                                v2 = v2
                            }
                        } else {
                            v62 := GenS3{F0: 65535, F1: (v14 | v14), F2: 0}
                            v4 = false
                            _ = v62
                        }
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    }
                default:
                    fmt.Println(v11)
                }
                v2 = "\n\t"
                _ = fi57
            }
        default:
            v63 := GenS6{F0: 2}
            _ = v63
        }
        _ = len(v16)
    }
    switch (((v2 + v2) + v2) + (v2 + (v2 + v2))) {
    case "foo bar":
        v18 = (v18 % ((-2147483648) | 1))
    case "a":
        v64 := GenS7{F0: 32767, F1: 65535, F2: (^(v14))}
        _ = v64
    default:
        for fi65 := 0; fi65 < 7; fi65++ {
            v66 := GenS8{F0: -32768}
            v2 = ("hello" + v2)
            _ = fi65
            _ = v66
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
