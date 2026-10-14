package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int64, p1 uint32, p2 uint8) int8 {
    p2 -= p2
    fmt.Println(p2, p2, p0)
    var v7 float32 = ((-2.25 + float32((1e10 + 0.1))) / -2.25)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 int8 = 100
    v2 := make(map[int]uint32)
    var v3 bool = (42 < 18446744073709551615)
    var v4 bool = v3
    var v5 int16 = 699
    v6 := make(map[int]string)
    var v8 uint = 4294967295
    v9 := GenS1{F0: 15, F1: 255}
    _ = len(v6)
    fmt.Println(v1)
    v10 := GenS1{F0: 65535, F1: 96}
    v11 := GenS1{F0: 65535, F1: 1}
    switch 127 {
    case 128:
        var v12 float64 = 1.5
        _ = v12
    case -174:
        var v13 int32 = -1
        _ = v13
    default:
        var v14 uint8 = 42
        _ = v14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS1{F0: 936, F1: 127}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 uint = (v8 % ((v8) | 1))
        switch (fi17 << ((v18) & 63)) {
        case 32767:
            var v19 uint8 = 42
            _ = v19
        case -9223372036854775808:
            var v20 string = "hello"
            _ = v20
        case 9223372036854775807:
            for fi21 := 0; fi21 < 17; fi21++ {
                var v22 uint = v8
                fmt.Println(v8, v5)
                v22 = (v22 % ((961) | 1))
                v23 := GenS1{F0: 542, F1: 207}
                fmt.Println(v4)
                var v24 bool = (!(32768 <= 9223372036854775807))
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
        }
        _ = fi17
        _ = v18
    }
    v25 := v2[-1]
    v5 += v5
    v5 += 42
    var v26 uint32 = v25
    var v27 uint8 = 255
    v28 := v6[-483]
    var v29 int8 = 100
    var v30 string = v28
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
