package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint16 }
type GenS2 struct { F0 string; F1 uint16 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{208, 2}
    v3 := []string{(("foo bar" + "") + ("a" + (("\n\t" + "\n\t") + ("\n\t" + "a"))))}
    var v4 int64 = (v1 + (v1 << ((18446744073709551615) & 63)))
    v5 := []uint32{}
    var v6 uint64 = 2
    var v7 int32 = 574
    for fi8 := 0; fi8 < 10; fi8++ {
        v1 += v1
        _ = fi8
    }
    v9 := GenS1{F0: v6, F1: 4294967295, F2: 255}
    fmt.Println(v1, v7, v7)
    _ = len(v2)
    var v10 string = ("hello" + "")
    if (!false) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v10 = ((v10 + (v10 + "hello")) + v10)
    }
    var v11 string = "\n\t"
    v1 = v1
    switch 0 {
    case 32767:
        v11 = (v11 + "a")
    case 127:
        if (!true) {
            v12 := GenS2{F0: v10, F1: 255}
            fmt.Println(v7, v7, v11)
            _ = v12
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
