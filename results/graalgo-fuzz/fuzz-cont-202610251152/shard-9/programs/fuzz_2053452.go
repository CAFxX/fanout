package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 bool; F1 bool; F2 int }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint64; F1 int8; F2 bool }
type GenS5 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32; F1 int64; F2 uint8 }
func genFn0(p0 string, p1 int) uint16 {
    var v7 uint32 = 534
    fmt.Println(v7)
    p0 = p0
    v8 := GenS1{F0: (^(v7)), F1: -1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 88
}
func genFn1(p0 int, p1 int32) int32 {
    fmt.Println(p1, p0)
    var v9 uint8 = 1
    if false {
        var v10 uint = 991
        fmt.Println(v9, v9, v10)
        _ = v10
    } else {
        for fi11 := 0; fi11 < 19; fi11++ {
            if ((0.0 == 0.0) || (2147483647 == 0)) {
                v12 := GenS2{F0: (false || (((!false) && (2 == 830)) || (!(!false)))), F1: (732 != -864), F2: 42}
                v13 := GenS2{F0: true, F1: (42 > 255), F2: (-(p0))}
                for fi14 := 0; fi14 < 3; fi14++ {
                    p1 = p1
                    _ = fi14
                }
                _ = v12
                _ = v13
            } else {
                if true {
                    fmt.Println(p1)
                } else {
                    fmt.Println(p1)
                    switch "foo bar" {
                    case "":
                        v15 := GenS1{F0: 4294967295, F1: -19}
                        _ = v15
                    case "hello":
                        p1 = (p1 % ((p1) | 1))
                    case "foo bar":
                        v16 := []int{fi11, 32767, fi11}
                        _ = v16
                    }
                }
            }
            _ = fi11
        }
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v9
    return p1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 string = "hello"
    var v4 int64 = 956
    v5 := make(map[int]bool)
    var v6 float64 = ((0.1 - ((3.14159 * 0.0) * (-2.25 - -2.25))) * -2.25)
    v2[2147483647] = false
    v3 = (v3 + ((v3 + v3) + (v3 + v3)))
    v17 := []int{9223372036854775807, 4294967295}
    v18 := GenS1{F0: 2147483647, F1: (v4 << ((485) & 63))}
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            fmt.Println(v3, fi20, v6)
            var v21 bool = (!true)
            _ = fi20
            _ = v21
        }
        var v22 string = "foo bar"
        _ = fi19
        _ = v22
    }
    v23 := GenS1{F0: 0, F1: v4}
    if (false || (!(!true))) {
        var v24 int32 = 666
        fmt.Println(v24)
        _ = v24
    } else {
        v25 := GenS2{F0: false, F1: false, F2: 32768}
        var v26 uint64 = 537
        _ = v25
        _ = v26
    }
    fmt.Println(v3, v3, v4)
    _ = len(v5)
    if true {
        for fi27 := 0; fi27 < 20; fi27++ {
            switch 1 {
            case 9223372036854775807:
                for fi28 := 0; fi28 < 16; fi28++ {
                    if (0.1 != -2.25) {
                        var v29 int32 = 0
                        if (255 == 1) {
                            if len(v17) > 0 {
                                v17[0] = fi27
                            }
                            v3 = v3
                        } else {
                            v6 -= (v6 - (((v6 * v6) / (3.14159 * v6)) - -2.25))
                            v29 = genFn1(fi28, v29)
                        }
                        if (false && false) {
                            switch (fi27 << ((255) & 63)) {
                            case 9223372036854775807:
                                if (int8(v29) != 4) {
                                    v30 := GenS3{F0: (v6 > (0.0 - -2.25)), F1: 42}
                                    delete(v1, 194)
                                    _ = v30
                                } else {
                                    if true {
                                        v17 = append(v17, 558)
                                        v4 += int64(v29)
                                        var v31 uint16 = 167
                                        _ = v31
                                    } else {
                                        var v32 float64 = ((0.1 - 1e10) + v6)
                                        v33 := GenS1{F0: 65535, F1: (v4 & (v4 & (v4 >> ((254) & 63))))}
                                        v6 = 0.0
                                        switch v3 {
                                        case "\n\t":
                                            v3 = (v3 + v3)
                                        case "":
                                            var v34 float64 = v6
                                            switch "" {
                                            case "hello":
                                                var v35 uint8 = 2
                                                _ = v35
                                            }
                                            if (32767 == 255) {
                                                v36 := GenS4{F0: 2, F1: 127, F2: (993 > int16(v4))}
                                                var v37 bool = true
                                                if (!false) {
                                                    v38 := GenS5{F0: 127, F1: 65535, F2: v4}
                                                    _ = v38
                                                }
                                                _ = v36
                                                _ = v37
                                            } else {
                                                switch ((v3 + v3) + v3) {
                                                case "foo bar":
                                                    var v39 uint = 42
                                                    _ = v39
                                                default:
                                                    v34 = -2.25
                                                }
                                            }
                                            _ = v34
                                        default:
                                            v3 = v3
                                        }
                                        _ = v32
                                        _ = v33
                                    }
                                    if (!false) {
                                        if (641 <= 938) {
                                            v3 = v3
                                        }
                                    }
                                    v5[42] = true
                                }
                            }
                        } else {
                            v40 := GenS1{F0: 0, F1: v4}
                            _ = v40
                        }
                        var v41 int64 = v4
                        _ = v29
                        _ = v41
                    } else {
                        if ((1 == 42) && true) {
                            v42 := GenS5{F0: 127, F1: 2, F2: (v4 << ((255) & 63))}
                            v1[65535] = 42
                            fmt.Println(fi27, fi28, fi28)
                            _ = v42
                        }
                        var v43 float64 = -2.25
                        if (-128 == 127) {
                            v3 = ((v3 + (v3 + v3)) + "\n\t")
                        }
                        _ = v43
                    }
                    v44 := GenS5{F0: 100, F1: 1, F2: v4}
                    v45 := GenS4{F0: 42, F1: 46, F2: (("" + "a") <= v3)}
                    _ = fi28
                    _ = v44
                    _ = v45
                }
            }
            _ = fi27
        }
        v46 := GenS3{F0: (false && (false && true)), F1: 42}
        _ = v46
    } else {
        for fi47 := 0; fi47 < 19; fi47++ {
            v6 = 3.14159
            for fi48 := 0; fi48 < 10; fi48++ {
                v6 = (v6 * (float64(float32(((v6 - 0.0) * v6))) - v6))
                v6 -= (0.1 - (0.0 - (0.1 * (v6 * v6))))
                _ = fi48
            }
            _ = fi47
        }
        fmt.Println(v6, v4, v3)
    }
    var v49 uint32 = 752
    if (65535 != 293) {
        v50 := GenS5{F0: 244, F1: 1, F2: v4}
        _ = v50
    }
    for fi51 := 0; fi51 < 3; fi51++ {
        if false {
            var v52 int32 = 0
            var v53 bool = false
            _ = v52
            _ = v53
        }
        var v54 int16 = -32768
        switch v3 {
        case "foo bar":
            if (42 != 65535) {
                v55 := GenS6{F0: 0}
                v56 := GenS7{F0: (v54 & (v54 - (v54 - v54)))}
                _ = v55
                _ = v56
            }
        case "hello":
            if ((v54 - v54) != (v54 % ((v54) | 1))) {
                switch fi51 {
                case 0:
                    var v57 uint64 = 2
                    _ = v57
                case 252:
                    delete(v5, (fi51 % ((127) | 1)))
                case -755:
                    for fi58 := 0; fi58 < 11; fi58++ {
                        fmt.Println(v49, v49)
                        _ = fi58
                    }
                default:
                    var v59 int16 = v54
                    _ = v59
                }
            }
        case "\n\t":
            if true {
                v54 = v54
                var v60 uint32 = (^(v49))
                v6 = v6
                _ = v60
            }
        }
        _ = fi51
        _ = v54
    }
    fmt.Println(v3)
    v6 += (((1.5 - v6) * float64((float32(v6) + (((float32(v6) - 3.14159) - (1e10 * 1e10)) - (3.14159 - (0.0 / 1.5)))))) / v6)
    if (1 >= 1000) {
        v61 := GenS5{F0: 1, F1: 38, F2: -1}
        var v62 int32 = -418
        if false {
            if (!((^(v49)) == v49)) {
                v63 := GenS8{F0: genFn1(9223372036854775807, int32(v49)), F1: (v4 * (^(v4))), F2: 127}
                _ = v63
            } else {
                for fi64 := 0; fi64 < 5; fi64++ {
                    v62 = 0
                    _ = fi64
                }
            }
            fmt.Println(v62, v3, v62)
            switch (v3 + "hello") {
            case "hello":
                v6 = float64((-2.25 - float32(v6)))
            case "\n\t":
                v3 = v3
            default:
                if false {
                    for fi65 := 0; fi65 < 11; fi65++ {
                        var v66 string = v3
                        v3 = ((v66 + ((v3 + v66) + "hello")) + (v66 + v66))
                        _ = fi65
                        _ = v66
                    }
                    var v67 uint64 = 1
                    v3 = (v3 + (v3 + v3))
                    _ = v67
                } else {
                    v6 = (((v6 + v6) / v6) / v6)
                    v4 -= v4
                }
            }
        } else {
            v68 := GenS7{F0: -677}
            _ = v68
        }
        _ = v61
        _ = v62
    }
    var v69 bool = true
    delete(v2, -139)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v69
}
