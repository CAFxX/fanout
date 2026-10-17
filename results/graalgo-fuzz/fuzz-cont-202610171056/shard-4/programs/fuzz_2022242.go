package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 bool }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0(p0 uint16, p1 int32) uint32 {
    p1 += -1
    fmt.Println(p0, p0)
    fmt.Println(p0, p1)
    var v7 uint16 = 1
    var v8 string = ""
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int64 = -1
    var v2 string = "\n\t"
    var v3 uint32 = 2147483648
    var v4 int32 = -1
    var v5 int32 = v4
    v6 := make(map[uint32]bool)
    v9 := GenS1{F0: 0, F1: (!false), F2: true}
    var v10 uint16 = 65535
    v2 = "\n\t"
    v11 := v6[v3]
    if true {
        var v12 string = (v2 + (v2 + ("" + ((v2 + ("hello" + "foo bar")) + "a"))))
        var v13 string = ((v2 + v12) + "")
        _ = v12
        _ = v13
    } else {
        if true {
            fmt.Println(v2, v4)
        } else {
            _ = len(v6)
            v10 = 0
            v10 = (v10 << ((uint(v10)) & 63))
        }
        fmt.Println(v5)
        for fi14 := 0; fi14 < 13; fi14++ {
            switch fi14 {
            case 9223372036854775807:
                fmt.Println(v4)
            default:
                v2 = ""
            }
            fmt.Println(v2, v3, v3)
            _ = fi14
        }
    }
    var v15 int = 1
    var v16 uint64 = 1
    if v11 {
        var v17 uint16 = (v10 << ((uint(v16)) & 63))
        _ = v17
    }
    v18 := []int{}
    v19 := GenS2{F0: 1, F1: v15}
    if v11 {
        var v20 int8 = -128
        var v21 uint64 = 42
        _ = v20
        _ = v21
    }
    switch v2 {
    case "foo bar":
        v22 := GenS1{F0: 42, F1: (v2 <= ("foo bar" + (v2 + ""))), F2: true}
        _ = v22
    default:
        fmt.Println(v2, v4, v3)
    }
    switch v2 {
    case "hello":
        v23 := GenS3{F0: (("a" + ("" + ("\n\t" + v2))) + v2), F1: (v10 - (v10 ^ v10))}
        _ = v23
    case "\n\t":
        var v24 int = v15
        _ = v24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 float32 = 3.14159
        fmt.Println(v4, fi25)
        _ = fi25
        _ = v26
    }
    v27 := GenS2{F0: 328, F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
