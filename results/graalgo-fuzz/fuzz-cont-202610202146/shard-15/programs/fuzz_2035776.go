package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int; F2 int32 }
func genFn0() bool {
    var v4 uint8 = 42
    v4 += v4
    v5 := []int{553, 128}
    v4 -= (v4 * (v4 / ((127) | 1)))
    if (0 <= -1) {
        switch "" {
        case "foo bar":
            v4 = (v4 * v4)
        case "\n\t":
            v4 = v4
        case "hello":
            fmt.Println(v4)
        }
        fmt.Println(v4)
        var v6 string = ""
        _ = v6
    }
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 int8, p1 int8, p2 bool) int {
    switch 128 {
    case -1:
        switch -2147483648 {
        case 0:
            fmt.Println(p2, p2)
        }
    default:
        if p2 {
            if (18446744073709551615 == 208) {
                fmt.Println(p1, p1)
                switch 127 {
                case 42:
                    for fi7 := 0; fi7 < 10; fi7++ {
                        p2 = ("hello" < "hello")
                        _ = fi7
                    }
                default:
                    var v8 uint32 = 42
                    _ = v8
                }
                if p2 {
                    v9 := []int{9223372036854775807, 0, -992, -716}
                    _ = v9
                } else {
                    var v10 uint32 = 2147483648
                    _ = v10
                }
            } else {
                var v11 uint = 0
                fmt.Println(p0, p1, p2)
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 uint8 = 128
                    for fi14 := 0; fi14 < 11; fi14++ {
                        var v15 int32 = 42
                        _ = fi14
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
            var v16 int = 42
            _ = v16
        } else {
            var v17 float32 = 1e10
            v18 := []int{-2147483648, -2147483648, -1, -2147483648}
            var v19 uint = 42
            _ = v17
            _ = v18
            _ = v19
        }
    }
    p2 = (-32768 == 1000)
    _ = p0
    _ = p1
    _ = p2
    return 847
}

func main() {
    var v1 int = 127
    var v2 uint64 = 2
    v3 := make(map[uint32]uint32)
    var v20 int = (-(v1))
    switch "\n\t" {
    case "hello":
        delete(v3, 2)
    case "":
        if true {
            var v21 int8 = 100
            _ = v21
        } else {
            v22 := GenS1{F0: 128, F1: (v1 / ((v20) | 1)), F2: -777}
            _ = v22
        }
    case "foo bar":
        var v23 uint16 = 0
        _ = v23
    }
    v2 -= v2
    if true {
        v20 += 2147483647
        var v24 string = ("foo bar" + (("foo bar" + "") + "a"))
        var v25 uint32 = 42
        _ = v24
        _ = v25
    }
    if false {
        v26 := GenS2{F0: 1}
        switch 128 {
        case 969:
            v20 = genFn1(25, 19, ((255 < 2) && (true && true)))
        default:
            v20 = (v1 * v1)
        }
        if true {
            fmt.Println(v20, v20)
            for fi27 := 0; fi27 < 1; fi27++ {
                switch int(v2) {
                case -886:
                    v2 -= 18446744073709551615
                default:
                    v2 = 4294967295
                }
                v1 = 127
                var v28 float64 = ((1e10 - 1e10) - -2.25)
                _ = fi27
                _ = v28
            }
            v2 += 42
        }
        _ = v26
    } else {
        fmt.Println(v2, v2, v1)
        v29 := GenS3{F0: false, F1: 2147483647, F2: 2147483647}
        _ = v29
    }
    if true {
        var v30 uint = 2
        _ = v30
    }
    fmt.Println(v2, v1, v2)
    if (false || (false || true)) {
        var v31 uint32 = 231
        v1 = -1
        _ = v31
    } else {
        fmt.Println(v1)
        v1 += (v1 | (-(v1)))
    }
    v32 := v3[0]
    delete(v3, v32)
    v33 := []int{42, v20, genFn1(2, -128, false)}
    v20 = -750
    fmt.Println(v1, v32, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
}
