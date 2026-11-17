package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -9223372036854775808
    var v3 int8 = -1
    var v4 uint8 = 126
    v3 = (v3 * 127)
    v5 := v1[("\n\t" + "\n\t")]
    var v6 float32 = float32(1e10)
    v2 += v2
    v2 += v2
    var v7 bool = (0.0 == -2.25)
    var v8 bool = false
    var v9 float32 = ((v6 * 0.1) - v6)
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{410, (-(fi10))}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 187}
    _ = len(v1)
    _ = len(v1)
    v13 := GenS1{F0: 548}
    if (2 < 979) {
        v14 := []int{-9223372036854775808, 4294967295, -2147483648}
        _ = v14
    } else {
        var v15 bool = v5
        _ = v15
    }
    v3 -= (-(v3))
    switch -2147483648 {
    case -1:
        switch ("" + "a") {
        case "a":
            v4 -= 95
        case "hello":
            var v16 uint64 = 2
            _ = v16
        }
    case 128:
        var v17 int16 = -242
        _ = v17
    case 537:
        v4 = v4
    default:
        switch "hello" {
        case "\n\t":
            v18 := GenS1{F0: 0}
            _ = v18
        }
    }
    var v19 uint64 = 1
    var v20 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
