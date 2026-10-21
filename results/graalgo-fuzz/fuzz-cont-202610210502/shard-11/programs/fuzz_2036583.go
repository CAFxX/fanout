package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 string; F2 uint32 }
func genFn0(p0 int64, p1 uint) uint16 {
    var v5 uint = (p1 & 255)
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 string = "hello"
        switch -9223372036854775808 {
        case -807:
            var v8 uint64 = 1
            _ = v8
        case 184:
            var v9 float64 = (1.5 * float64(1.5))
            _ = v9
        default:
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 int16 = 1
                _ = fi10
                _ = v11
            }
        }
        _ = fi6
        _ = v7
    }
    v12 := []int{-19}
    var v13 int32 = 0
    var v14 int = int(p0)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v5
    return 0
}

func main() {
    var v1 uint32 = 594
    var v2 int = 455
    var v3 int16 = 1
    v4 := []int{int(v3), v2}
    if (!(65535 == v1)) {
        var v15 uint8 = 255
        var v16 int = (^(v2))
        var v17 uint8 = 129
        _ = v15
        _ = v16
        _ = v17
    }
    v1 = (v1 / ((65535) | 1))
    var v18 uint64 = 42
    switch "\n\t" {
    case "\n\t":
        v19 := GenS1{F0: 183}
        _ = v19
    default:
        var v20 int8 = 100
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 bool = true
        var v23 int8 = 100
        _ = fi21
        _ = v22
        _ = v23
    }
    v3 = v3
    _ = len(v4)
    var v24 int16 = (v3 | 1)
    var v25 int32 = -2147483648
    fmt.Println(v25, v25, v18)
    var v26 uint16 = 980
    if false {
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 float32 = 0.1
            fmt.Println(v26)
            if true {
                v29 := GenS2{F0: v3, F1: "a", F2: (v1 % ((v1) | 1))}
                for fi30 := 0; fi30 < 8; fi30++ {
                    var v31 float64 = 1.5
                    fmt.Println(v25)
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
            _ = fi27
            _ = v28
        }
    }
    var v32 uint = 1
    v4 = append(v4, v2)
    switch (v2 ^ (v2 & v2)) {
    case -2147483648:
        v33 := GenS2{F0: v3, F1: "hello", F2: (v1 >> (((v32 * (v32 * v32))) & 63))}
        _ = v33
    }
    if (true && (!(87 < 81))) {
        if (104 < 127) {
            v1 = v1
            v2 += v2
            v34 := GenS2{F0: 1000, F1: ((("\n\t" + (("a" + "foo bar") + (("a" + "a") + "a"))) + "hello") + "hello"), F2: v1}
            _ = v34
        }
    }
    v26 -= v26
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
}
