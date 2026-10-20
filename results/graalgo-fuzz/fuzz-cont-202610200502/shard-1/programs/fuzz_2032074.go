package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint8; F1 string }
func genFn0(p0 bool, p1 int16, p2 int16) uint8 {
    fmt.Println(p2, p1, p0)
    if (!(255 < 1)) {
        var v7 string = "a"
        v7 = v7
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    v1 := []int{338, 2147483647}
    var v2 uint32 = 42
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    var v5 string = "foo bar"
    v6 := []string{(v5 + v5), (v5 + v5), v5, v5}
    var v8 float64 = (1.5 * 3.14159)
    v2 = v2
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := v3[(fi9 - (fi9 ^ fi9))]
        v2 = (v2 ^ (v2 * v2))
        for fi11 := 0; fi11 < 2; fi11++ {
            v10 = (v10 + v10)
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    v12 := GenS1{F0: 127, F1: 9223372036854775807}
    fmt.Println(v8)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 float32 = (0.1 / 3.14159)
        if false {
            var v15 uint32 = (v2 ^ 2147483648)
            fmt.Println(v8, v5)
            var v16 int64 = 42
            _ = v15
            _ = v16
        } else {
            if (0 != 32767) {
                delete(v4, "")
            } else {
                fmt.Println(v14, v5, v8)
                switch (((v5 + v5) + v5) + (v5 + "a")) {
                case "foo bar":
                    v17 := GenS2{F0: 2, F1: 996}
                    _ = v17
                case "":
                    fmt.Println(v5)
                case "hello":
                    v18 := GenS1{F0: 127, F1: 1}
                    _ = v18
                }
            }
            _ = len(v3)
            var v19 uint64 = 18446744073709551615
            _ = v19
        }
        _ = fi13
        _ = v14
    }
    v20 := GenS1{F0: 79, F1: -309}
    var v21 float32 = float32(v8)
    v22 := GenS3{F0: 21, F1: (v5 + (v5 + v5))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
