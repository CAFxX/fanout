package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 int32; F2 uint8 }
func genFn0(p0 int8, p1 int64) string {
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := []int{544, fi5, fi5}
        p0 = (p0 ^ p0)
        _ = fi5
        _ = v6
    }
    v7 := []int{-2147483648, 387, 128, 4294967295}
    _ = p0
    _ = p1
    _ = v7
    return ("foo bar" + "\n\t")
}

func main() {
    var v1 int16 = -32768
    var v2 bool = true
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    var v8 int32 = -2147483648
    var v9 uint = 42
    var v10 uint8 = 212
    switch ("a" + "foo bar") {
    case "":
        v11 := []int{128, 32768, 1}
        _ = v11
    case "foo bar":
        v8 = v8
    case "hello":
        var v12 uint16 = 325
        _ = v12
    default:
        if (4294967295 == 2147483648) {
            v13 := GenS1{F0: 0}
            v8 += (v8 & (v8 >> (((^(v9))) & 63)))
            _ = v13
        } else {
            v14 := GenS1{F0: 73}
            if true {
                v8 += -2147483648
                switch "hello" {
                case "foo bar":
                    v1 = 42
                case "hello":
                    for fi15 := 0; fi15 < 4; fi15++ {
                        v1 -= (v1 * 32767)
                        var v16 int32 = v8
                        for fi17 := 0; fi17 < 10; fi17++ {
                            v18 := GenS2{F0: 0, F1: (v9 | v9), F2: true}
                            var v19 string = "a"
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        _ = fi15
                        _ = v16
                    }
                }
            }
            if v2 {
                for fi20 := 0; fi20 < 3; fi20++ {
                    switch "a" {
                    case "a":
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v2 = false
                            _ = fi21
                        }
                    case "":
                        v22 := GenS3{F0: 995, F1: 927}
                        _ = v22
                    case "\n\t":
                        fmt.Println(fi20, v9, v10)
                    default:
                        v23 := GenS1{F0: 42}
                        _ = v23
                    }
                    v24 := GenS4{F0: 4294967295}
                    _ = fi20
                    _ = v24
                }
                var v25 int8 = 0
                v26 := GenS5{F0: 1, F1: 42, F2: v10}
                _ = v25
                _ = v26
            } else {
                switch "\n\t" {
                case "foo bar":
                    v27 := GenS5{F0: 887, F1: -1, F2: (^(v10))}
                    _ = v27
                case "hello":
                    v28 := GenS1{F0: 39}
                    _ = v28
                }
            }
            _ = v14
        }
    }
    var v29 float32 = (((1.5 - ((-2.25 * (0.1 / float32((0.0 - -2.25)))) / ((-2.25 - ((0.1 / 1.0) / -2.25)) / (-2.25 + 1.5)))) - 1e10) - (1.5 - (-2.25 - 0.0)))
    switch 32767 {
    case 9223372036854775807:
        v8 = -529
    }
    delete(v4, uint32(v10))
    for fi30 := 0; fi30 < 13; fi30++ {
        var v31 uint16 = 255
        var v32 string = (("a" + ((("foo bar" + "a") + "\n\t") + "hello")) + "\n\t")
        fmt.Println(v9, v9)
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 int16 = v1
    var v34 int64 = 661
    if v2 {
        for fi35 := 0; fi35 < 1; fi35++ {
            var v36 uint64 = 18446744073709551615
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 14; fi37++ {
            var v38 int32 = (v8 - 2147483647)
            switch fi37 {
            case -969:
                for fi39 := 0; fi39 < 9; fi39++ {
                    v40 := GenS1{F0: -32}
                    _ = fi39
                    _ = v40
                }
            }
            _ = fi37
            _ = v38
        }
    }
    v41 := GenS5{F0: 255, F1: 42, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v8
    _ = v9
}
