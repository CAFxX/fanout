package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int64 }

func main() {
    var v1 string = (("\n\t" + (("" + ("" + "hello")) + "a")) + "hello")
    var v2 uint8 = 2
    var v3 uint = 18446744073709551615
    v3 -= (v3 & 0)
    v4 := []int{-853}
    var v5 string = (v1 + v1)
    if (1 > -1) {
        var v6 int64 = 399
        _ = v6
    }
    var v7 bool = (v3 <= 65535)
    var v8 int32 = 173
    var v9 uint64 = 2
    fmt.Println(v3)
    v10 := GenS1{F0: 0, F1: int64(v2)}
    var v11 uint16 = 2
    if (2 != uint8(v11)) {
        switch -501 {
        case -332:
            var v12 uint = (^(v3))
            _ = v12
        case 32767:
            var v13 int64 = -9223372036854775808
            _ = v13
        default:
            if (1 < 42) {
                for fi14 := 0; fi14 < 19; fi14++ {
                    v15 := GenS2{F0: ((v5 + v1) + "")}
                    _ = fi14
                    _ = v15
                }
                var v16 float32 = 1e10
                _ = v16
            }
        }
    }
    v5 = v5
    var v17 uint16 = (v11 & 460)
    v18 := GenS3{F0: ((v1 + v1) + "hello"), F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
