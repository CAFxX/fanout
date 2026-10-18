package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int8; F2 int }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]uint32)
    var v3 int8 = 127
    v4 := make(map[uint32]bool)
    var v5 int64 = 0
    var v6 int32 = 1
    if (-1 >= 0) {
        switch "" {
        case "foo bar":
            v7 := GenS1{F0: 42}
            _ = v7
        }
        v1 -= v1
        var v8 uint8 = 255
        _ = v8
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            if (1 > 42) {
                fmt.Println(v3)
                for fi10 := 0; fi10 < 3; fi10++ {
                    switch ((("" + ((("" + ("a" + "a")) + "foo bar") + "hello")) + "") + "\n\t") {
                    case "foo bar":
                        v1 = 0.0
                    case "a":
                        var v11 int16 = -142
                        _ = v11
                    default:
                        v6 = v6
                    }
                    _ = fi10
                }
            } else {
                delete(v4, 2)
            }
            _ = fi9
        }
        fmt.Println(v1, v5, v3)
    }
    v3 = v3
    v5 = (^(v5))
    var v12 int64 = (v5 | -1)
    v13 := []int{127, -425}
    if (!false) {
        for fi14 := 0; fi14 < 5; fi14++ {
            _ = len(v4)
            fmt.Println(v5, v6, v1)
            _ = fi14
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 int8 = v3
            for fi17 := 0; fi17 < 19; fi17++ {
                if true {
                    v12 = -254
                    var v18 float32 = -2.25
                    _ = v18
                }
                var v19 string = "\n\t"
                _ = fi17
                _ = v19
            }
            if (255 == 18446744073709551615) {
                if (32767 != 255) {
                    v20 := GenS2{F0: 1000, F1: 65535}
                    v21 := GenS3{F0: v16, F1: 22}
                    _ = v20
                    _ = v21
                } else {
                    var v22 float32 = 3.14159
                    var v23 int64 = -829
                    if true {
                        v6 = (v6 * v6)
                        for fi24 := 0; fi24 < 12; fi24++ {
                            switch ((("hello" + "") + "\n\t") + ("\n\t" + "\n\t")) {
                            case "":
                                fmt.Println(fi24, v22)
                            case "\n\t":
                                v25 := GenS3{F0: (^(v3)), F1: 127}
                                _ = v25
                            case "hello":
                                if (v6 >= (v6 >> ((42) & 63))) {
                                    v6 -= (v6 / (((v6 % ((-653) | 1))) | 1))
                                    v26 := GenS4{F0: 0, F1: -67, F2: fi24}
                                    _ = v26
                                }
                            default:
                                v22 = 0.1
                            }
                            v3 += 42
                            _ = fi24
                        }
                    }
                    _ = v22
                    _ = v23
                }
                fmt.Println(v1)
            }
            _ = fi15
            _ = v16
        }
        v3 += v3
    }
    v1 -= (3.14159 / v1)
    v13 = append(v13, 2147483647)
    if true {
        v3 = v3
        for fi27 := 0; fi27 < 11; fi27++ {
            for fi28 := 0; fi28 < 18; fi28++ {
                if ((!false) && ((2 > 520) || (2 >= 4294967295))) {
                    var v29 uint16 = 32767
                    _ = v29
                } else {
                    var v30 uint8 = 24
                    _ = v30
                }
                _ = len(v4)
                _ = fi28
            }
            _ = fi27
        }
    } else {
        v5 -= -1
        v31 := GenS1{F0: 128}
        _ = v31
    }
    var v32 int32 = (v6 * 2147483647)
    if (!false) {
        fmt.Println(v6, v6, v3)
        v3 += (-(v3))
    } else {
        for fi33 := 0; fi33 < 18; fi33++ {
            if true {
                var v34 string = ("" + "")
                _ = len(v2)
                for fi35 := 0; fi35 < 10; fi35++ {
                    v6 = v6
                    v1 -= v1
                    _ = fi35
                }
                _ = v34
            }
            var v36 float64 = (v1 * (-2.25 / float64((((((-2.25 / -2.25) - 0.1) / (1.5 * (-2.25 * 1e10))) - 1e10) * 0.1))))
            v3 = -128
            _ = fi33
            _ = v36
        }
        var v37 bool = ((255 < 4294967295) || (!false))
        _ = v37
    }
    switch "hello" {
    case "":
        for fi38 := 0; fi38 < 9; fi38++ {
            var v39 bool = (-2.25 != (1.5 - 3.14159))
            var v40 int32 = v6
            var v41 int32 = (v40 << ((4294967295) & 63))
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
    case "\n\t":
        if (9223372036854775807 != v12) {
            var v42 uint8 = 127
            if false {
                delete(v4, 4294967295)
                var v43 string = "\n\t"
                _ = v43
            }
            var v44 int32 = v6
            _ = v42
            _ = v44
        }
    }
    v45 := GenS1{F0: 0}
    fmt.Println(v5, v5)
    fmt.Println(v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
