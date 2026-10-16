package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint64; F2 int16 }
func genFn0(p0 bool) uint8 {
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: 2147483648}
        fmt.Println(fi6, p0)
        _ = fi6
        _ = v7
    }
    var v8 int = -29
    _ = p0
    _ = v8
    return 0
}
func genFn1(p0 uint64) uint64 {
    v9 := []int{2147483647, 452, 415, 32767}
    p0 = p0
    _ = p0
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 172
    var v2 uint8 = v1
    var v3 float32 = 3.14159
    v4 := []uint32{1, 4294967295, 42, uint32(v2)}
    var v5 int16 = 308
    fmt.Println(v5)
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int8 = 127
        if ((0.0 * 1e10) <= 0.0) {
            if true {
                if (true || (0 >= 643)) {
                    if (v3 != -2.25) {
                        if len(v4) > 0 {
                            v4[0] = 4294967295
                        }
                        for fi12 := 0; fi12 < 20; fi12++ {
                            switch "hello" {
                            case "foo bar":
                                var v13 uint16 = 42
                                _ = v13
                            case "":
                                v5 += v5
                            default:
                                var v14 float64 = (0.0 / 0.1)
                                _ = v14
                            }
                            _ = fi12
                        }
                    } else {
                        var v15 int8 = 42
                        _ = v15
                    }
                    var v16 uint16 = 725
                    v11 -= v11
                    _ = v16
                } else {
                    v17 := GenS1{F0: 4294967295}
                    switch ("a" + "") {
                    case "foo bar":
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                    case "a":
                        switch (fi10 & -126) {
                        case 127:
                            fmt.Println(v5, v1, v3)
                        case -901:
                            var v18 uint32 = 2147483647
                            _ = v18
                        default:
                            v2 = v2
                        }
                    case "":
                        fmt.Println(v11, v11)
                    }
                    _ = v17
                }
                _ = len(v4)
                v4 = append(v4, 65535)
            } else {
                switch "foo bar" {
                case "\n\t":
                    for fi19 := 0; fi19 < 2; fi19++ {
                        fmt.Println(fi10, fi10, v3)
                        _ = fi19
                    }
                default:
                    if (v2 != v1) {
                        if len(v4) > 0 {
                            v4[0] = 176
                        }
                    } else {
                        v3 = (0.0 + ((v3 / v3) * ((float32(1.5) * (3.14159 - v3)) + v3)))
                        v20 := GenS2{F0: 372, F1: 42, F2: (v5 & (v5 / ((v5) | 1)))}
                        _ = v20
                    }
                }
                fmt.Println(v11, v5, fi10)
            }
        }
        _ = fi10
        _ = v11
    }
    var v21 uint64 = 42
    v5 = v5
    v22 := GenS1{F0: 2147483647}
    var v23 int32 = -829
    v24 := GenS2{F0: 119, F1: (v21 << ((4294967295) & 63)), F2: (v5 / ((v5) | 1))}
    v25 := GenS2{F0: 42, F1: v21, F2: v5}
    v5 -= (v5 / ((1000) | 1))
    if false {
        fmt.Println(v5, v3, v21)
    } else {
        fmt.Println(v2, v1, v23)
    }
    if len(v4) > 0 {
        v4[0] = 65535
    }
    var v26 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
