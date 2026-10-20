package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 string; F2 int }
type GenS5 struct { F0 string }
func genFn0(p0 int16) int8 {
    fmt.Println(p0)
    switch -495 {
    case 700:
        v4 := []int{4294967295}
        _ = v4
    case 128:
        v5 := GenS1{F0: true, F1: 789}
        _ = v5
    case -9223372036854775808:
        var v6 bool = ((^(p0)) != p0)
        _ = v6
    }
    if (!((0.0 <= 0.0) && false)) {
        var v7 uint16 = 42
        v7 += v7
        v8 := GenS1{F0: (-1 > 32767), F1: 2147483648}
        _ = v7
        _ = v8
    }
    var v9 int64 = 9223372036854775807
    v10 := GenS1{F0: false, F1: 2147483647}
    _ = p0
    _ = v10
    _ = v9
    return -1
}
func genFn1(p0 int8, p1 string) uint8 {
    v11 := GenS2{F0: 4294967295}
    var v12 uint = 42
    if false {
        v13 := GenS2{F0: 2}
        if true {
            for fi14 := 0; fi14 < 7; fi14++ {
                if false {
                    v15 := []int{}
                    v16 := GenS3{F0: 2147483648}
                    _ = v15
                    _ = v16
                } else {
                    p0 -= p0
                    v17 := []int{9223372036854775807, -2147483648, (fi14 ^ (fi14 | fi14))}
                    _ = len(v17)
                    _ = v17
                }
                p0 -= p0
                switch p1 {
                case "":
                    if false {
                        v12 += v12
                        var v18 int8 = (p0 * (p0 << (((v12 + (v12 & 65535))) & 63)))
                        _ = v18
                    }
                }
                _ = fi14
            }
            p1 = ("" + p1)
        }
        var v19 uint = v12
        _ = v13
        _ = v19
    } else {
        var v20 int = -2147483648
        v12 = v12
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]bool)
    var v3 bool = (100 >= -47)
    switch 34 {
    case -870:
        var v21 uint32 = 2147483647
        _ = v21
    }
    switch -326 {
    case -371:
        if v3 {
            v22 := GenS1{F0: v3, F1: 976}
            _ = v22
        }
    case 128:
        v3 = v3
    case 32768:
        v1 = v1
    default:
        switch 9223372036854775807 {
        case -9223372036854775808:
            v23 := GenS4{F0: "hello", F1: ("a" + "\n\t"), F2: 752}
            _ = v23
        }
    }
    v24 := []int{9223372036854775807}
    if false {
        if true {
            for fi25 := 0; fi25 < 17; fi25++ {
                for fi26 := 0; fi26 < 7; fi26++ {
                    var v27 string = "foo bar"
                    v3 = ((365 != 746) && false)
                    var v28 int32 = 696
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v3 = false
                v1 -= (v1 & -32768)
                _ = fi25
            }
            var v29 uint = 42
            v30 := GenS2{F0: 0}
            _ = v29
            _ = v30
        } else {
            v3 = v3
            if false {
                for fi31 := 0; fi31 < 3; fi31++ {
                    var v32 string = "\n\t"
                    _ = fi31
                    _ = v32
                }
                v3 = (2147483647 <= -9223372036854775808)
            } else {
                for fi33 := 0; fi33 < 16; fi33++ {
                    var v34 float64 = 1e10
                    _ = fi33
                    _ = v34
                }
            }
            fmt.Println(v1, v1)
        }
        fmt.Println(v3)
    } else {
        v35 := GenS5{F0: "hello"}
        _ = v35
    }
    v1 += (-(v1))
    for fi36 := 0; fi36 < 13; fi36++ {
        v3 = false
        fmt.Println(fi36, v3, v1)
        _ = fi36
    }
    var v37 uint = 4294967295
    fmt.Println(v1, v3)
    fmt.Println(v3, v1, v37)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v37
}
