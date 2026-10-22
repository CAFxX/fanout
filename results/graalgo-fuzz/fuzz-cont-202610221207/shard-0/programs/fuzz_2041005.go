package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 bool; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
func genFn0(p0 int8, p1 uint16, p2 uint) string {
    v4 := []int{-2147483648, 32767}
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int16 = 738
        if (!false) {
            v7 := GenS1{F0: p2}
            fmt.Println(p0, p1)
            fmt.Println(v6, p0)
            _ = v7
        } else {
            v8 := GenS2{F0: "\n\t", F1: true, F2: 4294967295}
            var v9 float64 = ((3.14159 * 0.0) - (1e10 / 1.5))
            _ = v8
            _ = v9
        }
        if (!false) {
            p2 = p2
        }
        _ = fi5
        _ = v6
    }
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    v10 := GenS3{F0: -65}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v4
    return "foo bar"
}

func main() {
    var v1 int32 = -349
    v2 := make(map[uint32]string)
    var v3 string = "hello"
    fmt.Println(v3)
    v3 = (v3 + "foo bar")
    v11 := []int{42, -9223372036854775808}
    _ = len(v11)
    v12 := GenS1{F0: 1}
    var v13 uint8 = 235
    v13 += (v13 & v13)
    v14 := GenS4{F0: (-713 == 406)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
