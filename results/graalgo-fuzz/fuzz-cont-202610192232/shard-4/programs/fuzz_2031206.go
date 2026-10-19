package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0(p0 uint8, p1 int64, p2 int64) string {
    var v5 bool = false
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{(fi6 & fi6), fi6, 1}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ("" + "\n\t")
}
func genFn1(p0 int, p1 int64, p2 uint64) int {
    if (0 == 2) {
        switch p0 {
        case 761:
            v8 := []int{p0, 2147483647}
            _ = v8
        case 0:
            if (-1 == -1) {
                fmt.Println(p2)
                p2 -= 162
            }
        case -1:
            p1 += int64(p0)
        default:
            var v9 int64 = (-(p1))
            _ = v9
        }
        v10 := []int{42, 32768}
        p0 -= (p0 * 42)
        _ = v10
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 uint = 255
    var v2 uint32 = 1
    var v3 int16 = 269
    v4 := []string{"foo bar", ("\n\t" + (("foo bar" + ("\n\t" + "foo bar")) + "a")), "\n\t"}
    for fi11 := 0; fi11 < 8; fi11++ {
        if true {
            var v12 bool = true
            if v12 {
                v4 = append(v4, "hello")
                if (27 <= 128) {
                    var v13 bool = ((v2 >> ((42) & 63)) <= (^(v2)))
                    _ = v13
                }
                _ = len(v4)
            } else {
                v2 = 2147483647
                var v14 bool = (6 > 83)
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 int16 = v3
                    switch "" {
                    case "a":
                        fmt.Println(v12, v14)
                    case "hello":
                        var v17 uint16 = 2
                        _ = v17
                    case "\n\t":
                        switch (("foo bar" + "hello") + ("a" + "hello")) {
                        case "foo bar":
                            v12 = v14
                        default:
                            if true {
                                v18 := GenS1{F0: 2}
                                var v19 uint64 = 18446744073709551615
                                var v20 int16 = -552
                                v2 -= (v2 - 141)
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                        }
                    default:
                        switch "a" {
                        case "":
                            v16 = (v3 + (v16 % (((^(v3))) | 1)))
                        default:
                            if v12 {
                                v21 := GenS1{F0: 2}
                                var v22 int8 = 122
                                switch 32768 {
                                case 1:
                                    switch "" {
                                    case "a":
                                        fmt.Println(fi15)
                                        v16 += 811
                                        var v23 bool = v14
                                        _ = v23
                                    case "\n\t":
                                        v2 += 2147483647
                                    default:
                                        v22 = (v22 & v22)
                                        var v24 float64 = 0.1
                                        v4 = append(v4, "a")
                                        _ = v24
                                    }
                                case 42:
                                    if v14 {
                                        var v25 uint64 = 18446744073709551615
                                        if (9223372036854775807 > (fi11 % ((fi11) | 1))) {
                                            _ = len(v4)
                                            v14 = (!(v12 && (!v14)))
                                            var v26 int32 = 42
                                            _ = v26
                                        } else {
                                            v2 = v2
                                            var v27 bool = v14
                                            _ = v27
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = ""
                                        }
                                        _ = v25
                                    }
                                case -3:
                                    v14 = v12
                                }
                                _ = v21
                                _ = v22
                            }
                        }
                    }
                    v1 = 4294967295
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            if (410 < -1) {
                fmt.Println(fi11)
            } else {
                var v28 bool = v12
                _ = len(v4)
                v29 := GenS2{F0: fi11}
                _ = v28
                _ = v29
            }
            _ = v12
        } else {
            fmt.Println(fi11, fi11, v3)
        }
        _ = fi11
    }
    if (1 != -49) {
        v30 := GenS1{F0: 42}
        if (32767 <= (v3 & v3)) {
            v2 = v2
        } else {
            if false {
                v4 = append(v4, "a")
            }
        }
        fmt.Println(v3, v2, v2)
        _ = v30
    }
    fmt.Println(v2)
    fmt.Println(v3)
    _ = len(v4)
    var v31 string = "\n\t"
    fmt.Println(v1, v1)
    if (false && (false || false)) {
        v32 := GenS3{F0: 255}
        _ = v32
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v2 = (v2 * 625)
    v33 := GenS2{F0: 32767}
    fmt.Println(v1, v1)
    v34 := GenS4{F0: (v2 / (((v2 * v2)) | 1)), F1: v1}
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
}
