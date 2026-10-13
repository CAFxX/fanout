package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 string; F1 int; F2 uint8 }

func main() {
    var v1 bool = (("foo bar" < ("foo bar" + "foo bar")) || false)
    var v2 uint64 = 2
    v3 := make(map[int]int)
    v2 += (^(v2))
    if (((!v1) && v1) && (2 > 901)) {
        v4 := GenS1{F0: v2, F1: 2}
        for fi5 := 0; fi5 < 18; fi5++ {
            v2 -= v2
            delete(v3, 163)
            var v6 uint64 = v2
            _ = fi5
            _ = v6
        }
        _ = v4
    } else {
        var v7 float64 = ((1.5 / 1.5) + 1.5)
        fmt.Println(v7)
        _ = v7
    }
    fmt.Println(v1, v1)
    var v8 float64 = -2.25
    fmt.Println(v2, v2)
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            if v1 {
                v3[127] = (fi10 ^ -1)
            }
            switch "foo bar" {
            case "":
                fmt.Println(fi9)
            case "a":
                fmt.Println(v1, v8, fi10)
            }
            v1 = v1
            _ = fi10
        }
        var v11 int8 = -1
        v2 -= v2
        _ = fi9
        _ = v11
    }
    var v12 bool = (634 != 0)
    v1 = (true || (v1 || ("foo bar" != ("a" + ("foo bar" + "")))))
    switch -9223372036854775808 {
    case 9223372036854775807:
        if (127 <= 127) {
            switch (("" + "foo bar") + "a") {
            case "hello":
                v13 := GenS1{F0: v2, F1: 2}
                _ = v13
            default:
                v14 := v3[-586]
                _ = v14
            }
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := GenS2{F0: "a", F1: 1}
                _ = fi15
                _ = v16
            }
        }
    default:
        var v17 bool = v12
        _ = v17
    }
    fmt.Println(v1, v12)
    v18 := GenS3{F0: "foo bar", F1: -678, F2: 255}
    var v19 uint64 = 766
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
}
