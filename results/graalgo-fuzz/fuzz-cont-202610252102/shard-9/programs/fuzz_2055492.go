package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 int; F1 string; F2 int8 }

func main() {
    var v1 int16 = 381
    v2 := make(map[uint32]uint32)
    v3 := []string{((("a" + "") + ("foo bar" + "a")) + ("foo bar" + "hello"))}
    if (42 < 0) {
        v1 += v1
    }
    v1 = 233
    fmt.Println(v1)
    v4 := GenS1{F0: false, F1: v1, F2: 2147483647}
    if true {
        switch ("foo bar" + "hello") {
        case "\n\t":
            var v5 int = -951
            _ = v5
        case "":
            fmt.Println(v1)
        }
        var v6 int32 = 1
        _ = v6
    }
    fmt.Println(v1)
    switch 195 {
    case 42:
        var v7 float64 = 1.5
        _ = v7
    default:
        fmt.Println(v1)
    }
    v1 -= 32767
    var v8 uint = 735
    v9 := GenS1{F0: (3.14159 > (0.1 / 3.14159)), F1: -1, F2: -991}
    for fi10 := 0; fi10 < 13; fi10++ {
        if true {
            for fi11 := 0; fi11 < 10; fi11++ {
                switch fi10 {
                case -262:
                    v8 = 632
                }
                _ = fi11
            }
            if len(v3) > 0 {
                v3[0] = ("foo bar" + "\n\t")
            }
        } else {
            _ = len(v2)
            for fi12 := 0; fi12 < 10; fi12++ {
                if false {
                    if true {
                        v8 = (v8 - 65535)
                    }
                    var v13 int = fi10
                    fmt.Println(fi12, v13, fi12)
                    var v14 uint8 = 24
                    v15 := GenS2{F0: -952, F1: "", F2: 1}
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    var v16 int32 = 42
                    _ = v16
                }
                var v17 uint8 = 44
                _ = fi12
                _ = v17
            }
        }
        _ = fi10
    }
    v1 -= 1000
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
