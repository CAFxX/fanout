package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint }
type GenS2 struct { F0 uint8; F1 int16 }

func main() {
    var v1 uint = 891
    v2 := []uint32{590, 42, 65535, 4294967295}
    v3 := make(map[string]bool)
    var v4 float64 = (3.14159 + 1.5)
    var v5 uint = 255
    var v6 float32 = (0.0 - 0.0)
    var v7 uint = v5
    var v8 int = 128
    switch v8 {
    case -9223372036854775808:
        for fi9 := 0; fi9 < 11; fi9++ {
            switch (v8 & (v8 | 32767)) {
            case 32767:
                for fi10 := 0; fi10 < 6; fi10++ {
                    if true {
                        if (1 < 1) {
                            v8 -= fi10
                            v11 := GenS1{F0: 556, F1: 65535, F2: v7}
                            switch v8 {
                            case -881:
                                if (!false) {
                                    var v12 uint16 = 2
                                    _ = v12
                                }
                            default:
                                var v13 uint64 = 0
                                _ = v13
                            }
                            _ = v11
                        }
                        v1 += (^(v1))
                    }
                    _ = fi10
                }
            }
            for fi14 := 0; fi14 < 11; fi14++ {
                v2 = append(v2, 4294967295)
                v15 := GenS2{F0: 2, F1: -993}
                var v16 uint = (v5 >> (((v5 << ((v7) & 63))) & 63))
                v17 := GenS2{F0: 127, F1: -507}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            v18 := GenS1{F0: 0, F1: 42, F2: 922}
            _ = fi9
            _ = v18
        }
    case 42:
        var v19 uint32 = 2147483648
        _ = v19
    default:
        fmt.Println(v7, v4, v6)
    }
    var v20 float64 = v4
    switch (^(v8)) {
    case 9223372036854775807:
        fmt.Println(v4, v5, v20)
    case 540:
        v5 = (v1 - (v1 % ((18446744073709551615) | 1)))
    default:
        v21 := GenS2{F0: 79, F1: 0}
        _ = v21
    }
    v3["hello"] = (42 != 18)
    var v22 uint8 = 31
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
