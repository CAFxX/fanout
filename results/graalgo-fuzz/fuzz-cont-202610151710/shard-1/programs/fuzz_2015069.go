package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint64; F1 string }

func main() {
    var v1 int64 = 1
    var v2 string = "a"
    var v3 int16 = 32767
    var v4 int64 = v1
    var v5 int = -2147483648
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(v4, v5, v3)
        var v7 float64 = 1.5
        var v8 uint = 18446744073709551615
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 648, F1: (v5 & 42), F2: true}
    v10 := []int{-1, v5}
    switch v2 {
    case "hello":
        v4 += v1
    case "":
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 bool = ((("\n\t" + ((v2 + (v2 + v2)) + (v2 + v2))) + v2) < "hello")
                v1 = 42
                var v14 uint32 = 239
                _ = fi12
                _ = v13
                _ = v14
            }
            if len(v10) > 0 {
                v10[0] = v5
            }
            v3 -= (v3 + (v3 | 32767))
            _ = fi11
        }
    case "a":
        v15 := GenS1{F0: -1, F1: (v5 >> ((1) & 63)), F2: (42 > 162)}
        _ = v15
    default:
        for fi16 := 0; fi16 < 16; fi16++ {
            v5 = (v5 ^ (^(v5)))
            _ = len(v10)
            _ = fi16
        }
    }
    if ((v1 < v4) || false) {
        v5 = (-(v5))
        for fi17 := 0; fi17 < 16; fi17++ {
            v5 = fi17
            _ = fi17
        }
        for fi18 := 0; fi18 < 13; fi18++ {
            if (1e10 <= 0.1) {
                v2 = ((v2 + ("a" + ("foo bar" + "a"))) + "hello")
            } else {
                var v19 uint64 = 2
                for fi20 := 0; fi20 < 15; fi20++ {
                    v21 := GenS2{F0: 620, F1: v1}
                    _ = fi20
                    _ = v21
                }
                var v22 uint8 = 31
                _ = v19
                _ = v22
            }
            if true {
                fmt.Println(fi18)
                for fi23 := 0; fi23 < 13; fi23++ {
                    switch fi18 {
                    case 127:
                        v1 -= 9223372036854775807
                    case 0:
                        v3 += (v3 / (((v3 << ((2) & 63))) | 1))
                    }
                    _ = fi23
                }
                var v24 int64 = v1
                _ = v24
            }
            _ = fi18
        }
    } else {
        fmt.Println(v5, v2)
        var v25 string = ("" + ((v2 + ((("" + (("hello" + v2) + ("foo bar" + (v2 + (("" + "") + v2))))) + (("hello" + v2) + v2)) + "\n\t")) + v2))
        fmt.Println(v1, v25, v2)
        _ = v25
    }
    v5 = -2147483648
    if (270 < 4294967295) {
        switch 128 {
        case 9223372036854775807:
            var v26 uint = 0
            _ = v26
        case -830:
            fmt.Println(v2, v4, v3)
        case -2147483648:
            v5 = (^(v5))
        }
        switch (v5 ^ (v5 ^ 314)) {
        case 9223372036854775807:
            v27 := GenS1{F0: 1, F1: (v5 & (v5 ^ (v5 >> ((18446744073709551615) & 63)))), F2: (42 < 42)}
            _ = v27
        default:
            var v28 float32 = 0.1
            _ = v28
        }
    } else {
        var v29 int = v5
        v3 -= (v3 << ((343) & 63))
        _ = v29
    }
    v1 -= (v4 & 1)
    v30 := GenS3{F0: 4294967295, F1: (v2 + "a")}
    var v31 int16 = (v3 >> ((336) & 63))
    v32 := GenS1{F0: -1, F1: v5, F2: true}
    var v33 float32 = (1.5 / 1.0)
    v34 := GenS3{F0: 897, F1: v2}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v9
}
