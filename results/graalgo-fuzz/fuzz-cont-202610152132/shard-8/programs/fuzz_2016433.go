package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int32) int32 {
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(fi5, p0)
        _ = fi5
    }
    v6 := []int{32767, 2147483647, -803}
    fmt.Println(p0)
    switch "" {
    case "foo bar":
        var v7 bool = true
        _ = v7
    case "":
        p0 += p0
    }
    _ = p0
    _ = v6
    return -919
}

func main() {
    var v1 int = 124
    var v2 int = v1
    var v3 bool = false
    v4 := make(map[uint32]string)
    v8 := GenS1{F0: 65535}
    var v9 float32 = (1e10 / (float32(((-2.25 / float64(0.1)) + (0.1 * ((1.5 + float64(-2.25)) / -2.25)))) + (1.5 + 1e10)))
    var v10 int64 = -47
    var v11 uint = 255
    v3 = v3
    fmt.Println(v9, v10)
    fmt.Println(v10)
    v1 = v2
    var v12 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
