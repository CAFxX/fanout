package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint32, p1 int16) uint16 {
    if true {
        for fi6 := 0; fi6 < 6; fi6++ {
            p0 = p0
            _ = fi6
        }
        v7 := []int{0, 939, 80}
        _ = v7
    } else {
        v8 := []int{-1, 153, 128, -9223372036854775808}
        v8 = append(v8, -9223372036854775808)
        _ = len(v8)
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS1{F0: (0.0 > float64(((-2.25 * 0.0) - (0.0 - 0.1))))}
            p0 += p0
            _ = fi10
            _ = v11
        }
        fmt.Println(fi9)
        _ = fi9
    }
    _ = p0
    _ = p1
    return 725
}

func main() {
    v1 := []uint32{2, 187, 42}
    var v2 uint16 = 187
    var v3 uint8 = 255
    v4 := make(map[int]string)
    v5 := make(map[uint32]int)
    for fi12 := 0; fi12 < 17; fi12++ {
        _ = len(v1)
        _ = fi12
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    if (340 >= -32768) {
        var v13 uint = 2
        v14 := GenS1{F0: (((("a" > "a") || ((!true) && true)) || true) && false)}
        _ = v13
        _ = v14
    }
    v3 += (v3 * (^(v3)))
    var v15 int64 = 9223372036854775807
    var v16 int = int(v2)
    var v17 string = "foo bar"
    var v18 int8 = -99
    var v19 int64 = (v15 / ((v15) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
