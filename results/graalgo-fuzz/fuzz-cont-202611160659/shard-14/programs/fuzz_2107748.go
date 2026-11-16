package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int; F1 bool; F2 bool }
func genFn0(p0 uint, p1 uint64) uint32 {
    var v7 uint = p0
    var v8 float32 = -2.25
    switch (("" + ("a" + "")) + "foo bar") {
    case "foo bar":
        if (((p0 * v7) < (p0 - (p0 >> ((255) & 63)))) && false) {
            v9 := GenS1{F0: 43, F1: 0}
            if (true && true) {
                switch (("a" + (("foo bar" + "foo bar") + ("" + ("a" + ("hello" + "a"))))) + "hello") {
                case "foo bar":
                    var v10 int = -9223372036854775808
                    _ = v10
                case "":
                    p1 -= p1
                case "a":
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v12 := GenS1{F0: 115, F1: (fi11 >> ((887) & 63))}
                        _ = fi11
                        _ = v12
                    }
                default:
                    var v13 uint8 = 205
                    _ = v13
                }
            }
            _ = v9
        } else {
            if true {
                switch -643 {
                case 32768:
                    v8 -= v8
                case 32767:
                    v14 := GenS2{F0: true}
                    _ = v14
                case -93:
                    p1 = 18446744073709551615
                default:
                    if true {
                        v15 := GenS1{F0: 255, F1: 4294967295}
                        _ = v15
                    }
                }
                switch 2147483647 {
                case 826:
                    if (!true) {
                        for fi16 := 0; fi16 < 6; fi16++ {
                            v17 := []int{-905}
                            _ = fi16
                            _ = v17
                        }
                        var v18 int64 = -1
                        _ = v18
                    }
                case 1:
                    var v19 string = ""
                    _ = v19
                case 0:
                    switch "" {
                    case "a":
                        for fi20 := 0; fi20 < 19; fi20++ {
                            v7 = v7
                            _ = fi20
                        }
                    case "\n\t":
                        p1 -= 2
                    }
                default:
                    v21 := GenS3{F0: 57, F1: v7}
                    _ = v21
                }
                var v22 int = 32768
                _ = v22
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483648
}
func genFn1(p0 uint32, p1 uint, p2 uint32) int32 {
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 float32 = (1.5 - (0.1 + 0.1))
        p1 = (p1 ^ 2)
        _ = fi23
        _ = v24
    }
    var v25 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 631
}

func main() {
    var v1 uint64 = 62
    var v2 uint32 = 2147483647
    var v3 float32 = 1.5
    v4 := make(map[int]uint32)
    var v5 string = ("hello" + ("a" + "\n\t"))
    v6 := []int{32767, 1}
    var v26 int8 = 42
    switch v5 {
    case "\n\t":
        for fi27 := 0; fi27 < 7; fi27++ {
            fmt.Println(v26, v5, v5)
            _ = fi27
        }
    case "hello":
        v26 = -103
    default:
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 float64 = (float64(0.1) - (float64(v3) - (3.14159 + (1e10 - 1.5))))
            _ = fi28
            _ = v29
        }
    }
    v3 = v3
    var v30 uint8 = 223
    var v31 int32 = 569
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 string = v5
        v34 := GenS2{F0: true}
        var v35 uint = 606
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    switch 1 {
    case 128:
        v30 -= (v30 % (((^(v30))) | 1))
    }
    switch v5 {
    case "hello":
        v5 = ((v5 + (((v5 + "") + "foo bar") + "a")) + v5)
    case "foo bar":
        var v36 uint32 = 248
        _ = v36
    default:
        switch 128 {
        case -9223372036854775808:
            if (true && (v31 <= (v31 / (((^(v31))) | 1)))) {
                _ = len(v4)
                for fi37 := 0; fi37 < 15; fi37++ {
                    var v38 int8 = (v26 & 100)
                    _ = len(v6)
                    delete(v4, 4294967295)
                    _ = fi37
                    _ = v38
                }
            } else {
                v39 := GenS4{F0: 285, F1: false, F2: (!((v31 % ((int32(v1)) | 1)) != v31))}
                _ = v39
            }
        case -890:
            v40 := GenS1{F0: (v30 & (v30 % (((v30 * v30)) | 1))), F1: -2147483648}
            _ = v40
        case -471:
            _ = len(v6)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
