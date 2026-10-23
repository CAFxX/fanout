package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 int32 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 string; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int16; F2 int16 }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8; F1 string }

func main() {
    var v1 uint = 2
    v2 := []uint32{}
    var v3 uint16 = 32767
    var v4 uint32 = 2
    var v5 string = "hello"
    for fi6 := 0; fi6 < 9; fi6++ {
        v3 += v3
        if (!true) {
            v1 -= (v1 >> ((uint(v3)) & 63))
            var v7 uint8 = 42
            for fi8 := 0; fi8 < 4; fi8++ {
                v3 = v3
                _ = len(v2)
                _ = fi8
            }
            _ = v7
        } else {
            var v9 int32 = 12
            v10 := GenS1{F0: -1, F1: -128}
            for fi11 := 0; fi11 < 11; fi11++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v1 = (v1 - 717)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v9 -= (v9 >> ((v1) & 63))
                _ = fi11
            }
            _ = v10
            _ = v9
        }
        if true {
            if len(v2) > 0 {
                v2[0] = uint32(v1)
            }
            var v12 int16 = 1
            switch fi6 {
            case 491:
                if (!((1 < 42) || (1 >= 2))) {
                    switch (fi6 >> (((v1 % ((803) | 1))) & 63)) {
                    case 4294967295:
                        if false {
                            for fi13 := 0; fi13 < 4; fi13++ {
                                switch (fi6 | fi13) {
                                case 32768:
                                    v3 -= (v3 | 733)
                                case 32767:
                                    var v14 float32 = 3.14159
                                    _ = v14
                                }
                                _ = fi13
                            }
                            fmt.Println(fi6, v1)
                            var v15 uint64 = 1
                            _ = v15
                        } else {
                            if true {
                                if (false && (false && (656 >= v4))) {
                                    fmt.Println(v4, v12)
                                    v16 := GenS1{F0: -1, F1: 1}
                                    _ = v16
                                }
                            }
                        }
                    default:
                        for fi17 := 0; fi17 < 20; fi17++ {
                            switch (fi17 ^ (fi17 | (fi17 >> ((920) & 63)))) {
                            case 9223372036854775807:
                                v3 = (v3 + uint16(v4))
                            }
                            if (false || true) {
                                v18 := GenS1{F0: (v12 + 505), F1: -58}
                                _ = v18
                            }
                            _ = fi17
                        }
                    }
                    var v19 float64 = (-2.25 * -2.25)
                    v12 -= (-(v12))
                    _ = v19
                } else {
                    var v20 int = (-(fi6))
                    switch 0 {
                    case 32767:
                        _ = len(v2)
                    case 1:
                        fmt.Println(v12, v12)
                    }
                    if true {
                        if (!false) {
                            if true {
                                v21 := GenS2{F0: v4, F1: (v20 * (-(fi6))), F2: -2147483648}
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    if false {
                                        v23 := GenS2{F0: v4, F1: (v20 << ((v1) & 63)), F2: 605}
                                        var v24 uint64 = 1
                                        var v25 float64 = 0.1
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                    _ = fi22
                                }
                                _ = v21
                            } else {
                                v26 := GenS2{F0: v4, F1: 127, F2: 194}
                                v27 := GenS3{F0: 646, F1: v12}
                                v2 = append(v2, 42)
                                _ = v26
                                _ = v27
                            }
                            var v28 int8 = int8(v1)
                            _ = v28
                        }
                    } else {
                        v29 := GenS4{F0: (("" + (v5 + v5)) + "a"), F1: 0, F2: 399}
                        var v30 uint = 396
                        _ = v29
                        _ = v30
                    }
                    _ = v20
                }
            case 489:
                fmt.Println(v12, v12, v3)
            case -235:
                v5 = v5
            default:
                v1 = uint(v4)
            }
            _ = v12
        } else {
            var v31 uint64 = 4294967295
            v4 = v4
            switch fi6 {
            case 0:
                v2 = append(v2, v4)
            case -1:
                v32 := GenS2{F0: v4, F1: fi6, F2: -1}
                _ = v32
            case 32767:
                switch -1 {
                case 0:
                    v5 = v5
                case 32767:
                    v33 := GenS5{F0: v1, F1: 2147483647}
                    _ = v33
                case 32768:
                    if len(v2) > 0 {
                        v2[0] = (v4 << (((^(v1))) & 63))
                    }
                default:
                    for fi34 := 0; fi34 < 11; fi34++ {
                        if len(v2) > 0 {
                            v2[0] = v4
                        }
                        _ = fi34
                    }
                }
            }
            _ = v31
        }
        _ = fi6
    }
    v5 = (((("hello" + v5) + (v5 + "foo bar")) + v5) + ((v5 + v5) + (v5 + (v5 + "hello"))))
    for fi35 := 0; fi35 < 8; fi35++ {
        v3 -= v3
        _ = fi35
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 int = (fi36 | -403)
        _ = fi36
        _ = v37
    }
    if (736 > -1) {
        var v38 uint32 = 42
        for fi39 := 0; fi39 < 6; fi39++ {
            v38 -= (v38 >> ((42) & 63))
            if true {
                switch 128 {
                case 127:
                    if (-128 <= 100) {
                        v1 = (^(v1))
                        if (!(0 <= -32768)) {
                            for fi40 := 0; fi40 < 12; fi40++ {
                                switch v5 {
                                case "foo bar":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "":
                                    var v41 string = v5
                                    _ = v41
                                case "a":
                                    switch (fi39 | (fi40 * (^(fi40)))) {
                                    case -9223372036854775808:
                                        var v42 float32 = float32(3.14159)
                                        _ = v42
                                    default:
                                        v3 -= (v3 * v3)
                                    }
                                default:
                                    fmt.Println(fi39, v38, fi39)
                                }
                                var v43 int64 = 0
                                _ = fi40
                                _ = v43
                            }
                            fmt.Println(v5, v38, v1)
                            var v44 int16 = 1
                            _ = v44
                        } else {
                            var v45 string = v5
                            v45 = "\n\t"
                            for fi46 := 0; fi46 < 19; fi46++ {
                                v47 := GenS5{F0: 4294967295, F1: (v38 & v4)}
                                if false {
                                    _ = len(v2)
                                    var v48 uint8 = 7
                                    _ = v48
                                } else {
                                    _ = len(v2)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v49 uint = (v1 - (^(v1)))
                                    _ = v49
                                }
                                var v50 int16 = 42
                                _ = fi46
                                _ = v47
                                _ = v50
                            }
                            _ = v45
                        }
                    }
                }
            }
            _ = fi39
        }
        v5 = "foo bar"
        _ = v38
    } else {
        fmt.Println(v5, v3)
        for fi51 := 0; fi51 < 13; fi51++ {
            for fi52 := 0; fi52 < 10; fi52++ {
                switch (fi52 % ((-1) | 1)) {
                case -1:
                    v4 += (v4 | v4)
                case -2147483648:
                    if false {
                        if (-9223372036854775808 < 171) {
                            var v53 int = -2147483648
                            v2 = append(v2, 65535)
                            _ = v53
                        } else {
                            v54 := GenS6{F0: 0, F1: 1, F2: 1}
                            v55 := GenS6{F0: 1, F1: 1000, F2: 1000}
                            fmt.Println(v4, fi52)
                            v4 += v4
                            if true {
                                switch 1 {
                                case 127:
                                    v1 = v1
                                case -652:
                                    if true {
                                        var v56 bool = true
                                        _ = v56
                                    }
                                default:
                                    v1 -= 0
                                }
                            } else {
                                fmt.Println(v3, v3, v3)
                                v57 := GenS7{F0: -9223372036854775808, F1: (^(v1))}
                                v4 = 986
                                _ = v57
                            }
                            _ = v54
                            _ = v55
                        }
                        var v58 int32 = 42
                        _ = v58
                    } else {
                        v5 = v5
                        if (true || (false && true)) {
                            switch (fi52 | fi52) {
                            case 2147483647:
                                fmt.Println(v1, fi51)
                            case 4294967295:
                                switch 127 {
                                case 499:
                                    var v59 int8 = int8(fi51)
                                    _ = v59
                                case 32767:
                                    v60 := GenS3{F0: int32(v4), F1: 32767}
                                    _ = v60
                                }
                            case 32768:
                                switch "hello" {
                                case "a":
                                    var v61 float64 = (3.14159 + 3.14159)
                                    _ = v61
                                case "":
                                    v5 = ("foo bar" + ("hello" + v5))
                                case "hello":
                                    fmt.Println(v4, v4)
                                }
                            }
                        } else {
                            if (true || true) {
                                if false {
                                    var v62 uint = v1
                                    _ = v62
                                }
                                if false {
                                    var v63 bool = (true || false)
                                    _ = v63
                                }
                                v5 = "a"
                            }
                        }
                        v1 = 1
                    }
                default:
                    _ = len(v2)
                }
                v3 = uint16(v1)
                v64 := GenS8{F0: 0}
                if true {
                    v3 -= (v3 & uint16(v4))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v2 = append(v2, (^(v4)))
                } else {
                    fmt.Println(v4)
                    v2 = append(v2, uint32(v1))
                }
                v2 = append(v2, (v4 + v4))
                _ = fi52
                _ = v64
            }
            v2 = append(v2, (v4 % ((v4) | 1)))
            v3 = 32767
            _ = fi51
        }
    }
    var v65 int8 = -1
    switch v5 {
    case "foo bar":
        v5 = v5
    }
    var v66 int32 = -26
    for fi67 := 0; fi67 < 8; fi67++ {
        v68 := GenS9{F0: 85, F1: (("\n\t" + "a") + v5)}
        _ = fi67
        _ = v68
    }
    var v69 int = 4294967295
    var v70 uint64 = 461
    var v71 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v69
    _ = v70
    _ = v71
}
