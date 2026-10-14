package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []uint32{65535}
    v2 := make(map[string]bool)
    var v3 uint8 = 2
    var v4 uint = 1
    v5 := []uint32{}
    v6 := make(map[string]bool)
    switch ((("hello" + "hello") + ("a" + "hello")) + ("" + "\n\t")) {
    case "foo bar":
        switch "" {
        case "foo bar":
            switch "" {
            case "":
                var v7 string = ""
                _ = v7
            case "a":
                v1 = append(v1, 937)
            default:
                for fi8 := 0; fi8 < 15; fi8++ {
                    _ = len(v2)
                    switch fi8 {
                    case -1:
                        var v9 int8 = 42
                        _ = v9
                    case 1:
                        var v10 int32 = int32(v4)
                        _ = v10
                    }
                    _ = fi8
                }
            }
        case "hello":
            v3 = v3
        }
    case "a":
        var v11 int16 = 1
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 15; fi12++ {
            v4 -= 18446744073709551615
            _ = fi12
        }
    }
    var v13 uint8 = (v3 << ((v4) & 63))
    var v14 int8 = -128
    v14 -= (v14 + (v14 | (v14 % ((102) | 1))))
    v13 = (v13 ^ v13)
    var v15 int16 = 421
    if true {
        v13 -= (^(v13))
        delete(v2, "foo bar")
        var v16 uint8 = v13
        _ = v16
    }
    v14 -= int8(v4)
    var v17 uint = (v4 % (((v4 + (v4 + (^(v4))))) | 1))
    v18 := GenS1{F0: true, F1: 18446744073709551615}
    var v19 uint32 = 0
    v13 = (v13 >> ((v4) & 63))
    var v20 float32 = float32(0.0)
    var v21 int8 = (v14 << ((v17) & 63))
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v22 := GenS2{F0: true}
        _ = v22
    default:
        switch -9223372036854775808 {
        case -9223372036854775808:
            var v23 string = ("\n\t" + ("\n\t" + (("foo bar" + ("" + ("foo bar" + "\n\t"))) + "a")))
            _ = v23
        case -83:
            fmt.Println(v14, v14)
        }
    }
    switch 128 {
    case 42:
        var v24 int32 = -2147483648
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
