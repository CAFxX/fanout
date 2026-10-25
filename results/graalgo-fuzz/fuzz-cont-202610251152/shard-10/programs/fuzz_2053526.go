package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint64, p1 uint64) int8 {
    v4 := GenS1{F0: false, F1: 1, F2: 255}
    v5 := GenS1{F0: (1 == -316), F1: 32768, F2: 127}
    v6 := GenS1{F0: (true || (!(!(true && true)))), F1: 9223372036854775807, F2: 2}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -1
    v3 := []int{-573}
    v3 = append(v3, 128)
    delete(v1, "hello")
    for fi7 := 0; fi7 < 19; fi7++ {
        v3 = append(v3, (-(fi7)))
        v2 = (v2 * (v2 + (v2 * v2)))
        _ = fi7
    }
    var v8 uint32 = 821
    delete(v1, "\n\t")
    for fi9 := 0; fi9 < 10; fi9++ {
        switch ("a" + "\n\t") {
        case "hello":
            _ = len(v3)
        case "a":
            v8 -= (^(v8))
        case "\n\t":
            fmt.Println(v8, v2, v2)
        default:
            switch "hello" {
            case "hello":
                v2 += v2
            case "a":
                var v10 uint16 = 42
                _ = v10
            }
        }
        if len(v3) > 0 {
            v3[0] = fi9
        }
        _ = len(v3)
        _ = fi9
    }
    fmt.Println(v2, v2)
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            fmt.Println(fi11, v8, v8)
            _ = fi12
        }
        if len(v3) > 0 {
            v3[0] = fi11
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            if len(v3) > 0 {
                v3[0] = fi11
            }
            _ = fi13
        }
        _ = fi11
    }
    v14 := GenS2{F0: 270, F1: -437}
    switch ("a" + "foo bar") {
    case "":
        fmt.Println(v8)
    case "hello":
        v8 += 657
    default:
        var v15 uint32 = v8
        _ = v15
    }
    v16 := GenS3{F0: 2147483647}
    if false {
        fmt.Println(v2)
        v3 = append(v3, -680)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
