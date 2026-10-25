package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint64; F2 string }
func genFn0(p0 uint16, p1 int16) int {
    switch 32767 {
    case 529:
        p1 = (p1 + (p1 & (p1 & (^(p1)))))
    }
    var v7 int64 = -1
    v8 := GenS1{F0: 42}
    p0 -= 2
    if ((int8(p1) > 127) || (255 < 128)) {
        var v9 int = 127
        v7 += (v7 ^ v7)
        var v10 float64 = 0.1
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 uint8 = 1
    v2 := []string{"", "\n\t", ("hello" + ("hello" + "hello"))}
    var v3 int16 = -32768
    var v4 int = 42
    v5 := make(map[int]uint32)
    v6 := []string{"hello", ("foo bar" + "hello"), "hello", ""}
    if false {
        v1 -= (v1 / (((v1 % ((v1) | 1))) | 1))
        var v11 int64 = 10
        _ = v11
    }
    var v12 string = "hello"
    v5[v4] = 22
    v12 = v12
    var v13 uint = 255
    v3 = v3
    if (!(!false)) {
        var v14 uint64 = 551
        _ = v14
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint32 = 65535
            var v17 uint8 = v1
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int64 = 1
        v19 := GenS2{F0: v12, F1: 18446744073709551615, F2: ((v12 + v12) + "hello")}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        _ = len(v5)
        v4 = v4
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
