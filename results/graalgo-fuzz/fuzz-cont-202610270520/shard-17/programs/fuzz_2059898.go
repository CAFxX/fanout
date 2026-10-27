package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint16; F2 uint }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 uint }
func genFn0() string {
    var v4 bool = (!(!((1 == 18446744073709551615) || (!true))))
    switch "\n\t" {
    case "hello":
        v5 := []int{42}
        _ = v5
    case "foo bar":
        v6 := []int{-2147483648, 4294967295}
        _ = v6
    case "\n\t":
        v4 = (!v4)
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 bool = v4
        var v9 int = fi7
        v10 := GenS1{F0: 537}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if (!v4) {
        switch 32768 {
        case 9223372036854775807:
            v11 := GenS1{F0: 2}
            _ = v11
        }
        if (-1 != 0) {
            var v12 string = "a"
            v13 := GenS2{F0: 0, F1: 2, F2: 255}
            _ = v12
            _ = v13
        }
        switch "hello" {
        case "foo bar":
            for fi14 := 0; fi14 < 4; fi14++ {
                if false {
                    var v15 float64 = (3.14159 - -2.25)
                    _ = v15
                }
                v16 := GenS2{F0: 1, F1: 2, F2: 4294967295}
                v4 = (v4 && ("" != ""))
                _ = fi14
                _ = v16
            }
        case "\n\t":
            fmt.Println(v4)
        }
    }
    v17 := GenS3{F0: 945, F1: -2147483648}
    _ = v17
    _ = v4
    return "a"
}
func genFn1(p0 uint, p1 int32, p2 uint16) uint32 {
    v18 := []int{614, -363, -2147483648, 1}
    var v19 uint8 = 255
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 uint64 = 916
        var v22 int64 = 0
        v23 := GenS2{F0: 1, F1: p2, F2: p0}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := GenS4{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v24
    return 327
}

func main() {
    var v1 int64 = 381
    v2 := make(map[uint32]string)
    var v3 uint64 = 691
    v25 := GenS1{F0: 42}
    if false {
        v3 = (^(v3))
        var v26 uint16 = 255
        _ = v26
    }
    switch 2147483647 {
    case 9223372036854775807:
        var v27 bool = (!false)
        _ = v27
    case 0:
        v28 := GenS5{F0: 1000, F1: 932}
        _ = v28
    case -9223372036854775808:
        var v29 int8 = 125
        _ = v29
    default:
        v1 = v1
    }
    _ = len(v2)
    v3 += 0
    v1 += v1
    var v30 bool = true
    var v31 int64 = -1
    fmt.Println(v1, v3)
    var v32 uint16 = 32767
    fmt.Println(v1, v30, v32)
    fmt.Println(v30, v1)
    v32 = 42
    switch -1 {
    case 128:
        v33 := v2[2147483647]
        _ = v33
    case 1:
        var v34 int64 = (^(v1))
        _ = v34
    case 2147483647:
        for fi35 := 0; fi35 < 15; fi35++ {
            for fi36 := 0; fi36 < 7; fi36++ {
                v37 := GenS1{F0: 1}
                var v38 bool = v30
                var v39 string = ("a" + "foo bar")
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
            }
            if (2 < 42) {
                v3 = (v3 * v3)
                fmt.Println(fi35, v30)
                v40 := GenS2{F0: 0, F1: (v32 >> ((255) & 63)), F2: 0}
                _ = v40
            }
            _ = fi35
        }
    }
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
