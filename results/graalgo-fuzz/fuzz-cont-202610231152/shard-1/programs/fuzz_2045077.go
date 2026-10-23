package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []int{281, 396, 852}
    var v2 uint16 = 255
    v3 := make(map[string]int)
    fmt.Println(v2)
    v2 += (v2 << ((42) & 63))
    v2 -= 119
    var v4 uint16 = (^(v2))
    v4 += 42
    v5 := GenS1{F0: 196}
    v2 = v2
    v6 := GenS1{F0: 2}
    var v7 uint8 = 29
    var v8 float32 = 1.5
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint = 42
        _ = fi9
        _ = v10
    }
    v2 = 1
    var v11 string = ""
    var v12 string = (("\n\t" + "\n\t") + ((v11 + ((v11 + "\n\t") + v11)) + "a"))
    v13 := v3[("foo bar" + "")]
    v12 = (v11 + v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
