package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int32; F1 int64 }
func genFn0(p0 int) uint8 {
    if (!(true || ("hello" <= ("" + (("hello" + "\n\t") + "\n\t"))))) {
        fmt.Println(p0)
    }
    v4 := []int{}
    p0 += 4294967295
    _ = p0
    _ = v4
    return 115
}
func genFn1(p0 string, p1 string) int {
    var v5 float32 = 0.0
    if false {
        var v6 int32 = 743
        _ = v6
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 float64 = float64(v5)
        v9 := GenS1{F0: 765, F1: 691}
        if false {
            if true {
                fmt.Println(v5, p0)
                v5 += v5
            } else {
                if true {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        if false {
                            fmt.Println(v8)
                            var v11 float32 = (1.5 + v5)
                            _ = v11
                        }
                        v12 := GenS1{F0: 776, F1: 4294967295}
                        v13 := []int{}
                        _ = fi10
                        _ = v12
                        _ = v13
                    }
                } else {
                    v14 := []int{}
                    _ = v14
                }
                if (-2147483648 != 1) {
                    if false {
                        fmt.Println(v8, v5)
                        p0 = (p1 + p1)
                    } else {
                        var v15 uint = 18446744073709551615
                        var v16 int8 = 100
                        _ = v15
                        _ = v16
                    }
                } else {
                    p0 = p1
                    p0 = ""
                }
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v17 uint64 = 42
    for fi18 := 0; fi18 < 16; fi18++ {
        if (-824 > 2147483647) {
            var v19 float64 = (0.1 - -2.25)
            for fi20 := 0; fi20 < 1; fi20++ {
                if ((!false) || ((false || (-527 == -2147483648)) && true)) {
                    var v21 uint32 = 0
                    _ = v21
                }
                v22 := GenS1{F0: 0, F1: 4294967295}
                v17 = 18446744073709551615
                p1 = (("foo bar" + p1) + p1)
                _ = fi20
                _ = v22
            }
            _ = v19
        } else {
            v5 -= v5
            p1 = ""
            p1 = (p0 + ("a" + "\n\t"))
        }
        var v23 uint32 = 0
        for fi24 := 0; fi24 < 16; fi24++ {
            switch "a" {
            case "hello":
                v25 := []int{fi24, fi18, fi18, -482}
                _ = v25
            }
            if (1 == 479) {
                v26 := GenS2{F0: -655, F1: 120}
                if true {
                    var v27 bool = ((fi24 != fi18) && (42 != -1))
                    _ = v27
                }
                _ = v26
            }
            v28 := []int{fi24, -297}
            if (32767 == 255) {
                v29 := GenS2{F0: 0, F1: -128}
                _ = v29
            } else {
                var v30 bool = false
                fmt.Println(p1, v5)
                _ = v30
            }
            _ = len(v28)
            fmt.Println(v23)
            _ = fi24
            _ = v28
        }
        _ = fi18
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v17
    _ = v5
    return 0
}

func main() {
    var v1 uint16 = 898
    v2 := []int{-9223372036854775808, -576, 128}
    v3 := make(map[int]int)
    v31 := GenS2{F0: 9223372036854775807, F1: 0}
    if len(v2) > 0 {
        v2[0] = -986
    }
    var v32 string = ("a" + "")
    var v33 int16 = 1000
    v34 := GenS3{F0: 9223372036854775807, F1: (v1 - v1)}
    _ = len(v2)
    v35 := GenS4{F0: -2147483648, F1: -1}
    v33 += -226
    var v36 float32 = 1e10
    var v37 int64 = 302
    var v38 int = 0
    var v39 int64 = (-(v37))
    var v40 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
