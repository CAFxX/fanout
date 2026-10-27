package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint64; F2 bool }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int64; F1 uint16; F2 int64 }
type GenS7 struct { F0 int16; F1 int8; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 212
    v3 := []string{"", "\n\t"}
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 int16 = int16(v2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi4
        _ = v5
    }
    fmt.Println(v2)
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            _ = len(v1)
            var v8 int64 = -1
            v9 := v1[136]
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v2 += v2
    var v10 int16 = 32767
    var v11 float64 = 3.14159
    v2 += v2
    v2 -= (v2 % ((2) | 1))
    if (!true) {
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 uint = uint(v2)
            var v14 bool = false
            if (18446744073709551615 != 18446744073709551615) {
                v15 := GenS1{F0: 128}
                _ = v15
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        delete(v1, 1)
        var v16 int16 = v10
        _ = v16
    } else {
        for fi17 := 0; fi17 < 10; fi17++ {
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            v18 := v1[65535]
            _ = fi17
            _ = v18
        }
    }
    if true {
        v19 := GenS2{F0: 548, F1: 4294967295, F2: false}
        _ = v19
    } else {
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 float32 = float32(v11)
            _ = fi20
            _ = v21
        }
        var v22 float32 = -2.25
        _ = v22
    }
    v23 := GenS2{F0: 1, F1: 539, F2: false}
    for fi24 := 0; fi24 < 16; fi24++ {
        v2 -= (v2 << ((536) & 63))
        v25 := GenS1{F0: (^(v2))}
        var v26 float32 = 0.0
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println(v10, v11)
    fmt.Println(v10, v11, v10)
    switch 260 {
    case 388:
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 float64 = 3.14159
            if true {
                v3 = append(v3, ("\n\t" + ""))
            } else {
                v2 += v2
                if false {
                    if (9223372036854775807 > int64(v2)) {
                        for fi29 := 0; fi29 < 3; fi29++ {
                            switch 9223372036854775807 {
                            case 0:
                                var v30 int32 = -427
                                _ = v30
                            default:
                                var v31 uint64 = 613
                                _ = v31
                            }
                            if (2 < 636) {
                                v32 := GenS3{F0: 2, F1: -828}
                                _ = v32
                            } else {
                                v28 += (v28 / v11)
                            }
                            var v33 int = fi27
                            _ = fi29
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 7; fi34++ {
                            fmt.Println(v2, fi27)
                            if true {
                                _ = len(v1)
                                delete(v1, 2147483647)
                            } else {
                                fmt.Println(v28, v10)
                                var v35 uint64 = 414
                                var v36 int32 = -42
                                _ = v35
                                _ = v36
                            }
                            _ = fi34
                        }
                        v2 += (^(v2))
                    }
                    for fi37 := 0; fi37 < 16; fi37++ {
                        var v38 uint8 = v2
                        fmt.Println(v38, fi37, v38)
                        var v39 float32 = -2.25
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                } else {
                    var v40 uint32 = 1
                    _ = v40
                }
            }
            v41 := GenS4{F0: 18446744073709551615, F1: (("foo bar" + ("" + ("" + "a"))) + "hello")}
            _ = fi27
            _ = v28
            _ = v41
        }
    default:
        v42 := GenS5{F0: 9223372036854775807, F1: (false && true)}
        _ = v42
    }
    switch 42 {
    case 9223372036854775807:
        switch 128 {
        case 4294967295:
            if (!false) {
                var v43 uint64 = 0
                var v44 float64 = (((v11 * (v11 * 0.1)) * v11) * v11)
                var v45 bool = false
                _ = v43
                _ = v44
                _ = v45
            } else {
                switch ("hello" + ("\n\t" + "")) {
                case "hello":
                    v10 = (v10 + -1)
                case "\n\t":
                    switch ("hello" + "hello") {
                    case "":
                        var v46 string = ("hello" + ("a" + "foo bar"))
                        _ = v46
                    case "\n\t":
                        fmt.Println(v2, v10)
                    case "a":
                        v11 -= v11
                    }
                default:
                    v11 -= 3.14159
                }
            }
        case 32767:
            switch "foo bar" {
            case "\n\t":
                var v47 int32 = int32(v10)
                _ = v47
            }
        }
    default:
        var v48 uint16 = 0
        _ = v48
    }
    if false {
        var v49 float64 = (3.14159 / v11)
        v2 = (v2 ^ v2)
        if true {
            switch 127 {
            case -9223372036854775808:
                v49 += 1e10
            case -2147483648:
                var v50 uint32 = 694
                _ = v50
            case 878:
                v49 = (v11 / 0.1)
            default:
                for fi51 := 0; fi51 < 1; fi51++ {
                    for fi52 := 0; fi52 < 8; fi52++ {
                        var v53 int8 = 100
                        if false {
                            v54 := GenS6{F0: int64(v2), F1: 1, F2: 1}
                            var v55 uint16 = 42
                            _ = v54
                            _ = v55
                        } else {
                            v56 := GenS7{F0: -891, F1: (v53 | (-(v53))), F2: v10}
                            var v57 int32 = 958
                            v11 = 0.0
                            switch (fi51 ^ fi51) {
                            case -2147483648:
                                v2 = v2
                            default:
                                switch fi52 {
                                case -1:
                                    v10 = (v10 >> ((516) & 63))
                                case 484:
                                    v49 = v49
                                case 411:
                                    var v58 int32 = v57
                                    _ = v58
                                }
                            }
                            delete(v1, 42)
                            _ = v56
                            _ = v57
                        }
                        _ = fi52
                        _ = v53
                    }
                    fmt.Println(fi51, fi51)
                    var v59 float64 = v11
                    _ = fi51
                    _ = v59
                }
            }
            v60 := GenS5{F0: 240, F1: true}
            _ = v60
        }
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
}
