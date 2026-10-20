package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint; F2 uint8 }

func main() {
    var v1 uint32 = 2147483648
    var v2 string = "foo bar"
    var v3 string = "foo bar"
    for fi4 := 0; fi4 < 17; fi4++ {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 int8 = -1
            switch 0 {
            case 128:
                v6 += v6
            case 1:
                var v7 int16 = -450
                _ = v7
            default:
                v2 = "a"
                var v8 int = (fi4 & fi5)
                fmt.Println(v8, v2)
                _ = v8
            }
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v9 := []int{-2147483648, 9223372036854775807}
    if (!true) {
        fmt.Println(v3, v1)
        v10 := GenS1{F0: -9223372036854775808, F1: 22}
        v11 := GenS2{F0: 42, F1: 359}
        _ = v10
        _ = v11
    } else {
        var v12 int = 128
        if false {
            if true {
                v13 := GenS1{F0: -1, F1: -76}
                _ = v13
            }
            var v14 string = (v3 + "hello")
            v12 = (-(v12))
            _ = v14
        } else {
            var v15 int = v12
            _ = v15
        }
        fmt.Println(v3)
        _ = v12
    }
    v16 := GenS3{F0: 42, F1: 42, F2: 129}
    fmt.Println(v2, v1)
    var v17 bool = false
    v17 = v17
    switch v2 {
    case "a":
        fmt.Println(v17, v2, v3)
    default:
        _ = len(v9)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
