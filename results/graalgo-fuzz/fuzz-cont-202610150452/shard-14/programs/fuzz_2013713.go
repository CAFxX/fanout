package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 bool }
func genFn0() string {
    v6 := GenS1{F0: 0}
    var v7 int16 = -1
    var v8 int64 = 1
    var v9 int32 = 2147483647
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return ("foo bar" + "")
}

func main() {
    var v1 uint32 = 42
    var v2 uint32 = v1
    var v3 float32 = (-2.25 * ((float32(-2.25) - 0.0) + 0.0))
    v4 := make(map[int]bool)
    var v5 uint8 = 28
    v10 := GenS2{F0: 2}
    v11 := GenS2{F0: 42}
    if (false && (0.1 > float64(((v3 - v3) / ((v3 + ((((float32(-2.25) - (v3 * v3)) - v3) + ((1.5 * v3) - -2.25)) + (0.1 + v3))) / ((0.0 / (v3 / 0.1)) - v3)))))) {
        fmt.Println(v5)
        for fi12 := 0; fi12 < 1; fi12++ {
            _ = len(v4)
            _ = fi12
        }
        var v13 int = -2147483648
        _ = v13
    } else {
        v2 = v2
        var v14 int = -2147483648
        var v15 float32 = 0.1
        _ = v14
        _ = v15
    }
    fmt.Println(v2, v3)
    v16 := GenS2{F0: 18446744073709551615}
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS2{F0: uint64(v1)}
        _ = fi17
        _ = v18
    }
    fmt.Println(v3)
    switch "" {
    case "hello":
        var v19 uint = 4294967295
        _ = v19
    case "\n\t":
        v4[4294967295] = (("\n\t" + "\n\t") != ("foo bar" + (("a" + "") + (((("foo bar" + (("\n\t" + ("\n\t" + ("hello" + "a"))) + ("foo bar" + ("" + ("" + "\n\t"))))) + "") + "hello") + "hello"))))
    case "foo bar":
        v20 := []int{9223372036854775807}
        _ = v20
    }
    var v21 int8 = 77
    if true {
        v22 := []int{128, 1}
        fmt.Println(v21)
        v2 += v2
        _ = v22
    } else {
        var v23 int8 = v21
        _ = v23
    }
    v2 -= 2147483648
    v24 := GenS3{F0: v1, F1: 435}
    if true {
        for fi25 := 0; fi25 < 11; fi25++ {
            if false {
                switch "foo bar" {
                case "foo bar":
                    v26 := []int{-382}
                    _ = v26
                }
            } else {
                v3 += (v3 - (v3 * v3))
                var v27 float32 = 1e10
                for fi28 := 0; fi28 < 13; fi28++ {
                    if (!false) {
                        v27 = ((((-2.25 / 3.14159) - (((-2.25 * ((v27 - v27) + v27)) - (v3 + (v3 - v3))) * ((1.5 * 1.5) + float32(1e10)))) / (-2.25 * 0.1)) * 3.14159)
                        v5 = (^(v5))
                        v4[42] = (32767 <= 0)
                    } else {
                        switch 2147483647 {
                        case 127:
                            delete(v4, (fi25 & fi28))
                        }
                    }
                    var v29 float32 = (v3 * float32((3.14159 + float64(v3))))
                    if (!(((fi28 * fi25) == (fi28 + 128)) && true)) {
                        if true {
                            v30 := GenS4{F0: false}
                            _ = v30
                        } else {
                            v31 := GenS4{F0: (42 < int64(fi28))}
                            v21 = v21
                            v32 := []int{}
                            _ = v31
                            _ = v32
                        }
                        v33 := []int{}
                        _ = v33
                    }
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
            if (65535 > 0) {
                v34 := GenS1{F0: 42}
                v2 += 0
                fmt.Println(v1)
                _ = v34
            }
            v35 := []int{int(v5), (fi25 * (fi25 ^ 127))}
            _ = fi25
            _ = v35
        }
        switch "\n\t" {
        case "foo bar":
            if false {
                v1 -= v1
                v21 += -1
            }
        default:
            switch "foo bar" {
            case "\n\t":
                _ = len(v4)
            case "a":
                fmt.Println(v3, v3)
            default:
                v3 -= -2.25
            }
        }
        var v36 uint32 = v2
        _ = v36
    } else {
        var v37 int = -452
        switch "foo bar" {
        case "hello":
            v37 -= 904
        case "foo bar":
            for fi38 := 0; fi38 < 9; fi38++ {
                v39 := []int{338, v37, 1, fi38}
                _ = fi38
                _ = v39
            }
        }
        _ = v37
    }
    var v40 int = 4294967295
    v41 := GenS1{F0: 1}
    v5 -= 234
    v42 := []int{(v40 / ((v40) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
