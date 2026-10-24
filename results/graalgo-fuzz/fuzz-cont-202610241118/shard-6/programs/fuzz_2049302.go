package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int32; F1 int8; F2 bool }
type GenS3 struct { F0 uint16; F1 int8; F2 uint }
func genFn0(p0 uint16, p1 string) int64 {
    var v4 uint = 0
    fmt.Println(p1)
    v5 := []int{42, 949, 32768}
    if false {
        if false {
            switch -2147483648 {
            case 400:
                p0 -= (^(p0))
            default:
                var v6 int = 32768
                _ = v6
            }
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    v3 := []int{4294967295}
    var v7 uint8 = 143
    for fi8 := 0; fi8 < 11; fi8++ {
        if (42 > 124) {
            if (-2.25 < 1e10) {
                if (2147483647 == -682) {
                    switch (fi8 & fi8) {
                    case -804:
                        var v9 int64 = 1
                        _ = v9
                    }
                } else {
                    v7 = (v7 | (^(v7)))
                    var v10 float32 = 3.14159
                    _ = v10
                }
            } else {
                v7 += v7
                var v11 float64 = 3.14159
                _ = v11
            }
            if (!(true && true)) {
                v7 -= 42
            } else {
                switch fi8 {
                case 1:
                    for fi12 := 0; fi12 < 6; fi12++ {
                        fmt.Println(v7, fi8, v7)
                        if false {
                            v3 = append(v3, fi8)
                            var v13 bool = true
                            _ = v13
                        }
                        switch "foo bar" {
                        case "":
                            if true {
                                var v14 int = fi12
                                _ = v14
                            } else {
                                switch fi12 {
                                case 42:
                                    switch ("a" + ("" + (("hello" + (("\n\t" + "foo bar") + (("\n\t" + ("" + (("\n\t" + "\n\t") + "\n\t"))) + (((((((((("foo bar" + ("foo bar" + "a")) + ((("hello" + "\n\t") + ("a" + "a")) + ("" + "\n\t"))) + (("a" + ("" + "hello")) + "")) + (("foo bar" + "\n\t") + "hello")) + "\n\t") + ("foo bar" + "hello")) + "foo bar") + "\n\t") + "\n\t") + ("\n\t" + ("" + "foo bar")))))) + (("foo bar" + ("foo bar" + "foo bar")) + ("a" + "\n\t"))))) {
                                    case "hello":
                                        var v15 bool = (false && ((-2147483648 <= (-(fi12))) && false))
                                        var v16 uint8 = v7
                                        fmt.Println(v15, fi8, v16)
                                        var v17 int = (-(fi8))
                                        var v18 float64 = 1.5
                                        var v19 int16 = -626
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    case "a":
                                        v7 = (^(v7))
                                    case "":
                                        v3 = append(v3, fi8)
                                    }
                                default:
                                    _ = len(v2)
                                }
                            }
                        case "foo bar":
                            v20 := GenS1{F0: "\n\t", F1: 127, F2: 18446744073709551615}
                            _ = v20
                        }
                        var v21 string = "\n\t"
                        _ = fi12
                        _ = v21
                    }
                case 32767:
                    var v22 bool = true
                    _ = v22
                default:
                    v23 := GenS1{F0: "a", F1: 202, F2: 1}
                    _ = v23
                }
            }
            fmt.Println(v7, fi8)
        }
        v2[2147483647] = "foo bar"
        var v24 uint16 = 0
        _ = fi8
        _ = v24
    }
    fmt.Println(v7)
    fmt.Println(v7)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v7)
    default:
        if false {
            var v25 int8 = 127
            _ = v25
        } else {
            v26 := v2[1]
            v26 = (v26 + (v26 + ("foo bar" + v26)))
            v27 := GenS1{F0: v26, F1: v7, F2: 2}
            _ = v26
            _ = v27
        }
    }
    switch ("a" + "a") {
    case "foo bar":
        v28 := GenS1{F0: ("a" + (("foo bar" + "\n\t") + (("foo bar" + ("foo bar" + "hello")) + ""))), F1: v7, F2: 0}
        _ = v28
    case "\n\t":
        var v29 int32 = 42
        _ = v29
    }
    v7 = (v7 ^ (v7 / (((^(v7))) | 1)))
    if false {
        if true {
            v7 = v7
            if (61 < 4294967295) {
                v7 -= v7
                _ = len(v1)
                var v30 int = -240
                _ = v30
            }
            if false {
                var v31 uint = 18446744073709551615
                v7 = 128
                fmt.Println(v7, v7)
                _ = v31
            } else {
                v7 -= v7
            }
        }
        v32 := GenS2{F0: -273, F1: 127, F2: ((9223372036854775807 <= 0) || true)}
        var v33 uint8 = 255
        _ = v32
        _ = v33
    } else {
        v7 += (v7 * v7)
    }
    v34 := GenS3{F0: 0, F1: 0, F2: 1}
    var v35 int32 = 61
    var v36 int32 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v7
}
