package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 bool }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 uint16; F1 uint; F2 int32 }
type GenS6 struct { F0 int64; F1 uint32; F2 int64 }
func genFn0(p0 int8, p1 int64, p2 uint64) int {
    var v6 uint = 652
    v7 := []int{42}
    var v8 bool = true
    if len(v7) > 0 {
        v7[0] = 917
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -587
}

func main() {
    v1 := []int{9223372036854775807, 1}
    var v2 uint16 = 42
    var v3 uint64 = 48
    var v4 int = -932
    var v5 int32 = 42
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint32 = 1
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 42
    v2 = (^(v2))
    var v12 string = "a"
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(fi13)
        _ = fi13
    }
    if false {
        var v14 uint8 = 204
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 = 2
        _ = v14
    } else {
        v3 += (v3 & v3)
    }
    v5 += (v5 - (v5 ^ (^(v5))))
    if (true && ((1 != 18446744073709551615) || false)) {
        if false {
            if (345 != 343) {
                v2 += (v2 + uint16(v11))
                for fi15 := 0; fi15 < 2; fi15++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v12 = "foo bar"
                    _ = fi15
                }
                v16 := GenS1{F0: 1, F1: false, F2: (true && false)}
                _ = v16
            } else {
                if (1 < 100) {
                    var v17 float32 = 3.14159
                    _ = v17
                }
                switch (v12 + (v12 + "foo bar")) {
                case "hello":
                    v18 := GenS2{F0: v11, F1: (0.0 <= float64(-2.25))}
                    _ = v18
                case "":
                    switch v12 {
                    case "":
                        var v19 uint8 = (v11 >> ((1) & 63))
                        _ = v19
                    case "\n\t":
                        if ((0.1 > (3.14159 + 1.5)) && false) {
                            if true {
                                v20 := GenS1{F0: -1, F1: false, F2: true}
                                v3 = (^(v3))
                                var v21 int32 = v5
                                _ = v20
                                _ = v21
                            }
                        }
                    default:
                        v2 = v2
                    }
                case "foo bar":
                    var v22 int64 = -9223372036854775808
                    _ = v22
                }
            }
        } else {
            v23 := GenS3{F0: (v4 + v4), F1: 2}
            var v24 uint = 975
            _ = v23
            _ = v24
        }
    }
    v25 := GenS1{F0: 936, F1: (2 > (v11 * 255)), F2: false}
    if len(v1) > 0 {
        v1[0] = (v4 & v4)
    }
    switch v12 {
    case "hello":
        if false {
            switch genFn0(1, 980, 662) {
            case 4294967295:
                v26 := GenS4{F0: v2, F1: -32768}
                _ = v26
            default:
                if true {
                    for fi27 := 0; fi27 < 11; fi27++ {
                        _ = len(v1)
                        var v28 int32 = v5
                        v3 += v3
                        _ = fi27
                        _ = v28
                    }
                    var v29 int64 = int64(v3)
                    _ = v29
                } else {
                    v11 -= 127
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v5 = -1
                }
            }
            v5 = 2147483647
            v5 += (-(v5))
        } else {
            if ((!true) || false) {
                for fi30 := 0; fi30 < 20; fi30++ {
                    for fi31 := 0; fi31 < 17; fi31++ {
                        switch "hello" {
                        case "hello":
                            v12 = v12
                        case "":
                            if true {
                                switch fi31 {
                                case -393:
                                    v2 -= (^(v2))
                                case 4294967295:
                                    fmt.Println(v3, v4)
                                default:
                                    fmt.Println(v12)
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v32 int = int(v11)
                                _ = v32
                            } else {
                                switch (fi31 - genFn0(-1, 1, v3)) {
                                case -2147483648:
                                    if len(v1) > 0 {
                                        v1[0] = int(v2)
                                    }
                                case 4294967295:
                                    _ = len(v1)
                                case 9223372036854775807:
                                    v3 += v3
                                default:
                                    fmt.Println(v11, v4)
                                }
                            }
                        default:
                            var v33 int64 = 9223372036854775807
                            _ = v33
                        }
                        fmt.Println(fi31)
                        _ = fi31
                    }
                    v34 := GenS2{F0: v11, F1: true}
                    if len(v1) > 0 {
                        v1[0] = -9223372036854775808
                    }
                    _ = fi30
                    _ = v34
                }
                if false {
                    _ = len(v1)
                    for fi35 := 0; fi35 < 12; fi35++ {
                        v5 += 719
                        var v36 int = 42
                        v3 = (^(v3))
                        _ = fi35
                        _ = v36
                    }
                    var v37 uint = 2
                    _ = v37
                } else {
                    for fi38 := 0; fi38 < 20; fi38++ {
                        v39 := GenS5{F0: uint16(v11), F1: uint(v2), F2: v5}
                        if len(v1) > 0 {
                            v1[0] = (v4 | fi38)
                        }
                        switch fi38 {
                        case 32767:
                            v4 = (v4 / (((^(fi38))) | 1))
                        default:
                            var v40 int8 = 127
                            _ = v40
                        }
                        _ = fi38
                        _ = v39
                    }
                    var v41 float32 = ((0.0 * (0.1 + 1e10)) - 1e10)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v41
                }
                var v42 float32 = 3.14159
                _ = v42
            }
            v3 -= 18446744073709551615
            var v43 uint64 = v3
            _ = v43
        }
    default:
        if true {
            v12 = ("a" + (v12 + v12))
            if (-289 >= -9223372036854775808) {
                v12 = (((v12 + (v12 + "foo bar")) + "\n\t") + "a")
            } else {
                var v44 uint64 = (v3 >> ((965) & 63))
                switch v12 {
                case "a":
                    v11 = uint8(v2)
                case "":
                    _ = len(v1)
                case "hello":
                    for fi45 := 0; fi45 < 1; fi45++ {
                        v4 = 4294967295
                        var v46 int64 = -9223372036854775808
                        _ = fi45
                        _ = v46
                    }
                }
                switch v12 {
                case "a":
                    v47 := GenS6{F0: int64(v2), F1: 2147483648, F2: 0}
                    _ = v47
                case "":
                    v11 = 127
                }
                _ = v44
            }
            v3 -= v3
        } else {
            v48 := GenS1{F0: 592, F1: true, F2: false}
            _ = v48
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
