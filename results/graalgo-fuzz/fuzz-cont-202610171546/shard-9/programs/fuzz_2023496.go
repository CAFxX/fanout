package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 127
    var v3 uint16 = 32767
    v4 := []uint32{4294967295}
    v3 += (v3 / ((2) | 1))
    v5 := GenS1{F0: 0, F1: (("" + ("" + "hello")) + "hello"), F2: 555}
    fmt.Println(v3)
    fmt.Println(v2, v2)
    _ = len(v4)
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 uint64 = 42
            switch "a" {
            case "a":
                var v9 uint64 = (v8 + 1)
                _ = v9
            case "hello":
                var v10 uint16 = (v3 & v3)
                _ = v10
            case "foo bar":
                if len(v4) > 0 {
                    v4[0] = 667
                }
            default:
                v2 -= (v2 << ((4294967295) & 63))
            }
            _ = fi7
            _ = v8
        }
        v11 := GenS1{F0: 418, F1: "hello", F2: fi6}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi6
        _ = v11
    }
    if true {
        v3 = (v3 % ((v3) | 1))
    } else {
        v3 += v3
        fmt.Println(v2, v3)
    }
    var v12 uint8 = 128
    var v13 uint8 = (v12 * v12)
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
