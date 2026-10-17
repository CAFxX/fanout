package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 bool }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint16; F2 int8 }
func genFn0(p0 string, p1 int32, p2 int16) bool {
    fmt.Println(p2, p1)
    v5 := []int{9223372036854775807, -9223372036854775808, 9223372036854775807, 533}
    switch "" {
    case "a":
        switch p0 {
        case "foo bar":
            p2 -= (p2 | p2)
        case "":
            fmt.Println(p0, p1)
        }
    case "hello":
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: "hello", F1: 2, F2: (false || (!true))}
            if (false || false) {
                for fi8 := 0; fi8 < 19; fi8++ {
                    if true {
                        p0 = ("hello" + "hello")
                        if true {
                            _ = len(v5)
                            var v9 uint16 = 686
                            _ = len(v5)
                            _ = v9
                        } else {
                            if false {
                                p2 = p2
                                p0 = "\n\t"
                                switch p0 {
                                case "a":
                                    var v10 uint8 = 1
                                    _ = v10
                                case "foo bar":
                                    fmt.Println(fi6, p0)
                                    var v11 uint64 = 168
                                    p0 = "a"
                                    p2 -= p2
                                    var v12 uint32 = 2147483647
                                    _ = v11
                                    _ = v12
                                case "":
                                    v13 := GenS2{F0: 65535, F1: p1}
                                    _ = v13
                                default:
                                    v14 := GenS1{F0: "a", F1: 983, F2: (!false)}
                                    _ = v14
                                }
                            }
                        }
                        p1 = p1
                    } else {
                        if len(v5) > 0 {
                            v5[0] = fi8
                        }
                        p1 += p1
                        v15 := GenS1{F0: p0, F1: 4294967295, F2: false}
                        _ = v15
                    }
                    v5 = append(v5, fi6)
                    _ = fi8
                }
                switch p0 {
                case "foo bar":
                    v5 = append(v5, (fi6 | fi6))
                default:
                    p0 = ("hello" + p0)
                }
                for fi16 := 0; fi16 < 1; fi16++ {
                    p2 = p2
                    var v17 string = p0
                    if true {
                        v5 = append(v5, fi16)
                        switch (p0 + p0) {
                        case "foo bar":
                            var v18 uint = 1
                            _ = v18
                        case "":
                            v19 := GenS1{F0: p0, F1: 2, F2: (0 > -1)}
                            _ = v19
                        case "hello":
                            v20 := GenS3{F0: 1, F1: 255}
                            _ = v20
                        }
                    }
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi6
            _ = v7
        }
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (42 <= 42)
}
func genFn1() uint16 {
    v21 := []int{1, -1}
    v21 = append(v21, 213)
    _ = v21
    return 1
}

func main() {
    v1 := []string{"hello", "a", ("a" + "\n\t"), "hello"}
    var v2 int = -9223372036854775808
    v3 := make(map[int]int)
    var v4 string = "foo bar"
    var v22 int8 = 123
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v23 uint32 = 418
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        v23 -= v23
        v25 := v3[v2]
        _ = fi24
        _ = v25
    }
    v2 = (v2 % ((v2) | 1))
    if (!(-1 > 0)) {
        var v26 float32 = (3.14159 * 1e10)
        _ = v26
    } else {
        switch ("foo bar" + ((("" + "foo bar") + v4) + v4)) {
        case "hello":
            v4 = ((("hello" + v4) + v4) + "foo bar")
        case "\n\t":
            var v27 uint = 2
            _ = v27
        }
        v4 = "a"
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        if true {
            v22 = v22
            var v29 float64 = -2.25
            v30 := GenS4{F0: 135, F1: 1, F2: (v22 / (((v22 ^ -18)) | 1))}
            _ = v29
            _ = v30
        } else {
            v22 = (v22 % (((v22 << ((255) & 63))) | 1))
            var v31 int8 = (^(v22))
            var v32 int8 = (v31 * (v22 + 1))
            _ = v31
            _ = v32
        }
        _ = fi28
    }
    var v33 int8 = (v22 >> ((42) & 63))
    fmt.Println(v4, v2)
    var v34 uint64 = 501
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
