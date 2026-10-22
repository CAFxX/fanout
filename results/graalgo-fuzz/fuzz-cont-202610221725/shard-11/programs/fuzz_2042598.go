package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int16; F2 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int }
func genFn0(p0 uint64) uint16 {
    v7 := []int{}
    var v8 uint32 = 0
    v9 := GenS1{F0: int(v8), F1: 339}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 993
}
func genFn1() uint8 {
    v10 := GenS2{F0: -46}
    if (!true) {
        switch "hello" {
        case "a":
            v11 := []int{}
            _ = v11
        }
        var v12 uint32 = 320
        var v13 float32 = float32(1e10)
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: 32768, F1: -2147483648}
        _ = v14
    }
    _ = v10
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 uint16 = 32767
    v4 := make(map[uint32]uint32)
    var v5 string = "a"
    v6 := []string{}
    var v15 uint32 = 4294967295
    v16 := GenS2{F0: 1}
    var v17 int8 = -2
    fmt.Println(v15, v17, v17)
    v17 = v17
    v18 := GenS3{F0: 1}
    switch 42 {
    case 127:
        var v19 bool = ((42 != -32768) || false)
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := v1[42]
        var v22 uint64 = 4294967295
        var v23 uint8 = 239
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 int32 = 2147483647
    v25 := GenS3{F0: genFn0(938)}
    if true {
        fmt.Println(v17, v3, v15)
        switch 584 {
        case -2147483648:
            fmt.Println(v15, v5, v15)
        }
        if false {
            switch 42 {
            case -1:
                var v26 uint32 = v15
                _ = v26
            default:
                if ((1.5 / 3.14159) > 0.1) {
                    _ = len(v4)
                    if len(v6) > 0 {
                        v6[0] = v5
                    }
                }
            }
            var v27 uint = 255
            var v28 string = ""
            _ = v27
            _ = v28
        }
    }
    v29 := GenS2{F0: v17}
    v30 := GenS4{F0: 1000, F1: 42, F2: 4294967295}
    v31 := GenS5{F0: 65535}
    v32 := GenS5{F0: 1}
    _ = len(v2)
    v33 := GenS6{F0: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
