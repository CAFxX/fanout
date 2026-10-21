package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int64; F2 bool }
type GenS3 struct { F0 int64; F1 bool; F2 int64 }
func genFn0(p0 int8, p1 uint16) int {
    v5 := GenS1{F0: 18446744073709551615}
    fmt.Println(p1)
    if true {
        v6 := GenS1{F0: 2}
        if false {
            switch "hello" {
            case "foo bar":
                var v7 uint = 4294967295
                _ = v7
            case "hello":
                fmt.Println(p0)
            default:
                switch "" {
                case "\n\t":
                    if false {
                        switch ("\n\t" + "") {
                        case "a":
                            var v8 uint64 = 472
                            _ = v8
                        case "foo bar":
                            switch "hello" {
                            case "":
                                p1 = p1
                            default:
                                if true {
                                    var v9 int16 = 1
                                    v10 := GenS1{F0: 255}
                                    for fi11 := 0; fi11 < 10; fi11++ {
                                        v12 := []int{-935, fi11}
                                        p0 -= p0
                                        if len(v12) > 0 {
                                            v12[0] = (fi11 * fi11)
                                        }
                                        _ = fi11
                                        _ = v12
                                    }
                                    _ = v10
                                    _ = v9
                                } else {
                                    var v13 int8 = -29
                                    if false {
                                        v14 := []int{0, 2147483647, 0}
                                        v15 := GenS1{F0: 0}
                                        p0 = v13
                                        _ = v14
                                        _ = v15
                                    }
                                    _ = v13
                                }
                            }
                        }
                        for fi16 := 0; fi16 < 9; fi16++ {
                            var v17 uint16 = (p1 + p1)
                            p1 += p1
                            var v18 int16 = -32768
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                case "a":
                    var v19 int = -769
                    _ = v19
                case "foo bar":
                    v20 := GenS1{F0: 4294967295}
                    _ = v20
                }
            }
        } else {
            switch ("hello" + ("foo bar" + ("a" + "a"))) {
            case "a":
                var v21 uint32 = 65535
                _ = v21
            default:
                p0 = (p0 << ((450) & 63))
            }
            fmt.Println(p1, p1)
            var v22 bool = (!(9223372036854775807 > 1))
            _ = v22
        }
        _ = v6
    } else {
        var v23 int64 = 865
        p1 -= (p1 >> ((1) & 63))
        v24 := GenS2{F0: 42, F1: (-(v23)), F2: (0 != 2147483647)}
        _ = v23
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v5
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 1
    v3 := make(map[uint32]bool)
    var v4 int64 = 9223372036854775807
    if true {
        if (!(!(true && true))) {
            v4 += (v4 & -452)
            var v25 float64 = 1.5
            _ = v25
        }
        var v26 int8 = -128
        _ = v26
    } else {
        v4 = (v4 % (((v4 - v4)) | 1))
        v4 -= 0
        var v27 int16 = 1000
        _ = v27
    }
    v28 := []int{}
    v29 := GenS1{F0: 18446744073709551615}
    v3[42] = true
    var v30 int16 = -32768
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    if (-595 >= -1) {
        v31 := GenS3{F0: (v4 / (((v4 ^ -836)) | 1)), F1: (false || true), F2: v4}
        var v32 uint32 = 358
        _ = v31
        _ = v32
    } else {
        var v33 int64 = (^(v4))
        _ = v33
    }
    var v34 string = "a"
    var v35 uint16 = 373
    for fi36 := 0; fi36 < 18; fi36++ {
        fmt.Println(v2)
        delete(v1, 42)
        _ = fi36
    }
    if false {
        var v37 int8 = -109
        _ = v37
    }
    delete(v3, 121)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
}
