package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }
type GenS2 struct { F0 uint16; F1 int32; F2 int }
func genFn0(p0 uint64, p1 uint) int64 {
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 uint = 611
        fmt.Println(fi4, p1, v5)
        _ = fi4
        _ = v5
    }
    fmt.Println(p1, p0)
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(fi6, p0, p0)
        v7 := GenS1{F0: "foo bar", F1: 32767, F2: (-(fi6))}
        p0 = uint64(p1)
        _ = fi6
        _ = v7
    }
    p1 += p1
    v8 := GenS2{F0: 970, F1: 2147483647, F2: 2147483647}
    _ = p0
    _ = p1
    _ = v8
    return 176
}

func main() {
    v1 := []int{-851, 128, 9223372036854775807, -856}
    var v2 string = "hello"
    v3 := []string{""}
    v2 = "hello"
    var v9 uint32 = 481
    v3 = append(v3, "\n\t")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = append(v1, 4294967295)
    if true {
        var v10 uint16 = 910
        _ = v10
    } else {
        fmt.Println(v2, v2)
    }
    v9 = 4294967295
    fmt.Println(v2, v9)
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v11 := GenS2{F0: uint16(v9), F1: 215, F2: 1}
    var v12 string = (v2 + ("hello" + ("" + (v2 + v2))))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, 4294967295)
    v3 = append(v3, v2)
    var v13 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
