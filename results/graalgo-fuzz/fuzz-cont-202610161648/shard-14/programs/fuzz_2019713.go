package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -1
    v3 := []string{"\n\t", ("a" + ""), "a", (("a" + ("a" + ("" + "\n\t"))) + "\n\t")}
    v4 := make(map[uint32]string)
    for fi5 := 0; fi5 < 20; fi5++ {
        v2 -= v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi5
    }
    v6 := GenS1{F0: -1, F1: (2147483648 <= 856), F2: v2}
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint64 = 680
        fmt.Println(v8, v2, v8)
        delete(v4, 889)
        _ = fi7
        _ = v8
    }
    if len(v3) > 0 {
        v3[0] = ("foo bar" + "hello")
    }
    _ = len(v3)
    if (!true) {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS2{F0: 65535, F1: false}
            switch 29 {
            case -1:
                var v11 uint = 4294967295
                _ = v11
            }
            _ = fi9
            _ = v10
        }
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS1{F0: -688, F1: false, F2: v2}
        v14 := GenS2{F0: 0, F1: (true || (-121 > 100))}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint8 = 255
    v15 = 0
    var v16 bool = (!true)
    v17 := GenS2{F0: 468, F1: (!true)}
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint64 = 907
        _ = fi18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
