package main

import (
    "fmt"
)


func main() {
    var v1 string = "\n\t"
    var v2 uint8 = 42
    var v3 float64 = 0.1
    v4 := []int{120, 123, -695}
    v1 = v1
    fmt.Println(v1)
    v3 = v3
    var v5 int32 = 0
    var v6 int32 = 144
    var v7 uint32 = 805
    var v8 string = v1
    var v9 int = -9223372036854775808
    for fi10 := 0; fi10 < 14; fi10++ {
        switch fi10 {
        case 840:
            var v11 int8 = -32
            _ = v11
        }
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
