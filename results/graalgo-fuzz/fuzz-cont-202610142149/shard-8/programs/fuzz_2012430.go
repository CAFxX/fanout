package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint }

func main() {
    var v1 uint32 = 2
    var v2 int32 = 0
    var v3 bool = false
    var v4 float32 = (float32((3.14159 * 0.0)) * (1e10 - -2.25))
    var v5 uint32 = v1
    v6 := make(map[uint32]bool)
    v5 = v1
    var v7 int = -737
    if v3 {
        v5 = v5
        v8 := []int{v7, 42, v7, v7}
        _ = v8
    } else {
        var v9 uint32 = (v5 | 65535)
        v10 := GenS1{F0: false}
        _ = v10
        _ = v9
    }
    var v11 uint8 = 2
    var v12 int64 = 311
    v4 = (1e10 - 1e10)
    fmt.Println(v11, v3)
    delete(v6, (v5 * (v1 << ((104) & 63))))
    if v3 {
        switch "hello" {
        case "":
            v7 = 32768
        case "\n\t":
            if ((float32(0.0) == ((-2.25 + v4) * 3.14159)) || true) {
                var v13 int = (v7 % ((895) | 1))
                _ = v13
            } else {
                switch (v7 + v7) {
                case 127:
                    switch v7 {
                    case -1:
                        var v14 int = (^(v7))
                        _ = v14
                    case 32767:
                        var v15 float32 = (0.0 / (v4 - (3.14159 - v4)))
                        _ = v15
                    }
                case -617:
                    v16 := []int{v7, 42, (^(v7))}
                    _ = v16
                case -819:
                    var v17 int8 = -115
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 2; fi18++ {
                        var v19 uint16 = 2
                        v20 := GenS1{F0: v3}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
                if v3 {
                    switch v7 {
                    case 4294967295:
                        v21 := GenS2{F0: v12, F1: 1}
                        _ = v21
                    default:
                        var v22 int64 = 9223372036854775807
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 18; fi23++ {
                        var v24 uint = 255
                        _ = fi23
                        _ = v24
                    }
                    v5 = v1
                }
            }
        }
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        v4 -= (3.14159 + v4)
        _ = fi25
    }
    fmt.Println(v7)
    for fi26 := 0; fi26 < 19; fi26++ {
        var v27 float32 = (-2.25 / v4)
        var v28 int = fi26
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := v6[v5]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
