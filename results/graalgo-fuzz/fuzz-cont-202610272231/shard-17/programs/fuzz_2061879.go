package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint8, p1 uint, p2 string) int {
    if (false || ((!((128 <= 0) && false)) && false)) {
        var v4 int = 128
        _ = v4
    } else {
        var v5 int32 = -951
        var v6 int64 = 1
        v7 := GenS1{F0: 42, F1: 599, F2: (127 > 0)}
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 int32 = 2147483647
    fmt.Println(v8, p1)
    for fi9 := 0; fi9 < 7; fi9++ {
        v8 = v8
        var v10 uint64 = 0
        _ = fi9
        _ = v10
    }
    var v11 uint8 = (p0 ^ 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return -237
}

func main() {
    var v1 int8 = 0
    var v2 uint16 = 1
    var v3 int32 = 0
    v2 += (v2 / ((v2) | 1))
    for fi12 := 0; fi12 < 4; fi12++ {
        v1 += (v1 % ((v1) | 1))
        v13 := GenS1{F0: 255, F1: -32768, F2: (((v3 > 0) && (((float64(((0.1 * ((float32(0.0) + (1.5 + (1e10 * (0.0 - 1.5)))) + 0.1)) / (-2.25 * 3.14159))) * 3.14159) + float64((0.1 + 3.14159))) > 1.5)) || true)}
        var v14 float32 = 0.1
        _ = fi12
        _ = v13
        _ = v14
    }
    v2 -= v2
    fmt.Println(v2)
    v15 := GenS2{F0: 42}
    v16 := []int{364, -9223372036854775808, 307}
    v16 = append(v16, genFn0(220, 4294967295, ("foo bar" + (("foo bar" + ("\n\t" + "")) + "foo bar"))))
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS1{F0: 255, F1: 1000, F2: (42 != 255)}
        v2 = (v2 & 0)
        switch "" {
        case "foo bar":
            switch "foo bar" {
            case "\n\t":
                _ = len(v16)
            case "foo bar":
                if false {
                    var v19 bool = ((("foo bar" + "hello") + "hello") != (("foo bar" + ("a" + "\n\t")) + ((("\n\t" + ("a" + "a")) + "hello") + ("\n\t" + ("\n\t" + "\n\t")))))
                    v1 = v1
                    if v19 {
                        v2 = (v2 % ((2) | 1))
                        for fi20 := 0; fi20 < 6; fi20++ {
                            v2 += v2
                            v3 = (v3 + 0)
                            _ = fi20
                        }
                    } else {
                        var v21 int32 = 2147483647
                        _ = v21
                    }
                    _ = v19
                }
            default:
                v2 += (^(v2))
            }
        case "\n\t":
            switch "" {
            case "foo bar":
                v2 += (v2 % ((v2) | 1))
            case "hello":
                v2 = (v2 + v2)
            case "":
                v3 = v3
            }
        default:
            var v22 uint8 = 2
            _ = v22
        }
        _ = fi17
        _ = v18
    }
    if true {
        fmt.Println(v3, v3)
        var v23 int16 = 42
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        if ((true || (2147483648 < 2)) && true) {
            switch (fi24 * fi24) {
            case 970:
                fmt.Println(fi24, v1, fi24)
            default:
                if len(v16) > 0 {
                    v16[0] = int(v3)
                }
            }
            v2 = v2
            for fi25 := 0; fi25 < 5; fi25++ {
                v3 = v3
                _ = fi25
            }
        }
        var v26 uint32 = 0
        _ = fi24
        _ = v26
    }
    v1 += 0
    if true {
        if true {
            v3 = v3
        } else {
            fmt.Println(v2, v2)
            var v27 float32 = (0.0 - -2.25)
            var v28 uint64 = 2
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 11; fi29++ {
            if false {
                v30 := GenS1{F0: 255, F1: -32768, F2: (v1 < v1)}
                _ = v30
            }
            v2 = v2
            _ = fi29
        }
        v1 += v1
    } else {
        v3 += (v3 - v3)
        if (v1 > v1) {
            switch -289 {
            case 0:
                v16 = append(v16, 247)
            case -2147483648:
                var v31 uint16 = (^(v2))
                _ = v31
            case 2147483647:
                var v32 float32 = 0.1
                _ = v32
            default:
                for fi33 := 0; fi33 < 9; fi33++ {
                    _ = len(v16)
                    switch "\n\t" {
                    case "":
                        v34 := GenS1{F0: 1, F1: int16(v1), F2: true}
                        _ = v34
                    case "a":
                        var v35 uint64 = 512
                        _ = v35
                    case "foo bar":
                        v36 := GenS1{F0: 2, F1: 1, F2: (true && false)}
                        _ = v36
                    default:
                        v37 := GenS2{F0: 0}
                        _ = v37
                    }
                    _ = fi33
                }
            }
        }
        _ = len(v16)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
