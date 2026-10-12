package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int8; F1 int; F2 string }

func main() {
    v1 := []string{}
    var v2 int32 = -1
    var v3 int32 = v2
    v4 := []int{}
    v5 := []int{}
    var v6 int = -1
    var v7 string = ("a" + "\n\t")
    var v8 uint8 = 241
    if (false && true) {
        _ = len(v4)
    } else {
        var v9 bool = true
        _ = v9
    }
    switch v7 {
    case "foo bar":
        v7 = v7
    case "\n\t":
        v10 := GenS1{F0: 42, F1: 3}
        _ = v10
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v11 string = v7
    for fi12 := 0; fi12 < 6; fi12++ {
        if false {
            var v13 uint8 = (v8 / ((v8) | 1))
            _ = v13
        }
        _ = fi12
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    switch v7 {
    case "hello":
        v11 = "hello"
    case "foo bar":
        v2 = (v3 >> ((18446744073709551615) & 63))
    case "\n\t":
        for fi14 := 0; fi14 < 16; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := GenS2{F0: -76, F1: fi15, F2: "hello"}
                v2 = v3
                _ = fi15
                _ = v16
            }
            v11 = v7
            for fi17 := 0; fi17 < 2; fi17++ {
                v18 := GenS1{F0: 499, F1: 255}
                v4 = append(v4, (fi14 * 2147483647))
                _ = fi17
                _ = v18
            }
            _ = fi14
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
