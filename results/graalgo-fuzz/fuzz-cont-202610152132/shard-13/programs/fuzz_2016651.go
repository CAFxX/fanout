package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int64) uint32 {
    v4 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 1
}

func main() {
    var v1 uint = 1
    var v2 uint64 = 907
    var v3 uint64 = v2
    v3 += 392
    v1 += v1
    v2 += (v3 - 2)
    var v5 uint8 = 128
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: true}
    fmt.Println(v5, v5)
    v8 := GenS1{F0: (!false)}
    v9 := GenS1{F0: true}
    var v10 int16 = -97
    var v11 int64 = 1
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint32 = genFn0(v11)
        _ = fi12
        _ = v13
    }
    var v14 int16 = (v10 & v10)
    v1 = (v1 << (((v1 << ((244) & 63))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
