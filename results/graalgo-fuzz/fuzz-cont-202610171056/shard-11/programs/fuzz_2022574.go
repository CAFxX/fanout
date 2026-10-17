package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 bool }
type GenS2 struct { F0 int; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint8 }
func genFn0(p0 int32, p1 bool) int16 {
    p0 = (p0 ^ (p0 ^ p0))
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: 127, F1: -1, F2: false}
            v10 := []int{32767, (fi7 & fi8)}
            _ = fi8
            _ = v10
            _ = v9
        }
        if (0 == 42) {
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 string = ("\n\t" + "foo bar")
                _ = fi11
                _ = v12
            }
            p1 = (128 < 2)
            for fi13 := 0; fi13 < 18; fi13++ {
                fmt.Println(fi13, fi7, fi13)
                fmt.Println(fi13, fi7)
                fmt.Println(p0, fi13, p1)
                p1 = p1
                _ = fi13
            }
        } else {
            p1 = true
        }
        switch ((("\n\t" + "foo bar") + ("" + "a")) + (("foo bar" + (("foo bar" + ("hello" + "\n\t")) + ("a" + "foo bar"))) + "foo bar")) {
        case "hello":
            var v14 string = "\n\t"
            _ = v14
        case "foo bar":
            switch fi7 {
            case 32768:
                switch fi7 {
                case 127:
                    switch ("hello" + "hello") {
                    case "":
                        for fi15 := 0; fi15 < 18; fi15++ {
                            p0 += p0
                            _ = fi15
                        }
                    case "\n\t":
                        p0 -= p0
                    default:
                        if (-400 == fi7) {
                            if true {
                                fmt.Println(p1, p1, p1)
                                var v16 uint64 = 204
                                v17 := GenS1{F0: 96, F1: 500, F2: true}
                                _ = v16
                                _ = v17
                            }
                            var v18 int = (^(fi7))
                            if (-64 != -3) {
                                v18 = fi7
                            }
                            _ = v18
                        } else {
                            var v19 int64 = -9223372036854775808
                            v20 := []int{fi7, 4294967295, fi7, 636}
                            _ = v19
                            _ = v20
                        }
                    }
                case -254:
                    p1 = false
                default:
                    switch "hello" {
                    case "a":
                        var v21 int = (fi7 - fi7)
                        _ = v21
                    case "foo bar":
                        switch -9223372036854775808 {
                        case 849:
                            var v22 float64 = 0.1
                            _ = v22
                        default:
                            v23 := []int{fi7, fi7}
                            _ = v23
                        }
                    case "":
                        var v24 uint16 = 977
                        _ = v24
                    }
                }
            default:
                v25 := GenS2{F0: fi7, F1: p1, F2: 0}
                _ = v25
            }
        }
        _ = fi7
    }
    var v26 bool = p1
    _ = p0
    _ = p1
    _ = v26
    return -32768
}

func main() {
    var v1 int8 = -128
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 uint32 = 4294967295
    v5 := []string{}
    v6 := make(map[uint32]uint32)
    var v27 uint = 42
    fmt.Println(v27, v4, v27)
    for fi28 := 0; fi28 < 16; fi28++ {
        v6[v4] = v4
        v1 -= 127
        _ = fi28
    }
    v29 := GenS1{F0: 127, F1: 1, F2: false}
    switch "hello" {
    case "":
        var v30 int = 485
        _ = v30
    default:
        var v31 uint32 = v4
        _ = v31
    }
    var v32 float32 = 0.0
    if false {
        v33 := GenS1{F0: 255, F1: -69, F2: false}
        v6[(^(v4))] = (v4 % ((v4) | 1))
        _ = v33
    }
    v32 = (v32 * v32)
    var v34 bool = false
    v35 := GenS3{F0: uint64(v4), F1: 0, F2: 197}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
