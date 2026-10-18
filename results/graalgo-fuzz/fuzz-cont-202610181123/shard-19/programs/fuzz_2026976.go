package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 string = ("" + ("\n\t" + (("" + (((("a" + "") + "") + "a") + "a")) + "hello")))
    v2 := make(map[string]uint32)
    var v3 uint = 224
    v4 := GenS1{F0: 0}
    fmt.Println(v1, v1)
    v5 := GenS1{F0: 2147483648}
    if (100 <= 0) {
        v3 = (v3 & (v3 * (v3 + v3)))
        v6 := GenS2{F0: (v1 + "a"), F1: -515}
        _ = v6
    }
    fmt.Println(v1, v1)
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(fi7, v3)
            var v9 int64 = 413
            v9 -= 1
            v10 := []int{-2147483648, -1}
            var v11 uint8 = 0
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            switch v1 {
            case "a":
                delete(v2, v1)
            case "\n\t":
                v13 := GenS3{F0: v3}
                var v14 int16 = -429
                if true {
                    var v15 uint = (v3 - 18446744073709551615)
                    _ = v15
                }
                _ = v13
                _ = v14
            default:
                fmt.Println(v3)
            }
            switch fi7 {
            case 4294967295:
                var v16 string = ((v1 + v1) + "hello")
                var v17 string = "hello"
                _ = v16
                _ = v17
            default:
                v18 := []int{(fi7 % ((fi7) | 1)), (^(fi7)), (fi12 >> ((v3) & 63))}
                _ = v18
            }
            _ = fi12
        }
        _ = fi7
    }
    if (2 >= 42) {
        v1 = (v1 + v1)
        v1 = v1
    } else {
        v19 := GenS2{F0: v1, F1: 32767}
        var v20 string = v1
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint16 = 42
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
