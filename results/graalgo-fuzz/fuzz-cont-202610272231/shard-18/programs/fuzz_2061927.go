package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 string }
func genFn0(p0 bool, p1 uint64, p2 int8) uint64 {
    if (9223372036854775807 > -868) {
        if (p0 && p0) {
            p0 = p0
        } else {
            if ((p2 * 0) < p2) {
                for fi5 := 0; fi5 < 16; fi5++ {
                    switch (fi5 - (fi5 - 32768)) {
                    case 32768:
                        p2 = -1
                    }
                    v6 := GenS1{F0: 471}
                    if p0 {
                        p1 = p1
                        for fi7 := 0; fi7 < 9; fi7++ {
                            var v8 int8 = (p2 - -128)
                            if p0 {
                                var v9 uint16 = 2
                                fmt.Println(fi7, fi7, v8)
                                v10 := []int{32767, 761, fi7}
                                _ = v10
                                _ = v9
                            } else {
                                p0 = true
                                v11 := GenS2{F0: p0, F1: 0}
                                var v12 int = (fi7 << ((255) & 63))
                                _ = v11
                                _ = v12
                            }
                            if (0.0 < ((3.14159 * float32(((1.5 * 0.1) * float64(1.5)))) * -2.25)) {
                                var v13 uint64 = 640
                                if (uint8(p1) != 1) {
                                    if p0 {
                                        v14 := []int{}
                                        p2 = 100
                                        _ = v14
                                    } else {
                                        p2 -= 1
                                        switch "hello" {
                                        case "a":
                                            var v15 int8 = p2
                                            _ = v15
                                        case "hello":
                                            v16 := GenS3{F0: 2, F1: 65535}
                                            _ = v16
                                        case "foo bar":
                                            var v17 uint32 = 494
                                            _ = v17
                                        default:
                                            fmt.Println(p2)
                                        }
                                    }
                                } else {
                                    var v18 float64 = 1.5
                                    v19 := []int{(-(fi7))}
                                    _ = v18
                                    _ = v19
                                }
                                if p0 {
                                    p0 = true
                                    v20 := GenS3{F0: 469, F1: 2}
                                    var v21 int32 = 2147483647
                                    _ = v20
                                    _ = v21
                                }
                                _ = v13
                            }
                            _ = fi7
                            _ = v8
                        }
                        for fi22 := 0; fi22 < 14; fi22++ {
                            var v23 uint = 198
                            _ = fi22
                            _ = v23
                        }
                    }
                    _ = fi5
                    _ = v6
                }
                var v24 uint16 = uint16(p1)
                var v25 float32 = 1e10
                _ = v24
                _ = v25
            }
            v26 := GenS1{F0: -322}
            _ = v26
        }
    } else {
        p2 -= (p2 + (p2 * (p2 | p2)))
    }
    p1 += p1
    var v27 int32 = 399
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 171
}
func genFn1() int {
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 bool = (!((!(false || (((!(!false)) && ((("foo bar" + "hello") + (("hello" + (("hello" + "foo bar") + "hello")) + "hello")) > ("\n\t" + ""))) && true))) || true))
        _ = fi28
        _ = v29
    }
    var v30 float32 = (-2.25 / 1.0)
    var v31 uint64 = genFn0(false, 380, 100)
    switch "hello" {
    case "\n\t":
        var v32 int16 = 42
        _ = v32
    }
    _ = v30
    _ = v31
    return -9223372036854775808
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]string)
    var v3 uint64 = 0
    var v4 float32 = 0.1
    if true {
        var v33 int = -449
        v34 := v2[v33]
        _ = v33
        _ = v34
    }
    var v35 int8 = (v1 - (v1 % ((0) | 1)))
    if ("\n\t" >= "\n\t") {
        v4 -= (v4 - v4)
        var v36 int16 = 1000
        _ = v36
    }
    v37 := GenS4{F0: ("foo bar" + ("\n\t" + (("a" + "a") + "")))}
    v38 := []int{-284, genFn1()}
    for fi39 := 0; fi39 < 7; fi39++ {
        v4 -= (3.14159 - 0.1)
        v4 = float32((0.1 + 3.14159))
        var v40 uint64 = v3
        _ = fi39
        _ = v40
    }
    fmt.Println(v4, v1)
    var v41 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v41
}
