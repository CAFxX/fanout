package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 string; F1 uint64; F2 uint8 }
func genFn0(p0 bool, p1 uint32, p2 string) int {
    if (0 == 1) {
        for fi4 := 0; fi4 < 11; fi4++ {
            switch (fi4 * (fi4 ^ (fi4 + fi4))) {
            case 32767:
                fmt.Println(fi4, p2)
            case -457:
                var v5 int16 = -501
                _ = v5
            default:
                v6 := []int{int(p1), 9223372036854775807}
                _ = v6
            }
            var v7 float64 = 3.14159
            fmt.Println(v7, v7, p0)
            _ = fi4
            _ = v7
        }
    } else {
        p0 = p0
        if (!p0) {
            fmt.Println(p2, p1, p2)
        } else {
            fmt.Println(p0, p2)
            var v8 float64 = 1.5
            _ = v8
        }
        p0 = true
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 int8, p1 uint8) int16 {
    v9 := []int{genFn0(true, 2, "a"), 1}
    v10 := GenS1{F0: (p1 | (p1 << ((4294967295) & 63))), F1: (128 != 515)}
    v11 := GenS1{F0: (p1 | p1), F1: ("\n\t" == "\n\t")}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    var v1 string = (((("foo bar" + "") + "a") + (((("foo bar" + "\n\t") + ("\n\t" + "hello")) + "hello") + "")) + "a")
    var v2 uint8 = 128
    var v3 uint = 701
    v12 := GenS1{F0: v2, F1: true}
    v2 -= (v2 + v2)
    fmt.Println(v2, v1)
    v13 := []int{-2147483648}
    _ = len(v13)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v14 := GenS1{F0: v2, F1: true}
    if (2 < v3) {
        v1 = "hello"
        var v15 uint32 = uint32(v3)
        _ = v15
    }
    v3 += v3
    v16 := GenS1{F0: (v2 | v2), F1: true}
    if false {
        for fi17 := 0; fi17 < 15; fi17++ {
            if (!false) {
                v3 = (v3 - v3)
            } else {
                v1 = "foo bar"
                var v18 float64 = 0.0
                v2 -= (v2 + v2)
                _ = v18
            }
            _ = fi17
        }
        var v19 int8 = 100
        v20 := GenS2{F0: v2, F1: v2}
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 int32 = 2147483647
            _ = fi21
            _ = v22
        }
        var v23 int32 = 2147483647
        v24 := GenS1{F0: (v2 << ((v3) & 63)), F1: (!true)}
        _ = v23
        _ = v24
    }
    var v25 uint32 = 2147483648
    v26 := GenS3{F0: v1, F1: 918, F2: (v2 * 90)}
    for fi27 := 0; fi27 < 4; fi27++ {
        if false {
            v2 -= 255
            fmt.Println(v2, v25)
            fmt.Println(v1)
        }
        _ = fi27
    }
    switch (v1 + v1) {
    case "":
        var v28 uint = v3
        _ = v28
    case "\n\t":
        var v29 uint64 = 18446744073709551615
        _ = v29
    case "a":
        v2 = (v2 << ((v3) & 63))
    default:
        if (-246 >= -2147483648) {
            fmt.Println(v3)
            switch 127 {
            case 4294967295:
                v25 += uint32(v2)
            case -2147483648:
                fmt.Println(v2, v25, v1)
            case -504:
                var v30 int16 = 115
                _ = v30
            }
        } else {
            v31 := GenS1{F0: (v2 / ((v2) | 1)), F1: true}
            var v32 uint64 = 2
            _ = v31
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
