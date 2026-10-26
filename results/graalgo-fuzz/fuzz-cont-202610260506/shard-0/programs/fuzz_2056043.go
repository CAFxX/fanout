package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint = 892
    var v2 int16 = 32767
    var v3 uint64 = 18446744073709551615
    var v4 int32 = 767
    var v5 uint32 = 0
    v6 := GenS1{F0: 127, F1: v3, F2: 1}
    fmt.Println(v3, v2)
    v7 := GenS2{F0: 42, F1: 42}
    v1 -= v1
    v8 := GenS3{F0: 0}
    v9 := []int{-490, 32768}
    if true {
        v5 = 4294967295
        v2 += v2
    } else {
        fmt.Println(v4, v1, v1)
        fmt.Println(v2)
    }
    if len(v9) > 0 {
        v9[0] = 0
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    switch 32768 {
    case 1:
        var v10 uint32 = (v5 / ((v5) | 1))
        _ = v10
    case 32767:
        v11 := GenS1{F0: 0, F1: (v3 * v3), F2: 1}
        _ = v11
    default:
        v12 := GenS1{F0: -2147483648, F1: v3, F2: 42}
        _ = v12
    }
    v13 := GenS3{F0: 1}
    v14 := GenS1{F0: 1, F1: uint64(v5), F2: 155}
    switch "hello" {
    case "":
        if (((false && false) && true) && (((v1 + v1) == v1) || (53 < 128))) {
            v4 += v4
            v9 = append(v9, 190)
            if len(v9) > 0 {
                v9[0] = 1
            }
        }
    case "foo bar":
        if (42 >= 0) {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                for fi16 := 0; fi16 < 15; fi16++ {
                    if (923 != (^(v5))) {
                        if true {
                            var v17 int16 = v2
                            var v18 float64 = 1e10
                            _ = v17
                            _ = v18
                        } else {
                            if len(v9) > 0 {
                                v9[0] = fi16
                            }
                        }
                    } else {
                        var v19 uint64 = (v3 >> (((v1 | 18446744073709551615)) & 63))
                        v19 = v19
                        _ = v19
                    }
                    if (!false) {
                        fmt.Println(fi15)
                        var v20 uint16 = 458
                        var v21 float32 = float32((0.0 * ((((3.14159 * (0.1 + -2.25)) + 0.1) - 1.5) + 1.5)))
                        _ = v20
                        _ = v21
                    } else {
                        v3 -= (v3 / ((uint64(v1)) | 1))
                        v2 = v2
                    }
                    if (!true) {
                        var v22 int8 = 0
                        if true {
                            v23 := GenS2{F0: 234, F1: 9223372036854775807}
                            _ = v23
                        } else {
                            v24 := GenS3{F0: -9223372036854775808}
                            _ = v24
                        }
                        v1 += (v1 ^ (v1 << ((466) & 63)))
                        _ = v22
                    }
                    _ = fi16
                }
                _ = fi15
            }
        } else {
            if ((^(v5)) >= (v5 | 404)) {
                v4 = v4
                fmt.Println(v1)
            }
        }
    case "\n\t":
        if (v4 != (v4 % ((-2147483648) | 1))) {
            if (!(781 >= -196)) {
                var v25 int64 = -814
                _ = v25
            } else {
                fmt.Println(v2, v1)
                v26 := GenS2{F0: 247, F1: 42}
                _ = v26
            }
            var v27 string = ((("hello" + "") + ("\n\t" + "foo bar")) + ("hello" + "a"))
            v3 = 4294967295
            _ = v27
        }
    default:
        _ = len(v9)
    }
    if len(v9) > 0 {
        v9[0] = 623
    }
    if (42 == 0) {
        if (-682 < 998) {
            fmt.Println(v2, v2, v3)
            if true {
                var v28 bool = false
                var v29 int16 = (v2 ^ (v2 * v2))
                v28 = ((((222 <= v3) || (v28 || (57 > 86))) && (!false)) && v28)
                _ = v28
                _ = v29
            } else {
                var v30 float64 = ((3.14159 / 1.5) * float64((0.1 - 1e10)))
                if false {
                    var v31 uint32 = (v5 * v5)
                    _ = v31
                }
                _ = v30
            }
            v5 -= v5
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
