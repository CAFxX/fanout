package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint16, p1 int32, p2 uint64) uint16 {
    fmt.Println(p1, p2)
    fmt.Println(p0, p0)
    if (!(42 < 100)) {
        if true {
            v7 := []int{9223372036854775807, 9223372036854775807}
            p2 -= p2
            _ = v7
        }
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{0}
    var v3 uint32 = 2
    v4 := []string{"hello", "foo bar", ("" + "hello")}
    var v5 bool = (false && (("foo bar" + ((("hello" + (("a" + "hello") + (("foo bar" + ("hello" + "a")) + "a"))) + ("" + "foo bar")) + "\n\t")) == ("foo bar" + "hello")))
    var v6 uint = 435
    v8 := GenS1{F0: 183, F1: 104}
    v3 += (^(v3))
    var v9 uint64 = 42
    fmt.Println(v6, v3, v5)
    if (1 <= 1) {
        if v5 {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 float64 = 3.14159
                switch fi10 {
                case -1:
                    fmt.Println(v3, v6, fi10)
                case 593:
                    v6 -= (^(v6))
                case 0:
                    var v12 int8 = -128
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            v13 := GenS2{F0: -28, F1: ((((1.5 - 1.5) * 0.1) + (float64(((1e10 + (-2.25 * (float32(3.14159) + 0.0))) - (((0.0 - 0.1) * 1e10) / 1e10))) + (3.14159 / 0.1))) == 3.14159)}
            _ = v13
        }
    } else {
        v14 := GenS1{F0: 2, F1: -887}
        if false {
            v15 := GenS2{F0: 735, F1: v5}
            switch 4294967295 {
            case 0:
                fmt.Println(v3)
            default:
                var v16 uint16 = 0
                _ = v16
            }
            _ = v15
        }
        var v17 string = ("a" + ("" + "\n\t"))
        _ = v14
        _ = v17
    }
    switch "" {
    case "\n\t":
        fmt.Println(v6, v9, v9)
    case "a":
        v5 = v5
    default:
        fmt.Println(v3)
    }
    if false {
        for fi18 := 0; fi18 < 12; fi18++ {
            if (825 < 18446744073709551615) {
                v3 = v3
                v19 := GenS1{F0: 128, F1: -895}
                fmt.Println(v6)
                _ = v19
            } else {
                fmt.Println(v6, v3)
                var v20 string = "hello"
                delete(v1, 65535)
                _ = v20
            }
            v6 = (v6 % ((v6) | 1))
            _ = fi18
        }
        _ = len(v2)
        _ = len(v4)
    } else {
        var v21 uint8 = uint8(v6)
        v22 := GenS3{F0: uint(v21), F1: 2, F2: (^(v21))}
        if v5 {
            if false {
                fmt.Println(v6, v5)
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 int8 = 42
                    v3 = v3
                    switch "\n\t" {
                    case "hello":
                        v21 += (v21 & v21)
                    case "\n\t":
                        if (v9 != (v9 & v9)) {
                            if false {
                                _ = len(v1)
                            } else {
                                v1[v3] = ("" <= ("foo bar" + "a"))
                            }
                        }
                    case "foo bar":
                        var v25 int16 = -931
                        _ = v25
                    }
                    _ = fi23
                    _ = v24
                }
            } else {
                switch ("hello" + ("foo bar" + "a")) {
                case "a":
                    for fi26 := 0; fi26 < 12; fi26++ {
                        var v27 int64 = int64(v9)
                        v2 = append(v2, 945)
                        v21 = 255
                        _ = fi26
                        _ = v27
                    }
                default:
                    v21 = 1
                }
            }
        } else {
            switch 1 {
            case 42:
                fmt.Println(v5, v3, v9)
            default:
                v6 = 2
            }
        }
        _ = v21
        _ = v22
    }
    v28 := GenS1{F0: 127, F1: -1}
    _ = len(v1)
    v3 = (v3 & v3)
    v5 = ((^(v3)) >= v3)
    v9 = 948
    delete(v1, 2)
    v29 := GenS4{F0: 2147483648, F1: 272}
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 uint = (v6 >> (((v6 >> ((2) & 63))) & 63))
        delete(v1, (v3 * v3))
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := GenS3{F0: 466, F1: 2, F2: 42}
            _ = fi32
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    var v34 float32 = ((1.5 * ((1e10 - (0.1 + 3.14159)) - 0.0)) + 1e10)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
