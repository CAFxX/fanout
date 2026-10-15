package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int64; F1 int8 }
func genFn0(p0 uint16, p1 string) uint32 {
    p0 -= 313
    p1 = ((((p1 + p1) + "foo bar") + p1) + p1)
    _ = p0
    _ = p1
    return 581
}

func main() {
    v1 := []uint32{2147483647, 2147483647}
    v2 := make(map[uint32]string)
    v3 := []int{127, 32768, 128, 832}
    var v4 float64 = 0.1
    _ = len(v2)
    fmt.Println(v4)
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: fi5, F1: 128}
        v4 += 1.5
        var v7 float32 = (1e10 / 0.1)
        _ = fi5
        _ = v6
        _ = v7
    }
    switch ("\n\t" + "foo bar") {
    case "a":
        var v8 int16 = -502
        _ = v8
    default:
        v4 -= 1.5
    }
    var v9 uint32 = 0
    v10 := GenS1{F0: 127, F1: 128}
    v4 = v4
    var v11 bool = true
    v4 = (0.1 + v4)
    for fi12 := 0; fi12 < 1; fi12++ {
        if v11 {
            v13 := GenS1{F0: (fi12 + (fi12 >> ((65535) & 63))), F1: 255}
            v14 := GenS1{F0: (fi12 | 1), F1: 127}
            switch fi12 {
            case 127:
                v9 = v9
            case 2147483647:
                if ((0 > 128) && (239 != 18446744073709551615)) {
                    if false {
                        v11 = (766 == 2147483647)
                        if (!(false && v11)) {
                            switch (-(fi12)) {
                            case -2147483648:
                                if true {
                                    for fi15 := 0; fi15 < 17; fi15++ {
                                        v16 := GenS1{F0: fi12, F1: 109}
                                        _ = fi15
                                        _ = v16
                                    }
                                    if v11 {
                                        fmt.Println(v4)
                                        v9 -= (v9 - genFn0(32767, "hello"))
                                        if len(v1) > 0 {
                                            v1[0] = 0
                                        }
                                    }
                                } else {
                                    fmt.Println(fi12)
                                    if true {
                                        var v17 float32 = 0.1
                                        v3 = append(v3, fi12)
                                        _ = v17
                                    }
                                }
                            case 2147483647:
                                fmt.Println(v9, fi12)
                            }
                            _ = len(v1)
                        } else {
                            v11 = false
                            switch "hello" {
                            case "\n\t":
                                var v18 float32 = ((3.14159 - 0.0) * 1e10)
                                _ = v18
                            case "hello":
                                switch (fi12 * (fi12 * (^(fi12)))) {
                                case -675:
                                    v11 = (0 >= 1)
                                default:
                                    v19 := GenS1{F0: fi12, F1: 2}
                                    _ = v19
                                }
                            default:
                                v20 := GenS1{F0: 1, F1: 1}
                                _ = v20
                            }
                        }
                    } else {
                        var v21 int64 = 0
                        _ = v21
                    }
                    var v22 uint64 = 1
                    v11 = (42 == 1)
                    _ = v22
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            case -9223372036854775808:
                _ = len(v1)
            }
            _ = v13
            _ = v14
        }
        _ = fi12
    }
    if (802 >= -2147483648) {
        var v23 int32 = -2147483648
        _ = v23
    }
    if v11 {
        var v24 float64 = (v4 + float64(0.0))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = append(v1, v9)
        _ = v24
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := GenS2{F0: -606, F1: -13}
        _ = fi25
        _ = v26
    }
    v27 := GenS1{F0: 9223372036854775807, F1: 221}
    v4 = (v4 + (v4 / (1.5 - (v4 - 3.14159))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}
