package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 string = "a"
    v3 := make(map[string]int)
    var v4 bool = false
    v5 := []int{}
    v6 := []uint32{}
    if false {
        var v7 bool = v4
        _ = v7
    }
    switch v2 {
    case "hello":
        fmt.Println(v2, v2)
    case "foo bar":
        fmt.Println(v2, v2)
    }
    var v8 uint8 = 230
    var v9 int16 = 32767
    v10 := GenS1{F0: 0}
    var v11 int16 = v9
    v5 = append(v5, -805)
    v9 += (v11 - (v9 % ((v11) | 1)))
    if false {
        v1[v2] = 569
        var v12 uint64 = 0
        switch "a" {
        case "":
            v13 := v3[("hello" + (v2 + "hello"))]
            _ = v13
        case "a":
            switch (v2 + "hello") {
            case "hello":
                v14 := GenS1{F0: 1}
                _ = v14
            case "a":
                var v15 string = v2
                _ = v15
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        case "hello":
            v12 = v12
        }
        _ = v12
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v17 := GenS1{F0: 42}
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 1}
    for fi19 := 0; fi19 < 9; fi19++ {
        v11 += (v9 & v11)
        _ = fi19
    }
    v9 = v11
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 uint16 = 137
        for fi22 := 0; fi22 < 18; fi22++ {
            if (18446744073709551615 < 418) {
                fmt.Println(v8)
            } else {
                v9 -= v11
            }
            v23 := GenS1{F0: v21}
            _ = fi22
            _ = v23
        }
        v3["foo bar"] = (^(fi20))
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
