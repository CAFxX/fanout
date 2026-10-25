package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 int }

func main() {
    var v1 uint8 = 255
    var v2 int32 = 100
    v3 := make(map[string]uint32)
    var v4 string = "\n\t"
    v5 := make(map[uint32]bool)
    v4 = "foo bar"
    v6 := GenS1{F0: 127}
    var v7 uint32 = 0
    var v8 uint16 = 0
    v8 -= 32767
    var v9 uint32 = (v7 - 105)
    v9 -= (v7 % ((4294967295) | 1))
    if ((((^(v9)) < (v9 / ((v7) | 1))) && (!(35 <= -323))) && false) {
        var v10 bool = ((1.5 != 3.14159) || (!(false || false)))
        _ = v10
    } else {
        if false {
            v11 := GenS2{F0: -106}
            for fi12 := 0; fi12 < 19; fi12++ {
                fmt.Println(fi12, v1, fi12)
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 uint16 = v8
                    v15 := GenS1{F0: -128}
                    v2 = v2
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 16; fi16++ {
                    switch fi16 {
                    case 541:
                        if (2 > v8) {
                            v3[v4] = (v7 + (v9 * (^(v9))))
                            v17 := v5[(v7 ^ (v9 | 65535))]
                            _ = v17
                        } else {
                            fmt.Println(v2, v1)
                        }
                    case 9223372036854775807:
                        switch (-(fi12)) {
                        case 863:
                            v18 := []int{(fi16 << ((42) & 63)), fi16, fi12}
                            _ = v18
                        case 737:
                            var v19 int64 = 9223372036854775807
                            _ = v19
                        case 127:
                            var v20 int16 = -1
                            _ = v20
                        default:
                            v21 := GenS3{F0: (v4 + (v4 + (v4 + v4))), F1: 42}
                            _ = v21
                        }
                    case 1:
                        switch fi12 {
                        case 1:
                            v7 -= 2147483648
                        case -1:
                            v9 = v9
                        default:
                            var v22 float32 = -2.25
                            _ = v22
                        }
                    default:
                        if (128 == (v1 * v1)) {
                            v8 += v8
                            v5[v7] = false
                            if (!false) {
                                switch (v4 + ("\n\t" + v4)) {
                                case "foo bar":
                                    v2 -= 2147483647
                                case "":
                                    v23 := []int{fi16, fi16, 9223372036854775807}
                                    _ = v23
                                }
                                v7 += (^(v7))
                            } else {
                                var v24 float32 = 0.0
                                var v25 string = ("" + v4)
                                var v26 uint = uint(v1)
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                            if (false || true) {
                                var v27 int8 = 40
                                v28 := v5[v7]
                                switch (fi16 % (((fi16 | (fi16 >> ((4294967295) & 63)))) | 1)) {
                                case 571:
                                    fmt.Println(v9, v9)
                                case 605:
                                    if v28 {
                                        var v29 uint = 154
                                        v7 = (v7 | uint32(v1))
                                        var v30 uint16 = v8
                                        _ = v29
                                        _ = v30
                                    } else {
                                        var v31 uint8 = 117
                                        _ = v31
                                    }
                                case 559:
                                    v32 := GenS4{F0: -205, F1: v2}
                                    _ = v32
                                }
                                v33 := GenS5{F0: fi12}
                                var v34 int64 = 0
                                _ = v27
                                _ = v28
                                _ = v33
                                _ = v34
                            } else {
                                switch "" {
                                case "a":
                                    var v35 uint16 = 1
                                    _ = v35
                                }
                            }
                        }
                    }
                    fmt.Println(v2, fi16)
                    var v36 uint = uint(v1)
                    _ = fi16
                    _ = v36
                }
                _ = fi12
            }
            _ = v11
        }
        v9 = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
