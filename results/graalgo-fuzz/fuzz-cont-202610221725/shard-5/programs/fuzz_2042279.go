package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint64 }
type GenS2 struct { F0 int; F1 uint8; F2 int64 }
type GenS3 struct { F0 int; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint64; F1 int32 }
func genFn0() uint8 {
    var v6 uint8 = 42
    fmt.Println(v6)
    v7 := GenS1{F0: 875, F1: 127, F2: 4294967295}
    v6 -= v6
    var v8 int = 263
    _ = v6
    _ = v7
    _ = v8
    return v6
}
func genFn1(p0 int, p1 uint16, p2 string) string {
    fmt.Println(p1, p1, p2)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    var v1 string = (("" + ((("foo bar" + "hello") + ("" + "\n\t")) + ("\n\t" + "a"))) + "foo bar")
    var v2 int = 32767
    var v3 uint8 = 128
    var v4 uint8 = 127
    v5 := make(map[int]int)
    v1 = ""
    v9 := GenS2{F0: v2, F1: v3, F2: int64(v4)}
    v10 := GenS2{F0: 42, F1: 0, F2: -547}
    v11 := []int{}
    var v12 int8 = -1
    v3 = 1
    v4 -= (v3 - genFn0())
    v12 -= (-(v12))
    switch (v1 + ("hello" + v1)) {
    case "a":
        v12 -= (v12 + v12)
    case "":
        fmt.Println(v4, v3, v2)
    case "hello":
        fmt.Println(v4, v12)
    default:
        var v13 float64 = 0.0
        _ = v13
    }
    if true {
        if (true && (42 < -1)) {
            v11 = append(v11, v2)
            v14 := GenS3{F0: (v2 * 4294967295), F1: (^(v4)), F2: (v12 + v12)}
            _ = v14
        } else {
            v1 = v1
            for fi15 := 0; fi15 < 16; fi15++ {
                for fi16 := 0; fi16 < 19; fi16++ {
                    v3 -= (v4 | v4)
                    _ = fi16
                }
                _ = fi15
            }
        }
        switch (("\n\t" + "") + ("a" + v1)) {
        case "a":
            v17 := GenS3{F0: (v2 % ((v2) | 1)), F1: (v3 % ((genFn0()) | 1)), F2: v12}
            _ = v17
        default:
            var v18 float32 = (0.1 / -2.25)
            _ = v18
        }
    }
    v19 := GenS2{F0: v2, F1: v4, F2: 42}
    _ = len(v11)
    v20 := GenS4{F0: 4294967295, F1: 2147483647}
    v12 = (v12 % ((42) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
