package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint }
func genFn0(p0 uint8, p1 int16) string {
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: int(p0), F1: 1, F2: 0}
        _ = fi6
        _ = v7
    }
    var v8 int32 = 0
    p1 = (p1 * int16(p0))
    for fi9 := 0; fi9 < 18; fi9++ {
        p1 = (-(p1))
        var v10 string = ("" + "\n\t")
        for fi11 := 0; fi11 < 20; fi11++ {
            p1 = p1
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    var v12 uint8 = p0
    _ = p0
    _ = p1
    _ = v12
    _ = v8
    return "a"
}
func genFn1(p0 int8) uint64 {
    if (-1 == 1) {
        var v13 int = 426
        _ = v13
    }
    v14 := GenS1{F0: -2147483648, F1: -2147483648, F2: 1}
    switch ("foo bar" + ("" + "hello")) {
    case "foo bar":
        var v15 bool = (!false)
        _ = v15
    case "\n\t":
        v16 := []int{-2147483648, -9223372036854775808, 335}
        _ = v16
    default:
        for fi17 := 0; fi17 < 17; fi17++ {
            p0 -= 1
            _ = fi17
        }
    }
    _ = p0
    _ = v14
    return 18446744073709551615
}

func main() {
    var v1 int32 = -573
    var v2 uint64 = 1
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    v5 := make(map[string]int)
    v2 -= 2
    if ((127 == -304) || (("a" + (("foo bar" + "\n\t") + ("a" + ("a" + "")))) > "\n\t")) {
        v3[""] = 2147483647
    } else {
        v18 := []int{2147483647, -385, 563}
        if true {
            if len(v18) > 0 {
                v18[0] = 9223372036854775807
            }
            v1 -= (v1 & v1)
            var v19 int32 = (v1 % (((v1 * v1)) | 1))
            _ = v19
        }
        for fi20 := 0; fi20 < 12; fi20++ {
            switch 42 {
            case 2147483647:
                v21 := GenS2{F0: -9223372036854775808}
                _ = v21
            case 1:
                v22 := GenS3{F0: 0, F1: 255, F2: 255}
                _ = v22
            case 0:
                var v23 string = (("foo bar" + (("hello" + ("\n\t" + "\n\t")) + "foo bar")) + "hello")
                _ = v23
            }
            switch "hello" {
            case "a":
                var v24 uint8 = 128
                _ = v24
            case "\n\t":
                _ = len(v5)
            }
            _ = fi20
        }
        _ = v18
    }
    if (true || (32767 >= 42)) {
        v25 := GenS3{F0: -50, F1: 42, F2: 2}
        _ = v25
    }
    var v26 uint64 = (v2 / ((18446744073709551615) | 1))
    v27 := []int{0, 32768, 127, 32767}
    switch "a" {
    case "\n\t":
        if (!((((("a" + (("a" + "foo bar") + "foo bar")) + (("" + "\n\t") + "a")) != "foo bar") || (true && true)) && true)) {
            var v28 int = 54
            _ = v28
        }
    case "a":
        var v29 uint32 = 497
        _ = v29
    case "hello":
        delete(v3, ("a" + "a"))
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        switch (fi30 ^ (fi30 >> ((676) & 63))) {
        case -46:
            if (873 >= 0) {
                if len(v27) > 0 {
                    v27[0] = fi30
                }
                delete(v5, "\n\t")
            } else {
                fmt.Println(v26, fi30, fi30)
                v2 = (^(v2))
            }
        case 1:
            if false {
                var v31 uint8 = 1
                _ = v31
            }
        case 956:
            var v32 uint = 0
            _ = v32
        default:
            var v33 int32 = v1
            _ = v33
        }
        _ = fi30
    }
    fmt.Println(v2, v26, v1)
    v34 := GenS1{F0: 128, F1: v1, F2: genFn1(1)}
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
