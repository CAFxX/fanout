package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int16 }

func main() {
    var v1 int = 0
    var v2 int32 = 0
    var v3 int8 = -1
    var v4 int = (v1 ^ v1)
    v5 := make(map[string]int)
    var v6 int16 = 32767
    v7 := v5[(("a" + "\n\t") + "foo bar")]
    switch (v7 & (v4 / ((v7) | 1))) {
    case 128:
        var v8 float32 = (0.0 * (0.0 / 1.0))
        _ = v8
    case 32767:
        v9 := GenS1{F0: 4294967295}
        _ = v9
    case 9223372036854775807:
        v10 := GenS1{F0: 4294967295}
        _ = v10
    default:
        var v11 uint = 255
        _ = v11
    }
    if false {
        var v12 uint = 713
        delete(v5, "")
        for fi13 := 0; fi13 < 6; fi13++ {
            delete(v5, "\n\t")
            _ = fi13
        }
        _ = v12
    }
    var v14 uint16 = 870
    switch v7 {
    case -930:
        var v15 int8 = v3
        _ = v15
    default:
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 int32 = (v2 * (-(v2)))
            if (((!true) || true) || (!(4294967295 > 0))) {
                v1 = -1
            } else {
                v18 := []int{v7, (v4 << ((42) & 63)), (v4 << ((0) & 63)), 127}
                _ = v18
            }
            _ = fi16
            _ = v17
        }
    }
    v14 += 32767
    v19 := []int{(v1 << ((418) & 63)), int(v2)}
    fmt.Println(v2, v3)
    v4 += v1
    v20 := GenS2{F0: 65535, F1: v6}
    var v21 uint8 = 1
    var v22 int = v4
    v4 = 127
    switch (v22 | -9223372036854775808) {
    case 125:
        fmt.Println(v22)
    case 32768:
        fmt.Println(v7, v21)
    default:
        v21 = (v21 / (((v21 & (v21 / ((128) | 1)))) | 1))
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        fmt.Println(v6)
        v24 := GenS1{F0: 65535}
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: 581, F1: (v6 << ((4294967295) & 63))}
    var v26 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
