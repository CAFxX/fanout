package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]string)
    var v3 uint8 = (v1 - 42)
    v3 = (v3 | (v1 | (v1 % (((v3 * (v1 * (v1 << ((0) & 63))))) | 1))))
    v4 := GenS1{F0: 204}
    v5 := GenS1{F0: 4294967295}
    for fi6 := 0; fi6 < 6; fi6++ {
        v3 = (v1 >> ((358) & 63))
        _ = fi6
    }
    switch "" {
    case "a":
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 805}
            _ = fi7
            _ = v8
        }
    }
    var v9 bool = (!false)
    if (4294967295 > 2147483647) {
        v10 := []int{128, 32768, 793}
        _ = v10
    }
    fmt.Println(v1, v9)
    v3 = (v3 << ((876) & 63))
    var v11 bool = true
    v12 := []int{4294967295, 32767, 4294967295, 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
