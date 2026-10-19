package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int }
func genFn0(p0 int16) int16 {
    var v5 uint16 = 32767
    fmt.Println(v5, p0)
    _ = p0
    _ = v5
    return p0
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 127
    v3 := make(map[int]int)
    var v4 bool = false
    if (!v4) {
        for fi6 := 0; fi6 < 1; fi6++ {
            v3[(v2 | 128)] = v2
            _ = fi6
        }
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 string = v1
            if ((true && ((!(-2147483648 <= 0)) || false)) || v4) {
                v1 = "foo bar"
                var v9 uint64 = 2
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        var v10 uint8 = 0
        _ = v10
    }
    if ((v4 && (!false)) && v4) {
        var v11 uint16 = 1
        v12 := GenS1{F0: 42}
        _ = v11
        _ = v12
    }
    v13 := []int{v2}
    v14 := GenS2{F0: v2, F1: v2}
    var v15 uint64 = 0
    var v16 int32 = 1
    v13 = append(v13, v2)
    v17 := GenS1{F0: 1}
    if v4 {
        fmt.Println(v2, v1)
    }
    fmt.Println(v2, v15, v1)
    v1 = v1
    var v18 int32 = v16
    var v19 bool = (-2147483648 == 0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
