package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int }
type GenS2 struct { F0 int64; F1 int }

func main() {
    var v1 int = 1
    var v2 uint16 = 87
    v3 := []uint32{42, 4294967295}
    v4 := make(map[string]bool)
    v5 := make(map[string]bool)
    v6 := make(map[string]bool)
    v2 = (v2 + v2)
    v7 := GenS1{F0: 255, F1: 0, F2: -2147483648}
    var v8 bool = (!true)
    var v9 float32 = 1.5
    if v8 {
        v10 := GenS1{F0: 1, F1: 42, F2: (v1 >> ((18446744073709551615) & 63))}
        for fi11 := 0; fi11 < 7; fi11++ {
            v9 = float32(3.14159)
            v2 -= (v2 / (((v2 ^ (v2 / ((32767) | 1)))) | 1))
            if (v8 && (("" + "foo bar") < "")) {
                v8 = v8
            } else {
                var v12 float64 = 3.14159
                v9 += (float32(v12) - float32(v12))
                _ = len(v3)
                _ = v12
            }
            _ = fi11
        }
        _ = v10
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS2{F0: -1, F1: (v1 * -986)}
        v8 = (v8 && v8)
        var v15 uint16 = 255
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 float64 = (3.14159 * 0.1)
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int8 = -1
        v9 += (v9 / -2.25)
        _ = fi18
        _ = v19
    }
    if false {
        v20 := GenS2{F0: 42, F1: v1}
        var v21 int32 = -1
        _ = v20
        _ = v21
    } else {
        var v22 int = 755
        _ = v22
    }
    fmt.Println(v9)
    fmt.Println(v2, v9)
    var v23 int64 = -1
    v2 -= (^(v2))
    var v24 int8 = 0
    v4[(("hello" + "") + "\n\t")] = (v8 || (true && false))
    for fi25 := 0; fi25 < 7; fi25++ {
        if (274 >= 1) {
            var v26 uint64 = 42
            fmt.Println(v1, v26)
            if (-2.25 > -2.25) {
                var v27 uint32 = 367
                _ = v27
            } else {
                if v8 {
                    if v8 {
                        v23 -= (^(v23))
                        for fi28 := 0; fi28 < 16; fi28++ {
                            var v29 uint8 = 204
                            v30 := GenS2{F0: (v23 + v23), F1: fi25}
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        v31 := v4["a"]
                        _ = v31
                    }
                } else {
                    v2 = v2
                }
            }
            _ = v26
        }
        v1 += v1
        _ = fi25
    }
    v2 += (v2 / ((255) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
