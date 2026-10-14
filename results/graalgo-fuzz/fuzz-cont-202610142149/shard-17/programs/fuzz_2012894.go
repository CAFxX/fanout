package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int64; F1 uint; F2 uint8 }
type GenS4 struct { F0 int16; F1 string; F2 bool }
type GenS5 struct { F0 int32; F1 bool; F2 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{42, -720}
    v3 := []string{("" + ("a" + ((("" + (("\n\t" + "hello") + ((("a" + ("\n\t" + "\n\t")) + "\n\t") + "a"))) + ("foo bar" + ((("hello" + "\n\t") + ("\n\t" + ("foo bar" + ""))) + "hello"))) + ""))), "a"}
    v4 := make(map[uint32]string)
    var v5 float32 = (((float32(3.14159) - 3.14159) - ((3.14159 - 1.5) - (0.0 / -2.25))) * 1e10)
    if (!false) {
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := GenS1{F0: fi6, F1: "\n\t", F2: 0}
            var v8 uint = 138
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 float32 = (1e10 + v5)
                _ = fi9
                _ = v10
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        v5 = ((float32(1.5) - 0.1) - v5)
        if true {
            if true {
                fmt.Println(v5)
                for fi11 := 0; fi11 < 20; fi11++ {
                    v5 = 1.5
                    _ = fi11
                }
                v5 -= v5
            }
            for fi12 := 0; fi12 < 7; fi12++ {
                _ = len(v1)
                v5 = v5
                _ = fi12
            }
        }
    }
    v5 = (float32(-2.25) / ((((v5 / (1.5 - v5)) + (v5 + v5)) / 1e10) + v5))
    v5 = (float32(((0.0 + -2.25) - -2.25)) / 1.0)
    v5 += v5
    fmt.Println(v5)
    switch ("foo bar" + ("\n\t" + "\n\t")) {
    case "\n\t":
        _ = len(v2)
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v5)
    case "":
        for fi13 := 0; fi13 < 18; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                if true {
                    if (788 > 32767) {
                        if ((fi13 != (fi14 << ((585) & 63))) && (41 != 0)) {
                            var v15 int64 = 9223372036854775807
                            switch fi13 {
                            case 29:
                                if (fi13 >= fi14) {
                                    v5 += float32(1.5)
                                }
                            case 0:
                                var v16 int16 = -894
                                _ = v16
                            case -9223372036854775808:
                                v17 := GenS2{F0: -1, F1: 4}
                                _ = v17
                            }
                            if (false || ("hello" == (("\n\t" + ((("" + (("foo bar" + "\n\t") + "")) + "hello") + "a")) + (("hello" + "a") + "\n\t")))) {
                                v15 = (v15 << ((4294967295) & 63))
                                fmt.Println(fi14, fi13, fi14)
                                v18 := GenS3{F0: -438, F1: 0, F2: 127}
                                fmt.Println(fi13, fi13, v5)
                                switch (("a" + ("hello" + "\n\t")) + "a") {
                                case "foo bar":
                                    _ = len(v1)
                                default:
                                    fmt.Println(fi13, v15)
                                }
                                _ = v18
                            } else {
                                if false {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    var v19 uint8 = 0
                                    switch (fi14 >> ((255) & 63)) {
                                    case 744:
                                        var v20 int8 = -128
                                        _ = v20
                                    case 127:
                                        if (1.5 != v5) {
                                            var v21 uint16 = 42
                                            _ = len(v1)
                                            _ = v21
                                        } else {
                                            fmt.Println(fi13, fi14, v5)
                                            v19 -= 1
                                        }
                                    default:
                                        var v22 bool = true
                                        v19 = v19
                                        var v23 float64 = (((0.0 * float64(v5)) / -2.25) + -2.25)
                                        _ = v22
                                        _ = v23
                                    }
                                    if (0.1 != (3.14159 / float64((1.5 * 3.14159)))) {
                                        if (-128 == 0) {
                                            v15 = (v15 & (v15 * (v15 | 603)))
                                            var v24 float64 = 0.0
                                            _ = v24
                                        }
                                    } else {
                                        var v25 uint16 = 456
                                        _ = v25
                                    }
                                    _ = v19
                                }
                            }
                            var v26 uint16 = 0
                            _ = v15
                            _ = v26
                        }
                    } else {
                        v27 := GenS4{F0: 612, F1: ("foo bar" + ""), F2: true}
                        var v28 float64 = (1.5 * ((1.5 * 1e10) - 1e10))
                        v5 = v5
                        v2 = append(v2, 684)
                        v28 = v28
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(fi13)
                } else {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                }
                _ = fi14
            }
            _ = fi13
        }
    default:
        var v29 uint32 = 65535
        _ = v29
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        v31 := GenS5{F0: 0, F1: (true || (0.0 > v5)), F2: (!false)}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        fmt.Println(v5)
        v33 := GenS5{F0: -2147483648, F1: true, F2: true}
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 float64 = 1.5
            switch "foo bar" {
            case "foo bar":
                if (!true) {
                    var v36 int64 = -175
                    v35 = v35
                    v37 := GenS2{F0: 0, F1: 0}
                    _ = v36
                    _ = v37
                } else {
                    var v38 int16 = 0
                    _ = v38
                }
            case "hello":
                if ("\n\t" <= ("\n\t" + ("hello" + ("\n\t" + "\n\t")))) {
                    v39 := GenS6{F0: 32768}
                    var v40 int = fi32
                    var v41 int64 = -1
                    var v42 int = (fi32 ^ fi32)
                    var v43 int64 = (^(v41))
                    _ = v39
                    _ = v40
                    _ = v41
                    _ = v42
                    _ = v43
                } else {
                    switch "hello" {
                    case "hello":
                        if (false && (2147483647 < 4294967295)) {
                            var v44 int32 = -2147483648
                            _ = v44
                        } else {
                            v45 := GenS7{F0: 2}
                            _ = v45
                        }
                    default:
                        var v46 uint8 = 128
                        _ = v46
                    }
                    var v47 uint32 = 2
                    var v48 int = int(v47)
                    v5 = v5
                    _ = v47
                    _ = v48
                }
            }
            v4[65535] = "foo bar"
            _ = fi34
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
