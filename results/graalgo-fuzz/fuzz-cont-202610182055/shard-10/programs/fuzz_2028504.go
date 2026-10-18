package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 int = 32768
    v3 -= v3
    v3 += v3
    v3 = -633
    v3 += -9223372036854775808
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := GenS1{F0: 1000}
        var v6 uint64 = 4294967295
        if true {
            var v7 int16 = 728
            _ = v7
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    v3 -= (v3 + (v3 - (v3 ^ v3)))
    v3 -= 6
    fmt.Println(v3)
    v8 := []int{(v3 >> ((42) & 63))}
    var v9 uint = 364
    fmt.Println(v3, v9)
    v10 := GenS1{F0: 1}
    switch v3 {
    case -1:
        v9 += v9
    case 128:
        var v11 float64 = -2.25
        _ = v11
    }
    fmt.Println(v3)
    switch v3 {
    case 128:
        fmt.Println(v3, v3)
    case 4294967295:
        v9 += (v9 & (v9 >> ((v9) & 63)))
    case -2147483648:
        fmt.Println(v9, v3)
    }
    var v12 float32 = 3.14159
    if (971 == uint16(v9)) {
        var v13 int64 = 42
        for fi14 := 0; fi14 < 2; fi14++ {
            v13 += v13
            var v15 uint = v9
            v3 = 1
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    var v16 uint32 = uint32(v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
