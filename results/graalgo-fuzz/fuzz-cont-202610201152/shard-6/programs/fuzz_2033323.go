package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0() uint32 {
    v7 := GenS1{F0: 32767, F1: 100}
    var v8 uint8 = 2
    v9 := []int{-9223372036854775808, -453}
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int16 = -1
    var v2 string = "hello"
    var v3 int64 = 42
    v4 := []uint32{1, 522, 65535, 2147483648}
    var v5 int32 = -2147483648
    v6 := make(map[int]string)
    var v10 uint64 = 735
    var v11 int = -1
    var v12 uint8 = 1
    var v13 uint = 65535
    v3 = 628
    switch int(v13) {
    case 121:
        var v14 uint32 = 671
        _ = v14
    case -612:
        fmt.Println(v2, v11, v5)
    case -2147483648:
        var v15 uint8 = v12
        _ = v15
    }
    fmt.Println(v10, v12)
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(v1)
        _ = fi16
    }
    fmt.Println(v1, v2, v13)
    if false {
        v5 -= v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
