package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8; F1 uint64; F2 int64 }
func genFn0(p0 int64, p1 int8) int8 {
    p0 += (p0 % (((p0 - (-(p0)))) | 1))
    if true {
        if false {
            var v5 int64 = 42
            fmt.Println(v5, p1)
            p1 = -76
            _ = v5
        } else {
            var v6 uint64 = 427
            v7 := GenS1{F0: 2147483647, F1: 255}
            if (!(-2147483648 < 42)) {
                v6 -= (v6 - 0)
            }
            _ = v6
            _ = v7
        }
        var v8 string = "foo bar"
        fmt.Println(v8, p1)
        _ = v8
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 int32 = 1
    v4 := make(map[uint32]bool)
    var v9 uint16 = 889
    v3 = 1
    var v10 int = 1
    fmt.Println(v10, v10)
    v11 := GenS2{F0: ("foo bar" + "a"), F1: "a"}
    v12 := GenS1{F0: 249, F1: (v9 >> ((65535) & 63))}
    if (true || false) {
        var v13 string = (((("a" + ("hello" + ("foo bar" + ("\n\t" + ("hello" + ("" + "a")))))) + "a") + "a") + "")
        _ = v13
    } else {
        fmt.Println(v3, v9, v10)
    }
    var v14 uint32 = 1
    var v15 float64 = 0.0
    switch "foo bar" {
    case "":
        v16 := GenS3{F0: (v10 >> ((738) & 63)), F1: -819}
        _ = v16
    case "a":
        var v17 int8 = -1
        _ = v17
    case "hello":
        v15 = float64(0.1)
    }
    if (0 < 127) {
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(v15, fi18, fi18)
            _ = fi18
        }
    }
    v14 -= 806
    if true {
        if false {
            fmt.Println(v15, v14)
            v19 := GenS4{F0: false}
            fmt.Println(v14, v9, v3)
            _ = v19
        }
    } else {
        v20 := []int{}
        _ = v20
    }
    v21 := GenS5{F0: 121, F1: 2, F2: 42}
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
