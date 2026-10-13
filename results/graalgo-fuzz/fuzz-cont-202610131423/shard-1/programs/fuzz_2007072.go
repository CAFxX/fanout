package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string; F1 int64; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }

func main() {
    var v1 float32 = (0.0 / 3.14159)
    var v2 bool = (0 != 9223372036854775807)
    v3 := make(map[int]int)
    var v4 int = 802
    var v5 uint = 2
    v3[9223372036854775807] = v4
    v6 := GenS1{F0: (!true), F1: 990, F2: 1}
    if (!false) {
        var v7 uint64 = 18446744073709551615
        _ = v7
    } else {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 float64 = (-2.25 + 1.5)
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := []int{fi8, fi8, -1}
                _ = len(v3)
                switch "" {
                case "a":
                    if ((!v2) || v2) {
                        var v12 uint64 = 4294967295
                        var v13 uint = (^(v5))
                        _ = v12
                        _ = v13
                    } else {
                        v14 := GenS1{F0: v2, F1: 857, F2: 0}
                        _ = v14
                    }
                case "":
                    v15 := GenS2{F0: ((("\n\t" + "foo bar") + "hello") + (("a" + "\n\t") + ("a" + "a"))), F1: 1, F2: (fi10 / ((32768) | 1))}
                    _ = v15
                case "hello":
                    var v16 int8 = -128
                    _ = v16
                }
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    }
    if (!(!("a" >= "hello"))) {
        v1 -= v1
        v17 := GenS3{F0: "hello"}
        fmt.Println(v1, v2, v4)
        _ = v17
    }
    v18 := GenS4{F0: v2}
    fmt.Println(v1, v2, v5)
    var v19 int32 = 1
    v20 := GenS2{F0: "foo bar", F1: -598, F2: v4}
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 int16 = 1
        var v23 int16 = 0
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v19, v5)
    var v24 int = -557
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
