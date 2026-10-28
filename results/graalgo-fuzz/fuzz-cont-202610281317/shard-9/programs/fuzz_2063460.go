package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int64, p1 int8) int {
    v7 := []int{2147483647, 9223372036854775807, 252, 32768}
    var v8 bool = false
    fmt.Println(p1, p1)
    var v9 float64 = 0.0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 32768
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint = 4294967295
    var v3 uint8 = 0
    var v4 uint16 = 265
    v5 := []string{(("\n\t" + "hello") + "\n\t"), ("hello" + "a"), "foo bar"}
    var v6 float32 = -2.25
    var v10 uint16 = (^(v4))
    switch 4294967295 {
    case -9223372036854775808:
        v4 = v4
    case 2147483647:
        fmt.Println(v4, v1, v10)
    default:
        v3 += (v3 % ((128) | 1))
    }
    var v11 int64 = 0
    var v12 uint = (v2 - (^(v2)))
    var v13 uint = uint(v3)
    var v14 int8 = -1
    if false {
        var v15 bool = ((v14 >> (((^(v2))) & 63)) >= v14)
        fmt.Println(v2)
        _ = v15
    }
    v2 -= v12
    v3 = v3
    fmt.Println(v10, v3)
    v16 := GenS1{F0: 762, F1: 100}
    if ((^(v1)) != v1) {
        v17 := GenS2{F0: ("\n\t" + ("\n\t" + (("a" + "hello") + "a"))), F1: 32767}
        var v18 uint = v13
        _ = v17
        _ = v18
    } else {
        v19 := GenS2{F0: "\n\t", F1: 1}
        var v20 uint64 = 2
        v14 += 42
        _ = v19
        _ = v20
    }
    if len(v5) > 0 {
        v5[0] = ("foo bar" + "foo bar")
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        if false {
            var v22 string = ((((("hello" + ("" + "\n\t")) + ("" + "a")) + "hello") + "") + "\n\t")
            switch (fi21 + fi21) {
            case -2147483648:
                var v23 int64 = v11
                _ = v23
            case 0:
                v24 := GenS2{F0: "\n\t", F1: -1}
                _ = v24
            case 1:
                for fi25 := 0; fi25 < 14; fi25++ {
                    var v26 string = (v22 + "hello")
                    _ = fi25
                    _ = v26
                }
            default:
                v27 := GenS3{F0: -2147483648}
                _ = v27
            }
            _ = v22
        }
        switch (fi21 % (((fi21 + genFn0(v11, -1))) | 1)) {
        case 128:
            v28 := GenS1{F0: -829, F1: v14}
            _ = v28
        default:
            switch (fi21 % (((fi21 >> ((uint(v1)) & 63))) | 1)) {
            case 267:
                switch fi21 {
                case 219:
                    v1 = (v1 & (v1 & (v1 >> ((4294967295) & 63))))
                case -650:
                    if (18446744073709551615 < 0) {
                        if len(v5) > 0 {
                            v5[0] = "a"
                        }
                    }
                case 32767:
                    if len(v5) > 0 {
                        v5[0] = "hello"
                    }
                }
            case 32768:
                v29 := GenS4{F0: v3}
                _ = v29
            case -259:
                v30 := GenS3{F0: 0}
                _ = v30
            }
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
