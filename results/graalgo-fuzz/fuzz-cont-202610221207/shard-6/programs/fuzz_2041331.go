package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int; F2 uint32 }
func genFn0(p0 bool, p1 bool, p2 int32) uint8 {
    p2 += p2
    v7 := GenS1{F0: ((("\n\t" + "hello") + "foo bar") + ""), F1: 18446744073709551615, F2: 670}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 127
}

func main() {
    v1 := []string{"", "", "foo bar"}
    var v2 int32 = -1
    v3 := make(map[string]int)
    var v4 uint64 = 4294967295
    var v5 int = -1
    v6 := make(map[int]string)
    v2 += (-(v2))
    fmt.Println(v5, v4, v2)
    v8 := GenS2{F0: (v5 + v5)}
    if ((v4 % (((v4 - (^(v4)))) | 1)) >= 18446744073709551615) {
        if len(v1) > 0 {
            v1[0] = "a"
        }
    } else {
        switch "hello" {
        case "":
            var v9 uint = 0
            _ = v9
        case "\n\t":
            v4 += (v4 % (((^(v4))) | 1))
        }
        var v10 int16 = 1000
        _ = v10
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if true {
        v4 -= 2
        var v11 string = "hello"
        switch v11 {
        case "hello":
            v1 = append(v1, v11)
        default:
            for fi12 := 0; fi12 < 11; fi12++ {
                v5 = v5
                _ = fi12
            }
        }
        _ = v11
    } else {
        switch "hello" {
        case "foo bar":
            switch v5 {
            case 42:
                fmt.Println(v2, v4, v4)
            case 9223372036854775807:
                fmt.Println(v4, v5, v5)
            case 698:
                fmt.Println(v4)
            default:
                v5 = 0
            }
        case "a":
            v13 := GenS2{F0: 128}
            _ = v13
        default:
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := v6[fi14]
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi14
                _ = v15
            }
        }
    }
    switch "a" {
    case "":
        var v16 int64 = 1
        _ = v16
    case "hello":
        switch (("foo bar" + ("hello" + "\n\t")) + "\n\t") {
        case "foo bar":
            fmt.Println(v2, v2, v4)
        case "":
            v4 += 0
        case "a":
            v4 -= (^(v4))
        }
    default:
        v17 := GenS2{F0: v5}
        _ = v17
    }
    var v18 uint8 = 128
    v19 := GenS3{F0: 1, F1: (v5 ^ 484), F2: uint32(v18)}
    var v20 uint8 = 42
    v21 := v6[int(v2)]
    var v22 float32 = 0.1
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v23 := GenS2{F0: (^(v5))}
    fmt.Println(v4, v2, v20)
    switch v5 {
    case 2147483647:
        var v24 int = (v5 / ((v5) | 1))
        _ = v24
    case -848:
        if true {
            v25 := GenS3{F0: 0, F1: int(v2), F2: 216}
            _ = v25
        }
    case 128:
        fmt.Println(v18)
    }
    var v26 bool = true
    var v27 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
