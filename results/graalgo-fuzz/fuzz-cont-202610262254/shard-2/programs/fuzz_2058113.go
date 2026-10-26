package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 int32, p2 uint) uint32 {
    var v6 int32 = -1
    v7 := []int{-988, 42, -874, 2147483647}
    fmt.Println(p1, p2, p1)
    v6 = (p1 << (((^(p2))) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint32) int16 {
    var v8 uint = 258
    switch ("foo bar" + (("a" + "foo bar") + "foo bar")) {
    case "foo bar":
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 bool = (false || true)
            _ = fi9
            _ = v10
        }
    case "\n\t":
        v11 := []int{9223372036854775807, -2147483648, 127, 32768}
        _ = v11
    default:
        p0 = 782
    }
    _ = p0
    _ = v8
    return -864
}

func main() {
    var v1 int = -1
    v2 := make(map[int]uint32)
    var v3 int32 = 0
    var v4 int64 = 42
    var v5 uint16 = 42
    var v12 int16 = 840
    var v13 uint8 = 127
    v14 := []int{v1, -156}
    var v15 float32 = (0.0 / 1.5)
    if true {
        if true {
            v4 -= (-(v4))
        }
        v13 += (v13 - (v13 & v13))
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 string = "foo bar"
            _ = fi16
            _ = v17
        }
    }
    v18 := GenS1{F0: true, F1: genFn0(v5, int32(v5), 4294967295)}
    var v19 uint32 = 2147483648
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            fmt.Println(v1, v5, v1)
            _ = fi21
        }
        fmt.Println(v15)
        _ = fi20
    }
    if (4294967295 >= 387) {
        switch ("foo bar" + "hello") {
        case "a":
            for fi22 := 0; fi22 < 9; fi22++ {
                switch -9223372036854775808 {
                case 128:
                    v13 = 255
                default:
                    var v23 bool = (-54 < -97)
                    _ = v23
                }
                _ = fi22
            }
        case "hello":
            v13 = 128
        case "\n\t":
            v24 := GenS2{F0: genFn1((v19 & 2147483647))}
            _ = v24
        default:
            for fi25 := 0; fi25 < 8; fi25++ {
                switch (((("\n\t" + ("foo bar" + "hello")) + ("foo bar" + ("foo bar" + "hello"))) + "") + "hello") {
                case "hello":
                    fmt.Println(fi25)
                case "":
                    var v26 uint32 = 2
                    _ = v26
                }
                _ = fi25
            }
        }
        v5 = 2
    }
    fmt.Println(v4)
    v27 := v2[(v1 - (v1 % ((2147483647) | 1)))]
    var v28 uint = 1
    if true {
        for fi29 := 0; fi29 < 4; fi29++ {
            fmt.Println(fi29, v4, v15)
            for fi30 := 0; fi30 < 20; fi30++ {
                if true {
                    switch (v1 | (fi29 % (((-(fi30))) | 1))) {
                    case 2147483647:
                        v3 -= (v3 % ((v3) | 1))
                    case 1:
                        var v31 uint64 = 0
                        _ = v31
                    case 128:
                        fmt.Println(v19)
                    default:
                        var v32 float32 = (v15 - v15)
                        _ = v32
                    }
                    v27 = (^(v19))
                    v13 = (v13 >> (((v28 | v28)) & 63))
                } else {
                    v19 -= genFn0((v5 * (v5 >> (((v28 >> (((^(v28))) & 63))) & 63))), (v3 | (v3 * 851)), v28)
                }
                var v33 int8 = int8(fi29)
                _ = fi30
                _ = v33
            }
            if len(v14) > 0 {
                v14[0] = v1
            }
            _ = fi29
        }
        var v34 uint32 = (v19 - genFn0((v5 >> (((v28 | (v28 << (((v28 + (v28 ^ 1))) & 63)))) & 63)), (v3 * (v3 / ((v3) | 1))), v28))
        _ = v34
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        v1 = (v1 << ((v28) & 63))
        fmt.Println(v1, v13, v3)
        _ = fi35
    }
    v2[(v1 * v1)] = (^(v27))
    v4 = 42
    v36 := GenS1{F0: false, F1: (v19 % ((v27) | 1))}
    var v37 bool = (v1 <= (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
