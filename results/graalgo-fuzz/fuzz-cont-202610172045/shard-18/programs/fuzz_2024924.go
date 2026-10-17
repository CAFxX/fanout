package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int }
type GenS2 struct { F0 uint; F1 int8; F2 int32 }

func main() {
    v1 := []string{"a", "foo bar", ("\n\t" + "foo bar")}
    var v2 int = 9223372036854775807
    v3 := make(map[uint32]int)
    var v4 uint = 18446744073709551615
    var v5 int64 = -1
    fmt.Println(v5, v4)
    if false {
        v6 := GenS1{F0: ((4294967295 < (v4 >> ((538) & 63))) || true), F1: false, F2: v2}
        if true {
            fmt.Println(v5, v4, v4)
            delete(v3, 2147483648)
            fmt.Println(v2)
        } else {
            var v7 uint16 = 1
            _ = v7
        }
        var v8 float32 = (1.5 - 1.5)
        _ = v6
        _ = v8
    }
    var v9 uint = v4
    v5 = v5
    var v10 uint16 = 0
    var v11 int32 = -108
    switch ("foo bar" + "hello") {
    case "":
        v10 -= (^(v10))
    }
    v10 = (v10 % (((v10 >> ((v4) & 63))) | 1))
    v12 := GenS2{F0: v9, F1: -13, F2: (v11 % ((int32(v10)) | 1))}
    var v13 uint64 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
