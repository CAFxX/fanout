package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int16; F1 bool; F2 int32 }
func genFn0(p0 int64) int8 {
    switch "" {
    case "foo bar":
        for fi4 := 0; fi4 < 7; fi4++ {
            v5 := GenS1{F0: 773, F1: 819, F2: 1}
            p0 = (p0 / (((p0 * (p0 + p0))) | 1))
            v6 := GenS1{F0: 287, F1: 1, F2: 1}
            _ = fi4
            _ = v5
            _ = v6
        }
    case "\n\t":
        if ((!true) && false) {
            p0 -= (p0 * p0)
            p0 -= p0
        } else {
            switch "" {
            case "\n\t":
                if (false && (((1e10 - 1e10) <= float32(1.5)) && true)) {
                    if (true && true) {
                        p0 += p0
                        p0 -= p0
                    }
                    if (-2.25 == float64(-2.25)) {
                        for fi7 := 0; fi7 < 18; fi7++ {
                            var v8 int = 2147483647
                            v9 := []int{(fi7 + fi7), (v8 << ((255) & 63)), 42, v8}
                            _ = fi7
                            _ = v8
                            _ = v9
                        }
                    }
                    fmt.Println(p0)
                }
            default:
                p0 += 0
            }
            var v10 string = ("" + "hello")
            _ = v10
        }
    default:
        if (!true) {
            var v11 float32 = 1e10
            for fi12 := 0; fi12 < 15; fi12++ {
                switch fi12 {
                case 691:
                    v13 := []int{fi12, fi12}
                    _ = v13
                case 128:
                    var v14 int = fi12
                    _ = v14
                }
                _ = fi12
            }
            v11 -= float32(-2.25)
            _ = v11
        } else {
            for fi15 := 0; fi15 < 6; fi15++ {
                p0 -= (p0 - p0)
                p0 = p0
                if (!false) {
                    var v16 string = "\n\t"
                    _ = v16
                }
                _ = fi15
            }
        }
    }
    switch 127 {
    case 32767:
        var v17 int = 4294967295
        _ = v17
    case 2147483647:
        var v18 int64 = p0
        _ = v18
    default:
        switch 4294967295 {
        case 128:
            p0 -= (p0 >> ((281) & 63))
        case 0:
            var v19 int8 = 1
            _ = v19
        }
    }
    var v20 string = ("foo bar" + "\n\t")
    _ = p0
    _ = v20
    return 127
}

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    var v21 bool = (!(65535 < 42))
    if true {
        v21 = (-1 != 2147483647)
    } else {
        switch "a" {
        case "foo bar":
            v21 = true
        case "\n\t":
            v22 := GenS2{F0: 9223372036854775807, F1: "hello"}
            _ = v22
        case "hello":
            v21 = v21
        }
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := GenS3{F0: -32768, F1: false, F2: 60}
            v25 := GenS2{F0: 2147483647, F1: "a"}
            switch "hello" {
            case "":
                if (95 == 100) {
                    v21 = false
                    v21 = v21
                }
            default:
                delete(v2, 4294967295)
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 uint = 65535
        _ = v26
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    var v27 float32 = ((0.1 / 1.0) * 0.0)
    v21 = v21
    fmt.Println(v27, v27)
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 int8 = 100
        v30 := GenS2{F0: -924, F1: ("foo bar" + "foo bar")}
        v31 := GenS3{F0: 1, F1: (v21 && true), F2: 0}
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 int = 2147483647
    v21 = (0 > 850)
    var v33 int32 = 639
    v33 = (v33 % ((1) | 1))
    var v34 bool = v21
    if true {
        var v35 string = "hello"
        if (v34 || v34) {
            v36 := GenS1{F0: 249, F1: 661, F2: 0}
            _ = v36
        }
        if (true && v21) {
            _ = len(v3)
            v35 = "foo bar"
        }
        _ = v35
    }
    var v37 int64 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v37
}
