package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 bool = true
    var v2 uint8 = 23
    v3 := make(map[uint32]string)
    v4 := make(map[string]string)
    v5 := make(map[int]bool)
    switch 32768 {
    case 0:
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := v3[2147483648]
            switch fi6 {
            case 670:
                switch "foo bar" {
                case "\n\t":
                    v8 := GenS1{F0: 1}
                    _ = v8
                default:
                    if v1 {
                        var v9 int8 = 5
                        v10 := GenS1{F0: 2}
                        _ = v10
                        _ = v9
                    }
                }
            }
            _ = fi6
            _ = v7
        }
    }
    v11 := []int{128, 2147483647, -738, 0}
    v2 += 127
    _ = len(v5)
    v4["foo bar"] = "a"
    v2 -= (^(v2))
    var v12 int8 = -1
    v12 += (v12 | 1)
    if v1 {
        v2 = v2
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    } else {
        v13 := GenS1{F0: 538}
        _ = len(v5)
        if len(v11) > 0 {
            v11[0] = -1
        }
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v1, fi14, v1)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
