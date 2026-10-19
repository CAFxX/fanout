package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
func genFn0() int64 {
    var v6 float32 = (0.0 - 0.0)
    if false {
        fmt.Println(v6)
    }
    var v7 int32 = 81
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []int{}
    var v2 int64 = 42
    var v3 int16 = 42
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    var v8 int8 = 1
    var v9 int8 = (v8 >> ((354) & 63))
    delete(v4, 760)
    v1 = append(v1, 1)
    v10 := GenS1{F0: false}
    fmt.Println(v8)
    var v11 uint64 = 0
    for fi12 := 0; fi12 < 12; fi12++ {
        delete(v5, (fi12 / ((4294967295) | 1)))
        _ = fi12
    }
    v13 := GenS2{F0: v2}
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v14 int32 = -2147483648
        v15 := GenS1{F0: true}
        _ = v14
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
