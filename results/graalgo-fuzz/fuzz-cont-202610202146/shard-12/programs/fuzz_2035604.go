package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    v3 := []uint32{2147483647}
    if ((true && false) && true) {
        v4 := GenS1{F0: 255, F1: 32767}
        _ = v4
    }
    switch "hello" {
    case "\n\t":
        var v5 int64 = 42
        _ = v5
    }
    var v6 int64 = -1
    var v7 uint16 = 1
    var v8 uint8 = 42
    _ = len(v1)
    v8 += 2
    v9 := GenS1{F0: v8, F1: v7}
    if false {
        switch ("hello" + ("foo bar" + "a")) {
        case "":
            v8 -= (v8 % (((^(v8))) | 1))
        case "\n\t":
            switch 128 {
            case 4294967295:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 887:
                v2 = (^(v2))
            default:
                fmt.Println(v8)
            }
        }
        var v10 int16 = 1000
        _ = v10
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS1{F0: (v8 * 36), F1: 32767}
            _ = fi11
            _ = v12
        }
        var v13 int = 185
        _ = v13
    }
    fmt.Println(v2)
    if (!((32767 > 42) && false)) {
        for fi14 := 0; fi14 < 11; fi14++ {
            fmt.Println(v8, v8, v8)
            if true {
                var v15 int = int(v8)
                _ = v15
            }
            _ = fi14
        }
        v7 = uint16(v8)
    }
    fmt.Println(v2, v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
