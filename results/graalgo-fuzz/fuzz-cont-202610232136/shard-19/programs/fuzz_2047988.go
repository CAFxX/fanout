package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint }
type GenS5 struct { F0 int8 }

func main() {
    var v1 uint32 = 4294967295
    var v2 float32 = (1e10 / ((-2.25 / (0.1 * 1.5)) + ((-2.25 + (1.5 * -2.25)) * (float32((1.5 / 1.0)) + 1.5))))
    v3 := make(map[string]string)
    var v4 uint8 = 1
    var v5 uint64 = 42
    v6 := GenS1{F0: v1, F1: 0}
    switch 9223372036854775807 {
    case -1:
        if ("\n\t" < "foo bar") {
            if (true || (false && false)) {
                v7 := GenS1{F0: v1, F1: 9223372036854775807}
                if (-1 != -150) {
                    var v8 uint16 = 975
                    var v9 string = ""
                    for fi10 := 0; fi10 < 1; fi10++ {
                        switch ("" + v9) {
                        case "hello":
                            var v11 int = (fi10 << ((42) & 63))
                            _ = v11
                        }
                        _ = fi10
                    }
                    _ = v8
                    _ = v9
                }
                _ = v7
            } else {
                var v12 uint16 = 607
                v13 := GenS2{F0: -1, F1: (v1 << ((482) & 63))}
                _ = v12
                _ = v13
            }
            v14 := []int{127, -2147483648, 9223372036854775807}
            _ = v14
        } else {
            var v15 int64 = -9223372036854775808
            var v16 int8 = 42
            v17 := GenS3{F0: 1, F1: v16}
            _ = v15
            _ = v16
            _ = v17
        }
    case 1:
        v5 -= v5
    }
    v5 = v5
    var v18 int16 = -32768
    v3[("hello" + "")] = "hello"
    v3[("" + "a")] = ("" + "")
    var v19 int16 = (v18 & (v18 << ((255) & 63)))
    switch "a" {
    case "":
        if true {
            fmt.Println(v2, v18)
        } else {
            v20 := GenS3{F0: 150, F1: 1}
            _ = v20
        }
    case "hello":
        if false {
            v21 := GenS1{F0: 42, F1: 0}
            if (false || ((false && false) || true)) {
                v22 := GenS4{F0: 65535, F1: (0 < -9223372036854775808), F2: 18446744073709551615}
                _ = v22
            }
            _ = v21
        }
    case "foo bar":
        fmt.Println(v5)
    default:
        var v23 int8 = 51
        _ = v23
    }
    switch 318 {
    case 2147483647:
        v24 := GenS5{F0: 117}
        _ = v24
    case -9223372036854775808:
        fmt.Println(v2, v2)
    default:
        v25 := GenS1{F0: (v1 ^ 754), F1: 42}
        _ = v25
    }
    var v26 string = (("\n\t" + "") + (("\n\t" + "") + "a"))
    for fi27 := 0; fi27 < 1; fi27++ {
        delete(v3, ((((v26 + "\n\t") + ((v26 + v26) + (v26 + v26))) + ((((("" + ((((v26 + ("hello" + (v26 + "a"))) + ((v26 + (("hello" + ("\n\t" + (("\n\t" + v26) + ((v26 + "hello") + (("a" + v26) + ""))))) + v26)) + "a")) + v26) + v26)) + v26) + v26) + "") + (v26 + ("a" + ("" + ("foo bar" + v26)))))) + v26))
        _ = fi27
    }
    if false {
        var v28 uint32 = (v1 << ((uint(v4)) & 63))
        var v29 int64 = 300
        fmt.Println(v1)
        _ = v28
        _ = v29
    } else {
        var v30 bool = (false || (!(v1 >= 0)))
        v31 := GenS1{F0: 0, F1: int(v5)}
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
