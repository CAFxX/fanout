package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint32 }
type GenS2 struct { F0 string; F1 uint64; F2 int64 }
func genFn0(p0 uint32, p1 string) int64 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 11; fi4++ {
        switch "a" {
        case "a":
            p0 = (p0 + (p0 * 42))
        case "hello":
            v5 := GenS1{F0: ((((("foo bar" + ("" + p1)) + p1) + (p1 + p1)) + p1) + ("a" + p1)), F1: 92, F2: 282}
            _ = v5
        }
        _ = fi4
    }
    _ = p0
    _ = p1
    return -424
}

func main() {
    v1 := []int{-9223372036854775808, 2147483647}
    v2 := []int{-2147483648}
    var v3 string = ("foo bar" + "foo bar")
    var v6 uint8 = 209
    v1 = append(v1, 32767)
    fmt.Println(v6)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = (v3 + "hello")
    var v7 uint64 = 4294967295
    var v8 int64 = genFn0(832, v3)
    for fi9 := 0; fi9 < 16; fi9++ {
        switch (("\n\t" + (("" + v3) + ("a" + "foo bar"))) + "foo bar") {
        case "\n\t":
            var v10 int16 = 1
            _ = v10
        case "":
            v2 = append(v2, (-(fi9)))
        default:
            switch "\n\t" {
            case "":
                v6 = v6
            case "a":
                v11 := GenS2{F0: "", F1: (v7 | (v7 & v7)), F2: v8}
                _ = v11
            }
        }
        _ = fi9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
