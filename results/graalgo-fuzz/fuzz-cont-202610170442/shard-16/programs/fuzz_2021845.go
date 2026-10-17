package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }

func main() {
    v1 := []string{("foo bar" + "a")}
    var v2 uint = 42
    v3 := make(map[uint32]int)
    var v4 uint8 = 127
    var v5 uint64 = 1
    v4 -= v4
    _ = len(v3)
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(v2, v5)
        _ = fi6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        if false {
            if false {
                v2 -= v2
            }
            v1 = append(v1, "hello")
        }
        v5 += (v5 >> (((^(v2))) & 63))
        _ = fi7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        v4 = v4
        v2 += v2
        _ = fi8
    }
    if false {
        v9 := GenS1{F0: "\n\t", F1: 0, F2: 95}
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(v2)
            fmt.Println(v4)
            _ = fi10
        }
        _ = v9
    }
    _ = len(v3)
    v11 := GenS1{F0: "a", F1: 0, F2: 0}
    v12 := GenS1{F0: (("" + "a") + ("foo bar" + (("a" + ((("a" + (((("" + (("foo bar" + ((("" + "\n\t") + "") + "foo bar")) + "\n\t")) + "hello") + "") + (("a" + ((("a" + "foo bar") + "hello") + ("\n\t" + "\n\t"))) + "hello"))) + (("" + (("" + "foo bar") + "hello")) + "a")) + "hello")) + (((("foo bar" + "") + "a") + "a") + "a")))), F1: -1, F2: 42}
    v4 -= (v4 & v4)
    v5 += 0
    var v13 int8 = -128
    v3[2] = 1
    fmt.Println(v2, v4)
    for fi14 := 0; fi14 < 4; fi14++ {
        switch ("a" + "a") {
        case "foo bar":
            v13 = v13
        default:
            var v15 uint8 = 85
            _ = v15
        }
        var v16 int16 = -459
        _ = fi14
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
