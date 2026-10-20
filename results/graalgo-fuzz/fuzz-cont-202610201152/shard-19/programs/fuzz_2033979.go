package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint8 = 128
    var v2 bool = (425 > 2147483647)
    v3 := make(map[int]int)
    var v4 float32 = 0.1
    v5 := []string{"", "", "a", ("\n\t" + "\n\t")}
    switch 2147483647 {
    case 35:
        var v6 string = ""
        _ = v6
    case -9223372036854775808:
        var v7 float32 = 0.0
        _ = v7
    case -584:
        delete(v3, -9223372036854775808)
    default:
        fmt.Println(v1)
    }
    _ = len(v3)
    v2 = v2
    fmt.Println(v4)
    switch 891 {
    case -2147483648:
        v8 := GenS1{F0: (!false)}
        _ = v8
    case 773:
        v4 -= v4
    default:
        var v9 float32 = v4
        _ = v9
    }
    v2 = (-2.25 < v4)
    delete(v3, 128)
    var v10 uint64 = 42
    v11 := GenS2{F0: true}
    v3[1] = 639
    v4 = (-2.25 / (v4 - (v4 / (float32((((0.0 / 3.14159) - -2.25) * 0.1)) / -2.25))))
    v2 = true
    var v12 int64 = 9223372036854775807
    delete(v3, 32767)
    switch 128 {
    case 517:
        v5 = append(v5, "")
    default:
        if (2 < v1) {
            var v13 uint32 = 2
            v2 = (!false)
            _ = v13
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
