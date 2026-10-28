package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int8; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 bool; F2 uint8 }

func main() {
    v1 := []int{}
    var v2 uint64 = 4294967295
    var v3 string = ""
    var v4 int64 = -9223372036854775808
    v5 := []string{"a", "\n\t", v3, v3}
    v6 := make(map[string]string)
    if false {
        fmt.Println(v4, v2)
        if true {
            _ = len(v6)
        } else {
            v7 := GenS1{F0: 1000, F1: 84, F2: v2}
            fmt.Println(v3, v2, v2)
            _ = v7
        }
    }
    var v8 uint16 = 1
    switch 9223372036854775807 {
    case 32767:
        fmt.Println(v8)
    case 4294967295:
        fmt.Println(v8, v3)
    case 714:
        for fi9 := 0; fi9 < 18; fi9++ {
            if (false || (!false)) {
                fmt.Println(v2)
            } else {
                if false {
                    var v10 int8 = -1
                    delete(v6, v3)
                    _ = v10
                }
                v1 = append(v1, 195)
                var v11 int8 = 10
                _ = v11
            }
            _ = fi9
        }
    default:
        var v12 int64 = v4
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            switch (v3 + ("" + v3)) {
            case "foo bar":
                if false {
                    v2 = 4294967295
                }
            case "a":
                v15 := GenS1{F0: -32768, F1: 42, F2: v2}
                _ = v15
            case "\n\t":
                v16 := GenS1{F0: 0, F1: 42, F2: v2}
                _ = v16
            default:
                _ = len(v6)
            }
            fmt.Println(v3, v2, v2)
            _ = fi14
        }
        v4 = -561
        if false {
            var v17 uint8 = 128
            if (v8 == (^(v8))) {
                if true {
                    v3 = v3
                    _ = len(v6)
                    var v18 uint32 = 2
                    _ = v18
                } else {
                    delete(v6, ("a" + (v3 + ((v3 + "\n\t") + v3))))
                    v2 = (v2 * v2)
                }
                v3 = v3
                var v19 float32 = (1e10 / (1.5 - ((3.14159 - 0.1) - 3.14159)))
                _ = v19
            }
            _ = v17
        }
        _ = fi13
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v4 = v4
        _ = fi20
    }
    v4 = (v4 / ((v4) | 1))
    var v21 string = (("hello" + "") + ((v3 + "foo bar") + "a"))
    if true {
        v4 -= v4
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v2 = 96
    } else {
        var v22 int8 = 1
        v23 := GenS2{F0: 2, F1: v22, F2: int16(v4)}
        _ = v22
        _ = v23
    }
    v6[(v3 + (("" + "a") + "foo bar"))] = v3
    for fi24 := 0; fi24 < 12; fi24++ {
        if ((!true) || (true && true)) {
            var v25 uint32 = 2147483648
            _ = len(v5)
            _ = v25
        } else {
            if true {
                if false {
                    if true {
                        switch 32767 {
                        case 1:
                            if (42 < -58) {
                                for fi26 := 0; fi26 < 12; fi26++ {
                                    if (uint8(v8) != uint8(v8)) {
                                        v27 := GenS3{F0: (v8 * v8), F1: 32, F2: 101}
                                        v3 = (("" + (("a" + v21) + v21)) + ("foo bar" + v21))
                                        _ = v27
                                    }
                                    fmt.Println(v2)
                                    v28 := GenS4{F0: (v4 & (v4 % (((v4 << ((35) & 63))) | 1)))}
                                    v2 += v2
                                    _ = fi26
                                    _ = v28
                                }
                            } else {
                                var v29 uint = 0
                                var v30 int64 = (v4 << ((v29) & 63))
                                _ = v29
                                _ = v30
                            }
                        case 127:
                            var v31 float32 = (((1.5 - -2.25) - 1.5) * float32((-2.25 - (1.5 * (0.0 * -2.25)))))
                            _ = v31
                        case 128:
                            var v32 int64 = v4
                            _ = v32
                        }
                        v5 = append(v5, "foo bar")
                        v33 := GenS5{F0: -2147483648, F1: true, F2: 184}
                        _ = v33
                    } else {
                        v2 -= v2
                        v8 = 2
                    }
                }
                v8 = v8
            }
        }
        _ = fi24
    }
    if (128 >= 42) {
        var v34 uint64 = 1
        v8 = v8
        v6["\n\t"] = (v3 + ("\n\t" + "hello"))
        _ = v34
    } else {
        v4 -= (v4 % (((-(v4))) | 1))
        var v35 bool = (2 == 814)
        _ = v35
    }
    var v36 float32 = (3.14159 / 1.0)
    v36 = ((v36 / 3.14159) / (((float32((1e10 - 1e10)) * 3.14159) - (v36 - -2.25)) - (v36 / v36)))
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
