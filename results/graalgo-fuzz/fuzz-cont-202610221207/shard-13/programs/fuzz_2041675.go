package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
func genFn0(p0 uint8, p1 uint64, p2 uint) uint8 {
    for fi4 := 0; fi4 < 8; fi4++ {
        switch "foo bar" {
        case "hello":
            var v5 string = "hello"
            _ = v5
        case "":
            v6 := []int{737, fi4}
            _ = v6
        default:
            p0 = uint8(p1)
        }
        _ = fi4
    }
    if false {
        v7 := GenS1{F0: 563, F1: p1, F2: -773}
        _ = v7
    }
    var v8 int32 = 1
    var v9 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 int32 = 42
    var v2 int = 0
    v3 := []int{320, v2, (-(v2)), v2}
    v1 = 0
    var v10 bool = false
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int16 = -1
        _ = fi11
        _ = v12
    }
    v3 = append(v3, v2)
    if len(v3) > 0 {
        v3[0] = 42
    }
    v2 += (v2 + (-(v2)))
    fmt.Println(v2, v2, v10)
    for fi13 := 0; fi13 < 17; fi13++ {
        v1 = (v1 >> ((65535) & 63))
        _ = fi13
    }
    var v14 uint64 = 252
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
}
