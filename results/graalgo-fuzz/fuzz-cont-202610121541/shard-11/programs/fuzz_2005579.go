package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int16; F1 int64; F2 string }
func genFn0(p0 uint32) int {
    if false {
        p0 += p0
        var v4 int64 = 9223372036854775807
        _ = v4
    } else {
        p0 -= (p0 << ((42) & 63))
        v5 := []int{-964}
        _ = v5
    }
    fmt.Println(p0)
    var v6 int = 128
    switch "foo bar" {
    case "hello":
        if false {
            for fi7 := 0; fi7 < 7; fi7++ {
                switch v6 {
                case 1:
                    fmt.Println(p0)
                case -366:
                    p0 = p0
                default:
                    var v8 bool = (true || (-683 < -2147483648))
                    _ = v8
                }
                var v9 int16 = 734
                _ = fi7
                _ = v9
            }
            v10 := GenS1{F0: 517, F1: "hello"}
            _ = v10
        } else {
            var v11 int32 = -2147483648
            v12 := GenS1{F0: 442, F1: "foo bar"}
            v13 := GenS1{F0: 18446744073709551615, F1: ("a" + ((("\n\t" + "foo bar") + "a") + ""))}
            _ = v11
            _ = v12
            _ = v13
        }
    case "foo bar":
        fmt.Println(v6)
    default:
        fmt.Println(v6)
    }
    v6 -= v6
    _ = p0
    _ = v6
    return (v6 + v6)
}
func genFn1(p0 uint8, p1 uint32) int16 {
    v14 := GenS1{F0: 395, F1: ("" + "hello")}
    p0 = (p0 >> ((888) & 63))
    fmt.Println(p1)
    var v15 uint16 = 65535
    if (967 != p1) {
        var v16 float64 = (((float64((0.0 * 0.1)) / 1e10) + (3.14159 + (3.14159 * (float64(-2.25) * (float64((1.5 - 0.0)) * 3.14159))))) * 1.5)
        _ = v16
    } else {
        if false {
            p0 -= p0
            for fi17 := 0; fi17 < 14; fi17++ {
                fmt.Println(fi17, fi17, p0)
                fmt.Println(fi17, p0)
                _ = fi17
            }
            var v18 int32 = 42
            _ = v18
        } else {
            fmt.Println(p0, p1)
        }
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return -32768
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]bool)
    v3 := []int{}
    v1 += (^(v1))
    v1 = v1
    var v19 uint64 = 2
    var v20 int16 = 42
    switch -128 {
    case 42:
        v19 = 1
    case 9223372036854775807:
        if false {
            var v21 uint32 = (^(v1))
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS2{F0: 53}
                _ = fi22
                _ = v23
            }
            fmt.Println(v20, v1, v20)
            _ = v21
        }
    default:
        v24 := GenS3{F0: 960, F1: (v19 / ((28) | 1))}
        _ = v24
    }
    if (!true) {
        var v25 uint = 0
        _ = v25
    } else {
        v26 := GenS4{F0: (false || (!false)), F1: ""}
        if (false && true) {
            var v27 bool = (!(-1 != -128))
            v28 := GenS5{F0: 32767, F1: int64(v20), F2: ("foo bar" + ("foo bar" + "foo bar"))}
            _ = v27
            _ = v28
        }
        var v29 float32 = (-2.25 - ((3.14159 / float32((3.14159 + (1.5 * (-2.25 * 1.5))))) - float32(1e10)))
        _ = v26
        _ = v29
    }
    if (-69 <= 100) {
        var v30 int8 = 100
        _ = v30
    } else {
        if false {
            v1 = (v1 + 42)
        } else {
            var v31 uint64 = (v19 - (v19 / ((v19) | 1)))
            for fi32 := 0; fi32 < 9; fi32++ {
                _ = len(v3)
                _ = fi32
            }
            var v33 int8 = 42
            _ = v31
            _ = v33
        }
        v34 := GenS4{F0: true, F1: "\n\t"}
        _ = v34
    }
    var v35 uint = 65535
    delete(v2, -131)
    v36 := GenS5{F0: v20, F1: 469, F2: "hello"}
    switch "\n\t" {
    case "foo bar":
        v1 -= v1
    default:
        fmt.Println(v35, v35)
    }
    var v37 int = 0
    for fi38 := 0; fi38 < 13; fi38++ {
        if true {
            var v39 int16 = (v20 & (v20 * v20))
            if false {
                v39 = v20
                fmt.Println(v39)
            }
            delete(v2, v37)
            _ = v39
        }
        var v40 bool = false
        v37 = (v37 & genFn0(v1))
        _ = fi38
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
