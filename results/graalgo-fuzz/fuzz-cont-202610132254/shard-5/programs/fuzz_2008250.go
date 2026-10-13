package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int = 9223372036854775807
    v2 := []uint32{2147483648, 65535}
    var v3 uint64 = 258
    v3 -= (^(v3))
    v1 = (^(v1))
    v4 := GenS1{F0: 42}
    v3 -= (v3 | (v3 >> ((65535) & 63)))
    switch "\n\t" {
    case "":
        var v5 int = v1
        _ = v5
    case "\n\t":
        var v6 uint64 = v3
        _ = v6
    case "hello":
        var v7 int32 = -2147483648
        _ = v7
    default:
        _ = len(v2)
    }
    _ = len(v2)
    switch v1 {
    case 9223372036854775807:
        switch 128 {
        case 127:
            switch "foo bar" {
            case "hello":
                var v8 int8 = -96
                _ = v8
            case "":
                if false {
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                    for fi9 := 0; fi9 < 12; fi9++ {
                        var v10 int16 = 0
                        _ = fi9
                        _ = v10
                    }
                    if false {
                        var v11 int32 = 2147483647
                        fmt.Println(v1, v11, v1)
                        _ = v11
                    } else {
                        for fi12 := 0; fi12 < 9; fi12++ {
                            var v13 int32 = -2147483648
                            _ = fi12
                            _ = v13
                        }
                    }
                } else {
                    var v14 int64 = 9223372036854775807
                    _ = v14
                }
            }
        }
    }
    fmt.Println(v3, v1)
    var v15 int = (v1 * v1)
    v16 := GenS2{F0: 213}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
