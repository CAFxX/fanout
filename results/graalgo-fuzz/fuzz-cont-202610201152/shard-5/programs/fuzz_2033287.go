package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 string, p1 bool) bool {
    p1 = p1
    p0 = ("" + p0)
    p1 = p1
    p0 = p0
    _ = p0
    _ = p1
    return p1
}
func genFn1(p0 int32, p1 int16) uint32 {
    p1 += 64
    if false {
        var v5 int16 = (^(p1))
        if true {
            for fi6 := 0; fi6 < 2; fi6++ {
                v7 := []int{876, -285, 32767, fi6}
                _ = len(v7)
                switch int(p1) {
                case -2147483648:
                    v8 := GenS1{F0: -1, F1: -128, F2: 128}
                    _ = v8
                default:
                    if false {
                        fmt.Println(p1)
                        var v9 string = ""
                        v9 = ""
                        _ = v9
                    } else {
                        fmt.Println(fi6)
                        p0 = p0
                    }
                }
                _ = fi6
                _ = v7
            }
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := []int{(fi10 + -742), int(p0), -1, 4294967295}
            for fi12 := 0; fi12 < 17; fi12++ {
                v5 = (p1 & v5)
                _ = fi12
            }
            v13 := GenS2{F0: 638}
            _ = fi10
            _ = v11
            _ = v13
        }
        _ = v5
    } else {
        v14 := GenS3{F0: -68, F1: 65535}
        _ = v14
    }
    _ = p0
    _ = p1
    return 2147483648
}

func main() {
    var v1 float64 = 0.0
    v2 := []string{"", (("\n\t" + ("foo bar" + ("\n\t" + "foo bar"))) + "hello")}
    var v3 int8 = 127
    var v4 int8 = v3
    v15 := GenS4{F0: 0}
    v16 := GenS3{F0: 42, F1: 1}
    if false {
        v4 += (v3 / ((-102) | 1))
    }
    var v17 bool = false
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = (("a" + "\n\t") + "")
    }
    fmt.Println(v4, v17)
    switch "" {
    case "a":
        v17 = v17
    case "":
        var v18 uint8 = 128
        _ = v18
    case "foo bar":
        v19 := GenS2{F0: 42}
        _ = v19
    default:
        v2 = append(v2, ((("foo bar" + "a") + (("" + ("\n\t" + "")) + "hello")) + (("\n\t" + (("hello" + "hello") + "a")) + "")))
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 bool = true
        _ = len(v2)
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: 205, F1: 100, F2: 2}
    v4 = (v4 - 127)
    var v23 int64 = 0
    var v24 bool = v17
    var v25 uint8 = 181
    var v26 uint8 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
