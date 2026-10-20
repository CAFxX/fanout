package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 int8; F1 uint16; F2 uint16 }

func main() {
    v1 := []string{"hello", "a", ("a" + "\n\t"), (("" + ((("\n\t" + "\n\t") + "") + "hello")) + ("foo bar" + "\n\t"))}
    v2 := []uint32{2147483648, 65535}
    var v3 uint64 = 20
    var v4 int16 = 42
    var v5 uint64 = v3
    v6 := GenS1{F0: 18446744073709551615}
    v4 -= -374
    v5 += v3
    v5 = v5
    _ = len(v1)
    v7 := GenS2{F0: 255, F1: -348, F2: 1}
    switch "" {
    case "\n\t":
        var v8 int = 147
        _ = v8
    case "foo bar":
        v3 = v5
    }
    fmt.Println(v4)
    var v9 string = "a"
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS3{F0: v4}
            _ = fi11
            _ = v12
        }
        v13 := GenS4{F0: 255, F1: 42}
        v4 -= -1
        _ = fi10
        _ = v13
    }
    v14 := GenS5{F0: 0, F1: 42, F2: 255}
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
