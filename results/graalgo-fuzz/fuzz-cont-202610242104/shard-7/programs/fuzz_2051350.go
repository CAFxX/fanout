package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS2 struct { F0 int }

func main() {
    v1 := []string{"\n\t", ("foo bar" + "hello")}
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    var v4 uint16 = 251
    delete(v3, "foo bar")
    v4 = (v4 << ((700) & 63))
    var v5 bool = false
    v5 = (true || ((42 < 18446744073709551615) && v5))
    _ = len(v2)
    v4 -= v4
    switch "a" {
    case "":
        v3[(("hello" + "\n\t") + "hello")] = -9223372036854775808
    case "a":
        if v5 {
            switch 32767 {
            case -812:
                if v5 {
                    if (42 >= -2147483648) {
                        fmt.Println(v5)
                    }
                    var v6 uint = 4294967295
                    _ = v6
                } else {
                    if v5 {
                        var v7 uint = 18446744073709551615
                        _ = v7
                    } else {
                        if (!v5) {
                            var v8 uint64 = 0
                            v4 = (v4 / ((v4) | 1))
                            for fi9 := 0; fi9 < 19; fi9++ {
                                v8 -= v8
                                v8 += v8
                                _ = fi9
                            }
                            _ = v8
                        }
                        v4 -= v4
                    }
                }
            case 9223372036854775807:
                for fi10 := 0; fi10 < 3; fi10++ {
                    v11 := GenS1{F0: (^(v4)), F1: 0, F2: 0}
                    var v12 uint = 18446744073709551615
                    v13 := GenS2{F0: (^(fi10))}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            default:
                for fi14 := 0; fi14 < 18; fi14++ {
                    v5 = true
                    _ = fi14
                }
            }
            v5 = false
        }
    case "foo bar":
        for fi15 := 0; fi15 < 13; fi15++ {
            v16 := GenS1{F0: (v4 ^ v4), F1: 521, F2: 42}
            fmt.Println(v5, fi15, fi15)
            _ = fi15
            _ = v16
        }
    default:
        v17 := GenS2{F0: -1}
        _ = v17
    }
    var v18 uint16 = 786
    var v19 uint = 255
    v5 = v5
    var v20 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
