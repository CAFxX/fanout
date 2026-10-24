package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 uint }
func genFn0() uint8 {
    v4 := GenS1{F0: 42, F1: 128}
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int8 = 45
        _ = fi5
        _ = v6
    }
    if ((79 == 1) || true) {
        var v7 int64 = 972
        v8 := []int{-1}
        _ = v7
        _ = v8
    }
    var v9 uint32 = 42
    v10 := []int{42, 4294967295, -227, 1}
    _ = v10
    _ = v4
    _ = v9
    return 42
}

func main() {
    var v1 int32 = -1
    var v2 uint16 = 1
    v3 := []int{32767, -335, 4294967295, 127}
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int = 2147483647
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: -1, F1: 255}
    if true {
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 uint = 65535
            v16 := GenS1{F0: 997, F1: 255}
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 int = 454
        fmt.Println(v1)
        _ = v17
    }
    v18 := GenS2{F0: 124, F1: 169}
    var v19 uint32 = 65535
    if false {
        v1 = 216
        _ = len(v3)
    }
    var v20 float32 = float32(0.0)
    var v21 int = -1
    v22 := GenS2{F0: 127, F1: 128}
    if (!false) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        _ = len(v3)
        var v23 uint8 = 0
        v21 = (v21 ^ (v21 << ((376) & 63)))
        _ = v23
    }
    v20 = v20
    v21 -= v21
    if true {
        switch (v21 % ((-2147483648) | 1)) {
        case -2147483648:
            var v24 float32 = v20
            _ = v24
        default:
            if ("hello" > "a") {
                if true {
                    var v25 int = 398
                    var v26 int64 = 1
                    if true {
                        if (32767 != -1) {
                            v27 := GenS3{F0: v26}
                            _ = v27
                        } else {
                            var v28 uint = uint(v2)
                            v29 := GenS3{F0: (v26 % (((v26 >> ((v28) & 63))) | 1))}
                            _ = v28
                            _ = v29
                        }
                        v3 = append(v3, v21)
                        fmt.Println(v1, v19)
                    }
                    _ = v25
                    _ = v26
                } else {
                    v1 -= 2147483647
                    fmt.Println(v1, v1)
                }
                v30 := GenS4{F0: true, F1: 18446744073709551615}
                _ = v30
            } else {
                v2 -= 42
                fmt.Println(v20, v21)
                if len(v3) > 0 {
                    v3[0] = (v21 % ((127) | 1))
                }
            }
        }
        var v31 uint16 = 65535
        _ = v31
    } else {
        _ = len(v3)
    }
    v19 -= v19
    var v32 uint16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
}
