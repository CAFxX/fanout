package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }

func main() {
    var v1 int64 = 836
    v2 := make(map[string]int)
    var v3 int32 = -2147483648
    v4 := make(map[uint32]bool)
    var v5 int = 32768
    var v6 uint32 = 850
    v6 += (v6 % ((v6) | 1))
    var v7 uint16 = 2
    v8 := GenS1{F0: (18446744073709551615 != 241), F1: -1, F2: v7}
    v5 += (v5 << ((1) & 63))
    if ((false || true) && (!false)) {
        var v9 int64 = (v1 & v1)
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint8 = 117
        var v12 string = "hello"
        _ = fi10
        _ = v11
        _ = v12
    }
    v6 += v6
    v1 = v1
    v13 := GenS1{F0: false, F1: 100, F2: v7}
    fmt.Println(v6, v3)
    switch "foo bar" {
    case "a":
        var v14 int64 = int64(v6)
        _ = v14
    case "foo bar":
        if false {
            v15 := GenS1{F0: (247 <= 128), F1: 42, F2: v7}
            fmt.Println(v7)
            _ = len(v4)
            _ = v15
        }
    case "":
        v16 := []int{v5, (v5 * -543), -1}
        _ = v16
    }
    var v17 int8 = 118
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
