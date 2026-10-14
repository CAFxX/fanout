package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint64; F1 uint8; F2 string }
func genFn0() uint8 {
    var v5 float64 = 0.0
    v6 := GenS1{F0: 42}
    v7 := GenS2{F0: 42, F1: 42}
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint16 = 0
        fmt.Println(fi8)
        var v10 uint64 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1() bool {
    var v11 float64 = 0.1
    v12 := []int{9223372036854775807}
    if (((3.14159 >= 0.0) || (true || true)) || true) {
        v11 = v11
    } else {
        var v13 int = 9223372036854775807
        v13 = (v13 * 0)
        _ = v13
    }
    v11 = (v11 / 1.5)
    _ = v11
    _ = v12
    return (!true)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = false
    var v3 int = -1
    var v4 int8 = 127
    switch -871 {
    case 1:
        if v2 {
            var v14 uint = 1
            var v15 float32 = float32(0.0)
            var v16 int16 = 0
            _ = v14
            _ = v15
            _ = v16
        } else {
            delete(v1, 4294967295)
        }
    case 42:
        v2 = v2
    default:
        v17 := GenS2{F0: (v4 >> ((823) & 63)), F1: 495}
        _ = v17
    }
    v18 := GenS2{F0: (v4 % ((v4) | 1)), F1: 65535}
    var v19 float32 = 1.5
    v20 := GenS1{F0: 393}
    var v21 int32 = -2147483648
    v19 = (1.5 / (3.14159 * float32(1.5)))
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS3{F0: v3, F1: -1}
        var v24 int16 = -32768
        _ = fi22
        _ = v23
        _ = v24
    }
    switch v3 {
    case 128:
        var v25 int16 = -1
        _ = v25
    default:
        if v2 {
            fmt.Println(v2, v19)
            if (65535 != 2) {
                for fi26 := 0; fi26 < 3; fi26++ {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        var v28 int32 = v21
                        var v29 bool = (!(42 < 18446744073709551615))
                        var v30 int32 = 42
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    _ = len(v1)
                    v31 := GenS4{F0: 42, F1: 42, F2: "a"}
                    _ = fi26
                    _ = v31
                }
                var v32 uint16 = 893
                _ = v32
            }
        }
    }
    var v33 float32 = 0.1
    fmt.Println(v3, v33)
    v34 := GenS3{F0: v3, F1: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
