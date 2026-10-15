package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint16) uint64 {
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{(fi6 - fi6)}
        _ = fi6
        _ = v7
    }
    var v8 int32 = -2147483648
    v9 := GenS1{F0: -87, F1: 2, F2: "\n\t"}
    _ = p0
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]string)
    v3 := []uint32{137, 212, 2}
    var v4 float32 = 3.14159
    var v5 uint64 = 793
    var v10 uint16 = 0
    v11 := GenS1{F0: 0, F1: 165, F2: (("foo bar" + ("hello" + "hello")) + "\n\t")}
    if false {
        v12 := GenS2{F0: 15}
        _ = v12
    } else {
        var v13 int16 = 548
        _ = v13
    }
    v14 := GenS1{F0: 68, F1: 90, F2: "foo bar"}
    _ = len(v2)
    var v15 float32 = v4
    var v16 uint = (^(v1))
    fmt.Println(v5, v16)
    var v17 bool = (!false)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
