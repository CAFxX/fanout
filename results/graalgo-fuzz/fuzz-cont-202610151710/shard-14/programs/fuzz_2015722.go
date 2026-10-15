package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 bool; F1 string; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint16; F1 uint64 }
func genFn0() string {
    var v5 string = ("a" + "hello")
    switch -2147483648 {
    case -9223372036854775808:
        v6 := GenS1{F0: true, F1: "a"}
        _ = v6
    case 32768:
        v7 := GenS1{F0: false, F1: ""}
        _ = v7
    default:
        fmt.Println(v5)
    }
    _ = v5
    return v5
}

func main() {
    var v1 uint8 = 113
    v2 := []int{650, -513}
    var v3 uint = 375
    v4 := []int{}
    _ = len(v4)
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int64 = 0
        _ = fi8
        _ = v9
    }
    var v10 uint = (v3 ^ (^(v3)))
    v11 := GenS2{F0: (431 < int64(v3)), F1: "a", F2: 454}
    v12 := GenS2{F0: (42 > 679), F1: "foo bar", F2: 2147483647}
    fmt.Println(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (true && true) {
        if len(v2) > 0 {
            v2[0] = 1
        }
        v13 := GenS3{F0: v1, F1: 2147483647}
        _ = v13
    }
    v14 := GenS4{F0: 2147483647, F1: (65535 <= uint32(v3))}
    var v15 int32 = 1
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS5{F0: 42, F1: 205}
        _ = len(v4)
        _ = fi16
        _ = v17
    }
    var v18 int64 = 1
    switch -9223372036854775808 {
    case 127:
        for fi19 := 0; fi19 < 14; fi19++ {
            v15 += v15
            _ = fi19
        }
    case -1000:
        fmt.Println(v3, v10, v1)
    case 42:
        var v20 int8 = 1
        _ = v20
    }
    v21 := GenS2{F0: true, F1: ("\n\t" + "a"), F2: 2147483647}
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 int64 = v18
        v2 = append(v2, (fi22 + fi22))
        _ = fi22
        _ = v23
    }
    fmt.Println(v18, v10)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
