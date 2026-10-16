package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 bool }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }
func genFn0(p0 int64, p1 uint8) int32 {
    var v7 int8 = 54
    switch "" {
    case "foo bar":
        switch "a" {
        case "":
            fmt.Println(p0, v7, p1)
        }
    default:
        var v8 int32 = 739
        _ = v8
    }
    v7 = 0
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "\n\t"
    v3 := make(map[uint32]bool)
    var v4 int16 = 0
    var v5 uint64 = 2
    var v6 int32 = -897
    v1[2147483648] = 9223372036854775807
    v9 := GenS1{F0: 42, F1: -732, F2: false}
    switch -9223372036854775808 {
    case 32767:
        v10 := GenS2{F0: "foo bar", F1: 2}
        _ = v10
    case 2147483647:
        switch v2 {
        case "foo bar":
            v11 := GenS2{F0: "hello", F1: 26}
            _ = v11
        case "":
            v4 -= 42
        case "a":
            v6 += v6
        default:
            v5 = v5
        }
    default:
        var v12 float64 = 0.1
        _ = v12
    }
    if true {
        var v13 int = -1
        _ = v13
    } else {
        var v14 int8 = 100
        var v15 uint = 757
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            v2 = "a"
            v2 = "a"
            _ = fi17
        }
        delete(v1, 747)
        switch (v2 + "foo bar") {
        case "a":
            _ = len(v1)
        }
        _ = fi16
    }
    v18 := GenS2{F0: v2, F1: 127}
    v19 := GenS3{F0: (!false)}
    v20 := GenS1{F0: 843, F1: 0, F2: (0 == (v4 * 1))}
    v4 = v4
    if false {
        v21 := []int{}
        fmt.Println(v2, v2)
        _ = v21
    } else {
        v22 := GenS4{F0: 444}
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
