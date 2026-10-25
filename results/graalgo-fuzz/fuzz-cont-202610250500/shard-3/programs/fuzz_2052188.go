package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
func genFn0() uint64 {
    if false {
        v7 := []int{9223372036854775807, -1, 128}
        _ = v7
    }
    switch "foo bar" {
    case "foo bar":
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := []int{32767, (fi8 * (fi8 - fi8))}
            fmt.Println(fi8)
            _ = len(v9)
            _ = fi8
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 string = "foo bar"
            for fi12 := 0; fi12 < 1; fi12++ {
                fmt.Println(v11, fi12, fi12)
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
    }
    return 1
}

func main() {
    var v1 float32 = (0.1 * -2.25)
    v2 := make(map[uint32]int)
    var v3 string = "a"
    var v4 uint64 = 1
    v5 := make(map[uint32]string)
    var v6 int64 = -9223372036854775808
    v6 = (v6 | v6)
    v3 = ((v3 + "") + (v3 + "hello"))
    if true {
        switch (v3 + "a") {
        case "hello":
            v6 -= (v6 * (^(v6)))
        case "foo bar":
            v6 = (v6 ^ -9223372036854775808)
        }
    } else {
        fmt.Println(v1, v3, v4)
    }
    var v13 string = (v3 + (v3 + ((((v3 + v3) + v3) + ((v3 + (v3 + v3)) + (v3 + "hello"))) + "a")))
    var v14 float32 = v1
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := []int{(^(fi15)), fi15, (fi15 >> ((1) & 63)), fi15}
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 648, F1: v3, F2: 42}
    v3 = "hello"
    v18 := []int{32767, 115}
    var v19 uint32 = 130
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 int = (-(fi20))
        _ = fi20
        _ = v21
    }
    var v22 int64 = v6
    for fi23 := 0; fi23 < 3; fi23++ {
        v13 = v3
        v6 = v6
        _ = fi23
    }
    fmt.Println(v14, v1, v14)
    v24 := GenS1{F0: v19, F1: "a", F2: -186}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
