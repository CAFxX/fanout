package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int16; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 0
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 1
    var v5 int8 = -1
    var v6 int16 = 42
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: v6, F1: uint(v2)}
        _ = fi7
        _ = v8
    }
    fmt.Println(v4, v2, v6)
    var v9 int32 = 1
    switch "\n\t" {
    case "hello":
        v10 := v1[32768]
        _ = v10
    case "a":
        switch "" {
        case "\n\t":
            v5 = 0
        }
    case "":
        v9 += (v9 ^ 2147483647)
    }
    if (1 != v4) {
        v4 = 204
    } else {
        v1[42] = -2147483648
    }
    fmt.Println(v5, v6, v4)
    v6 -= (v6 >> ((255) & 63))
    switch 0 {
    case -910:
        for fi11 := 0; fi11 < 12; fi11++ {
            v5 += (^(v5))
            if (542 == uint32(v4)) {
                if (1e10 != (-2.25 + 1.5)) {
                    switch "" {
                    case "a":
                        if true {
                            v6 -= 32767
                            var v12 int8 = v5
                            fmt.Println(v9, v4, v12)
                            _ = v12
                        } else {
                            if true {
                                if true {
                                    switch (^(fi11)) {
                                    case 4294967295:
                                        var v13 uint16 = 711
                                        _ = v13
                                    }
                                    v14 := v3[2147483647]
                                    _ = v14
                                }
                                v4 = (v4 & 218)
                                if (true && ((v6 | v6) >= v6)) {
                                    v15 := []int{(fi11 - fi11), 9223372036854775807, fi11}
                                    v2 = 1000
                                    v2 = 18446744073709551615
                                    _ = v15
                                } else {
                                    for fi16 := 0; fi16 < 9; fi16++ {
                                        switch fi16 {
                                        case 1:
                                            v17 := GenS1{F0: v6, F1: 18446744073709551615}
                                            _ = v17
                                        case 9223372036854775807:
                                            v18 := GenS2{F0: (fi11 + (fi16 - (fi16 / ((0) | 1)))), F1: (fi16 - fi16)}
                                            _ = v18
                                        }
                                        v19 := GenS2{F0: (fi16 * fi16), F1: fi11}
                                        var v20 int32 = v9
                                        _ = fi16
                                        _ = v19
                                        _ = v20
                                    }
                                    v21 := GenS2{F0: 0, F1: (fi11 % (((fi11 ^ (-(fi11)))) | 1))}
                                    var v22 int64 = 42
                                    _ = v21
                                    _ = v22
                                }
                            } else {
                                if (((fi11 >= fi11) && true) || false) {
                                    switch int(v2) {
                                    case 1:
                                        v9 -= v9
                                    default:
                                        v2 = v2
                                    }
                                } else {
                                    for fi23 := 0; fi23 < 3; fi23++ {
                                        v4 += (v4 - (v4 | v4))
                                        _ = fi23
                                    }
                                    if ((v6 != (^(v6))) || false) {
                                        v9 = 42
                                        var v24 uint8 = v4
                                        _ = v24
                                    } else {
                                        v25 := GenS1{F0: (v6 * (v6 << ((2) & 63))), F1: 4294967295}
                                        _ = v25
                                    }
                                }
                                v2 -= 18446744073709551615
                                switch "foo bar" {
                                case "a":
                                    if (65535 != 0) {
                                        v6 = (v6 & -655)
                                    }
                                case "\n\t":
                                    if (59 != 0) {
                                        if true {
                                            var v26 uint16 = 2
                                            v2 = (v2 ^ v2)
                                            _ = v26
                                        } else {
                                            if (!(524 == v9)) {
                                                if true {
                                                    v27 := []int{fi11}
                                                    v4 -= 255
                                                    _ = v27
                                                }
                                            }
                                            switch -2147483648 {
                                            case 32768:
                                                v28 := GenS3{F0: (v6 % ((-1) | 1)), F1: 0}
                                                _ = v28
                                            default:
                                                v29 := []int{}
                                                _ = v29
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    case "hello":
                        if (1 < 1) {
                            v6 = v6
                        } else {
                            switch 9223372036854775807 {
                            case -965:
                                v4 = (v4 << ((42) & 63))
                            }
                            v2 = (v2 << ((967) & 63))
                        }
                    }
                }
                var v30 uint = 18446744073709551615
                _ = v30
            } else {
                var v31 bool = false
                if false {
                    var v32 int8 = -1
                    _ = v32
                }
                _ = v31
            }
            _ = fi11
        }
    }
    v33 := []int{}
    v9 -= v9
    _ = len(v1)
    if len(v33) > 0 {
        v33[0] = 32767
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
