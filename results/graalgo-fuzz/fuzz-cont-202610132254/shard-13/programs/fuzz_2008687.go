package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []string{}
    var v2 string = "a"
    var v3 uint64 = 42
    fmt.Println(v3)
    switch v2 {
    case "foo bar":
        v2 = v2
    case "a":
        v4 := GenS1{F0: (v3 << ((65535) & 63))}
        _ = v4
    case "\n\t":
        var v5 uint = 255
        _ = v5
    }
    switch -2147483648 {
    case 127:
        var v6 int64 = -9223372036854775808
        _ = v6
    }
    v3 += v3
    v2 = v2
    v7 := GenS1{F0: (v3 * (^(v3)))}
    v3 += v3
    switch "" {
    case "foo bar":
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                if len(v1) > 0 {
                    v1[0] = ""
                }
                fmt.Println(fi9, v2, fi8)
                _ = fi9
            }
            v10 := GenS1{F0: (v3 ^ v3)}
            _ = fi8
            _ = v10
        }
    case "hello":
        v3 = 18446744073709551615
    case "":
        if (true || true) {
            var v11 int64 = 67
            _ = v11
        }
    default:
        fmt.Println(v2)
    }
    if len(v1) > 0 {
        v1[0] = v2
    }
    v12 := GenS2{F0: 935}
    var v13 uint32 = 640
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
}
