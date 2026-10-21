package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
func genFn0(p0 uint64, p1 int32, p2 int) int64 {
    var v5 uint64 = p0
    v6 := GenS1{F0: 464, F1: (p2 % ((p2) | 1))}
    fmt.Println(v5, v5)
    for fi7 := 0; fi7 < 15; fi7++ {
        switch "" {
        case "\n\t":
            switch p2 {
            case 9223372036854775807:
                switch "\n\t" {
                case "foo bar":
                    v5 = 2
                default:
                    v5 = v5
                }
            }
        case "a":
            p2 -= (fi7 & (fi7 - -463))
        case "":
            if (2 == 1) {
                v5 -= v5
            } else {
                p2 = (fi7 + (-(p2)))
            }
        }
        _ = fi7
    }
    var v8 uint = 256
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v8
    return -545
}
func genFn1(p0 uint64) uint64 {
    p0 += (p0 % ((716) | 1))
    var v9 uint16 = 1
    _ = p0
    _ = v9
    return (p0 * uint64(v9))
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]bool)
    var v3 float32 = float32(-2.25)
    var v4 uint64 = 4294967295
    fmt.Println(v1, v1, v3)
    switch -9223372036854775808 {
    case 2147483647:
        if true {
            var v10 uint16 = (v1 % (((v1 % (((v1 / (((v1 >> ((1) & 63))) | 1))) | 1))) | 1))
            fmt.Println(v3, v4, v10)
            _ = v10
        }
    default:
        if (408 != 1000) {
            v4 = (v4 << ((65535) & 63))
        }
    }
    var v11 string = ""
    var v12 int64 = 1
    v13 := GenS1{F0: v12, F1: -9223372036854775808}
    v1 = 711
    var v14 bool = true
    v12 -= genFn0(18446744073709551615, -2147483648, -9223372036854775808)
    switch (("\n\t" + "foo bar") + (v11 + (v11 + (v11 + "a")))) {
    case "a":
        switch (v11 + ("foo bar" + v11)) {
        case "":
            var v15 uint16 = v1
            _ = v15
        default:
            fmt.Println(v12, v3, v14)
        }
    case "":
        var v16 string = (v11 + ("\n\t" + v11))
        _ = v16
    case "\n\t":
        v11 = ("a" + v11)
    }
    if true {
        v11 = (v11 + "")
    } else {
        v11 = (v11 + "a")
        fmt.Println(v1, v14, v3)
        var v17 uint8 = 128
        _ = v17
    }
    v18 := GenS1{F0: v12, F1: -1}
    var v19 uint16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
