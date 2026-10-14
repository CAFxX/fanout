package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int }
type GenS2 struct { F0 bool; F1 int16 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 861
    var v3 int = 1
    var v4 uint = 0
    v5 := []string{(("foo bar" + "\n\t") + "hello"), "\n\t", "a"}
    var v6 uint8 = 42
    switch ((("foo bar" + ("" + "a")) + "") + (((("hello" + "\n\t") + ("foo bar" + "")) + ("" + "a")) + "a")) {
    case "hello":
        var v7 int16 = -702
        _ = v7
    default:
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    }
    fmt.Println(v4)
    var v8 uint32 = (v2 >> ((v4) & 63))
    for fi9 := 0; fi9 < 16; fi9++ {
        delete(v1, (v3 - fi9))
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 uint64 = 18446744073709551615
            v3 += v3
            _ = fi10
            _ = v11
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi9
    }
    delete(v1, 9223372036854775807)
    v12 := GenS1{F0: "hello", F1: 9223372036854775807, F2: v3}
    var v13 uint16 = 953
    v3 += v3
    var v14 string = ("" + "hello")
    fmt.Println(v6)
    fmt.Println(v2, v2)
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 bool = false
        v17 := GenS2{F0: v16, F1: 1000}
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS2{F0: v16, F1: 1000}
            v14 = (("hello" + "") + (v14 + v14))
            var v20 int64 = 42
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    if true {
        v8 = v2
    }
    var v21 uint16 = v13
    var v22 int = (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
