package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int8 }
func genFn0(p0 uint64, p1 uint64) int8 {
    v5 := GenS1{F0: "foo bar"}
    v6 := []int{744, 99}
    var v7 string = "hello"
    var v8 string = v7
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int64 = 0
    v2 := []int{}
    v3 := []string{"\n\t", "", "a", "hello"}
    v4 := make(map[uint32]string)
    var v9 uint = 255
    var v10 uint16 = 65535
    fmt.Println(v10, v1, v10)
    v1 -= (v1 * v1)
    v11 := GenS1{F0: "foo bar"}
    if ((v9 | 0) > 909) {
        v10 = 337
        var v12 int64 = v1
        _ = v12
    } else {
        v10 = uint16(v9)
        var v13 uint64 = 517
        _ = v13
    }
    fmt.Println(v1)
    v14 := v4[514]
    var v15 uint32 = 0
    v16 := GenS2{F0: v1, F1: -128}
    if (!true) {
        v17 := GenS2{F0: (^(v1)), F1: 36}
        _ = v17
    }
    fmt.Println(v10)
    var v18 int16 = 0
    v14 = ("a" + (((v14 + v14) + (v14 + v14)) + (v14 + (v14 + "a"))))
    v19 := GenS2{F0: (^(v1)), F1: genFn0(0, 18446744073709551615)}
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
