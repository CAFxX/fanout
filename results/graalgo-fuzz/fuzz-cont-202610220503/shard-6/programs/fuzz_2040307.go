package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 string = "hello"
    var v2 float32 = ((1e10 + (1.5 - 0.0)) - (float32(0.1) - 1.5))
    var v3 int8 = 1
    var v4 int64 = 9223372036854775807
    var v5 uint64 = 1
    v6 := make(map[uint32]int)
    if (false || true) {
        switch v1 {
        case "hello":
            v7 := []int{128}
            _ = v7
        case "foo bar":
            v1 = v1
        }
        v8 := GenS1{F0: 4294967295}
        if true {
            fmt.Println(v3)
            v4 -= 337
        }
        _ = v8
    }
    var v9 uint64 = v5
    var v10 bool = (194 >= -2147483648)
    v11 := []int{128, 42}
    v2 = v2
    v12 := GenS1{F0: 375}
    var v13 string = (("a" + "hello") + "foo bar")
    v2 -= (float32(0.1) * v2)
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 bool = v10
        _ = fi14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
