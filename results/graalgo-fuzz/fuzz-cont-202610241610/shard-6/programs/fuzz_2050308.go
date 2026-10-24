package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint32 }

func main() {
    v1 := []uint32{4294967295, 1, 402, 4294967295}
    var v2 bool = (0 != 32767)
    var v3 uint8 = 255
    var v4 int8 = -128
    var v5 int64 = 1
    _ = len(v1)
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v6 uint8 = (v3 + (v3 | (v3 + 127)))
        _ = v6
    }
    v3 = (v3 ^ 255)
    _ = len(v1)
    v7 := GenS1{F0: 42, F1: (v5 & 237), F2: 4294967295}
    var v8 int8 = -1
    fmt.Println(v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        if v2 {
            v1 = append(v1, 2147483648)
            switch "a" {
            case "":
                for fi10 := 0; fi10 < 16; fi10++ {
                    var v11 float32 = 0.1
                    _ = fi10
                    _ = v11
                }
            case "a":
                v4 = (^(v8))
            case "hello":
                v2 = false
            }
        }
        if len(v1) > 0 {
            v1[0] = 42
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        switch ("foo bar" + "a") {
        case "hello":
            var v13 string = "\n\t"
            _ = v13
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint64 = 18446744073709551615
                _ = fi14
                _ = v15
            }
        }
        v8 += v8
        if len(v1) > 0 {
            v1[0] = 65535
        }
        _ = fi12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
