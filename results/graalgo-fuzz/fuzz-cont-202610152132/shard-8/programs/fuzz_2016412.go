package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint32; F1 int32; F2 int }
func genFn0(p0 int64) uint8 {
    p0 = 0
    v5 := []int{-934, -9223372036854775808}
    _ = p0
    _ = v5
    return 154
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 1
    v3 := make(map[int]int)
    var v4 uint32 = 4294967295
    v1[("hello" + "hello")] = "a"
    var v6 uint8 = 235
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint16 = 695
        fmt.Println(v8)
        delete(v1, "a")
        _ = fi7
        _ = v8
    }
    if ("a" > "a") {
        switch "hello" {
        case "":
            v4 -= uint32(v2)
        case "\n\t":
            v9 := []int{}
            _ = v9
        case "a":
            for fi10 := 0; fi10 < 2; fi10++ {
                fmt.Println(v6, v6)
                var v11 int = fi10
                _ = fi10
                _ = v11
            }
        }
        v12 := GenS1{F0: (false && true), F1: 32767}
        _ = v12
    }
    v13 := GenS1{F0: true, F1: 1000}
    v14 := []int{127}
    var v15 uint64 = 42
    if false {
        fmt.Println(v15)
        switch "\n\t" {
        case "foo bar":
            var v16 int = 128
            _ = v16
        case "hello":
            if false {
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 int64 = 195
                    var v19 float32 = 1e10
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
        case "\n\t":
            v14 = append(v14, -2147483648)
        default:
            v20 := GenS2{F0: v4, F1: -494, F2: 0}
            _ = v20
        }
    }
    switch "\n\t" {
    case "":
        var v21 int = -9223372036854775808
        _ = v21
    case "a":
        v15 -= (^(v15))
    }
    switch 2147483647 {
    case 0:
        _ = len(v14)
    }
    if (725 < 9223372036854775807) {
        _ = len(v1)
    }
    v2 = uint(v4)
    if (false && true) {
        for fi22 := 0; fi22 < 5; fi22++ {
            v2 = (v2 - v2)
            v6 = (v6 % (((v6 - v6)) | 1))
            _ = fi22
        }
        v23 := GenS2{F0: v4, F1: 1, F2: 2147483647}
        if true {
            v24 := GenS2{F0: v4, F1: 566, F2: 32768}
            fmt.Println(v2, v4, v4)
            _ = v24
        } else {
            var v25 float64 = 3.14159
            _ = v25
        }
        _ = v23
    }
    fmt.Println(v15, v6, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
