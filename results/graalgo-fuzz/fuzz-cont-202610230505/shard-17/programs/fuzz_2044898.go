package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint32 }

func main() {
    var v1 uint8 = 1
    var v2 int64 = -887
    var v3 int32 = 42
    v4 := GenS1{F0: v1, F1: 65535, F2: 2147483647}
    fmt.Println(v2, v3, v3)
    v2 -= (v2 + v2)
    v3 += (v3 << ((954) & 63))
    v3 = v3
    switch -1 {
    case 973:
        v1 = v1
    case 42:
        v5 := []int{-2147483648}
        _ = v5
    default:
        var v6 int64 = v2
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 1, F1: 652, F2: 1}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (^(v1)), F1: 2147483647, F2: 4294967295}
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint8 = (v1 & (v1 >> ((0) & 63)))
        v12 := []int{(fi10 ^ (fi10 << ((782) & 63))), fi10, fi10}
        if (1000 != -32768) {
            v13 := GenS1{F0: v11, F1: 2147483647, F2: 2147483648}
            v3 += v3
            _ = v13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
