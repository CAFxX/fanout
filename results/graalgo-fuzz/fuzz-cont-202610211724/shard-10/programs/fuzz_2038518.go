package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int64; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 string; F2 int8 }
func genFn0(p0 int16) uint {
    v7 := GenS1{F0: -9223372036854775808, F1: 2}
    var v8 string = "hello"
    var v9 int = -2147483648
    var v10 string = "a"
    var v11 uint32 = 983
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 208
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -2147483648
    var v3 uint16 = 215
    v4 := []string{(("foo bar" + "\n\t") + "a"), ("a" + ""), "a", ""}
    var v5 int64 = -1
    v6 := make(map[uint32]string)
    var v12 uint32 = 42
    var v13 uint32 = 1
    var v14 float64 = ((float64((float32(0.0) + -2.25)) - (3.14159 - ((1.5 / 1e10) - (1e10 - 1e10)))) - ((float64((float32((1e10 / float64(1.5))) - 1.5)) / -2.25) * 3.14159))
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int32 = v2
        _ = fi15
        _ = v16
    }
    switch "hello" {
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            switch "foo bar" {
            case "hello":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "foo bar":
                if false {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v5, v14)
                    default:
                        fmt.Println(v14, v5)
                    }
                    fmt.Println(v3, v2)
                    if true {
                        v17 := GenS1{F0: (v5 & (-(v5))), F1: 890}
                        _ = v17
                    }
                } else {
                    v2 += (v2 << ((420) & 63))
                    fmt.Println(v13, v12, v2)
                    switch "a" {
                    case "":
                        var v18 bool = true
                        _ = v18
                    default:
                        v19 := GenS1{F0: 42, F1: 4294967295}
                        _ = v19
                    }
                }
            default:
                v14 += v14
            }
        case "a":
            if (false || true) {
                _ = len(v4)
            }
        case "hello":
            if true {
                var v20 bool = (("\n\t" == "") || (-319 < -203))
                _ = v20
            }
        default:
            for fi21 := 0; fi21 < 4; fi21++ {
                fmt.Println(v3)
                v5 -= (^(v5))
                _ = fi21
            }
        }
    case "":
        v13 += 919
    case "hello":
        var v22 float64 = v14
        _ = v22
    default:
        v13 = 911
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        v2 = v2
        var v24 float64 = v14
        v24 = (v14 - v24)
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        if (true && true) {
            var v26 float64 = v14
            _ = v26
        } else {
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 int32 = v2
                _ = fi27
                _ = v28
            }
            var v29 float32 = 0.0
            _ = v29
        }
        _ = len(v6)
        _ = fi25
    }
    fmt.Println(v14, v14)
    if true {
        v3 = (^(v3))
        v30 := GenS2{F0: v5, F1: -9223372036854775808, F2: 2}
        _ = v30
    }
    if (!(-2.25 < 0.1)) {
        v31 := GenS3{F0: 880}
        _ = v31
    } else {
        v12 = v12
    }
    v32 := GenS1{F0: 9223372036854775807, F1: 2}
    v2 += (v2 % ((v2) | 1))
    var v33 string = "hello"
    var v34 uint = 0
    delete(v6, 0)
    v3 = (v3 << (((v34 % ((v34) | 1))) & 63))
    var v35 string = (v33 + v33)
    for fi36 := 0; fi36 < 4; fi36++ {
        v37 := GenS4{F0: 42, F1: v35, F2: 0}
        for fi38 := 0; fi38 < 4; fi38++ {
            v34 += (v34 / ((4294967295) | 1))
            var v39 int32 = (^(v2))
            _ = fi38
            _ = v39
        }
        v40 := GenS2{F0: (v5 ^ (v5 * 32)), F1: (v5 >> ((v34) & 63)), F2: 727}
        _ = fi36
        _ = v37
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
