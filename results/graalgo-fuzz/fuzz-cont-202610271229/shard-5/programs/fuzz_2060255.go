package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 string }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 int8, p1 uint, p2 uint32) bool {
    switch ("" + "") {
    case "\n\t":
        if true {
            var v7 int = 1
            var v8 int64 = 0
            v7 = v7
            _ = v7
            _ = v8
        } else {
            v9 := []int{902, 0}
            v10 := GenS1{F0: -9223372036854775808, F1: 32767, F2: "\n\t"}
            _ = v10
            _ = v9
        }
    case "hello":
        for fi11 := 0; fi11 < 13; fi11++ {
            fmt.Println(fi11, fi11)
            var v12 string = "hello"
            _ = fi11
            _ = v12
        }
    default:
        var v13 float64 = (float64(((0.1 - float32(3.14159)) * float32(1e10))) / (1e10 * 3.14159))
        _ = v13
    }
    p0 -= p0
    v14 := []int{0, -9223372036854775808, -848}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (32767 > 0)
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{42}
    v3 := []string{}
    v4 := make(map[string]string)
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    var v15 float32 = float32(float64((0.0 - (0.1 - 3.14159))))
    var v16 uint64 = 2
    v16 -= v16
    var v17 float32 = v15
    if false {
        v18 := GenS2{F0: 9223372036854775807, F1: 1}
        v19 := GenS1{F0: 0, F1: 1, F2: ("\n\t" + (("hello" + "hello") + ""))}
        _ = v18
        _ = v19
    } else {
        v16 -= (v16 + (^(v16)))
        switch ("hello" + ("" + ((("hello" + "foo bar") + "foo bar") + "hello"))) {
        case "":
            if true {
                v17 += (v15 + v15)
                fmt.Println(v17, v16)
            } else {
                var v20 int32 = 42
                v21 := GenS3{F0: (^(v16)), F1: 0}
                v16 -= 42
                _ = v20
                _ = v21
            }
        case "a":
            switch 42 {
            case -9223372036854775808:
                v16 = (v16 & (v16 | v16))
            default:
                v17 += -2.25
            }
        case "foo bar":
            switch -329 {
            case 183:
                v17 = float32((0.0 / (0.1 + (0.1 - float64(v17)))))
            default:
                v22 := GenS4{F0: "a", F1: 42}
                _ = v22
            }
        default:
            fmt.Println(v16, v15, v15)
        }
    }
    var v23 bool = (821 >= 418)
    for fi24 := 0; fi24 < 14; fi24++ {
        fmt.Println(v17, v17, v23)
        _ = fi24
    }
    v3 = append(v3, ("a" + "a"))
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 int = fi25
        v26 -= (-(fi25))
        _ = fi25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
