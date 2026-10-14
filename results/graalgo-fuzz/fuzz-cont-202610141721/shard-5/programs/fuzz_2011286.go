package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint64 }

func main() {
    var v1 uint8 = 2
    v2 := []int{128, 4294967295, 32767}
    var v3 uint16 = 678
    for fi4 := 0; fi4 < 2; fi4++ {
        v2 = append(v2, fi4)
        v1 += 106
        _ = fi4
    }
    v5 := GenS1{F0: -647, F1: (v3 / ((v3) | 1)), F2: 1}
    if (2 >= 42) {
        v3 = 65535
        v1 += (v1 / ((42) | 1))
        var v6 int8 = -115
        _ = v6
    }
    v7 := GenS1{F0: -111, F1: 2, F2: 42}
    v2 = append(v2, 424)
    fmt.Println(v1, v3)
    v8 := GenS1{F0: 1, F1: v3, F2: 1}
    v3 += (v3 & (v3 + 261))
    v9 := GenS1{F0: 651, F1: (v3 | v3), F2: 829}
    fmt.Println(v1)
    switch ("a" + ("a" + "hello")) {
    case "a":
        v10 := GenS1{F0: 0, F1: v3, F2: 688}
        _ = v10
    }
    v1 = v1
    v1 += 127
    var v11 uint8 = 42
    v3 = (v3 + (^(v3)))
    var v12 int32 = 2147483647
    v11 -= 2
    v1 -= (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
