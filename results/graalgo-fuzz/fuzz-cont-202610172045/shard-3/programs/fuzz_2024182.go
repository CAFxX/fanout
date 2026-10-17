package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int; F1 int8 }
func genFn0(p0 uint16, p1 int64, p2 bool) int8 {
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(p0, p0, p1)
        fmt.Println(fi6, p0, p0)
        _ = fi6
    }
    var v7 int16 = -32768
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 int8, p1 uint64) string {
    p0 += p0
    p0 -= (^(p0))
    switch "foo bar" {
    case "\n\t":
        v9 := GenS1{F0: genFn0(65535, -9223372036854775808, false)}
        _ = v9
    case "hello":
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return "hello"
}

func main() {
    var v1 uint16 = 98
    var v2 float32 = 1.5
    var v3 int16 = 42
    v4 := make(map[string]int)
    v5 := make(map[string]int)
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := []int{fi10, fi10}
        _ = fi10
        _ = v11
    }
    var v12 int64 = 910
    if false {
        v13 := []int{-9223372036854775808, -9223372036854775808}
        delete(v5, "a")
        fmt.Println(v12)
        _ = v13
    } else {
        if (v3 != -32768) {
            fmt.Println(v2)
        } else {
            fmt.Println(v1, v12, v2)
        }
        var v14 int64 = 9223372036854775807
        for fi15 := 0; fi15 < 2; fi15++ {
            for fi16 := 0; fi16 < 18; fi16++ {
                v17 := GenS1{F0: 0}
                v18 := GenS2{F0: 974}
                var v19 uint8 = 0
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            if ((false && ((42 > 42) && (("" + "\n\t") < (("foo bar" + "hello") + ("a" + ("" + "\n\t")))))) && (42 < 1)) {
                v2 -= v2
            } else {
                var v20 uint64 = 2
                switch (fi15 ^ -444) {
                case -9223372036854775808:
                    v21 := GenS2{F0: (v1 / ((409) | 1))}
                    _ = v21
                case -2147483648:
                    v22 := v4[""]
                    _ = v22
                default:
                    v23 := GenS3{F0: genFn0(v1, (v14 & v14), (!true)), F1: (v3 >> ((1) & 63))}
                    _ = v23
                }
                _ = v20
            }
            _ = fi15
        }
        _ = v14
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 int8 = 127
        v26 := GenS4{F0: (fi24 | (fi24 ^ fi24)), F1: genFn0(1, (v12 + v12), (0 == 255))}
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 int8 = (v25 / (((v25 / (((v25 ^ (v25 << ((0) & 63)))) | 1))) | 1))
            _ = fi27
            _ = v28
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    var v29 int32 = int32(v12)
    fmt.Println(v12, v2)
    v1 = v1
    v30 := v4[((("a" + "foo bar") + "hello") + "foo bar")]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
