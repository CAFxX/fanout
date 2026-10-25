package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint32 = 569
    var v2 uint16 = 255
    var v3 float32 = 1e10
    var v4 uint = 1
    var v5 uint = 717
    var v6 float32 = 1.5
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: fi7}
        _ = fi7
        _ = v8
    }
    v9 := []int{0}
    var v10 uint = 65535
    switch "foo bar" {
    case "":
        switch "a" {
        case "a":
            var v11 uint16 = 305
            _ = v11
        }
    case "foo bar":
        v12 := GenS1{F0: 2147483647}
        _ = v12
    }
    fmt.Println(v10)
    v10 += 4294967295
    v1 -= v1
    var v13 uint16 = 0
    var v14 uint32 = v1
    switch "a" {
    case "a":
        if len(v9) > 0 {
            v9[0] = -2147483648
        }
    case "":
        var v15 float64 = 1.5
        _ = v15
    case "foo bar":
        var v16 float64 = (float64(((v3 * 1.5) + v3)) / 0.1)
        _ = v16
    }
    v14 += 95
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
