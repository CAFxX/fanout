package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int; F1 bool; F2 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 bool; F2 uint32 }

func main() {
    var v1 bool = true
    var v2 int64 = -9223372036854775808
    v3 := make(map[uint32]string)
    v4 := []int{1, 128, -759, 2147483647}
    v5 := make(map[uint32]string)
    switch (("a" + "\n\t") + (("a" + "") + "hello")) {
    case "hello":
        if (0 >= -32768) {
            if v1 {
                if ((^(v2)) <= v2) {
                    if ((!v1) || false) {
                        if v1 {
                            v6 := GenS1{F0: 4294967295}
                            _ = v6
                        } else {
                            for fi7 := 0; fi7 < 18; fi7++ {
                                for fi8 := 0; fi8 < 4; fi8++ {
                                    fmt.Println(fi8, fi7)
                                    if ((fi7 << ((255) & 63)) <= (fi7 | (fi7 & (fi7 * fi8)))) {
                                        v5[2147483648] = "foo bar"
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if (fi7 == (^(fi8))) {
                                            var v9 int32 = 2147483647
                                            v9 = 2147483647
                                            v2 += (-(v2))
                                            _ = v9
                                        }
                                    }
                                    v1 = true
                                    v2 = v2
                                    if (18446744073709551615 < 94) {
                                        var v10 uint64 = 4294967295
                                        _ = v10
                                    }
                                    _ = fi8
                                }
                                fmt.Println(v2, v2)
                                _ = fi7
                            }
                            fmt.Println(v2)
                            v11 := GenS2{F0: 251, F1: 932, F2: 0}
                            _ = v11
                        }
                        var v12 int = 127
                        switch v12 {
                        case 42:
                            if v1 {
                                for fi13 := 0; fi13 < 8; fi13++ {
                                    v14 := GenS1{F0: 539}
                                    v15 := GenS2{F0: 0, F1: 42, F2: 152}
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                                v4 = append(v4, (v12 * v12))
                            }
                        default:
                            _ = len(v3)
                        }
                        _ = v12
                    } else {
                        var v16 bool = (!(0.1 <= ((1e10 / 1.5) + float64((1.5 / (3.14159 + 1.5))))))
                        if (v1 || v1) {
                            v17 := GenS1{F0: 4294967295}
                            _ = v17
                        }
                        _ = v16
                    }
                } else {
                    var v18 float64 = float64((((1e10 * ((-2.25 / 0.1) - ((0.0 + float32(0.0)) / 1.5))) - 0.1) * 0.1))
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v18
                }
                _ = len(v3)
            }
            v1 = v1
            for fi19 := 0; fi19 < 8; fi19++ {
                v2 = v2
                v1 = (v2 == v2)
                _ = fi19
            }
        }
    case "\n\t":
        if true {
            v20 := GenS2{F0: 128, F1: 1000, F2: 127}
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 uint32 = 500
        var v23 float64 = float64(0.1)
        _ = fi21
        _ = v22
        _ = v23
    }
    if true {
        if v1 {
            if (false && (509 != 4294967295)) {
                fmt.Println(v2, v2)
            }
        }
    } else {
        v24 := GenS3{F0: 854, F1: 71}
        v1 = v1
        _ = v24
    }
    var v25 int = -444
    v26 := GenS4{F0: v25, F1: (!((((((0.0 / ((0.1 + 0.1) * -2.25)) / (((float64(1.5) - ((1e10 + -2.25) + (1e10 / 1.5))) / 1.0) / (((1.5 / 0.1) / 1.5) * 0.1))) + 1e10) + float64((float32(1.5) + (0.1 - -2.25)))) - (float64(1.5) - (1.5 * (3.14159 / 3.14159)))) != float64(-2.25))), F2: -32768}
    fmt.Println(v2)
    var v27 int32 = 0
    fmt.Println(v27)
    if (18446744073709551615 != 42) {
        v4 = append(v4, (-(v25)))
    } else {
        v2 = (v2 ^ v2)
    }
    var v28 int32 = (v27 >> ((0) & 63))
    for fi29 := 0; fi29 < 20; fi29++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi29
    }
    v30 := GenS1{F0: 325}
    fmt.Println(v25, v28)
    switch "a" {
    case "":
        if v1 {
            switch "hello" {
            case "a":
                for fi31 := 0; fi31 < 6; fi31++ {
                    var v32 uint16 = 42
                    _ = fi31
                    _ = v32
                }
            }
        } else {
            v33 := GenS5{F0: 32767}
            _ = v33
        }
    case "a":
        v34 := GenS6{F0: 18446744073709551615, F1: ((v25 >> ((42) & 63)) != 0), F2: 4294967295}
        _ = v34
    }
    fmt.Println(v28, v1, v2)
    if v1 {
        switch "foo bar" {
        case "foo bar":
            _ = len(v4)
        case "\n\t":
            switch "a" {
            case "":
                v3[0] = "foo bar"
            case "hello":
                fmt.Println(v28)
            default:
                var v35 int16 = 32767
                _ = v35
            }
        }
        v28 += -2147483648
    } else {
        switch ((("a" + (("hello" + ("hello" + "")) + ("a" + "\n\t"))) + "") + "") {
        case "a":
            if (0 >= 1) {
                v25 = v25
            }
        }
        var v36 int32 = 0
        switch (v25 + (-(v25))) {
        case 127:
            _ = len(v4)
        case 4294967295:
            for fi37 := 0; fi37 < 8; fi37++ {
                fmt.Println(v36, v1, v27)
                for fi38 := 0; fi38 < 20; fi38++ {
                    var v39 uint16 = 814
                    _ = fi38
                    _ = v39
                }
                _ = fi37
            }
        case 814:
            for fi40 := 0; fi40 < 15; fi40++ {
                var v41 uint = 255
                v28 = v36
                _ = fi40
                _ = v41
            }
        default:
            var v42 int32 = v36
            _ = v42
        }
        _ = v36
    }
    v43 := GenS2{F0: 1, F1: 32767, F2: 134}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v5
}
