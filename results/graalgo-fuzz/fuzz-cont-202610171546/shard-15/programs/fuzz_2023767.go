package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 bool; F1 int16; F2 int }
type GenS6 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -2147483648
    v3 := []uint32{}
    v4 := []int{9223372036854775807, -2147483648, -975, 2147483647}
    v1[("foo bar" + (("\n\t" + ((("\n\t" + ("hello" + "\n\t")) + "hello") + "a")) + "foo bar"))] = ""
    switch ("" + "\n\t") {
    case "foo bar":
        v1["a"] = ("hello" + ("" + "foo bar"))
    case "":
        for fi5 := 0; fi5 < 14; fi5++ {
            if (255 >= 1) {
                v6 := GenS1{F0: "\n\t", F1: 1000}
                if false {
                    v7 := GenS2{F0: 2}
                    fmt.Println(v2, fi5)
                    for fi8 := 0; fi8 < 7; fi8++ {
                        v2 -= v2
                        fmt.Println(fi8)
                        v9 := GenS3{F0: int64(v2), F1: 2147483648, F2: 2}
                        _ = fi8
                        _ = v9
                    }
                    _ = v7
                } else {
                    fmt.Println(fi5)
                    if false {
                        if ((fi5 << ((255) & 63)) <= (fi5 / ((fi5) | 1))) {
                            var v10 uint = 690
                            _ = v10
                        } else {
                            for fi11 := 0; fi11 < 6; fi11++ {
                                v2 = -2147483648
                                var v12 uint = 42
                                _ = fi11
                                _ = v12
                            }
                            v13 := GenS4{F0: -128, F1: 422}
                            if (!(0 > 42)) {
                                for fi14 := 0; fi14 < 12; fi14++ {
                                    v15 := GenS3{F0: -9223372036854775808, F1: 2, F2: 858}
                                    _ = fi14
                                    _ = v15
                                }
                            }
                            _ = v13
                        }
                    } else {
                        v2 += (v2 ^ (v2 + (v2 & 1)))
                        v2 = -1
                        if true {
                            switch "a" {
                            case "\n\t":
                                var v16 int16 = 504
                                _ = v16
                            default:
                                if false {
                                    v2 = v2
                                    fmt.Println(fi5, fi5)
                                } else {
                                    if (3.14159 >= (1.5 - 1e10)) {
                                        if (!false) {
                                            for fi17 := 0; fi17 < 18; fi17++ {
                                                var v18 uint16 = 2
                                                fmt.Println(v2)
                                                _ = fi17
                                                _ = v18
                                            }
                                        } else {
                                            fmt.Println(fi5, v2)
                                            fmt.Println(v2)
                                        }
                                        var v19 uint64 = 4294967295
                                        _ = v19
                                    }
                                    v2 -= v2
                                    if false {
                                        var v20 float32 = ((1.5 * 0.1) + 1e10)
                                        _ = v20
                                    }
                                }
                            }
                        }
                    }
                }
                var v21 int32 = v2
                _ = v21
                _ = v6
            } else {
                var v22 uint64 = 18446744073709551615
                _ = v22
            }
            for fi23 := 0; fi23 < 8; fi23++ {
                v2 += (v2 ^ v2)
                v24 := GenS5{F0: (2147483648 >= 2147483648), F1: 1, F2: 550}
                _ = fi23
                _ = v24
            }
            v2 += v2
            _ = fi5
        }
    default:
        v2 = 0
    }
    v1["\n\t"] = ("foo bar" + "a")
    for fi25 := 0; fi25 < 19; fi25++ {
        var v26 int = -633
        for fi27 := 0; fi27 < 18; fi27++ {
            if ((!false) || ((!(!(1 == 1))) && true)) {
                if (!true) {
                    fmt.Println(v26)
                    var v28 uint = 255
                    var v29 float64 = -2.25
                    _ = v28
                    _ = v29
                }
            }
            _ = len(v3)
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    if (-2.25 < 1.5) {
        switch 383 {
        case 127:
            var v30 uint32 = 0
            _ = v30
        }
        for fi31 := 0; fi31 < 13; fi31++ {
            v2 = -551
            _ = fi31
        }
    } else {
        v2 = v2
    }
    v2 = (-(v2))
    v32 := GenS4{F0: 42, F1: -9223372036854775808}
    var v33 int8 = 42
    v34 := GenS6{F0: 2, F1: 853}
    var v35 float32 = 1e10
    v35 += (((v35 * (v35 - (-2.25 * 0.0))) / 1.5) + v35)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v36 uint8 = 91
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
