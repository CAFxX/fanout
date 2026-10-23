package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int32 }
func genFn0(p0 int16, p1 int8) bool {
    v4 := []int{-9223372036854775808, 9223372036854775807}
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    _ = p0
    _ = p1
    _ = v4
    return (523 <= 758)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "a"
    var v3 bool = (!false)
    if (42 > 1) {
        if true {
            if ((1 <= 42) && false) {
                v5 := GenS1{F0: 65535}
                var v6 int32 = 1
                _ = v5
                _ = v6
            } else {
                for fi7 := 0; fi7 < 3; fi7++ {
                    v8 := GenS1{F0: 255}
                    _ = fi7
                    _ = v8
                }
            }
            var v9 int64 = 1
            _ = v9
        }
        v2 = v2
        v2 = (v2 + (v2 + (v2 + v2)))
    }
    var v10 float32 = 1.5
    v10 = 1e10
    if ((128 != 4294967295) && false) {
        var v11 uint = 520
        if (0.0 < 1.5) {
            fmt.Println(v3, v2)
            v10 = (1e10 * float32(-2.25))
            v12 := GenS2{F0: 2, F1: 813, F2: -938}
            _ = v12
        }
        v13 := GenS2{F0: 42, F1: 27, F2: -2147483648}
        _ = v11
        _ = v13
    } else {
        v3 = v3
    }
    v14 := GenS2{F0: 824, F1: 2, F2: 69}
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := []int{(fi15 + -316), 42, fi16}
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := []int{(fi18 >> ((40) & 63)), 32767}
            var v20 int8 = -1
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi15
    }
    switch 973 {
    case -507:
        v21 := GenS2{F0: 2, F1: 429, F2: 0}
        _ = v21
    case 351:
        v10 = (v10 / v10)
    case -673:
        if ((v3 && true) || (!((645 == 1) || v3))) {
            _ = len(v1)
            var v22 int16 = -32768
            _ = v22
        }
    }
    v2 = ""
    fmt.Println(v10)
    for fi23 := 0; fi23 < 5; fi23++ {
        fmt.Println(v3, v3, v10)
        _ = fi23
    }
    var v24 bool = (42 >= 251)
    v24 = v3
    var v25 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
