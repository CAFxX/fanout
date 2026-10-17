package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := []int{32767}
    v2 := make(map[string]int)
    var v3 int32 = 2147483647
    v4 := []string{"", "\n\t"}
    v4 = append(v4, ("hello" + "\n\t"))
    var v5 int64 = 42
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: 363, F1: 255}
    switch "a" {
    case "a":
        v8 := GenS1{F0: 621, F1: 81}
        _ = v8
    case "\n\t":
        fmt.Println(v5, v3)
    case "foo bar":
        if false {
            for fi9 := 0; fi9 < 4; fi9++ {
                for fi10 := 0; fi10 < 11; fi10++ {
                    if false {
                        v3 = (v3 << ((123) & 63))
                        fmt.Println(v3, v5, v5)
                        var v11 int32 = -696
                        _ = v11
                    }
                    delete(v2, "a")
                    v3 = v3
                    _ = fi10
                }
                _ = fi9
            }
        } else {
            var v12 int = 2147483647
            v13 := GenS1{F0: 2, F1: 255}
            _ = v12
            _ = v13
        }
    }
    v14 := GenS1{F0: 255, F1: 42}
    v1 = append(v1, 663)
    var v15 float32 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
