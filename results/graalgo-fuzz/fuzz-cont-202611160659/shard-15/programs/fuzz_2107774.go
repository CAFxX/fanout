package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int32 }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint) uint16 {
    p0 = p0
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 float64 = 3.14159
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: false, F2: 2147483647}
    _ = p0
    _ = v7
    return 230
}

func main() {
    v1 := []int{601, -118, 42}
    v2 := []uint32{}
    v3 := []string{"\n\t"}
    var v4 string = ""
    var v8 string = v4
    var v9 float64 = -2.25
    var v10 float64 = (v9 * v9)
    v11 := GenS2{F0: 0, F1: 2}
    var v12 uint64 = 525
    var v13 int64 = 42
    v13 = v13
    var v14 int = 32768
    v15 := GenS1{F0: 63, F1: (true && ((369 != 1) && ((2147483647 < -1) || true))), F2: 413}
    v4 = (v8 + "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
