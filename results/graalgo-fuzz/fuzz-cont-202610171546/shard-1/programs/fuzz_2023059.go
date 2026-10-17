package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint64; F1 int64; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 18446744073709551615
    var v3 int8 = -128
    if (true || true) {
        delete(v1, 4294967295)
        var v4 float32 = ((0.0 - (0.1 - 1e10)) / 1e10)
        _ = v4
    } else {
        switch "hello" {
        case "":
            v5 := GenS1{F0: 2, F1: 1000, F2: -2147483648}
            _ = v5
        default:
            v3 += (v3 | (-(v3)))
        }
    }
    v6 := GenS2{F0: 94, F1: -1}
    var v7 uint = 65535
    var v8 bool = false
    v3 = v3
    if ((v3 != 1) && (((v2 | 1) < (v2 & 4294967295)) || (v8 || v8))) {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := v1[2147483647]
            switch "hello" {
            case "\n\t":
                switch "a" {
                case "foo bar":
                    var v11 int = fi9
                    _ = v11
                }
            case "a":
                var v12 int16 = 42
                _ = v12
            default:
                if (530 == (v2 ^ 503)) {
                    if ((!((v2 + v2) <= 1)) && ((3.14159 + 3.14159) <= 1.5)) {
                        v8 = ((false && (v2 != (^(v2)))) && v8)
                        v13 := []int{(fi9 ^ fi9)}
                        _ = v13
                    }
                    var v14 string = ("hello" + (("" + "a") + "a"))
                    switch "\n\t" {
                    case "":
                        v10 = (!v10)
                    }
                    _ = v14
                }
            }
            switch "" {
            case "a":
                v15 := []int{-1}
                _ = v15
            case "\n\t":
                for fi16 := 0; fi16 < 6; fi16++ {
                    v2 -= 2
                    var v17 int8 = v3
                    _ = fi16
                    _ = v17
                }
            default:
                var v18 uint8 = 0
                _ = v18
            }
            _ = fi9
            _ = v10
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := v1[65535]
            for fi21 := 0; fi21 < 6; fi21++ {
                switch -966 {
                case 0:
                    v2 -= 987
                case 148:
                    v22 := []int{2147483647, (-(fi19)), -2147483648}
                    _ = v22
                }
                v8 = ((32767 > 669) || (false || (672 > 2)))
                var v23 int = fi19
                _ = fi21
                _ = v23
            }
            _ = fi19
            _ = v20
        }
        var v24 uint32 = 877
        _ = v24
    }
    switch (("hello" + "hello") + "a") {
    case "":
        fmt.Println(v2, v3)
    case "\n\t":
        if true {
            for fi25 := 0; fi25 < 18; fi25++ {
                v26 := GenS1{F0: 255, F1: 42, F2: fi25}
                if ((false || (!v8)) && v8) {
                    v2 -= uint64(v7)
                    v2 += v2
                    switch ("a" + (("" + "foo bar") + "a")) {
                    case "hello":
                        for fi27 := 0; fi27 < 15; fi27++ {
                            fmt.Println(fi25)
                            _ = fi27
                        }
                    }
                }
                switch ("foo bar" + (("\n\t" + "hello") + ("hello" + "\n\t"))) {
                case "hello":
                    var v28 float64 = ((float64((float32(0.1) / (1.5 - 0.0))) / 0.1) - 3.14159)
                    _ = v28
                case "\n\t":
                    v29 := []int{(fi25 % ((-2147483648) | 1)), -2147483648, int(v2)}
                    _ = v29
                default:
                    for fi30 := 0; fi30 < 1; fi30++ {
                        fmt.Println(fi30)
                        _ = fi30
                    }
                }
                _ = fi25
                _ = v26
            }
        }
    case "foo bar":
        v3 += (v3 / (((^(v3))) | 1))
    default:
        var v31 int16 = -1
        _ = v31
    }
    switch "\n\t" {
    case "a":
        v32 := GenS3{F0: (v2 - v2), F1: 362, F2: 153}
        _ = v32
    default:
        fmt.Println(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
