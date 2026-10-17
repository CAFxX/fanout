package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 18446744073709551615
    v3 := []uint32{2147483648, 425, 65535, 695}
    for fi4 := 0; fi4 < 2; fi4++ {
        for fi5 := 0; fi5 < 4; fi5++ {
            v2 = (^(v2))
            fmt.Println(fi4)
            _ = fi5
        }
        _ = fi4
    }
    v6 := GenS1{F0: (42 >= -345), F1: 1}
    fmt.Println(v2)
    switch "\n\t" {
    case "\n\t":
        v7 := GenS2{F0: (v2 << (((v2 + v2)) & 63))}
        _ = v7
    case "hello":
        v8 := GenS3{F0: -9223372036854775808, F1: 2147483647}
        _ = v8
    case "a":
        v9 := GenS2{F0: 2}
        _ = v9
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v2)
    case "":
        if (719 > 561) {
            v10 := GenS1{F0: true, F1: -32768}
            for fi11 := 0; fi11 < 16; fi11++ {
                switch (fi11 ^ fi11) {
                case -9223372036854775808:
                    switch ("foo bar" + "\n\t") {
                    case "\n\t":
                        fmt.Println(v2)
                    case "":
                        v2 -= v2
                    }
                }
                _ = fi11
            }
            _ = v10
        }
    }
    var v12 int32 = 42
    fmt.Println(v12, v2)
    fmt.Println(v2, v12)
    var v13 int8 = -1
    v12 -= (v12 - v12)
    var v14 string = ""
    fmt.Println(v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v6
}
