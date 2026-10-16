package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
func genFn0() string {
    var v6 float32 = (0.0 * 1.5)
    v7 := GenS1{F0: 1000}
    if (v6 < (float32((1.5 / ((float64(v6) + 0.0) / (0.0 / float64(float32(float64(v6))))))) - v6)) {
        v8 := GenS1{F0: -1}
        v9 := []int{-2147483648, -9223372036854775808}
        var v10 uint16 = 242
        _ = v10
        _ = v8
        _ = v9
    } else {
        v11 := []int{4294967295}
        _ = v11
    }
    _ = v6
    _ = v7
    return "\n\t"
}

func main() {
    var v1 int32 = 104
    var v2 uint32 = 65535
    v3 := []uint32{v2, (^(v2)), v2}
    v4 := []string{"\n\t"}
    var v5 uint64 = 1
    v5 = v5
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 int8 = 127
        v14 := GenS1{F0: 337}
        if false {
            v1 = (v1 + (v1 ^ -126))
            for fi15 := 0; fi15 < 1; fi15++ {
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi15
            }
        } else {
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 uint32 = v2
                if (false && true) {
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                } else {
                    v1 = v1
                    if true {
                        v13 = v13
                    }
                }
                var v18 int32 = v1
                v13 = (-(v13))
                var v19 uint8 = 1
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            switch "hello" {
            case "hello":
                v2 = 42
            default:
                v2 = v2
            }
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    v5 = (v5 ^ v5)
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 int64 = 265
        _ = fi20
        _ = v21
    }
    if false {
        _ = len(v3)
        switch "foo bar" {
        case "hello":
            v5 -= v5
        case "":
            v22 := GenS2{F0: 294}
            _ = v22
        case "\n\t":
            for fi23 := 0; fi23 < 12; fi23++ {
                v4 = append(v4, (((("a" + "hello") + "a") + "a") + ("\n\t" + ("a" + ("a" + "foo bar")))))
                var v24 uint32 = 2147483647
                v24 -= (v24 - v2)
                _ = fi23
                _ = v24
            }
        }
    } else {
        var v25 string = "a"
        switch v25 {
        case "\n\t":
            v1 += (^(v1))
        case "foo bar":
            switch 2147483647 {
            case 32:
                fmt.Println(v25)
            case -1:
                var v26 int16 = -173
                _ = v26
            default:
                v1 -= (v1 * v1)
            }
        }
        _ = v25
    }
    fmt.Println(v1)
    v27 := GenS2{F0: v5}
    v28 := GenS2{F0: (^(v5))}
    if (int8(v2) <= -1) {
        fmt.Println(v2, v1, v2)
        switch 624 {
        case -846:
            var v29 uint8 = 1
            _ = v29
        }
    } else {
        var v30 string = "foo bar"
        var v31 int16 = 32767
        _ = v30
        _ = v31
    }
    var v32 uint64 = 960
    v2 += (v2 * 940)
    switch "foo bar" {
    case "\n\t":
        var v33 int16 = 291
        _ = v33
    case "hello":
        if len(v3) > 0 {
            v3[0] = v2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
