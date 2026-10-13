package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int64; F1 string }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 float32 = 3.14159
    fmt.Println(v4, v1)
    for fi5 := 0; fi5 < 10; fi5++ {
        if false {
            if (!false) {
                for fi6 := 0; fi6 < 8; fi6++ {
                    if ((fi5 / ((fi5) | 1)) < fi5) {
                        fmt.Println(fi5, fi6, fi6)
                    } else {
                        fmt.Println(fi5, v1, v4)
                    }
                    fmt.Println(v4, fi6)
                    v7 := GenS1{F0: (1 == 1), F1: ""}
                    _ = fi6
                    _ = v7
                }
                v3[""] = ""
            }
        } else {
            v1 -= v1
        }
        for fi8 := 0; fi8 < 2; fi8++ {
            v1 = v1
            v4 = 0.0
            var v9 int32 = -23
            _ = fi8
            _ = v9
        }
        _ = fi5
    }
    if true {
        v1 = (1e10 + v1)
    } else {
        for fi10 := 0; fi10 < 19; fi10++ {
            v3[""] = ""
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 float64 = v1
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 20; fi13++ {
                switch "a" {
                case "\n\t":
                    switch 32767 {
                    case -1:
                        v4 += float32(v1)
                    case 32767:
                        if true {
                            v3["foo bar"] = ""
                        } else {
                            var v14 int = (fi10 % ((fi10) | 1))
                            var v15 int = (v14 | (v14 & -419))
                            v16 := GenS2{F0: 42, F1: 2}
                            v1 = v1
                            v17 := GenS3{F0: -754, F1: ("hello" + "hello")}
                            _ = v14
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    }
                default:
                    var v18 int32 = 26
                    _ = v18
                }
                v4 -= v4
                v4 += v4
                _ = fi13
            }
            _ = fi10
        }
        if false {
            for fi19 := 0; fi19 < 4; fi19++ {
                for fi20 := 0; fi20 < 1; fi20++ {
                    v21 := GenS4{F0: -83, F1: "\n\t"}
                    fmt.Println(v4, fi20)
                    _ = fi20
                    _ = v21
                }
                var v22 uint = 18446744073709551615
                fmt.Println(fi19, v22)
                _ = fi19
                _ = v22
            }
        }
        v23 := []int{128}
        _ = v23
    }
    if (176 < 0) {
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := []int{32767}
            fmt.Println(v1)
            _ = fi24
            _ = v25
        }
    } else {
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 int64 = -953
            _ = fi26
            _ = v27
        }
        v28 := []int{9223372036854775807, 934, 2147483647}
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        v4 = v4
        fmt.Println(v4, v1)
        _ = fi29
    }
    v30 := []int{-596, 9223372036854775807}
    v4 = v4
    if (127 <= 1) {
        v30 = append(v30, 1)
        v4 -= v4
    } else {
        v4 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
