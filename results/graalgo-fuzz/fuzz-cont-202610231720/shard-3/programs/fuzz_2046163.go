package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint8, p1 uint8, p2 uint64) bool {
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 uint32 = 2147483648
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        p2 = p2
        v7 := GenS1{F0: 0}
        v8 := GenS1{F0: 0}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        if (fi9 == fi9) {
            var v10 float64 = (0.0 * 0.1)
            v11 := GenS1{F0: 18446744073709551615}
            _ = v10
            _ = v11
        } else {
            p2 -= p2
            p0 = p0
        }
        if false {
            v12 := []int{4294967295, -1, (fi9 * (fi9 | (fi9 + 2147483647)))}
            v13 := GenS2{F0: -363, F1: -34}
            _ = v12
            _ = v13
        } else {
            var v14 uint = 65535
            var v15 uint64 = 42
            var v16 bool = true
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi9
    }
    v17 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return false
}

func main() {
    var v1 int16 = 585
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    for fi18 := 0; fi18 < 3; fi18++ {
        if true {
            fmt.Println(fi18)
            v19 := []int{(fi18 >> ((158) & 63))}
            _ = v19
        }
        var v20 float32 = 1e10
        _ = fi18
        _ = v20
    }
    if true {
        var v21 uint64 = 18446744073709551615
        _ = v21
    }
    if (561 < 885) {
        _ = len(v2)
        v1 = (v1 | v1)
        switch ("foo bar" + ("" + (("a" + ("hello" + "hello")) + ("\n\t" + ("" + ("a" + "\n\t")))))) {
        case "foo bar":
            fmt.Println(v1)
        case "a":
            for fi22 := 0; fi22 < 11; fi22++ {
                v23 := []int{(fi22 << ((255) & 63)), (-(fi22)), (-(fi22))}
                _ = fi22
                _ = v23
            }
        default:
            switch "foo bar" {
            case "hello":
                fmt.Println(v1)
            case "foo bar":
                v2[32767] = 4294967295
            case "\n\t":
                v24 := GenS3{F0: 4294967295, F1: 0}
                _ = v24
            }
        }
    } else {
        v1 += 42
        if ("hello" <= (("hello" + ("" + "a")) + "foo bar")) {
            var v25 uint64 = 998
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 string = ""
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 5; fi28++ {
                if (((("a" + "") + "") + (("foo bar" + "") + "foo bar")) == "\n\t") {
                    v25 += 42
                    v1 = v1
                    v29 := GenS2{F0: 9223372036854775807, F1: v1}
                    _ = v29
                }
                _ = fi28
            }
            _ = v25
        } else {
            v30 := GenS4{F0: 2, F1: v1}
            for fi31 := 0; fi31 < 19; fi31++ {
                v1 = v1
                v1 = v1
                switch (fi31 >> ((4294967295) & 63)) {
                case 9223372036854775807:
                    v1 = (v1 - v1)
                case 381:
                    for fi32 := 0; fi32 < 6; fi32++ {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v1, fi31)
                        case "\n\t":
                            fmt.Println(v1, fi31)
                        }
                        _ = fi32
                    }
                }
                _ = fi31
            }
            _ = v30
        }
    }
    v33 := GenS3{F0: 4294967295, F1: -2147483648}
    v34 := []int{42, 128, 943}
    if len(v34) > 0 {
        v34[0] = -9223372036854775808
    }
    var v35 bool = true
    v1 += -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
