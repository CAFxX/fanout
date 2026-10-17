package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint64, p2 uint16) int16 {
    var v4 int64 = -871
    v5 := GenS1{F0: 128}
    var v6 int64 = (v4 ^ -548)
    if true {
        v7 := []int{2147483647, 127}
        _ = v7
    }
    var v8 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    return -1
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]int)
    var v3 float64 = v1
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint32 = 42
        _ = fi9
        _ = v10
    }
    v11 := v2[225]
    v12 := GenS1{F0: 206}
    fmt.Println(v3)
    if true {
        var v13 uint16 = 32767
        _ = v13
    }
    fmt.Println(v3, v3, v3)
    var v14 string = ("\n\t" + "foo bar")
    v3 += (float64((1.5 + ((-2.25 * 0.1) + ((1e10 - (((1.5 - 0.1) * 0.1) * float32(v3))) / -2.25)))) * v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
}
