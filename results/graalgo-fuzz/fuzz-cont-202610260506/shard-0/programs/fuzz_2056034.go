package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 uint32; F2 uint }
type GenS6 struct { F0 int8; F1 int8; F2 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{((((("foo bar" + "hello") + "\n\t") + ("foo bar" + "")) + ("\n\t" + "foo bar")) + "\n\t"), "a", "", ("a" + "\n\t")}
    v3 := make(map[int]int)
    var v4 uint8 = 255
    var v5 uint = 65535
    var v6 uint64 = 2
    if (2147483647 <= 1) {
        for fi7 := 0; fi7 < 2; fi7++ {
            if (824 > 0) {
                switch "" {
                case "a":
                    v8 := GenS1{F0: (!false), F1: 101}
                    _ = v8
                case "foo bar":
                    var v9 uint16 = 1
                    _ = v9
                case "":
                    switch "hello" {
                    case "\n\t":
                        var v10 uint64 = v6
                        _ = v10
                    default:
                        v4 -= uint8(v5)
                    }
                default:
                    v6 -= (v6 * v6)
                }
            }
            v4 -= v4
            switch ("\n\t" + (("foo bar" + "hello") + "hello")) {
            case "hello":
                if (42 > -795) {
                    v11 := GenS1{F0: false, F1: int8(v5)}
                    v12 := GenS2{F0: 0}
                    _ = v11
                    _ = v12
                }
            case "a":
                for fi13 := 0; fi13 < 16; fi13++ {
                    _ = len(v2)
                    delete(v3, fi13)
                    var v14 int8 = 42
                    var v15 int8 = int8(v6)
                    if false {
                        var v16 float32 = -2.25
                        v16 -= -2.25
                        if false {
                            if len(v2) > 0 {
                                v2[0] = (("" + "") + (("hello" + "a") + "a"))
                            }
                        } else {
                            var v17 int16 = 0
                            switch "" {
                            case "a":
                                fmt.Println(v15, v17, v4)
                            case "\n\t":
                                var v18 float64 = 0.0
                                _ = v18
                            case "foo bar":
                                var v19 float64 = (1e10 / 1.0)
                                _ = v19
                            default:
                                fmt.Println(fi7)
                            }
                            var v20 float64 = ((3.14159 + 3.14159) * 0.1)
                            _ = v17
                            _ = v20
                        }
                        _ = v16
                    } else {
                        v21 := GenS3{F0: 65535}
                        v5 += v5
                        var v22 float64 = (float64(-2.25) + 1.5)
                        _ = v21
                        _ = v22
                    }
                    v23 := v3[fi13]
                    if (2147483647 < 65535) {
                        v1[0] = 2
                        if ((v5 != (v5 & 306)) || (-2147483648 != -2147483648)) {
                            v23 = fi7
                            if len(v2) > 0 {
                                v2[0] = ("foo bar" + "")
                            }
                            v24 := GenS3{F0: 103}
                            _ = v24
                        }
                        fmt.Println(fi13)
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v23
                }
            }
            _ = fi7
        }
        v5 = 4294967295
    }
    if (false && (9223372036854775807 > -994)) {
        v25 := GenS1{F0: true, F1: 127}
        _ = v25
    }
    v3[437] = 42
    v4 += 132
    v6 = v6
    fmt.Println(v6, v5)
    switch 127 {
    case 943:
        v5 += (v5 | (v5 | v5))
    case -181:
        var v26 uint8 = (v4 | v4)
        _ = v26
    default:
        for fi27 := 0; fi27 < 16; fi27++ {
            v28 := GenS4{F0: 177}
            _ = fi27
            _ = v28
        }
    }
    if (!false) {
        var v29 int16 = -32768
        v30 := GenS5{F0: 1, F1: 65535, F2: 375}
        _ = v29
        _ = v30
    } else {
        v6 = v6
        for fi31 := 0; fi31 < 7; fi31++ {
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 int8 = 100
                _ = fi32
                _ = v33
            }
            v5 = uint(v6)
            v2 = append(v2, "")
            _ = fi31
        }
        _ = len(v3)
    }
    v4 -= v4
    for fi34 := 0; fi34 < 6; fi34++ {
        v4 -= v4
        v35 := GenS6{F0: 42, F1: 36, F2: -2147483648}
        var v36 int64 = 0
        _ = fi34
        _ = v35
        _ = v36
    }
    delete(v1, 0)
    v37 := GenS4{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
