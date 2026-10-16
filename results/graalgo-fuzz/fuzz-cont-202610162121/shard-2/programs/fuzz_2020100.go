package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int64 }
func genFn0(p0 int) string {
    v5 := []int{p0, p0, p0, 4294967295}
    switch (("" + "a") + ((("\n\t" + "") + "hello") + "a")) {
    case "foo bar":
        p0 += 880
    case "a":
        fmt.Println(p0)
    default:
        if true {
            var v6 uint8 = 1
            _ = v6
        } else {
            fmt.Println(p0)
        }
    }
    v7 := GenS1{F0: (p0 | p0), F1: 473, F2: 128}
    for fi8 := 0; fi8 < 7; fi8++ {
        p0 += fi8
        var v9 uint8 = 88
        var v10 int8 = 100
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v7
    return (((("" + "hello") + "hello") + ("foo bar" + "a")) + ("\n\t" + (("\n\t" + "\n\t") + "foo bar")))
}
func genFn1() string {
    var v11 uint = 2
    if true {
        var v12 uint64 = 709
        v13 := GenS1{F0: 1, F1: 2147483647, F2: 0}
        _ = v12
        _ = v13
    }
    _ = v11
    return "hello"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 42
    var v3 int8 = 1
    var v4 float64 = 0.1
    var v14 int64 = -1
    _ = len(v1)
    v2 -= v2
    v15 := GenS2{F0: -32768, F1: 1}
    v4 = (1.5 - v4)
    var v16 string = "\n\t"
    v17 := []int{9223372036854775807, 4294967295}
    var v18 int16 = -682
    switch int(v14) {
    case 128:
        v3 -= (v3 >> ((0) & 63))
    default:
        v16 = v16
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    v19 := GenS2{F0: -32768, F1: 42}
    var v20 int8 = v3
    if true {
        if false {
            var v21 int32 = 898
            _ = v21
        }
        fmt.Println(v16, v2)
    } else {
        if len(v17) > 0 {
            v17[0] = 128
        }
        switch "\n\t" {
        case "a":
            v22 := GenS3{F0: v2, F1: 18446744073709551615}
            _ = v22
        case "":
            for fi23 := 0; fi23 < 20; fi23++ {
                delete(v1, 0)
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(v14)
                    var v25 int64 = (v14 & (v14 ^ int64(v18)))
                    fmt.Println(v16, fi24)
                    if (246 < 535) {
                        v1[1] = 1
                        v26 := GenS4{F0: (-(v14))}
                        _ = v26
                    } else {
                        v14 = v25
                        v25 += v14
                        v3 = (^(v20))
                    }
                    var v27 uint64 = 712
                    _ = fi24
                    _ = v25
                    _ = v27
                }
                _ = fi23
            }
        case "\n\t":
            switch "" {
            case "hello":
                var v28 int8 = v20
                _ = v28
            case "":
                v29 := GenS2{F0: (v18 / ((v18) | 1)), F1: 1}
                _ = v29
            case "a":
                fmt.Println(v3, v20)
            }
        default:
            if true {
                var v30 int8 = v20
                var v31 int16 = (v18 + 1)
                _ = v30
                _ = v31
            } else {
                v14 = 98
            }
        }
        if (2147483648 != 0) {
            fmt.Println(v2, v3)
            v32 := GenS3{F0: v2, F1: 2}
            for fi33 := 0; fi33 < 1; fi33++ {
                switch (fi33 ^ 128) {
                case -336:
                    fmt.Println(v18)
                }
                _ = fi33
            }
            _ = v32
        }
    }
    var v34 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
}
