package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 bool = (0.1 <= (float32(0.0) * -2.25))
    var v2 bool = (226 == 127)
    v3 := make(map[string]string)
    _ = len(v3)
    v4 := []int{}
    v5 := GenS1{F0: -623, F1: 108}
    v2 = (!(18446744073709551615 >= 42))
    switch "\n\t" {
    case "hello":
        switch "foo bar" {
        case "a":
            v2 = v2
        }
    case "\n\t":
        v1 = v2
    }
    var v6 int8 = 127
    v7 := GenS1{F0: 32768, F1: v6}
    if (-645 < 1) {
        fmt.Println(v1)
    } else {
        v2 = (!v2)
        var v8 uint8 = 0
        for fi9 := 0; fi9 < 13; fi9++ {
            if len(v4) > 0 {
                v4[0] = (fi9 * fi9)
            }
            v1 = v2
            fmt.Println(v1, v8, v1)
            _ = fi9
        }
        _ = v8
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (!(!false)) {
            switch (fi10 & fi10) {
            case 0:
                v11 := GenS2{F0: 32767, F1: 18446744073709551615}
                _ = v11
            default:
                v4 = append(v4, fi10)
            }
            v12 := GenS1{F0: fi10, F1: (v6 | v6)}
            switch fi10 {
            case 120:
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 uint8 = 0
                    _ = fi13
                    _ = v14
                }
            case 9223372036854775807:
                v2 = (fi10 > fi10)
            default:
                var v15 float32 = (1.5 * (0.0 - 3.14159))
                _ = v15
            }
            _ = v12
        }
        _ = fi10
    }
    var v16 int32 = -64
    fmt.Println(v16, v2)
    _ = len(v3)
    if (!((563 == 1) && (false && v1))) {
        v2 = true
    }
    delete(v3, (("foo bar" + "foo bar") + ""))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
