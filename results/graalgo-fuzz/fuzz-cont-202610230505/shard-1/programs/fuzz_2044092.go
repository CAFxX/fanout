package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint8; F1 bool; F2 int32 }
func genFn0(p0 uint32, p1 uint64) int64 {
    if (-897 > -2147483648) {
        switch "a" {
        case "a":
            switch 42 {
            case -2147483648:
                p0 -= p0
            case 127:
                p0 = p0
            case -9223372036854775808:
                v5 := []int{32768}
                _ = v5
            default:
                fmt.Println(p1, p1)
            }
        }
        fmt.Println(p1)
        switch 546 {
        case 32768:
            switch "foo bar" {
            case "":
                for fi6 := 0; fi6 < 3; fi6++ {
                    v7 := []int{fi6, -936, fi6}
                    _ = fi6
                    _ = v7
                }
            case "\n\t":
                v8 := []int{2147483647}
                _ = v8
            case "a":
                p0 -= p0
            default:
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 uint16 = 65535
                    if (127 > 128) {
                        var v11 uint = 1
                        p0 = (p0 & p0)
                        _ = v11
                    }
                    switch ((("\n\t" + ("hello" + ("foo bar" + ((("a" + (("foo bar" + "\n\t") + ("\n\t" + ("" + ("foo bar" + ("foo bar" + ("a" + "\n\t"))))))) + "") + ((("\n\t" + "hello") + "hello") + ""))))) + "") + "a") {
                    case "hello":
                        for fi12 := 0; fi12 < 8; fi12++ {
                            var v13 uint = 42
                            _ = fi12
                            _ = v13
                        }
                    case "foo bar":
                        var v14 int32 = 2147483647
                        _ = v14
                    case "a":
                        var v15 uint32 = (p0 / ((65535) | 1))
                        _ = v15
                    }
                    _ = fi9
                    _ = v10
                }
            }
        }
    }
    v16 := GenS1{F0: 0, F1: ((127 > 82) && true)}
    _ = p0
    _ = p1
    _ = v16
    return 187
}
func genFn1(p0 uint8, p1 int16) int32 {
    if (2 > 2) {
        fmt.Println(p1, p1)
        for fi17 := 0; fi17 < 5; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                fmt.Println(fi17)
                _ = fi18
            }
            _ = fi17
        }
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        p1 = p1
        if (true && false) {
            v20 := GenS1{F0: p0, F1: true}
            v21 := GenS1{F0: (p0 - p0), F1: (!true)}
            _ = v20
            _ = v21
        }
        v22 := GenS2{F0: 245, F1: (false || true), F2: -2147483648}
        _ = fi19
        _ = v22
    }
    if true {
        var v23 float32 = 1e10
        switch "\n\t" {
        case "a":
            v23 = v23
        case "hello":
            p1 += (p1 - (p1 | (p1 & (p1 - 1000))))
        }
        _ = v23
    } else {
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 int32 = 2147483647
            switch fi24 {
            case 610:
                v25 = 1
            default:
                v26 := GenS2{F0: 0, F1: (p0 == (^(p0))), F2: (v25 % ((v25) | 1))}
                _ = v26
            }
            v25 += v25
            _ = fi24
            _ = v25
        }
        v27 := []int{163, -635, 777}
        _ = v27
    }
    v28 := GenS2{F0: (p0 - p0), F1: (0 >= -371), F2: -2147483648}
    _ = p0
    _ = p1
    _ = v28
    return -2147483648
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -719
    v3 := []int{}
    var v4 int32 = (v2 * v2)
    switch "" {
    case "":
        var v29 float32 = 1.5
        _ = v29
    case "\n\t":
        var v30 int = -2147483648
        _ = v30
    case "foo bar":
        if (!false) {
            switch 747 {
            case 720:
                var v31 int8 = 1
                _ = v31
            }
            fmt.Println(v2)
            v1[("a" + "a")] = -30
        }
    }
    v2 -= 0
    v2 = -107
    v32 := GenS1{F0: 38, F1: true}
    v2 -= v2
    var v33 int64 = 519
    fmt.Println(v4, v4, v33)
    v34 := GenS2{F0: 220, F1: (!true), F2: v4}
    var v35 bool = true
    v2 -= genFn1(128, 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
