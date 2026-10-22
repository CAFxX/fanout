package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 uint8 }
func genFn0() int8 {
    var v5 int64 = 42
    var v6 uint16 = 0
    v5 = (^(v5))
    _ = v5
    _ = v6
    return 100
}

func main() {
    v1 := []int{}
    var v2 int16 = -942
    v3 := []int{-9223372036854775808, 32767}
    v4 := []string{"", "hello", "\n\t"}
    var v7 int = 42
    switch "a" {
    case "a":
        for fi8 := 0; fi8 < 15; fi8++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 uint = 4294967295
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    default:
        var v11 int16 = v2
        _ = v11
    }
    v2 += (v2 / ((0) | 1))
    v12 := GenS1{F0: 0, F1: -9223372036854775808, F2: 42}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v7 = (v7 ^ 830)
    v1 = append(v1, (v7 | 9223372036854775807))
    v1 = append(v1, 1)
    fmt.Println(v7, v2)
    if true {
        for fi13 := 0; fi13 < 9; fi13++ {
            if (65535 != 2) {
                v14 := GenS1{F0: -13, F1: 9223372036854775807, F2: 287}
                _ = v14
            } else {
                fmt.Println(fi13, v2, fi13)
                if ("hello" <= "") {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        v7 -= (v7 ^ (fi13 & 848))
                        var v16 float64 = 3.14159
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi15
                        _ = v16
                    }
                } else {
                    for fi17 := 0; fi17 < 11; fi17++ {
                        v7 -= int(v2)
                        _ = fi17
                    }
                    var v18 int32 = -111
                    if len(v3) > 0 {
                        v3[0] = (^(fi13))
                    }
                    _ = v18
                }
            }
            if len(v1) > 0 {
                v1[0] = fi13
            }
            v7 += fi13
            _ = fi13
        }
        if false {
            var v19 float32 = float32(0.1)
            _ = v19
        }
        var v20 int = v7
        _ = v20
    } else {
        var v21 uint64 = 443
        _ = v21
    }
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi22 := 0; fi22 < 15; fi22++ {
            v23 := GenS2{F0: 128, F1: ("a" + "hello")}
            _ = fi22
            _ = v23
        }
        v24 := GenS3{F0: 127, F1: 42, F2: 127}
        _ = v24
    } else {
        if true {
            if false {
                fmt.Println(v7, v2)
                switch ("a" + "") {
                case "foo bar":
                    var v25 int = v7
                    _ = v25
                case "hello":
                    for fi26 := 0; fi26 < 20; fi26++ {
                        v27 := GenS3{F0: 0, F1: 0, F2: 1}
                        _ = fi26
                        _ = v27
                    }
                }
                v4 = append(v4, ("" + ""))
            } else {
                var v28 int = 128
                if (4294967295 == 42) {
                    for fi29 := 0; fi29 < 10; fi29++ {
                        for fi30 := 0; fi30 < 7; fi30++ {
                            v3 = append(v3, v28)
                            _ = fi30
                        }
                        for fi31 := 0; fi31 < 14; fi31++ {
                            v32 := GenS4{F0: ((^(v28)) == v7), F1: 104}
                            v33 := GenS2{F0: 1, F1: ""}
                            if true {
                                var v34 uint8 = 1
                                _ = v34
                            }
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        _ = fi29
                    }
                    for fi35 := 0; fi35 < 1; fi35++ {
                        for fi36 := 0; fi36 < 20; fi36++ {
                            var v37 float64 = 0.0
                            v3 = append(v3, v7)
                            if ((v28 & fi35) < -943) {
                                v38 := GenS5{F0: 0}
                                var v39 uint64 = 0
                                if (255 > 65535) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v3 = append(v3, v7)
                                    v2 = v2
                                    var v40 uint8 = 80
                                    v3 = append(v3, fi35)
                                    _ = v40
                                } else {
                                    v1 = append(v1, -833)
                                    var v41 int = (-(v28))
                                    v42 := GenS6{F0: 815, F1: 65535}
                                    _ = v41
                                    _ = v42
                                }
                                _ = v38
                                _ = v39
                            } else {
                                if (float32(v37) != (1.5 - (-2.25 - -2.25))) {
                                    v7 -= (^(v7))
                                    switch 9223372036854775807 {
                                    case -2147483648:
                                        v7 = v28
                                    case 679:
                                        fmt.Println(fi35)
                                    }
                                    v43 := GenS4{F0: ("hello" == "a"), F1: 1}
                                    _ = v43
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 1
                                    }
                                }
                            }
                            switch -1 {
                            case 128:
                                var v44 uint64 = 628
                                _ = v44
                            case 9223372036854775807:
                                var v45 uint64 = 1
                                _ = v45
                            case 353:
                                _ = len(v4)
                            default:
                                if false {
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = 0
                                        }
                                    } else {
                                        var v46 uint8 = 42
                                        _ = v46
                                    }
                                    v28 += v28
                                }
                            }
                            v47 := GenS7{F0: 1}
                            var v48 int64 = 42
                            _ = fi36
                            _ = v37
                            _ = v47
                            _ = v48
                        }
                        _ = fi35
                    }
                }
                var v49 uint64 = 357
                _ = v28
                _ = v49
            }
            for fi50 := 0; fi50 < 15; fi50++ {
                v7 += 0
                _ = fi50
            }
        }
        switch "" {
        case "hello":
            fmt.Println(v2, v7)
        case "foo bar":
            var v51 bool = (!true)
            _ = v51
        }
    }
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
