package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 bool }

func main() {
    var v1 uint = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 int8 = 100
    var v5 int = 128
    v6 := []uint32{2}
    var v7 int8 = 127
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 string = "foo bar"
        if (true && true) {
            var v10 uint8 = 128
            v4 = (v7 | v7)
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    if true {
        v4 = -128
    }
    v7 -= (v4 | -128)
    v11 := GenS1{F0: (v7 % (((v4 - (v4 >> ((0) & 63)))) | 1)), F1: 0, F2: (false || false)}
    _ = len(v3)
    switch ("hello" + ("" + ("foo bar" + ("hello" + "\n\t")))) {
    case "hello":
        v5 = -1
    case "foo bar":
        v5 = (v5 + v5)
    }
    delete(v3, "hello")
    v5 += v5
    if (false || true) {
        fmt.Println(v5, v7, v7)
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS1{F0: 72, F1: 1, F2: (true && false)}
            v14 := GenS1{F0: (v4 - v4), F1: 0, F2: false}
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 4; fi15++ {
            v1 = 18446744073709551615
            v16 := GenS1{F0: (v7 | 127), F1: 0, F2: false}
            _ = fi15
            _ = v16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
