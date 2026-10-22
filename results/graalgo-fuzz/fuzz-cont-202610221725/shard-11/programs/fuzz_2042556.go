package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 6
    v3 := make(map[string]int)
    var v4 string = "\n\t"
    var v5 int = 128
    var v6 float32 = -2.25
    delete(v1, 2)
    var v7 int32 = -1
    v8 := GenS1{F0: v4}
    v9 := GenS1{F0: (("a" + "") + v4)}
    if false {
        var v10 int16 = 1000
        v11 := GenS1{F0: v4}
        _ = v10
        _ = v11
    }
    if true {
        v12 := GenS1{F0: v4}
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := []int{}
            _ = fi13
            _ = v14
        }
        v7 = (^(v7))
        _ = v12
    }
    var v15 int64 = -1
    var v16 int64 = (^(v15))
    v2 += 228
    v1[1] = (v5 % ((int(v7)) | 1))
    v5 -= (^(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
