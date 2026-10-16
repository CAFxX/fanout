package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int32; F2 uint32 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int64, p1 int64) uint {
    v7 := GenS1{F0: ("hello" + "\n\t"), F1: 42, F2: 42}
    v8 := GenS2{F0: (true && false)}
    v9 := GenS2{F0: false}
    v10 := []int{9223372036854775807, 42}
    v11 := GenS2{F0: (("" > "foo bar") || (523 != 957))}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 65535
}
func genFn1(p0 uint, p1 int8, p2 string) int32 {
    v12 := GenS3{F0: 42, F1: 1, F2: 2147483648}
    var v13 int32 = 42
    var v14 float64 = 3.14159
    fmt.Println(v13)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return v13
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    v4 := make(map[string]bool)
    var v5 bool = false
    var v6 uint64 = 1
    if false {
        var v15 uint32 = 182
        _ = v15
    }
    if (uint32(v6) == 42) {
        v2[4294967295] = "\n\t"
        v6 = 18446744073709551615
    }
    v6 -= (v6 - (v6 >> ((801) & 63)))
    fmt.Println(v6, v6)
    var v16 uint = genFn0(-1, 1)
    var v17 uint8 = 2
    v18 := GenS2{F0: v5}
    v19 := GenS4{F0: 1, F1: -207}
    var v20 int64 = 696
    switch 2147483647 {
    case -9223372036854775808:
        v16 = 42
    case 9223372036854775807:
        v21 := v4[""]
        _ = v21
    }
    var v22 int16 = -81
    v23 := []int{820, -699}
    if false {
        v16 -= genFn0(v20, 234)
        for fi24 := 0; fi24 < 4; fi24++ {
            for fi25 := 0; fi25 < 1; fi25++ {
                v20 = v20
                switch "\n\t" {
                case "a":
                    var v26 float64 = 1.5
                    _ = v26
                case "foo bar":
                    var v27 uint8 = 1
                    _ = v27
                case "\n\t":
                    var v28 uint = (v16 ^ genFn0(int64(v6), v20))
                    _ = v28
                }
                fmt.Println(v20, fi25)
                _ = fi25
            }
            if (!v5) {
                fmt.Println(v17, fi24)
                fmt.Println(v17, v17, v16)
                v29 := GenS5{F0: -3}
                _ = v29
            } else {
                v30 := GenS2{F0: v5}
                if ((1.5 / ((-2.25 / float32(1.5)) - 1e10)) > 3.14159) {
                    var v31 uint16 = 42
                    _ = v31
                }
                v17 = (v17 - v17)
                _ = v30
            }
            _ = fi24
        }
        var v32 int64 = (v20 & v20)
        _ = v32
    }
    var v33 uint64 = 128
    delete(v3, 611)
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
