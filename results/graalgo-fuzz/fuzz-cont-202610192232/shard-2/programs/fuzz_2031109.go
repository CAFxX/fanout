package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int }

func main() {
    var v1 int16 = 32767
    var v2 int8 = -1
    var v3 bool = false
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    if (124 > 42) {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: v2, F1: 2}
            v8 := []int{9223372036854775807, -1, (fi6 & (fi6 / ((fi6) | 1)))}
            var v9 int = fi6
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        v10 := GenS2{F0: 4294967295}
        _ = v10
    } else {
        if true {
            fmt.Println(v2, v2, v1)
            if (true && v3) {
                var v11 bool = v3
                _ = v11
            } else {
                fmt.Println(v3, v1)
                v12 := []int{}
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 18; fi13++ {
            fmt.Println(v1)
            var v14 uint16 = 393
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                _ = len(v5)
            case "":
                var v15 float64 = ((3.14159 - 1.5) * float64(3.14159))
                _ = v15
            }
            _ = fi13
            _ = v14
        }
    }
    fmt.Println(v1, v3)
    fmt.Println(v1, v2)
    var v16 uint64 = 0
    v17 := []int{-9223372036854775808}
    var v18 uint = 255
    v17 = append(v17, int(v2))
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
