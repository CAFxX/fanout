package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }
type GenS2 struct { F0 int64; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint64; F1 int32; F2 int8 }
type GenS4 struct { F0 int; F1 uint16; F2 bool }
type GenS5 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 string = ""
    var v3 float64 = 0.1
    v4 := make(map[string]uint32)
    var v5 int32 = 1
    for fi6 := 0; fi6 < 14; fi6++ {
        if true {
            fmt.Println(v5)
        }
        for fi7 := 0; fi7 < 12; fi7++ {
            fmt.Println(v5, v3)
            _ = fi7
        }
        v5 = v5
        _ = fi6
    }
    _ = len(v4)
    var v8 int16 = 0
    var v9 uint64 = 219
    v10 := GenS1{F0: true, F1: -308, F2: v2}
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 uint32 = 2147483647
        for fi13 := 0; fi13 < 13; fi13++ {
            switch v2 {
            case "hello":
                if false {
                    if true {
                        if true {
                            if ((0 == 42) || ((v12 << ((1) & 63)) != v12)) {
                                v8 -= (v8 << ((42) & 63))
                                v2 = ""
                                var v14 uint32 = 220
                                _ = v14
                            }
                            if true {
                                var v15 string = v2
                                v9 = 0
                                _ = v15
                            } else {
                                v9 = (^(v9))
                                v2 = v2
                                if true {
                                    v8 = 383
                                } else {
                                    v8 -= (v8 & (v8 / ((0) | 1)))
                                    var v16 int16 = (v8 ^ (v8 >> ((116) & 63)))
                                    v17 := []int{fi13, 161, fi11}
                                    if true {
                                        if (v8 >= -1) {
                                            v5 = (^(v5))
                                            fmt.Println(v12, v8)
                                        } else {
                                            var v18 uint8 = 236
                                            v19 := GenS2{F0: -1, F1: (v9 * (v9 * v9)), F2: int16(v9)}
                                            _ = v18
                                            _ = v19
                                        }
                                        switch (("\n\t" + "") + (v2 + "a")) {
                                        case "hello":
                                            fmt.Println(v16, fi11, v16)
                                        case "a":
                                            if (((303 == 65535) || (false || (255 == 42))) || ((((v3 / 1e10) + -2.25) - 0.1) <= ((3.14159 * (v3 - v3)) * float64(((0.1 + 1e10) * 0.0))))) {
                                                v5 = -454
                                                var v20 uint8 = 53
                                                _ = v20
                                            }
                                        }
                                    } else {
                                        v17 = append(v17, (fi13 ^ (^(fi13))))
                                        switch v2 {
                                        case "hello":
                                            fmt.Println(v9)
                                        }
                                        v21 := GenS3{F0: (v9 >> ((196) & 63)), F1: 253, F2: 0}
                                        _ = v21
                                    }
                                    v3 = 3.14159
                                    _ = v16
                                    _ = v17
                                }
                            }
                        }
                        v22 := GenS4{F0: (fi13 + 0), F1: 32767, F2: false}
                        v23 := GenS5{F0: (!((42 < 42) && true)), F1: 255}
                        _ = v22
                        _ = v23
                    }
                }
            default:
                _ = len(v1)
            }
            _ = fi13
        }
        v12 -= 65535
        _ = fi11
        _ = v12
    }
    var v24 uint64 = 2
    var v25 int64 = -9223372036854775808
    v26 := v1[4294967295]
    fmt.Println(v26)
    delete(v4, v2)
    v27 := v4[v2]
    var v28 uint16 = 951
    var v29 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
