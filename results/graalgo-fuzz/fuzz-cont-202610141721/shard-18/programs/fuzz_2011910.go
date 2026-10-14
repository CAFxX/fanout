package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 string, p1 int8, p2 uint8) uint32 {
    var v6 string = p0
    fmt.Println(p0, p2, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483648
}

func main() {
    var v1 int = 9223372036854775807
    var v2 int8 = 127
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    var v5 int8 = v2
    fmt.Println(v1)
    var v7 uint16 = 42
    fmt.Println(v2)
    v8 := GenS1{F0: 42}
    var v9 uint = 42
    var v10 string = "\n\t"
    var v11 uint32 = 953
    v11 = 42
    v7 = (v7 - 1)
    v2 += -91
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(v9)
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
