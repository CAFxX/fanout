package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 bool }

func main() {
    var v1 string = "\n\t"
    var v2 bool = false
    v3 := make(map[int]int)
    fmt.Println(v1, v2)
    v1 = ("a" + v1)
    v1 = "a"
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 uint64 = 42
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint16 = 75
            var v8 int8 = 0
            if (false || ((v7 ^ v7) != v7)) {
                switch ("" + ("hello" + v1)) {
                case "":
                    var v9 string = v1
                    _ = v9
                case "\n\t":
                    delete(v3, (fi6 + 863))
                default:
                    var v10 uint = 42
                    _ = v10
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi4
        _ = v5
    }
    v3[2147483647] = -101
    v11 := []int{127, 127}
    v12 := GenS1{F0: 71, F1: -161, F2: (1 != 255)}
    fmt.Println(v2, v1)
    v2 = (!(255 <= 255))
    v13 := GenS1{F0: 1, F1: -1, F2: false}
    v14 := GenS1{F0: 245, F1: -670, F2: v2}
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
