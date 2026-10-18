package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    var v1 string = "foo bar"
    v2 := []string{(v1 + "foo bar"), "", v1, v1}
    var v3 uint16 = 1
    v4 := make(map[int]int)
    v3 = v3
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint32 = 525
        _ = fi5
        _ = v6
    }
    var v7 string = "\n\t"
    switch -9223372036854775808 {
    case 32767:
        fmt.Println(v1, v1)
    case -2147483648:
        fmt.Println(v1, v1)
    case 2147483647:
        v7 = v1
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v8 := v4[-1]
    v3 += (v3 + (v3 % ((v3) | 1)))
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println(v8, v3, v1)
    fmt.Println(v7, v8)
    v1 = (((v7 + v1) + v7) + ((v7 + (v7 + (v7 + (("hello" + v1) + v7)))) + v7))
    var v9 float64 = ((-2.25 + float64(0.1)) * float64(1.5))
    _ = len(v4)
    v1 = ("a" + (v7 + v7))
    if false {
        switch v8 {
        case 27:
            v10 := GenS1{F0: 0, F1: 78}
            _ = v10
        }
        var v11 uint16 = v3
        _ = v11
    } else {
        fmt.Println(v1, v8, v9)
        v7 = ("a" + ("hello" + ""))
        v3 = 32767
    }
    if true {
        var v12 uint8 = 71
        _ = v12
    } else {
        var v13 uint32 = 0
        var v14 uint32 = v13
        var v15 int64 = 0
        _ = v13
        _ = v14
        _ = v15
    }
    v1 = ""
    var v16 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
