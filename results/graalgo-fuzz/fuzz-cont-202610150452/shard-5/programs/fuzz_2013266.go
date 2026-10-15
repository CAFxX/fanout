package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint16; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 bool; F2 uint16 }
func genFn0(p0 int16, p1 int64) string {
    switch "\n\t" {
    case "foo bar":
        fmt.Println(p1, p0)
    case "\n\t":
        fmt.Println(p1, p0)
    case "hello":
        v6 := GenS1{F0: 68, F1: "a"}
        _ = v6
    }
    v7 := GenS2{F0: 518}
    v8 := GenS3{F0: 127, F1: 0, F2: 995}
    v9 := []int{-82}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return ((("a" + "foo bar") + "hello") + ("hello" + "a"))
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    var v5 float64 = -2.25
    v10 := []int{32768, 32768, 0, 127}
    if true {
        var v11 float64 = v5
        _ = v11
    }
    if (42 != 1) {
        v12 := GenS4{F0: (("a" + "") >= "a"), F1: 4294967295, F2: 1}
        v13 := GenS5{F0: 95}
        _ = v12
        _ = v13
    } else {
        var v14 bool = v1
        if false {
            if (true || true) {
                fmt.Println(v14, v14, v14)
                var v15 int32 = -2147483648
                fmt.Println(v15)
                _ = v15
            }
            fmt.Println(v1, v5)
            fmt.Println(v14, v1)
        }
        v16 := GenS6{F0: 199, F1: (!(18446744073709551615 != 0)), F2: 706}
        _ = v14
        _ = v16
    }
    var v17 uint8 = 197
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 float64 = (0.1 - (v5 - (v5 * v5)))
        fmt.Println(v5, fi18)
        v20 := GenS6{F0: -1, F1: (0 <= (v17 / ((2) | 1))), F2: 2}
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println(v17, v1)
    var v21 uint64 = 148
    var v22 bool = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
