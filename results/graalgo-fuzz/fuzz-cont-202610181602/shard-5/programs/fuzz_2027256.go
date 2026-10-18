package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 string }

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]string)
    var v3 int = 32767
    v4 := []uint32{2, 65535}
    var v5 uint = 679
    _ = len(v2)
    fmt.Println(v5)
    v6 := GenS1{F0: 1, F1: 1}
    v7 := GenS2{F0: "hello"}
    v1 = (v1 & v1)
    var v8 string = "hello"
    fmt.Println(v1, v8, v8)
    switch v8 {
    case "a":
        v9 := GenS2{F0: ("foo bar" + "")}
        _ = v9
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if false {
        fmt.Println(v3, v3)
        switch v3 {
        case 1:
            v4 = append(v4, 444)
        case 0:
            v5 = uint(v1)
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            v8 = v8
            var v11 int = 0
            if (false && true) {
                var v12 int8 = 127
                v5 -= (v5 | (v5 >> ((255) & 63)))
                _ = v12
            }
            _ = fi10
            _ = v11
        }
    } else {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 int = fi13
            fmt.Println(v5, v3)
            _ = fi13
            _ = v14
        }
        v15 := GenS2{F0: (("foo bar" + "foo bar") + v8)}
        _ = v15
    }
    fmt.Println(v1)
    v1 += v1
    switch v8 {
    case "a":
        for fi16 := 0; fi16 < 1; fi16++ {
            v1 -= (v1 ^ (v1 - v1))
            v8 = ((v8 + v8) + v8)
            _ = fi16
        }
    default:
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := v2[(fi17 | fi17)]
            _ = fi17
            _ = v18
        }
    }
    fmt.Println(v5, v3, v3)
    var v19 uint = uint(v1)
    if false {
        v1 -= (v1 * (v1 + v1))
        var v20 int = v3
        _ = v20
    }
    v19 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
