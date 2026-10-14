package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    var v3 string = ("" + "foo bar")
    v3 = "foo bar"
    fmt.Println(v3)
    var v4 int8 = 25
    for fi5 := 0; fi5 < 14; fi5++ {
        v4 = (-(v4))
        fmt.Println(v4)
        v6 := v2[4294967295]
        _ = fi5
        _ = v6
    }
    if true {
        v4 += v4
    }
    v7 := GenS1{F0: 639, F1: 127, F2: 42}
    v8 := GenS2{F0: 1}
    switch v3 {
    case "a":
        switch v3 {
        case "\n\t":
            v9 := v1[((v3 + (v3 + (((v3 + "\n\t") + ((v3 + v3) + "\n\t")) + ""))) + (v3 + "hello"))]
            _ = v9
        }
    case "hello":
        for fi10 := 0; fi10 < 10; fi10++ {
            fmt.Println(v3, v4, v3)
            var v11 uint64 = 42
            _ = fi10
            _ = v11
        }
    default:
        v3 = (v3 + "\n\t")
    }
    v12 := v2[939]
    v3 = "hello"
    v13 := GenS1{F0: 4294967295, F1: (v4 / (((v4 / (((v4 - v4)) | 1))) | 1)), F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
