package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 0
    v3 := make(map[string]bool)
    v4 := []int{128, 0, 2147483647}
    switch "\n\t" {
    case "hello":
        v5 := GenS1{F0: 1, F1: 1, F2: 42}
        _ = v5
    case "":
        var v6 uint64 = 42
        _ = v6
    default:
        v7 := GenS2{F0: 208, F1: 65535}
        _ = v7
    }
    v2 = 255
    var v8 uint64 = 897
    var v9 float32 = 0.1
    v2 = v2
    for fi10 := 0; fi10 < 5; fi10++ {
        v9 = (1.5 * -2.25)
        if true {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int8 = -53
                var v13 uint64 = (^(v8))
                _ = fi11
                _ = v12
                _ = v13
            }
            switch ((((("hello" + ("\n\t" + "")) + "") + "hello") + ("hello" + ("a" + "hello"))) + "") {
            case "a":
                v8 -= v8
            case "foo bar":
                fmt.Println(fi10, fi10, v8)
            case "hello":
                v3[""] = (65535 != 2147483647)
            default:
                var v14 float32 = (0.0 + (3.14159 / v9))
                _ = v14
            }
        }
        _ = fi10
    }
    if len(v4) > 0 {
        v4[0] = 32767
    }
    var v15 int16 = 32767
    v15 += v15
    if len(v4) > 0 {
        v4[0] = 127
    }
    if (1 != -1) {
        for fi16 := 0; fi16 < 4; fi16++ {
            fmt.Println(v9, v8)
            switch ("a" + ("a" + "a")) {
            case "a":
                for fi17 := 0; fi17 < 7; fi17++ {
                    v2 += v2
                    v18 := GenS3{F0: false, F1: 0}
                    _ = fi17
                    _ = v18
                }
            case "hello":
                var v19 uint16 = (v2 & 1)
                _ = v19
            case "":
                for fi20 := 0; fi20 < 8; fi20++ {
                    v2 += (v2 | (v2 + v2))
                    v9 -= float32(-2.25)
                    _ = fi20
                }
            }
            v15 = v15
            _ = fi16
        }
        switch (("foo bar" + ("\n\t" + "")) + "") {
        case "\n\t":
            fmt.Println(v9, v8)
        case "hello":
            if true {
                var v21 int64 = -1
                _ = v21
            } else {
                v8 = (v8 | (^(v8)))
                var v22 string = ("hello" + ("hello" + "foo bar"))
                v4 = append(v4, 2147483647)
                _ = v22
            }
        default:
            delete(v3, ("hello" + ("foo bar" + "hello")))
        }
        v23 := GenS3{F0: ((float32(1.5) - v9) == v9), F1: 773}
        _ = v23
    } else {
        if false {
            v24 := GenS2{F0: 127, F1: 2147483648}
            _ = v24
        } else {
            v25 := GenS4{F0: 1}
            fmt.Println(v2, v2)
            v26 := GenS4{F0: -2147483648}
            _ = v25
            _ = v26
        }
    }
    if false {
        if (42 > 18446744073709551615) {
            var v27 bool = (1 != -51)
            fmt.Println(v2, v8)
            _ = v27
        } else {
            v28 := GenS3{F0: true, F1: uint(v2)}
            delete(v1, "hello")
            _ = v28
        }
    } else {
        if true {
            var v29 float32 = float32(0.1)
            v2 = 148
            _ = v29
        } else {
            var v30 uint16 = v2
            _ = v30
        }
        v8 -= v8
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
