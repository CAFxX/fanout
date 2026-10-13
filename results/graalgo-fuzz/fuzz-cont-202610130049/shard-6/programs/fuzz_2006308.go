package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 bool = true
    v2 := []string{""}
    var v3 uint64 = 42
    var v4 int32 = -1
    v5 := []uint32{}
    v6 := []string{}
    v7 := GenS1{F0: 704}
    v8 := GenS1{F0: 0}
    v4 += v4
    var v9 int64 = 1
    var v10 uint32 = 4294967295
    var v11 uint32 = 2
    v1 = (v1 || v1)
    if ((v10 & 4294967295) <= (v10 * 790)) {
        v9 = (v9 >> ((4294967295) & 63))
    } else {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        v11 -= (^(v10))
    }
    if v1 {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 int16 = 32767
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v14 float32 = 0.1
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    var v15 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
