package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
func genFn0() int8 {
    v5 := []int{}
    v6 := GenS1{F0: -32768, F1: 2147483647}
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 int = -1
    var v3 int32 = 665
    var v4 int16 = 42
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{(v2 ^ int(v4)), 4294967295, (fi7 + 4294967295)}
        var v9 uint8 = 255
        _ = fi7
        _ = v8
        _ = v9
    }
    v4 = -846
    v2 -= v2
    if (false && (!((v2 - 0) > (v2 << ((18446744073709551615) & 63))))) {
        var v10 int8 = 127
        var v11 int8 = -123
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: (v4 | (v4 / ((v4) | 1))), F1: int32(v2)}
    delete(v1, "\n\t")
    v13 := GenS1{F0: 42, F1: (v3 ^ (v3 + 42))}
    if true {
        for fi14 := 0; fi14 < 10; fi14++ {
            v3 -= v3
            v15 := []int{fi14}
            fmt.Println(fi14)
            _ = fi14
            _ = v15
        }
        var v16 uint64 = 187
        v16 += v16
        _ = v16
    } else {
        fmt.Println(v2)
    }
    v4 += 40
    fmt.Println(v4, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
