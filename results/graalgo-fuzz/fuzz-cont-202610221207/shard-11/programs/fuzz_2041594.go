package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int64; F2 int32 }
func genFn0(p0 bool, p1 uint8) int32 {
    if true {
        switch ("a" + ((("" + "\n\t") + "\n\t") + "")) {
        case "\n\t":
            p0 = ((p0 || false) || (!true))
        }
        var v4 uint = 42
        _ = v4
    } else {
        var v5 int64 = 42
        _ = v5
    }
    var v6 uint8 = (^(p1))
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 float64 = 0.1
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 657
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 341
    var v3 float32 = (1.5 * 3.14159)
    var v9 uint = 823
    v10 := GenS1{F0: 4294967295, F1: 939, F2: true}
    var v11 float64 = (3.14159 * 0.0)
    v12 := []int{128}
    var v13 string = ((("\n\t" + ("" + "hello")) + ("a" + ("\n\t" + "foo bar"))) + "\n\t")
    fmt.Println(v11)
    v14 := GenS2{F0: false}
    v15 := GenS3{F0: -358, F1: 9223372036854775807, F2: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
