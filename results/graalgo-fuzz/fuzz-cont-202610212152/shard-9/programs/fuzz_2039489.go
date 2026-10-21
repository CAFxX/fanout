package main

import (
    "fmt"
)


func main() {
    var v1 int32 = 0
    var v2 int32 = v1
    var v3 uint16 = 255
    v4 := make(map[string]int)
    v5 := []uint32{65535, 2}
    var v6 float64 = float64((3.14159 + 0.0))
    v3 = v3
    fmt.Println(v1, v6)
    var v7 float32 = ((-2.25 / (1e10 + (0.0 - -2.25))) * 1e10)
    var v8 bool = true
    var v9 uint8 = 0
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 uint16 = v3
            v1 = -2147483648
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    var v13 uint = 42
    var v14 int = -1
    v9 -= 255
    if v8 {
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
