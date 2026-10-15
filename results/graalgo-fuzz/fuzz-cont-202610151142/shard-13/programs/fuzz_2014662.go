package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []int{2147483647, -646}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    switch 42 {
    case -745:
        for fi4 := 0; fi4 < 12; fi4++ {
            var v5 string = "a"
            _ = fi4
            _ = v5
        }
    case 32768:
        for fi6 := 0; fi6 < 3; fi6++ {
            fmt.Println(fi6)
            fmt.Println(fi6)
            if false {
                var v7 uint32 = 868
                _ = v7
            }
            _ = fi6
        }
    case 128:
        var v8 int8 = 0
        _ = v8
    }
    v1 = append(v1, 127)
    var v9 int32 = 1
    v1 = append(v1, 32768)
    var v10 uint16 = 42
    var v11 uint16 = 2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        switch "hello" {
        case "":
            fmt.Println(v10, v11, v11)
        case "\n\t":
            fmt.Println(v9, v11)
        default:
            v12 := GenS1{F0: 0}
            _ = v12
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
