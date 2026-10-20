package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 string; F1 int8; F2 int64 }
type GenS3 struct { F0 int64; F1 int; F2 string }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint) string {
    switch 9223372036854775807 {
    case 128:
        v4 := GenS1{F0: 255, F1: (("hello" + "\n\t") + "")}
        _ = v4
    case 127:
        var v5 float32 = ((-2.25 + ((0.0 * float32((0.0 / 1.5))) * (1.5 - 1.5))) / (float32(((((0.0 + 1.5) - 0.1) / -2.25) - float64(-2.25))) / 1.5))
        _ = v5
    default:
        if false {
            switch "a" {
            case "\n\t":
                if false {
                    v6 := []int{1, 775}
                    _ = v6
                }
            case "foo bar":
                v7 := GenS1{F0: 2, F1: ((((((("foo bar" + ("hello" + "hello")) + "a") + "\n\t") + ("foo bar" + ("a" + "a"))) + "hello") + "a") + (("hello" + "hello") + ((("" + ("foo bar" + ("\n\t" + "foo bar"))) + "\n\t") + "hello")))}
                _ = v7
            case "a":
                if false {
                    var v8 uint16 = 65535
                    for fi9 := 0; fi9 < 10; fi9++ {
                        p0 = (^(p0))
                        var v10 uint32 = 366
                        _ = fi9
                        _ = v10
                    }
                    v11 := GenS1{F0: 42, F1: ("" + "hello")}
                    _ = v11
                    _ = v8
                }
            }
            var v12 int16 = -32768
            _ = v12
        }
    }
    if true {
        switch 520 {
        case 42:
            v13 := GenS2{F0: ((("\n\t" + ("a" + "a")) + "") + "hello"), F1: 100, F2: -1}
            _ = v13
        case 127:
            var v14 uint = (p0 | 42)
            _ = v14
        }
        v15 := []int{128}
        _ = v15
    }
    v16 := []int{9223372036854775807, 42}
    p0 -= 2
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    _ = p0
    _ = v16
    return "hello"
}

func main() {
    var v1 uint = 1
    var v2 int64 = 1
    var v3 float32 = -2.25
    var v17 int16 = 672
    switch "foo bar" {
    case "foo bar":
        v18 := GenS3{F0: v2, F1: -675, F2: "\n\t"}
        _ = v18
    default:
        v19 := []int{-365}
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v1 -= (v1 | (v1 * v1))
        var v21 string = ("" + "foo bar")
        var v22 uint64 = 565
        _ = fi20
        _ = v21
        _ = v22
    }
    switch "foo bar" {
    case "hello":
        switch 2147483647 {
        case 2147483647:
            v23 := []int{-1}
            _ = v23
        case -1:
            var v24 uint = 657
            _ = v24
        case -350:
            v25 := []int{32767}
            _ = v25
        default:
            switch -547 {
            case -230:
                v26 := GenS4{F0: (("" + "") + (("foo bar" + "a") + "\n\t")), F1: 9223372036854775807}
                _ = v26
            }
        }
    }
    fmt.Println(v3)
    var v27 uint8 = 255
    v17 -= (v17 + (-(v17)))
    v28 := GenS5{F0: 2, F1: 4294967295}
    fmt.Println(v27, v2, v17)
    v29 := GenS2{F0: "foo bar", F1: 127, F2: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
