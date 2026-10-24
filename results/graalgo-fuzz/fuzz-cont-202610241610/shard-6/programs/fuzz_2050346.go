package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 uint16, p1 int8) uint {
    p1 = (p1 % (((p1 ^ 0)) | 1))
    v7 := []int{751, 127}
    var v8 bool = true
    var v9 int = 782
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 uint32, p1 string, p2 uint16) string {
    p2 = 190
    v10 := []int{}
    var v11 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return ("a" + (p1 + "foo bar"))
}

func main() {
    var v1 bool = false
    var v2 int32 = 534
    v3 := []string{"foo bar", "hello", "foo bar", "a"}
    var v4 uint32 = 961
    var v5 int = 127
    var v6 float32 = -2.25
    var v12 int16 = 243
    switch "a" {
    case "foo bar":
        var v13 int32 = (v2 * 2147483647)
        _ = v13
    case "":
        fmt.Println(v2)
    case "a":
        v12 += (-(v12))
    }
    v1 = ((-(v2)) < v2)
    switch (v5 + v5) {
    case 1:
        for fi14 := 0; fi14 < 7; fi14++ {
            for fi15 := 0; fi15 < 15; fi15++ {
                v16 := GenS1{F0: (v4 % ((v4) | 1))}
                _ = fi15
                _ = v16
            }
            v17 := GenS2{F0: "\n\t", F1: 9223372036854775807}
            fmt.Println(v4, fi14)
            _ = fi14
            _ = v17
        }
    }
    switch "hello" {
    case "hello":
        switch "" {
        case "hello":
            if v1 {
                v18 := GenS2{F0: (("a" + "foo bar") + ("a" + "hello")), F1: 42}
                v4 = (v4 + (v4 & v4))
                _ = v18
            } else {
                v4 = (v4 & v4)
                var v19 int64 = -9223372036854775808
                switch (v5 ^ (v5 % (((v5 >> ((0) & 63))) | 1))) {
                case 32767:
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                case 32768:
                    if (v1 || (-2147483648 > v2)) {
                        for fi20 := 0; fi20 < 20; fi20++ {
                            v12 += -318
                            _ = fi20
                        }
                    }
                default:
                    v3 = append(v3, (((("a" + "hello") + ("hello" + "\n\t")) + "hello") + ""))
                }
                _ = v19
            }
        case "a":
            for fi21 := 0; fi21 < 18; fi21++ {
                v22 := GenS2{F0: "", F1: 9223372036854775807}
                _ = fi21
                _ = v22
            }
        case "foo bar":
            v23 := GenS2{F0: "", F1: -956}
            _ = v23
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v24 uint32 = v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v25 float64 = 1e10
    var v26 uint64 = 88
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
