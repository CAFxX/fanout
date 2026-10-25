package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int64; F2 int32 }
type GenS3 struct { F0 int32 }
func genFn0() int64 {
    if (1 == 1000) {
        for fi5 := 0; fi5 < 14; fi5++ {
            switch (-(fi5)) {
            case -1:
                var v6 float64 = float64((1e10 - 0.0))
                _ = v6
            case 0:
                var v7 string = (("\n\t" + ("hello" + "a")) + "hello")
                _ = v7
            case -848:
                var v8 string = "a"
                _ = v8
            }
            _ = fi5
        }
        v9 := GenS1{F0: 0}
        var v10 float32 = 1e10
        _ = v10
        _ = v9
    } else {
        var v11 float32 = 0.1
        v12 := []int{128, 791, 32768, -9223372036854775808}
        _ = v11
        _ = v12
    }
    var v13 int16 = 1
    var v14 bool = true
    switch ("foo bar" + (("a" + "a") + ("" + ""))) {
    case "":
        v15 := GenS2{F0: 42, F1: 24, F2: -1}
        _ = v15
    }
    v16 := []int{-1}
    _ = v13
    _ = v14
    _ = v16
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 240
    var v2 float64 = (0.0 + 0.0)
    v3 := []string{"\n\t", "hello"}
    v4 := []string{"foo bar", "foo bar", (("" + "a") + "")}
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 uint32 = 2147483648
        v2 -= v2
        var v19 int8 = 127
        _ = fi17
        _ = v18
        _ = v19
    }
    if true {
        var v20 float64 = ((v2 / float64(0.1)) * ((1e10 * 0.1) / v2))
        _ = v20
    }
    if (((("a" + (("hello" + ("a" + "\n\t")) + ("hello" + (((((((("\n\t" + ("a" + "a")) + "foo bar") + "") + "foo bar") + "hello") + "") + "a") + (("hello" + ("foo bar" + "hello")) + "hello"))))) + (("hello" + ("\n\t" + "foo bar")) + "hello")) + "hello") >= (("\n\t" + "\n\t") + "a")) {
        v1 = (^(v1))
    } else {
        for fi21 := 0; fi21 < 8; fi21++ {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            v4 = append(v4, "foo bar")
            _ = fi21
        }
    }
    var v22 int8 = 1
    v23 := GenS1{F0: 70}
    switch "foo bar" {
    case "":
        v24 := GenS1{F0: 18446744073709551615}
        _ = v24
    case "foo bar":
        fmt.Println(v22, v2)
    case "\n\t":
        v2 -= (v2 / v2)
    default:
        if true {
            v2 = v2
        }
    }
    v2 = (v2 - v2)
    v25 := GenS1{F0: uint64(v1)}
    v26 := GenS3{F0: -603}
    var v27 int16 = -7
    v27 = v27
    if false {
        v2 += (1.5 * v2)
        var v28 int16 = (-(v27))
        var v29 uint8 = 2
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
