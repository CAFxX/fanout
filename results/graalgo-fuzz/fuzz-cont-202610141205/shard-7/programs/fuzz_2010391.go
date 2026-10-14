package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 bool = true
    v5 := GenS1{F0: 9223372036854775807}
    if false {
        v4 = v4
        for fi6 := 0; fi6 < 6; fi6++ {
            fmt.Println(v4, fi6)
            v7 := GenS2{F0: 27}
            _ = fi6
            _ = v7
        }
    }
    var v8 string = ((("foo bar" + "\n\t") + ("foo bar" + "a")) + "\n\t")
    v8 = v8
    var v9 int64 = -1
    if v4 {
        switch 127 {
        case 9223372036854775807:
            var v10 uint32 = 4294967295
            _ = v10
        case -2147483648:
            v11 := []int{-1, -672, 127, 579}
            _ = v11
        case 1:
            var v12 uint64 = 4294967295
            _ = v12
        default:
            v9 = v9
        }
        delete(v3, 538)
        v8 = v8
    }
    var v13 bool = ((528 != 0) && false)
    if false {
        var v14 float32 = 1.5
        for fi15 := 0; fi15 < 14; fi15++ {
            if (true && (-2.25 == (1.5 / v14))) {
                var v16 uint64 = 743
                _ = v16
            } else {
                var v17 uint8 = 0
                delete(v2, v8)
                var v18 int16 = int16(fi15)
                _ = v17
                _ = v18
            }
            var v19 float64 = 0.1
            var v20 uint = 18446744073709551615
            _ = fi15
            _ = v19
            _ = v20
        }
        _ = len(v3)
        _ = v14
    }
    if (!v13) {
        v21 := GenS3{F0: -482, F1: 635}
        _ = v21
    } else {
        v22 := GenS2{F0: -117}
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
