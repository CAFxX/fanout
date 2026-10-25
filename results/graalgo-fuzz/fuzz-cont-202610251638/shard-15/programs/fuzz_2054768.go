package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16, p1 int8, p2 int8) uint {
    p1 = (^(p2))
    switch ("a" + "a") {
    case "foo bar":
        p0 -= p0
    case "hello":
        v7 := []int{-135}
        _ = v7
    case "a":
        fmt.Println(p0, p2, p0)
    default:
        p0 = (p0 ^ p0)
    }
    if false {
        for fi8 := 0; fi8 < 8; fi8++ {
            p1 = 127
            if false {
                var v9 float32 = (-2.25 - -2.25)
                var v10 bool = false
                _ = v10
                _ = v9
            }
            if (!(4294967295 == 0)) {
                if (!false) {
                    if (true && (829 != 0)) {
                        var v11 uint8 = 42
                        _ = v11
                    }
                } else {
                    for fi12 := 0; fi12 < 10; fi12++ {
                        if true {
                            var v13 bool = true
                            fmt.Println(p0, p2)
                            _ = v13
                        } else {
                            var v14 int = 945
                            _ = v14
                        }
                        _ = fi12
                    }
                }
            }
            _ = fi8
        }
        var v15 uint16 = p0
        _ = v15
    } else {
        if (4294967295 == 0) {
            var v16 int32 = 154
            fmt.Println(v16)
            for fi17 := 0; fi17 < 15; fi17++ {
                if ((-2.25 > (-2.25 / (0.1 * 1e10))) || (true && false)) {
                    v18 := []int{(fi17 / ((127) | 1)), (fi17 ^ 32767)}
                    _ = v18
                } else {
                    fmt.Println(p1, fi17, p1)
                }
                _ = fi17
            }
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    v1 := []int{-785, 4294967295, 32768}
    v2 := make(map[int]uint32)
    var v3 float32 = -2.25
    v4 := make(map[int]int)
    var v5 int32 = -789
    v6 := []string{("hello" + ("a" + "a")), "hello"}
    var v19 int16 = -32768
    v2[2147483647] = 2
    v4[4294967295] = 14
    var v20 float32 = (3.14159 - v3)
    fmt.Println(v20)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        switch "\n\t" {
        case "hello":
            v3 = (v3 - v20)
        case "\n\t":
            if (0.1 < float64(((-2.25 - v20) / v3))) {
                switch fi21 {
                case 4294967295:
                    var v22 uint8 = 127
                    _ = v22
                case 0:
                    v23 := v2[fi21]
                    _ = v23
                case -885:
                    fmt.Println(v20, v19)
                default:
                    fmt.Println(fi21)
                }
            } else {
                switch fi21 {
                case 127:
                    switch "\n\t" {
                    case "a":
                        v19 = v19
                    }
                case 32767:
                    fmt.Println(fi21, v20, v3)
                case -9223372036854775808:
                    switch "hello" {
                    case "a":
                        v3 -= v3
                    case "":
                        if len(v6) > 0 {
                            v6[0] = ((("" + "\n\t") + "hello") + "")
                        }
                    default:
                        for fi24 := 0; fi24 < 1; fi24++ {
                            v20 = v3
                            _ = fi24
                        }
                    }
                default:
                    v3 += (v3 - 1.5)
                }
                v6 = append(v6, "a")
            }
        case "foo bar":
            v20 = 1.5
        }
        v3 = v3
        _ = fi21
    }
    _ = len(v4)
    v25 := GenS1{F0: 32768, F1: 0, F2: v19}
    switch -455 {
    case 32767:
        var v26 int64 = 9223372036854775807
        _ = v26
    default:
        var v27 float32 = float32(1.5)
        _ = v27
    }
    v19 += v19
    v28 := GenS2{F0: (v5 / ((0) | 1))}
    v3 = v3
    var v29 uint32 = 0
    v4[42] = 977
    if len(v1) > 0 {
        v1[0] = -35
    }
    var v30 int64 = -709
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
