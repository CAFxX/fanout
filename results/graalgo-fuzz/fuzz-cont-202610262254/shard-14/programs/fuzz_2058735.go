package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint }

func main() {
    var v1 uint16 = 652
    var v2 int32 = 2147483647
    var v3 string = ("foo bar" + "foo bar")
    var v4 string = v3
    var v5 bool = ((18446744073709551615 != 4294967295) || (!true))
    v6 := GenS1{F0: -9223372036854775808, F1: 42}
    if (((3.14159 + 1.5) / 3.14159) != 0.1) {
        v7 := GenS1{F0: -938, F1: 128}
        _ = v7
    }
    if (960 > 0) {
        v5 = ("" == v3)
    } else {
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                v2 -= v2
                fmt.Println(v3, v4, fi9)
                if ((v3 != "foo bar") && (!v5)) {
                    if v5 {
                        switch (fi9 * 2147483647) {
                        case 128:
                            v10 := []int{(fi9 << ((1) & 63)), fi8, fi9}
                            _ = v10
                        case -1:
                            switch 996 {
                            case -981:
                                v11 := GenS1{F0: -170, F1: 42}
                                _ = v11
                            case 128:
                                v3 = ((((("hello" + (v4 + (("a" + ("\n\t" + v4)) + ""))) + v4) + (v3 + v3)) + v4) + (v3 + (("a" + v4) + v3)))
                            case 482:
                                var v12 uint32 = 340
                                _ = v12
                            }
                        default:
                            var v13 int32 = (v2 - 1)
                            _ = v13
                        }
                        v1 = 623
                        var v14 uint32 = 1
                        v2 = (v2 - v2)
                        _ = v14
                    }
                    fmt.Println(v4, v5, fi9)
                    v15 := GenS2{F0: -374, F1: 65535}
                    _ = v15
                }
                _ = fi9
            }
            v16 := GenS3{F0: -325}
            _ = fi8
            _ = v16
        }
        v17 := []int{128, -1}
        _ = v17
    }
    switch 32767 {
    case 32768:
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS2{F0: 1, F1: 2}
            _ = fi18
            _ = v19
        }
    }
    if (!false) {
        v20 := GenS4{F0: 9223372036854775807, F1: 231, F2: 1}
        v21 := []int{-1}
        _ = v20
        _ = v21
    }
    v4 = (v4 + v3)
    var v22 uint32 = 42
    var v23 int = 42
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 int = -9223372036854775808
        _ = fi24
        _ = v25
    }
    fmt.Println(v3, v2, v23)
    fmt.Println(v23)
    v3 = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
