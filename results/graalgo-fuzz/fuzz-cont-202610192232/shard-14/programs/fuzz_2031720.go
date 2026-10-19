package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 string }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]int)
    var v3 bool = true
    v4 := make(map[string]string)
    var v5 uint64 = uint64(v1)
    var v6 int64 = 0
    var v7 int64 = (v6 % (((v6 * (v6 + v6))) | 1))
    var v8 int32 = 1
    v5 = 2
    delete(v4, "\n\t")
    switch 128 {
    case 128:
        if (128 > 1) {
            var v9 uint16 = 988
            v10 := []int{32767}
            var v11 float64 = 0.0
            _ = v10
            _ = v11
            _ = v9
        } else {
            fmt.Println(v1, v8, v8)
        }
    case 9223372036854775807:
        var v12 int64 = (v6 + (v7 / ((v7) | 1)))
        _ = v12
    case -2147483648:
        switch 128 {
        case 386:
            fmt.Println(v5)
        case -56:
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 uint = 42
                var v15 uint32 = 2
                _ = fi13
                _ = v14
                _ = v15
            }
        case 127:
            var v16 int8 = -91
            _ = v16
        }
    }
    switch 128 {
    case 208:
        fmt.Println(v1)
    }
    v7 = 819
    v17 := v4[(("\n\t" + "") + ("hello" + "\n\t"))]
    v1 += v1
    switch (v17 + v17) {
    case "a":
        var v18 uint32 = 681
        _ = v18
    case "":
        fmt.Println(v3, v1)
    default:
        fmt.Println(v8, v6)
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        switch (v17 + v17) {
        case "a":
            var v20 float32 = 1e10
            _ = v20
        }
        if v3 {
            v7 = v7
            v1 = (v1 >> ((331) & 63))
            var v21 int64 = (^(v6))
            _ = v21
        }
        _ = fi19
    }
    if (255 <= 65535) {
        v22 := GenS1{F0: 0}
        switch -922 {
        case 127:
            var v23 uint32 = 2
            _ = v23
        case 9223372036854775807:
            v24 := GenS2{F0: 1, F1: "hello"}
            _ = v24
        default:
            var v25 float32 = ((0.0 - 1e10) / 3.14159)
            _ = v25
        }
        v26 := GenS2{F0: v1, F1: v17}
        _ = v22
        _ = v26
    }
    var v27 string = v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
