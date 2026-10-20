package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int64; F1 int8; F2 string }
type GenS3 struct { F0 uint; F1 int32 }
func genFn0(p0 int32, p1 int32) int64 {
    for fi4 := 0; fi4 < 17; fi4++ {
        fmt.Println(p1, p0, p0)
        var v5 int16 = 82
        _ = fi4
        _ = v5
    }
    var v6 int16 = 32767
    var v7 int8 = 0
    p1 = 42
    var v8 uint8 = 127
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 uint32, p1 uint, p2 uint) bool {
    for fi9 := 0; fi9 < 13; fi9++ {
        if false {
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 uint8 = 127
                if ((p0 & (p0 >> ((246) & 63))) != p0) {
                    if (-128 >= 1) {
                        fmt.Println(fi9, p2)
                        v11 += v11
                        p1 = p1
                    }
                    var v12 int16 = 0
                    p0 += (p0 & (p0 ^ p0))
                    if true {
                        if false {
                            v12 -= v12
                        } else {
                            if ((3.14159 - 1e10) < 1e10) {
                                fmt.Println(p0, p2)
                                var v13 uint64 = 42
                                _ = v13
                            }
                        }
                        p1 -= p2
                        switch fi10 {
                        case -793:
                            var v14 int16 = (^(v12))
                            _ = v14
                        case 4294967295:
                            v12 += v12
                        }
                    }
                    _ = v12
                } else {
                    var v15 float32 = -2.25
                    switch (("\n\t" + "foo bar") + "foo bar") {
                    case "\n\t":
                        v16 := []int{int(p1), fi9}
                        _ = v16
                    }
                    fmt.Println(v11, p0)
                    _ = v15
                }
                _ = fi10
                _ = v11
            }
        }
        fmt.Println(p1)
        switch ("\n\t" + "a") {
        case "\n\t":
            p0 = 2147483647
        case "foo bar":
            v17 := GenS1{F0: 42, F1: -1}
            _ = v17
        case "":
            for fi18 := 0; fi18 < 3; fi18++ {
                fmt.Println(fi9)
                fmt.Println(fi9, p1)
                switch "a" {
                case "foo bar":
                    v19 := GenS2{F0: 9223372036854775807, F1: -128, F2: "a"}
                    _ = v19
                }
                fmt.Println(fi9, fi9, fi18)
                p1 -= 2
                var v20 int32 = 1
                v21 := GenS3{F0: p2, F1: v20}
                _ = fi18
                _ = v20
                _ = v21
            }
        default:
            p2 += 0
        }
        _ = fi9
    }
    p1 = p1
    var v22 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return false
}

func main() {
    var v1 int = -703
    var v2 int64 = 577
    var v3 uint64 = 0
    switch (v1 << ((936) & 63)) {
    case 520:
        v1 = v1
    case 2147483647:
        if (127 != 1) {
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 int8 = -128
                _ = fi23
                _ = v24
            }
            var v25 uint64 = v3
            _ = v25
        }
    }
    v26 := []int{(v1 | v1), v1, v1}
    var v27 float64 = -2.25
    var v28 uint = 18446744073709551615
    fmt.Println(v2)
    v29 := GenS2{F0: v2, F1: 42, F2: "foo bar"}
    var v30 uint8 = 88
    if (1 <= -538) {
        v3 -= (v3 + v3)
        for fi31 := 0; fi31 < 18; fi31++ {
            for fi32 := 0; fi32 < 8; fi32++ {
                v28 = (^(v28))
                var v33 uint8 = 127
                var v34 bool = ((^(v28)) > v28)
                var v35 string = "hello"
                v36 := GenS2{F0: (v2 | (-(v2))), F1: 100, F2: v35}
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
                _ = v36
            }
            _ = fi31
        }
    }
    v2 = v2
    fmt.Println(v30)
    switch "\n\t" {
    case "\n\t":
        var v37 string = ("a" + "")
        _ = v37
    default:
        fmt.Println(v27, v28)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
