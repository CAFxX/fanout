package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint8 }

func main() {
    var v1 int = 32767
    var v2 string = "foo bar"
    v3 := []string{}
    var v4 bool = false
    var v5 string = ("foo bar" + v2)
    var v6 uint = 0
    var v7 uint32 = 584
    v3 = append(v3, v5)
    if false {
        fmt.Println(v1)
    } else {
        v8 := GenS1{F0: -1}
        if v4 {
            v7 = 935
            v9 := GenS1{F0: -1}
            var v10 float64 = (float64((1e10 + float32(0.0))) + 1.5)
            _ = v10
            _ = v9
        } else {
            switch v1 {
            case 4294967295:
                fmt.Println(v7, v1, v7)
            case 9223372036854775807:
                var v11 string = v5
                _ = v11
            }
            fmt.Println(v5, v1, v4)
        }
        _ = v8
    }
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = ("\n\t" + v2)
    }
    var v12 string = ""
    v13 := GenS2{F0: v1, F1: 2}
    var v14 float64 = (float64(float32(1.5)) / (1e10 - (-2.25 - (-2.25 + (1.5 - float64(0.1))))))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
