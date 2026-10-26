package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []string{}
    v2 := []int{-858, -215, 128}
    var v3 float32 = 1e10
    var v4 int32 = -439
    v5 := make(map[uint32]int)
    var v6 int64 = 9223372036854775807
    var v7 float64 = ((0.0 * ((0.0 * 1.5) + float64((((v3 * v3) + v3) + (0.1 + 3.14159))))) * 1e10)
    var v8 uint8 = 167
    fmt.Println(v3)
    var v9 uint8 = 2
    for fi10 := 0; fi10 < 20; fi10++ {
        v4 += v4
        v11 := GenS1{F0: ((!false) || false)}
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 uint32 = 4294967295
            _ = len(v5)
            v9 = (^(v9))
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v3 = v3
    v14 := GenS1{F0: ("a" > ("foo bar" + ("foo bar" + ("" + "foo bar"))))}
    v9 = 0
    var v15 int32 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
