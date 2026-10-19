package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int64; F2 int8 }

func main() {
    var v1 uint16 = 32767
    v2 := []string{((("\n\t" + (("foo bar" + "a") + ("\n\t" + "hello"))) + ("\n\t" + "")) + "\n\t"), "\n\t"}
    v3 := make(map[int]int)
    var v4 uint32 = 0
    v5 := make(map[int]string)
    v6 := []string{"\n\t", "foo bar", "foo bar", "a"}
    v7 := GenS1{F0: 362, F1: v4}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 32767 {
    case -136:
        switch int(v4) {
        case 42:
            v8 := GenS1{F0: -584, F1: v4}
            _ = v8
        default:
            v1 = (v1 & v1)
        }
    case 0:
        var v9 int64 = -9223372036854775808
        _ = v9
    default:
        v10 := GenS2{F0: 1, F1: 0, F2: -128}
        _ = v10
    }
    fmt.Println(v4)
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v4)
    case "hello":
        var v11 int64 = -1
        _ = v11
    case "":
        for fi12 := 0; fi12 < 4; fi12++ {
            for fi13 := 0; fi13 < 16; fi13++ {
                v1 -= 32767
                var v14 uint64 = 18446744073709551615
                delete(v3, 4294967295)
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    var v15 int = 731
    v15 = 559
    fmt.Println(v15, v1, v1)
    if true {
        switch v15 {
        case 2147483647:
            v1 -= (v1 / ((v1) | 1))
        }
        v1 = (v1 - v1)
        v15 += (v15 | (v15 >> ((481) & 63)))
    } else {
        v4 += (v4 & (^(v4)))
        fmt.Println(v15)
    }
    var v16 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
