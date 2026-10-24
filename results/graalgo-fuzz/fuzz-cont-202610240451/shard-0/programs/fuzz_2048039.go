package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 bool }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 uint16 }

func main() {
    var v1 bool = ((false || false) || (false && true))
    var v2 int32 = 42
    v3 := []uint32{0, 732, 1, 2147483647}
    var v4 string = "hello"
    var v5 bool = (9223372036854775807 < 1)
    if len(v3) > 0 {
        v3[0] = 2
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        if ((!((!true) && v1)) || (((!(-2 <= 100)) || (42 > -1)) || true)) {
            if false {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                if true {
                    v7 := GenS1{F0: 65535, F1: v5, F2: v5}
                    for fi8 := 0; fi8 < 20; fi8++ {
                        v4 = (v4 + v4)
                        _ = fi8
                    }
                    _ = v7
                }
                switch fi6 {
                case 32768:
                    fmt.Println(v5, v5)
                case 934:
                    v5 = ((!false) && (v4 >= v4))
                default:
                    v9 := GenS1{F0: 747, F1: (3.14159 > 1e10), F2: (v1 || true)}
                    _ = v9
                }
            }
        }
        _ = fi6
    }
    v3 = append(v3, 42)
    fmt.Println(v5, v4, v4)
    v4 = "\n\t"
    v1 = (-32768 != -1)
    v10 := GenS2{F0: 771, F1: -620}
    if len(v3) > 0 {
        v3[0] = 482
    }
    var v11 int8 = 36
    for fi12 := 0; fi12 < 12; fi12++ {
        v2 = 0
        _ = fi12
    }
    v4 = v4
    if v1 {
        if (v2 <= v2) {
            v11 = (v11 | 100)
            fmt.Println(v4, v1)
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                if v1 {
                    if len(v3) > 0 {
                        v3[0] = 65535
                    }
                    var v14 float32 = 1e10
                    v15 := GenS2{F0: 42, F1: 42}
                    _ = v14
                    _ = v15
                } else {
                    if (v5 || false) {
                        switch v4 {
                        case "a":
                            v1 = v5
                        case "hello":
                            for fi16 := 0; fi16 < 9; fi16++ {
                                if ((223 > 1) || v5) {
                                    v17 := GenS1{F0: 65535, F1: ((v2 == 0) || (v2 > -2147483648)), F2: false}
                                    var v18 int64 = 1
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                    v19 := GenS3{F0: 7}
                                    v11 -= (v11 + 118)
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                                v11 += v11
                                _ = fi16
                            }
                        case "\n\t":
                            v20 := GenS3{F0: (v11 & v11)}
                            _ = v20
                        default:
                            if (4294967295 < 65535) {
                                _ = len(v3)
                                switch 1 {
                                case 128:
                                    v3 = append(v3, 521)
                                case 9223372036854775807:
                                    if ((v2 != (v2 + -1)) && (int64(v2) < -570)) {
                                        v4 = ""
                                        for fi21 := 0; fi21 < 8; fi21++ {
                                            if (v1 && false) {
                                                v22 := GenS2{F0: 42, F1: 0}
                                                if v5 {
                                                    v2 += 247
                                                    var v23 int16 = 1
                                                    _ = v23
                                                }
                                                _ = v22
                                            } else {
                                                fmt.Println(fi13)
                                                if false {
                                                    v1 = (1 <= fi13)
                                                    v24 := GenS4{F0: 739}
                                                    _ = v24
                                                } else {
                                                    var v25 uint32 = 2147483647
                                                    _ = v25
                                                }
                                            }
                                            v2 -= (v2 >> ((4294967295) & 63))
                                            switch (v4 + (v4 + v4)) {
                                            case "":
                                                var v26 int = (fi13 << ((240) & 63))
                                                _ = v26
                                            case "a":
                                                if (v11 <= v11) {
                                                    if v1 {
                                                        var v27 float64 = ((3.14159 / 1.5) + 1.5)
                                                        _ = v27
                                                    } else {
                                                        var v28 float32 = 1e10
                                                        v4 = v4
                                                        _ = v28
                                                    }
                                                } else {
                                                    var v29 int = (fi13 >> ((18446744073709551615) & 63))
                                                    if v1 {
                                                        if (v1 || v1) {
                                                            _ = len(v3)
                                                            var v30 uint64 = 1
                                                            var v31 int16 = 32767
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = 65535
                                                        }
                                                    }
                                                    _ = v29
                                                }
                                                v2 -= (v2 * v2)
                                            }
                                            v3 = append(v3, 65535)
                                            v1 = (2 != 47)
                                            v5 = (v1 || (!(fi13 > (fi21 / (((fi21 | (fi13 >> ((2) & 63)))) | 1)))))
                                            _ = fi21
                                        }
                                        v32 := GenS5{F0: 97, F1: 1}
                                        _ = v32
                                    } else {
                                        switch fi13 {
                                        case 4294967295:
                                            var v33 uint64 = 0
                                            _ = v33
                                        }
                                    }
                                }
                            }
                        }
                        var v34 uint16 = 65535
                        _ = v34
                    }
                    _ = len(v3)
                }
                fmt.Println(v5, v11, v11)
                v1 = (!false)
                _ = fi13
            }
            v35 := GenS5{F0: v2, F1: 2}
            _ = v35
        }
        var v36 int16 = -45
        _ = v36
    }
    if true {
        fmt.Println(v2, v1, v1)
        var v37 int8 = (v11 ^ (v11 & (v11 ^ v11)))
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
