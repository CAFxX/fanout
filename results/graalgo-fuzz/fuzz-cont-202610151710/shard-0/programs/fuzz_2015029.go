package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -9223372036854775808
    v3 := []uint32{}
    v4 := make(map[uint32]uint32)
    var v5 float32 = (3.14159 + (float32(-2.25) - (0.1 / 1.0)))
    v3 = append(v3, 2)
    v5 -= v5
    v6 := GenS1{F0: 76, F1: 0}
    v4[643] = 2147483647
    v2 = v2
    v7 := GenS2{F0: 255, F1: 517, F2: -205}
    _ = len(v3)
    var v8 int16 = -32768
    var v9 int16 = 373
    v8 -= -32768
    fmt.Println(v8, v8, v2)
    switch "hello" {
    case "a":
        if true {
            switch -9223372036854775808 {
            case 42:
                switch 9223372036854775807 {
                case 4294967295:
                    fmt.Println(v2)
                case 32767:
                    var v10 string = "foo bar"
                    _ = v10
                case -341:
                    if false {
                        v11 := GenS1{F0: 183, F1: 4294967295}
                        if false {
                            if true {
                                v9 += 32767
                            } else {
                                fmt.Println(v9)
                                if true {
                                    fmt.Println(v8, v9, v8)
                                    v12 := GenS1{F0: 132, F1: 288}
                                    v2 += (v2 % ((9223372036854775807) | 1))
                                    _ = v12
                                } else {
                                    delete(v1, "foo bar")
                                }
                            }
                            for fi13 := 0; fi13 < 9; fi13++ {
                                for fi14 := 0; fi14 < 17; fi14++ {
                                    v5 -= v5
                                    if (!false) {
                                        v15 := GenS3{F0: 255, F1: 42, F2: fi13}
                                        v16 := GenS1{F0: 255, F1: 562}
                                        var v17 float64 = (float64(v5) - 1e10)
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                    } else {
                                        var v18 uint16 = 42
                                        _ = v18
                                    }
                                    if (!true) {
                                        fmt.Println(v8, v9)
                                        v8 = v9
                                    }
                                    _ = fi14
                                }
                                v19 := GenS4{F0: 0, F1: 0}
                                _ = fi13
                                _ = v19
                            }
                        } else {
                            v20 := v4[657]
                            _ = v20
                        }
                        _ = v11
                    } else {
                        v21 := GenS3{F0: 255, F1: 255, F2: 42}
                        for fi22 := 0; fi22 < 13; fi22++ {
                            switch (fi22 * int(v2)) {
                            case 1:
                                v5 = 1.5
                            default:
                                v23 := GenS1{F0: 128, F1: 2}
                                _ = v23
                            }
                            _ = fi22
                        }
                        var v24 int16 = v9
                        _ = v21
                        _ = v24
                    }
                }
            case 32767:
                _ = len(v4)
            case 127:
                v2 = -622
            }
        }
    case "":
        v2 = v2
    case "foo bar":
        v25 := GenS5{F0: 163, F1: 865}
        _ = v25
    default:
        v26 := v1["\n\t"]
        _ = v26
    }
    switch -2147483648 {
    case 128:
        for fi27 := 0; fi27 < 17; fi27++ {
            v3 = append(v3, 591)
            v28 := GenS1{F0: 1, F1: 894}
            if true {
                for fi29 := 0; fi29 < 13; fi29++ {
                    switch ("a" + ("hello" + ("" + "foo bar"))) {
                    case "hello":
                        var v30 int64 = 1
                        _ = v30
                    default:
                        var v31 uint64 = 818
                        _ = v31
                    }
                    _ = fi29
                }
            }
            _ = fi27
            _ = v28
        }
    case -1:
        v32 := GenS2{F0: 858, F1: 2, F2: -32768}
        _ = v32
    }
    v9 = (^(v9))
    v2 += 42
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 uint16 = 90
        var v35 string = "a"
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 float64 = float64(1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
