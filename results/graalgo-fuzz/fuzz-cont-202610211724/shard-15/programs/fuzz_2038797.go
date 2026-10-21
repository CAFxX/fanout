package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint32; F1 bool; F2 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 string }
func genFn0() uint32 {
    v7 := GenS1{F0: 317, F1: 1}
    var v8 bool = (634 > 0)
    var v9 uint32 = 2147483648
    var v10 string = ("" + "foo bar")
    var v11 uint64 = 42
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return v9
}
func genFn1(p0 uint64) int {
    var v12 int8 = 42
    v13 := GenS1{F0: 2, F1: 128}
    v12 -= v12
    v14 := []int{0, -1}
    if true {
        fmt.Println(v12, v12)
        v12 -= v12
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 9223372036854775807
}

func main() {
    v1 := []int{4294967295, -1, -982}
    var v2 int32 = -223
    var v3 bool = (299 == 183)
    v4 := []string{"hello", ""}
    v5 := []uint32{42}
    var v6 int64 = -558
    if (v3 || false) {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 uint8 = 1
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            for fi18 := 0; fi18 < 1; fi18++ {
                fmt.Println(v6)
                _ = len(v5)
                _ = fi18
            }
            _ = fi17
        }
        v2 = -433
    } else {
        if false {
            fmt.Println(v3)
            v19 := GenS1{F0: 11, F1: 128}
            var v20 uint32 = 0
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 8; fi21++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v22 := GenS1{F0: 65535, F1: 127}
                v23 := GenS1{F0: 1, F1: 253}
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 int16 = -1
            _ = v24
        }
        fmt.Println(v3, v2)
        if false {
            for fi25 := 0; fi25 < 18; fi25++ {
                for fi26 := 0; fi26 < 7; fi26++ {
                    v2 = (v2 >> ((18446744073709551615) & 63))
                    _ = fi26
                }
                for fi27 := 0; fi27 < 13; fi27++ {
                    v28 := GenS2{F0: (v6 << ((274) & 63)), F1: 1, F2: 1}
                    var v29 int64 = v6
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = fi25
            }
            var v30 int = 9223372036854775807
            switch "" {
            case "foo bar":
                var v31 uint64 = 18446744073709551615
                _ = v31
            default:
                var v32 uint32 = 1
                _ = v32
            }
            _ = v30
        }
    }
    var v33 int8 = 70
    for fi34 := 0; fi34 < 5; fi34++ {
        v3 = (true && v3)
        switch -1 {
        case 32768:
            var v35 string = "foo bar"
            _ = v35
        case 42:
            fmt.Println(v33)
        }
        if ("" != "foo bar") {
            if true {
                fmt.Println(fi34, fi34)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v36 uint16 = 0
                _ = v36
            } else {
                if (2 <= 497) {
                    var v37 float64 = float64(0.0)
                    if v3 {
                        for fi38 := 0; fi38 < 10; fi38++ {
                            switch 127 {
                            case 0:
                                if true {
                                    v33 = v33
                                }
                            case 829:
                                var v39 uint16 = 1
                                _ = v39
                            case -1:
                                if ((-1 == (v33 << ((42) & 63))) || (v2 > (v2 >> ((18446744073709551615) & 63)))) {
                                    v2 = v2
                                    var v40 bool = (v33 >= (v33 & (v33 | (-(v33)))))
                                    if false {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v41 := GenS3{F0: 4294967295, F1: 100}
                                        v6 -= (v6 << ((2) & 63))
                                        _ = v41
                                    }
                                    _ = v40
                                } else {
                                    v2 -= (-(v2))
                                }
                            default:
                                v42 := GenS1{F0: 255, F1: 183}
                                _ = v42
                            }
                            v37 = v37
                            _ = fi38
                        }
                        if len(v5) > 0 {
                            v5[0] = 867
                        }
                    } else {
                        for fi43 := 0; fi43 < 20; fi43++ {
                            if v3 {
                                v44 := GenS4{F0: v33}
                                var v45 bool = true
                                if len(v4) > 0 {
                                    v4[0] = "foo bar"
                                }
                                if len(v1) > 0 {
                                    v1[0] = (fi43 ^ genFn1(15))
                                }
                                _ = v44
                                _ = v45
                            } else {
                                v6 = v6
                            }
                            var v46 uint = 223
                            _ = fi43
                            _ = v46
                        }
                        v47 := GenS5{F0: 1, F1: "a"}
                        switch "\n\t" {
                        case "foo bar":
                            v6 = v6
                        case "":
                            var v48 int64 = (v6 - v6)
                            _ = v48
                        case "\n\t":
                            v49 := GenS6{F0: 475, F1: v3, F2: v6}
                            _ = v49
                        default:
                            var v50 bool = true
                            _ = v50
                        }
                        _ = v47
                    }
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    _ = v37
                } else {
                    var v51 int8 = v33
                    _ = v51
                }
                for fi52 := 0; fi52 < 8; fi52++ {
                    v2 = v2
                    v5 = append(v5, 2147483648)
                    fmt.Println(v2, v2)
                    _ = fi52
                }
                switch "" {
                case "":
                    if (487 >= 730) {
                        var v53 int8 = (v33 << ((0) & 63))
                        switch "" {
                        case "a":
                            switch (fi34 - (fi34 ^ int(v6))) {
                            case -9223372036854775808:
                                for fi54 := 0; fi54 < 1; fi54++ {
                                    fmt.Println(fi34, v53)
                                    v53 = v53
                                    _ = fi54
                                }
                            case 32768:
                                v1 = append(v1, (fi34 * (fi34 << ((333) & 63))))
                            case 0:
                                fmt.Println(v3, v33, v33)
                            default:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        case "":
                            var v55 int8 = 0
                            _ = v55
                        case "\n\t":
                            _ = len(v1)
                        }
                        _ = v53
                    }
                case "\n\t":
                    for fi56 := 0; fi56 < 7; fi56++ {
                        v6 = v6
                        _ = fi56
                    }
                case "foo bar":
                    _ = len(v5)
                default:
                    v3 = (1e10 == (3.14159 + ((float32(float64((0.0 - float32((0.0 - 1.5))))) + ((((float32(0.1) * 1.5) + float32(0.0)) / float32(0.1)) * (1.5 - (((1.5 + 1e10) * -2.25) - -2.25)))) / 1.0)))
                }
            }
            if v3 {
                v57 := GenS1{F0: 65535, F1: 128}
                for fi58 := 0; fi58 < 1; fi58++ {
                    v59 := GenS6{F0: 207, F1: (false && false), F2: v6}
                    v33 = (v33 ^ (v33 - v33))
                    _ = fi58
                    _ = v59
                }
                if len(v5) > 0 {
                    v5[0] = genFn0()
                }
                _ = v57
            } else {
                v2 += v2
                _ = len(v1)
                var v60 uint16 = 115
                _ = v60
            }
        } else {
            var v61 int8 = v33
            _ = v61
        }
        _ = fi34
    }
    for fi62 := 0; fi62 < 9; fi62++ {
        var v63 uint32 = genFn0()
        v1 = append(v1, fi62)
        _ = fi62
        _ = v63
    }
    v64 := GenS7{F0: "\n\t"}
    var v65 uint64 = 391
    for fi66 := 0; fi66 < 6; fi66++ {
        switch "\n\t" {
        case "a":
            switch 931 {
            case 2147483647:
                var v67 float32 = (float32(1.5) * (((3.14159 * ((1e10 + 1e10) + 1.5)) + 1.5) + 3.14159))
                _ = v67
            default:
                v3 = (!false)
            }
        case "hello":
            v68 := GenS8{F0: 65535, F1: (v65 >> ((1) & 63))}
            _ = v68
        case "\n\t":
            var v69 int64 = (v6 >> ((627) & 63))
            _ = v69
        }
        fmt.Println(v6, v33)
        _ = fi66
    }
    v65 = v65
    var v70 int16 = -237
    fmt.Println(v2)
    switch "\n\t" {
    case "a":
        v71 := GenS9{F0: ("hello" + "")}
        _ = v71
    case "foo bar":
        for fi72 := 0; fi72 < 6; fi72++ {
            fmt.Println(v65, v33)
            _ = fi72
        }
    }
    fmt.Println(v33)
    v6 += v6
    var v73 float64 = 1e10
    var v74 int8 = v33
    for fi75 := 0; fi75 < 14; fi75++ {
        fmt.Println(v2, v65, v3)
        switch "" {
        case "hello":
            v73 = (0.0 / 3.14159)
        case "":
            _ = len(v4)
        }
        _ = fi75
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v70:", v70)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
    _ = v70
    _ = v73
    _ = v74
}
