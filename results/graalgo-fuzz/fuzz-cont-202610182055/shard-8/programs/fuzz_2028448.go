package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }

func main() {
    var v1 uint32 = 56
    var v2 uint16 = 42
    v3 := make(map[int]string)
    v4 := []string{"a", ("hello" + ""), "\n\t"}
    if true {
        var v5 int32 = -1
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 int = (fi6 << ((18446744073709551615) & 63))
            v8 := GenS1{F0: fi6, F1: (3.14159 <= 1e10)}
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    } else {
        switch "" {
        case "a":
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: (fi9 % ((fi9) | 1)), F1: true}
                switch fi9 {
                case 1:
                    fmt.Println(v2, v2)
                }
                _ = fi9
                _ = v10
            }
        case "hello":
            v1 -= (v1 % ((v1) | 1))
        default:
            v11 := GenS1{F0: 2147483647, F1: false}
            _ = v11
        }
        v3[9223372036854775807] = "foo bar"
    }
    var v12 uint = 521
    v12 = v12
    v1 = v1
    v12 = v12
    for fi13 := 0; fi13 < 3; fi13++ {
        v12 -= v12
        if true {
            var v14 int8 = 64
            fmt.Println(v14, v2)
            _ = v14
        } else {
            fmt.Println(v12, v1, fi13)
            if false {
                fmt.Println(v12)
            } else {
                fmt.Println(fi13, v2)
                delete(v3, fi13)
                v15 := v3[2147483647]
                _ = v15
            }
        }
        _ = fi13
    }
    var v16 int = 128
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := v3[fi17]
        v12 = 0
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
