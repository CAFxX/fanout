package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 909
    v3 := make(map[int]string)
    for fi4 := 0; fi4 < 11; fi4++ {
        switch "a" {
        case "":
            v5 := GenS1{F0: 0, F1: (2 <= 2), F2: 1000}
            _ = v5
        default:
            v6 := GenS1{F0: 127, F1: (false && false), F2: 1}
            _ = v6
        }
        var v7 float32 = 0.1
        _ = fi4
        _ = v7
    }
    fmt.Println(v2)
    fmt.Println(v2)
    v8 := GenS1{F0: 0, F1: true, F2: -1}
    var v9 int16 = -1
    if true {
        var v10 uint64 = 2
        v2 = v2
        for fi11 := 0; fi11 < 2; fi11++ {
            switch ("hello" + "") {
            case "":
                if false {
                    var v12 float32 = -2.25
                    _ = v12
                }
            case "a":
                var v13 int64 = int64(v10)
                _ = v13
            }
            _ = fi11
        }
        _ = v10
    }
    var v14 uint8 = 128
    v14 = 135
    var v15 float64 = float64(3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
