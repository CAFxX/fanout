package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int8 = -45
    var v2 uint64 = 0
    v3 := make(map[string]string)
    if true {
        switch "hello" {
        case "hello":
            v2 = 42
        case "\n\t":
            var v4 uint64 = 42
            _ = v4
        }
        switch 4294967295 {
        case 9223372036854775807:
            v5 := v3["\n\t"]
            _ = v5
        case -1:
            v2 -= 1
        case -878:
            var v6 int64 = -1
            _ = v6
        default:
            v7 := []int{}
            _ = v7
        }
        fmt.Println(v1, v2)
    }
    fmt.Println(v2, v1)
    var v8 uint = 0
    fmt.Println(v2, v1, v8)
    var v9 uint = 733
    fmt.Println(v8, v9, v1)
    var v10 int32 = 331
    var v11 int16 = -143
    if false {
        switch 32768 {
        case -915:
            v12 := GenS1{F0: 127}
            _ = v12
        case 4294967295:
            if true {
                var v13 int16 = -1
                _ = v13
            } else {
                v2 = 357
            }
        case 32768:
            var v14 int32 = (v10 % (((v10 >> (((^(v8))) & 63))) | 1))
            _ = v14
        default:
            v15 := []int{-617, 32767, -1, int(v10)}
            _ = v15
        }
    } else {
        var v16 float64 = 1e10
        _ = v16
    }
    v2 = v2
    var v17 uint64 = 42
    switch -531 {
    case -157:
        v11 = 0
    }
    if true {
        var v18 int8 = (v1 % (((^(v1))) | 1))
        var v19 string = ("hello" + ((("a" + "") + "foo bar") + "\n\t"))
        _ = v18
        _ = v19
    }
    _ = len(v3)
    v10 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
