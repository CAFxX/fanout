package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64, p1 string) uint {
    p0 -= p0
    switch "\n\t" {
    case "a":
        v7 := []int{883, 32768, 2147483647, 0}
        _ = v7
    }
    var v8 int8 = 100
    v9 := GenS1{F0: p0, F1: 128, F2: p0}
    var v10 string = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]string)
    var v3 int8 = -128
    var v4 string = "\n\t"
    var v5 uint16 = 32767
    var v6 int32 = -2147483648
    delete(v2, 2147483648)
    var v11 int16 = 671
    fmt.Println(v3, v4, v4)
    var v12 int = 710
    v13 := GenS2{F0: -2147483648}
    v14 := GenS1{F0: 42, F1: 127, F2: 1}
    v1 += v1
    var v15 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
