package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int64; F2 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool }

func main() {
    v1 := []uint32{}
    v2 := []uint32{2, 2147483647}
    v3 := make(map[int]uint32)
    var v4 int32 = -2147483648
    var v5 uint16 = 501
    v6 := GenS1{F0: true}
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            fmt.Println(fi7, fi7, v5)
            var v9 int = (fi7 % (((fi8 << ((65535) & 63))) | 1))
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 int16 = -32768
            if (34 < 127) {
                v5 = v5
                v5 += 334
                if true {
                    v5 = 0
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                    fmt.Println(v5)
                } else {
                    fmt.Println(v5)
                    v12 := GenS2{F0: 151}
                    _ = v12
                }
            } else {
                var v13 uint64 = 0
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        _ = fi7
    }
    switch "foo bar" {
    case "foo bar":
        var v14 int = 2
        _ = v14
    default:
        if true {
            v5 += 114
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 uint8 = 136
                if false {
                    v16 += (v16 << ((0) & 63))
                } else {
                    if true {
                        fmt.Println(v5)
                    }
                    v17 := GenS1{F0: (!true)}
                    delete(v3, int(v5))
                    _ = v17
                }
                v5 = (v5 ^ 724)
                _ = fi15
                _ = v16
            }
        }
    }
    if false {
        v4 += -519
    } else {
        var v18 float32 = 1.5
        switch "foo bar" {
        case "foo bar":
            if true {
                fmt.Println(v5)
                fmt.Println(v18, v5)
            } else {
                for fi19 := 0; fi19 < 11; fi19++ {
                    if true {
                        fmt.Println(v5, v18, v5)
                        if false {
                            for fi20 := 0; fi20 < 18; fi20++ {
                                v21 := GenS3{F0: v5, F1: -9223372036854775808, F2: 2147483647}
                                if (3.14159 <= 1.5) {
                                    var v22 float32 = v18
                                    _ = v22
                                } else {
                                    var v23 float32 = (v18 + (v18 + (v18 + v18)))
                                    v24 := GenS2{F0: 127}
                                    fmt.Println(v4)
                                    _ = v23
                                    _ = v24
                                }
                                if ((((v5 - (v5 / ((v5) | 1))) > 219) || (363 == 18446744073709551615)) || (("hello" + (("\n\t" + "hello") + "foo bar")) >= ("hello" + "hello"))) {
                                    v25 := GenS1{F0: (0 < 1000)}
                                    switch "hello" {
                                    case "foo bar":
                                        v26 := GenS4{F0: 32767}
                                        _ = v26
                                    }
                                    _ = v25
                                }
                                _ = fi20
                                _ = v21
                            }
                        }
                        for fi27 := 0; fi27 < 10; fi27++ {
                            v28 := GenS2{F0: 225}
                            if (!(-42 <= -566)) {
                                fmt.Println(fi19, v5, v5)
                                v29 := GenS5{F0: (-(fi19)), F1: v4, F2: 0}
                                var v30 uint = 587
                                v5 -= (v5 + 255)
                                v18 = v18
                                v4 -= -1
                                _ = v29
                                _ = v30
                            } else {
                                v3[(fi27 * 747)] = 2147483647
                                v18 = v18
                            }
                            _ = len(v3)
                            _ = fi27
                            _ = v28
                        }
                    } else {
                        if (-1 > 42) {
                            fmt.Println(v5, v5, v4)
                        } else {
                            fmt.Println(v5)
                        }
                        var v31 int = fi19
                        var v32 int8 = 0
                        _ = v31
                        _ = v32
                    }
                    _ = fi19
                }
                v33 := GenS1{F0: (false || (41 <= 32768))}
                _ = v33
            }
        case "a":
            var v34 int64 = 0
            _ = v34
        default:
            v35 := GenS6{F0: 99}
            _ = v35
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v18
    }
    v36 := GenS3{F0: v5, F1: 0, F2: (v4 | (v4 & v4))}
    if (true && true) {
        switch ("" + (("foo bar" + "\n\t") + ("\n\t" + "\n\t"))) {
        case "hello":
            var v37 uint16 = v5
            _ = v37
        case "\n\t":
            if (42 != 128) {
                switch 128 {
                case 42:
                    fmt.Println(v4)
                }
                v5 = v5
            } else {
                switch (((("hello" + "foo bar") + ("a" + "foo bar")) + "") + "a") {
                case "\n\t":
                    if true {
                        var v38 uint16 = (v5 << ((553) & 63))
                        var v39 uint8 = 125
                        _ = v38
                        _ = v39
                    } else {
                        var v40 int32 = (-(v4))
                        if len(v1) > 0 {
                            v1[0] = 2
                        }
                        _ = len(v2)
                        _ = v40
                    }
                default:
                    v41 := GenS7{F0: ((false && false) && true)}
                    _ = v41
                }
            }
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        for fi42 := 0; fi42 < 11; fi42++ {
            var v43 uint = 817
            for fi44 := 0; fi44 < 3; fi44++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v45 uint = 42
                _ = fi44
                _ = v45
            }
            _ = fi42
            _ = v43
        }
        if len(v2) > 0 {
            v2[0] = 42
        }
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        if true {
            fmt.Println(v5, fi46)
            var v47 uint64 = 42
            switch "\n\t" {
            case "":
                if false {
                    var v48 int = -2147483648
                    _ = v48
                }
            case "\n\t":
                v49 := GenS8{F0: -664}
                _ = v49
            default:
                v50 := GenS9{F0: true}
                _ = v50
            }
            _ = v47
        }
        _ = fi46
    }
    var v51 int16 = 42
    v5 -= v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
