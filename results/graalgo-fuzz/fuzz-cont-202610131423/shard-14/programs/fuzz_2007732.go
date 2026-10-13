package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int32; F1 uint64; F2 int16 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int16; F1 uint8 }
func genFn0() string {
    v4 := []int{-2147483648, 9223372036854775807, -9223372036854775808}
    if (648 != 755) {
        if len(v4) > 0 {
            v4[0] = 32768
        }
        v5 := GenS1{F0: 100, F1: 255, F2: -128}
        v6 := GenS2{F0: 0, F1: 2, F2: 61}
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 string = (("foo bar" + "hello") + ("hello" + "\n\t"))
            v8 = "a"
            if false {
                v8 = (("\n\t" + (v8 + v8)) + v8)
                v9 := GenS3{F0: 42, F1: 1, F2: 0}
                _ = len(v4)
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        switch ("a" + ("a" + "foo bar")) {
        case "":
            var v10 bool = (-2.25 < -2.25)
            _ = v10
        case "foo bar":
            var v11 int8 = 100
            _ = v11
        case "a":
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 float32 = 0.0
                _ = fi12
                _ = v13
            }
        }
        switch 1 {
        case -891:
            v14 := GenS1{F0: 37, F1: 72, F2: 42}
            _ = v14
        default:
            if (((-2.25 / 0.1) + 1e10) < 1e10) {
                var v15 uint = 1
                _ = v15
            }
        }
    }
    switch 9223372036854775807 {
    case -1:
        v16 := GenS4{F0: (!(-2.25 != 1e10)), F1: 32767}
        _ = v16
    }
    _ = v4
    return "foo bar"
}

func main() {
    var v1 int16 = 0
    var v2 int16 = v1
    v3 := make(map[int]bool)
    if (true && false) {
        for fi17 := 0; fi17 < 1; fi17++ {
            fmt.Println(v1, v1, v2)
            v1 = (v2 % ((v1) | 1))
            _ = fi17
        }
    }
    v1 -= v1
    var v18 uint32 = 2147483648
    v19 := []int{128, 9223372036854775807, -1}
    switch ("\n\t" + "a") {
    case "hello":
        var v20 uint64 = 2
        _ = v20
    case "a":
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 int16 = (v2 % (((v1 / ((v1) | 1))) | 1))
            var v23 int16 = v22
            _ = fi21
            _ = v22
            _ = v23
        }
    default:
        var v24 uint8 = 114
        _ = v24
    }
    v2 = (v1 / (((v2 * -532)) | 1))
    var v25 int8 = 127
    v25 = (v25 << ((2) & 63))
    var v26 uint = 4294967295
    if false {
        var v27 int8 = 27
        var v28 uint64 = 63
        v29 := GenS2{F0: 396, F1: v18, F2: (v18 / ((2147483648) | 1))}
        _ = v27
        _ = v28
        _ = v29
    } else {
        v30 := GenS3{F0: 0, F1: 18446744073709551615, F2: (v2 ^ v1)}
        for fi31 := 0; fi31 < 4; fi31++ {
            v32 := GenS5{F0: 42, F1: 127}
            var v33 int64 = 0
            _ = fi31
            _ = v32
            _ = v33
        }
        v2 -= 0
        _ = v30
    }
    switch "foo bar" {
    case "\n\t":
        if len(v19) > 0 {
            v19[0] = int(v26)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
