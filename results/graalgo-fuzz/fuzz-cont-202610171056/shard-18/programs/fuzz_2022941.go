package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int8; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint64) int32 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 uint8 = 127
        v7 := []int{}
        p0 = (p0 * p0)
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{291, 32768}
    p0 -= (p0 ^ 4294967295)
    _ = p0
    _ = v8
    return 42
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint8 = 1
    var v3 int8 = 100
    var v4 uint16 = 211
    if true {
        if (genFn0(v1) != 528) {
            v3 -= int8(v1)
        }
        var v9 uint8 = (v2 - (v2 | (^(v2))))
        _ = v9
    } else {
        v10 := GenS1{F0: false, F1: 42}
        var v11 float64 = ((1.5 - -2.25) / 1e10)
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := GenS1{F0: true, F1: -16}
            _ = fi12
            _ = v13
        }
        _ = v10
        _ = v11
    }
    v1 -= 405
    for fi14 := 0; fi14 < 6; fi14++ {
        if (false || true) {
            if true {
                var v15 float64 = (1.5 / (0.1 + -2.25))
                v2 = uint8(v1)
                _ = v15
            } else {
                var v16 uint = 42
                v4 = 914
                _ = v16
            }
            v17 := GenS2{F0: (^(v4)), F1: 744, F2: 172}
            v18 := []int{fi14}
            _ = v17
            _ = v18
        }
        v3 = (v3 | v3)
        v19 := []int{208, fi14, 253}
        _ = fi14
        _ = v19
    }
    var v20 uint8 = (v2 | 1)
    if false {
        switch "hello" {
        case "":
            v21 := []int{0}
            _ = v21
        default:
            v22 := []int{-2147483648}
            _ = v22
        }
        var v23 int64 = -740
        v24 := GenS3{F0: 2147483647, F1: 2147483648}
        _ = v23
        _ = v24
    }
    fmt.Println(v1)
    fmt.Println(v4)
    var v25 bool = false
    if v25 {
        v25 = (v25 || ((!v25) && (v25 || v25)))
        fmt.Println(v25, v2, v2)
    } else {
        if v25 {
            v26 := GenS4{F0: v3, F1: 65535, F2: (^(v20))}
            var v27 uint16 = 65535
            v4 = v4
            _ = v26
            _ = v27
        } else {
            v28 := GenS5{F0: 4294967295, F1: 42}
            fmt.Println(v1)
            switch 1 {
            case -583:
                v29 := []int{}
                _ = v29
            }
            _ = v28
        }
        v30 := []int{2147483647, -462}
        v25 = ((1 == 4294967295) && ((^(v20)) == (v2 - (v20 >> ((255) & 63)))))
        _ = v30
    }
    switch 2147483647 {
    case 0:
        var v31 uint16 = 603
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
}
