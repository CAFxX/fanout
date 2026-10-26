package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int16; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint16, p1 uint8) uint {
    var v7 int8 = -1
    v8 := GenS1{F0: ((((("foo bar" + ("foo bar" + "\n\t")) + "\n\t") + "") + "a") + ("\n\t" + "foo bar"))}
    p0 = 0
    switch "" {
    case "a":
        var v9 uint64 = 4294967295
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS2{F0: 106, F1: 370}
            _ = fi10
            _ = v11
        }
    case "\n\t":
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint16 = p0
            p1 -= (p1 / ((127) | 1))
            switch "foo bar" {
            case "\n\t":
                v14 := []int{}
                _ = v14
            }
            _ = fi12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 72
}
func genFn1(p0 int8, p1 uint32, p2 uint16) string {
    for fi15 := 0; fi15 < 15; fi15++ {
        p0 = (^(p0))
        var v16 uint = 255
        _ = fi15
        _ = v16
    }
    v17 := GenS3{F0: 32767, F1: 32767, F2: 32767}
    for fi18 := 0; fi18 < 17; fi18++ {
        p0 = p0
        fmt.Println(p2, p2)
        p0 = 1
        _ = fi18
    }
    var v19 int8 = (p0 + -21)
    v20 := GenS4{F0: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v19
    _ = v20
    return ""
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 32767
    var v3 int64 = -816
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    var v6 float64 = 0.1
    if false {
        var v21 int8 = -128
        switch "a" {
        case "":
            v22 := GenS1{F0: "\n\t"}
            _ = v22
        case "\n\t":
            for fi23 := 0; fi23 < 20; fi23++ {
                if (genFn0(42, 220) >= 4294967295) {
                    v24 := GenS5{F0: 128, F1: 913}
                    _ = v24
                } else {
                    v21 -= v21
                }
                _ = fi23
            }
        case "a":
            v6 -= (0.1 - v6)
        }
        _ = v21
    }
    fmt.Println(v6, v2, v6)
    fmt.Println(v3, v2, v3)
    v25 := GenS4{F0: 2}
    v1["\n\t"] = 42
    var v26 int16 = v2
    if (-9223372036854775808 != 189) {
        fmt.Println(v3, v3, v6)
        if false {
            fmt.Println(v26)
            v26 = (-(v26))
            var v27 int32 = 354
            _ = v27
        }
        v3 += (v3 ^ (v3 % (((v3 | 358)) | 1)))
    } else {
        var v28 bool = true
        v26 += (v2 % ((v26) | 1))
        fmt.Println(v2, v6)
        _ = v28
    }
    if false {
        fmt.Println(v6, v2)
        var v29 int8 = -128
        _ = v29
    }
    v2 += -1
    if true {
        v6 += float64(0.0)
        v30 := GenS1{F0: (("a" + "a") + "hello")}
        _ = v30
    } else {
        v3 = (v3 * v3)
    }
    if false {
        var v31 int8 = 1
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
