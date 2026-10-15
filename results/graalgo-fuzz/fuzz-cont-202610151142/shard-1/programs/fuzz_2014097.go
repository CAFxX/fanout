package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int64; F1 bool; F2 string }
func genFn0(p0 int64, p1 uint64) uint {
    var v7 int8 = -128
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 bool = (true && false)
        v9 = (((float32(((((-2.25 - 0.1) / 1.5) + float64(0.0)) * 0.0)) * 1e10) - 3.14159) != 0.1)
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{4294967295}
    for fi12 := 0; fi12 < 13; fi12++ {
        if true {
            var v13 int8 = (-(v7))
            v14 := GenS1{F0: false}
            var v15 float32 = float32(-2.25)
            _ = v13
            _ = v14
            _ = v15
        }
        var v16 string = ""
        _ = fi12
        _ = v16
    }
    var v17 string = ""
    _ = p0
    _ = p1
    _ = v11
    _ = v17
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 string, p2 uint64) uint16 {
    if (!(p1 == "\n\t")) {
        fmt.Println(p2)
        var v18 int32 = -1
        var v19 int8 = 100
        _ = v18
        _ = v19
    }
    p1 = p1
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 100
    var v3 float32 = 1.5
    v4 := make(map[uint32]bool)
    var v5 bool = false
    var v6 uint64 = 4294967295
    var v20 int64 = -1
    v21 := v1[-2147483648]
    v5 = ((false || ((-(v20)) > 0)) || (3.14159 < (-2.25 * ((0.1 * 3.14159) - 0.0))))
    v22 := []int{-2147483648, 1}
    if (-1 == -1) {
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        if v5 {
            v23 := GenS2{F0: 42, F1: 1}
            v24 := GenS3{F0: v20, F1: v5, F2: ("foo bar" + "")}
            v25 := GenS1{F0: ((("foo bar" + "") + "a") != "a")}
            _ = v23
            _ = v24
            _ = v25
        }
        v21 += 505
    } else {
        var v26 uint = 1
        fmt.Println(v6, v6)
        _ = v26
    }
    var v27 uint32 = (v21 >> ((2) & 63))
    v20 -= 464
    fmt.Println(v20, v6)
    _ = len(v1)
    switch ("foo bar" + "") {
    case "\n\t":
        for fi28 := 0; fi28 < 14; fi28++ {
            if (v5 || v5) {
                v27 -= (^(v21))
            } else {
                switch fi28 {
                case -614:
                    var v29 int = (-(fi28))
                    _ = v29
                case 9223372036854775807:
                    var v30 uint16 = 42
                    _ = v30
                }
                v20 = v20
            }
            v4[(v27 + v27)] = (0 > 127)
            _ = fi28
        }
    case "foo bar":
        if (!false) {
            v22 = append(v22, 1)
        } else {
            if len(v22) > 0 {
                v22[0] = 4294967295
            }
            for fi31 := 0; fi31 < 7; fi31++ {
                if (!(v3 < v3)) {
                    v6 = v6
                    var v32 int = (-(fi31))
                    fmt.Println(v21)
                    _ = v32
                }
                _ = fi31
            }
        }
    default:
        switch ("hello" + "\n\t") {
        case "a":
            var v33 uint16 = 42
            _ = v33
        case "foo bar":
            for fi34 := 0; fi34 < 18; fi34++ {
                for fi35 := 0; fi35 < 12; fi35++ {
                    fmt.Println(v27, v3)
                    delete(v4, v21)
                    var v36 int8 = v2
                    _ = fi35
                    _ = v36
                }
                for fi37 := 0; fi37 < 15; fi37++ {
                    if v5 {
                        var v38 float32 = 0.0
                        _ = v38
                    }
                    _ = fi37
                }
                _ = fi34
            }
        }
    }
    var v39 string = ("\n\t" + "a")
    v6 = 2
    var v40 int = -259
    v39 = v39
    fmt.Println(v6, v5, v6)
    var v41 uint8 = 60
    v40 -= (v40 - 2147483647)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
