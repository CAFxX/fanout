package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint64 }
func genFn0(p0 int32) int32 {
    var v7 int64 = 9223372036854775807
    p0 += 42
    _ = p0
    _ = v7
    return (p0 & (p0 - (p0 ^ -2147483648)))
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 2
    var v3 bool = (true || (false && true))
    var v4 bool = true
    v5 := []int{-1, 293, int(v2), 42}
    var v6 uint64 = 4294967295
    v8 := GenS1{F0: 0, F1: (v4 && v3), F2: (v6 + (v6 << ((v2) & 63)))}
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: 852, F1: v3, F2: v6}
        _ = fi9
        _ = v10
    }
    var v11 int16 = 1
    var v12 string = "hello"
    var v13 uint = 0
    fmt.Println(v13)
    if v3 {
        var v14 int8 = 42
        switch 527 {
        case -2147483648:
            v4 = true
        case 1:
            fmt.Println(v14)
        case 32767:
            v15 := GenS1{F0: 2147483647, F1: false, F2: 4294967295}
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            fmt.Println(v11)
            _ = fi16
        }
        _ = v14
    } else {
        fmt.Println(v2, v13)
        var v17 int = -454
        _ = v17
    }
    var v18 bool = v4
    _ = len(v1)
    v19 := GenS1{F0: -1, F1: (42 > (v13 * 65535)), F2: (v6 - (v6 >> ((v13) & 63)))}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
