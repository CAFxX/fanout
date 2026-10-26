package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint16 {
    var v7 uint64 = 1
    v8 := GenS1{F0: (v7 | v7)}
    if (true || false) {
        if false {
            var v9 uint16 = 255
            _ = v9
        }
    }
    _ = v7
    _ = v8
    return 601
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = (1.5 + 0.1)
    var v3 int = 0
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    var v6 uint8 = 255
    var v10 uint64 = v1
    v11 := []int{(v3 - (v3 << ((945) & 63))), v3}
    var v12 int = v3
    if false {
        var v13 string = "hello"
        _ = v13
    } else {
        v14 := v4[v3]
        _ = v14
    }
    var v15 int32 = -14
    v3 -= 140
    v15 += (-(v15))
    v16 := v4[4294967295]
    fmt.Println(v15, v12, v15)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
