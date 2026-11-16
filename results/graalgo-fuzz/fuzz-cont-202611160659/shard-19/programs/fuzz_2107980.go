package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 int8; F1 uint16; F2 uint }

func main() {
    var v1 float64 = 1e10
    var v2 uint32 = 42
    var v3 float64 = v1
    var v4 uint64 = 18446744073709551615
    v5 := []int{176, 9223372036854775807, 4294967295}
    v5 = append(v5, -1)
    v1 -= 0.1
    v2 = 300
    var v6 uint8 = 1
    fmt.Println(v1, v1, v1)
    v1 = float64((3.14159 + float32((-2.25 * v3))))
    if false {
        v1 += ((0.0 + (float64(0.1) - float64(0.1))) + 0.0)
        if false {
            switch "foo bar" {
            case "\n\t":
                for fi7 := 0; fi7 < 15; fi7++ {
                    v1 -= v1
                    _ = fi7
                }
            }
            for fi8 := 0; fi8 < 19; fi8++ {
                if (0 > fi8) {
                    var v9 string = "a"
                    _ = v9
                }
                if (!(false || false)) {
                    switch fi8 {
                    case 32767:
                        switch (fi8 & fi8) {
                        case -6:
                            for fi10 := 0; fi10 < 2; fi10++ {
                                var v11 string = "\n\t"
                                var v12 uint = uint(v2)
                                if (!(939 > 42)) {
                                    switch 0 {
                                    case 42:
                                        if true {
                                            var v13 float64 = (v3 + 0.1)
                                            _ = len(v5)
                                            _ = v13
                                        } else {
                                            fmt.Println(v4, v6)
                                            switch v11 {
                                            case "\n\t":
                                                var v14 float64 = 1.5
                                                _ = v14
                                            case "hello":
                                                fmt.Println(v3, v12, v11)
                                            default:
                                                var v15 int32 = -2147483648
                                                _ = v15
                                            }
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                    case -250:
                                        v4 = (v4 - (v4 >> ((v12) & 63)))
                                    case -2147483648:
                                        if (true && true) {
                                            v16 := GenS1{F0: true, F1: v11}
                                            var v17 string = v11
                                            _ = v16
                                            _ = v17
                                        }
                                    default:
                                        v18 := GenS2{F0: 2147483647, F1: 9223372036854775807}
                                        _ = v18
                                    }
                                    var v19 float32 = (float32(v3) - 1.5)
                                    _ = v19
                                } else {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v20 uint16 = uint16(v12)
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v20
                                }
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                        case 0:
                            for fi21 := 0; fi21 < 6; fi21++ {
                                var v22 uint32 = v2
                                if true {
                                    if len(v5) > 0 {
                                        v5[0] = (fi8 >> ((2) & 63))
                                    }
                                    v4 = v4
                                }
                                _ = fi21
                                _ = v22
                            }
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v3 += v3
                }
                if true {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        fmt.Println(v3)
                        v3 = v1
                        _ = fi23
                    }
                    switch (fi8 | fi8) {
                    case -1:
                        fmt.Println(v2, v4)
                    case 2147483647:
                        if true {
                            if (!true) {
                                fmt.Println(v2, v3)
                                v3 = ((((3.14159 + v3) + 0.0) + 3.14159) * 3.14159)
                                v4 = (v4 & (v4 * v4))
                            }
                            var v24 uint = 42
                            _ = v24
                        }
                    default:
                        fmt.Println(v4, v4, v6)
                    }
                } else {
                    v3 += v1
                }
                _ = fi8
            }
        }
    } else {
        switch 32768 {
        case 2147483647:
            fmt.Println(v2, v4, v1)
        case -1:
            var v25 uint8 = (v6 << ((0) & 63))
            _ = v25
        default:
            fmt.Println(v3)
        }
    }
    var v26 uint64 = (^(v4))
    for fi27 := 0; fi27 < 16; fi27++ {
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 int = fi28
            _ = fi28
            _ = v29
        }
        _ = len(v5)
        _ = fi27
    }
    if false {
        v30 := GenS1{F0: (-102 > -32768), F1: "\n\t"}
        if len(v5) > 0 {
            v5[0] = -565
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            switch "hello" {
            case "foo bar":
                fmt.Println(fi31)
            case "a":
                var v32 uint16 = 65535
                _ = v32
            case "\n\t":
                var v33 uint16 = 461
                _ = v33
            default:
                if (1 != 968) {
                    v1 = 3.14159
                    v4 = (v4 + (v26 / (((v26 << ((42) & 63))) | 1)))
                    if true {
                        if false {
                            if (false || true) {
                                fmt.Println(v1, v6, v3)
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    if len(v5) > 0 {
                                        v5[0] = fi34
                                    }
                                    v35 := GenS1{F0: true, F1: ("hello" + "a")}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v36 float64 = 1.5
                                    v37 := GenS3{F0: (!true), F1: (^(v26))}
                                    v38 := GenS3{F0: true, F1: (v4 % ((2) | 1))}
                                    var v39 int8 = 96
                                    if true {
                                        var v40 float32 = 3.14159
                                        _ = v40
                                    } else {
                                        _ = len(v5)
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    var v41 bool = (v6 == (v6 % (((v6 << ((1) & 63))) | 1)))
                                    var v42 int16 = 271
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                    _ = v41
                                    _ = v42
                                }
                            } else {
                                if false {
                                    v43 := GenS4{F0: 255}
                                    _ = v43
                                } else {
                                    v2 += v2
                                }
                                v3 -= 0.0
                                v44 := GenS3{F0: ((32767 >= 42) || true), F1: (^(v4))}
                                _ = v44
                            }
                            switch (fi31 / (((fi31 - (fi31 ^ -22))) | 1)) {
                            case 296:
                                v1 += 0.0
                            case 346:
                                if (222 == 984) {
                                    var v45 uint16 = 1
                                    fmt.Println(v26, v26, v6)
                                    v46 := GenS4{F0: v6}
                                    _ = v45
                                    _ = v46
                                } else {
                                    for fi47 := 0; fi47 < 7; fi47++ {
                                        _ = len(v5)
                                        v48 := GenS5{F0: -1, F1: 42}
                                        var v49 uint8 = v6
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                    }
                                }
                            case -192:
                                v50 := GenS6{F0: 165, F1: true}
                                _ = v50
                            }
                        } else {
                            v51 := GenS7{F0: -57, F1: 65535, F2: 1}
                            _ = v51
                        }
                        var v52 uint = 255
                        _ = v52
                    } else {
                        for fi53 := 0; fi53 < 17; fi53++ {
                            v54 := GenS2{F0: -400, F1: 1}
                            _ = fi53
                            _ = v54
                        }
                        if true {
                            for fi55 := 0; fi55 < 9; fi55++ {
                                var v56 int32 = -2147483648
                                _ = fi55
                                _ = v56
                            }
                        }
                        if (4294967295 > 4294967295) {
                            v6 -= (v6 + 247)
                        } else {
                            var v57 int = (-(fi31))
                            var v58 uint8 = (v6 - 128)
                            _ = v57
                            _ = v58
                        }
                    }
                } else {
                    var v59 float32 = 0.1
                    v4 = v4
                    _ = v59
                }
            }
            _ = len(v5)
            switch "" {
            case "hello":
                v5 = append(v5, fi31)
            case "\n\t":
                var v60 bool = false
                _ = v60
            default:
                v5 = append(v5, (-(fi31)))
            }
            _ = fi31
        }
        _ = v30
    }
    var v61 uint16 = 65535
    v1 += v3
    for fi62 := 0; fi62 < 16; fi62++ {
        for fi63 := 0; fi63 < 20; fi63++ {
            v1 += v1
            _ = fi63
        }
        _ = fi62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
