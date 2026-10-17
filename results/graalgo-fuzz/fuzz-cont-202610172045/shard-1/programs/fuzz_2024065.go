package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 string; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[int]int)
    var v2 string = "\n\t"
    var v3 uint32 = 426
    var v4 string = ("\n\t" + (v2 + ("a" + "foo bar")))
    fmt.Println(v3, v2, v3)
    _ = len(v1)
    if true {
        delete(v1, 32768)
        switch -9223372036854775808 {
        case 9223372036854775807:
            var v5 uint64 = 143
            _ = v5
        case 1:
            var v6 int8 = 0
            _ = v6
        }
    }
    if (118 == 42) {
        var v7 int32 = 42
        _ = v7
    }
    v4 = v4
    v8 := GenS1{F0: 0, F1: 544}
    var v9 int8 = 100
    delete(v1, 32767)
    if (true || (32768 >= -815)) {
        v1[4294967295] = 127
    }
    v1[32767] = -938
    var v10 uint = 845
    if (23 != v9) {
        fmt.Println(v9, v2)
    }
    if false {
        v11 := GenS2{F0: v4, F1: 54, F2: 145}
        v12 := GenS3{F0: -9223372036854775808, F1: -336}
        _ = v11
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
