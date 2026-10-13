package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int16 = 50
    var v2 uint16 = 190
    var v3 string = "a"
    v4 := make(map[string]uint32)
    var v5 int8 = 42
    v6 := make(map[int]uint32)
    v7 := []int{128}
    v8 := GenS1{F0: 1, F1: 18446744073709551615}
    var v9 int32 = 42
    var v10 int64 = -9223372036854775808
    if (false && (!((((float64(0.0) + 1.5) * 3.14159) + 3.14159) >= 0.0))) {
        fmt.Println(v1, v5, v10)
    }
    var v11 uint16 = 65535
    if false {
        for fi12 := 0; fi12 < 2; fi12++ {
            if (false && false) {
                var v13 int64 = (v10 ^ (^(v10)))
                var v14 uint64 = 18446744073709551615
                _ = v13
                _ = v14
            } else {
                var v15 uint = 65535
                _ = v15
            }
            v16 := GenS1{F0: (^(v9)), F1: 79}
            if true {
                v17 := GenS2{F0: v9}
                fmt.Println(v5, v3, v10)
                switch v3 {
                case "":
                    var v18 int8 = 0
                    _ = v18
                }
                _ = v17
            } else {
                v5 += int8(v10)
                v5 = 100
            }
            _ = fi12
            _ = v16
        }
        v2 = 824
    } else {
        var v19 uint32 = 136
        var v20 string = ("hello" + v3)
        switch 32768 {
        case -2147483648:
            _ = len(v7)
        case -9223372036854775808:
            var v21 int16 = (v1 & 42)
            _ = v21
        default:
            var v22 string = ""
            _ = v22
        }
        _ = v19
        _ = v20
    }
    var v23 uint32 = 1
    var v24 string = v3
    v5 += v5
    var v25 float64 = 1e10
    switch ("a" + v24) {
    case "foo bar":
        if len(v7) > 0 {
            v7[0] = -1
        }
    case "\n\t":
        v26 := GenS2{F0: 1}
        _ = v26
    }
    if (((166 >= 2) && false) && false) {
        var v27 uint16 = v11
        var v28 uint64 = 0
        var v29 uint64 = v28
        _ = v27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
