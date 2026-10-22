package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 string; F2 int8 }
type GenS5 struct { F0 int }

func main() {
    v1 := []string{("" + "foo bar"), "", "\n\t", "hello"}
    var v2 bool = (127 <= 127)
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    if false {
        v2 = true
    }
    var v5 int8 = 1
    var v6 uint32 = 589
    switch ("hello" + ("" + "")) {
    case "hello":
        var v7 int64 = -1
        _ = v7
    case "foo bar":
        for fi8 := 0; fi8 < 20; fi8++ {
            if len(v1) > 0 {
                v1[0] = ""
            }
            _ = fi8
        }
    case "a":
        v2 = v2
    }
    v2 = false
    var v9 int16 = 795
    if v2 {
        v2 = v2
    }
    delete(v3, v6)
    switch "a" {
    case "":
        for fi10 := 0; fi10 < 15; fi10++ {
            switch "" {
            case "a":
                if (-342 > -1) {
                    _ = len(v4)
                } else {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        _ = len(v4)
                        v12 := GenS1{F0: ("foo bar" + ("\n\t" + ("a" + "foo bar"))), F1: int16(v6), F2: 0}
                        v5 = v5
                        _ = fi11
                        _ = v12
                    }
                    if ("hello" != ((("" + "\n\t") + "hello") + "hello")) {
                        if v2 {
                            if true {
                                for fi13 := 0; fi13 < 2; fi13++ {
                                    v14 := GenS2{F0: v6, F1: v6, F2: 255}
                                    v4[("a" + "hello")] = v6
                                    _ = fi13
                                    _ = v14
                                }
                                v1 = append(v1, ("foo bar" + (("\n\t" + (("foo bar" + (("" + ("\n\t" + (("a" + "foo bar") + (("\n\t" + "\n\t") + "a")))) + "\n\t")) + ("\n\t" + "\n\t"))) + "\n\t")))
                            }
                        } else {
                            v15 := GenS1{F0: "", F1: 42, F2: (v6 ^ 4294967295)}
                            fmt.Println(v2, v6)
                            fmt.Println(v2)
                            _ = v15
                        }
                        v1 = append(v1, ("" + "foo bar"))
                    } else {
                        v5 = (v5 | 77)
                        var v16 int = 4294967295
                        _ = v16
                    }
                }
            case "foo bar":
                if v2 {
                    v2 = (v2 || (0 <= 1))
                    var v17 uint8 = 255
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v17
                }
            case "":
                var v18 int32 = 0
                _ = v18
            }
            _ = fi10
        }
    case "foo bar":
        var v19 float32 = (((((float32((float64(((float32(0.1) + (-2.25 * 0.0)) + 0.0)) / 1.0)) / (1e10 / ((0.1 - 3.14159) / 0.1))) / ((1.5 + 1e10) * (1.5 - 0.0))) * 3.14159) + (1.5 - 0.1)) / 1.0)
        _ = v19
    case "\n\t":
        var v20 int = -230
        _ = v20
    default:
        switch "a" {
        case "\n\t":
            var v21 int = -9223372036854775808
            _ = v21
        case "foo bar":
            if (!v2) {
                if v2 {
                    switch -48 {
                    case -924:
                        if false {
                            var v22 uint8 = 140
                            _ = v22
                        } else {
                            if (2147483647 >= -2147483648) {
                                for fi23 := 0; fi23 < 18; fi23++ {
                                    fmt.Println(v6, v9, v5)
                                    _ = len(v4)
                                    v9 = v9
                                    _ = fi23
                                }
                            } else {
                                var v24 uint64 = 4294967295
                                var v25 int64 = 42
                                _ = v24
                                _ = v25
                            }
                            var v26 uint16 = 2
                            _ = v26
                        }
                    default:
                        if false {
                            fmt.Println(v2, v5)
                            var v27 int8 = 56
                            _ = v27
                        } else {
                            v28 := GenS2{F0: (v6 >> ((42) & 63)), F1: v6, F2: 1}
                            fmt.Println(v6)
                            _ = v28
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                } else {
                    switch "" {
                    case "":
                        v5 = v5
                    case "foo bar":
                        for fi29 := 0; fi29 < 4; fi29++ {
                            _ = len(v3)
                            if (!v2) {
                                v5 -= (v5 * v5)
                            }
                            v9 -= v9
                            _ = fi29
                        }
                    default:
                        if ((v5 | v5) < -73) {
                            for fi30 := 0; fi30 < 10; fi30++ {
                                var v31 uint16 = uint16(v6)
                                _ = fi30
                                _ = v31
                            }
                            if (65535 < 238) {
                                for fi32 := 0; fi32 < 7; fi32++ {
                                    v33 := GenS3{F0: -2147483648}
                                    for fi34 := 0; fi34 < 5; fi34++ {
                                        switch "a" {
                                        case "\n\t":
                                            v9 -= 32767
                                        case "hello":
                                            var v35 float32 = 1e10
                                            _ = v35
                                        }
                                        v36 := GenS4{F0: 2, F1: ("foo bar" + "foo bar"), F2: v5}
                                        v2 = (!(4294967295 == v6))
                                        var v37 string = "foo bar"
                                        v9 = v9
                                        var v38 uint64 = 0
                                        v5 -= (-(v5))
                                        switch 9223372036854775807 {
                                        case 288:
                                            v39 := GenS2{F0: 0, F1: v6, F2: 0}
                                            _ = v39
                                        case -9223372036854775808:
                                            var v40 float64 = 0.1
                                            _ = v40
                                        }
                                        _ = fi34
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                    v2 = (true || (!v2))
                                    _ = fi32
                                    _ = v33
                                }
                            }
                            fmt.Println(v5, v5)
                        } else {
                            v41 := GenS3{F0: 764}
                            for fi42 := 0; fi42 < 12; fi42++ {
                                var v43 uint16 = 32767
                                v9 = (v9 >> ((4294967295) & 63))
                                _ = fi42
                                _ = v43
                            }
                            _ = v41
                        }
                    }
                    fmt.Println(v6)
                    switch 128 {
                    case 4294967295:
                        v44 := GenS1{F0: "", F1: v9, F2: v6}
                        _ = v44
                    }
                }
            }
        default:
            if len(v1) > 0 {
                v1[0] = "a"
            }
        }
    }
    var v45 float32 = -2.25
    var v46 uint = 653
    for fi47 := 0; fi47 < 20; fi47++ {
        var v48 uint8 = 1
        switch ((("hello" + "hello") + "") + "") {
        case "foo bar":
            if false {
                fmt.Println(v48, v46)
                v46 += (v46 % ((255) | 1))
                var v49 int64 = 9223372036854775807
                _ = v49
            } else {
                v9 = (-(v9))
                switch "a" {
                case "\n\t":
                    delete(v4, (("a" + "") + ("" + "\n\t")))
                case "foo bar":
                    v50 := GenS5{F0: fi47}
                    _ = v50
                case "":
                    _ = len(v3)
                }
            }
        default:
            for fi51 := 0; fi51 < 12; fi51++ {
                var v52 int = fi51
                _ = fi51
                _ = v52
            }
        }
        var v53 float32 = (((v45 / float32(1.5)) - v45) * v45)
        _ = fi47
        _ = v48
        _ = v53
    }
    v45 = (v45 + (v45 * ((((v45 + v45) * float32((0.1 * float64(1e10)))) / v45) - (-2.25 - float32(1.5)))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v9
}
