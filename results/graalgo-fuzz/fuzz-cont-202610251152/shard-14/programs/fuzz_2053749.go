package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint; F1 bool }

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]string)
    var v3 uint32 = 2
    v3 -= 0
    v4 := []int{32767}
    if (-9223372036854775808 >= -1) {
        for fi5 := 0; fi5 < 2; fi5++ {
            delete(v2, "\n\t")
            var v6 bool = (true || false)
            v7 := GenS1{F0: -977, F1: 9223372036854775807, F2: (fi5 + fi5)}
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 int16 = 212
        if true {
            var v9 uint32 = (v3 ^ 42)
            delete(v2, "hello")
            _ = v9
        }
        _ = v8
    } else {
        v3 = (v3 & (^(v3)))
        v1 = (v1 - v1)
        if true {
            v3 -= (v3 * v3)
        } else {
            if (!(uint(v1) <= 255)) {
                if len(v4) > 0 {
                    v4[0] = 32767
                }
                v1 += (v1 >> ((4294967295) & 63))
            }
            fmt.Println(v3, v1)
            if (4294967295 != (^(v3))) {
                var v10 uint64 = 199
                _ = v10
            }
        }
    }
    v11 := GenS1{F0: 0, F1: -497, F2: 32767}
    v1 = (^(v1))
    _ = len(v2)
    if true {
        switch ("a" + "hello") {
        case "a":
            var v12 uint = 481
            _ = v12
        case "":
            v13 := v2["hello"]
            _ = v13
        case "foo bar":
            var v14 int16 = 1000
            _ = v14
        default:
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 float64 = 0.0
                fmt.Println(v1, v3)
                v3 += 2147483648
                _ = fi15
                _ = v16
            }
        }
    } else {
        v17 := GenS1{F0: -639, F1: 58, F2: 2147483647}
        var v18 uint16 = 65535
        fmt.Println(v18, v18, v3)
        _ = v17
        _ = v18
    }
    switch -12 {
    case 9223372036854775807:
        v1 = (v1 - (v1 * (v1 ^ v1)))
    case 0:
        v19 := GenS1{F0: 0, F1: 544, F2: -694}
        _ = v19
    default:
        var v20 uint64 = 0
        _ = v20
    }
    var v21 float64 = -2.25
    switch "\n\t" {
    case "a":
        switch ("foo bar" + (((("" + "a") + ("\n\t" + "hello")) + "\n\t") + "\n\t")) {
        case "":
            if (-170 <= 2147483647) {
                v22 := GenS2{F0: 1, F1: 415}
                _ = v22
            } else {
                if (129 < v1) {
                    v3 = 77
                    v21 += v21
                }
                delete(v2, (("hello" + "foo bar") + "a"))
            }
        case "hello":
            v21 = 0.0
        case "\n\t":
            for fi23 := 0; fi23 < 17; fi23++ {
                v21 += ((((v21 + (float64(1.5) * v21)) * ((0.0 + (v21 / 1.0)) / 1.0)) / v21) / 1.0)
                for fi24 := 0; fi24 < 2; fi24++ {
                    v1 = 134
                    v3 += (^(v3))
                    v21 = v21
                    _ = fi24
                }
                if (v1 <= 0) {
                    v25 := v2[("" + "")]
                    v3 = v3
                    if true {
                        var v26 uint32 = v3
                        v3 = (v3 - (^(v3)))
                        _ = v26
                    } else {
                        v27 := GenS3{F0: 723, F1: (v1 != v1)}
                        fmt.Println(v1, v3, v25)
                        _ = v27
                    }
                    _ = v25
                } else {
                    var v28 float64 = (3.14159 - ((float64(0.1) - 1e10) / v21))
                    _ = v28
                }
                _ = fi23
            }
        default:
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 uint8 = v1
                _ = fi29
                _ = v30
            }
        }
    }
    switch 287 {
    case 32767:
        for fi31 := 0; fi31 < 7; fi31++ {
            switch (fi31 % ((fi31) | 1)) {
            case 327:
                v2["\n\t"] = "foo bar"
            }
            _ = fi31
        }
    case 9223372036854775807:
        v32 := v2["foo bar"]
        _ = v32
    default:
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
    }
    v1 -= v1
    fmt.Println(v1, v21, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
