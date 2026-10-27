package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 string; F2 int64 }
type GenS4 struct { F0 int }
func genFn0() string {
    switch "hello" {
    case "hello":
        if true {
            for fi5 := 0; fi5 < 1; fi5++ {
                var v6 int8 = 127
                _ = fi5
                _ = v6
            }
            v7 := GenS1{F0: "hello"}
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 uint32 = 779
                v10 := GenS1{F0: ("foo bar" + "a")}
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        } else {
            if (!(32767 < 128)) {
                if (65 == 2147483647) {
                    var v11 int32 = 482
                    v12 := []int{1, -635}
                    if false {
                        v11 -= v11
                        for fi13 := 0; fi13 < 17; fi13++ {
                            if len(v12) > 0 {
                                v12[0] = (fi13 + fi13)
                            }
                            _ = fi13
                        }
                        var v14 float64 = float64(-2.25)
                        _ = v14
                    }
                    _ = v11
                    _ = v12
                }
                var v15 uint16 = 255
                _ = v15
            }
            v16 := GenS1{F0: "foo bar"}
            _ = v16
        }
    case "\n\t":
        var v17 float64 = 1e10
        _ = v17
    default:
        var v18 float64 = 1.5
        _ = v18
    }
    var v19 uint16 = 65535
    var v20 float64 = 0.1
    v21 := GenS1{F0: "hello"}
    _ = v19
    _ = v20
    _ = v21
    return "hello"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 385
    var v3 float64 = -2.25
    var v4 int = -2147483648
    v1[394] = v4
    v22 := []int{42}
    fmt.Println(v3, v3, v2)
    for fi23 := 0; fi23 < 7; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            v2 = (^(v2))
            fmt.Println(v2, fi23, v3)
            switch "\n\t" {
            case "\n\t":
                fmt.Println(fi24)
                delete(v1, 4294967295)
                var v25 float64 = 1e10
                _ = v25
            case "a":
                if true {
                    fmt.Println(fi24, fi23)
                    v22 = append(v22, (v4 % ((v4) | 1)))
                    v3 = v3
                } else {
                    v26 := GenS2{F0: 1}
                    fmt.Println(fi24)
                    _ = v26
                }
            case "foo bar":
                v27 := GenS3{F0: 0, F1: "", F2: 0}
                _ = v27
            default:
                var v28 int64 = -1
                _ = v28
            }
            _ = fi24
        }
        _ = fi23
    }
    var v29 int8 = 127
    v3 -= (0.1 + v3)
    if (v4 < 2147483647) {
        v29 = v29
        v30 := GenS4{F0: (-(v4))}
        _ = v30
    } else {
        v4 -= v4
        var v31 float32 = 0.0
        _ = v31
    }
    var v32 float64 = ((v3 + -2.25) - v3)
    var v33 uint32 = 65535
    v34 := GenS3{F0: 293, F1: ("foo bar" + ""), F2: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
