package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 uint32 = 169
    var v2 uint8 = 255
    var v3 uint = 42
    v4 := []string{(((("a" + ("a" + "a")) + ("hello" + "foo bar")) + "\n\t") + (("\n\t" + "hello") + "\n\t")), "\n\t", (("hello" + "a") + "hello"), "a"}
    var v5 int16 = -32768
    v6 := make(map[uint32]int)
    var v7 int32 = -2147483648
    switch (((("\n\t" + (((("" + (("" + ((("" + "a") + "hello") + ("\n\t" + "\n\t"))) + ("foo bar" + ("" + "a")))) + "hello") + "foo bar") + "")) + "\n\t") + "\n\t") + "\n\t") {
    case "":
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint8 = (v2 >> ((154) & 63))
            v5 += v5
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int = -2147483648
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    case "a":
        var v12 bool = (!(v3 >= uint(v2)))
        _ = v12
    case "\n\t":
        v13 := GenS1{F0: 2, F1: true, F2: 4294967295}
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        switch "hello" {
        case "a":
            var v15 int32 = int32(v5)
            _ = v15
        default:
            v3 = v3
        }
        _ = fi14
    }
    v16 := GenS2{F0: 266, F1: 1, F2: (false && true)}
    v17 := GenS3{F0: -114}
    v18 := GenS4{F0: (v1 ^ (v1 - v1)), F1: 18446744073709551615}
    var v19 int = -9223372036854775808
    v5 -= (v5 + v5)
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
