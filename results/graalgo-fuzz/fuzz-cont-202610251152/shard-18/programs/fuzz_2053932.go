package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() bool {
    v4 := []int{32768, 486}
    var v5 int16 = -1
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 int = 127
    var v2 int64 = -292
    v3 := make(map[uint32]int)
    if false {
        v6 := GenS1{F0: ""}
        _ = v6
    }
    v7 := GenS1{F0: "\n\t"}
    var v8 int = (v1 - (v1 & v1))
    for fi9 := 0; fi9 < 2; fi9++ {
        if true {
            switch 1 {
            case 1:
                v2 = -530
            }
            v3[0] = v8
            switch "hello" {
            case "\n\t":
                v10 := v3[251]
                _ = v10
            case "a":
                fmt.Println(v8, fi9, v1)
            }
        } else {
            fmt.Println(v2, v1, v2)
        }
        var v11 float64 = (3.14159 / 1.0)
        _ = fi9
        _ = v11
    }
    fmt.Println(v1)
    fmt.Println(v2, v2, v1)
    fmt.Println(v2)
    var v12 int16 = 293
    v8 = (v1 >> ((2) & 63))
    delete(v3, 2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
