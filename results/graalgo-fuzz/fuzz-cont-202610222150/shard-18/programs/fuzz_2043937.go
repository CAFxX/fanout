package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{32768, -971, 1}
    var v3 uint32 = 2147483647
    var v4 int = 128
    fmt.Println(v4, v3)
    v4 = (v4 & v4)
    v5 := v1["\n\t"]
    v6 := v1[((v5 + v5) + "\n\t")]
    v6 = "a"
    v7 := GenS1{F0: -32768, F1: 32767, F2: 158}
    v8 := GenS1{F0: 1, F1: -32768, F2: (v3 << ((266) & 63))}
    v3 -= v3
    v4 -= 4294967295
    if false {
        v6 = v6
        v9 := GenS2{F0: 42}
        var v10 float64 = float64(3.14159)
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint = 1
            var v13 uint64 = 42
            _ = fi11
            _ = v12
            _ = v13
        }
        fmt.Println(v3, v4)
        v14 := GenS3{F0: 25}
        _ = v14
    }
    fmt.Println(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
