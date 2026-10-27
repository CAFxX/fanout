package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint }

func main() {
    v1 := []int{-1, 42, 32768}
    v2 := []string{"a", "foo bar", "a"}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    if false {
        var v5 int = 128
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 bool = false
            v2 = append(v2, "a")
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi6
            _ = v7
        }
        v5 = (-(v5))
        _ = v5
    } else {
        delete(v4, "hello")
        var v8 uint16 = 140
        v8 = (v8 + v8)
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: 1, F1: (-2.25 <= (0.1 * ((3.14159 - 0.0) + -2.25))), F2: 0}
        var v11 uint32 = 2147483647
        switch "hello" {
        case "foo bar":
            switch "\n\t" {
            case "a":
                v11 += 42
            case "foo bar":
                _ = len(v3)
            }
        case "hello":
            var v12 int8 = -52
            _ = v12
        case "a":
            v11 = 0
        default:
            if true {
                v13 := v4[("a" + "foo bar")]
                if false {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v13 = (fi9 | v13)
                        _ = fi14
                    }
                    v2 = append(v2, ("\n\t" + "hello"))
                    v4[""] = -803
                }
                _ = v13
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v15 float32 = (1.5 - 0.1)
    v16 := v3[4294967295]
    fmt.Println(v16, v15)
    var v17 float32 = (v15 / v15)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v15, v15, v16)
    var v18 uint32 = 742
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
