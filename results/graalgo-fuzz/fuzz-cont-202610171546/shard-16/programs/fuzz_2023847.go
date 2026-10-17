package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint16; F2 int8 }
func genFn0(p0 uint8) uint16 {
    if (false && false) {
        v5 := GenS1{F0: -457}
        var v6 int32 = 732
        _ = v5
        _ = v6
    } else {
        switch 21 {
        case 0:
            var v7 float32 = ((0.0 * float32((0.1 - 1.5))) * (float32((1.5 - 3.14159)) / 1.0))
            _ = v7
        case 42:
            var v8 uint32 = 922
            _ = v8
        default:
            fmt.Println(p0)
        }
        v9 := GenS2{F0: (p0 >> ((42) & 63)), F1: -61}
        _ = v9
    }
    var v10 uint32 = 2
    if (!(!(((1.5 * 0.1) >= 0.1) && true))) {
        v10 += (v10 | (v10 | (v10 & (v10 / (((v10 - (v10 * 4294967295))) | 1)))))
    }
    var v11 int16 = -1
    for fi12 := 0; fi12 < 10; fi12++ {
        if (!true) {
            v13 := GenS1{F0: v11}
            v14 := GenS3{F0: v11, F1: (p0 ^ p0)}
            _ = v13
            _ = v14
        } else {
            var v15 float32 = float32(3.14159)
            _ = v15
        }
        _ = fi12
    }
    _ = p0
    _ = v10
    _ = v11
    return 65535
}
func genFn1(p0 uint16, p1 int8, p2 int32) string {
    switch "\n\t" {
    case "\n\t":
        var v16 int = 42
        _ = v16
    case "hello":
        v17 := []int{-1, 32768, 0, -504}
        _ = v17
    case "a":
        fmt.Println(p1, p2, p1)
    }
    v18 := GenS1{F0: 1}
    var v19 uint16 = (^(p0))
    fmt.Println(v19)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return ("" + "foo bar")
}

func main() {
    var v1 int64 = 386
    var v2 int32 = -1
    var v3 int64 = v1
    v4 := []string{("a" + "a"), "hello", (("\n\t" + "hello") + ""), "foo bar"}
    v3 -= v1
    fmt.Println(v3, v2)
    var v20 int16 = 0
    if true {
        if len(v4) > 0 {
            v4[0] = "a"
        }
    }
    _ = len(v4)
    for fi21 := 0; fi21 < 11; fi21++ {
        v1 -= v1
        var v22 float32 = 1.5
        fmt.Println(v22, v22, v22)
        _ = fi21
        _ = v22
    }
    switch (("foo bar" + "") + ("a" + (("" + ("a" + "foo bar")) + ""))) {
    case "":
        v23 := GenS1{F0: v20}
        _ = v23
    }
    fmt.Println(v1)
    for fi24 := 0; fi24 < 19; fi24++ {
        for fi25 := 0; fi25 < 19; fi25++ {
            if true {
                switch "\n\t" {
                case "a":
                    var v26 uint16 = 320
                    _ = v26
                case "\n\t":
                    v3 = v1
                default:
                    var v27 float64 = 1.5
                    _ = v27
                }
                v1 -= v1
            } else {
                var v28 int32 = 396
                fmt.Println(fi25, v20, v2)
                _ = v28
            }
            v29 := GenS4{F0: 42, F1: 255, F2: 100}
            var v30 float64 = 0.1
            _ = fi25
            _ = v29
            _ = v30
        }
        switch -2147483648 {
        case 42:
            var v31 uint8 = 0
            _ = v31
        case 1:
            v4 = append(v4, ((((("foo bar" + ("hello" + ("a" + ("foo bar" + "")))) + ("hello" + (("a" + "") + (("hello" + ((("\n\t" + ("hello" + ("hello" + "\n\t"))) + "hello") + "hello")) + "\n\t")))) + ("foo bar" + "foo bar")) + "") + ((("" + (("\n\t" + "\n\t") + ((((("a" + "foo bar") + "foo bar") + "hello") + "a") + "foo bar"))) + (("\n\t" + "\n\t") + ("\n\t" + "\n\t"))) + "\n\t")))
        }
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
