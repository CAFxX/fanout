package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 string; F1 uint; F2 int }
type GenS5 struct { F0 bool; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8 }

func main() {
    v1 := []int{}
    var v2 float64 = 3.14159
    var v3 uint64 = 2
    var v4 int64 = 9223372036854775807
    var v5 int64 = v4
    var v6 uint32 = 0
    fmt.Println(v4, v5)
    v5 += (v4 | v4)
    if false {
        var v7 int8 = int8(v3)
        switch -13 {
        case 32767:
            if ((v4 - v5) > v5) {
                if (false && (((("\n\t" + "foo bar") + "hello") + "") >= "a")) {
                    v8 := GenS1{F0: 0, F1: 0}
                    var v9 int64 = 1
                    if (v3 != (v3 + 314)) {
                        v10 := GenS2{F0: int16(v9), F1: v3}
                        fmt.Println(v9)
                        switch "a" {
                        case "":
                            var v11 int16 = 42
                            _ = v11
                        case "a":
                            for fi12 := 0; fi12 < 4; fi12++ {
                                switch fi12 {
                                case 42:
                                    switch fi12 {
                                    case 128:
                                        switch fi12 {
                                        case 42:
                                            if (!true) {
                                                for fi13 := 0; fi13 < 15; fi13++ {
                                                    v5 = v5
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = fi13
                                                }
                                            } else {
                                                v2 = -2.25
                                                if (255 != 700) {
                                                    v14 := GenS1{F0: (-(v7)), F1: -22}
                                                    var v15 int64 = v4
                                                    var v16 uint32 = (v6 / ((v6) | 1))
                                                    _ = v14
                                                    _ = v15
                                                    _ = v16
                                                }
                                                v17 := GenS2{F0: 235, F1: v3}
                                                _ = v17
                                            }
                                        case 2147483647:
                                            v5 -= 0
                                        default:
                                            v9 = v9
                                        }
                                    case 32768:
                                        fmt.Println(v7, v7, fi12)
                                    case 282:
                                        if true {
                                            fmt.Println(v6, v6)
                                            fmt.Println(v4)
                                        }
                                    }
                                default:
                                    v18 := GenS2{F0: 1000, F1: (v3 & (v3 ^ v3))}
                                    _ = v18
                                }
                                v7 = (v7 << ((868) & 63))
                                _ = fi12
                            }
                        case "foo bar":
                            v1 = append(v1, 9223372036854775807)
                        default:
                            if (2147483647 >= 42) {
                                if false {
                                    fmt.Println(v7)
                                    for fi19 := 0; fi19 < 12; fi19++ {
                                        for fi20 := 0; fi20 < 17; fi20++ {
                                            switch "foo bar" {
                                            case "\n\t":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            default:
                                                _ = len(v1)
                                            }
                                            v21 := GenS1{F0: (v7 | v7), F1: 42}
                                            switch fi19 {
                                            case -1:
                                                switch "hello" {
                                                case "hello":
                                                    v3 = (v3 * 4294967295)
                                                case "\n\t":
                                                    if len(v1) > 0 {
                                                        v1[0] = fi20
                                                    }
                                                    v22 := GenS1{F0: v7, F1: -72}
                                                    _ = v22
                                                default:
                                                    if (!((4294967295 == 4294967295) && false)) {
                                                        v23 := GenS3{F0: 2, F1: v7}
                                                        if (-1 != -2147483648) {
                                                            v24 := GenS4{F0: ((("" + ("hello" + (("foo bar" + "a") + ((("" + "\n\t") + ("a" + "\n\t")) + ("hello" + (("" + "") + "\n\t")))))) + ("hello" + ("a" + "foo bar"))) + (("\n\t" + "") + ("foo bar" + "\n\t"))), F1: 2, F2: (fi19 & -9223372036854775808)}
                                                            _ = v24
                                                        } else {
                                                            v6 = (v6 * (v6 * 147))
                                                        }
                                                        _ = v23
                                                    } else {
                                                        switch fi20 {
                                                        case -1:
                                                            var v25 int16 = -1
                                                            _ = v25
                                                        default:
                                                            v6 = v6
                                                        }
                                                        v6 -= (^(v6))
                                                        if len(v1) > 0 {
                                                            v1[0] = -1
                                                        }
                                                    }
                                                    var v26 bool = (uint(v6) <= 0)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v26
                                                }
                                            case 42:
                                                fmt.Println(fi19, v6)
                                            }
                                            _ = fi20
                                            _ = v21
                                        }
                                        v4 += (v4 | (v5 / ((v5) | 1)))
                                        _ = fi19
                                    }
                                }
                            } else {
                                for fi27 := 0; fi27 < 15; fi27++ {
                                    v9 = v5
                                    v4 = v9
                                    var v28 uint32 = (v6 << ((0) & 63))
                                    _ = fi27
                                    _ = v28
                                }
                                if (!(v6 < 2)) {
                                    fmt.Println(v4)
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                } else {
                                    var v29 string = "hello"
                                    var v30 float32 = -2.25
                                    _ = v29
                                    _ = v30
                                }
                            }
                        }
                        _ = v10
                    } else {
                        var v31 string = "a"
                        _ = v31
                    }
                    _ = v8
                    _ = v9
                } else {
                    for fi32 := 0; fi32 < 1; fi32++ {
                        switch 137 {
                        case -471:
                            var v33 int64 = (v4 - v5)
                            _ = v33
                        case -323:
                            v7 = (v7 | 1)
                        case 721:
                            _ = len(v1)
                        default:
                            switch (("" + ((("a" + "foo bar") + "foo bar") + "hello")) + "foo bar") {
                            case "foo bar":
                                var v34 float64 = (v2 - v2)
                                _ = v34
                            case "":
                                v4 -= v5
                            case "hello":
                                v35 := GenS1{F0: (v7 << ((645) & 63)), F1: int32(v7)}
                                _ = v35
                            }
                        }
                        v36 := GenS3{F0: 228, F1: 73}
                        v6 -= (v6 - v6)
                        _ = fi32
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 20; fi37++ {
                        if ((2147483648 <= v6) || false) {
                            for fi38 := 0; fi38 < 20; fi38++ {
                                v1 = append(v1, (fi38 << ((42) & 63)))
                                var v39 uint8 = 2
                                _ = fi38
                                _ = v39
                            }
                            var v40 string = "foo bar"
                            _ = v40
                        }
                        fmt.Println(v3, fi37)
                        _ = fi37
                    }
                }
                var v41 uint = 0
                _ = v41
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v42 uint = 42
                _ = v42
            }
        case -9223372036854775808:
            v43 := GenS5{F0: ((380 >= 1) && (-101 < 89)), F1: v3, F2: -1}
            _ = v43
        }
        fmt.Println(v5)
        _ = v7
    } else {
        v44 := GenS6{F0: 0}
        var v45 int32 = -1
        var v46 bool = false
        _ = v44
        _ = v45
        _ = v46
    }
    if (false || ((!(true || false)) || true)) {
        var v47 int8 = 0
        _ = v47
    } else {
        v5 = (^(v5))
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = v5
    var v48 uint32 = (^(v6))
    v4 = (-(v5))
    v48 -= v6
    if ((42 > 65535) || (1e10 != 3.14159)) {
        var v49 uint8 = 80
        var v50 uint8 = (^(v49))
        _ = v49
        _ = v50
    } else {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v51 := GenS7{F0: 250}
            _ = v51
        }
    }
    v3 = 2
    v5 += v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
