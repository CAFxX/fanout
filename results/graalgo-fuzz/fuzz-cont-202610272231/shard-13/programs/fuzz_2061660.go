package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint16 }
func genFn0(p0 string, p1 bool) uint16 {
    switch ((p0 + (p0 + ((p0 + (p0 + ((("a" + (p0 + "hello")) + ((("foo bar" + "foo bar") + p0) + "")) + p0))) + "hello"))) + p0) {
    case "":
        if p1 {
            if ((128 < 105) && ("hello" >= "foo bar")) {
                var v5 int = -375
                v6 := []int{(v5 ^ 721), v5, (v5 ^ (v5 - (v5 >> ((42) & 63)))), v5}
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v5
                _ = v6
            } else {
                v7 := []int{32768, 745, 128}
                switch ("" + "foo bar") {
                case "\n\t":
                    switch 153 {
                    case -1:
                        for fi8 := 0; fi8 < 20; fi8++ {
                            var v9 float32 = (1.5 * ((-2.25 + 1.5) - float32(3.14159)))
                            _ = fi8
                            _ = v9
                        }
                    case 32767:
                        fmt.Println(p0, p1)
                    }
                case "":
                    p0 = "hello"
                case "a":
                    p1 = (false || false)
                }
                _ = v7
            }
            p1 = (!(p1 && true))
        }
    case "foo bar":
        var v10 uint32 = 222
        _ = v10
    default:
        var v11 uint8 = 128
        _ = v11
    }
    p1 = p1
    switch p0 {
    case "hello":
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint, p1 int, p2 uint32) string {
    fmt.Println(p2, p2, p0)
    v12 := GenS1{F0: false}
    switch (("a" + ("a" + "a")) + "") {
    case "a":
        v13 := []int{-1, (p1 >> (((^(p0))) & 63))}
        _ = v13
    case "foo bar":
        v14 := GenS2{F0: 42, F1: 56, F2: 42}
        _ = v14
    case "\n\t":
        var v15 int16 = -61
        _ = v15
    default:
        if false {
            p0 = 18446744073709551615
            switch (p1 - 0) {
            case 9223372036854775807:
                p1 -= 881
            }
        } else {
            var v16 bool = true
            _ = v16
        }
    }
    var v17 string = "hello"
    v17 = v17
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v17
    return v17
}

func main() {
    var v1 float64 = 1e10
    var v2 string = "\n\t"
    v3 := make(map[uint32]uint32)
    var v4 bool = (-1 == 283)
    v18 := GenS1{F0: (v4 && false)}
    var v19 int32 = -246
    switch v2 {
    case "hello":
        for fi20 := 0; fi20 < 14; fi20++ {
            v19 = (^(v19))
            _ = fi20
        }
    default:
        var v21 float32 = 0.1
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 int8 = 42
        var v24 int64 = -9223372036854775808
        var v25 uint32 = 1
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    v26 := GenS2{F0: genFn0(("\n\t" + v2), (v1 != (0.0 / 1.0))), F1: 42, F2: genFn0("hello", ((v4 && false) || (v4 || true)))}
    var v27 float64 = 1e10
    v27 = v1
    switch 227 {
    case -1:
        var v28 int32 = 42
        _ = v28
    default:
        if true {
            v29 := GenS2{F0: genFn0(((v2 + v2) + v2), false), F1: 116, F2: 255}
            var v30 int8 = -29
            fmt.Println(v19)
            _ = v29
            _ = v30
        }
    }
    v31 := []int{-9223372036854775808, 32768, 4294967295, -726}
    var v32 uint64 = 1
    fmt.Println(v19)
    var v33 int64 = -92
    v3[124] = 4294967295
    v27 = v1
    var v34 int16 = -32768
    var v35 int = 900
    var v36 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
