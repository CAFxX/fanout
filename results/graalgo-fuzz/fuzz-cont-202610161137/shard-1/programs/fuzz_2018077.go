package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int32 {
    switch "foo bar" {
    case "hello":
        var v6 int64 = 452
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{fi7}
        v9 := GenS1{F0: 100}
        v8 = append(v8, 127)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float64 = 0.0
    fmt.Println(v10)
    v11 := []int{633, 9223372036854775807, -9223372036854775808}
    _ = v10
    _ = v11
    return -1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 1
    var v3 uint16 = 2
    var v4 string = "foo bar"
    var v5 uint64 = 4294967295
    v12 := []int{127, 127}
    v2 = (v2 | v2)
    v5 = v5
    var v13 int64 = -126
    v14 := GenS1{F0: 59}
    if (0 <= -584) {
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := GenS1{F0: -1}
            switch "\n\t" {
            case "foo bar":
                var v17 uint8 = 68
                _ = v17
            default:
                switch -730 {
                case 42:
                    var v18 float32 = 1e10
                    _ = v18
                }
            }
            _ = fi15
            _ = v16
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            v5 = (v5 | v5)
            _ = fi19
        }
    } else {
        fmt.Println(v4, v3, v5)
    }
    var v20 int8 = -98
    var v21 float64 = ((float64(0.1) * -2.25) - (0.0 - (-2.25 + 0.1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
