package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []uint32{4294967295, 2147483648}
    var v2 uint32 = 139
    v3 := make(map[string]string)
    var v4 int16 = -220
    var v5 int = -9223372036854775808
    fmt.Println(v4)
    v6 := v3[("foo bar" + ("foo bar" + "a"))]
    var v7 uint64 = 825
    var v8 bool = true
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint32 = v2
        v6 = v6
        var v11 int32 = -466
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = len(v3)
    v12 := v3[v6]
    fmt.Println(v2, v5, v7)
    switch "\n\t" {
    case "foo bar":
        switch 1 {
        case -2147483648:
            switch "hello" {
            case "":
                var v13 uint16 = 8
                _ = v13
            case "a":
                v14 := GenS1{F0: 0}
                _ = v14
            case "hello":
                v8 = v8
            default:
                if v8 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
        case -109:
            if v8 {
                v15 := GenS1{F0: 18446744073709551615}
                fmt.Println(v6, v6)
                v16 := GenS1{F0: 42}
                _ = v15
                _ = v16
            }
        }
    case "a":
        v3[v6] = "hello"
    case "hello":
        v7 = 0
    default:
        switch ("foo bar" + v12) {
        case "\n\t":
            v8 = v8
        case "a":
            var v17 int32 = 42
            _ = v17
        default:
            v8 = (true && ((true && v8) || v8))
        }
    }
    v18 := GenS1{F0: 65535}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
