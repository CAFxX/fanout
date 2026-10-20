package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 string }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 string }
func genFn0(p0 string, p1 uint64, p2 string) int {
    var v5 float64 = (-2.25 + 3.14159)
    v6 := []int{0, 127, 0, 127}
    if (255 > 785) {
        var v7 float32 = float32((3.14159 * (v5 - v5)))
        var v8 uint16 = 65535
        _ = v7
        _ = v8
    }
    switch -1 {
    case 127:
        var v9 uint = 18446744073709551615
        _ = v9
    case 42:
        fmt.Println(v5, v5, p2)
    case -9223372036854775808:
        fmt.Println(p0)
    default:
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                p0 = "hello"
                v6 = append(v6, (fi10 ^ (fi11 / ((fi10) | 1))))
                var v12 float32 = (1e10 * 1e10)
                fmt.Println(p2, v5, p0)
                p0 = "\n\t"
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -2147483648
}
func genFn1(p0 uint8, p1 uint16) string {
    var v13 float32 = 0.1
    p0 = p0
    v14 := []int{-459, 4294967295}
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int32 = -2147483648
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        fmt.Println(p0, p1)
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return "foo bar"
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    var v4 int16 = 42
    var v18 int = 654
    v19 := []int{238, v18, (v18 & 1), (v18 - (v18 & (v18 | 9223372036854775807)))}
    v1 -= v1
    v3[847] = (v18 << (((v1 << ((v1) & 63))) & 63))
    v20 := GenS1{F0: -2147483648, F1: v18, F2: ("foo bar" + "")}
    v18 = (-(v18))
    var v21 int64 = 1
    for fi22 := 0; fi22 < 4; fi22++ {
        v4 -= v4
        switch ("" + "\n\t") {
        case "hello":
            for fi23 := 0; fi23 < 16; fi23++ {
                if (true && false) {
                    fmt.Println(fi23)
                } else {
                    if true {
                        v21 = v21
                        v24 := GenS2{F0: v4, F1: 0}
                        if len(v19) > 0 {
                            v19[0] = (fi23 * 4294967295)
                        }
                        v18 -= (v18 | (-(fi22)))
                        v21 -= v21
                        fmt.Println(v21, fi23)
                        var v25 string = "\n\t"
                        _ = v24
                        _ = v25
                    } else {
                        var v26 int32 = -250
                        switch v18 {
                        case 53:
                            v18 = (v18 + fi22)
                        default:
                            v4 = v4
                        }
                        v27 := GenS3{F0: ""}
                        _ = v26
                        _ = v27
                    }
                }
                if (2 > 42) {
                    v28 := GenS3{F0: "hello"}
                    _ = v28
                } else {
                    delete(v2, 548)
                }
                switch "\n\t" {
                case "hello":
                    var v29 int16 = -32768
                    _ = v29
                }
                _ = fi23
            }
        }
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        _ = fi22
    }
    fmt.Println(v18)
    for fi30 := 0; fi30 < 20; fi30++ {
        if (65535 <= 18446744073709551615) {
            fmt.Println(v1, v21)
            switch "a" {
            case "a":
                v4 += (^(v4))
            case "hello":
                fmt.Println(v21)
            case "\n\t":
                v3[148] = genFn0((("hello" + "a") + ""), 899, "hello")
            }
            for fi31 := 0; fi31 < 9; fi31++ {
                var v32 string = "a"
                var v33 float64 = 1e10
                _ = fi31
                _ = v32
                _ = v33
            }
        } else {
            v34 := GenS1{F0: -1, F1: fi30, F2: ""}
            v21 = (v21 + v21)
            v18 += genFn0(("hello" + ""), 556, ("foo bar" + ("hello" + "foo bar")))
            _ = v34
        }
        v21 -= v21
        var v35 uint64 = 1
        _ = fi30
        _ = v35
    }
    fmt.Println(v18)
    fmt.Println(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
