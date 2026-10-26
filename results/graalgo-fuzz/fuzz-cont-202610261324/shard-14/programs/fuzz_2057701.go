package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 string }
func genFn0(p0 bool, p1 uint64) int64 {
    v7 := []int{2147483647}
    if len(v7) > 0 {
        v7[0] = 42
    }
    v8 := GenS1{F0: p0, F1: 1}
    if len(v7) > 0 {
        v7[0] = -9223372036854775808
    }
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -623
}
func genFn1() uint64 {
    switch 9223372036854775807 {
    case 9223372036854775807:
        v9 := GenS2{F0: 0, F1: 673}
        _ = v9
    case 127:
        var v10 uint32 = 0
        _ = v10
    case -1:
        var v11 uint = 18446744073709551615
        _ = v11
    }
    var v12 uint64 = 280
    var v13 int64 = 42
    for fi14 := 0; fi14 < 9; fi14++ {
        fmt.Println(fi14)
        for fi15 := 0; fi15 < 17; fi15++ {
            fmt.Println(fi14)
            _ = fi15
        }
        _ = fi14
    }
    _ = v12
    _ = v13
    return 121
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 float32 = 1e10
    v4 := make(map[int]string)
    v5 := []int{901}
    var v6 uint64 = 0
    for fi16 := 0; fi16 < 5; fi16++ {
        v6 += v6
        v17 := GenS3{F0: 65535}
        var v18 int32 = -2147483648
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v3, v3)
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    var v19 int8 = 42
    if false {
        v20 := GenS3{F0: 65535}
        v21 := GenS4{F0: "\n\t", F1: 255}
        _ = v20
        _ = v21
    }
    v22 := GenS5{F0: 689}
    v6 -= 0
    fmt.Println(v3, v6, v3)
    fmt.Println(v6, v6, v6)
    for fi23 := 0; fi23 < 17; fi23++ {
        _ = len(v2)
        _ = fi23
    }
    v24 := GenS6{F0: 329, F1: (true && (!(!true)))}
    switch (("hello" + "hello") + (("" + ("a" + "foo bar")) + "")) {
    case "\n\t":
        v25 := GenS7{F0: ""}
        _ = v25
    case "hello":
        v26 := GenS5{F0: 2147483647}
        _ = v26
    }
    var v27 float32 = (1e10 + ((v3 + 3.14159) - float32(1.5)))
    v28 := v2["a"]
    fmt.Println(v6, v6)
    if ((0.1 - 3.14159) != 3.14159) {
        var v29 bool = false
        var v30 int = v28
        _ = v29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
