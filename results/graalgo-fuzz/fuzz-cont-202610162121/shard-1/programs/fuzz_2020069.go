package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
type GenS4 struct { F0 uint16; F1 bool }
func genFn0() int16 {
    var v4 uint64 = 457
    fmt.Println(v4)
    var v5 uint32 = 0
    for fi6 := 0; fi6 < 20; fi6++ {
        v4 += 0
        _ = fi6
    }
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := []int{42}
    var v2 uint = 255
    v3 := make(map[uint32]bool)
    v1 = append(v1, 32767)
    var v7 uint64 = 4294967295
    fmt.Println(v2)
    var v8 float64 = 1e10
    if (-2147483648 == -2147483648) {
        v9 := GenS1{F0: (v7 & 1), F1: 255}
        v8 -= (v8 + 0.1)
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 uint = 18446744073709551615
            v7 += (^(v7))
            v2 -= (v11 | v11)
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        var v12 uint = v2
        _ = v12
    }
    _ = len(v1)
    if (true || ((!false) && (genFn0() == -32768))) {
        delete(v3, 65535)
        v13 := GenS1{F0: v7, F1: 216}
        _ = v13
    }
    v14 := GenS2{F0: (v7 | (v7 * uint64(v2)))}
    var v15 uint8 = 127
    v7 = 235
    switch "hello" {
    case "foo bar":
        v16 := GenS3{F0: (v15 - v15), F1: 1, F2: -22}
        _ = v16
    }
    switch ("a" + ("\n\t" + "")) {
    case "\n\t":
        var v17 string = ("\n\t" + "foo bar")
        _ = v17
    case "":
        v2 = 2
    default:
        v2 = (^(v2))
    }
    var v18 int32 = 2147483647
    v19 := GenS4{F0: 920, F1: ((-2.25 * 3.14159) >= (1e10 - float64((3.14159 * float32(((v8 / v8) * v8))))))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
