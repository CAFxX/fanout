package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int16 }
func genFn0() uint8 {
    v5 := []int{}
    var v6 float32 = float32(-2.25)
    _ = v5
    _ = v6
    return 128
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    var v3 int = 348
    v4 := make(map[string]int)
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 42, F1: -1, F2: -32768}
        v3 -= 128
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: -655, F2: 0}
    v10 := []int{v3, (v3 / (((v3 ^ (v3 % ((-2147483648) | 1)))) | 1))}
    v11 := GenS1{F0: 291, F1: 42, F2: -215}
    var v12 string = "\n\t"
    fmt.Println(v3)
    v13 := GenS1{F0: -9223372036854775808, F1: 32767, F2: -878}
    _ = len(v2)
    v14 := GenS1{F0: 0, F1: 1, F2: -320}
    v12 = "hello"
    var v15 int = (v3 >> ((18446744073709551615) & 63))
    if true {
        v3 = 32768
    }
    v10 = append(v10, 167)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
