package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 string }

func main() {
    var v1 float32 = 1e10
    v2 := []uint32{}
    var v3 uint8 = 2
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]bool)
    var v6 int32 = -722
    var v7 uint8 = 255
    var v8 int32 = v6
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: 0, F1: true, F2: (("a" + "\n\t") + "hello")}
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 168
    v8 = v6
    switch "a" {
    case "":
        var v12 uint16 = v11
        _ = v12
    }
    v5[uint32(v11)] = true
    switch -108 {
    case 204:
        var v13 int8 = 100
        _ = v13
    case 208:
        v2 = append(v2, 295)
    }
    fmt.Println(v7, v3, v6)
    if false {
        v7 = v3
        v14 := GenS1{F0: 1, F1: (!(!("foo bar" > "foo bar"))), F2: ""}
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v8, v7)
        }
        _ = v14
    } else {
        var v15 bool = true
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 float64 = ((0.1 - 3.14159) + -2.25)
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS1{F0: 18446744073709551615, F1: (fi18 == (^(fi18))), F2: ""}
            _ = len(v5)
            _ = fi18
            _ = v19
        }
        v2 = append(v2, 2147483648)
        _ = fi16
        _ = v17
    }
    var v20 int16 = 42
    v6 -= (v6 % ((v6) | 1))
    var v21 float32 = ((-2.25 + v1) * v1)
    v2 = append(v2, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
