package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }

func main() {
    var v1 uint32 = 65535
    var v2 int8 = -1
    v3 := make(map[uint32]string)
    v4 := []string{}
    switch (("\n\t" + ((("hello" + "\n\t") + "a") + "foo bar")) + "hello") {
    case "hello":
        v1 = (v1 % ((v1) | 1))
    case "":
        switch 0 {
        case -683:
            v2 += 127
        case 1:
            v1 += v1
        case 42:
            if (3 != 42) {
                for fi5 := 0; fi5 < 12; fi5++ {
                    var v6 uint8 = 231
                    var v7 int64 = 1
                    _ = fi5
                    _ = v6
                    _ = v7
                }
                v8 := GenS1{F0: -9223372036854775808, F1: 42}
                _ = v8
            }
        }
    default:
        v1 = v1
    }
    var v9 uint16 = 65535
    v10 := GenS1{F0: 1, F1: 0}
    var v11 float64 = 1.5
    v12 := GenS1{F0: 316, F1: 0}
    if false {
        var v13 int32 = 0
        _ = v13
    }
    fmt.Println(v11, v2, v9)
    if true {
        fmt.Println(v1, v2, v2)
    } else {
        v1 = v1
    }
    v11 = 0.1
    var v14 int64 = 358
    var v15 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
