package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint16 {
    var v6 uint64 = 4294967295
    switch 0 {
    case 941:
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := []int{32768, 9223372036854775807, 735, fi7}
            _ = fi7
            _ = v8
        }
    case 9223372036854775807:
        var v9 int32 = 2147483647
        _ = v9
    }
    _ = v6
    return 65535
}

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]bool)
    var v3 int64 = 143
    var v4 uint8 = 42
    v5 := make(map[uint32]int)
    switch "" {
    case "foo bar":
        v10 := []int{-991}
        _ = v10
    default:
        v11 := GenS1{F0: 65535}
        _ = v11
    }
    v12 := GenS1{F0: 330}
    delete(v5, 0)
    var v13 int8 = 22
    v14 := v5[2147483648]
    var v15 int16 = -32768
    v16 := []int{127, (v14 - v14)}
    v5[0] = v14
    v4 += v4
    fmt.Println(v13, v15, v15)
    _ = len(v2)
    var v17 float32 = (((float32(((1e10 + (1e10 / ((-2.25 + float64((1.5 / 3.14159))) - 1.5))) - -2.25)) * 0.0) - ((3.14159 * float32(3.14159)) + float32(1.5))) * 1.5)
    var v18 uint16 = 255
    v17 = v17
    v5[0] = (v14 | v14)
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
