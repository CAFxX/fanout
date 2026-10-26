package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint16 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 42
    var v3 uint64 = 2
    v4 := []string{"\n\t", "foo bar", "foo bar"}
    var v5 uint8 = 0
    v2 += (^(v2))
    var v6 uint16 = 65535
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    v2 = v2
    v7 := GenS1{F0: -2147483648, F1: -2147483648, F2: v6}
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(v2, v5)
        delete(v1, fi8)
        switch fi8 {
        case 32767:
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 bool = (true && true)
                v11 := GenS2{F0: ((("hello" + "foo bar") + (("foo bar" + "hello") + ("" + "a"))) + "a")}
                v12 := GenS1{F0: -2147483648, F1: 0, F2: v6}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        case -814:
            v5 = 127
        }
        _ = fi8
    }
    v5 = 142
    var v13 uint16 = (v6 - 428)
    v6 += (v13 ^ (v6 & (v6 + (v6 >> ((2) & 63)))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
