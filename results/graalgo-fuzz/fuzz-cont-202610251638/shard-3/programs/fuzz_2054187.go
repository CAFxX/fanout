package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 0
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    var v5 uint64 = 1
    var v6 uint16 = 0
    if (!(((v2 + (v2 % (((v2 << ((65535) & 63))) | 1))) <= v2) && true)) {
        fmt.Println(v6, v5, v6)
        v6 = 594
        var v7 uint32 = 42
        _ = v7
    } else {
        var v8 uint = 18446744073709551615
        fmt.Println(v6)
        _ = v8
    }
    if (-1 <= 765) {
        fmt.Println(v2, v2, v5)
    } else {
        var v9 uint8 = 255
        v4["a"] = "\n\t"
        var v10 uint32 = uint32(v6)
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        if true {
            v12 := []int{(fi11 ^ (fi11 % ((fi11) | 1))), 4294967295}
            var v13 int8 = 120
            v2 = (v2 & 2147483647)
            _ = v12
            _ = v13
        } else {
            var v14 float64 = (-2.25 / 1.0)
            v15 := GenS1{F0: -128, F1: 0}
            _ = v14
            _ = v15
        }
        switch (-(fi11)) {
        case 42:
            fmt.Println(fi11, v2, v6)
        }
        _ = fi11
    }
    switch 779 {
    case 32768:
        var v16 float32 = -2.25
        _ = v16
    }
    var v17 float64 = 3.14159
    v18 := GenS1{F0: 28, F1: -9223372036854775808}
    var v19 uint = 42
    switch 127 {
    case -707:
        v20 := GenS2{F0: "\n\t", F1: -1}
        _ = v20
    case 4294967295:
        v17 -= v17
    }
    v17 = v17
    switch "\n\t" {
    case "a":
        v21 := GenS3{F0: 42}
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
