package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
type GenS2 struct { F0 string; F1 int64; F2 string }
type GenS3 struct { F0 int64; F1 uint32 }

func main() {
    var v1 string = "foo bar"
    var v2 int16 = 1
    v3 := make(map[int]int)
    v4 := []string{v1, ("\n\t" + v1), v1, v1}
    v5 := []string{v1, v1, (v1 + "hello"), (("hello" + v1) + "hello")}
    var v6 uint64 = 1
    v7 := GenS1{F0: int(v6), F1: 127, F2: 42}
    v1 = v1
    v6 = (v6 & v6)
    v8 := GenS1{F0: 32767, F1: 128, F2: 2}
    _ = len(v3)
    fmt.Println(v1)
    var v9 int8 = -127
    v10 := GenS2{F0: v1, F1: -9223372036854775808, F2: (("a" + "hello") + (("hello" + v1) + "foo bar"))}
    var v11 int = 127
    v9 = (v9 % (((^(v9))) | 1))
    var v12 int64 = 0
    switch (((("hello" + v1) + v1) + (v1 + (v1 + v1))) + (v1 + "")) {
    case "hello":
        v4 = append(v4, v1)
    }
    if false {
        var v13 float32 = 1e10
        if (false && true) {
            var v14 uint16 = 123
            v15 := GenS3{F0: v12, F1: 42}
            _ = v14
            _ = v15
        } else {
            v16 := GenS1{F0: (v11 | 0), F1: 9, F2: 54}
            var v17 uint8 = 128
            _ = v16
            _ = v17
        }
        if true {
            if (v1 >= "hello") {
                delete(v3, (-(v11)))
            }
        }
        _ = v13
    }
    fmt.Println(v6, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
