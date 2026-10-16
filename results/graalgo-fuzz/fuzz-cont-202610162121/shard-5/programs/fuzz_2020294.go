package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int, p1 uint8) int8 {
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := []int{32767}
        _ = fi4
        _ = v5
    }
    if true {
        p1 = (p1 << ((65535) & 63))
        v6 := GenS1{F0: -801, F1: 0}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 2; fi7++ {
            for fi8 := 0; fi8 < 3; fi8++ {
                p0 -= 4294967295
                var v9 float64 = (0.1 * (1.5 * 0.0))
                v10 := []int{(fi8 % ((32768) | 1))}
                v11 := GenS2{F0: 56}
                switch ("a" + "a") {
                case "hello":
                    fmt.Println(p1, fi8)
                case "":
                    v12 := GenS1{F0: 0, F1: -128}
                    _ = v12
                case "foo bar":
                    _ = len(v10)
                }
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            var v13 int64 = 47
            _ = fi7
            _ = v13
        }
        fmt.Println(p0, p1)
    }
    v14 := []int{p0, 42, (^(p0))}
    v15 := GenS2{F0: 127}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 int16 = 42
    var v2 int = 32767
    v3 := make(map[string]string)
    if true {
        v3["a"] = ("" + "\n\t")
        fmt.Println(v1, v2)
    }
    if false {
        if true {
            fmt.Println(v2)
            switch "" {
            case "hello":
                v1 = 0
            case "foo bar":
                var v16 float32 = -2.25
                _ = v16
            default:
                fmt.Println(v1)
            }
            v1 = (v1 ^ (v1 + (v1 / (((v1 - -32768)) | 1))))
        }
        var v17 uint32 = 42
        _ = v17
    }
    var v18 int16 = (v1 & (v1 | v1))
    v19 := GenS2{F0: -76}
    var v20 uint16 = 255
    _ = len(v3)
    switch v2 {
    case -1:
        fmt.Println(v2, v2, v20)
    default:
        v20 = v20
    }
    var v21 uint64 = 281
    v22 := GenS2{F0: 1}
    v18 -= 1000
    var v23 int64 = -1
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
