package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 int64 }
func genFn0(p0 int64, p1 string, p2 int8) int64 {
    v7 := GenS1{F0: 657}
    v8 := GenS1{F0: 518}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 9223372036854775807
}
func genFn1(p0 int8, p1 int8, p2 string) bool {
    var v9 float64 = float64((((3.14159 - 0.0) / (-2.25 / 3.14159)) + 0.0))
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (!false)
}

func main() {
    var v1 string = ("\n\t" + "")
    var v2 uint = 1
    v3 := make(map[string]string)
    var v4 uint8 = 6
    v5 := []uint32{}
    var v6 float32 = 1e10
    v11 := GenS2{F0: 0, F1: ("foo bar" + v1), F2: "hello"}
    fmt.Println(v6)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v6, v4)
    if (100 == -1) {
        v1 = v1
        var v12 int = 127
        _ = v12
    }
    v13 := GenS3{F0: 1}
    v14 := GenS1{F0: (^(v2))}
    if (false || false) {
        fmt.Println(v6)
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
