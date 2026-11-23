package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int16; F1 bool; F2 int16 }

func main() {
    var v1 int8 = 1
    v2 := []uint32{2, 4294967295, 706}
    v3 := make(map[uint32]int)
    var v4 uint = 65535
    v5 := []int{-305, 32767, -9223372036854775808}
    v6 := make(map[int]uint32)
    v4 -= 255
    fmt.Println(v1, v1)
    v7 := GenS1{F0: 0, F1: 76}
    switch ("" + ("foo bar" + "foo bar")) {
    case "a":
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 int8 = 31
            if (0.1 != 1.5) {
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 uint64 = 1
                    _ = fi10
                    _ = v11
                }
                v4 += 0
                switch fi8 {
                case -486:
                    if (1 == (v4 & (v4 | 504))) {
                        switch (fi8 - 2147483647) {
                        case 1:
                            var v12 int64 = 0
                            _ = v12
                        case 127:
                            if len(v2) > 0 {
                                v2[0] = 2147483647
                            }
                        default:
                            for fi13 := 0; fi13 < 13; fi13++ {
                                if true {
                                    switch "" {
                                    case "\n\t":
                                        var v14 float32 = (((1.5 * float32((-2.25 * 1e10))) / 1.0) / ((-2.25 + -2.25) * (-2.25 * (0.0 + 0.1))))
                                        _ = v14
                                    default:
                                        v1 = (v1 + (v1 ^ (v9 << (((v4 ^ v4)) & 63))))
                                    }
                                }
                                _ = fi13
                            }
                        }
                        for fi15 := 0; fi15 < 17; fi15++ {
                            v9 += (v9 ^ 127)
                            _ = fi15
                        }
                    }
                case -463:
                    switch (("\n\t" + ((("a" + "a") + ("a" + "")) + "a")) + (((("a" + ("hello" + ("foo bar" + ("a" + "foo bar")))) + "a") + ("a" + "foo bar")) + ("\n\t" + "hello"))) {
                    case "":
                        if ((fi8 >> ((659) & 63)) >= (fi8 ^ fi8)) {
                            if ((42 == 2147483647) || true) {
                                for fi16 := 0; fi16 < 7; fi16++ {
                                    var v17 uint16 = 902
                                    _ = fi16
                                    _ = v17
                                }
                                if false {
                                    var v18 int = (fi8 + fi8)
                                    _ = len(v3)
                                    v4 += 404
                                    _ = v18
                                }
                            } else {
                                for fi19 := 0; fi19 < 2; fi19++ {
                                    var v20 bool = (false && (670 == 2))
                                    if (!v20) {
                                        if (42 >= 2) {
                                            var v21 bool = v20
                                            _ = v21
                                        }
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v22 bool = ((!false) || (42 <= 571))
                                            _ = v22
                                        case "":
                                            var v23 int64 = -1
                                            _ = v23
                                        case "hello":
                                            if len(v2) > 0 {
                                                v2[0] = 453
                                            }
                                        default:
                                            var v24 uint16 = 356
                                            _ = v24
                                        }
                                        var v25 uint32 = 24
                                        _ = v25
                                    }
                                    v26 := GenS1{F0: 255, F1: (fi19 & fi8)}
                                    var v27 int = (fi19 | fi19)
                                    var v28 string = "\n\t"
                                    v28 = v28
                                    _ = fi19
                                    _ = v20
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            }
                            v1 -= (v9 % (((-(v9))) | 1))
                        }
                    }
                }
            }
            _ = fi8
            _ = v9
        }
    default:
        if false {
            for fi29 := 0; fi29 < 16; fi29++ {
                for fi30 := 0; fi30 < 13; fi30++ {
                    v1 = -128
                    if true {
                        v31 := GenS2{F0: 379}
                        _ = v31
                    } else {
                        var v32 int64 = -1
                        _ = v32
                    }
                    _ = fi30
                }
                _ = fi29
            }
            v4 = v4
            for fi33 := 0; fi33 < 2; fi33++ {
                if true {
                    fmt.Println(v4, fi33, v4)
                    var v34 int8 = 42
                    _ = v34
                }
                _ = fi33
            }
        } else {
            var v35 int32 = -1
            var v36 uint16 = 281
            if false {
                for fi37 := 0; fi37 < 15; fi37++ {
                    v38 := GenS3{F0: (v36 & 705), F1: 127}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi37
                    _ = v38
                }
                v39 := v3[0]
                _ = v39
            }
            _ = v35
            _ = v36
        }
    }
    fmt.Println(v4)
    var v40 float64 = (3.14159 - 1.5)
    v4 -= v4
    var v41 int = 0
    v42 := GenS4{F0: 42, F1: 255}
    v43 := GenS5{F0: int16(v4), F1: true, F2: -58}
    var v44 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
