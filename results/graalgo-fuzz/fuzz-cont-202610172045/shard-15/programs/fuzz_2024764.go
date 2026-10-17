package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int = 237
    v2 := make(map[string]bool)
    var v3 int = v1
    if ((v3 <= (v1 ^ (v3 % ((-2147483648) | 1)))) && (65535 != 42)) {
        v3 = 4294967295
        v1 -= (v1 * (v3 % (((v3 + v1)) | 1)))
    }
    if (4294967295 <= 2) {
        var v4 string = "\n\t"
        v3 += 128
        _ = v4
    } else {
        var v5 string = "hello"
        v3 = v3
        v1 = v3
        _ = v5
    }
    var v6 int = (^(v3))
    v7 := GenS1{F0: -12}
    var v8 bool = false
    switch (v3 - (v1 * (v3 / ((127) | 1)))) {
    case 462:
        for fi9 := 0; fi9 < 5; fi9++ {
            v8 = (!true)
            for fi10 := 0; fi10 < 11; fi10++ {
                if false {
                    v3 -= (fi9 / (((v6 >> ((196) & 63))) | 1))
                    v11 := GenS1{F0: 100}
                    _ = v11
                } else {
                    v12 := GenS1{F0: 72}
                    _ = v12
                }
                fmt.Println(v1, fi9)
                _ = fi10
            }
            var v13 int32 = 1
            _ = fi9
            _ = v13
        }
    case 850:
        v14 := GenS2{F0: 255, F1: -2147483648}
        _ = v14
    case 32767:
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := GenS2{F0: 255, F1: -2147483648}
            v17 := GenS3{F0: 65535}
            _ = fi15
            _ = v16
            _ = v17
        }
    default:
        switch 32768 {
        case 904:
            switch "hello" {
            case "a":
                if v8 {
                    v1 += 127
                } else {
                    v2[("hello" + (("a" + "foo bar") + "foo bar"))] = v8
                    fmt.Println(v1, v3)
                    switch "" {
                    case "a":
                        var v18 int8 = 82
                        _ = v18
                    default:
                        var v19 float64 = 1e10
                        _ = v19
                    }
                }
            case "\n\t":
                v1 -= v1
            default:
                v20 := v2[("foo bar" + "")]
                _ = v20
            }
        case -13:
            var v21 int = 0
            _ = v21
        case 2147483647:
            v1 = v1
        }
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 bool = (-231 == 0)
        _ = fi22
        _ = v23
    }
    v6 += v1
    v24 := GenS2{F0: 2, F1: -784}
    v25 := GenS3{F0: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
