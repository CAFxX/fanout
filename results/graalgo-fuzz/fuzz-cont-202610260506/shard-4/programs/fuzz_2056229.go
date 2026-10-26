package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint64 }
type GenS2 struct { F0 int8; F1 string; F2 uint16 }
type GenS3 struct { F0 uint16; F1 string; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int }

func main() {
    var v1 bool = (1 <= 2)
    v2 := []uint32{2147483648, 2147483648, 4294967295, 2}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 uint16 = 0
    v6 := make(map[uint32]int)
    fmt.Println(v1, v5)
    var v7 uint32 = 544
    fmt.Println(v7, v1, v5)
    v1 = (!true)
    v8 := GenS1{F0: 724, F1: ("a" + (("foo bar" + "hello") + "a")), F2: uint64(v5)}
    switch 32768 {
    case 831:
        var v9 uint32 = v7
        _ = v9
    case -9223372036854775808:
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS1{F0: (v7 >> ((2) & 63)), F1: "", F2: 628}
            _ = fi10
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS2{F0: 1, F1: "foo bar", F2: 0}
            v14 := GenS3{F0: 462, F1: ("foo bar" + "foo bar"), F2: fi12}
            var v15 bool = v1
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    }
    fmt.Println(v7)
    var v16 uint32 = v7
    fmt.Println(v7, v7)
    if v1 {
        fmt.Println(v1, v16)
        if v1 {
            v7 += 432
            v1 = false
            if false {
                v17 := GenS4{F0: v16}
                _ = v17
            }
        }
    } else {
        v1 = v1
        v18 := GenS5{F0: 42}
        var v19 float32 = (0.1 * -2.25)
        _ = v18
        _ = v19
    }
    var v20 int8 = 0
    var v21 uint64 = 700
    fmt.Println(v7, v21, v5)
    var v22 int = 0
    if (!v1) {
        v23 := v6[v7]
        v22 -= -2147483648
        v16 += (v7 * v7)
        _ = v23
    }
    fmt.Println(v16)
    fmt.Println(v7, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
