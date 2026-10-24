package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint8 = 166
    var v2 uint = 238
    var v3 bool = true
    var v4 uint = v2
    v5 := make(map[string]string)
    v6 := make(map[string]int)
    if false {
        switch "foo bar" {
        case "foo bar":
            if (127 >= -7) {
                fmt.Println(v3, v3)
                var v7 int64 = 0
                _ = v7
            }
        case "hello":
            var v8 int32 = -1
            _ = v8
        default:
            var v9 int64 = -1
            _ = v9
        }
    } else {
        var v10 float64 = 3.14159
        _ = v10
    }
    v4 = 18446744073709551615
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: -1, F1: -468}
        _ = fi11
        _ = v12
    }
    v13 := v5["\n\t"]
    var v14 uint8 = (v1 ^ v1)
    fmt.Println(v2)
    v15 := GenS2{F0: 42}
    switch v13 {
    case "\n\t":
        fmt.Println(v14, v14, v13)
    }
    v16 := []int{416}
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := v6[v13]
        fmt.Println(fi17, v3, fi17)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 string = v13
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint = v4
        fmt.Println(v22, v2, fi21)
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
