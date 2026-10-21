package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 uint32; F2 int }
func genFn0() bool {
    var v5 uint8 = 243
    v6 := GenS1{F0: 0}
    v7 := GenS2{F0: 33, F1: true}
    v8 := GenS3{F0: 54, F1: 1}
    for fi9 := 0; fi9 < 9; fi9++ {
        if (3.14159 < 0.0) {
            v5 -= (v5 << ((18446744073709551615) & 63))
        }
        if (!false) {
            v5 -= 42
            var v10 int32 = 986
            _ = v10
        } else {
            switch "" {
            case "hello":
                var v11 float32 = 0.1
                _ = v11
            }
            v5 += 236
        }
        _ = fi9
    }
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 uint64 = 4294967295
    var v2 string = "hello"
    v3 := make(map[uint32]bool)
    var v4 int16 = -1
    var v12 uint64 = (v1 >> ((4294967295) & 63))
    if false {
        var v13 int8 = 1
        v13 = -128
        _ = v13
    }
    v14 := v3[2147483647]
    v15 := GenS3{F0: 2, F1: -1}
    fmt.Println(v2, v1, v14)
    var v16 uint = 2
    v3[337] = (v16 < (^(v16)))
    if (65535 == v16) {
        v17 := GenS4{F0: v4}
        _ = v17
    } else {
        v18 := []int{595}
        var v19 int8 = 42
        if v14 {
            v4 = int16(v1)
            v14 = false
        }
        _ = v18
        _ = v19
    }
    v20 := GenS5{F0: 2, F1: 1, F2: 32767}
    if true {
        fmt.Println(v4, v4, v4)
        v21 := GenS2{F0: (v16 | v16), F1: v14}
        _ = v21
    }
    switch -2147483648 {
    case -450:
        v22 := GenS2{F0: (^(v16)), F1: v14}
        _ = v22
    case -1:
        fmt.Println(v4, v16, v14)
    case 32767:
        switch v2 {
        case "":
            v1 -= (v12 | (^(v1)))
        }
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 int16 = v4
        v25 := GenS2{F0: v16, F1: (0 == 203)}
        v26 := GenS3{F0: 335, F1: -1}
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
