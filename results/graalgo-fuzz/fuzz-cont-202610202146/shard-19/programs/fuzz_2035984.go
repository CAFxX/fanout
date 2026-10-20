package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{}
    var v3 string = ("\n\t" + "hello")
    var v4 uint = 27
    var v5 uint16 = 42
    var v6 float32 = 0.0
    var v7 string = (v3 + "\n\t")
    for fi8 := 0; fi8 < 5; fi8++ {
        switch -627 {
        case 32768:
            if true {
                v9 := GenS1{F0: -289, F1: 236, F2: -9223372036854775808}
                _ = v9
            } else {
                v10 := v1[v3]
                v7 = v7
                if (0 >= 2) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = v10
            }
        default:
            var v11 uint8 = 1
            _ = v11
        }
        v4 = v4
        _ = fi8
    }
    fmt.Println(v4)
    fmt.Println(v5)
    if len(v2) > 0 {
        v2[0] = int(v5)
    }
    var v12 uint16 = v5
    var v13 uint64 = 1
    v4 = v4
    v14 := GenS1{F0: 0, F1: 1, F2: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
