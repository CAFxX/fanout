package main

import (
    "fmt"
)

func genFn0(p0 uint32) uint8 {
    switch "hello" {
    case "\n\t":
        p0 = (p0 * 65535)
    }
    var v7 uint = 2
    _ = p0
    _ = v7
    return 2
}

func main() {
    v1 := []string{"\n\t", "\n\t", ""}
    v2 := make(map[uint32]int)
    var v3 bool = (42 >= 2)
    var v4 uint8 = 50
    v5 := make(map[uint32]string)
    var v6 int8 = 78
    var v8 uint = 255
    _ = len(v5)
    var v9 float64 = float64((-2.25 / 1.5))
    v6 -= (v6 << ((v8) & 63))
    v3 = false
    _ = len(v2)
    if v3 {
        v8 = 259
        var v10 int8 = v6
        v3 = ((!(v9 > ((((float64(1.5) * float64(-2.25)) - -2.25) * (1e10 / ((v9 - v9) / v9))) + ((((((v9 / float64((1.5 / ((0.0 * ((0.1 + float32(1e10)) - 1.5)) + 0.1)))) * (-2.25 / v9)) * v9) / v9) + v9) - -2.25)))) && (!false))
        _ = v10
    }
    v4 = v4
    var v11 uint16 = 1
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
