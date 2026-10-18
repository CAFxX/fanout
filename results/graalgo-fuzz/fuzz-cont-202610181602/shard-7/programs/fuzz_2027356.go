package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float32 = float32(0.1)
    var v2 int16 = -1
    var v3 float64 = 1.5
    var v4 int8 = 41
    var v5 uint16 = 109
    var v6 int8 = (v4 >> ((97) & 63))
    var v7 int16 = 32767
    switch ("hello" + "") {
    case "":
        v4 = (-(v4))
    default:
        v7 = 1000
    }
    v8 := []int{419, 940}
    if len(v8) > 0 {
        v8[0] = 32768
    }
    v3 -= v3
    v9 := GenS1{F0: 4294967295}
    var v10 int32 = 42
    v11 := GenS2{F0: 2}
    v7 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
