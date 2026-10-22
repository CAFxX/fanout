package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
func genFn0() string {
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := []int{(-(fi5)), -2147483648, fi5, fi5}
        _ = fi5
        _ = v6
    }
    if ((!(("\n\t" > "a") || (false || (true && false)))) && false) {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 float64 = (-2.25 / 1e10)
            fmt.Println(v8, fi7)
            _ = fi7
            _ = v8
        }
    }
    switch (("" + (("hello" + "\n\t") + "foo bar")) + (("\n\t" + "\n\t") + ("" + "\n\t"))) {
    case "":
        var v9 uint8 = 255
        _ = v9
    case "hello":
        v10 := []int{9223372036854775807, -2147483648}
        _ = v10
    case "\n\t":
        var v11 float32 = float32(1.5)
        _ = v11
    default:
        var v12 uint64 = 0
        _ = v12
    }
    return (("\n\t" + ((("hello" + "\n\t") + "\n\t") + ("hello" + ""))) + (("a" + ("" + ("\n\t" + "\n\t"))) + "foo bar"))
}
func genFn1(p0 uint8, p1 uint, p2 string) bool {
    if false {
        p0 = 0
    }
    v13 := []int{127, -9223372036854775808, 128}
    var v14 bool = (((-2.25 / (float64((3.14159 + (1.5 - 1.5))) - float64(-2.25))) * 0.1) != 1.5)
    for fi15 := 0; fi15 < 8; fi15++ {
        switch "\n\t" {
        case "":
            fmt.Println(p0)
        }
        p2 = "\n\t"
        if ((0 == 377) && false) {
            for fi16 := 0; fi16 < 19; fi16++ {
                switch 575 {
                case -847:
                    var v17 uint8 = (p0 >> (((p1 + 608)) & 63))
                    switch "hello" {
                    case "":
                        if v14 {
                            var v18 uint32 = 65535
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            var v19 uint16 = 65535
                            fmt.Println(v14, v18, v19)
                            var v20 float32 = 0.1
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    case "a":
                        p0 = 2
                    default:
                        var v21 int16 = 42
                        _ = v21
                    }
                    var v22 uint16 = 1
                    _ = v17
                    _ = v22
                case 4294967295:
                    v23 := GenS1{F0: 958, F1: 1}
                    _ = v23
                default:
                    v24 := GenS1{F0: 0, F1: uint32(p0)}
                    _ = v24
                }
                switch "foo bar" {
                case "hello":
                    v14 = true
                case "a":
                    fmt.Println(fi15, p0, p1)
                default:
                    var v25 float32 = (1e10 + 1e10)
                    _ = v25
                }
                _ = fi16
            }
            var v26 int8 = 1
            var v27 float64 = (1e10 + 1e10)
            _ = v26
            _ = v27
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return v14
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 0
    var v3 uint = 269
    var v4 float64 = 0.1
    v1[65535] = "hello"
    v28 := []int{2147483647, 9223372036854775807, 2147483647, -1}
    v4 += 1.5
    var v29 int = 32767
    fmt.Println(v3, v29)
    fmt.Println(v2, v3, v4)
    if true {
        if (false && true) {
            v29 += 32767
            v3 += 4294967295
            if true {
                v30 := v1[2147483647]
                _ = v30
            } else {
                fmt.Println(v29, v4, v2)
            }
        } else {
            v1[2] = "a"
            fmt.Println(v29, v3, v29)
        }
    } else {
        fmt.Println(v2)
        var v31 uint8 = 128
        _ = v31
    }
    if (!true) {
        v2 -= 18446744073709551615
        v2 -= v2
    } else {
        v2 = v3
        for fi32 := 0; fi32 < 14; fi32++ {
            _ = len(v1)
            _ = fi32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
