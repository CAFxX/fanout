package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 int; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 1.5
    v3 := []string{("foo bar" + ("a" + "foo bar"))}
    var v4 float32 = 1e10
    if false {
        v5 := GenS1{F0: -32768}
        v2 -= (0.1 - v2)
        _ = v5
    }
    v6 := GenS2{F0: 167, F1: "\n\t"}
    var v7 float32 = v4
    v7 = ((((0.1 * v7) - 3.14159) / v4) / float32((0.0 + 3.14159)))
    v4 = ((float32(-2.25) / ((v4 - float32((v2 * v2))) - v4)) - v4)
    if true {
        var v8 uint = 18446744073709551615
        for fi9 := 0; fi9 < 1; fi9++ {
            v7 = v7
            switch "a" {
            case "foo bar":
                v8 -= 5
            case "a":
                v2 += v2
            }
            v10 := GenS3{F0: 1}
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        v11 := GenS4{F0: 0, F1: false}
        var v12 int32 = 42
        _ = v11
        _ = v12
    }
    switch -694 {
    case -118:
        v13 := GenS3{F0: 42}
        _ = v13
    }
    _ = len(v3)
    v14 := GenS5{F0: -290, F1: 0, F2: 2147483647}
    if (!true) {
        var v15 uint8 = 0
        fmt.Println(v2, v7, v15)
        var v16 string = ("hello" + "foo bar")
        _ = v15
        _ = v16
    } else {
        fmt.Println(v2)
        var v17 uint16 = 0
        v7 = v4
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        switch fi18 {
        case 39:
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 uint64 = 18446744073709551615
                _ = fi19
                _ = v20
            }
        }
        var v21 uint = 18446744073709551615
        _ = fi18
        _ = v21
    }
    v22 := GenS3{F0: 343}
    var v23 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
