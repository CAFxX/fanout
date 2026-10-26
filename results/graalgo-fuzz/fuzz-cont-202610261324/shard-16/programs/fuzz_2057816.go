package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() string {
    var v4 string = ("\n\t" + "\n\t")
    var v5 uint32 = 2
    switch v4 {
    case "hello":
        fmt.Println(v5)
    case "foo bar":
        var v6 int32 = -423
        _ = v6
    }
    _ = v4
    _ = v5
    return v4
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]int)
    v3 := []int{1}
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int = (-(fi7))
        v8 = (^(v8))
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint = 42
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    _ = len(v2)
    _ = len(v3)
    var v11 string = "hello"
    v12 := GenS1{F0: 0}
    _ = len(v2)
    v11 = "foo bar"
    var v13 float64 = (-2.25 + -2.25)
    fmt.Println(v1)
    for fi14 := 0; fi14 < 19; fi14++ {
        v13 -= v13
        var v15 bool = (float32((0.1 * v13)) >= (-2.25 / (0.0 + (1e10 - float32((v13 + -2.25))))))
        var v16 int8 = 127
        _ = fi14
        _ = v15
        _ = v16
    }
    if (0 != -1) {
        fmt.Println(v11)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
