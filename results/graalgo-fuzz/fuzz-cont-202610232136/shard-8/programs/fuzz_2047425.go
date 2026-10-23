package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 string }
func genFn0(p0 uint16, p1 int16) uint8 {
    var v6 uint16 = p0
    for fi7 := 0; fi7 < 2; fi7++ {
        p1 -= (p1 & p1)
        _ = fi7
    }
    var v8 uint64 = 364
    v8 = v8
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 128
}
func genFn1() uint16 {
    var v9 uint16 = 65535
    v9 += 453
    if false {
        v9 = v9
    }
    _ = v9
    return v9
}

func main() {
    var v1 uint64 = 42
    var v2 string = "foo bar"
    var v3 int64 = 1
    v4 := []string{(v2 + v2)}
    var v5 uint8 = 1
    var v10 string = ((("foo bar" + ((((v2 + v2) + v2) + ((v2 + (v2 + v2)) + (v2 + v2))) + v2)) + "") + v2)
    var v11 float32 = 1e10
    var v12 float32 = v11
    var v13 uint32 = 2
    if (uint32(v5) <= 65535) {
        if true {
            switch 128 {
            case 42:
                v14 := GenS1{F0: 100, F1: -2147483648}
                _ = v14
            }
            fmt.Println(v12, v2, v5)
        }
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS2{F0: v10}
            v17 := GenS2{F0: (v10 + "hello")}
            var v18 bool = (((-2.25 - -2.25) < 0.1) && true)
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    }
    switch 2147483647 {
    case 32768:
        fmt.Println(v2, v11)
    default:
        if true {
            if true {
                var v19 uint = 611
                _ = len(v4)
                _ = v19
            } else {
                fmt.Println(v1, v13, v13)
            }
            switch 32768 {
            case 32768:
                fmt.Println(v11, v2)
            }
            var v20 int16 = 1
            _ = v20
        }
    }
    fmt.Println(v12, v11)
    switch 4294967295 {
    case 290:
        var v21 int16 = -421
        _ = v21
    default:
        if (((v3 / (((v3 ^ (v3 | v3))) | 1)) < v3) && false) {
            var v22 uint = 255
            fmt.Println(v11)
            v22 = v22
            _ = v22
        }
    }
    v10 = v2
    var v23 int64 = (v3 | 1)
    v11 = 0.0
    v24 := GenS1{F0: -55, F1: 344}
    v25 := GenS2{F0: (v10 + v10)}
    v10 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
