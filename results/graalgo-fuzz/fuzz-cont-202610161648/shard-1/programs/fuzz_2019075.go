package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int16 = -32768
    var v2 float64 = 1e10
    var v3 int64 = -9223372036854775808
    fmt.Println(v2, v2)
    var v4 string = (("foo bar" + (("" + "foo bar") + ("foo bar" + "\n\t"))) + ("a" + ("a" + ("hello" + "a"))))
    switch v4 {
    case "hello":
        switch 1 {
        case -9223372036854775808:
            v1 += 1000
        case -902:
            var v5 uint32 = 169
            _ = v5
        case 42:
            var v6 bool = (v3 <= -9223372036854775808)
            _ = v6
        }
    case "\n\t":
        switch -192 {
        case -9223372036854775808:
            v7 := GenS1{F0: 198, F1: 42}
            _ = v7
        }
    case "a":
        for fi8 := 0; fi8 < 12; fi8++ {
            v4 = v4
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := []int{(fi8 << ((463) & 63))}
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    default:
        var v11 uint16 = 2
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(v3, fi12, v3)
        _ = fi12
    }
    v13 := []int{-588}
    v4 = ""
    if true {
        v4 = ("a" + "hello")
    } else {
        var v14 int16 = (v1 >> ((126) & 63))
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 float64 = v2
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v2 = v2
    if true {
        v17 := GenS1{F0: 0, F1: 0}
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
