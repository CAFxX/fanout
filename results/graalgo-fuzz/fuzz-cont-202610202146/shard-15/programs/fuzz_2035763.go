package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint }

func main() {
    v1 := []string{"", (("\n\t" + "\n\t") + "hello"), "\n\t"}
    v2 := []string{}
    var v3 int8 = 127
    var v4 uint = 4294967295
    v5 := GenS1{F0: 1, F1: "hello"}
    fmt.Println(v4, v3)
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 float64 = ((1.5 / 1e10) / 3.14159)
        _ = fi6
        _ = v7
    }
    v3 += (v3 & v3)
    v1 = append(v1, "")
    v3 -= v3
    v8 := GenS2{F0: (v4 * (v4 / (((v4 - (v4 ^ v4))) | 1))), F1: false}
    var v9 float64 = 0.1
    if (0 != -216) {
        switch "" {
        case "a":
            var v10 int = 1
            _ = v10
        case "hello":
            var v11 int = 0
            _ = v11
        default:
            v3 += (v3 >> ((v4) & 63))
        }
    } else {
        var v12 float64 = 1e10
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := GenS3{F0: (fi13 / ((fi13) | 1)), F1: 18446744073709551615, F2: 0}
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 17; fi16++ {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                switch "" {
                case "\n\t":
                    v4 = v4
                case "a":
                    var v17 uint = v4
                    _ = v17
                }
                var v18 float64 = v9
                v19 := GenS4{F0: (v4 - v4)}
                _ = fi16
                _ = v18
                _ = v19
            }
            v12 = v12
            _ = fi15
        }
        _ = v12
    }
    v20 := GenS2{F0: 301, F1: (v4 <= (^(v4)))}
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        if (!false) {
            var v22 int16 = -1
            _ = v22
        } else {
            v9 += v9
        }
        if (9223372036854775807 <= 683) {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            fmt.Println(v3)
            for fi23 := 0; fi23 < 11; fi23++ {
                fmt.Println(fi21)
                if len(v2) > 0 {
                    v2[0] = "a"
                }
                v4 = (v4 - v4)
                v3 = 1
                _ = fi23
            }
        } else {
            if ((0.1 - 3.14159) == 0.0) {
                fmt.Println(v9, v4, fi21)
            }
            if (true || (!(!false))) {
                var v24 uint64 = 2
                var v25 int64 = -9223372036854775808
                _ = v24
                _ = v25
            }
        }
        v1 = append(v1, "foo bar")
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
