package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{9223372036854775807, 1, -9223372036854775808}
    var v3 uint64 = 2
    var v4 int64 = -317
    v5 := []string{("a" + ("" + "\n\t")), "foo bar", "foo bar", "\n\t"}
    var v6 float64 = (1e10 + 1.5)
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(v4)
        for fi8 := 0; fi8 < 8; fi8++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v3 = v3
            fmt.Println(v4, fi8)
            _ = fi8
        }
        v6 += v6
        _ = fi7
    }
    v4 = -595
    var v9 uint32 = 2147483648
    for fi10 := 0; fi10 < 7; fi10++ {
        if (false || true) {
            switch "hello" {
            case "\n\t":
                var v11 uint8 = 161
                _ = v11
            case "a":
                var v12 uint32 = (v9 - v9)
                _ = v12
            }
            v1[2147483647] = (fi10 < 2147483647)
            fmt.Println(v9, v9)
        }
        _ = fi10
    }
    var v13 float32 = 0.1
    if ("a" > "hello") {
        switch ((("foo bar" + "a") + "foo bar") + ((("hello" + "\n\t") + "hello") + "a")) {
        case "hello":
            var v14 int = -512
            _ = v14
        case "a":
            for fi15 := 0; fi15 < 2; fi15++ {
                v6 = v6
                _ = fi15
            }
        case "":
            for fi16 := 0; fi16 < 4; fi16++ {
                if len(v2) > 0 {
                    v2[0] = (fi16 / ((fi16) | 1))
                }
                _ = fi16
            }
        }
        _ = len(v1)
    }
    if true {
        v17 := GenS1{F0: 100, F1: 1, F2: false}
        _ = v17
    }
    v18 := GenS1{F0: -100, F1: 646, F2: true}
    fmt.Println(v4, v3, v6)
    var v19 int64 = int64(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
