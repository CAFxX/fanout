package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 bool; F1 int16 }

func main() {
    v1 := []string{"foo bar", "\n\t", ("" + "hello")}
    var v2 int8 = -9
    var v3 uint16 = 2
    v4 := make(map[uint32]int)
    var v5 int8 = -22
    v6 := make(map[string]string)
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    switch "" {
    case "hello":
        v5 -= (v2 >> ((1) & 63))
    case "":
        v5 -= 27
    case "a":
        v3 = v3
    }
    fmt.Println(v5, v5)
    for fi7 := 0; fi7 < 16; fi7++ {
        v3 += (v3 & v3)
        v8 := GenS1{F0: 0, F1: "\n\t"}
        _ = fi7
        _ = v8
    }
    var v9 int8 = (v2 - v5)
    var v10 int32 = -638
    switch "foo bar" {
    case "hello":
        switch -555 {
        case 4294967295:
            v11 := GenS1{F0: 2, F1: ("" + ("foo bar" + "\n\t"))}
            _ = v11
        default:
            v5 = 1
        }
    }
    v12 := GenS2{F0: true, F1: 1}
    v2 += v5
    v10 = (v10 ^ (v10 - -959))
    for fi13 := 0; fi13 < 6; fi13++ {
        if true {
            switch "hello" {
            case "\n\t":
                fmt.Println(v3)
            case "hello":
                v5 = -54
            default:
                v2 -= v2
            }
        }
        v3 -= v3
        var v14 uint = 897
        _ = fi13
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 6; fi15++ {
            v1 = append(v1, ("a" + ""))
            var v16 int8 = 31
            _ = fi15
            _ = v16
        }
        v1 = append(v1, ("foo bar" + ("a" + "")))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
