package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
func genFn0() int {
    var v5 int16 = 42
    var v6 int = -521
    v7 := []int{(v6 | 0), (v6 >> ((334) & 63)), v6}
    fmt.Println(v5, v6)
    fmt.Println(v6)
    _ = v5
    _ = v6
    _ = v7
    return (v6 - (v6 ^ v6))
}
func genFn1(p0 int64, p1 bool, p2 string) int64 {
    p1 = (!p1)
    p2 = p2
    var v8 int = genFn0()
    p2 = (("foo bar" + "a") + "hello")
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p0
}

func main() {
    var v1 bool = (0 >= 42)
    v2 := []string{"a", "hello", ""}
    var v3 uint8 = 255
    var v4 float64 = ((1.5 / 1.0) / 3.14159)
    v4 += v4
    v4 += ((v4 / (v4 * 1.5)) + (v4 * 1e10))
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v4 = v4
        if ((1000 < 0) && v1) {
            v4 = (v4 / 1e10)
        } else {
            v4 = (v4 / -2.25)
            v9 := GenS1{F0: 2}
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS2{F0: 0}
                var v12 uint8 = 127
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        }
    }
    switch "foo bar" {
    case "foo bar":
        _ = len(v2)
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        if v1 {
            switch fi13 {
            case -606:
                v4 = v4
            case 32767:
                v3 = v3
            case 452:
                for fi14 := 0; fi14 < 10; fi14++ {
                    v3 = (v3 >> ((187) & 63))
                    if v1 {
                        var v15 uint = 4294967295
                        v16 := GenS2{F0: -1}
                        _ = v15
                        _ = v16
                    }
                    _ = fi14
                }
            }
            if (0.0 == 1.5) {
                if v1 {
                    v2 = append(v2, (("\n\t" + "\n\t") + (("hello" + "a") + "foo bar")))
                    v3 = v3
                } else {
                    if (v4 == v4) {
                        v17 := GenS3{F0: (v3 * v3), F1: 233}
                        v1 = v1
                        switch "foo bar" {
                        case "\n\t":
                            if v1 {
                                var v18 int32 = -1
                                _ = v18
                            } else {
                                _ = len(v2)
                                v4 = 1e10
                                var v19 float32 = -2.25
                                _ = v19
                            }
                        case "foo bar":
                            v1 = (2 == 446)
                        default:
                            for fi20 := 0; fi20 < 9; fi20++ {
                                fmt.Println(v3, v4)
                                if v1 {
                                    fmt.Println(v1)
                                    var v21 uint32 = 708
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                    var v22 int32 = 1
                                    var v23 float32 = -2.25
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                _ = fi20
                            }
                        }
                        _ = v17
                    }
                    fmt.Println(fi13, v3, v4)
                    fmt.Println(v4, v1, v1)
                }
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 uint64 = 100
                    var v26 int32 = -2147483648
                    v2 = append(v2, ("" + "a"))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                switch "hello" {
                case "hello":
                    v27 := GenS4{F0: 340}
                    _ = v27
                case "a":
                    for fi28 := 0; fi28 < 18; fi28++ {
                        var v29 int = 9223372036854775807
                        v30 := GenS4{F0: 4294967295}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case "foo bar":
                    var v31 uint64 = 42
                    _ = v31
                }
            }
        } else {
            var v32 float32 = -2.25
            _ = v32
        }
        if (v1 && (v4 > v4)) {
            v33 := GenS3{F0: (v3 - v3), F1: 255}
            _ = v33
        } else {
            v1 = (18446744073709551615 < 969)
            switch "" {
            case "a":
                _ = len(v2)
            default:
                switch (fi13 & fi13) {
                case -9223372036854775808:
                    switch (("" + "\n\t") + "hello") {
                    case "":
                        v2 = append(v2, ("\n\t" + ("a" + "")))
                    default:
                        for fi34 := 0; fi34 < 6; fi34++ {
                            v3 -= v3
                            v35 := GenS5{F0: (float64((1.5 * 1e10)) >= (-2.25 * v4))}
                            _ = fi34
                            _ = v35
                        }
                    }
                case -1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v36 float64 = v4
                    _ = v36
                }
            }
        }
        _ = fi13
    }
    if true {
        var v37 int32 = 1
        var v38 int16 = 950
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v37
        _ = v38
    } else {
        switch 4294967295 {
        case -9223372036854775808:
            switch 489 {
            case 315:
                fmt.Println(v4)
            }
        case -359:
            switch genFn0() {
            case 128:
                v1 = v1
            default:
                var v39 uint = 255
                _ = v39
            }
        default:
            fmt.Println(v4)
        }
    }
    v3 -= v3
    v1 = v1
    v3 += 23
    v40 := GenS5{F0: (v1 && (((false || false) && v1) || (true && (65535 > 42))))}
    if v1 {
        v4 = v4
        v3 += (v3 >> ((255) & 63))
    } else {
        if v1 {
            var v41 uint16 = 42
            _ = v41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}
