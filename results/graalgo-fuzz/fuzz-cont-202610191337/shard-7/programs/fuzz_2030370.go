package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint32; F2 uint }

func main() {
    var v1 uint32 = 821
    var v2 uint64 = 0
    v3 := []int{}
    v4 := make(map[uint32]bool)
    var v5 uint32 = (^(v1))
    if (1000 > 309) {
        for fi6 := 0; fi6 < 6; fi6++ {
            switch (fi6 >> ((255) & 63)) {
            case 488:
                v7 := GenS1{F0: 42, F1: -559}
                _ = v7
            case 128:
                v8 := GenS2{F0: 65535}
                _ = v8
            default:
                for fi9 := 0; fi9 < 18; fi9++ {
                    var v10 float64 = (1e10 + ((1e10 + 1e10) * 0.0))
                    v11 := GenS1{F0: 255, F1: -1}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            v1 = (v5 ^ (v1 + (v5 % ((v1) | 1))))
            v12 := GenS3{F0: 58, F1: (v5 % ((v1) | 1)), F2: 18446744073709551615}
            _ = fi6
            _ = v12
        }
        fmt.Println(v5)
        if true {
            fmt.Println(v5, v1)
        }
    }
    _ = len(v4)
    var v13 int32 = 1
    switch (("foo bar" + ((("" + ("\n\t" + "")) + "hello") + "foo bar")) + "\n\t") {
    case "":
        switch -2147483648 {
        case 128:
            v5 -= v1
        default:
            var v14 bool = (true || (!true))
            _ = v14
        }
    case "hello":
        v15 := GenS3{F0: -115, F1: v5, F2: 65535}
        _ = v15
    case "a":
        var v16 uint64 = 876
        _ = v16
    }
    var v17 uint8 = 127
    for fi18 := 0; fi18 < 16; fi18++ {
        if true {
            switch "foo bar" {
            case "":
                _ = len(v3)
            case "\n\t":
                fmt.Println(v2, v17, v5)
            case "a":
                var v19 float32 = 1.5
                _ = v19
            }
        }
        _ = fi18
    }
    v4[v1] = ((("foo bar" + "a") != ("" + "hello")) && (("foo bar" < "a") && (255 < 728)))
    v20 := v4[(v5 * 1)]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
