package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint }
func genFn0(p0 int, p1 uint64) uint32 {
    fmt.Println(p0)
    v4 := []int{}
    if true {
        fmt.Println(p1, p0)
        for fi5 := 0; fi5 < 8; fi5++ {
            p1 = p1
            p1 = p1
            v4 = append(v4, fi5)
            _ = fi5
        }
    }
    var v6 bool = (0.0 == 1.5)
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 65535
}
func genFn1(p0 int32) uint32 {
    if false {
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := []int{575, (fi7 >> ((0) & 63)), -946}
            var v9 float64 = (-2.25 / 1.0)
            if (0.0 <= (0.1 + ((3.14159 / 1.5) + 1e10))) {
                v8 = append(v8, fi7)
                switch "foo bar" {
                case "a":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                default:
                    if false {
                        if ("hello" != "\n\t") {
                            var v10 float32 = ((0.0 + 1e10) / float32(v9))
                            v11 := GenS1{F0: 100, F1: 255}
                            _ = v10
                            _ = v11
                        } else {
                            switch "hello" {
                            case "foo bar":
                                v9 -= (float64(1e10) + float64((1e10 / -2.25)))
                            case "a":
                                var v12 uint8 = 127
                                _ = v12
                            }
                        }
                        v13 := GenS2{F0: 42, F1: ((-1 > -353) && (-1 != -61)), F2: 212}
                        var v14 int = -439
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint32 = 2147483648
                        var v16 string = "\n\t"
                        _ = v15
                        _ = v16
                    }
                }
                p0 += (p0 >> ((2) & 63))
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        v17 := []int{2147483647, 2147483647, 32767, -328}
        fmt.Println(p0)
        _ = v17
    } else {
        var v18 uint = 941
        switch (("hello" + ("" + ("\n\t" + "foo bar"))) + "foo bar") {
        case "a":
            p0 = 42
        case "hello":
            var v19 int8 = -128
            _ = v19
        case "":
            for fi20 := 0; fi20 < 14; fi20++ {
                v18 += (v18 | (v18 % ((v18) | 1)))
                fmt.Println(p0)
                v21 := GenS1{F0: 42, F1: 30}
                _ = fi20
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 8; fi22++ {
            if (!true) {
                for fi23 := 0; fi23 < 6; fi23++ {
                    switch (fi22 - (fi23 | 127)) {
                    case -1:
                        var v24 int8 = -127
                        _ = v24
                    case 32768:
                        v18 -= 18446744073709551615
                    case 1:
                        if ((1 != 2147483647) && (!(4294967295 == 592))) {
                            fmt.Println(p0)
                            fmt.Println(v18, fi23, fi22)
                            v25 := GenS2{F0: 2, F1: (float32((1.5 + 0.0)) <= float32((0.0 - 3.14159))), F2: 42}
                            _ = v25
                        }
                    default:
                        p0 -= (p0 & (p0 % ((p0) | 1)))
                    }
                    v26 := GenS3{F0: 0, F1: -52}
                    _ = fi23
                    _ = v26
                }
            }
            _ = fi22
        }
        _ = v18
    }
    v27 := []int{128, -78}
    switch 127 {
    case 0:
        if len(v27) >= 2 {
            v27 = v27[1:len(v27)]
        }
    case 551:
        p0 = (p0 << ((0) & 63))
    default:
        v28 := GenS4{F0: ((("" + ("hello" + "")) + "hello") + ("foo bar" + "foo bar")), F1: 4294967295}
        _ = v28
    }
    p0 -= p0
    switch 4294967295 {
    case -1:
        var v29 uint64 = 2
        _ = v29
    case 32768:
        _ = len(v27)
    default:
        if (1000 >= 42) {
            p0 = p0
            switch "\n\t" {
            case "hello":
                var v30 int8 = 127
                _ = v30
            case "foo bar":
                for fi31 := 0; fi31 < 16; fi31++ {
                    p0 = (p0 - -2147483648)
                    for fi32 := 0; fi32 < 9; fi32++ {
                        fmt.Println(fi31)
                        var v33 float32 = (1e10 + 1e10)
                        v33 -= 3.14159
                        _ = fi32
                        _ = v33
                    }
                    v27 = append(v27, (fi31 ^ 862))
                    _ = fi31
                }
            case "\n\t":
                fmt.Println(p0)
            default:
                var v34 int8 = 0
                _ = v34
            }
            _ = len(v27)
        }
    }
    _ = p0
    _ = v27
    return 417
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 255
    var v3 string = "hello"
    v3 = (v3 + v3)
    v35 := []int{0}
    var v36 int64 = int64(v2)
    v37 := GenS1{F0: 127, F1: 128}
    for fi38 := 0; fi38 < 8; fi38++ {
        v35 = append(v35, -9223372036854775808)
        var v39 int8 = 127
        _ = fi38
        _ = v39
    }
    v40 := GenS5{F0: (v2 * v2), F1: 443}
    var v41 int8 = -96
    switch "a" {
    case "\n\t":
        if true {
            v36 = 634
        } else {
            v3 = v3
        }
    case "":
        v41 = (v41 << ((65535) & 63))
    default:
        var v42 int16 = 965
        _ = v42
    }
    var v43 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v40
    _ = v41
    _ = v43
}
