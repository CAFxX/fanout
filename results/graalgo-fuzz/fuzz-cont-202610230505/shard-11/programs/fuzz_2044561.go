package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 string }

func main() {
    v1 := []int{}
    v2 := []int{}
    v3 := make(map[string]uint32)
    var v4 int = 396
    v5 := make(map[string]string)
    fmt.Println(v4)
    delete(v3, "")
    fmt.Println(v4)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 string = "foo bar"
        switch v7 {
        case "":
            if (false && (91 < 0)) {
                var v8 uint32 = 2
                var v9 int8 = 100
                var v10 string = v7
                _ = v10
                _ = v8
                _ = v9
            }
        case "\n\t":
            switch 127 {
            case 9223372036854775807:
                v11 := GenS1{F0: "foo bar", F1: 65535, F2: v7}
                _ = v11
            case 466:
                var v12 uint64 = 42
                _ = v12
            case 127:
                var v13 string = v7
                _ = v13
            }
        case "a":
            if (!false) {
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 int = (fi6 + -9223372036854775808)
                    _ = fi14
                    _ = v15
                }
                fmt.Println(fi6)
            } else {
                for fi16 := 0; fi16 < 9; fi16++ {
                    v4 = v4
                    _ = fi16
                }
                v4 = -585
            }
        default:
            fmt.Println(fi6, v7, fi6)
        }
        _ = fi6
        _ = v7
    }
    var v17 int16 = 1000
    fmt.Println(v4, v17)
    if false {
        var v18 string = "a"
        _ = v18
    } else {
        v4 -= v4
        v4 = (v4 / (((v4 - (v4 * (v4 + (v4 | v4))))) | 1))
    }
    var v19 float32 = (0.0 * float32(3.14159))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
