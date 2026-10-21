package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint64 }

func main() {
    var v1 string = ((((((("\n\t" + "foo bar") + "") + "hello") + "hello") + "") + "hello") + ("hello" + "a"))
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 65535
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 uint64 = 299
        fmt.Println(fi5, v1, v6)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: (^(v4)), F1: 127, F2: 4294967295}
    v8 := GenS1{F0: (^(v4)), F1: 100, F2: 2}
    if (581 == 0) {
        fmt.Println(v4, v1)
    }
    v1 = "hello"
    fmt.Println(v1)
    v9 := GenS2{F0: 339}
    switch v1 {
    case "\n\t":
        switch v1 {
        case "\n\t":
            var v10 uint64 = 18446744073709551615
            _ = v10
        default:
            var v11 uint32 = v4
            _ = v11
        }
    case "a":
        fmt.Println(v1, v4)
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS3{F0: 983, F1: 829}
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
