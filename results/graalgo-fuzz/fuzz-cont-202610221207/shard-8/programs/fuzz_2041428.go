package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint; F2 uint16 }

func main() {
    var v1 int32 = 0
    var v2 int32 = (v1 + v1)
    v3 := []int{-9223372036854775808}
    var v4 int = -863
    var v5 uint64 = 18446744073709551615
    v6 := make(map[int]uint32)
    if false {
        switch 1 {
        case -1:
            var v7 uint64 = v5
            _ = v7
        }
    }
    var v8 uint8 = 243
    var v9 bool = false
    fmt.Println(v9)
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            v6[475] = 4294967295
            v1 = v2
            _ = fi11
        }
        v12 := GenS1{F0: v5}
        v1 -= v2
        _ = fi10
        _ = v12
    }
    var v13 float32 = (float32(float64((1.5 * 1e10))) * 0.1)
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 int32 = (v2 << ((2) & 63))
        var v16 int8 = 127
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint32 = 115
    switch (v4 ^ v4) {
    case 42:
        var v18 uint32 = (v17 ^ 2147483648)
        _ = v18
    case 128:
        v9 = false
    case -731:
        if v9 {
            var v19 int = v4
            v20 := GenS2{F0: uint32(v8)}
            _ = v19
            _ = v20
        } else {
            if v9 {
                var v21 int8 = -69
                _ = v21
            }
            fmt.Println(v8)
            v5 = v5
        }
    }
    var v22 uint = 18446744073709551615
    v23 := GenS3{F0: "a", F1: uint(v8), F2: 42}
    var v24 int = (v4 | -828)
    v24 = v4
    fmt.Println(v8, v1, v13)
    var v25 int64 = -1
    var v26 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
