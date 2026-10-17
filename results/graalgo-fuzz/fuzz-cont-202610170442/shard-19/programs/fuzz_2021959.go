package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int }
type GenS2 struct { F0 bool }
func genFn0() int8 {
    v7 := []int{713, 492}
    var v8 int8 = 100
    v9 := GenS1{F0: 2, F1: 2, F2: -9223372036854775808}
    v10 := GenS1{F0: 2, F1: 0, F2: int(v8)}
    switch -9223372036854775808 {
    case 42:
        v8 = (v8 | v8)
    case 411:
        var v11 int = 127
        _ = v11
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (v8 ^ (v8 % ((100) | 1)))
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = float32(float64(0.0))
    var v3 uint64 = 0
    var v4 string = "foo bar"
    v5 := make(map[uint32]string)
    var v6 uint16 = 130
    v12 := []int{42, 237, -9223372036854775808, -932}
    switch v4 {
    case "foo bar":
        v6 += v6
    case "hello":
        switch -9223372036854775808 {
        case 0:
            v6 += v6
        case 128:
            v5[1] = v4
        }
    case "\n\t":
        switch v4 {
        case "a":
            var v13 int64 = 9223372036854775807
            _ = v13
        case "":
            var v14 int = -9223372036854775808
            _ = v14
        case "\n\t":
            v15 := GenS1{F0: 42, F1: v6, F2: 127}
            _ = v15
        }
    }
    v6 = v6
    var v16 uint8 = 168
    v16 = 0
    v17 := GenS2{F0: (!(!(!true)))}
    v18 := GenS1{F0: v16, F1: v6, F2: 4294967295}
    v6 += (v6 << ((2) & 63))
    v19 := v1[-9223372036854775808]
    v20 := GenS2{F0: (1 == 42)}
    v6 += 2
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 int = 32768
        delete(v5, 65535)
        v4 = v4
        _ = fi21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
