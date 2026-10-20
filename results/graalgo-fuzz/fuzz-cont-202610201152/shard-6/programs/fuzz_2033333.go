package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    var v1 int8 = -1
    v2 := []string{"hello"}
    v3 := make(map[string]uint32)
    v4 := []uint32{}
    if (!(504 == 1)) {
        for fi5 := 0; fi5 < 19; fi5++ {
            v1 -= (v1 / ((v1) | 1))
            var v6 uint32 = 42
            v6 += 2147483647
            _ = fi5
            _ = v6
        }
    } else {
        v1 = (v1 << ((255) & 63))
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: false, F1: 42}
        _ = fi7
        _ = v8
    }
    switch "a" {
    case "":
        switch 127 {
        case 2147483647:
            v4 = append(v4, 667)
        }
    }
    var v9 uint32 = 4294967295
    if (false && ((0.1 != (-2.25 - 1e10)) || ((0.1 * (3.14159 - 1.5)) > 0.1))) {
        if false {
            v10 := GenS2{F0: false, F1: (v9 ^ v9)}
            _ = v10
        }
    }
    if true {
        switch 42 {
        case 1:
            var v11 uint64 = 4294967295
            _ = v11
        default:
            if (!false) {
                if true {
                    for fi12 := 0; fi12 < 4; fi12++ {
                        v9 -= (v9 % ((v9) | 1))
                        _ = fi12
                    }
                    for fi13 := 0; fi13 < 18; fi13++ {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v1 = (-(v1))
                            _ = fi14
                        }
                        switch fi13 {
                        case 1:
                            for fi15 := 0; fi15 < 20; fi15++ {
                                v16 := GenS3{F0: (^(v9)), F1: 2147483647}
                                var v17 int64 = 1
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        case 537:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case 9223372036854775807:
                            for fi18 := 0; fi18 < 7; fi18++ {
                                switch "a" {
                                case "a":
                                    v9 = v9
                                case "hello":
                                    v1 = 127
                                case "":
                                    var v19 int64 = 641
                                    _ = v19
                                }
                                _ = fi18
                            }
                        }
                        switch "a" {
                        case "\n\t":
                            for fi20 := 0; fi20 < 8; fi20++ {
                                fmt.Println(fi13, v1, v1)
                                _ = fi20
                            }
                        case "":
                            var v21 uint16 = 479
                            _ = v21
                        default:
                            var v22 int32 = 1
                            _ = v22
                        }
                        _ = fi13
                    }
                    for fi23 := 0; fi23 < 16; fi23++ {
                        if true {
                            for fi24 := 0; fi24 < 16; fi24++ {
                                if false {
                                    v25 := GenS3{F0: 65535, F1: 2147483647}
                                    var v26 float64 = 1e10
                                    if ((290 < 720) && (-2147483648 != 42)) {
                                        var v27 uint16 = 42
                                        var v28 int = (fi23 & fi23)
                                        _ = v27
                                        _ = v28
                                    } else {
                                        var v29 int32 = -2147483648
                                        _ = v29
                                    }
                                    _ = v25
                                    _ = v26
                                }
                                _ = fi24
                            }
                        } else {
                            v30 := v3[("" + ("" + "\n\t"))]
                            v31 := GenS3{F0: v9, F1: 0}
                            _ = v30
                            _ = v31
                        }
                        if (!true) {
                            var v32 uint64 = 0
                            _ = v32
                        } else {
                            var v33 uint32 = (v9 << ((2) & 63))
                            for fi34 := 0; fi34 < 13; fi34++ {
                                v9 -= v9
                                var v35 uint32 = 423
                                _ = fi34
                                _ = v35
                            }
                            _ = v33
                        }
                        fmt.Println(v9, v9)
                        _ = fi23
                    }
                } else {
                    fmt.Println(v9, v1)
                }
            }
        }
    } else {
        v9 -= v9
        if len(v4) > 0 {
            v4[0] = 2147483648
        }
    }
    v9 += v9
    v1 = -46
    var v36 uint64 = 1
    var v37 uint32 = (v9 >> ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v9
}
