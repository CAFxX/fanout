package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 551
    var v3 int32 = 1
    var v4 uint32 = 4294967295
    v5 := []string{(("\n\t" + "") + ""), (("a" + ("foo bar" + "")) + "a")}
    for fi6 := 0; fi6 < 11; fi6++ {
        v5 = append(v5, "foo bar")
        _ = fi6
    }
    if true {
        v1 = v1
        if (1e10 < 0.0) {
            _ = len(v5)
            var v7 uint16 = v2
            _ = v7
        } else {
            var v8 int16 = 42
            _ = v8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS1{F0: 42}
            _ = fi9
            _ = v10
        }
    }
    switch ((("hello" + "hello") + "\n\t") + "foo bar") {
    case "a":
        var v11 float32 = 1.5
        _ = v11
    }
    v2 -= (^(v2))
    var v12 uint8 = 42
    if (!false) {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        v4 -= (v4 ^ 397)
        if len(v5) > 0 {
            v5[0] = ("hello" + "\n\t")
        }
    } else {
        v13 := GenS2{F0: 100, F1: 630}
        _ = v13
    }
    var v14 int16 = 42
    v12 = (^(v12))
    v14 -= int16(v12)
    fmt.Println(v1, v12, v1)
    if (true && ((^(v12)) > (v12 + v12))) {
        v5 = append(v5, (("foo bar" + "hello") + "foo bar"))
        if len(v5) > 0 {
            v5[0] = "a"
        }
        v15 := GenS2{F0: 127, F1: 15}
        _ = v15
    }
    fmt.Println(v1, v4, v12)
    v4 = v4
    var v16 int16 = (v14 + (v14 & (v14 % ((v14) | 1))))
    if (!false) {
        fmt.Println(v4, v12, v4)
        v5 = append(v5, "a")
    }
    var v17 int = -1
    fmt.Println(v1)
    switch ("" + ("a" + "a")) {
    case "\n\t":
        v12 += (v12 << ((508) & 63))
    case "foo bar":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        v17 = -131
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
