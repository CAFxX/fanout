package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int16 = 1000
    var v2 bool = (32 >= 2)
    v3 := make(map[uint32]int)
    v4 := []int{}
    var v5 int32 = 42
    v6 := make(map[int]bool)
    if v2 {
        v1 -= v1
        if (728 != 204) {
            v6[32767] = v2
            var v7 int64 = 684
            _ = v7
        } else {
            switch "" {
            case "foo bar":
                fmt.Println(v1, v2)
            case "a":
                fmt.Println(v1)
            case "\n\t":
                var v8 int8 = 42
                _ = v8
            default:
                v9 := GenS1{F0: 55}
                _ = v9
            }
            v2 = false
            var v10 int64 = 0
            _ = v10
        }
        fmt.Println(v5, v5)
    }
    fmt.Println(v2, v5)
    var v11 uint32 = 4294967295
    v12 := GenS1{F0: 65535}
    v2 = ((-2147483648 <= (v5 >> ((181) & 63))) && v2)
    var v13 uint32 = v11
    if (!(v2 || (v2 && ((3.14159 / float64(((1e10 * (-2.25 + 0.1)) - (1.5 * 1e10)))) > 0.0)))) {
        switch "hello" {
        case "a":
            var v14 uint = 0
            _ = v14
        case "hello":
            for fi15 := 0; fi15 < 5; fi15++ {
                _ = len(v4)
                v16 := GenS1{F0: 42}
                _ = fi15
                _ = v16
            }
        case "\n\t":
            switch "a" {
            case "foo bar":
                if v2 {
                    v2 = (4294967295 < 88)
                } else {
                    var v17 int16 = (v1 * v1)
                    v6[4294967295] = (0 <= 4294967295)
                    v2 = (1.5 < 3.14159)
                    _ = v17
                }
            }
        }
        var v18 int32 = v5
        _ = v18
    } else {
        var v19 float32 = 0.0
        _ = v19
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
