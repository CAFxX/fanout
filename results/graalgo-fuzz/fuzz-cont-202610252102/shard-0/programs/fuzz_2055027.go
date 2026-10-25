package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
func genFn0(p0 int8, p1 string) int16 {
    var v7 int16 = 42
    if false {
        var v8 int8 = p0
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return v7
}

func main() {
    v1 := []int{9223372036854775807}
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    var v4 float32 = 1e10
    var v5 uint64 = 4294967295
    var v6 float32 = v4
    v6 = (v4 + 0.1)
    if true {
        var v9 int32 = 287
        switch "foo bar" {
        case "hello":
            var v10 int64 = -756
            _ = v10
        case "\n\t":
            var v11 int32 = (v9 | -232)
            _ = v11
        case "a":
            for fi12 := 0; fi12 < 7; fi12++ {
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 bool = true
                    if (42 < 127) {
                        var v15 uint = 18446744073709551615
                        v1 = append(v1, fi13)
                        var v16 float32 = (v6 + 1e10)
                        _ = v15
                        _ = v16
                    } else {
                        switch -9223372036854775808 {
                        case 32767:
                            v17 := GenS1{F0: -1, F1: 255}
                            _ = v17
                        default:
                            v18 := GenS1{F0: int16(v9), F1: 255}
                            _ = v18
                        }
                        var v19 uint64 = 1
                        _ = v19
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v6 -= 3.14159
                    _ = fi13
                    _ = v14
                }
                var v20 uint8 = 58
                _ = fi12
                _ = v20
            }
        }
        _ = v9
    }
    fmt.Println(v5)
    v5 = v5
    v21 := GenS1{F0: 1000, F1: uint(v5)}
    fmt.Println(v4, v6)
    var v22 bool = (true || false)
    _ = len(v1)
    var v23 float32 = float32(float64(v6))
    v24 := GenS1{F0: 0, F1: 42}
    v22 = ((v4 != v23) && ((float32(1e10) >= v6) || ((genFn0(-128, "foo bar") < 1) || (1.5 > float64(v4)))))
    if v22 {
        fmt.Println(v4)
        var v25 float32 = 0.0
        var v26 int32 = -467
        _ = v25
        _ = v26
    } else {
        var v27 int32 = -556
        switch -9223372036854775808 {
        case -9223372036854775808:
            fmt.Println(v27, v27)
        case 42:
            v28 := GenS1{F0: 1, F1: 0}
            _ = v28
        default:
            var v29 int8 = 0
            _ = v29
        }
        v23 -= (0.0 + 1.5)
        _ = v27
    }
    fmt.Println(v5, v6, v5)
    v30 := v2[-9223372036854775808]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
