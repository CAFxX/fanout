package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    var v4 float64 = 1.5
    v4 = (v4 + float64(1e10))
    v5 := GenS1{F0: ("a" + ("hello" + (("a" + "foo bar") + "\n\t"))), F1: 255, F2: -9223372036854775808}
    v6 := []int{}
    fmt.Println(v4)
    var v7 uint64 = 18446744073709551615
    if (!(128 != 42)) {
        switch (("foo bar" + "\n\t") + "a") {
        case "foo bar":
            var v8 uint8 = 239
            _ = v8
        case "":
            if (9 <= 2147483648) {
                v4 = v4
            }
        case "a":
            switch 264 {
            case 32768:
                if false {
                    var v9 float32 = ((0.0 + -2.25) - 0.1)
                    _ = v9
                }
            case -1:
                var v10 uint32 = 2147483647
                _ = v10
            }
        default:
            v4 = v4
        }
    } else {
        v7 -= (^(v7))
        fmt.Println(v4, v7)
    }
    v11 := GenS1{F0: ("foo bar" + "a"), F1: 2, F2: 9223372036854775807}
    v4 -= v4
    v1[""] = "foo bar"
    var v12 int = 32768
    if (false && true) {
        if (1000 == 1) {
            for fi13 := 0; fi13 < 4; fi13++ {
                fmt.Println(fi13, fi13)
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 int8 = 1
                    _ = fi14
                    _ = v15
                }
                switch "\n\t" {
                case "hello":
                    switch "a" {
                    case "\n\t":
                        for fi16 := 0; fi16 < 15; fi16++ {
                            var v17 uint64 = (v7 * v7)
                            v18 := GenS2{F0: v7}
                            var v19 int16 = 1000
                            _ = fi16
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                    case "a":
                        if (-1 == -669) {
                            v20 := GenS3{F0: fi13}
                            if true {
                                v21 := GenS4{F0: 60, F1: false}
                                var v22 uint32 = 483
                                var v23 int32 = 897
                                _ = v21
                                _ = v22
                                _ = v23
                            } else {
                                v6 = append(v6, 42)
                                var v24 uint8 = 72
                                fmt.Println(v24, v12, v7)
                                _ = v24
                            }
                            _ = v20
                        }
                    }
                case "a":
                    fmt.Println(v12, v7, v12)
                case "":
                    for fi25 := 0; fi25 < 12; fi25++ {
                        var v26 int16 = 1000
                        switch "\n\t" {
                        case "foo bar":
                            if ((1 <= 979) && ((false || (!true)) && (-1 >= 292))) {
                                v12 -= (fi25 << ((18446744073709551615) & 63))
                            } else {
                                fmt.Println(v26, v12, v7)
                                if len(v6) > 0 {
                                    v6[0] = fi13
                                }
                            }
                        case "a":
                            if (!false) {
                                switch ("" + (("\n\t" + ("hello" + "a")) + "\n\t")) {
                                case "hello":
                                    if false {
                                        v12 = v12
                                        var v27 uint32 = 0
                                        _ = v27
                                    }
                                }
                            } else {
                                if (v4 < (-2.25 - v4)) {
                                    v4 = v4
                                } else {
                                    if false {
                                        delete(v1, (("foo bar" + "hello") + "a"))
                                    } else {
                                        if (2147483647 >= 65535) {
                                            if false {
                                                var v28 bool = (v12 < 9223372036854775807)
                                                v4 -= (1e10 * (v4 - (float64((float32(((1.5 - v4) + ((float64(3.14159) * v4) * (((v4 * float64((float32(v4) * 1e10))) / 1.5) - v4)))) * ((-2.25 * 1e10) / 0.1))) + v4)))
                                                if ((!v28) && v28) {
                                                    v26 = v26
                                                    var v29 bool = (((0 != 764) || (v4 > v4)) || v28)
                                                    v30 := GenS5{F0: 2}
                                                    _ = v29
                                                    _ = v30
                                                } else {
                                                    fmt.Println(v28, v12, v7)
                                                    var v31 float64 = v4
                                                    fmt.Println(v7, v26, v28)
                                                    v12 += 9223372036854775807
                                                    _ = v31
                                                }
                                                _ = v28
                                            } else {
                                                v26 -= v26
                                                v32 := GenS6{F0: 42}
                                                v26 = v26
                                                v7 = (^(v7))
                                                var v33 int32 = 1
                                                _ = v32
                                                _ = v33
                                            }
                                        } else {
                                            v34 := GenS4{F0: 2147483647, F1: true}
                                            v26 = (v26 + 32767)
                                            v35 := GenS2{F0: 18446744073709551615}
                                            _ = v34
                                            _ = v35
                                        }
                                        fmt.Println(fi13, v7, v12)
                                    }
                                    if (v7 == v7) {
                                        v26 += (v26 / ((746) | 1))
                                    }
                                    var v36 uint32 = 224
                                    _ = v36
                                }
                                v37 := GenS3{F0: int(v7)}
                                _ = v37
                            }
                        case "":
                            _ = len(v6)
                        default:
                            v12 = (fi25 - fi25)
                        }
                        fmt.Println(v12, v4, v26)
                        _ = fi25
                        _ = v26
                    }
                default:
                    var v38 float64 = v4
                    _ = v38
                }
                _ = fi13
            }
            for fi39 := 0; fi39 < 15; fi39++ {
                fmt.Println(fi39, v4, v7)
                _ = fi39
            }
            switch v12 {
            case -2147483648:
                fmt.Println(v4)
            case 9223372036854775807:
                v12 = (-(v12))
            case 32768:
                var v40 float64 = 3.14159
                _ = v40
            }
        } else {
            delete(v2, 2)
        }
        switch "a" {
        case "foo bar":
            var v41 float64 = (v4 + (((v4 + (v4 * 0.1)) - (v4 * (1e10 / v4))) * (v4 * 1.5)))
            _ = v41
        default:
            switch "hello" {
            case "foo bar":
                for fi42 := 0; fi42 < 4; fi42++ {
                    for fi43 := 0; fi43 < 1; fi43++ {
                        v6 = append(v6, fi43)
                        _ = fi43
                    }
                    _ = fi42
                }
            case "hello":
                var v44 bool = (!false)
                _ = v44
            }
        }
    }
    switch ("\n\t" + (("" + "a") + "")) {
    case "":
        v12 -= (v12 / ((v12) | 1))
    case "foo bar":
        v45 := GenS1{F0: "\n\t", F1: 255, F2: v12}
        _ = v45
    default:
        v46 := GenS1{F0: "a", F1: 1, F2: v12}
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
