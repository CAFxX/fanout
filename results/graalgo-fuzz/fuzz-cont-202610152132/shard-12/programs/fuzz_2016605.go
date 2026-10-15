package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 378
    v3 := make(map[uint32]int)
    var v4 uint64 = 42
    var v5 uint8 = 2
    if true {
        v6 := GenS1{F0: uint32(v4)}
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    v9 := GenS1{F0: 0}
    if false {
        var v10 uint8 = (v5 * v5)
        v5 = v5
        _ = v10
    }
    var v11 uint64 = 201
    var v12 bool = (v2 >= v2)
    v4 += (v11 / ((v11) | 1))
    var v13 float64 = 3.14159
    var v14 int16 = int16(v4)
    var v15 string = "\n\t"
    v14 -= (v14 / ((int16(v5)) | 1))
    v16 := GenS1{F0: 2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
