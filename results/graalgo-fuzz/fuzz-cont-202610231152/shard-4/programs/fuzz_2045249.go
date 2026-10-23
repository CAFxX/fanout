package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
func genFn0(p0 int16) int64 {
    fmt.Println(p0)
    var v4 int16 = (p0 - p0)
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]bool)
    var v3 int8 = -128
    for fi5 := 0; fi5 < 1; fi5++ {
        v1 = v1
        _ = fi5
    }
    delete(v2, 2147483647)
    var v6 float64 = 1e10
    if true {
        fmt.Println(v6)
    } else {
        if (-2147483648 == 42) {
            switch 42 {
            case 32767:
                var v7 uint8 = 70
                _ = v7
            case -1:
                for fi8 := 0; fi8 < 5; fi8++ {
                    fmt.Println(fi8, v3)
                    _ = fi8
                }
            case 0:
                var v9 float64 = ((1.5 + 3.14159) - v6)
                _ = v9
            default:
                var v10 string = (("hello" + "hello") + "")
                _ = v10
            }
            for fi11 := 0; fi11 < 8; fi11++ {
                v6 = v6
                _ = fi11
            }
        }
    }
    v6 += float64(float32(float64(0.1)))
    var v12 uint = 0
    if false {
        var v13 uint8 = 100
        _ = v13
    } else {
        switch 32767 {
        case -921:
            fmt.Println(v3, v3)
        }
    }
    fmt.Println(v12, v6)
    if true {
        v12 -= (v12 ^ 0)
    } else {
        fmt.Println(v12, v1, v1)
        for fi14 := 0; fi14 < 2; fi14++ {
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := GenS1{F0: 1, F1: 9223372036854775807}
                _ = fi15
                _ = v16
            }
            v17 := GenS1{F0: 2147483647, F1: -166}
            _ = fi14
            _ = v17
        }
    }
    var v18 int16 = -1
    v19 := GenS1{F0: int32(v18), F1: 666}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
}
