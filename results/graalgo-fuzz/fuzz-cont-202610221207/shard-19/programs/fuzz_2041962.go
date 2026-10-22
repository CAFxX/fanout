package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint16 = 201
    v2 := []uint32{4294967295, 577}
    var v3 string = (("a" + ("\n\t" + (("" + "a") + "a"))) + "")
    v4 := make(map[string]int)
    var v5 string = "a"
    v6 := []string{v3, (v3 + "a"), v5, v5}
    v1 -= 42
    v3 = ""
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(v1, v3, v1)
        v8 := v4[("a" + (v5 + "a"))]
        _ = fi7
        _ = v8
    }
    if (42 <= -128) {
        fmt.Println(v1)
    } else {
        v9 := v4[v5]
        _ = v9
    }
    var v10 uint = 255
    v10 = 42
    if (-2147483648 >= 0) {
        switch -212 {
        case 4294967295:
            v11 := GenS1{F0: v3}
            _ = v11
        default:
            for fi12 := 0; fi12 < 14; fi12++ {
                v5 = v5
                _ = fi12
            }
        }
    }
    v4[(v5 + v5)] = -655
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
