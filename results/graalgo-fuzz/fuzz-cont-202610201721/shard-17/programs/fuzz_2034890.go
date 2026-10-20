package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    var v4 int8 = 122
    switch "foo bar" {
    case "\n\t":
        v5 := GenS1{F0: 453}
        _ = v5
    case "foo bar":
        var v6 uint16 = 80
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int8 = (v4 + v4)
            v10 := v3["foo bar"]
            var v11 uint8 = 127
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi7
    }
    v3["a"] = true
    fmt.Println(v1, v1)
    if (true && false) {
        switch ("hello" + "foo bar") {
        case "":
            v1 -= 18446744073709551615
        case "hello":
            v1 = (v1 ^ v1)
        default:
            if (115 >= 126) {
                if (!false) {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        v13 := GenS1{F0: 0}
                        v14 := []int{}
                        var v15 int32 = -1
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    var v16 uint = 1
                    _ = v16
                } else {
                    switch 1 {
                    case 4294967295:
                        v17 := GenS2{F0: v4}
                        _ = v17
                    case -1:
                        var v18 uint64 = v1
                        _ = v18
                    default:
                        fmt.Println(v1, v4)
                    }
                }
            } else {
                var v19 float32 = ((((0.0 / -2.25) / -2.25) - 0.1) / -2.25)
                _ = v19
            }
        }
    } else {
        var v20 uint64 = (v1 ^ v1)
        v20 = 983
        _ = v20
    }
    fmt.Println(v4, v1)
    fmt.Println(v1, v4)
    v4 = (^(v4))
    v21 := GenS1{F0: 2147483648}
    v22 := []int{-2147483648, -862, -9223372036854775808}
    var v23 uint64 = (v1 - 0)
    v4 += 1
    for fi24 := 0; fi24 < 5; fi24++ {
        v23 -= 42
        _ = fi24
    }
    switch "\n\t" {
    case "a":
        if true {
            _ = len(v22)
            var v25 int32 = 42
            _ = v25
        }
    case "foo bar":
        fmt.Println(v23)
    case "hello":
        var v26 float32 = (1e10 - 3.14159)
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
