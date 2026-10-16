package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint = 1
    var v2 uint16 = 572
    var v3 int16 = int16(v1)
    var v4 bool = true
    v5 := GenS1{F0: 255}
    fmt.Println(v2)
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{fi6, fi6, (fi6 & fi6)}
        for fi8 := 0; fi8 < 14; fi8++ {
            if v4 {
                v9 := GenS1{F0: v2}
                _ = v9
            }
            var v10 int32 = -1
            v3 += (v3 << (((v1 & v1)) & 63))
            _ = fi8
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    if v4 {
        fmt.Println(v2, v3)
        v11 := []int{}
        _ = v11
    }
    var v12 int32 = 1
    for fi13 := 0; fi13 < 17; fi13++ {
        fmt.Println(v2, v3, v12)
        _ = fi13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 float32 = 0.0
        _ = fi14
        _ = v15
    }
    switch "hello" {
    case "":
        fmt.Println(v3, v3, v4)
    case "hello":
        fmt.Println(v1)
    case "\n\t":
        var v16 int16 = v3
        _ = v16
    }
    var v17 int = 127
    v18 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
