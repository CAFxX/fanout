package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 int16; F1 int32; F2 int32 }

func main() {
    var v1 int = -107
    v2 := make(map[string]string)
    v3 := []uint32{307, 151}
    var v4 bool = (-571 == 42)
    v5 := GenS1{F0: 127, F1: -32768}
    var v6 float32 = 3.14159
    if (((v6 + ((1.5 - v6) / (v6 * v6))) <= v6) || ("a" >= "foo bar")) {
        for fi7 := 0; fi7 < 3; fi7++ {
            delete(v2, "hello")
            _ = fi7
        }
        var v8 int64 = -406
        _ = v8
    }
    if v4 {
        _ = len(v2)
        v4 = v4
        if v4 {
            fmt.Println(v1, v1, v6)
            v6 += (0.1 - float32(1e10))
        } else {
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 int = (-(fi9))
                fmt.Println(v1)
                _ = fi9
                _ = v10
            }
            _ = len(v2)
        }
    } else {
        switch "foo bar" {
        case "":
            fmt.Println(v1)
        }
    }
    v11 := GenS2{F0: 42, F1: "\n\t"}
    var v12 uint16 = 0
    v12 = (v12 ^ v12)
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint = 332
        var v15 float64 = float64(1.5)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 bool = (v4 && ((v4 || (!(1 < 127))) && v4))
    if (!v16) {
        v16 = (!(-759 <= 1))
        v6 -= 1.5
    } else {
        v1 -= v1
        if (("hello" + (("a" + ("hello" + "")) + ((("hello" + ("foo bar" + (("foo bar" + "\n\t") + "\n\t"))) + ("" + ((("" + "") + ((("hello" + ("a" + ("foo bar" + "foo bar"))) + (("foo bar" + ("hello" + "foo bar")) + ("a" + ((("hello" + ("" + "\n\t")) + "\n\t") + "\n\t")))) + "\n\t")) + "hello"))) + ("\n\t" + "\n\t")))) >= "") {
            v17 := GenS3{F0: 65535, F1: 1, F2: 0}
            var v18 uint = 1
            _ = v17
            _ = v18
        }
        if true {
            var v19 bool = v4
            var v20 int = (v1 | v1)
            var v21 float32 = (((v6 - (0.1 / float32(1.5))) * (((-2.25 / v6) * ((0.1 / -2.25) - v6)) / v6)) + v6)
            _ = v19
            _ = v20
            _ = v21
        } else {
            v22 := GenS4{F0: 480, F1: v16}
            if (!true) {
                fmt.Println(v12)
            }
            switch "\n\t" {
            case "hello":
                v23 := GenS5{F0: -220, F1: 42}
                _ = v23
            default:
                v1 = (v1 / (((v1 % ((v1) | 1))) | 1))
            }
            _ = v22
        }
    }
    var v24 int64 = 0
    switch (-(v1)) {
    case 127:
        v3 = append(v3, 2)
    case -1:
        v25 := v2[(((("a" + ((("\n\t" + "foo bar") + "\n\t") + ("" + ""))) + "a") + ("" + "\n\t")) + "foo bar")]
        _ = v25
    default:
        v26 := GenS6{F0: 1000, F1: 2147483647, F2: -2147483648}
        _ = v26
    }
    v27 := GenS6{F0: -32768, F1: -684, F2: 1}
    switch v1 {
    case 827:
        fmt.Println(v24, v16, v24)
    case -1:
        var v28 uint16 = v12
        _ = v28
    default:
        for fi29 := 0; fi29 < 4; fi29++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v30 float64 = (1.5 - 1e10)
            _ = len(v2)
            _ = fi29
            _ = v30
        }
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
