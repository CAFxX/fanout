package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int8 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint64; F1 bool; F2 int }

func main() {
    var v1 int16 = 641
    v2 := []string{("hello" + ("foo bar" + ("" + (((("hello" + ("\n\t" + ("\n\t" + "hello"))) + ("a" + "hello")) + ("a" + "foo bar")) + "hello")))), "foo bar", ("a" + "\n\t")}
    var v3 uint64 = 4294967295
    var v4 int16 = (v1 & (v1 + v1))
    if false {
        if ((float64(float32(-2.25)) > 0.1) || false) {
            v3 -= (v3 + (v3 >> ((1) & 63)))
            v4 -= v1
        }
        switch -1 {
        case 32767:
            v5 := GenS1{F0: 127, F1: false, F2: -34}
            _ = v5
        }
    }
    switch 759 {
    case 35:
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    default:
        v6 := GenS2{F0: 65535, F1: 4294967295}
        _ = v6
    }
    v4 += v1
    v7 := GenS3{F0: v3, F1: false, F2: 0}
    var v8 int8 = -106
    v3 -= 0
    switch 534 {
    case -91:
        if len(v2) > 0 {
            v2[0] = "a"
        }
    case -652:
        v4 += 458
    }
    v1 = (v4 >> ((2) & 63))
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(v1, v1)
        var v11 uint8 = 136
        _ = fi10
        _ = v11
    }
    if false {
        switch -309 {
        case -998:
            v1 = (v1 | 32767)
        default:
            var v12 int = 128
            _ = v12
        }
        fmt.Println(v4, v9, v9)
    }
    var v13 int = 1
    for fi14 := 0; fi14 < 15; fi14++ {
        v4 = -1
        _ = fi14
    }
    switch (v13 & v13) {
    case 1:
        fmt.Println(v13, v1)
    }
    if false {
        for fi15 := 0; fi15 < 15; fi15++ {
            v9 = ("a" + "foo bar")
            v9 = v9
            _ = fi15
        }
        if true {
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 float64 = 0.0
                _ = fi16
                _ = v17
            }
        }
        _ = len(v2)
    } else {
        v18 := GenS1{F0: 255, F1: true, F2: -105}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
