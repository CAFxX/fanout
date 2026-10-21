package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint8; F1 bool }
func genFn0(p0 int16, p1 int8) int8 {
    fmt.Println(p1, p1)
    v7 := GenS1{F0: 931, F1: -1}
    v8 := GenS2{F0: true, F1: 48}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 uint16, p1 uint8) uint16 {
    v9 := GenS3{F0: 127, F1: true}
    var v10 int32 = 2147483647
    fmt.Println(v10, p1)
    v11 := []int{674, 9223372036854775807, -1, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 729
}

func main() {
    v1 := []int{486}
    var v2 float32 = 1e10
    var v3 uint32 = 2147483647
    v4 := make(map[int]bool)
    var v5 uint16 = 444
    var v6 int64 = -918
    fmt.Println(v5, v5)
    _ = len(v1)
    v4[32768] = true
    v6 += 9223372036854775807
    v3 = uint32(v5)
    v2 = v2
    for fi12 := 0; fi12 < 2; fi12++ {
        _ = len(v4)
        if true {
            var v13 float64 = 0.1
            v14 := GenS3{F0: 127, F1: true}
            switch ((("foo bar" + "foo bar") + "\n\t") + "\n\t") {
            case "\n\t":
                var v15 int16 = 1000
                _ = v15
            }
            _ = v13
            _ = v14
        }
        v4[(-(fi12))] = true
        _ = fi12
    }
    var v16 uint64 = 2
    fmt.Println(v16, v5)
    var v17 uint32 = 481
    delete(v4, 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
