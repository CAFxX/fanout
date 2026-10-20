package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0(p0 int, p1 int16, p2 uint32) uint32 {
    p2 += (^(p2))
    p1 = (p1 << ((735) & 63))
    v4 := []int{(p0 << ((uint(p2)) & 63)), 127}
    var v5 int64 = 784
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (^(p2))
}
func genFn1() string {
    var v6 int32 = 826
    v7 := []int{}
    if (!(true && false)) {
        switch "\n\t" {
        case "hello":
            for fi8 := 0; fi8 < 10; fi8++ {
                if (("" + "foo bar") == "") {
                    for fi9 := 0; fi9 < 10; fi9++ {
                        v6 = (v6 ^ v6)
                        switch ("foo bar" + (("hello" + "") + "")) {
                        case "\n\t":
                            v6 += v6
                        }
                        var v10 uint32 = 698
                        _ = fi9
                        _ = v10
                    }
                    v6 += v6
                } else {
                    v6 += 42
                }
                _ = len(v7)
                _ = fi8
            }
        case "a":
            fmt.Println(v6)
        case "foo bar":
            v6 += v6
        }
    } else {
        switch "a" {
        case "\n\t":
            v6 += (-(v6))
        }
        if len(v7) > 0 {
            v7[0] = -9223372036854775808
        }
    }
    _ = v6
    _ = v7
    return ""
}

func main() {
    var v1 int = 32767
    var v2 uint8 = 37
    var v3 uint64 = 1
    v11 := []int{}
    fmt.Println(v1)
    var v12 int = v1
    v13 := GenS1{F0: 42, F1: 100}
    var v14 float32 = 0.1
    v12 -= 169
    var v15 uint = 18446744073709551615
    for fi16 := 0; fi16 < 7; fi16++ {
        switch "a" {
        case "a":
            var v17 int32 = 1
            _ = v17
        case "hello":
            var v18 int32 = -453
            _ = v18
        case "":
            if len(v11) > 0 {
                v11[0] = (^(v12))
            }
        default:
            v11 = append(v11, -1)
        }
        var v19 bool = true
        if v19 {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            var v20 bool = (((("\n\t" + "foo bar") + "a") + ("\n\t" + ("\n\t" + ("foo bar" + (("" + "hello") + (("foo bar" + "\n\t") + "hello")))))) <= "hello")
            for fi21 := 0; fi21 < 2; fi21++ {
                switch (fi16 << ((42) & 63)) {
                case 42:
                    v14 -= float32(0.0)
                case 128:
                    v1 = (v12 / (((fi21 ^ 42)) | 1))
                case 0:
                    var v22 int64 = -9223372036854775808
                    _ = v22
                default:
                    if v20 {
                        switch (^(v1)) {
                        case 2147483647:
                            if v19 {
                                fmt.Println(v3, v2)
                                var v23 float32 = (v14 - (v14 - v14))
                                _ = v23
                            } else {
                                fmt.Println(v1, fi21, v20)
                                if v19 {
                                    if true {
                                        fmt.Println(v15, fi16)
                                        var v24 int8 = 79
                                        v15 = v15
                                        if (v15 == (v15 ^ v15)) {
                                            v20 = v20
                                            var v25 int32 = 42
                                            var v26 int8 = (v24 ^ v24)
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = v24
                                    }
                                }
                                if len(v11) > 0 {
                                    v11[0] = -2147483648
                                }
                            }
                        }
                        var v27 uint8 = v2
                        var v28 uint64 = (v3 ^ 4294967295)
                        var v29 int64 = 42
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        if len(v11) > 0 {
                            v11[0] = -2147483648
                        }
                    }
                }
                v30 := GenS1{F0: 1000, F1: -70}
                _ = fi21
                _ = v30
            }
            _ = v20
        } else {
            if (("\n\t" + ("\n\t" + "hello")) >= (("" + ((("a" + ("foo bar" + "a")) + "a") + ("\n\t" + "a"))) + "\n\t")) {
                v31 := GenS1{F0: -1, F1: 100}
                var v32 int = v12
                switch v12 {
                case 0:
                    if false {
                        for fi33 := 0; fi33 < 6; fi33++ {
                            var v34 uint16 = 806
                            _ = fi33
                            _ = v34
                        }
                    }
                case 127:
                    v35 := GenS1{F0: 32767, F1: -79}
                    _ = v35
                }
                _ = v31
                _ = v32
            }
            v3 -= 18446744073709551615
        }
        _ = fi16
        _ = v19
    }
    fmt.Println(v12, v14)
    v15 = 18446744073709551615
    fmt.Println(v12)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    if (1 <= uint8(v15)) {
        var v36 float64 = 1e10
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
