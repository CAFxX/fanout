package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int32) uint32 {
    var v4 uint32 = 485
    fmt.Println(v4, v4)
    v4 += 675
    _ = p0
    _ = v4
    return (v4 << ((4294967295) & 63))
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    if true {
        var v5 float32 = 1.5
        _ = v5
    } else {
        v6 := []int{1, 9223372036854775807}
        v7 := GenS1{F0: 42}
        switch 42 {
        case -818:
            var v8 float64 = -2.25
            _ = v8
        default:
            var v9 string = "hello"
            _ = v9
        }
        _ = v6
        _ = v7
    }
    var v10 uint32 = 286
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int32 = 0
        var v13 int = (fi11 % (((-(fi11))) | 1))
        _ = fi11
        _ = v12
        _ = v13
    }
    v10 -= 4294967295
    fmt.Println(v10)
    v10 -= genFn0(-1)
    var v14 uint8 = 127
    if (uint16(v14) != 2) {
        for fi15 := 0; fi15 < 5; fi15++ {
            switch fi15 {
            case -286:
                v16 := GenS1{F0: 42}
                _ = v16
            case 0:
                delete(v1, (fi15 * fi15))
            case 32767:
                v14 += (v14 - v14)
            }
            _ = fi15
        }
        var v17 string = ("foo bar" + "hello")
        _ = v17
    } else {
        for fi18 := 0; fi18 < 14; fi18++ {
            v14 = v14
            v10 += (v10 >> ((0) & 63))
            var v19 string = "hello"
            _ = fi18
            _ = v19
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
}
