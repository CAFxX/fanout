package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int; F1 int; F2 uint8 }

func main() {
    var v1 uint8 = 0
    var v2 uint32 = 57
    var v3 uint16 = 42
    var v4 float64 = 3.14159
    var v5 float64 = float64((0.1 + -2.25))
    var v6 float32 = float32(v4)
    fmt.Println(v6, v3, v2)
    if true {
        switch 1 {
        case 610:
            v7 := GenS1{F0: -581, F1: v3}
            _ = v7
        default:
            v6 = -2.25
        }
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint32 = (v2 << ((42) & 63))
        v10 := []int{fi8, -332}
        _ = len(v10)
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v3, v3)
    if true {
        if true {
            var v11 int8 = 127
            _ = v11
        }
    }
    if true {
        v12 := GenS2{F0: 127, F1: 1, F2: v1}
        _ = v12
    } else {
        v6 = float32(v4)
        v5 += v4
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 string = "a"
        var v15 float64 = v5
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 string = "\n\t"
    var v17 string = v16
    var v18 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
