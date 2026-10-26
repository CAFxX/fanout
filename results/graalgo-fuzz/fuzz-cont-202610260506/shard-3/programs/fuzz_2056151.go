package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int16; F2 uint64 }
func genFn0() uint16 {
    var v4 int = -442
    var v5 uint64 = 777
    var v6 int16 = -32768
    _ = v4
    _ = v5
    _ = v6
    return 728
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint32 = (v1 & 4294967295)
    v3 := make(map[uint32]bool)
    var v7 uint32 = 219
    v7 -= (v1 << ((242) & 63))
    var v8 string = ("a" + "hello")
    v9 := GenS1{F0: v8}
    v7 = v1
    if true {
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int64 = 0
            _ = fi10
            _ = v11
        }
        fmt.Println(v8, v2, v7)
    } else {
        var v12 uint = 4294967295
        _ = v12
    }
    var v13 bool = ((998 < -383) && ((!true) || true))
    v14 := GenS2{F0: -1, F1: -1, F2: 1}
    v13 = v13
    v2 -= v2
    for fi15 := 0; fi15 < 12; fi15++ {
        v7 = 0
        _ = fi15
    }
    fmt.Println(v2, v7, v13)
    if (4294967295 >= 792) {
        v16 := []int{1, 9223372036854775807, 0, 4294967295}
        v7 -= v1
        _ = v16
    } else {
        var v17 int8 = 1
        if true {
            fmt.Println(v8, v8)
        }
        _ = v17
    }
    v13 = ((!(true && (!true))) && false)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
