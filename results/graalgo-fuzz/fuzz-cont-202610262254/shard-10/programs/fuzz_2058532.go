package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint }

func main() {
    var v1 int = -9223372036854775808
    v2 := []string{}
    var v3 int8 = -1
    var v4 int64 = 42
    var v5 uint8 = 128
    _ = len(v2)
    if ((float32(1e10) <= -2.25) || (true && (0 >= 2147483648))) {
        if false {
            fmt.Println(v4, v3, v4)
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        v4 = v4
        _ = fi6
    }
    var v7 uint16 = 353
    switch v1 {
    case 127:
        v8 := GenS1{F0: (v4 * v4)}
        _ = v8
    default:
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS2{F0: 4294967295}
            _ = fi9
            _ = v10
        }
    }
    switch (v1 >> ((2) & 63)) {
    case 32767:
        v11 := GenS1{F0: (v4 / ((v4) | 1))}
        _ = v11
    case -1:
        var v12 int64 = (v4 / (((v4 & (v4 ^ v4))) | 1))
        _ = v12
    case 9223372036854775807:
        if len(v2) > 0 {
            v2[0] = ("foo bar" + ("foo bar" + ("a" + ("\n\t" + "\n\t"))))
        }
    default:
        if (!(42 > 151)) {
            v13 := GenS1{F0: v4}
            _ = v13
        } else {
            if (false && true) {
                var v14 int64 = (v4 % ((0) | 1))
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v15 string = "foo bar"
                    v15 = v15
                    _ = v15
                }
                v3 += v3
                _ = v14
            } else {
                var v16 int8 = (v3 ^ 127)
                v17 := GenS3{F0: "a", F1: 65535, F2: 0}
                _ = v16
                _ = v17
            }
            var v18 int = 2147483647
            _ = v18
        }
    }
    v19 := GenS3{F0: (("foo bar" + "\n\t") + "\n\t"), F1: 2, F2: 399}
    v2 = append(v2, "a")
    v20 := GenS4{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
