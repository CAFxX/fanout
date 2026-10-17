package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 float32 = (float32(-2.25) * (0.1 - 1.5))
    var v2 bool = false
    v3 := make(map[string]string)
    v4 := []uint32{42, 2147483648}
    var v5 int16 = 42
    v6 := GenS1{F0: 955}
    var v7 uint16 = 65535
    fmt.Println(v5, v5, v7)
    var v8 float64 = 1.5
    var v9 uint16 = (v7 - v7)
    for fi10 := 0; fi10 < 2; fi10++ {
        _ = len(v3)
        _ = fi10
    }
    var v11 int8 = -128
    var v12 int64 = -1
    for fi13 := 0; fi13 < 10; fi13++ {
        v1 = ((3.14159 - 3.14159) + v1)
        _ = fi13
    }
    var v14 string = ""
    switch 32768 {
    case 128:
        var v15 uint64 = 4294967295
        _ = v15
    case 127:
        for fi16 := 0; fi16 < 5; fi16++ {
            v3[v14] = v14
            v3["hello"] = v14
            v8 -= 1.5
            _ = fi16
        }
    }
    var v17 float32 = (3.14159 - v1)
    var v18 float32 = float32(v8)
    var v19 int8 = v11
    var v20 int16 = v5
    if v2 {
        switch "\n\t" {
        case "\n\t":
            v12 = 0
        case "hello":
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 float32 = 3.14159
                v9 = 2
                switch 32767 {
                case 0:
                    v18 = (v1 * 0.0)
                case 1:
                    if (!(18446744073709551615 == 18446744073709551615)) {
                        v14 = v14
                        if len(v4) > 0 {
                            v4[0] = 11
                        }
                        if (2 <= 261) {
                            v7 += 7
                            var v23 string = "\n\t"
                            if true {
                                v11 += (v19 ^ (v19 | v19))
                                var v24 float64 = v8
                                v12 = (v12 | (-(v12)))
                                _ = v24
                            }
                            _ = v23
                        }
                    } else {
                        v25 := GenS1{F0: v7}
                        for fi26 := 0; fi26 < 14; fi26++ {
                            var v27 string = v14
                            var v28 int64 = (^(v12))
                            fmt.Println(v18, v1, v19)
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        v19 -= v19
                        _ = v25
                    }
                }
                _ = fi21
                _ = v22
            }
        default:
            if (-374 > -9223372036854775808) {
                v9 = (v7 - v9)
                v29 := GenS2{F0: (v7 & v7), F1: 563}
                _ = len(v3)
                _ = v29
            }
        }
        v18 = 1e10
    }
    fmt.Println(v11, v17)
    v14 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
