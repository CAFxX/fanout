package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint32 }

func main() {
    var v1 int64 = -1
    v2 := make(map[string]string)
    var v3 uint64 = 2
    v4 := make(map[int]uint32)
    v5 := []uint32{1, 0, 552, 656}
    var v6 uint8 = 80
    v7 := GenS1{F0: "\n\t", F1: false, F2: 0}
    var v8 float32 = -2.25
    v1 = -1
    _ = len(v5)
    v6 = v6
    fmt.Println(v3)
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 int32 = 2147483647
        v1 = (v1 << ((220) & 63))
        _ = fi9
        _ = v10
    }
    switch -1 {
    case 128:
        var v11 float32 = float32(3.14159)
        _ = v11
    case -10:
        if true {
            v12 := GenS2{F0: 0}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 13; fi13++ {
                v14 := GenS1{F0: "foo bar", F1: true, F2: 787}
                _ = fi13
                _ = v14
            }
            if true {
                if (32767 == 929) {
                    var v15 uint8 = (^(v6))
                    var v16 float64 = 3.14159
                    fmt.Println(v15, v15, v16)
                    _ = v15
                    _ = v16
                }
                switch "\n\t" {
                case "\n\t":
                    var v17 bool = true
                    _ = v17
                case "a":
                    var v18 float32 = v8
                    _ = v18
                }
                fmt.Println(v1, v6, v8)
            } else {
                for fi19 := 0; fi19 < 7; fi19++ {
                    if false {
                        if (!true) {
                            if (!(42 >= (v1 | (v1 - v1)))) {
                                var v20 int64 = (v1 | (v1 & v1))
                                for fi21 := 0; fi21 < 7; fi21++ {
                                    var v22 float32 = v8
                                    v1 = v1
                                    _ = fi21
                                    _ = v22
                                }
                                v23 := GenS1{F0: "foo bar", F1: false, F2: 869}
                                _ = v20
                                _ = v23
                            }
                        } else {
                            if true {
                                var v24 string = "a"
                                v8 = v8
                                v5 = append(v5, 42)
                                _ = v24
                            }
                            v25 := GenS3{F0: (v6 + v6), F1: int64(v3), F2: 811}
                            fmt.Println(v6, v1, v8)
                            _ = v25
                        }
                    } else {
                        switch ("hello" + "foo bar") {
                        case "hello":
                            fmt.Println(v6, v8)
                        }
                    }
                    for fi26 := 0; fi26 < 4; fi26++ {
                        v3 = v3
                        if true {
                            v5 = append(v5, 2)
                            v27 := GenS2{F0: 42}
                            _ = v27
                        }
                        var v28 float32 = ((0.0 / v8) / v8)
                        fmt.Println(v1, fi26, v1)
                        v3 += 2
                        _ = fi26
                        _ = v28
                    }
                    _ = fi19
                }
                for fi29 := 0; fi29 < 18; fi29++ {
                    var v30 uint = 0
                    if (true || true) {
                        if (!false) {
                            if false {
                                switch "foo bar" {
                                case "\n\t":
                                    v30 += (v30 / ((v30) | 1))
                                default:
                                    v3 -= (v3 | v3)
                                }
                                v1 = (v1 | (v1 & v1))
                                for fi31 := 0; fi31 < 5; fi31++ {
                                    var v32 int = fi29
                                    _ = fi31
                                    _ = v32
                                }
                            } else {
                                for fi33 := 0; fi33 < 20; fi33++ {
                                    v34 := GenS1{F0: ("foo bar" + "\n\t"), F1: ((-331 > fi29) || (!false)), F2: (v30 ^ v30)}
                                    _ = len(v2)
                                    _ = fi33
                                    _ = v34
                                }
                                v5 = append(v5, 1)
                            }
                        } else {
                            fmt.Println(v1, v1)
                        }
                    }
                    for fi35 := 0; fi35 < 10; fi35++ {
                        fmt.Println(fi29, fi29)
                        v1 += (v1 ^ (v1 | -379))
                        var v36 int8 = -1
                        var v37 int8 = (^(v36))
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = fi29
                    _ = v30
                }
            }
            _ = len(v2)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
