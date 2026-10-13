package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint; F1 string; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int16) string {
    v5 := GenS1{F0: 437, F1: -1}
    switch -2147483648 {
    case 42:
        if (!false) {
            for fi6 := 0; fi6 < 8; fi6++ {
                if false {
                    v7 := GenS1{F0: 42, F1: -1000}
                    v8 := GenS2{F0: (3.14159 <= 3.14159), F1: ""}
                    fmt.Println(p0, fi6)
                    _ = v7
                    _ = v8
                }
                v9 := []int{fi6, fi6, 0}
                _ = fi6
                _ = v9
            }
            p0 -= (p0 + (^(p0)))
        } else {
            var v10 uint = 1
            p0 = 1
            var v11 uint8 = 0
            _ = v10
            _ = v11
        }
    default:
        p0 = p0
    }
    p0 = (p0 & (p0 ^ p0))
    switch 2147483647 {
    case 532:
        for fi12 := 0; fi12 < 19; fi12++ {
            p0 += (p0 | p0)
            switch "\n\t" {
            case "\n\t":
                v13 := []int{32768}
                _ = v13
            case "foo bar":
                v14 := GenS2{F0: ((fi12 / ((826) | 1)) != 1), F1: ""}
                _ = v14
            case "a":
                v15 := []int{}
                _ = v15
            default:
                switch "\n\t" {
                case "":
                    if (-588 > 0) {
                        v16 := GenS1{F0: 1, F1: -12}
                        var v17 uint64 = 0
                        _ = v16
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 16; fi18++ {
                            switch "foo bar" {
                            case "foo bar":
                                var v19 bool = true
                                _ = v19
                            }
                            var v20 string = ("foo bar" + "hello")
                            var v21 uint8 = 1
                            _ = fi18
                            _ = v20
                            _ = v21
                        }
                        var v22 int8 = -71
                        _ = v22
                    }
                }
            }
            switch fi12 {
            case 128:
                v23 := GenS2{F0: (!true), F1: (("a" + "hello") + "foo bar")}
                _ = v23
            default:
                var v24 uint = 463
                _ = v24
            }
            _ = fi12
        }
    case 32768:
        for fi25 := 0; fi25 < 16; fi25++ {
            v26 := GenS1{F0: 0, F1: 0}
            _ = fi25
            _ = v26
        }
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = v5
    return "hello"
}
func genFn1(p0 uint, p1 uint8) string {
    v27 := []int{}
    p0 = p0
    _ = p0
    _ = p1
    _ = v27
    return ""
}

func main() {
    var v1 float32 = (0.0 - float32(float64(1.5)))
    var v2 uint = 1
    var v3 uint8 = 2
    v4 := []uint32{4294967295, 537}
    v28 := GenS3{F0: v2, F1: (("hello" + "hello") + "foo bar"), F2: 65535}
    var v29 uint64 = 18446744073709551615
    switch "hello" {
    case "":
        fmt.Println(v3, v1)
    case "hello":
        switch "foo bar" {
        case "":
            fmt.Println(v1)
        default:
            switch "" {
            case "a":
                var v30 int32 = -2147483648
                _ = v30
            }
        }
    case "a":
        fmt.Println(v29)
    default:
        fmt.Println(v29, v1)
    }
    v31 := GenS4{F0: 4294967295, F1: v3}
    fmt.Println(v1, v3)
    for fi32 := 0; fi32 < 2; fi32++ {
        v29 -= (^(v29))
        var v33 float64 = float64(0.1)
        v3 = v3
        _ = fi32
        _ = v33
    }
    v1 = float32((float64(v1) - 1.5))
    v1 -= v1
    var v34 int8 = -128
    v35 := GenS5{F0: v34}
    v36 := GenS4{F0: (v29 * (v29 | 1)), F1: v3}
    for fi37 := 0; fi37 < 2; fi37++ {
        for fi38 := 0; fi38 < 3; fi38++ {
            switch ("a" + "foo bar") {
            case "a":
                v34 = (v34 << ((18446744073709551615) & 63))
            case "hello":
                switch fi38 {
                case 42:
                    v39 := GenS1{F0: 42, F1: 2147483647}
                    _ = v39
                default:
                    var v40 uint64 = v29
                    _ = v40
                }
            case "\n\t":
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            var v41 uint8 = (^(v3))
            var v42 uint64 = (^(v29))
            _ = fi38
            _ = v41
            _ = v42
        }
        _ = fi37
    }
    switch "hello" {
    case "foo bar":
        v29 += v29
    case "a":
        var v43 uint16 = 65535
        _ = v43
    case "":
        var v44 float32 = v1
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
