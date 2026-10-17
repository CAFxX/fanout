package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int64, p2 bool) int8 {
    fmt.Println(p0)
    p0 += (p0 ^ (p0 >> ((1) & 63)))
    var v6 float64 = float64(3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 100
}

func main() {
    var v1 uint32 = 638
    var v2 int = -593
    var v3 uint = 255
    var v4 uint8 = 128
    var v5 int16 = -32768
    v5 = v5
    v7 := GenS1{F0: v1}
    v4 += (v4 >> ((v3) & 63))
    if false {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int64 = 0
            var v10 uint64 = 489
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    var v11 uint = 143
    v11 = 300
    var v12 int8 = 0
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint16 = 875
        _ = fi13
        _ = v14
    }
    var v15 bool = (!(!true))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
