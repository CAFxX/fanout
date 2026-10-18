package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint16; F2 string }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int8; F1 uint32 }
func genFn0(p0 int64) int8 {
    v5 := GenS1{F0: 2}
    fmt.Println(p0)
    if (-169 != -1) {
        var v6 int32 = -980
        v7 := []int{}
        p0 = (p0 | p0)
        _ = v6
        _ = v7
    } else {
        p0 -= (p0 / (((p0 + -1)) | 1))
        v8 := GenS2{F0: "", F1: 331, F2: ("" + ("a" + "\n\t"))}
        _ = v8
    }
    _ = p0
    _ = v5
    return 108
}

func main() {
    var v1 uint16 = 32767
    v2 := []string{"hello", (("foo bar" + "hello") + "")}
    v3 := make(map[uint32]string)
    var v4 uint = 471
    var v9 int16 = 959
    fmt.Println(v1)
    v10 := GenS3{F0: 42, F1: 127}
    fmt.Println(v9, v9)
    var v11 int64 = 9223372036854775807
    if false {
        _ = len(v2)
    }
    var v12 bool = ("foo bar" < "foo bar")
    var v13 int32 = 2147483647
    if (v12 || (float32(-2.25) != 1e10)) {
        var v14 float32 = 3.14159
        _ = v14
    } else {
        v15 := GenS4{F0: -1, F1: 65535}
        _ = v15
    }
    fmt.Println(v1, v12)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
