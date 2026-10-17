package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 236
    var v3 string = "hello"
    var v4 uint16 = 2
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]int)
    if false {
        var v7 int64 = 1
        v2 += (v2 | (^(v2)))
        v8 := GenS1{F0: 9223372036854775807, F1: (v4 * v4), F2: 42}
        _ = v7
        _ = v8
    }
    v3 = "a"
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(v3, fi9, v4)
        _ = fi9
    }
    v4 += v4
    switch v3 {
    case "foo bar":
        v3 = v3
    case "hello":
        var v10 int64 = 0
        _ = v10
    case "":
        var v11 uint64 = 127
        _ = v11
    default:
        v12 := GenS1{F0: 1, F1: 65535, F2: 77}
        _ = v12
    }
    v3 = (v3 + "foo bar")
    switch 0 {
    case -260:
        v2 += v2
    default:
        delete(v1, -9223372036854775808)
    }
    if ((v3 + "hello") > v3) {
        for fi13 := 0; fi13 < 1; fi13++ {
            if (fi13 >= fi13) {
                v4 -= 849
                fmt.Println(v2)
            } else {
                switch "hello" {
                case "":
                    v14 := v5[65535]
                    _ = v14
                case "foo bar":
                    for fi15 := 0; fi15 < 16; fi15++ {
                        fmt.Println(fi15, v3, v4)
                        var v16 float32 = 1.5
                        _ = fi15
                        _ = v16
                    }
                default:
                    var v17 int = fi13
                    _ = v17
                }
                var v18 float32 = -2.25
                _ = v18
            }
            _ = fi13
        }
    }
    switch (v3 + v3) {
    case "a":
        v4 = v4
    case "\n\t":
        fmt.Println(v2)
    default:
        v4 += (v4 ^ (v4 >> ((1) & 63)))
    }
    v3 = v3
    v4 = v4
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 uint64 = 0
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        fmt.Println(v2)
        for fi22 := 0; fi22 < 8; fi22++ {
            v2 = v2
            var v23 int8 = 42
            v4 += (v4 & 1)
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    var v24 string = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
