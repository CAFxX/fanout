package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint8; F1 uint8 }
func genFn0() uint32 {
    v7 := GenS1{F0: -930, F1: false}
    var v8 int32 = 0
    if (853 == 340) {
        switch 42 {
        case -9223372036854775808:
            v9 := []int{0}
            _ = v9
        }
    } else {
        v10 := GenS1{F0: -1, F1: (32768 == -915)}
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 float32 = (1e10 - 1.5)
            for fi13 := 0; fi13 < 1; fi13++ {
                v8 = 1
                v8 += 1
                switch ("hello" + "foo bar") {
                case "a":
                    if false {
                        switch (fi13 << ((0) & 63)) {
                        case -9223372036854775808:
                            v14 := []int{(fi13 ^ 0), fi11}
                            _ = v14
                        case 127:
                            v12 -= -2.25
                        case 927:
                            if (2 <= 42) {
                                if (2 >= 445) {
                                    v8 = 2147483647
                                }
                            } else {
                                var v15 int = fi13
                                v16 := GenS1{F0: -1, F1: false}
                                if true {
                                    v12 = 0.1
                                } else {
                                    var v17 int8 = -1
                                    _ = v17
                                }
                                _ = v15
                                _ = v16
                            }
                        default:
                            if (v8 != v8) {
                                fmt.Println(v12, fi11)
                            }
                        }
                        if false {
                            v18 := []int{79}
                            if len(v18) > 0 {
                                v18[0] = fi11
                            }
                            fmt.Println(v8, v12)
                            _ = v18
                        } else {
                            fmt.Println(fi13)
                        }
                    }
                case "hello":
                    var v19 uint8 = 1
                    _ = v19
                case "\n\t":
                    v8 = (v8 % (((^(v8))) | 1))
                default:
                    v20 := GenS2{F0: 19, F1: 42}
                    _ = v20
                }
                if true {
                    var v21 int16 = 5
                    v22 := GenS1{F0: -94, F1: (false || (224 <= 127))}
                    _ = v21
                    _ = v22
                }
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v23 int = 2147483647
    _ = v23
    _ = v7
    _ = v8
    return 639
}

func main() {
    var v1 int32 = 42
    v2 := []uint32{2, 1, 805, 2}
    var v3 int8 = -11
    v4 := []int{}
    v5 := []int{9223372036854775807, 128}
    var v6 uint8 = 168
    fmt.Println(v3, v3, v3)
    v3 -= -69
    v3 -= -1
    switch "hello" {
    case "\n\t":
        if true {
            var v24 int = -9223372036854775808
            _ = v24
        }
    case "hello":
        v3 = v3
    case "foo bar":
        var v25 uint32 = 917
        _ = v25
    }
    fmt.Println(v3)
    v6 = (v6 | v6)
    var v26 uint8 = v6
    v6 -= 146
    v3 = (v3 * v3)
    fmt.Println(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 -= v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
