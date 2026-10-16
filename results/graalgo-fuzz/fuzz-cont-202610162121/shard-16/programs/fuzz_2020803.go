package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint) uint32 {
    var v6 uint32 = 886
    v7 := []int{32767, 0, -2147483648}
    _ = len(v7)
    _ = p0
    _ = v6
    _ = v7
    return (^(v6))
}
func genFn1() int32 {
    var v8 uint32 = 268
    v9 := GenS1{F0: 42, F1: 128}
    var v10 int8 = 0
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (42 == 42)
    var v3 float64 = 0.1
    var v4 float64 = (v3 - ((v3 * float64(3.14159)) * ((0.0 / 3.14159) - v3)))
    var v5 int64 = -1
    var v11 uint8 = 103
    fmt.Println(v4, v2, v4)
    v12 := GenS1{F0: -1, F1: 2147483647}
    fmt.Println(v4, v5)
    v13 := GenS2{F0: (4294967295 > 42)}
    v14 := []int{4294967295, 32768, 9223372036854775807}
    var v15 int64 = (v5 + (v5 % ((v5) | 1)))
    fmt.Println(v4, v3)
    if (1 >= (v11 | 0)) {
        switch ("a" + "foo bar") {
        case "":
            _ = len(v14)
        }
        switch (("foo bar" + "") + ("" + "foo bar")) {
        case "a":
            v16 := GenS3{F0: 42}
            _ = v16
        }
        v3 = (v3 - v4)
    } else {
        v4 += v3
        var v17 uint64 = 659
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
