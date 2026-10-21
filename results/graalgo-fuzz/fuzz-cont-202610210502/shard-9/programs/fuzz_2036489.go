package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    var v1 bool = (1 != 42)
    var v2 int = 1
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    v5 := make(map[uint32]string)
    var v6 int32 = 2147483647
    var v7 float32 = (3.14159 / (1e10 + (1.5 * 3.14159)))
    v1 = (!(!false))
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            delete(v5, 695)
            _ = fi9
        }
        var v10 float32 = v7
        switch (-(fi8)) {
        case 27:
            v6 = (v6 & int32(v2))
        }
        _ = fi8
        _ = v10
    }
    if false {
        v11 := GenS1{F0: "foo bar", F1: (v6 >> ((1) & 63)), F2: 939}
        _ = v11
    }
    var v12 int8 = -1
    var v13 int = v2
    switch ("foo bar" + (("a" + "foo bar") + (("foo bar" + "\n\t") + "a"))) {
    case "hello":
        var v14 int = v13
        _ = v14
    case "foo bar":
        v15 := []int{(v13 + (-(v13))), 2147483647, v13, v2}
        _ = v15
    case "\n\t":
        v16 := GenS2{F0: 42, F1: (v1 || false)}
        _ = v16
    default:
        delete(v4, 9223372036854775807)
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 uint64 = 5
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := GenS3{F0: (v6 >> ((2) & 63)), F1: 1}
        fmt.Println(v2)
        var v21 uint32 = 65535
        _ = fi19
        _ = v20
        _ = v21
    }
    v6 = 0
    v7 += v7
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
