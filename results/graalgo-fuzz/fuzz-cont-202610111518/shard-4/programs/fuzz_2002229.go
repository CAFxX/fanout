package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int8 }
type GenS2 struct { F0 bool; F1 uint8; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int }
func genFn0(p0 string, p1 int8) int {
    fmt.Println(p0)
    v4 := GenS1{F0: 23, F1: true, F2: 0}
    if false {
        p1 = p1
        for fi5 := 0; fi5 < 7; fi5++ {
            if (true && false) {
                p0 = p0
                p1 = p1
                v6 := GenS2{F0: (65535 >= 2), F1: 2, F2: 9223372036854775807}
                _ = v6
            }
            v7 := []int{}
            if true {
                fmt.Println(fi5)
                var v8 string = "foo bar"
                _ = v8
            } else {
                v9 := GenS2{F0: false, F1: 42, F2: -653}
                for fi10 := 0; fi10 < 4; fi10++ {
                    switch p0 {
                    case "foo bar":
                        if ((-2.25 + (0.1 - 0.0)) <= (0.0 - -2.25)) {
                            switch (p0 + ((p0 + (p0 + p0)) + ((p0 + p0) + "\n\t"))) {
                            case "":
                                var v11 string = "foo bar"
                                _ = v11
                            case "\n\t":
                                fmt.Println(fi10, fi5, fi5)
                            case "a":
                                if (42 < -741) {
                                    v12 := GenS3{F0: p0}
                                    var v13 float64 = 3.14159
                                    _ = v12
                                    _ = v13
                                } else {
                                    v14 := GenS4{F0: p0, F1: (p0 + p0)}
                                    _ = v14
                                }
                            }
                            p0 = "hello"
                        }
                    case "\n\t":
                        var v15 float64 = 1e10
                        _ = v15
                    }
                    _ = fi10
                }
                _ = v9
            }
            _ = fi5
            _ = v7
        }
        switch -305 {
        case -9223372036854775808:
            p1 = p1
        }
    } else {
        v16 := []int{-9223372036854775808, 42, 32768, 365}
        _ = v16
    }
    var v17 uint16 = 0
    _ = p0
    _ = p1
    _ = v17
    _ = v4
    return -741
}
func genFn1() uint64 {
    switch "a" {
    case "a":
        v18 := GenS5{F0: genFn0("a", 1)}
        _ = v18
    case "hello":
        var v19 float32 = 1e10
        _ = v19
    default:
        var v20 int = genFn0(("" + (("\n\t" + "foo bar") + "\n\t")), 42)
        _ = v20
    }
    var v21 float64 = 0.1
    _ = v21
    return 678
}

func main() {
    var v1 uint16 = 0
    var v2 bool = (9223372036854775807 <= -1)
    var v3 int8 = -1
    v3 += (v3 % (((v3 << ((933) & 63))) | 1))
    var v22 float64 = 1e10
    v3 -= (v3 ^ (v3 | v3))
    v3 = 1
    var v23 int32 = -223
    v24 := []int{genFn0(((("foo bar" + "hello") + ((("a" + "hello") + "a") + "hello")) + ("\n\t" + "")), (v3 + v3))}
    if false {
        if v2 {
            var v25 int = 9223372036854775807
            _ = v25
        } else {
            _ = len(v24)
            switch 165 {
            case 0:
                var v26 int8 = -128
                _ = v26
            case 128:
                v23 = v23
            default:
                fmt.Println(v22)
            }
        }
        v27 := GenS4{F0: "a", F1: ""}
        var v28 bool = false
        _ = v27
        _ = v28
    }
    fmt.Println(v22, v1, v22)
    var v29 float64 = (v22 - v22)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
}
