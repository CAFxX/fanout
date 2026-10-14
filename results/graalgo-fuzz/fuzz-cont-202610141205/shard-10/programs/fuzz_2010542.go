package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint }

func main() {
    var v1 float32 = ((0.0 - 1.5) - 1e10)
    var v2 uint64 = 4294967295
    var v3 uint32 = 4294967295
    v2 = (^(v2))
    v2 -= (v2 + v2)
    v1 = ((v1 - 0.0) / (v1 - ((float32(1e10) + 0.0) - v1)))
    var v4 int64 = -995
    var v5 uint64 = v2
    if (1 != v3) {
        var v6 string = "\n\t"
        if (((v3 % ((v3) | 1)) > (v3 ^ v3)) && (47 > 0)) {
            var v7 uint8 = 0
            _ = v7
        } else {
            v8 := []int{-588}
            _ = v8
        }
        fmt.Println(v4)
        _ = v6
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        switch fi9 {
        case 4294967295:
            fmt.Println(v4)
        case 2147483647:
            if true {
                if (0 == (fi9 * (fi9 | (-(fi9))))) {
                    v5 -= v5
                    v5 = 1
                    var v10 uint32 = v3
                    _ = v10
                } else {
                    v3 = (v3 & v3)
                }
                fmt.Println(v2, v1, v5)
            } else {
                fmt.Println(v5)
            }
        case -859:
            if false {
                v4 -= -9223372036854775808
            } else {
                var v11 uint16 = 2
                var v12 string = "hello"
                v2 = v5
                _ = v11
                _ = v12
            }
        }
        if (!(int32(v5) <= -2147483648)) {
            v5 -= 453
        } else {
            var v13 uint16 = 1
            _ = v13
        }
        switch "foo bar" {
        case "hello":
            v4 = v4
        default:
            fmt.Println(v2, v1)
        }
        _ = fi9
    }
    fmt.Println(v3)
    switch (("foo bar" + "hello") + "foo bar") {
    case "a":
        v4 += (v4 % (((v4 / ((v4) | 1))) | 1))
    case "\n\t":
        if (200 != v5) {
            if (v4 <= v4) {
                v14 := GenS1{F0: 65535, F1: 9223372036854775807, F2: "hello"}
                v3 = v3
                _ = v14
            } else {
                v15 := GenS2{F0: 16}
                v16 := GenS3{F0: 18446744073709551615, F1: 42, F2: 18446744073709551615}
                for fi17 := 0; fi17 < 8; fi17++ {
                    switch (fi17 ^ (fi17 >> ((42) & 63))) {
                    case 415:
                        v5 += (v2 << ((65535) & 63))
                    case -2147483648:
                        v18 := []int{(^(fi17)), 890, fi17, (^(fi17))}
                        _ = v18
                    case 0:
                        var v19 float64 = float64(v1)
                        _ = v19
                    default:
                        var v20 int8 = 0
                        _ = v20
                    }
                    if (true && true) {
                        if false {
                            var v21 int32 = -516
                            v3 = 0
                            _ = v21
                        }
                    }
                    var v22 uint64 = 0
                    _ = fi17
                    _ = v22
                }
                _ = v15
                _ = v16
            }
        }
    case "hello":
        var v23 int64 = (-(v4))
        _ = v23
    }
    v4 = v4
    v5 = (v5 + v2)
    v2 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
