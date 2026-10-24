package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint32 }

func main() {
    var v1 uint = 65535
    var v2 float64 = 3.14159
    v3 := []int{-2147483648, 9223372036854775807}
    var v4 int16 = 0
    v5 := []uint32{65535}
    v6 := make(map[uint32]uint32)
    switch -945 {
    case -1:
        if true {
            v1 = (v1 + v1)
            v6[1] = 2147483648
        }
    case 4294967295:
        fmt.Println(v1)
    case 127:
        var v7 uint64 = 2
        _ = v7
    }
    var v8 uint32 = 668
    if (204 == 43) {
        var v9 float64 = v2
        switch 9223372036854775807 {
        case 4294967295:
            v9 -= ((-2.25 / float64(-2.25)) * (v9 * (v9 + float64(0.0))))
        case 32767:
            v2 = (float64(1.5) * (-2.25 - (((((float64((float32(0.0) / 1e10)) * v9) + 0.0) * v2) - v9) - (-2.25 * (v9 / 1.5)))))
        }
        _ = v9
    }
    v1 = (v1 - v1)
    if true {
        v1 += 65535
    } else {
        if (255 >= 1) {
            v10 := GenS1{F0: -1, F1: v1}
            _ = v10
        }
        if len(v5) > 0 {
            v5[0] = v8
        }
        if false {
            if false {
                fmt.Println(v4)
                for fi11 := 0; fi11 < 9; fi11++ {
                    v2 -= (1e10 * (float64(1.5) * float64(((0.1 / 3.14159) * float32(v2)))))
                    _ = fi11
                }
                if ((1.5 != float64(1.5)) && false) {
                    var v12 uint16 = 0
                    _ = v12
                } else {
                    switch 9223372036854775807 {
                    case 127:
                        switch 4294967295 {
                        case 4294967295:
                            v6[(v8 + v8)] = 1
                        case 127:
                            fmt.Println(v1, v1, v2)
                        }
                    case -2147483648:
                        for fi13 := 0; fi13 < 12; fi13++ {
                            if (!(uint8(v8) == uint8(v8))) {
                                for fi14 := 0; fi14 < 6; fi14++ {
                                    var v15 uint32 = (^(v8))
                                    var v16 int64 = 0
                                    var v17 int8 = 42
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                                v18 := GenS1{F0: -57, F1: 386}
                                _ = v18
                            }
                            fmt.Println(v2)
                            if (v1 == (v1 | (^(v1)))) {
                                if false {
                                    var v19 int16 = (^(v4))
                                    _ = v19
                                }
                                switch fi13 {
                                case 464:
                                    var v20 float32 = ((3.14159 * (((0.0 * 1.5) - 1e10) + -2.25)) - (((0.1 - 0.0) + 1e10) / ((3.14159 / -2.25) / 1.0)))
                                    _ = v20
                                case 32768:
                                    if false {
                                        for fi21 := 0; fi21 < 4; fi21++ {
                                            switch -675 {
                                            case 9223372036854775807:
                                                v22 := GenS2{F0: ("hello" + ("hello" + "foo bar")), F1: 42}
                                                _ = v22
                                            }
                                            var v23 int8 = 39
                                            _ = fi21
                                            _ = v23
                                        }
                                        _ = len(v6)
                                    } else {
                                        v24 := GenS2{F0: (("foo bar" + "") + ""), F1: 2}
                                        var v25 uint32 = (v8 ^ 564)
                                        v3 = append(v3, fi13)
                                        _ = v24
                                        _ = v25
                                    }
                                }
                                v1 = (v1 << (((v1 - 1)) & 63))
                            }
                            _ = fi13
                        }
                    case 4294967295:
                        fmt.Println(v2)
                    }
                }
            }
            v6[v8] = 377
        }
    }
    fmt.Println(v2, v2)
    delete(v6, uint32(v1))
    v1 -= (v1 / (((v1 % ((v1) | 1))) | 1))
    var v26 uint = 175
    v2 = 3.14159
    v1 = 18446744073709551615
    if false {
        v27 := GenS3{F0: -789, F1: (v8 % (((v8 * v8)) | 1))}
        _ = v27
    }
    var v28 float64 = (0.1 + (v2 - v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
