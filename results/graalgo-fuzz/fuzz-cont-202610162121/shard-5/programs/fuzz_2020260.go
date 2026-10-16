package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]string)
    var v3 int32 = 42
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 float32 = -2.25
        _ = fi4
        _ = v5
    }
    fmt.Println(v3, v1)
    fmt.Println(v3)
    v1 = (v1 >> ((42) & 63))
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{}
        if (true || (false || (!((-2147483648 <= v3) || (42 > 4294967295))))) {
            var v8 uint = 18446744073709551615
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    var v9 int64 = v1
    var v10 int64 = (^(v9))
    v10 -= v9
    fmt.Println(v1, v3)
    var v11 bool = false
    v2[0] = (("\n\t" + "\n\t") + "")
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS1{F0: 372}
        v14 := []int{-1}
        _ = fi12
        _ = v13
        _ = v14
    }
    switch "hello" {
    case "foo bar":
        v10 -= 0
    case "a":
        fmt.Println(v10, v3)
    case "":
        if (v11 && (4294967295 < 2)) {
            v15 := GenS2{F0: 42, F1: 73}
            _ = v15
        } else {
            v1 -= (v1 % (((v10 ^ (v10 % ((v1) | 1)))) | 1))
        }
    }
    v16 := GenS3{F0: 0, F1: 1}
    fmt.Println(v3)
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS1{F0: 2}
        _ = fi17
        _ = v18
    }
    fmt.Println(v10, v9)
    v19 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
