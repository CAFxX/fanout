package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }

func main() {
    var v1 float32 = 1.5
    var v2 uint64 = 12
    v3 := []string{"\n\t", (("foo bar" + "\n\t") + "a"), "hello", ("foo bar" + "")}
    var v4 float64 = (0.0 + 0.0)
    var v5 float64 = v4
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int64 = 42
        v8 := GenS1{F0: (fi6 | (fi6 * fi6)), F1: 18446744073709551615}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 255
    v10 := GenS1{F0: 4294967295, F1: v2}
    var v11 int = -2147483648
    fmt.Println(v11)
    v12 := GenS2{F0: false}
    var v13 int64 = 1
    _ = len(v3)
    v9 = v9
    v14 := GenS1{F0: 0, F1: (v2 >> ((18446744073709551615) & 63))}
    v15 := GenS2{F0: ("a" <= "hello")}
    for fi16 := 0; fi16 < 6; fi16++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v17 := GenS1{F0: -844, F1: (v2 % ((42) | 1))}
        v2 -= 2
        _ = fi16
        _ = v17
    }
    if true {
        v18 := GenS1{F0: 127, F1: v2}
        _ = v18
    } else {
        v19 := GenS3{F0: (-633 >= 1000)}
        _ = v19
    }
    fmt.Println(v9, v4, v9)
    if (false || false) {
        fmt.Println(v11)
        _ = len(v3)
        for fi20 := 0; fi20 < 9; fi20++ {
            v3 = append(v3, "")
            _ = fi20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
