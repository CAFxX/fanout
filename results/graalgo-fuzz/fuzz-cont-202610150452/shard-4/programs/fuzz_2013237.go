package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int16 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int64; F2 int32 }
func genFn0(p0 int32) int32 {
    v4 := GenS1{F0: ((("hello" + "") + ("a" + "hello")) + "hello"), F1: -489, F2: -302}
    v5 := GenS2{F0: "", F1: 125, F2: 196}
    _ = p0
    _ = v4
    _ = v5
    return (-(p0))
}

func main() {
    v1 := []int{}
    var v2 int32 = -2147483648
    var v3 uint = 65535
    v3 += v3
    v2 = genFn0(v2)
    v3 -= 269
    var v6 bool = (((!true) || (18446744073709551615 >= 380)) || true)
    var v7 uint8 = 98
    var v8 string = "foo bar"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 := GenS3{F0: 1, F1: -799, F2: (v2 >> (((v3 / (((v3 + v3)) | 1))) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
