package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int64; F2 string }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint64, p1 int) int8 {
    p1 = (p1 + (-(p1)))
    v5 := GenS1{F0: 18446744073709551615}
    if true {
        switch "hello" {
        case "hello":
            switch p1 {
            case 1:
                if false {
                    var v6 int64 = -9223372036854775808
                    var v7 int = p1
                    _ = v6
                    _ = v7
                }
            case 979:
                p1 = p1
            }
        default:
            v8 := GenS1{F0: 89}
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 8; fi9++ {
            p1 = 127
            p0 = (p0 % ((1) | 1))
            _ = fi9
        }
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 float32 = 1e10
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 float32 = (0.1 - (-2.25 + 0.1))
            _ = fi12
            _ = v13
        }
    }
    var v14 int32 = -706
    _ = p0
    _ = p1
    _ = v14
    _ = v5
    return 127
}

func main() {
    var v1 string = "foo bar"
    v2 := []string{}
    var v3 int = 2147483647
    var v4 int16 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v15 := GenS2{F0: (^(v4))}
    v16 := GenS3{F0: 127, F1: 9223372036854775807, F2: "hello"}
    switch 0 {
    case -538:
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    default:
        var v17 float64 = (float64(float32(3.14159)) * 3.14159)
        _ = v17
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v18 := GenS1{F0: 2}
    var v19 float64 = 1.5
    var v20 int = v3
    if len(v2) > 0 {
        v2[0] = (v1 + (v1 + v1))
    }
    var v21 uint8 = 175
    if len(v2) > 0 {
        v2[0] = v1
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 uint8 = v21
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 int16 = (v4 - v4)
            _ = fi24
            _ = v25
        }
        v3 = fi22
        _ = fi22
        _ = v23
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := GenS4{F0: ((false && (!(42 < 93))) || ((("hello" + v1) + ((("foo bar" + (v1 + v1)) + (("\n\t" + v1) + ((v1 + (v1 + v1)) + "\n\t"))) + "")) != (v1 + (((("hello" + v1) + (v1 + (v1 + ((((v1 + ("foo bar" + "a")) + v1) + ("foo bar" + ((v1 + (v1 + "\n\t")) + v1))) + v1)))) + v1) + "foo bar")))), F1: v3}
        var v28 int32 = -2147483648
        if false {
            _ = len(v2)
            fmt.Println(v4, v19, v19)
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS5{F0: v21, F1: 2147483647}
    v21 = v21
    var v30 string = ""
    v2 = append(v2, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
