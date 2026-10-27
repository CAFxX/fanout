package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0() uint {
    v5 := []int{9223372036854775807, 1}
    if false {
        var v6 uint = 1
        _ = len(v5)
        v6 = 862
        _ = v6
    }
    v7 := GenS1{F0: 65535}
    v8 := GenS1{F0: 289}
    _ = v5
    _ = v7
    _ = v8
    return 507
}
func genFn1() int {
    var v9 int16 = -1
    fmt.Println(v9)
    _ = v9
    return 32768
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = false
    v3 := []int{42, 285, 540, -822}
    v4 := []int{-2147483648, -2147483648, -964}
    for fi10 := 0; fi10 < 5; fi10++ {
        if v2 {
            v2 = (-1 <= 42)
        } else {
            v2 = (9223372036854775807 <= 0)
            v2 = (14 > 42)
            if (65535 > 65535) {
                fmt.Println(fi10)
            }
        }
        v11 := GenS1{F0: 0}
        var v12 uint32 = 2147483647
        _ = fi10
        _ = v11
        _ = v12
    }
    v1[2147483647] = v2
    fmt.Println(v2)
    v2 = v2
    v13 := GenS1{F0: 255}
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int64 = 42
        var v16 uint = 0
        var v17 uint64 = 615
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    fmt.Println(v2)
    switch 32767 {
    case -2147483648:
        if v2 {
            if (-128 < 39) {
                switch 2147483647 {
                case -341:
                    v2 = (0 >= 0)
                default:
                    v1[553] = (255 >= 33)
                }
                if v2 {
                    v1[42] = false
                    var v18 uint64 = 18446744073709551615
                    _ = v18
                }
            }
            if (!v2) {
                v19 := GenS1{F0: 920}
                _ = v19
            } else {
                var v20 int16 = 42
                _ = v20
            }
        } else {
            v21 := GenS1{F0: 0}
            var v22 int8 = -14
            _ = v21
            _ = v22
        }
    case -1:
        fmt.Println(v2)
    }
    fmt.Println(v2)
    v23 := GenS2{F0: -411, F1: 127}
    v24 := GenS3{F0: 458}
    var v25 uint32 = 2147483648
    var v26 int32 = 0
    var v27 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
