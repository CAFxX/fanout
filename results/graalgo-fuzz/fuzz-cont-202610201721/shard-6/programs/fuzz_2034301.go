package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint, p1 int, p2 int16) int32 {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int8 = 127
        _ = fi7
        _ = v8
    }
    v9 := []int{-266}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint64 = 0
    var v2 int64 = 278
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    v5 := []string{("" + "\n\t"), "", ""}
    v6 := []int{-713, 593}
    v2 -= v2
    var v10 uint8 = 42
    var v11 int64 = (v2 << ((18446744073709551615) & 63))
    v10 = 255
    if ((("a" > "foo bar") && (!(!false))) || (255 > (v10 - (v10 - 127)))) {
        for fi12 := 0; fi12 < 12; fi12++ {
            v1 -= (v1 & v1)
            _ = fi12
        }
        var v13 int8 = 100
        _ = v13
    }
    var v14 int64 = -539
    var v15 float64 = ((float64(((1.5 + 0.0) / (0.1 / 3.14159))) - 1e10) * 0.1)
    v16 := GenS1{F0: (v2 << ((2) & 63))}
    v17 := v4[732]
    for fi18 := 0; fi18 < 1; fi18++ {
        fmt.Println(v17, v15)
        var v19 string = "a"
        v11 += 1
        _ = fi18
        _ = v19
    }
    v2 += v11
    fmt.Println(v10)
    var v20 int32 = -1
    if true {
        v14 = (v14 / ((v14) | 1))
    }
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
