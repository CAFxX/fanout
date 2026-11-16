package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 2
    v3 := []string{"a", "", "foo bar", "hello"}
    v4 := make(map[string]bool)
    v5 := []int{32767, 9223372036854775807, 1}
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int8 = -1
            _ = fi8
            _ = v9
        }
        v2 = v2
        _ = fi7
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v2 += 631
        _ = fi10
    }
    var v11 float32 = 0.0
    delete(v1, ("" + "foo bar"))
    fmt.Println(v2)
    var v12 float32 = float32((((1.5 / 1.0) + 1e10) + 0.0))
    var v13 string = "hello"
    var v14 uint16 = v2
    var v15 int8 = 76
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 uint = 255
        var v18 string = (v13 + "hello")
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
