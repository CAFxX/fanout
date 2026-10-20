package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS5 struct { F0 bool; F1 int32; F2 int8 }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-853, -41, -1, 127}
    v3 := make(map[int]bool)
    var v4 uint16 = 898
    fmt.Println(v4)
    v4 += v4
    for fi5 := 0; fi5 < 13; fi5++ {
        v4 = (v4 & 42)
        v4 = (v4 >> ((18446744073709551615) & 63))
        switch fi5 {
        case -1:
            if (42 <= 868) {
                var v6 int16 = 42
                v4 -= v4
                v7 := GenS1{F0: -1, F1: 128}
                _ = v6
                _ = v7
            } else {
                v8 := GenS2{F0: 2, F1: "a"}
                _ = v8
            }
        default:
            if false {
                v2 = append(v2, -2147483648)
                fmt.Println(v4)
            } else {
                for fi9 := 0; fi9 < 3; fi9++ {
                    v4 = 32767
                    v2 = append(v2, 127)
                    switch (fi5 - fi9) {
                    case 42:
                        var v10 float32 = 0.1
                        _ = v10
                    case -2147483648:
                        var v11 int = 587
                        _ = v11
                    }
                    _ = fi9
                }
                delete(v3, fi5)
                switch (fi5 << ((4294967295) & 63)) {
                case 713:
                    var v12 uint8 = 107
                    _ = v12
                case 2147483647:
                    var v13 uint8 = 71
                    _ = v13
                default:
                    if true {
                        for fi14 := 0; fi14 < 13; fi14++ {
                            v15 := GenS3{F0: 0, F1: 18446744073709551615}
                            v16 := GenS1{F0: 9223372036854775807, F1: 255}
                            var v17 int64 = 946
                            _ = fi14
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        switch "a" {
                        case "foo bar":
                            switch "hello" {
                            case "a":
                                switch fi5 {
                                case 0:
                                    _ = len(v2)
                                }
                            case "\n\t":
                                v18 := GenS1{F0: -760, F1: 128}
                                _ = v18
                            }
                        case "hello":
                            switch ("\n\t" + ("foo bar" + "a")) {
                            case "\n\t":
                                if false {
                                    for fi19 := 0; fi19 < 20; fi19++ {
                                        v20 := GenS4{F0: 111, F1: (v4 * 11), F2: 9223372036854775807}
                                        _ = fi19
                                        _ = v20
                                    }
                                    var v21 uint = 4294967295
                                    var v22 int64 = 0
                                    _ = v21
                                    _ = v22
                                }
                            default:
                                var v23 string = "\n\t"
                                _ = v23
                            }
                        case "\n\t":
                            if false {
                                _ = len(v2)
                                var v24 int32 = -1
                                v25 := GenS5{F0: false, F1: (v24 + v24), F2: 76}
                                _ = v24
                                _ = v25
                            } else {
                                switch ("\n\t" + ("foo bar" + "foo bar")) {
                                case "hello":
                                    var v26 int64 = 858
                                    _ = v26
                                case "":
                                    v1[fi5] = true
                                case "\n\t":
                                    if ((!true) && false) {
                                        var v27 float64 = -2.25
                                        _ = v27
                                    }
                                }
                            }
                        }
                        var v28 uint64 = 613
                        _ = v28
                    }
                }
            }
        }
        _ = fi5
    }
    _ = len(v2)
    _ = len(v2)
    v1[32767] = true
    v4 += v4
    v4 -= (v4 >> ((255) & 63))
    var v29 uint32 = 42
    if (true && true) {
        v2 = append(v2, 32767)
        if (((0.1 / 0.1) + (1e10 + float64(float32(((float64(float32(1e10)) * 0.0) + float64(1e10)))))) > -2.25) {
            var v30 float32 = 3.14159
            for fi31 := 0; fi31 < 6; fi31++ {
                var v32 float64 = 1e10
                _ = fi31
                _ = v32
            }
            _ = v30
        } else {
            v29 = v29
            if true {
                switch "a" {
                case "a":
                    var v33 uint16 = (v4 % ((v4) | 1))
                    _ = v33
                }
                fmt.Println(v29, v29)
            }
            v29 += 4294967295
        }
        v29 = v29
    } else {
        v29 += v29
        switch -475 {
        case 606:
            _ = len(v2)
        default:
            var v34 bool = false
            _ = v34
        }
    }
    var v35 int = 42
    switch v35 {
    case 716:
        if (v4 > v4) {
            v4 = (^(v4))
        }
    case 4294967295:
        delete(v1, 920)
    default:
        var v36 int = (v35 ^ -2147483648)
        _ = v36
    }
    var v37 float64 = 0.0
    v38 := GenS6{F0: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v4
}
