package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint8 }

func main() {
    var v1 bool = (!true)
    var v2 float64 = -2.25
    v3 := make(map[string]uint32)
    var v4 int32 = 1
    var v5 int32 = v4
    v5 = v5
    if true {
        var v6 uint = 941
        var v7 uint = (v6 / ((v6) | 1))
        _ = v6
        _ = v7
    } else {
        v2 += v2
        v8 := []int{-1, 9223372036854775807, 9223372036854775807}
        v9 := GenS1{F0: 127, F1: 255, F2: 2147483648}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 bool = (0.0 > (float32(v2) * (1e10 / float32(((v2 / (v2 * v2)) - v2)))))
        fmt.Println(v4)
        _ = fi10
        _ = v11
    }
    if v1 {
        var v12 int64 = -9223372036854775808
        if v1 {
            if (1 <= v4) {
                switch 261 {
                case 2147483647:
                    v13 := GenS1{F0: 1, F1: 42, F2: 42}
                    _ = v13
                case 9223372036854775807:
                    var v14 int64 = -1
                    _ = v14
                case 684:
                    v15 := v3["hello"]
                    _ = v15
                default:
                    if true {
                        v1 = v1
                        var v16 float64 = float64(1e10)
                        _ = v16
                    }
                }
                v17 := []int{366, -9223372036854775808, 4294967295, 9223372036854775807}
                fmt.Println(v1, v2, v12)
                _ = v17
            }
        }
        _ = len(v3)
        _ = v12
    } else {
        v5 += v4
        for fi18 := 0; fi18 < 8; fi18++ {
            if (!v1) {
                v2 += 1e10
                v5 -= (v4 | -1)
            }
            v19 := v3[(("foo bar" + ("foo bar" + "")) + (("" + (("\n\t" + "hello") + "")) + "a"))]
            _ = fi18
            _ = v19
        }
    }
    v20 := GenS1{F0: 128, F1: 42, F2: 2147483647}
    v1 = v1
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 uint32 = 65535
        v23 := GenS2{F0: 0, F1: 42, F2: 2}
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 string = ("foo bar" + "a")
    var v25 int64 = 0
    switch 1 {
    case 1:
        var v26 int16 = -456
        _ = v26
    case 274:
        v27 := GenS1{F0: int(v25), F1: 282, F2: 4294967295}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
