package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 string; F2 int16 }
type GenS8 struct { F0 uint32; F1 int64; F2 uint64 }

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 989
    var v3 uint64 = 18446744073709551615
    v2 += (v3 >> ((18446744073709551615) & 63))
    v4 := []int{364}
    fmt.Println(v2, v3, v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        v5 := GenS1{F0: 974}
        v2 += (v2 ^ v2)
        v6 := GenS2{F0: 32767}
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := GenS2{F0: 1000}
            _ = fi7
            _ = v8
        }
        fmt.Println(v3, v2)
        _ = len(v4)
    }
    var v9 int = int(v2)
    v10 := GenS1{F0: 0}
    fmt.Println(v9)
    fmt.Println(v1)
    v11 := GenS3{F0: false}
    _ = len(v4)
    v1 = uint16(v2)
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v1, v1)
            if false {
                if len(v4) > 0 {
                    v4[0] = fi12
                }
            } else {
                fmt.Println(v3, fi13, v1)
                v4 = append(v4, 128)
            }
            if true {
                v3 -= (v3 | v3)
                if (42 < 303) {
                    v9 = -266
                    v14 := GenS4{F0: true, F1: false, F2: 205}
                    var v15 int64 = 643
                    var v16 float64 = float64(0.1)
                    var v17 uint = 255
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    if (2 <= 508) {
                        v2 += 1
                        if (!false) {
                            if true {
                                if true {
                                    v3 += v3
                                }
                                if (v2 >= (v3 + v3)) {
                                    v18 := GenS5{F0: (v3 >> ((2) & 63))}
                                    v4 = append(v4, fi13)
                                    _ = v18
                                } else {
                                    v3 += 0
                                    if (!true) {
                                        fmt.Println(v1, v2)
                                        if false {
                                            var v19 int = fi12
                                            v4 = append(v4, (fi13 / ((v19) | 1)))
                                            var v20 int16 = 0
                                            _ = v19
                                            _ = v20
                                        }
                                    }
                                    var v21 uint32 = 2147483647
                                    var v22 float64 = (1.5 / 3.14159)
                                    var v23 uint = 35
                                    var v24 int8 = 0
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                switch "a" {
                                case "":
                                    if ((-2.25 + 3.14159) <= (3.14159 * ((0.1 * 0.1) * 0.1))) {
                                        if (v1 > 65535) {
                                            var v25 bool = false
                                            v4 = append(v4, fi12)
                                            _ = v25
                                        } else {
                                            v2 -= (v3 - v2)
                                            fmt.Println(v2, v2, v1)
                                        }
                                        v26 := GenS2{F0: int16(v3)}
                                        var v27 int64 = 9223372036854775807
                                        _ = v26
                                        _ = v27
                                    }
                                }
                            } else {
                                if (-32768 == 754) {
                                    if ((0.1 >= 1e10) && true) {
                                        var v28 bool = (1.5 >= 1.5)
                                        if (v28 && v28) {
                                            v28 = (!(fi12 < (fi13 / ((-1) | 1))))
                                            v28 = v28
                                        } else {
                                            if v28 {
                                                v29 := GenS6{F0: 88}
                                                v30 := GenS1{F0: 294}
                                                v28 = true
                                                _ = v29
                                                _ = v30
                                            }
                                        }
                                        switch "foo bar" {
                                        case "":
                                            v31 := GenS5{F0: v3}
                                            _ = v31
                                        case "\n\t":
                                            v28 = (v28 || (!v28))
                                        case "foo bar":
                                            v1 = 1
                                        default:
                                            if (!false) {
                                                v32 := GenS7{F0: 76, F1: ("" + "a"), F2: 926}
                                                v28 = (127 >= 772)
                                                _ = v32
                                            } else {
                                                v2 = uint64(v1)
                                                var v33 uint = uint(v1)
                                                fmt.Println(fi12)
                                                _ = v33
                                            }
                                        }
                                        _ = v28
                                    }
                                    fmt.Println(v9, v1)
                                } else {
                                    fmt.Println(fi12)
                                    v3 -= v2
                                }
                                v4 = append(v4, 334)
                                if (!(42 == -491)) {
                                    var v34 bool = true
                                    if false {
                                        v4 = append(v4, (fi13 % ((2147483647) | 1)))
                                        var v35 int16 = 42
                                        _ = v35
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    _ = v34
                                } else {
                                    fmt.Println(v3)
                                }
                            }
                            fmt.Println(v3, v3)
                            if len(v4) > 0 {
                                v4[0] = fi12
                            }
                        }
                    } else {
                        v36 := GenS1{F0: -32768}
                        _ = v36
                    }
                    var v37 int8 = 42
                    _ = v37
                }
            }
            switch -1 {
            case 2147483647:
                fmt.Println(v2, v2)
            case 9223372036854775807:
                var v38 string = "hello"
                _ = v38
            case 32767:
                v3 = v3
            default:
                var v39 string = "a"
                _ = v39
            }
            v40 := GenS4{F0: (0 <= 128), F1: false, F2: 4294967295}
            v4 = append(v4, -421)
            var v41 uint32 = 1
            _ = fi13
            _ = v40
            _ = v41
        }
        switch "a" {
        case "hello":
            for fi42 := 0; fi42 < 9; fi42++ {
                var v43 float64 = 0.1
                v2 -= 4294967295
                _ = fi42
                _ = v43
            }
        default:
            v44 := GenS8{F0: 4294967295, F1: -316, F2: v3}
            _ = v44
        }
        v45 := GenS4{F0: false, F1: true, F2: 4294967295}
        _ = fi12
        _ = v45
    }
    var v46 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v9
}
