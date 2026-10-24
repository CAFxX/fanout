package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint16 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int16; F1 string }
func genFn0(p0 int32, p1 string, p2 bool) int8 {
    v5 := GenS1{F0: p1, F1: 0, F2: 840}
    var v6 string = (((p1 + (p1 + ("a" + "\n\t"))) + "foo bar") + ("\n\t" + ""))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 127
}

func main() {
    var v1 uint16 = 2
    var v2 int64 = 1
    var v3 int8 = 100
    v4 := make(map[uint32]int)
    v3 += (^(v3))
    v3 -= v3
    if false {
        if ((18446744073709551615 >= 900) || true) {
            var v7 int64 = 350
            v3 = v3
            switch "\n\t" {
            case "a":
                var v8 uint = 42
                _ = v8
            case "hello":
                fmt.Println(v7, v1)
            case "foo bar":
                v1 = (v1 & 255)
            default:
                switch 0 {
                case -915:
                    v7 = v2
                }
            }
            _ = v7
        } else {
            v9 := GenS2{F0: -9223372036854775808, F1: 4294967295}
            v10 := []int{int(v1), 127}
            _ = v10
            _ = v9
        }
    } else {
        var v11 uint32 = 65535
        for fi12 := 0; fi12 < 9; fi12++ {
            switch (fi12 ^ fi12) {
            case 32768:
                if false {
                    fmt.Println(v3, v11)
                    var v13 float32 = (0.1 / 1.5)
                    if true {
                        _ = len(v4)
                    }
                    _ = v13
                } else {
                    var v14 float64 = (((0.1 - 1e10) - (0.0 - 0.1)) * 3.14159)
                    _ = v14
                }
            default:
                v15 := []int{-9223372036854775808, int(v11), fi12, 42}
                _ = v15
            }
            switch "\n\t" {
            case "a":
                for fi16 := 0; fi16 < 13; fi16++ {
                    var v17 float32 = (((((1e10 + (((-2.25 / 1.5) + 0.1) / (1e10 * 1.5))) - float32(((float64(1.5) - float64(0.0)) - 0.1))) * float32(((3.14159 - 3.14159) - 0.0))) * 1e10) + (((1e10 + 0.1) / 0.1) + float32(1e10)))
                    _ = fi16
                    _ = v17
                }
            case "hello":
                var v18 uint8 = 211
                _ = v18
            }
            switch "a" {
            case "a":
                var v19 uint32 = v11
                _ = v19
            case "\n\t":
                fmt.Println(fi12)
            default:
                if false {
                    fmt.Println(v1, v11)
                    v20 := GenS3{F0: 0, F1: ("" + ("a" + (("foo bar" + "\n\t") + "a")))}
                    for fi21 := 0; fi21 < 18; fi21++ {
                        fmt.Println(v2, fi12, fi21)
                        _ = fi21
                    }
                    _ = v20
                } else {
                    var v22 int8 = v3
                    var v23 int = (fi12 | fi12)
                    for fi24 := 0; fi24 < 7; fi24++ {
                        var v25 uint16 = 818
                        var v26 uint64 = 1
                        var v27 int8 = (v22 ^ (v3 % (((^(v22))) | 1)))
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = v22
                    _ = v23
                }
            }
            _ = fi12
        }
        _ = v11
    }
    var v28 uint8 = 2
    for fi29 := 0; fi29 < 18; fi29++ {
        for fi30 := 0; fi30 < 1; fi30++ {
            delete(v4, 703)
            _ = fi30
        }
        fmt.Println(fi29, fi29, v2)
        _ = fi29
    }
    var v31 uint8 = 230
    switch "a" {
    case "":
        var v32 int32 = 1
        _ = v32
    case "\n\t":
        var v33 bool = (!true)
        _ = v33
    default:
        fmt.Println(v31, v3, v3)
    }
    v34 := []int{0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v34
    _ = v4
}
