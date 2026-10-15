package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
func genFn0(p0 int, p1 int32) uint64 {
    v7 := GenS1{F0: 42, F1: -9223372036854775808}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 string, p1 uint64, p2 bool) int {
    if p2 {
        p0 = (p0 + "hello")
        var v9 float64 = 1.5
        v10 := GenS1{F0: 0, F1: 719}
        _ = v10
        _ = v9
    }
    switch ("\n\t" + (("hello" + (("hello" + (p0 + p0)) + "\n\t")) + p0)) {
    case "":
        v11 := []int{42, 9223372036854775807}
        _ = v11
    case "\n\t":
        var v12 uint = 4294967295
        _ = v12
    case "a":
        if p2 {
            p0 = (p0 + p0)
        } else {
            var v13 uint64 = (p1 & 444)
            _ = v13
        }
    default:
        switch 4294967295 {
        case 0:
            switch ("" + ("\n\t" + p0)) {
            case "a":
                v14 := GenS1{F0: -9223372036854775808, F1: 318}
                _ = v14
            case "":
                if false {
                    var v15 string = p0
                    _ = v15
                }
            case "foo bar":
                v16 := GenS1{F0: 42, F1: 42}
                _ = v16
            }
        case 42:
            switch "hello" {
            case "\n\t":
                if (!(!(255 != 127))) {
                    v17 := []int{}
                    var v18 uint16 = 0
                    v19 := GenS1{F0: -9223372036854775808, F1: -1}
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    p2 = (65535 > uint(p1))
                    fmt.Println(p2, p0)
                    for fi20 := 0; fi20 < 7; fi20++ {
                        v21 := GenS2{F0: 77, F1: -936, F2: 0}
                        for fi22 := 0; fi22 < 1; fi22++ {
                            var v23 float64 = float64((float32(-2.25) * float32((1.5 / (3.14159 / float64(0.1))))))
                            var v24 int64 = int64(p1)
                            p2 = (32767 > 368)
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = fi20
                        _ = v21
                    }
                }
            case "":
                v25 := GenS3{F0: p2}
                _ = v25
            }
        }
    }
    p2 = (!true)
    _ = p0
    _ = p1
    _ = p2
    return -75
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []string{"foo bar", ("hello" + "hello"), "a"}
    v3 := []uint32{(v1 + v1), v1, (^(v1)), 363}
    v4 := []string{"a"}
    var v5 int32 = 2147483647
    v6 := []int{42, 1, 127, 402}
    v5 += 670
    if ((-2.25 - (0.1 - 0.1)) != 1e10) {
        v5 -= 42
        for fi26 := 0; fi26 < 16; fi26++ {
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 int = (fi26 ^ (-(fi26)))
                fmt.Println(fi26, v1)
                _ = fi27
                _ = v28
            }
            v1 += 42
            _ = fi26
        }
    } else {
        _ = len(v3)
    }
    if false {
        var v29 int = -2147483648
        if (928 != 119) {
            v30 := GenS3{F0: (!(-196 >= genFn1(("foo bar" + "\n\t"), 2, ("\n\t" != "\n\t"))))}
            _ = v30
        }
        v29 += v29
        _ = v29
    } else {
        var v31 uint = 18446744073709551615
        _ = v31
    }
    v5 = v5
    switch -2147483648 {
    case -1:
        v32 := GenS4{F0: 42}
        _ = v32
    default:
        v33 := GenS3{F0: true}
        _ = v33
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        v34 := GenS4{F0: -1}
        var v35 bool = true
        if (!v35) {
            fmt.Println(v5, v5)
        }
        _ = v34
        _ = v35
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
