package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint16 }
type GenS2 struct { F0 int; F1 string; F2 uint }
type GenS3 struct { F0 int64 }
func genFn0() int8 {
    var v6 uint32 = 187
    var v7 uint16 = 20
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v6, v6)
        _ = fi8
    }
    _ = v6
    _ = v7
    return -128
}
func genFn1(p0 int16, p1 int32, p2 int) uint64 {
    if false {
        switch 42 {
        case 426:
            for fi9 := 0; fi9 < 3; fi9++ {
                p0 = p0
                for fi10 := 0; fi10 < 5; fi10++ {
                    if false {
                        if true {
                            v11 := []int{p2, -1}
                            _ = v11
                        } else {
                            var v12 uint64 = 380
                            _ = v12
                        }
                        var v13 uint32 = 0
                        switch p2 {
                        case 128:
                            v13 = 4294967295
                            v14 := []int{fi10}
                            p2 = (-(fi10))
                            p1 = p1
                            _ = v14
                        case -779:
                            var v15 bool = true
                            _ = v15
                        case 32768:
                            fmt.Println(p0)
                        default:
                            var v16 int8 = 0
                            _ = v16
                        }
                        var v17 float32 = (((1.5 - 1e10) * 0.0) / 1.5)
                        _ = v13
                        _ = v17
                    }
                    if true {
                        var v18 bool = false
                        fmt.Println(fi9, p2)
                        fmt.Println(v18, fi9)
                        _ = v18
                    } else {
                        v19 := GenS1{F0: p2, F1: 0, F2: 1}
                        v20 := GenS2{F0: p2, F1: "hello", F2: 356}
                        v21 := []int{fi9, fi10}
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    _ = fi10
                }
                v22 := []int{p2, (fi9 + (p2 | (p2 >> ((1) & 63))))}
                _ = fi9
                _ = v22
            }
        }
        switch "\n\t" {
        case "\n\t":
            v23 := GenS1{F0: -143, F1: 9223372036854775807, F2: 888}
            _ = v23
        case "hello":
            v24 := []int{}
            _ = v24
        case "foo bar":
            var v25 uint32 = 4294967295
            _ = v25
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 int = fi26
            p0 = (p0 + p0)
            _ = fi26
            _ = v27
        }
    }
    p1 += p1
    var v28 string = ("\n\t" + (("" + ("a" + ("a" + "a"))) + (("foo bar" + "a") + (("foo bar" + ((("" + "\n\t") + "foo bar") + ("hello" + "hello"))) + (("\n\t" + ("a" + "a")) + "\n\t")))))
    switch v28 {
    case "a":
        v29 := []int{(p2 / ((9223372036854775807) | 1))}
        _ = v29
    case "hello":
        v28 = (v28 + "foo bar")
    }
    v30 := []int{4294967295, p2, (p2 - (p2 - p2)), p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v30
    return 1
}

func main() {
    v1 := []string{""}
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]int)
    var v4 int32 = -1
    var v5 uint64 = 2
    var v31 bool = ((643 != 1) && ((^(v2)) < 4294967295))
    v2 += v2
    v5 -= genFn1(-570, v4, -492)
    v32 := GenS2{F0: 1, F1: (("" + (((("foo bar" + "hello") + ("\n\t" + ("a" + ("" + ("\n\t" + "\n\t"))))) + "foo bar") + "")) + "foo bar"), F2: 18446744073709551615}
    var v33 float64 = ((float64(3.14159) * 1e10) / -2.25)
    switch "" {
    case "foo bar":
        var v34 int16 = -258
        _ = v34
    default:
        fmt.Println(v2, v4)
    }
    v33 -= 0.1
    if v31 {
        if v31 {
            _ = len(v3)
        } else {
            if v31 {
                switch 0 {
                case 2147483647:
                    var v35 bool = (1 >= 127)
                    _ = v35
                }
                v36 := GenS3{F0: 42}
                _ = len(v1)
                _ = v36
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
