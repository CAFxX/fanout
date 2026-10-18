package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int32, p1 int, p2 uint) int {
    var v6 float32 = (0.1 * 0.0)
    v7 := []int{p1, p1}
    p0 = p0
    if len(v7) > 0 {
        v7[0] = p1
    }
    var v8 float32 = (v6 + 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{-686, -1, 32767, 42}
    var v3 uint64 = 18446744073709551615
    var v4 uint16 = 65535
    var v5 string = ("hello" + "hello")
    fmt.Println(v5, v3, v4)
    if true {
        v4 -= v4
        var v9 uint16 = 850
        fmt.Println(v4)
        _ = v9
    }
    v10 := GenS1{F0: 1}
    v4 = (^(v4))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch (v5 + "\n\t") {
    case "a":
        v5 = v5
    case "hello":
        var v11 int16 = -599
        _ = v11
    }
    if (-128 > 103) {
        v12 := GenS1{F0: 255}
        _ = v12
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v13 float64 = 3.14159
        _ = v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS1{F0: 0}
        fmt.Println(v3)
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 int32 = 322
            v4 = v4
            v19 := GenS1{F0: 74}
            if false {
                var v20 int16 = 1
                if true {
                    var v21 int32 = 42
                    v20 -= int16(v4)
                    switch int(v20) {
                    case -519:
                        delete(v1, 1)
                    case 32768:
                        if false {
                            switch (fi17 * -2147483648) {
                            case 128:
                                v5 = v5
                            case 2147483647:
                                v22 := GenS1{F0: 1}
                                _ = v22
                            case -582:
                                _ = len(v2)
                            default:
                                _ = len(v2)
                            }
                            var v23 int = fi17
                            var v24 bool = false
                            _ = v23
                            _ = v24
                        } else {
                            if len(v2) > 0 {
                                v2[0] = fi17
                            }
                            var v25 int8 = 127
                            v26 := GenS1{F0: 127}
                            _ = v25
                            _ = v26
                        }
                        switch fi17 {
                        case 745:
                            switch v5 {
                            case "a":
                                if false {
                                    delete(v1, 42)
                                    var v27 uint16 = 484
                                    v27 -= (v27 & v4)
                                    _ = v27
                                }
                            case "":
                                v2 = append(v2, fi16)
                            }
                        default:
                            v21 -= v21
                        }
                    case 128:
                        if false {
                            var v28 int8 = 42
                            _ = v28
                        } else {
                            if true {
                                var v29 int = 2147483647
                                v2 = append(v2, (fi16 - 127))
                                v30 := GenS1{F0: 253}
                                v3 = v3
                                _ = v29
                                _ = v30
                            } else {
                                v3 = (v3 & (v3 % (((v3 << ((0) & 63))) | 1)))
                                delete(v1, 2147483648)
                            }
                        }
                    default:
                        switch ((("" + (((v5 + (v5 + (v5 + ("a" + v5)))) + v5) + (v5 + ("a" + v5)))) + v5) + "\n\t") {
                        case "hello":
                            fmt.Println(fi16, v21, fi16)
                            var v31 int8 = -1
                            _ = v31
                        case "\n\t":
                            var v32 float64 = ((((1e10 + 1e10) * 1.5) / 3.14159) / 3.14159)
                            _ = v32
                        default:
                            fmt.Println(v3, fi17, v4)
                        }
                    }
                    _ = v21
                }
                v33 := GenS1{F0: 236}
                _ = v20
                _ = v33
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi34 := 0; fi34 < 6; fi34++ {
            v4 -= v4
            var v35 uint32 = 65535
            var v36 uint16 = v4
            _ = fi34
            _ = v35
            _ = v36
        }
        var v37 uint = 255
        _ = fi16
        _ = v37
    }
    switch 128 {
    case 0:
        v38 := GenS1{F0: 2}
        _ = v38
    }
    fmt.Println(v3, v4, v4)
    v39 := GenS2{F0: 191, F1: (!true), F2: 198}
    if ((32767 >= -18) && ((32767 != 1) && (v5 < (v5 + "hello")))) {
        if true {
            for fi40 := 0; fi40 < 18; fi40++ {
                var v41 uint16 = 166
                _ = fi40
                _ = v41
            }
        } else {
            if true {
                for fi42 := 0; fi42 < 3; fi42++ {
                    var v43 int64 = 9223372036854775807
                    v5 = (v5 + "a")
                    _ = fi42
                    _ = v43
                }
            } else {
                switch 4294967295 {
                case 42:
                    switch v5 {
                    case "foo bar":
                        v2 = append(v2, 32767)
                    case "":
                        v44 := GenS3{F0: -399, F1: 2}
                        _ = v44
                    }
                case -664:
                    switch 42 {
                    case 132:
                        switch 66 {
                        case 42:
                            if (0.0 <= 0.1) {
                                if (0 == -1) {
                                    v4 -= v4
                                    v5 = (v5 + v5)
                                    v45 := GenS4{F0: 830, F1: 42}
                                    _ = v45
                                } else {
                                    v46 := GenS5{F0: 32767, F1: 1000}
                                    for fi47 := 0; fi47 < 3; fi47++ {
                                        v2 = append(v2, fi47)
                                        for fi48 := 0; fi48 < 11; fi48++ {
                                            var v49 string = (v5 + "")
                                            fmt.Println(fi47, v5, fi48)
                                            switch fi48 {
                                            case 4294967295:
                                                v5 = (((v5 + "a") + "foo bar") + ("foo bar" + ((v49 + "\n\t") + (v49 + (v49 + (((v49 + "") + v5) + v5))))))
                                            default:
                                                v3 = uint64(v4)
                                            }
                                            _ = fi48
                                            _ = v49
                                        }
                                        var v50 float32 = 3.14159
                                        _ = fi47
                                        _ = v50
                                    }
                                    for fi51 := 0; fi51 < 2; fi51++ {
                                        if len(v2) > 0 {
                                            v2[0] = (fi51 >> ((0) & 63))
                                        }
                                        switch -961 {
                                        case 127:
                                            if false {
                                                if true {
                                                    v4 -= v4
                                                    v5 = "\n\t"
                                                    if true {
                                                        switch "hello" {
                                                        case "a":
                                                            var v52 uint8 = 132
                                                            _ = v52
                                                        case "hello":
                                                            if (0 < 42) {
                                                                var v53 int64 = -1
                                                                var v54 bool = (v3 >= 42)
                                                                var v55 int32 = 2147483647
                                                                _ = v53
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                        case "foo bar":
                                                            var v56 int64 = -9223372036854775808
                                                            _ = v56
                                                        default:
                                                            if (!true) {
                                                                var v57 int16 = 516
                                                                v5 = v5
                                                                v58 := GenS6{F0: v4}
                                                                _ = v57
                                                                _ = v58
                                                            }
                                                        }
                                                    }
                                                }
                                            } else {
                                                if ((v3 ^ (v3 % ((18446744073709551615) | 1))) <= v3) {
                                                    var v59 string = ("hello" + v5)
                                                    var v60 uint16 = (v4 << ((18446744073709551615) & 63))
                                                    _ = v59
                                                    _ = v60
                                                }
                                                v4 -= 1
                                                if false {
                                                    _ = len(v2)
                                                }
                                            }
                                        case 0:
                                            v4 = (^(v4))
                                        }
                                        var v61 int = (fi51 ^ (fi51 | (-(fi51))))
                                        _ = fi51
                                        _ = v61
                                    }
                                    _ = v46
                                }
                                fmt.Println(v3)
                                for fi62 := 0; fi62 < 20; fi62++ {
                                    var v63 int8 = 127
                                    _ = fi62
                                    _ = v63
                                }
                            } else {
                                if true {
                                    v3 = (v3 * (v3 % ((v3) | 1)))
                                } else {
                                    var v64 uint32 = 1
                                    v5 = "foo bar"
                                    _ = v64
                                }
                            }
                        default:
                            switch "hello" {
                            case "\n\t":
                                var v65 string = v5
                                _ = v65
                            }
                        }
                    }
                case 998:
                    switch 10 {
                    case 127:
                        if len(v2) > 0 {
                            v2[0] = -2147483648
                        }
                    }
                }
            }
        }
        for fi66 := 0; fi66 < 13; fi66++ {
            v4 = v4
            _ = fi66
        }
    }
    var v67 float64 = -2.25
    for fi68 := 0; fi68 < 15; fi68++ {
        v3 = 0
        _ = fi68
    }
    fmt.Println(v4)
    v4 = (v4 << ((2) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v67
}
