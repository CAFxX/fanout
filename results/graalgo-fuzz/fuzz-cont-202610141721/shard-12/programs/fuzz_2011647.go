package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS2 struct { F0 int8; F1 bool }
func genFn0(p0 int8, p1 string, p2 int32) int {
    v4 := []int{}
    p2 += p2
    v5 := GenS1{F0: 42, F1: 127, F2: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 128
}

func main() {
    var v1 float32 = (((-2.25 / 1.0) * -2.25) * float32(1e10))
    v2 := make(map[uint32]int)
    var v3 string = ""
    var v6 int16 = -32768
    v7 := GenS1{F0: 680, F1: 255, F2: 790}
    if false {
        var v8 int32 = int32(v6)
        switch -1 {
        case -69:
            var v9 float64 = 1e10
            _ = v9
        }
        _ = v8
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            fmt.Println(v6)
            _ = fi10
        }
    }
    if (((((v6 * (v6 | -476)) != v6) || ((!(80 >= -1)) || false)) && ((false && true) && false)) && (1 < 0)) {
        v3 = ""
        var v11 int32 = 382
        _ = v11
    } else {
        if true {
            v6 -= v6
            if true {
                fmt.Println(v1, v1)
                v3 = ""
                v6 += v6
            }
            v1 = v1
        }
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        if (!(!(!true))) {
            v3 = v3
        }
        _ = fi12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 float64 = (((float64((v1 * v1)) * 1e10) + 1e10) * float64(v1))
        v6 = (v6 / (((v6 | (^(v6)))) | 1))
        _ = fi13
        _ = v14
    }
    var v15 uint32 = 1
    v16 := GenS1{F0: 722, F1: 1, F2: 170}
    switch v3 {
    case "a":
        var v17 int32 = -1
        _ = v17
    case "hello":
        for fi18 := 0; fi18 < 11; fi18++ {
            v15 = (^(v15))
            _ = fi18
        }
    case "\n\t":
        var v19 int64 = -354
        _ = v19
    default:
        fmt.Println(v3, v6)
    }
    v6 = 1000
    if (true || true) {
        v20 := GenS1{F0: 2, F1: 2, F2: -9223372036854775808}
        var v21 float64 = 1.5
        _ = v20
        _ = v21
    }
    switch 2147483647 {
    case 32768:
        if (1 > 2) {
            v22 := GenS1{F0: 2, F1: 1, F2: -9223372036854775808}
            _ = v22
        }
    case 32767:
        fmt.Println(v3, v3, v15)
    }
    v23 := GenS2{F0: 100, F1: (v6 <= v6)}
    fmt.Println(v1, v1, v1)
    v2[2147483648] = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v6
    _ = v7
}
