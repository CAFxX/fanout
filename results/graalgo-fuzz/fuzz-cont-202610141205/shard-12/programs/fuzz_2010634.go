package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS2 struct { F0 int; F1 int32; F2 uint64 }

func main() {
    var v1 float32 = ((1.5 - 1.5) * 0.1)
    v2 := make(map[string]string)
    var v3 float32 = v1
    var v4 uint64 = 42
    var v5 int = 9223372036854775807
    v6 := make(map[int]int)
    v4 += (v4 % (((v4 * (v4 << ((255) & 63)))) | 1))
    if true {
        fmt.Println(v1, v4)
        v7 := GenS1{F0: (v4 % ((v4) | 1)), F1: (^(v4)), F2: -1}
        _ = v7
    }
    if true {
        var v8 uint8 = uint8(v4)
        var v9 string = "\n\t"
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        v1 += 0.1
        v11 := GenS2{F0: 127, F1: -2147483648, F2: v4}
        v12 := GenS2{F0: (fi10 - fi10), F1: 557, F2: v4}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 string = ("foo bar" + ("\n\t" + (("\n\t" + (((((("" + "") + "a") + (("" + "\n\t") + ("" + "foo bar"))) + "hello") + "hello") + "hello")) + ("foo bar" + ("foo bar" + "")))))
    if false {
        switch (^(v5)) {
        case 1:
            _ = len(v2)
        case 32767:
            v1 -= v3
        case -9223372036854775808:
            switch v5 {
            case -946:
                _ = len(v6)
            case -2147483648:
                v13 = v13
            default:
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    if true {
                        var v14 string = (("hello" + ("a" + "hello")) + (v13 + v13))
                        _ = v14
                    }
                default:
                    fmt.Println(v1)
                }
            }
        default:
            v15 := []int{}
            _ = v15
        }
        var v16 int8 = -128
        var v17 uint16 = 2
        _ = v16
        _ = v17
    } else {
        fmt.Println(v5)
        v4 = (^(v4))
    }
    var v18 int64 = 23
    delete(v6, v5)
    v1 -= 0.1
    var v19 uint64 = 0
    for fi20 := 0; fi20 < 7; fi20++ {
        if false {
            var v21 string = v13
            v22 := []int{v5, fi20, (fi20 / (((fi20 ^ fi20)) | 1)), v5}
            var v23 uint32 = 0
            _ = v21
            _ = v22
            _ = v23
        } else {
            fmt.Println(v3, v13)
            v24 := []int{v5, fi20}
            _ = v24
        }
        v25 := []int{4294967295, v5, v5}
        v26 := GenS2{F0: (fi20 >> ((18446744073709551615) & 63)), F1: 1, F2: (v4 / (((^(v19))) | 1))}
        _ = fi20
        _ = v25
        _ = v26
    }
    _ = len(v2)
    v27 := []int{}
    var v28 int = 591
    v2[(v13 + v13)] = (v13 + (v13 + ((("" + "hello") + ((v13 + v13) + v13)) + v13)))
    fmt.Println(v1, v3, v1)
    if true {
        _ = len(v27)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
