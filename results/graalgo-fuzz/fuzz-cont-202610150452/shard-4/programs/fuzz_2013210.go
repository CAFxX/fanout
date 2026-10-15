package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{2147483648, 2, 2147483648}
    var v3 string = ("hello" + "foo bar")
    v4 := make(map[uint32]int)
    v5 := []string{((v3 + "hello") + ("" + v3))}
    var v6 string = ("a" + "hello")
    fmt.Println(v3)
    v6 = "foo bar"
    if false {
        v6 = "a"
        v6 = v6
    } else {
        var v7 int16 = 1000
        if true {
            if (!true) {
                v5 = append(v5, "foo bar")
            } else {
                v8 := GenS1{F0: 541, F1: 1, F2: 2}
                _ = v8
            }
            v3 = v3
            v9 := GenS1{F0: 4294967295, F1: 0, F2: 18446744073709551615}
            _ = v9
        } else {
            v10 := GenS1{F0: 2147483647, F1: 1, F2: 42}
            _ = v10
        }
        _ = v7
    }
    v4[2] = 128
    v11 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 718}
    v4[475] = 128
    v6 = v6
    var v12 int32 = 2147483647
    var v13 uint32 = 712
    if (!((!(1000 >= -13)) || (!true))) {
        fmt.Println(v12)
        v3 = v3
    }
    if false {
        v14 := v4[(v13 ^ v13)]
        v15 := GenS2{F0: (v13 % ((v13) | 1))}
        _ = v14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
