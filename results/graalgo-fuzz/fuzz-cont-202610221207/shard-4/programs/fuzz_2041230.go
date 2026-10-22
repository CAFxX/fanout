package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int16; F1 uint32; F2 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 string; F1 uint64 }
type GenS10 struct { F0 int8; F1 int }

func main() {
    var v1 string = "a"
    v2 := []int{32767, 127, -414, 42}
    var v3 int = -1
    switch v3 {
    case 0:
        if true {
            v3 += (-(v3))
            v4 := GenS1{F0: 612}
            fmt.Println(v3, v1)
            _ = v4
        } else {
            for fi5 := 0; fi5 < 19; fi5++ {
                v3 += 2147483647
                _ = fi5
            }
            switch (v3 / ((-428) | 1)) {
            case 572:
                if (1e10 < 0.0) {
                    if (false || false) {
                        v6 := GenS1{F0: 558}
                        v2 = append(v2, 259)
                        v3 += -2147483648
                        _ = v6
                    }
                } else {
                    v1 = v1
                }
            case 42:
                for fi7 := 0; fi7 < 13; fi7++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v8 := GenS2{F0: 2147483647}
                    v3 = (v3 | (fi7 / ((v3) | 1)))
                    _ = fi7
                    _ = v8
                }
            case 128:
                for fi9 := 0; fi9 < 13; fi9++ {
                    fmt.Println(v1, fi9, v3)
                    v10 := GenS2{F0: 1}
                    _ = fi9
                    _ = v10
                }
            default:
                if (0 == 128) {
                    var v11 uint8 = 127
                    for fi12 := 0; fi12 < 8; fi12++ {
                        fmt.Println(v3, v1)
                        switch v1 {
                        case "foo bar":
                            for fi13 := 0; fi13 < 10; fi13++ {
                                fmt.Println(fi12, fi13)
                                var v14 uint32 = 383
                                if false {
                                    v14 -= v14
                                    v15 := GenS3{F0: 1000, F1: 0}
                                    _ = v15
                                } else {
                                    v3 = 564
                                    var v16 int8 = 1
                                    fmt.Println(fi13)
                                    _ = v16
                                }
                                _ = fi13
                                _ = v14
                            }
                        case "":
                            _ = len(v2)
                        case "hello":
                            if (true && (2147483648 <= 2147483647)) {
                                var v17 float64 = 3.14159
                                v18 := GenS2{F0: 725}
                                _ = v17
                                _ = v18
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        default:
                            fmt.Println(v11, v11, v3)
                        }
                        _ = fi12
                    }
                    var v19 int32 = 0
                    _ = v11
                    _ = v19
                }
            }
        }
    case -1:
        switch (v3 / ((128) | 1)) {
        case 1:
            if len(v2) > 0 {
                v2[0] = 32768
            }
        case 2147483647:
            v3 = v3
        }
    default:
        fmt.Println(v3)
    }
    var v20 float64 = float64(0.0)
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 17; fi22++ {
            v2 = append(v2, fi21)
            var v23 uint8 = 130
            v1 = v1
            if (true && true) {
                v20 = v20
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v1, v23)
            } else {
                v24 := GenS4{F0: 276, F1: 0, F2: (39 != 1)}
                v23 = (^(v23))
                switch v1 {
                case "hello":
                    fmt.Println(fi22, v3, v3)
                default:
                    v25 := GenS5{F0: 42}
                    _ = v25
                }
                v20 = v20
                _ = v24
            }
            _ = fi22
            _ = v23
        }
        for fi26 := 0; fi26 < 14; fi26++ {
            if false {
                v2 = append(v2, fi21)
                v3 -= v3
            } else {
                v1 = (v1 + v1)
                switch v1 {
                case "a":
                    v3 = -130
                    v2 = append(v2, fi26)
                case "hello":
                    v3 = fi21
                case "\n\t":
                    if false {
                        var v27 int16 = 32767
                        _ = v27
                    }
                }
                v28 := GenS6{F0: 667, F1: 128, F2: 1000}
                _ = v28
            }
            v29 := GenS7{F0: 437}
            v1 = "a"
            _ = fi26
            _ = v29
        }
        fmt.Println(v3)
        _ = fi21
    }
    v30 := GenS8{F0: 669}
    v31 := GenS9{F0: "\n\t", F1: 18446744073709551615}
    var v32 bool = (-2147483648 == 2147483647)
    for fi33 := 0; fi33 < 2; fi33++ {
        fmt.Println(v32, fi33)
        _ = fi33
    }
    var v34 float32 = (-2.25 / 1e10)
    fmt.Println(v20, v34, v34)
    v35 := GenS10{F0: 2, F1: (-(v3))}
    v32 = (!true)
    if (int8(v3) < -128) {
        v3 += 4294967295
        v3 = (-(v3))
    } else {
        v2 = append(v2, 979)
    }
    var v36 int64 = -9223372036854775808
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = (v3 + v3)
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        if len(v2) > 0 {
            v2[0] = fi37
        }
        fmt.Println(v34, v36, fi37)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi37
    }
    var v38 int = v3
    var v39 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v39
}
