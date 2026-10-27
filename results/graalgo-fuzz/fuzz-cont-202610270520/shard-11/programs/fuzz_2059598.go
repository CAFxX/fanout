package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int8 }
func genFn0(p0 uint32) string {
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return (("hello" + "hello") + "")
}

func main() {
    v1 := []uint32{1, 65535, 2147483647}
    var v2 uint8 = 2
    var v3 string = (((("foo bar" + "a") + (("foo bar" + ("foo bar" + "hello")) + "a")) + ("foo bar" + "\n\t")) + "\n\t")
    var v4 bool = false
    v1 = append(v1, 2147483647)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := GenS1{F0: 1, F1: v2, F2: -1}
        var v7 int64 = 212
        switch v3 {
        case "\n\t":
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 uint32 = 2147483648
                _ = fi8
                _ = v9
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = v2
    var v10 int8 = 0
    for fi11 := 0; fi11 < 4; fi11++ {
        switch (v3 + v3) {
        case "":
            var v12 uint32 = 421
            _ = v12
        case "hello":
            v4 = (-817 > int16(v10))
        case "\n\t":
            var v13 int64 = 42
            _ = v13
        }
        fmt.Println(fi11, v10, v10)
        v10 += (v10 / ((v10) | 1))
        _ = fi11
    }
    switch "a" {
    case "":
        v2 = 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
