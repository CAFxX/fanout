package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int32; F1 int16; F2 int }
type GenS4 struct { F0 int }
func genFn0(p0 int, p1 int32) uint {
    var v7 bool = ("" < "hello")
    switch "" {
    case "\n\t":
        var v8 uint64 = 2
        _ = v8
    }
    switch p0 {
    case 127:
        var v9 int64 = 1
        _ = v9
    case -866:
        var v10 int8 = 100
        _ = v10
    default:
        var v11 float32 = (0.0 + float32(1e10))
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 bool = v7
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 42
    var v3 int64 = -9223372036854775808
    v4 := []string{}
    var v5 int16 = 276
    var v6 int = -2147483648
    fmt.Println(v6, v3, v3)
    var v14 bool = (!false)
    v6 -= v6
    var v15 bool = (42 < 0)
    switch "\n\t" {
    case "a":
        v16 := GenS1{F0: 2, F1: ("a" + ("a" + "a"))}
        _ = v16
    case "foo bar":
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := GenS2{F0: 0, F1: 0}
            _ = fi17
            _ = v18
        }
    default:
        v2 = v2
    }
    v19 := GenS2{F0: 2, F1: v3}
    for fi20 := 0; fi20 < 1; fi20++ {
        fmt.Println(v3, v2)
        v21 := GenS3{F0: 249, F1: v5, F2: (^(fi20))}
        _ = fi20
        _ = v21
    }
    var v22 uint8 = (^(v2))
    for fi23 := 0; fi23 < 9; fi23++ {
        v24 := GenS1{F0: 951, F1: "hello"}
        if v15 {
            v25 := GenS4{F0: (fi23 >> ((18446744073709551615) & 63))}
            _ = v25
        } else {
            if (!v15) {
                fmt.Println(v14, v15)
            } else {
                if (v15 && (v5 > v5)) {
                    v14 = ((283 == 42) || (v5 >= 0))
                }
                delete(v1, ("foo bar" + "\n\t"))
                var v26 uint = 0
                _ = v26
            }
        }
        _ = len(v1)
        _ = fi23
        _ = v24
    }
    var v27 uint8 = 42
    v5 += 1
    if false {
        if (363 >= 0) {
            v27 = v22
            _ = len(v4)
        }
        switch (v6 % (((^(v6))) | 1)) {
        case 580:
            v27 = (v2 / (((v2 + v2)) | 1))
        }
    } else {
        for fi28 := 0; fi28 < 8; fi28++ {
            if v15 {
                fmt.Println(v3, fi28)
            }
            _ = fi28
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
