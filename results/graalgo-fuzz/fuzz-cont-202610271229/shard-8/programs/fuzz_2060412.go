package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 uint16; F1 int16; F2 int }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int32; F1 int; F2 string }
type GenS5 struct { F0 string }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{1, 65535, 65535}
    var v3 float64 = 0.1
    v4 := []uint32{794, 2147483648, 659}
    v5 := make(map[int]int)
    var v6 uint64 = 2
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: 892, F1: fi8, F2: 2}
            v6 = (^(v6))
            var v10 int32 = -1
            _ = fi8
            _ = v10
            _ = v9
        }
        v6 = v6
        _ = fi7
    }
    _ = len(v4)
    v6 = v6
    var v11 uint64 = 1
    v12 := v5[-9223372036854775808]
    v12 = v12
    v6 += 18446744073709551615
    var v13 string = ("\n\t" + ("hello" + "a"))
    switch v12 {
    case -13:
        switch v13 {
        case "hello":
            var v14 uint8 = 227
            _ = v14
        default:
            for fi15 := 0; fi15 < 12; fi15++ {
                v13 = (((v13 + v13) + v13) + (v13 + (("hello" + (("foo bar" + v13) + v13)) + "\n\t")))
                _ = fi15
            }
        }
    case 32768:
        fmt.Println(v11, v13)
    case 2147483647:
        fmt.Println(v12, v6, v3)
    default:
        var v16 int32 = 661
        _ = v16
    }
    if (-1 <= 100) {
        for fi17 := 0; fi17 < 11; fi17++ {
            v2 = append(v2, 2)
            _ = fi17
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            switch v13 {
            case "":
                switch v13 {
                case "":
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                default:
                    fmt.Println(v3, fi18)
                }
            case "\n\t":
                fmt.Println(v6)
            default:
                var v19 int16 = -1
                _ = v19
            }
            switch v13 {
            case "foo bar":
                for fi20 := 0; fi20 < 3; fi20++ {
                    v3 = v3
                    var v21 float32 = (1.5 / 1.5)
                    v22 := GenS1{F0: -2147483648, F1: fi20, F2: 1}
                    var v23 int8 = 1
                    v24 := v5[(fi18 * fi20)]
                    v24 -= v12
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            case "\n\t":
                v13 = (v13 + "")
            case "hello":
                if true {
                    var v25 uint64 = 4294967295
                    _ = v25
                } else {
                    v12 = (-(v12))
                    var v26 int8 = 100
                    _ = v26
                }
            }
            _ = fi18
        }
        if (false && (18446744073709551615 != 244)) {
            _ = len(v1)
            v12 = 604
        } else {
            v6 = (v11 << ((636) & 63))
            if true {
                for fi27 := 0; fi27 < 6; fi27++ {
                    delete(v5, 1)
                    delete(v5, 1)
                    _ = fi27
                }
                for fi28 := 0; fi28 < 12; fi28++ {
                    delete(v1, 577)
                    if (!false) {
                        delete(v5, (v12 % (((v12 ^ fi28)) | 1)))
                    } else {
                        for fi29 := 0; fi29 < 18; fi29++ {
                            v30 := GenS2{F0: 2, F1: 883, F2: 506}
                            var v31 uint64 = 0
                            switch 0 {
                            case -823:
                                v3 = (3.14159 / v3)
                            case -2147483648:
                                v6 += (v11 & v11)
                            case 1:
                                fmt.Println(v12, v12, fi28)
                            }
                            var v32 uint16 = 154
                            v31 += (v31 << ((42) & 63))
                            _ = fi29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    }
                    var v33 uint32 = 2147483648
                    _ = fi28
                    _ = v33
                }
            } else {
                if true {
                    var v34 int16 = 32767
                    v3 -= (((0.1 + -2.25) + v3) + v3)
                    _ = len(v2)
                    _ = v34
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        var v35 int64 = 9223372036854775807
                        _ = v35
                    case "a":
                        _ = len(v1)
                    case "hello":
                        v13 = "a"
                    default:
                        var v36 uint8 = 121
                        _ = v36
                    }
                    v37 := GenS1{F0: 42, F1: (v12 | (-(v12))), F2: 0}
                    v4 = append(v4, 65535)
                    _ = v37
                }
                var v38 uint = 0
                _ = v38
            }
            v11 -= (^(v11))
        }
    } else {
        v39 := GenS3{F0: 648, F1: 67}
        _ = v39
    }
    for fi40 := 0; fi40 < 16; fi40++ {
        switch 32767 {
        case 32767:
            switch v13 {
            case "":
                v41 := GenS4{F0: 456, F1: v12, F2: "a"}
                _ = v41
            default:
                for fi42 := 0; fi42 < 20; fi42++ {
                    v43 := v1[v12]
                    v44 := GenS1{F0: 42, F1: (fi40 << ((0) & 63)), F2: 568}
                    if false {
                        v45 := GenS1{F0: -2147483648, F1: 32768, F2: 983}
                        _ = v45
                    }
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            }
        case -293:
            var v46 int16 = 1
            _ = v46
        default:
            v47 := GenS5{F0: v13}
            _ = v47
        }
        _ = fi40
    }
    fmt.Println(v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
