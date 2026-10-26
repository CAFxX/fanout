package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int64; F2 int32 }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]int)
    var v3 string = "foo bar"
    v1 -= v1
    fmt.Println(v3)
    var v4 uint32 = 834
    if (!true) {
        for fi5 := 0; fi5 < 8; fi5++ {
            fmt.Println(v1, v1, fi5)
            fmt.Println(v3, v1)
            v6 := []int{662}
            _ = fi5
            _ = v6
        }
        if (!(3.14159 <= (1.5 * 0.0))) {
            v4 -= (v4 / ((v4) | 1))
        } else {
            v7 := GenS1{F0: "a"}
            var v8 string = v3
            v9 := GenS1{F0: ((v3 + v8) + (v8 + v3))}
            _ = v7
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 string = v3
            switch fi10 {
            case -1:
                v1 = v1
            case -2147483648:
                v11 = v3
            case -9223372036854775808:
                if false {
                    fmt.Println(v1, fi10, v1)
                }
            }
            _ = fi10
            _ = v11
        }
    }
    v12 := v2[v4]
    if (false && true) {
        var v13 int8 = (^(v1))
        _ = v13
    }
    var v14 int64 = 42
    v15 := GenS2{F0: 1, F1: 42, F2: int32(v14)}
    delete(v2, 0)
    var v16 int = (v12 % ((v12) | 1))
    var v17 uint = 2
    var v18 uint = (^(v17))
    var v19 uint = v18
    v20 := []int{(^(v12)), (v12 / ((884) | 1)), (v12 ^ 302)}
    var v21 uint8 = 7
    v22 := GenS3{F0: (v4 | v4), F1: (-(v14)), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
