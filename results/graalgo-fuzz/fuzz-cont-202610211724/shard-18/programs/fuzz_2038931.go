package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 960
    var v3 bool = ("" > ("" + ""))
    v3 = (42 >= 48)
    var v4 uint16 = 0
    var v5 int8 = 42
    var v6 int32 = -805
    v3 = false
    for fi7 := 0; fi7 < 4; fi7++ {
        _ = len(v1)
        v6 += int32(v2)
        if (((710 > (v2 | v2)) && v3) || v3) {
            for fi8 := 0; fi8 < 4; fi8++ {
                v5 = (v5 - v5)
                _ = fi8
            }
            var v9 int32 = v6
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: (v5 * -128), F1: (v6 * v6), F2: v5}
    v11 := []int{42, 331, -917}
    var v12 int64 = 466
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
