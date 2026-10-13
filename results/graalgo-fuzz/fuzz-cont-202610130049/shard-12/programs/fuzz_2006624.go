package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint = 0
    var v2 uint = (v1 | (v1 & v1))
    var v3 string = "a"
    v3 = "foo bar"
    var v4 int = 127
    v2 = (v1 >> ((2) & 63))
    v1 -= 207
    var v5 float64 = (0.0 / (3.14159 + float64((((3.14159 - 3.14159) - (3.14159 * 0.0)) * 0.1))))
    fmt.Println(v2, v3, v3)
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: (false && (v2 != v2))}
        switch (v4 % ((128) | 1)) {
        case 985:
            v8 := GenS1{F0: true}
            _ = v8
        case -110:
            v9 := GenS1{F0: false}
            _ = v9
        default:
            var v10 float64 = ((float64((float32(v5) + (0.0 / 1.5))) * v5) + ((((-2.25 + (3.14159 / v5)) * v5) / v5) / 3.14159))
            _ = v10
        }
        var v11 bool = true
        _ = fi6
        _ = v11
        _ = v7
    }
    var v12 int32 = 42
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v12)
        v14 := []int{fi13, -9223372036854775808, v4}
        _ = fi13
        _ = v14
    }
    switch ("" + (v3 + (v3 + v3))) {
    case "":
        if ((0 > 42) || true) {
            if false {
                v15 := []int{}
                v12 = v12
                fmt.Println(v2)
                _ = v15
            }
            v16 := []int{-9223372036854775808}
            switch v4 {
            case 32768:
                var v17 int16 = 64
                _ = v17
            case -12:
                var v18 float64 = (float64(0.0) * v5)
                _ = v18
            case 2147483647:
                var v19 uint8 = 127
                _ = v19
            default:
                v20 := GenS2{F0: v12}
                _ = v20
            }
            _ = v16
        }
    default:
        v12 = v12
    }
    v21 := []int{(v4 | v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
