package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint16 }
func genFn0() uint16 {
    if false {
        v5 := GenS1{F0: 114, F1: -128, F2: "hello"}
        if ((42 == 9223372036854775807) && (((3.14159 <= ((-2.25 + 0.0) + float64(0.1))) && ("a" <= ("" + ("hello" + ("a" + "a"))))) && false)) {
            var v6 int16 = 312
            _ = v6
        } else {
            var v7 string = "hello"
            v7 = ("" + ("" + v7))
            switch v7 {
            case "foo bar":
                v8 := GenS2{F0: 1}
                _ = v8
            case "\n\t":
                var v9 int8 = 1
                _ = v9
            default:
                v10 := GenS3{F0: 1, F1: -2147483648, F2: 2}
                _ = v10
            }
            _ = v7
        }
        v11 := []int{68, 0, -169}
        _ = v11
        _ = v5
    } else {
        v12 := []int{42, -2147483648}
        var v13 uint16 = 0
        v14 := GenS1{F0: -979, F1: 0, F2: ("a" + "a")}
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 uint8 = 127
    v15 += 128
    v16 := GenS4{F0: (v15 & (v15 % ((v15) | 1))), F1: 32767}
    _ = v15
    _ = v16
    return 0
}
func genFn1(p0 int64, p1 uint, p2 uint32) uint16 {
    var v17 int = -9223372036854775808
    var v18 float32 = 1.5
    switch "\n\t" {
    case "foo bar":
        if ((false || true) || (false || ((v17 - (v17 % (((v17 % ((2147483647) | 1))) | 1))) <= -9223372036854775808))) {
            switch (v17 << ((40) & 63)) {
            case 714:
                if true {
                    var v19 uint = p1
                    v17 = 128
                    _ = v19
                } else {
                    switch (v17 ^ 42) {
                    case -61:
                        v20 := GenS4{F0: 255, F1: 1000}
                        _ = v20
                    }
                    v21 := []int{}
                    switch v17 {
                    case 128:
                        var v22 int16 = -32768
                        _ = v22
                    }
                    _ = v21
                }
            case -64:
                p1 += 1
            }
        } else {
            v23 := GenS5{F0: 131}
            fmt.Println(v18)
            _ = v23
        }
    case "hello":
        fmt.Println(p2)
    case "\n\t":
        p0 -= (p0 | p0)
    default:
        p0 -= -1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return genFn0()
}

func main() {
    var v1 uint32 = 0
    v2 := []int{0, -92, 42, 128}
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    var v24 uint32 = (v1 / (((v1 / ((v1) | 1))) | 1))
    fmt.Println(v24, v24)
    if true {
        delete(v4, 4294967295)
    } else {
        var v25 string = "hello"
        if (true && true) {
            var v26 int16 = 309
            _ = v26
        } else {
            switch -647 {
            case 128:
                v24 += (v24 >> ((255) & 63))
            default:
                var v27 int8 = -128
                _ = v27
            }
            v28 := GenS1{F0: 0, F1: -69, F2: (("\n\t" + ("a" + "hello")) + ("\n\t" + "\n\t"))}
            _ = v28
        }
        if len(v2) > 0 {
            v2[0] = 1
        }
        _ = v25
    }
    switch -718 {
    case 741:
        v24 -= v24
    }
    var v29 int64 = -850
    var v30 int8 = -60
    v30 = v30
    for fi31 := 0; fi31 < 12; fi31++ {
        fmt.Println(fi31, v24)
        if false {
            switch fi31 {
            case 127:
                v24 -= 42
            default:
                _ = len(v2)
            }
            fmt.Println(fi31, fi31)
        }
        _ = fi31
    }
    v32 := GenS1{F0: 1, F1: (v30 >> ((1) & 63)), F2: "foo bar"}
    fmt.Println(v30, v1)
    if (0 == 0) {
        var v33 uint8 = 127
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
}
