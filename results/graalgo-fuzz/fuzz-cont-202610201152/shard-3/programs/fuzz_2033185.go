package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 bool }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }

func main() {
    var v1 float32 = -2.25
    var v2 uint8 = 0
    v3 := []uint32{2, 496}
    var v4 float64 = (1e10 - (0.1 / -2.25))
    var v5 int64 = -978
    var v6 int64 = 868
    if false {
        var v7 uint8 = (v2 % (((^(v2))) | 1))
        var v8 int64 = v6
        v9 := GenS1{F0: 127, F1: 615, F2: ((v2 - (v2 + v2)) < v7)}
        _ = v7
        _ = v8
        _ = v9
    }
    if len(v3) > 0 {
        v3[0] = 518
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int32 = 42
        v1 += v1
        if len(v3) > 0 {
            v3[0] = 42
        }
        _ = fi10
        _ = v11
    }
    if (v5 == v6) {
        var v12 bool = true
        v13 := GenS1{F0: 6, F1: v5, F2: v12}
        _ = v12
        _ = v13
    }
    var v14 uint64 = 1
    for fi15 := 0; fi15 < 18; fi15++ {
        fmt.Println(v14)
        _ = fi15
    }
    v16 := GenS1{F0: 6, F1: (-(v5)), F2: false}
    if (3.14159 < v4) {
        v6 = v6
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        if (0 == (v2 * 42)) {
            var v17 uint32 = 465
            _ = v17
        } else {
            v5 -= (v6 * v6)
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            fmt.Println(v2)
            _ = fi18
        }
    }
    if (-9223372036854775808 == (v5 * (-(v6)))) {
        switch 4294967295 {
        case 175:
            fmt.Println(v2)
        case -9223372036854775808:
            var v19 int8 = 42
            _ = v19
        }
        switch -483 {
        case 42:
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 float64 = v4
                _ = fi20
                _ = v21
            }
        default:
            if false {
                fmt.Println(v2, v1)
            }
        }
    } else {
        v22 := GenS1{F0: 103, F1: int64(v14), F2: true}
        v3 = append(v3, 2147483647)
        switch "" {
        case "":
            v23 := GenS1{F0: 100, F1: 1, F2: false}
            _ = v23
        default:
            var v24 string = (((("a" + "\n\t") + ("hello" + "a")) + "a") + "hello")
            _ = v24
        }
        _ = v22
    }
    v6 -= (v6 % (((v6 << ((65535) & 63))) | 1))
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 uint = 255
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        fmt.Println(v5, v2, v4)
        v6 = 42
        _ = fi27
    }
    v28 := GenS2{F0: v6, F1: v14, F2: 42}
    var v29 int64 = v5
    for fi30 := 0; fi30 < 18; fi30++ {
        v4 += (1.5 / (-2.25 + 0.1))
        _ = fi30
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    fmt.Println(v1, v29, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
