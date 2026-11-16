package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }

func main() {
    v1 := []string{(("hello" + "a") + ""), "a", "\n\t", "hello"}
    var v2 int32 = 1
    v3 := []int{2147483647, 1}
    var v4 int16 = 42
    v5 := []uint32{0, 2147483647, 115, 1}
    var v6 int32 = v2
    var v7 int = 32767
    v4 -= -1
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 int8 = -1
        for fi10 := 0; fi10 < 18; fi10++ {
            v2 -= (v6 | v2)
            v7 -= (^(v7))
            v6 += 2147483647
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    switch "hello" {
    case "":
        var v11 int = (v7 * (v7 & 4294967295))
        _ = v11
    case "foo bar":
        if true {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 int8 = -128
                _ = fi12
                _ = v13
            }
            if ((!(!(-2.25 < 0.1))) || (true && false)) {
                var v14 bool = (9223372036854775807 == 32767)
                fmt.Println(v14, v2, v7)
                _ = v14
            }
        }
    }
    fmt.Println(v7)
    switch (v7 - (v7 % ((v7) | 1))) {
    case -668:
        v4 = v4
    }
    switch "foo bar" {
    case "foo bar":
        v15 := GenS1{F0: ("a" + ("a" + "\n\t")), F1: -206}
        _ = v15
    case "\n\t":
        fmt.Println(v7)
    case "":
        v6 = v2
    }
    switch (v7 * (^(v7))) {
    case 83:
        fmt.Println(v6)
    case 2147483647:
        var v16 int64 = 9223372036854775807
        _ = v16
    }
    v2 += v2
    for fi17 := 0; fi17 < 17; fi17++ {
        v7 = fi17
        _ = fi17
    }
    fmt.Println(v6)
    var v18 uint64 = 966
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
