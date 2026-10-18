package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint32; F2 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32 }

func main() {
    v1 := []uint32{2, 86, 2147483647, 17}
    v2 := []uint32{4294967295, 4294967295, 2147483647, 592}
    v3 := make(map[int]bool)
    var v4 uint16 = 0
    var v5 int32 = 0
    v6 := GenS1{F0: 59, F1: 1, F2: v5}
    var v7 float64 = 0.0
    v8 := GenS2{F0: 2147483647}
    var v9 uint64 = 18446744073709551615
    if true {
        switch "a" {
        case "hello":
            v10 := v3[-914]
            _ = v10
        case "\n\t":
            v11 := GenS2{F0: 2}
            _ = v11
        case "foo bar":
            v9 += v9
        default:
            v7 = v7
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            switch -496 {
            case 32767:
                v5 = (^(v5))
            case 0:
                fmt.Println(v9, fi12, fi12)
            case 127:
                for fi13 := 0; fi13 < 10; fi13++ {
                    if true {
                        switch 706 {
                        case 2147483647:
                            switch (fi13 + 42) {
                            case 9223372036854775807:
                                if false {
                                    fmt.Println(v5, v9, v9)
                                    v14 := GenS1{F0: -1, F1: (v9 / (((v9 / ((v9) | 1))) | 1)), F2: -1}
                                    if (!(v5 == v5)) {
                                        if false {
                                            var v15 int8 = -128
                                            _ = v15
                                        }
                                    }
                                    _ = v14
                                } else {
                                    var v16 uint64 = v9
                                    _ = v16
                                }
                            }
                        }
                        v17 := GenS1{F0: 127, F1: (v9 | 18446744073709551615), F2: 1}
                        v18 := GenS2{F0: 0}
                        if false {
                            delete(v3, (fi12 << ((1) & 63)))
                            if (uint32(v4) > 2147483647) {
                                v5 = (-(v5))
                                var v19 int = -2147483648
                                _ = v19
                            } else {
                                v20 := GenS3{F0: 57, F1: v4}
                                _ = v20
                            }
                            v7 += 3.14159
                        } else {
                            v21 := GenS4{F0: 42, F1: 65535, F2: 127}
                            _ = v21
                        }
                        if (0.1 != -2.25) {
                            v22 := GenS1{F0: 99, F1: (v9 >> ((18446744073709551615) & 63)), F2: 2147483647}
                            v4 += (v4 ^ (^(v4)))
                            _ = v22
                        } else {
                            if false {
                                if (1 > 220) {
                                    v7 += (v7 + 0.0)
                                    var v23 int32 = v5
                                    _ = v23
                                } else {
                                    v4 = (v4 << ((290) & 63))
                                    v24 := GenS3{F0: 1, F1: v4}
                                    _ = v24
                                }
                            } else {
                                fmt.Println(fi12, v4, v5)
                            }
                            if false {
                                v25 := v3[(fi12 + 42)]
                                switch -1 {
                                case -9223372036854775808:
                                    v5 = v5
                                default:
                                    switch fi13 {
                                    case 74:
                                        v26 := GenS1{F0: -128, F1: 18446744073709551615, F2: (v5 * (-(v5)))}
                                        _ = v26
                                    case -1:
                                        var v27 bool = (!v25)
                                        _ = v27
                                    case 695:
                                        fmt.Println(fi12, v7, fi13)
                                    }
                                }
                                fmt.Println(v5, v25)
                                _ = v25
                            } else {
                                var v28 uint8 = 255
                                delete(v3, (-(fi13)))
                                v4 += (^(v4))
                                _ = v28
                            }
                            v29 := GenS5{F0: "hello"}
                            _ = v29
                        }
                        _ = v17
                        _ = v18
                    } else {
                        _ = len(v1)
                    }
                    v30 := GenS6{F0: 2147483648}
                    _ = fi13
                    _ = v30
                }
            }
            _ = fi12
        }
    } else {
        var v31 uint = 18446744073709551615
        v9 -= v9
        _ = v31
    }
    var v32 int64 = 9223372036854775807
    v5 -= v5
    var v33 uint32 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
