package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []uint32{1, 1, 0}
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    v4 := []string{"foo bar", "\n\t"}
    var v5 int = 128
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if ("" >= "hello") {
        v4 = append(v4, ("\n\t" + "foo bar"))
    } else {
        var v6 int = v5
        if false {
            fmt.Println(v5, v6)
            for fi7 := 0; fi7 < 3; fi7++ {
                switch "" {
                case "a":
                    if false {
                        var v8 int8 = -128
                        _ = v8
                    }
                }
                _ = fi7
            }
        }
        _ = v6
    }
    var v9 uint64 = 1
    var v10 int32 = -1
    if false {
        var v11 uint = 314
        v12 := GenS1{F0: v11}
        _ = v11
        _ = v12
    } else {
        switch -741 {
        case 32768:
            fmt.Println(v5, v9)
        case 127:
            fmt.Println(v9)
        default:
            v13 := GenS1{F0: 604}
            _ = v13
        }
        switch "a" {
        case "":
            v3[42] = ((true || false) || true)
        }
        v14 := GenS1{F0: 4294967295}
        _ = v14
    }
    _ = len(v3)
    fmt.Println(v5)
    v15 := GenS2{F0: 156}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
