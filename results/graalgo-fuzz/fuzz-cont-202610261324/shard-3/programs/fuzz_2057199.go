package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 uint64 = 2
    var v2 int32 = 1
    var v3 float64 = 1.5
    var v4 int = -54
    var v5 uint64 = v1
    var v6 uint16 = 1
    if true {
        var v7 uint32 = 65535
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(v5, v2)
        _ = fi8
    }
    if (int32(v5) >= -2147483648) {
        v9 := []int{761, -2147483648, (v4 + 1)}
        var v10 uint64 = (v1 ^ (v5 >> ((18446744073709551615) & 63)))
        v4 = v4
        _ = v10
        _ = v9
    } else {
        fmt.Println(v6, v2, v6)
        if true {
            fmt.Println(v1)
        } else {
            var v11 uint64 = v1
            v3 -= v3
            var v12 string = "foo bar"
            _ = v11
            _ = v12
        }
    }
    v13 := GenS1{F0: 42}
    if false {
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := []int{(^(fi14)), (fi14 * fi14)}
            _ = fi14
            _ = v15
        }
        v16 := GenS2{F0: v4}
        _ = v16
    }
    fmt.Println(v1)
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int16 = 32767
        var v19 int64 = 593
        v20 := []int{(-(v4)), (^(v4))}
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    switch (v4 / (((v4 * (v4 + v4))) | 1)) {
    case -276:
        var v21 int32 = v2
        _ = v21
    case -388:
        v22 := GenS1{F0: 18446744073709551615}
        _ = v22
    case -10:
        if (42 <= 434) {
            var v23 float32 = (float32((v3 * v3)) * 0.1)
            var v24 float64 = float64(v23)
            fmt.Println(v3, v24)
            _ = v23
            _ = v24
        }
    default:
        v25 := GenS3{F0: (v2 / (((-(v2))) | 1))}
        _ = v25
    }
    switch "\n\t" {
    case "":
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 float64 = v3
            _ = fi26
            _ = v27
        }
    default:
        var v28 bool = true
        _ = v28
    }
    var v29 int16 = 221
    if false {
        var v30 uint = 2
        if false {
            var v31 float64 = v3
            for fi32 := 0; fi32 < 20; fi32++ {
                v33 := GenS4{F0: false, F1: 4294967295}
                v5 = 18446744073709551615
                v2 = (v2 + v2)
                _ = fi32
                _ = v33
            }
            _ = v31
        } else {
            if false {
                v5 = 42
                var v34 uint16 = uint16(v30)
                _ = v34
            } else {
                v35 := []int{}
                v5 += (v1 % ((18446744073709551615) | 1))
                v36 := GenS1{F0: (v30 - 727)}
                _ = v35
                _ = v36
            }
            var v37 uint8 = 234
            _ = v37
        }
        var v38 float32 = float32(v3)
        _ = v30
        _ = v38
    }
    var v39 bool = ((v5 & (^(v1))) == 2)
    for fi40 := 0; fi40 < 7; fi40++ {
        for fi41 := 0; fi41 < 12; fi41++ {
            v42 := GenS1{F0: uint(v6)}
            v43 := GenS5{F0: (v5 << ((1) & 63))}
            v44 := []int{(^(v4)), (fi40 - (v4 + 56)), (v4 | (fi41 / (((fi41 ^ v4)) | 1))), fi41}
            _ = fi41
            _ = v42
            _ = v43
            _ = v44
        }
        _ = fi40
    }
    var v45 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
