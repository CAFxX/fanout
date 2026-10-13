package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int64; F1 int16; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint16; F2 string }
type GenS6 struct { F0 uint }
func genFn0() uint8 {
    switch 275 {
    case 178:
        v7 := GenS1{F0: 32767, F1: 32, F2: -357}
        _ = v7
    case 32767:
        var v8 int32 = 2147483647
        _ = v8
    }
    var v9 string = ""
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            if (-759 == 0) {
                v9 = v9
            }
            fmt.Println(fi11)
            _ = fi11
        }
        _ = fi10
    }
    switch 32767 {
    case 42:
        switch 2147483647 {
        case -2147483648:
            fmt.Println(v9)
        default:
            var v12 string = "\n\t"
            _ = v12
        }
    case 175:
        v13 := GenS1{F0: 311, F1: 65535, F2: 1}
        _ = v13
    default:
        v14 := GenS2{F0: 181, F1: 1}
        _ = v14
    }
    _ = v9
    return 199
}

func main() {
    var v1 float64 = float64(1.5)
    var v2 int8 = -118
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    var v5 int8 = (v2 ^ (v2 - (-(v2))))
    var v6 int8 = v5
    v15 := GenS3{F0: 1, F1: -1, F2: -1}
    if (false || true) {
        switch -1 {
        case 32767:
            if (false && true) {
                v3["hello"] = (948 <= 0)
            }
        default:
            v16 := []int{}
            _ = v16
        }
        if true {
            fmt.Println(v6, v5)
            fmt.Println(v1, v5, v5)
        } else {
            var v17 float32 = float32(-2.25)
            _ = v17
        }
    }
    _ = len(v3)
    fmt.Println(v6, v6)
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := GenS3{F0: -101, F1: -1, F2: -9223372036854775808}
        fmt.Println(fi18)
        var v20 float32 = 1.5
        _ = fi18
        _ = v19
        _ = v20
    }
    v1 -= v1
    var v21 uint32 = 42
    var v22 int32 = -710
    v23 := GenS4{F0: 80, F1: v21}
    for fi24 := 0; fi24 < 5; fi24++ {
        switch (^(fi24)) {
        case 412:
            var v25 int16 = 1
            _ = v25
        case 32767:
            for fi26 := 0; fi26 < 4; fi26++ {
                v27 := GenS5{F0: v21, F1: 4, F2: "hello"}
                v28 := GenS6{F0: uint(v21)}
                v5 += (v6 | v6)
                _ = fi26
                _ = v27
                _ = v28
            }
        default:
            var v29 int = fi24
            _ = v29
        }
        _ = fi24
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 int64 = -584
            v33 := GenS3{F0: v32, F1: 1000, F2: (v32 | v32)}
            var v34 uint = 556
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        _ = fi30
    }
    v35 := []int{-850, 1, 42}
    fmt.Println(v22)
    v22 -= v22
    var v36 uint32 = (v21 | (v21 ^ (^(v21))))
    v5 += (v6 % ((v6) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
