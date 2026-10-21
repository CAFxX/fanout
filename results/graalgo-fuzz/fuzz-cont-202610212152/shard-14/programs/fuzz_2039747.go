package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
type GenS4 struct { F0 string; F1 uint8 }

func main() {
    v1 := []string{"hello", ("a" + ("" + "a")), "\n\t"}
    var v2 bool = (697 <= 1)
    var v3 uint8 = 255
    var v4 bool = ((((v2 || (423 != -2147483648)) || v2) || v2) && (("foo bar" + "foo bar") >= (("foo bar" + "a") + "a")))
    if false {
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 int8 = 100
            _ = fi5
            _ = v6
        }
        v4 = true
    }
    v3 += (v3 >> ((755) & 63))
    var v7 int8 = 1
    var v8 uint8 = v3
    switch int(v3) {
    case -306:
        v9 := GenS1{F0: (v3 ^ (v8 & 128)), F1: true}
        _ = v9
    case -336:
        v7 = v7
    case -9223372036854775808:
        for fi10 := 0; fi10 < 10; fi10++ {
            v3 = (^(v3))
            _ = fi10
        }
    }
    v4 = true
    if true {
        switch (("foo bar" + "hello") + (("a" + ("\n\t" + "a")) + "foo bar")) {
        case "":
            var v11 float32 = ((0.1 * 0.1) + 0.0)
            _ = v11
        default:
            for fi12 := 0; fi12 < 20; fi12++ {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                v13 := GenS1{F0: v8, F1: v4}
                v7 = 100
                _ = fi12
                _ = v13
            }
        }
        v14 := GenS2{F0: -15}
        var v15 int16 = 42
        _ = v14
        _ = v15
    } else {
        if v2 {
            v16 := GenS1{F0: 42, F1: (true || v4)}
            v8 = 1
            _ = v16
        } else {
            v7 += (v7 >> ((2) & 63))
            fmt.Println(v4, v2)
        }
        v7 = v7
        v17 := GenS2{F0: 1}
        _ = v17
    }
    switch "a" {
    case "hello":
        for fi18 := 0; fi18 < 18; fi18++ {
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            _ = fi18
        }
    case "\n\t":
        fmt.Println(v3, v7, v3)
    default:
        v2 = false
    }
    if v2 {
        v19 := GenS3{F0: 417, F1: 0, F2: -2147483648}
        for fi20 := 0; fi20 < 6; fi20++ {
            switch fi20 {
            case 42:
                var v21 string = "\n\t"
                _ = v21
            default:
                var v22 float64 = (3.14159 / ((3.14159 * float64(-2.25)) - ((3.14159 * ((-2.25 / 1.0) + -2.25)) * (1.5 - 1.5))))
                _ = v22
            }
            v23 := GenS4{F0: ("" + ("foo bar" + ("" + "a"))), F1: (v3 + (v8 / (((v3 >> ((65535) & 63))) | 1)))}
            _ = fi20
            _ = v23
        }
        var v24 float64 = -2.25
        _ = v19
        _ = v24
    }
    var v25 int = -683
    v4 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
