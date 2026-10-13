package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int16; F1 int32; F2 bool }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint) int64 {
    v6 := GenS1{F0: 42, F1: 1000}
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 9223372036854775807
}
func genFn1() uint32 {
    v7 := []int{}
    v8 := GenS2{F0: 32768, F1: false}
    switch -1 {
    case -1:
        var v9 uint64 = 4294967295
        _ = v9
    default:
        switch -9223372036854775808 {
        case -2147483648:
            switch "foo bar" {
            case "a":
                v10 := GenS3{F0: 42, F1: 2147483647}
                _ = v10
            case "foo bar":
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            default:
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 int = fi11
                    _ = fi11
                    _ = v12
                }
            }
        case -9223372036854775808:
            v7 = append(v7, 225)
        }
    }
    var v13 int32 = 0
    _ = v13
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 float64 = 3.14159
    var v2 int16 = 32767
    v3 := make(map[string]string)
    v4 := []string{"\n\t", "foo bar", (("hello" + "foo bar") + ("\n\t" + ""))}
    var v5 int16 = (v2 + (-(v2)))
    _ = len(v3)
    var v14 int64 = 9223372036854775807
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    if (0.1 >= (((-2.25 / 1.0) - -2.25) + 0.1)) {
        var v15 uint = 2
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := GenS2{F0: (fi16 ^ 595), F1: (2147483647 == 0)}
        _ = fi16
        _ = v17
    }
    var v18 bool = true
    if ((-2.25 <= (3.14159 / 1.0)) && v18) {
        v19 := GenS4{F0: (v5 * v2), F1: 0}
        v1 -= 3.14159
        _ = v19
    } else {
        var v20 float64 = (v1 - (v1 / v1))
        if v18 {
            v18 = (186 <= 141)
            if len(v4) > 0 {
                v4[0] = ("" + "foo bar")
            }
            v21 := GenS5{F0: v2, F1: -519, F2: ((!v18) || ((1e10 > 0.1) || v18))}
            _ = v21
        }
        delete(v3, "\n\t")
        _ = v20
    }
    v22 := GenS1{F0: 177, F1: v5}
    fmt.Println(v2, v1, v5)
    if ((v18 && false) || v18) {
        switch "" {
        case "":
            v18 = v18
        case "hello":
            var v23 uint = 4294967295
            _ = v23
        default:
            fmt.Println(v2, v18, v14)
        }
        v18 = (971 == 478)
        v2 += v2
    }
    v2 -= (-(v5))
    if ("foo bar" > ("foo bar" + "")) {
        v24 := GenS6{F0: -1}
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        v5 = (-(v5))
        var v26 int = (fi25 - (-(fi25)))
        v2 -= v5
        _ = fi25
        _ = v26
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v14, v5, v18)
    if (!false) {
        var v27 uint16 = 340
        if (!(!false)) {
            var v28 int64 = -9223372036854775808
            _ = v28
        } else {
            for fi29 := 0; fi29 < 17; fi29++ {
                v5 += v5
                if v18 {
                    fmt.Println(v14, v2)
                    for fi30 := 0; fi30 < 1; fi30++ {
                        var v31 int64 = (-(v14))
                        _ = fi30
                        _ = v31
                    }
                    var v32 int64 = (^(v14))
                    _ = v32
                }
                _ = fi29
            }
        }
        _ = v27
    }
    v33 := GenS4{F0: 42, F1: (^(v14))}
    if (!(!((v5 % ((v5) | 1)) <= v2))) {
        var v34 uint16 = 5
        fmt.Println(v18)
        _ = v34
    } else {
        switch int(v5) {
        case 4294967295:
            var v35 int32 = -2147483648
            _ = v35
        case 0:
            var v36 uint = 42
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
