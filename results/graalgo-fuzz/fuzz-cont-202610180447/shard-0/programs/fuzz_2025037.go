package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint16; F1 bool; F2 int32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint16, p2 int32) int8 {
    if (!false) {
        p1 = (^(p1))
        p2 += (^(p2))
        fmt.Println(p0, p0)
    }
    fmt.Println(p2)
    fmt.Println(p0)
    var v4 uint64 = 433
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 uint16) int64 {
    var v5 bool = false
    v6 := GenS1{F0: (422 != 932), F1: 42}
    switch "foo bar" {
    case "foo bar":
        var v7 int16 = -1
        _ = v7
    case "\n\t":
        v8 := []int{2147483647}
        _ = v8
    case "":
        switch "foo bar" {
        case "hello":
            if (-873 > 1) {
                var v9 float32 = (0.1 + 0.1)
                _ = v9
            } else {
                if v5 {
                    v5 = false
                } else {
                    if ((("hello" + "hello") + "\n\t") > "foo bar") {
                        v10 := []int{4294967295}
                        fmt.Println(v5, p0)
                        _ = v10
                    }
                    fmt.Println(p0, v5)
                    v5 = (!true)
                }
            }
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 float64 = 0.1
    v2 := []string{"", "a", "hello"}
    v3 := make(map[int]string)
    v11 := GenS1{F0: false, F1: 1}
    var v12 int8 = 0
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 string = ""
        v12 = v12
        v14 = ((((("" + v14) + "hello") + v14) + v14) + "")
        _ = fi13
        _ = v14
    }
    var v15 string = "hello"
    v2 = append(v2, ((v15 + (("a" + "a") + ((v15 + "a") + (v15 + ("hello" + (v15 + v15)))))) + v15))
    v16 := GenS1{F0: (-2.25 <= 1e10), F1: 18446744073709551615}
    if len(v2) > 0 {
        v2[0] = (v15 + v15)
    }
    v2 = append(v2, ("\n\t" + v15))
    if (42 < 18446744073709551615) {
        _ = len(v3)
        if (genFn0(42, 2, -2147483648) != (v12 - (v12 + genFn0(362, 32767, 0)))) {
            if true {
                v17 := GenS1{F0: false, F1: 789}
                fmt.Println(v15, v15, v1)
                _ = v17
            }
            v18 := GenS1{F0: true, F1: 4294967295}
            switch 32767 {
            case -235:
                v1 -= v1
            default:
                switch 32767 {
                case 4294967295:
                    var v19 uint8 = 255
                    _ = v19
                case 32768:
                    v1 -= (v1 * 1.5)
                default:
                    var v20 uint64 = 4294967295
                    _ = v20
                }
            }
            _ = v18
        } else {
            var v21 int64 = genFn1(1)
            fmt.Println(v15, v12)
            v12 = v12
            _ = v21
        }
        var v22 uint16 = 1
        _ = v22
    } else {
        v23 := GenS2{F0: -93, F1: "\n\t"}
        _ = len(v2)
        _ = v23
    }
    fmt.Println(v1, v1, v1)
    v24 := GenS2{F0: 1, F1: v15}
    var v25 int32 = -447
    v26 := GenS2{F0: 434, F1: v15}
    for fi27 := 0; fi27 < 7; fi27++ {
        for fi28 := 0; fi28 < 12; fi28++ {
            if (genFn1(2) != 273) {
                v1 -= ((v1 + (v1 - v1)) / v1)
                v29 := GenS3{F0: 32767, F1: (59 >= 828), F2: 1}
                v25 -= (-(v25))
                fmt.Println(fi28, v12, v1)
                if (false || false) {
                    fmt.Println(v15, fi28)
                    if (false || (2147483648 != 2)) {
                        v30 := GenS2{F0: 9223372036854775807, F1: "foo bar"}
                        _ = v30
                    } else {
                        var v31 int = 416
                        _ = v31
                    }
                    var v32 uint16 = 0
                    v33 := GenS4{F0: 760}
                    var v34 uint32 = 137
                    if false {
                        v35 := GenS2{F0: fi28, F1: v15}
                        _ = v35
                    }
                    _ = v32
                    _ = v33
                    _ = v34
                } else {
                    v15 = v15
                    v15 = v15
                }
                _ = v29
            } else {
                _ = len(v3)
                v36 := v3[(fi27 >> ((4294967295) & 63))]
                var v37 string = ("hello" + "\n\t")
                _ = v36
                _ = v37
            }
            if true {
                v12 -= genFn0(0, 65535, v25)
                fmt.Println(v1)
            }
            _ = fi28
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
