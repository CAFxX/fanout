package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]bool)
    var v3 int = 128
    var v4 int8 = 0
    for fi5 := 0; fi5 < 16; fi5++ {
        v4 += v4
        _ = fi5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            switch "a" {
            case "":
                var v8 int8 = (v4 + (v4 + v4))
                _ = v8
            }
            _ = fi7
        }
        fmt.Println(v3, v4)
        _ = fi6
    }
    v4 += (v4 * (-(v4)))
    var v9 uint = 255
    switch (v3 % ((v3) | 1)) {
    case 9223372036854775807:
        v2[v3] = (!(v1 >= (v1 - 9223372036854775807)))
    case 885:
        v10 := []int{(v3 % ((v3) | 1)), (v3 / (((v3 | v3)) | 1)), -519}
        _ = v10
    case 4294967295:
        for fi11 := 0; fi11 < 5; fi11++ {
            switch "a" {
            case "hello":
                var v12 uint64 = 2
                _ = v12
            case "foo bar":
                v1 = 0
            case "":
                v13 := GenS1{F0: 78, F1: (fi11 + 32767), F2: v9}
                _ = v13
            default:
                delete(v2, (^(v3)))
            }
            _ = fi11
        }
    }
    var v14 bool = (760 != 0)
    var v15 uint = v9
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
