package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 string; F2 int8 }
type GenS3 struct { F0 int32; F1 int32; F2 bool }
func genFn0(p0 uint32, p1 int16) uint {
    switch "\n\t" {
    case "":
        if (5 <= 1) {
            p0 -= (p0 / (((p0 ^ p0)) | 1))
        }
    }
    fmt.Println(p1, p0)
    p0 -= (p0 % (((p0 | 42)) | 1))
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 bool = false
        _ = fi6
        _ = v7
    }
    p1 -= (p1 & (-(p1)))
    _ = p0
    _ = p1
    return 168
}

func main() {
    var v1 uint32 = 42
    var v2 int64 = -265
    var v3 uint = 255
    v4 := []string{"foo bar", "a", "foo bar", ""}
    var v5 uint32 = (^(v1))
    if (!false) {
        var v8 uint64 = uint64(v3)
        fmt.Println(v3, v8)
        fmt.Println(v5, v8, v3)
        _ = v8
    }
    v9 := GenS1{F0: "a"}
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    var v11 int64 = v2
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(fi12, v11)
        _ = fi12
    }
    switch "hello" {
    case "foo bar":
        v13 := GenS1{F0: ("a" + "foo bar")}
        _ = v13
    }
    v14 := GenS2{F0: (2 != 42), F1: "foo bar", F2: -1}
    v15 := GenS3{F0: -59, F1: -1, F2: (true && ((221 != 65535) || false))}
    if (!((int8(v1) <= 0) || false)) {
        fmt.Println(v1)
        if len(v4) > 0 {
            v4[0] = "a"
        }
    }
    v5 += v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v2, v2, v2)
    var v16 float64 = 0.1
    var v17 int8 = 0
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
