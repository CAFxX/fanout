package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    v3 := []uint32{65535, 65535, 2147483647, 2147483648}
    var v4 int = -9223372036854775808
    v5 := make(map[string]bool)
    fmt.Println(v4, v4)
    v2 = (^(v2))
    switch ("hello" + "\n\t") {
    case "\n\t":
        v4 = 4294967295
    case "":
        v6 := GenS1{F0: 32767, F1: true, F2: 4294967295}
        _ = v6
    default:
        var v7 uint32 = 2147483647
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: 1, F1: true, F2: 4294967295}
        var v10 string = ("" + "\n\t")
        v3 = append(v3, 65535)
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = "foo bar"
    v2 = v2
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(v2)
        var v13 int = (fi12 | fi12)
        _ = fi12
        _ = v13
    }
    fmt.Println(v11, v4)
    var v14 string = v11
    var v15 string = v11
    v14 = v14
    v11 = (((((v11 + v11) + ((v11 + "foo bar") + v11)) + "") + v15) + ((v11 + ("\n\t" + (v14 + ("a" + v11)))) + "hello"))
    fmt.Println(v15)
    v15 = v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
