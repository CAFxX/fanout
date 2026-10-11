package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint16 }

func main() {
    var v1 int16 = -32768
    var v2 string = "hello"
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    switch 2147483647 {
    case 42:
        v2 = v2
    case -1:
        v6 := GenS1{F0: "foo bar", F1: (false || (v1 != (v1 | v1)))}
        _ = v6
    case 32767:
        switch v2 {
        case "hello":
            var v7 float64 = 1e10
            _ = v7
        }
    default:
        v8 := []int{42, 189}
        _ = v8
    }
    var v9 int = -570
    v9 = v9
    v10 := GenS1{F0: v2, F1: (!(v9 <= (v9 ^ (v9 * v9))))}
    var v11 uint64 = 42
    v12 := []int{(v9 | v9)}
    fmt.Println(v1, v11, v9)
    v13 := GenS2{F0: 0, F1: 1, F2: "\n\t"}
    v14 := GenS3{F0: (v11 % ((18446744073709551615) | 1))}
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS3{F0: v11}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v18 := GenS1{F0: ("a" + "\n\t"), F1: (!(v2 <= v2))}
        if true {
            v11 = (^(v11))
            v2 = ""
        }
        _ = fi17
        _ = v18
    }
    v19 := GenS4{F0: 127, F1: 847}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
