package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int8 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 int = -2147483648
    var v5 uint32 = 4294967295
    v6 := []uint32{v5, (v5 % (((v5 | v5)) | 1))}
    fmt.Println(v5, v5)
    if false {
        fmt.Println(v5, v5)
    }
    var v7 int = (v4 % ((-290) | 1))
    _ = len(v3)
    switch ("a" + "foo bar") {
    case "hello":
        v8 := GenS1{F0: 597, F1: false, F2: 0}
        _ = v8
    case "foo bar":
        _ = len(v3)
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        _ = len(v3)
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 float64 = -2.25
            if true {
                _ = len(v1)
                v12 := GenS2{F0: 2147483647}
                var v13 int = -807
                _ = v12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v7 = 32767
    v14 := GenS1{F0: 510, F1: true, F2: -128}
    v5 += 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
