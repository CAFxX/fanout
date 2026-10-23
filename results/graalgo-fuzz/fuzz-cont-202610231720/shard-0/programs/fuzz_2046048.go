package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int64 }

func main() {
    var v1 int = 1
    v2 := make(map[string]int)
    var v3 bool = (false && ((!((1e10 + (float32(0.1) - 3.14159)) >= float32((((float64(float32(1e10)) / 1.0) / 1.5) / 1e10)))) && ((2147483647 < 42) && false)))
    v3 = (475 > -32768)
    v4 := []int{v1, v1, 127}
    if (0.1 > (0.1 + (1e10 / float32(1e10)))) {
        delete(v2, "\n\t")
    } else {
        v5 := GenS1{F0: v1, F1: -1, F2: 9223372036854775807}
        v6 := v2["\n\t"]
        v3 = (18446744073709551615 <= 1)
        _ = v5
        _ = v6
    }
    v3 = v3
    v3 = v3
    v1 = v1
    var v7 string = "a"
    v1 += v1
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint = 824
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: v1, F1: 32767, F2: 441}
    if v3 {
        v3 = ((963 == 179) && v3)
        v1 = -613
        var v11 uint32 = 65535
        _ = v11
    }
    var v12 uint16 = 2
    var v13 int8 = -66
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
