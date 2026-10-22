package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int32; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint16 }
func genFn0() uint8 {
    if false {
        var v5 uint16 = 409
        var v6 int = 4294967295
        v7 := []int{(v6 ^ 190), 32768, v6, v6}
        _ = v5
        _ = v6
        _ = v7
    } else {
        var v8 int64 = -764
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: 144, F1: -770}
            v8 -= v8
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    v11 := []int{641}
    v12 := GenS2{F0: 82, F1: 765, F2: -136}
    var v13 int8 = 100
    _ = v11
    _ = v12
    _ = v13
    return 2
}

func main() {
    var v1 uint = 573
    v2 := make(map[uint32]string)
    var v3 float32 = float32(3.14159)
    var v4 uint = v1
    switch 128 {
    case 42:
        v4 = (v1 * 703)
    case 128:
        fmt.Println(v1, v1)
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint32 = 462
        if (-313 > 42) {
            v15 = (^(v15))
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v1)
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 int32 = 2147483647
        _ = fi16
        _ = v17
    }
    fmt.Println(v3, v1)
    v1 -= (^(v4))
    fmt.Println(v1, v3)
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 int = (fi18 * 128)
        var v20 int8 = -128
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := v2[540]
    v1 -= (v1 & (v4 + (v1 >> (((v1 ^ v1)) & 63))))
    var v22 float32 = -2.25
    v21 = "a"
    v21 = v21
    if false {
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 int32 = -238
            v25 := GenS1{F0: 255, F1: (v24 % ((v24) | 1))}
            var v26 int16 = 0
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 uint16 = 0
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 7; fi29++ {
            v30 := []int{}
            v31 := GenS2{F0: 1, F1: 42, F2: 1}
            _ = fi29
            _ = v30
            _ = v31
        }
    } else {
        v21 = v21
        v32 := []int{42}
        fmt.Println(v4, v21, v3)
        _ = v32
    }
    if true {
        var v33 int64 = 9223372036854775807
        v34 := GenS3{F0: 32767}
        _ = v33
        _ = v34
    }
    if false {
        v22 = 1e10
    } else {
        switch 128 {
        case 712:
            fmt.Println(v4, v21)
        case 662:
            v4 = v4
        default:
            var v35 int8 = 0
            _ = v35
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 uint64 = 809
            _ = fi36
            _ = v37
        }
    }
    v4 -= (v1 << (((v1 & v1)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
