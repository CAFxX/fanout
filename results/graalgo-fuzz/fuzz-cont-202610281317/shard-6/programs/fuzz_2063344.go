package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int16; F1 uint16; F2 uint64 }
func genFn0(p0 int32, p1 uint8) int {
    p1 = (p1 / (((p1 & (p1 & p1))) | 1))
    p1 = p1
    var v5 bool = (0 < (-(p0)))
    if (v5 && (-2.25 == (((0.1 - 0.1) + float64(1e10)) * -2.25))) {
        if (!false) {
            v6 := []int{1, -2147483648, 1, -495}
            _ = v6
        } else {
            v7 := GenS1{F0: 412, F1: 2, F2: 616}
            if (v5 || v5) {
                var v8 uint64 = 18446744073709551615
                p0 -= p0
                var v9 float64 = 0.1
                _ = v8
                _ = v9
            }
            _ = v7
        }
    } else {
        fmt.Println(p1, p0)
        if false {
            fmt.Println(p0, p1, v5)
        } else {
            var v10 int64 = 480
            _ = v10
        }
    }
    v11 := []int{-456, 32767, 32768, 4294967295}
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return 4294967295
}

func main() {
    var v1 string = ("foo bar" + "\n\t")
    v2 := make(map[int]string)
    v3 := []int{32767, -571, 0, 32768}
    var v4 uint8 = 237
    fmt.Println(v4, v1)
    var v12 int16 = 1000
    v13 := GenS2{F0: 127, F1: (v1 + "hello")}
    v12 = 32767
    v4 = 42
    _ = len(v3)
    v14 := GenS2{F0: 2147483647, F1: ("hello" + v1)}
    switch genFn0(1, v4) {
    case 9223372036854775807:
        v1 = v1
    case 128:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v15 float64 = -2.25
    var v16 uint = 255
    if (-32768 <= v12) {
        if false {
            v17 := GenS1{F0: -53, F1: v16, F2: 18446744073709551615}
            _ = v17
        }
        v18 := GenS3{F0: (^(v12)), F1: 255, F2: 608}
        _ = v18
    }
    fmt.Println(v4, v16)
    for fi19 := 0; fi19 < 18; fi19++ {
        if true {
            v1 = "foo bar"
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
