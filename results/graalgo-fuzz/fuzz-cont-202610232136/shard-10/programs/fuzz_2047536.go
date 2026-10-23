package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int32 }
type GenS2 struct { F0 int8; F1 int32; F2 uint16 }

func main() {
    var v1 uint64 = 2
    var v2 bool = (true && true)
    var v3 bool = v2
    var v4 bool = v2
    var v5 uint32 = 4294967295
    fmt.Println(v2, v1, v2)
    var v6 uint16 = 1
    switch -1 {
    case 127:
        var v7 int8 = 1
        _ = v7
    case 2147483647:
        var v8 uint = 2
        _ = v8
    case 4294967295:
        v3 = (!false)
    }
    v6 += (v6 | 632)
    switch ("" + "") {
    case "foo bar":
        v6 = (v6 - v6)
    case "hello":
        var v9 uint = 42
        _ = v9
    default:
        switch ("hello" + ("foo bar" + (("" + "hello") + ""))) {
        case "\n\t":
            if v4 {
                var v10 uint8 = 42
                _ = v10
            } else {
                if v2 {
                    v11 := []int{-2147483648, -9223372036854775808, 14}
                    _ = v11
                }
                for fi12 := 0; fi12 < 18; fi12++ {
                    for fi13 := 0; fi13 < 8; fi13++ {
                        fmt.Println(v5, fi13)
                        if v3 {
                            if ("" != ("foo bar" + "")) {
                                v14 := GenS1{F0: -426, F1: (-(fi13)), F2: 42}
                                _ = v14
                            } else {
                                var v15 float64 = -2.25
                                if v4 {
                                    v16 := []int{127, 42, (fi13 >> ((42) & 63))}
                                    var v17 float64 = 0.1
                                    _ = v16
                                    _ = v17
                                }
                                _ = v15
                            }
                            var v18 int16 = 0
                            var v19 int64 = 784
                            _ = v18
                            _ = v19
                        }
                        _ = fi13
                    }
                    _ = fi12
                }
                v5 += 65535
            }
        default:
            if v3 {
                for fi20 := 0; fi20 < 2; fi20++ {
                    fmt.Println(v4)
                    v21 := GenS2{F0: 42, F1: 185, F2: v6}
                    _ = fi20
                    _ = v21
                }
                v22 := GenS1{F0: 0, F1: 9223372036854775807, F2: 42}
                var v23 uint32 = 918
                _ = v22
                _ = v23
            } else {
                fmt.Println(v6, v6, v5)
                for fi24 := 0; fi24 < 12; fi24++ {
                    var v25 float32 = float32(((3.14159 + 1e10) - 0.1))
                    if (v25 <= (v25 + v25)) {
                        for fi26 := 0; fi26 < 3; fi26++ {
                            if v3 {
                                v27 := []int{(fi26 * fi24), fi26, fi24, (fi26 << ((924) & 63))}
                                _ = v27
                            } else {
                                var v28 uint32 = 0
                                v29 := []int{127, (fi24 << ((42) & 63)), (fi24 << ((65535) & 63))}
                                var v30 float64 = float64(float32(1.5))
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            _ = fi26
                        }
                        var v31 uint = 4294967295
                        _ = v31
                    } else {
                        v1 = 4294967295
                    }
                    _ = fi24
                    _ = v25
                }
            }
        }
    }
    v32 := GenS1{F0: 471, F1: 4294967295, F2: -1}
    v33 := []int{2147483647, -713}
    if len(v33) > 0 {
        v33[0] = 1
    }
    var v34 int64 = 0
    var v35 int32 = 42
    var v36 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
