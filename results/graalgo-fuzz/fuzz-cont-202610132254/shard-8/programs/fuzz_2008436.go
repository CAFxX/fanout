package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint }
type GenS2 struct { F0 uint8; F1 bool; F2 int16 }
func genFn0(p0 bool, p1 bool) bool {
    v6 := GenS1{F0: 32768, F1: "\n\t", F2: 65535}
    p0 = (166 > 100)
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 string = ""
    v4 := make(map[int]uint32)
    var v5 float64 = 0.0
    if true {
        var v7 uint = 862
        _ = v7
    }
    v8 := []int{-524}
    switch v3 {
    case "hello":
        if (true || false) {
            var v9 int64 = -1
            _ = v9
        }
    case "\n\t":
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 int16 = -859
            if false {
                v12 := GenS2{F0: 47, F1: (false || (true || true)), F2: 42}
                _ = v12
            } else {
                if (45 > 127) {
                    if ((((((v3 + (((("hello" + ("" + (("a" + v3) + ("a" + v3)))) + ("\n\t" + v3)) + ("a" + v3)) + v3)) + v3) + "hello") + v3) + "\n\t") > v3) {
                        fmt.Println(fi10, fi10)
                        v11 = v11
                    } else {
                        if true {
                            var v13 uint64 = 18446744073709551615
                            _ = v13
                        }
                    }
                }
                fmt.Println(v5, v3)
                var v14 uint32 = 292
                _ = v14
            }
            _ = fi10
            _ = v11
        }
    case "":
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    v15 := v2[807]
    var v16 int32 = -746
    if len(v8) > 0 {
        v8[0] = v15
    }
    v5 = float64(float32((1e10 + -2.25)))
    fmt.Println(v16, v5, v15)
    var v17 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
