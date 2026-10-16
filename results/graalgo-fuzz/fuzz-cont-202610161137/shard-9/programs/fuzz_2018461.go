package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
func genFn0(p0 uint32, p1 uint64, p2 string) bool {
    var v4 uint8 = 130
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := []int{fi5, 304}
        p2 = p2
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: ((p2 + "foo bar") + (p2 + "a"))}
    var v8 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    _ = v8
    return (true && true)
}

func main() {
    var v1 uint32 = 543
    v2 := make(map[string]string)
    v3 := []uint32{v1, (v1 * v1), 1, (^(v1))}
    var v9 uint64 = 1
    switch "a" {
    case "\n\t":
        v2["hello"] = "a"
    }
    var v10 int16 = -1
    if ("" == ("hello" + "\n\t")) {
        var v11 int = -9223372036854775808
        _ = v11
    } else {
        for fi12 := 0; fi12 < 2; fi12++ {
            fmt.Println(v1, v9, v1)
            var v13 float32 = ((1e10 + (1.5 - 1.5)) / -2.25)
            switch ((("a" + "foo bar") + "hello") + "hello") {
            case "hello":
                for fi14 := 0; fi14 < 19; fi14++ {
                    v13 = v13
                    _ = fi14
                }
            case "":
                v9 = (v9 * v9)
            case "foo bar":
                v1 += 488
            default:
                var v15 float64 = (3.14159 * 3.14159)
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        fmt.Println(v10)
    }
    var v16 int64 = -868
    _ = len(v3)
    fmt.Println(v16)
    fmt.Println(v1, v9, v10)
    fmt.Println(v16)
    var v17 int = 2147483647
    var v18 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
