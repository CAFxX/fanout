package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []uint32{461}
    var v2 uint = 65535
    var v3 int8 = 42
    var v4 int = 321
    var v5 uint16 = 2
    v6 := make(map[int]uint32)
    fmt.Println(v2, v3, v3)
    v7 := GenS1{F0: 18446744073709551615, F1: 1, F2: -625}
    if (!false) {
        _ = len(v1)
        v5 -= 255
    } else {
        delete(v6, -370)
    }
    v8 := GenS1{F0: 1, F1: 42, F2: 1}
    for fi9 := 0; fi9 < 7; fi9++ {
        switch (("foo bar" + ("hello" + "foo bar")) + (("foo bar" + (("\n\t" + "foo bar") + ("foo bar" + "a"))) + ("" + ("\n\t" + "")))) {
        case "foo bar":
            v2 = (v2 - v2)
        }
        _ = fi9
    }
    v5 = (v5 & (v5 % ((1) | 1)))
    if false {
        for fi10 := 0; fi10 < 5; fi10++ {
            switch (("foo bar" + (((("hello" + "") + (((("\n\t" + "hello") + "foo bar") + "hello") + "hello")) + "foo bar") + "foo bar")) + "foo bar") {
            case "foo bar":
                for fi11 := 0; fi11 < 14; fi11++ {
                    v12 := GenS2{F0: (v2 * v2), F1: 127}
                    v4 = (v4 << ((v2) & 63))
                    _ = len(v6)
                    _ = fi11
                    _ = v12
                }
            case "":
                var v13 string = "\n\t"
                _ = v13
            case "hello":
                v14 := v6[4294967295]
                _ = v14
            default:
                v2 = v2
            }
            if len(v1) > 0 {
                v1[0] = 646
            }
            _ = fi10
        }
        if true {
            v15 := GenS3{F0: ("foo bar" + "hello"), F1: v3}
            var v16 float32 = -2.25
            var v17 bool = (2147483648 < 947)
            _ = v15
            _ = v16
            _ = v17
        }
    } else {
        v4 += -251
        if false {
            fmt.Println(v4)
            _ = len(v6)
        } else {
            v18 := GenS2{F0: v2, F1: (v3 - (v3 - v3))}
            _ = v18
        }
        fmt.Println(v4)
    }
    var v19 int32 = -924
    var v20 int16 = 32767
    _ = len(v6)
    v19 += (^(v19))
    if (!false) {
        fmt.Println(v3)
    } else {
        fmt.Println(v3, v19)
        var v21 uint64 = 0
        for fi22 := 0; fi22 < 6; fi22++ {
            v21 = (^(v21))
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 uint = (^(v2))
                fmt.Println(fi23, v3)
                if ((3.14159 + 1e10) <= float64((1e10 / 1e10))) {
                    var v25 int8 = 99
                    v20 = v20
                    _ = v25
                } else {
                    if false {
                        var v26 float32 = 0.0
                        var v27 float32 = (v26 / ((v26 - (1.5 * (1.5 / -2.25))) / v26))
                        switch fi23 {
                        case 746:
                            if false {
                                switch (fi22 - -1) {
                                case 981:
                                    v28 := GenS1{F0: v21, F1: 998, F2: 0}
                                    _ = v28
                                case 42:
                                    switch v4 {
                                    case 2147483647:
                                        var v29 float32 = float32(-2.25)
                                        fmt.Println(v3, v4)
                                        if false {
                                            fmt.Println(v5)
                                            var v30 int = fi22
                                            _ = v30
                                        }
                                        _ = v29
                                    default:
                                        if false {
                                            v31 := GenS1{F0: 260, F1: 9, F2: 42}
                                            if false {
                                                if (95 > 124) {
                                                    var v32 int = (fi22 | 1)
                                                    if true {
                                                        v33 := GenS4{F0: (v5 + (v5 >> ((18446744073709551615) & 63)))}
                                                        _ = v33
                                                    } else {
                                                        v6[fi22] = 4294967295
                                                    }
                                                    fmt.Println(v19, v2, v4)
                                                    fmt.Println(v20)
                                                    if len(v1) > 0 {
                                                        v1[0] = 147
                                                    }
                                                    _ = v32
                                                }
                                                v19 = (v19 / ((-2147483648) | 1))
                                            }
                                            _ = v31
                                        }
                                    }
                                case 32767:
                                    v24 = (v2 << (((v2 << ((v24) & 63))) & 63))
                                }
                            } else {
                                v19 -= v19
                                var v34 int = fi23
                                _ = v34
                            }
                        default:
                            v5 -= v5
                        }
                        _ = v26
                        _ = v27
                    }
                }
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        _ = v21
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        for fi36 := 0; fi36 < 9; fi36++ {
            v20 = v20
            _ = fi36
        }
        _ = fi35
    }
    v20 = v20
    v4 += (^(v4))
    for fi37 := 0; fi37 < 1; fi37++ {
        v38 := GenS2{F0: (v2 ^ v2), F1: v3}
        _ = fi37
        _ = v38
    }
    var v39 uint32 = 4294967295
    _ = len(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
