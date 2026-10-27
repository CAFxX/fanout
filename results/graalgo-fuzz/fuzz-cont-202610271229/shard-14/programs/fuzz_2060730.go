package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 245
    v3 := make(map[string]string)
    var v4 bool = false
    var v5 uint64 = 4294967295
    v2 -= uint8(v5)
    v4 = (!(!true))
    var v6 bool = true
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 879}
        var v9 int64 = -9223372036854775808
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int = fi10
            switch (fi10 ^ 4294967295) {
            case 710:
                var v13 uint64 = v5
                _ = v13
            default:
                var v14 float64 = 1e10
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        var v15 int16 = 1
        var v16 string = "foo bar"
        _ = fi10
        _ = v15
        _ = v16
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v6, v4)
    case "hello":
        var v17 float32 = ((0.1 * 0.1) - 3.14159)
        _ = v17
    case "\n\t":
        fmt.Println(v2)
    default:
        switch "hello" {
        case "hello":
            var v18 uint8 = v2
            _ = v18
        case "foo bar":
            v5 -= (^(v5))
        }
    }
    if true {
        for fi19 := 0; fi19 < 13; fi19++ {
            for fi20 := 0; fi20 < 11; fi20++ {
                v2 -= 126
                v2 += v2
                _ = fi20
            }
            _ = fi19
        }
        v21 := GenS1{F0: 1}
        _ = v21
    } else {
        var v22 int8 = 39
        _ = v22
    }
    switch 746 {
    case 42:
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := []int{}
            _ = fi23
            _ = v24
        }
    case -2147483648:
        v6 = true
    case -1:
        v3["a"] = "foo bar"
    default:
        switch 937 {
        case 903:
            v25 := []int{127}
            _ = v25
        default:
            switch "hello" {
            case "\n\t":
                v26 := GenS1{F0: 255}
                _ = v26
            case "foo bar":
                v27 := []int{}
                _ = v27
            }
        }
    }
    v28 := []int{-9223372036854775808, 127, -694, 128}
    delete(v1, 36)
    var v29 float32 = float32((1e10 - 0.0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
