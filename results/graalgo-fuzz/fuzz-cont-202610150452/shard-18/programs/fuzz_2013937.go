package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint16) string {
    p0 -= 792
    var v5 int64 = 1
    var v6 int8 = 1
    fmt.Println(v5)
    v5 = (v5 ^ (v5 | v5))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return ""
}
func genFn1(p0 uint16, p1 int64) int {
    switch "a" {
    case "a":
        v7 := []int{}
        _ = v7
    case "hello":
        p1 = p1
    case "":
        var v8 uint32 = 42
        _ = v8
    }
    p1 += (p1 | p1)
    p1 -= (p1 ^ p1)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]string)
    var v3 uint16 = 1
    var v4 int32 = -2147483648
    var v9 float64 = 0.0
    if true {
        v10 := GenS1{F0: (v3 | 32767), F1: -259}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(v3, v4, v3)
            _ = fi11
        }
        v9 -= 0.1
        delete(v2, ("hello" + "\n\t"))
    }
    v2["hello"] = ("a" + (("\n\t" + (("hello" + "a") + "hello")) + ""))
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 string = "a"
        fmt.Println(v9, fi12, v4)
        var v14 int16 = -32768
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: 0}
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS1{F0: (v3 | v3), F1: 357}
        _ = fi16
        _ = v17
    }
    delete(v2, "a")
    switch (("\n\t" + "\n\t") + "foo bar") {
    case "foo bar":
        fmt.Println(v4, v3, v4)
    default:
        if true {
            switch 127 {
            case -782:
                if true {
                    v18 := []int{70, -1, 127}
                    _ = v18
                } else {
                    var v19 string = "a"
                    v1 = v1
                    var v20 bool = (0 >= 18446744073709551615)
                    _ = v19
                    _ = v20
                }
            case 127:
                if (true || true) {
                    if false {
                        for fi21 := 0; fi21 < 12; fi21++ {
                            switch "a" {
                            case "foo bar":
                                v3 += (v3 - 42)
                            case "\n\t":
                                _ = len(v2)
                            case "":
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    v23 := GenS1{F0: 633, F1: fi21}
                                    var v24 int16 = -1
                                    var v25 float64 = (v9 - v9)
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                            default:
                                v1 -= 0.1
                            }
                            fmt.Println(fi21)
                            v4 += v4
                            _ = fi21
                        }
                        var v26 string = "a"
                        _ = v26
                    }
                } else {
                    switch -17 {
                    case 1:
                        v4 += v4
                    case -3:
                        var v27 int = genFn1(789, 42)
                        _ = v27
                    case -812:
                        v9 = ((v9 - (3.14159 + v9)) * (v9 + v9))
                    default:
                        fmt.Println(v1)
                    }
                }
            default:
                var v28 int64 = 9223372036854775807
                _ = v28
            }
            fmt.Println(v3, v3, v3)
            fmt.Println(v3)
        }
    }
    v3 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
