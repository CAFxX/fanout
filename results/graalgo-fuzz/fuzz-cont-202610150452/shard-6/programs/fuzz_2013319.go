package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 uint64; F1 uint8 }
func genFn0() uint {
    var v5 bool = true
    v5 = (65535 == 4294967295)
    _ = v5
    return 0
}

func main() {
    var v1 int32 = 2147483647
    var v2 string = "a"
    v3 := make(map[uint32]int)
    var v4 int = -9223372036854775808
    v6 := []int{int(v1), (v4 << ((genFn0()) & 63)), (v4 + v4)}
    v3[2147483647] = v4
    if (0 != -193) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi7 := 0; fi7 < 17; fi7++ {
            if ((!false) && true) {
                if true {
                    if (((float32(0.0) + -2.25) != 0.0) && false) {
                        if (18446744073709551615 >= 2) {
                            var v8 uint32 = 58
                            _ = v8
                        } else {
                            if ((!true) || ((42 > 165) && true)) {
                                if true {
                                    v4 += v4
                                    if (false && false) {
                                        v4 -= 2147483647
                                        var v9 int = 128
                                        var v10 uint = 18446744073709551615
                                        _ = v10
                                        _ = v9
                                    }
                                    var v11 int16 = 1000
                                    _ = v11
                                }
                                var v12 string = "hello"
                                v13 := GenS1{F0: 42}
                                _ = v12
                                _ = v13
                            }
                            if false {
                                switch fi7 {
                                case 127:
                                    fmt.Println(v1)
                                case -2147483648:
                                    var v14 int64 = -9223372036854775808
                                    _ = v14
                                case 32767:
                                    fmt.Println(v2)
                                default:
                                    v15 := GenS1{F0: 109}
                                    _ = v15
                                }
                            }
                            v1 -= -274
                        }
                        var v16 int64 = 1
                        _ = v16
                    }
                    v4 += v4
                    var v17 uint16 = 42
                    _ = v17
                }
            } else {
                var v18 uint32 = 2147483648
                if false {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        if true {
                            var v20 uint32 = 42
                            _ = v20
                        } else {
                            switch (v2 + ((v2 + v2) + v2)) {
                            case "a":
                                v18 -= (^(v18))
                            case "\n\t":
                                v18 = (v18 ^ v18)
                            case "hello":
                                v1 += (v1 ^ (v1 / ((v1) | 1)))
                            default:
                                var v21 int32 = v1
                                _ = v21
                            }
                            v22 := GenS2{F0: -1, F1: 1, F2: 18446744073709551615}
                            _ = v22
                        }
                        _ = fi19
                    }
                } else {
                    if false {
                        switch (v4 ^ (fi7 * v4)) {
                        case -1:
                            switch (v2 + (v2 + ((v2 + v2) + v2))) {
                            case "a":
                                var v23 int64 = 0
                                _ = v23
                            default:
                                if (!(!(!false))) {
                                    var v24 int = 216
                                    _ = v24
                                }
                            }
                        }
                    }
                }
                _ = v18
            }
            _ = fi7
        }
        v2 = ((v2 + v2) + "\n\t")
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        if len(v6) > 0 {
            v6[0] = (fi25 >> ((65535) & 63))
        }
        if (2147483647 > 781) {
            v4 += v4
            var v26 int = (fi25 / ((v4) | 1))
            var v27 int = fi25
            _ = v26
            _ = v27
        }
        if ((1.5 + (0.1 * -2.25)) < float32(((0.0 - 1.5) - (float64(((0.0 * ((0.0 + 0.1) / 1.0)) - 0.1)) * (-2.25 * 1e10))))) {
            var v28 int8 = 1
            v29 := GenS3{F0: 18446744073709551615, F1: 42}
            _ = v28
            _ = v29
        } else {
            v2 = ("\n\t" + (("\n\t" + ((v2 + ((v2 + (v2 + "foo bar")) + v2)) + v2)) + v2))
            var v30 string = "a"
            _ = v30
        }
        _ = fi25
    }
    v31 := GenS2{F0: 9223372036854775807, F1: 732, F2: 2}
    v1 = v1
    var v32 uint8 = 228
    v33 := GenS2{F0: 0, F1: 2, F2: 2}
    if len(v6) > 0 {
        v6[0] = (v4 | (v4 * 9223372036854775807))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v6
}
