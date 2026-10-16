package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 string = "\n\t"
    var v3 uint64 = 833
    for fi4 := 0; fi4 < 15; fi4++ {
        v2 = v2
        switch ((v2 + "hello") + "foo bar") {
        case "\n\t":
            if true {
                v5 := []int{fi4, (-(fi4)), (fi4 << ((42) & 63)), fi4}
                v3 = (v3 & (v3 * 759))
                _ = len(v5)
                _ = v5
            }
        }
        _ = fi4
    }
    var v6 float32 = (3.14159 * float32(0.0))
    var v7 uint8 = 137
    v7 -= 127
    var v8 float64 = -2.25
    var v9 int32 = 2147483647
    _ = len(v1)
    v10 := GenS1{F0: (v9 - -1), F1: 870, F2: 2147483648}
    delete(v1, -9223372036854775808)
    delete(v1, 42)
    v1[42] = 32767
    v1[128] = -1
    var v11 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
