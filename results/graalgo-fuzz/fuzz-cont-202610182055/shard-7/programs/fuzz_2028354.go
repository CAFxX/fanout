package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int32 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint8; F1 uint16; F2 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 3
    var v3 string = ("a" + "hello")
    fmt.Println(v2, v2)
    if (-57 <= 1) {
        v4 := GenS1{F0: 32768, F1: 65535, F2: 2147483647}
        _ = v4
    } else {
        v5 := GenS2{F0: 1, F1: false}
        var v6 int8 = -1
        _ = v5
        _ = v6
    }
    var v7 float64 = 3.14159
    fmt.Println(v2, v2)
    v3 = v3
    switch 128 {
    case 738:
        fmt.Println(v7, v7)
    case 4294967295:
        v2 = v2
    default:
        v2 = (v2 & (v2 << ((855) & 63)))
    }
    var v8 int8 = 1
    for fi9 := 0; fi9 < 4; fi9++ {
        switch (v3 + (((((("foo bar" + ((v3 + v3) + "foo bar")) + v3) + ("a" + "hello")) + v3) + v3) + "foo bar")) {
        case "a":
            v10 := GenS1{F0: (fi9 + fi9), F1: 42, F2: 42}
            _ = v10
        default:
            for fi11 := 0; fi11 < 11; fi11++ {
                fmt.Println(fi11, v3, fi9)
                _ = fi11
            }
        }
        var v12 uint64 = 0
        _ = fi9
        _ = v12
    }
    var v13 int16 = v2
    v3 = v3
    if true {
        v1[(v3 + v3)] = (v3 > (v3 + ("foo bar" + "\n\t")))
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 bool = (65535 >= 42)
            var v17 bool = (42 != 4294967295)
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := []int{(fi14 % (((fi14 | (^(fi14)))) | 1)), 1, 32767}
        _ = fi14
        _ = v18
    }
    var v19 uint32 = 4294967295
    v1[("foo bar" + v3)] = (2 > 1)
    fmt.Println(v8)
    _ = len(v1)
    for fi20 := 0; fi20 < 16; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS3{F0: 127, F1: 2, F2: v8}
            var v23 uint8 = 128
            fmt.Println(fi20, v19, fi21)
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 int = 39
        switch (((v3 + (v3 + ("foo bar" + "hello"))) + ((v3 + v3) + ("foo bar" + v3))) + v3) {
        case "foo bar":
            for fi25 := 0; fi25 < 1; fi25++ {
                v26 := GenS1{F0: fi25, F1: v19, F2: -1}
                if (true || (true && true)) {
                    fmt.Println(v3, fi25, fi20)
                    var v27 int = fi20
                    var v28 float32 = 0.0
                    v3 = ((v3 + v3) + (v3 + (v3 + v3)))
                    if false {
                        _ = len(v1)
                        v29 := v1[(v3 + v3)]
                        _ = v29
                    }
                    _ = v27
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
        case "hello":
            for fi30 := 0; fi30 < 17; fi30++ {
                fmt.Println(v13)
                delete(v1, "hello")
                var v31 int64 = 42
                _ = fi30
                _ = v31
            }
        default:
            switch (fi20 - v24) {
            case 127:
                var v32 int32 = 0
                _ = v32
            case 636:
                v8 = v8
            case 128:
                v3 = (v3 + "hello")
            default:
                var v33 float64 = v7
                _ = v33
            }
        }
        _ = fi20
        _ = v24
    }
    switch 1 {
    case -9223372036854775808:
        var v34 int32 = 789
        _ = v34
    case 0:
        v3 = ("a" + "foo bar")
    default:
        v1[v3] = false
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
