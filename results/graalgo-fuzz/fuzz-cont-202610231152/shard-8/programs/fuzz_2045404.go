package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int8) int {
    fmt.Println(p0)
    p0 = -1
    var v6 float32 = float32(1.5)
    fmt.Println(p0, p0)
    if false {
        p0 -= 0
    } else {
        v6 = v6
    }
    _ = p0
    _ = v6
    return 0
}
func genFn1(p0 uint64, p1 uint16, p2 int8) bool {
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(fi7, p0)
        v8 := []int{(fi7 % ((fi7) | 1))}
        _ = fi7
        _ = v8
    }
    var v9 int16 = 42
    v9 += v9
    v10 := []int{0, -140, 308}
    v9 = (-(v9))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (("hello" + "") < "\n\t")
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 uint64 = 4294967295
    v5 := make(map[string]uint32)
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := []int{fi11, 843}
        var v13 float64 = float64(1.5)
        fmt.Println(v13, v13)
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v4)
    switch "a" {
    case "":
        v4 -= (v4 * v4)
    case "a":
        v14 := GenS1{F0: (v4 % ((v4) | 1)), F1: true, F2: -830}
        _ = v14
    }
    var v15 uint8 = 117
    v16 := GenS2{F0: 100}
    var v17 string = ""
    v15 = (v15 % ((v15) | 1))
    for fi18 := 0; fi18 < 6; fi18++ {
        fmt.Println(v17, v4)
        _ = fi18
    }
    v19 := GenS2{F0: 100}
    v3[v17] = 834
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
