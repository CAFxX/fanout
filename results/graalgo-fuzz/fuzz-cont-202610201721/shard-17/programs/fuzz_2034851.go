package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint64; F1 int }
func genFn0(p0 bool, p1 int8) uint16 {
    if p0 {
        var v7 uint8 = 255
        _ = v7
    }
    var v8 uint16 = 255
    fmt.Println(v8, p0, p0)
    p1 = p1
    _ = p0
    _ = p1
    _ = v8
    return (v8 - (v8 - 506))
}

func main() {
    var v1 uint16 = 905
    var v2 bool = true
    var v3 float64 = 0.0
    var v4 string = ""
    v5 := make(map[uint32]uint32)
    var v6 int = 42
    var v9 uint64 = 18446744073709551615
    var v10 uint = 450
    fmt.Println(v3)
    v11 := []int{0, v6}
    v9 -= (v9 << ((v10) & 63))
    v12 := GenS1{F0: -48, F1: (v9 - uint64(v10)), F2: 100}
    v2 = (!v2)
    var v13 int32 = 2147483647
    v6 = 2147483647
    _ = len(v5)
    switch (((((v4 + ("hello" + "")) + v4) + "a") + ((v4 + ((v4 + v4) + v4)) + "hello")) + ((v4 + v4) + ((v4 + "a") + v4))) {
    case "":
        v14 := GenS2{F0: (v9 % (((v9 >> ((856) & 63))) | 1)), F1: (v6 | v6)}
        _ = v14
    default:
        if len(v11) > 0 {
            v11[0] = v6
        }
    }
    var v15 int64 = 42
    fmt.Println(v6)
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 uint = v10
        _ = fi16
        _ = v17
    }
    v5[2] = 65535
    var v18 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
