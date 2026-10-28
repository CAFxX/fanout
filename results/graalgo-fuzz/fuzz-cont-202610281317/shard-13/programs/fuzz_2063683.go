package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 383
    v4 := GenS1{F0: 0, F1: 65535, F2: -32768}
    if true {
        switch "hello" {
        case "hello":
            v3 = v3
        case "a":
            var v5 uint8 = 2
            _ = v5
        }
        v1[0] = 0
    }
    var v6 int32 = -2147483648
    v7 := GenS1{F0: -1, F1: 42, F2: 42}
    _ = len(v2)
    var v8 int = 32768
    v3 -= (v3 << ((42) & 63))
    v9 := GenS2{F0: 0}
    switch v8 {
    case 4294967295:
        if (2 > 698) {
            fmt.Println(v6, v3)
            if false {
                v8 -= -232
                var v10 uint8 = 162
                var v11 int16 = -32768
                _ = v10
                _ = v11
            } else {
                if true {
                    var v12 int = 9223372036854775807
                    if false {
                        var v13 uint8 = 255
                        var v14 int64 = 9223372036854775807
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint64 = 42
                        _ = v15
                    }
                    v16 := []int{v12, (-(v8))}
                    _ = v12
                    _ = v16
                }
            }
            v6 = v6
        }
    case -733:
        v17 := v1[743]
        _ = v17
    case -1:
        v3 -= v3
    default:
        v3 = v3
    }
    v18 := GenS1{F0: 0, F1: 65535, F2: -759}
    var v19 float64 = 1.5
    for fi20 := 0; fi20 < 8; fi20++ {
        fmt.Println(v19, v19)
        if true {
            v19 = float64(0.1)
            var v21 int64 = 0
            var v22 float64 = v19
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            v8 = 939
            v19 -= v19
            _ = fi23
        }
        _ = fi20
    }
    var v24 int64 = -1
    switch 1 {
    case -820:
        v25 := GenS3{F0: 1, F1: 2147483647, F2: 128}
        _ = v25
    case -1:
        v26 := GenS2{F0: 186}
        _ = v26
    case 906:
        v27 := GenS2{F0: 110}
        _ = v27
    default:
        v3 = 0
    }
    if ((v8 - (v8 ^ 9223372036854775807)) < (v8 + (v8 - (-(v8))))) {
        v28 := GenS1{F0: (v24 << ((42) & 63)), F1: 77, F2: 1000}
        var v29 uint16 = 255
        _ = v28
        _ = v29
    } else {
        var v30 uint32 = 18
        if false {
            var v31 float64 = v19
            _ = len(v1)
            _ = v31
        } else {
            var v32 uint64 = 18446744073709551615
            _ = v32
        }
        delete(v2, (^(v30)))
        _ = v30
    }
    var v33 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
