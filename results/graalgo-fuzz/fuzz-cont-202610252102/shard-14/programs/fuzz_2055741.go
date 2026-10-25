package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint }
func genFn0(p0 bool) int64 {
    switch (("a" + "foo bar") + "foo bar") {
    case "\n\t":
        p0 = p0
    case "a":
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := GenS1{F0: (("" + "\n\t") + "foo bar"), F1: 255}
            v7 := GenS1{F0: ("" + "foo bar"), F1: 145}
            _ = fi5
            _ = v6
            _ = v7
        }
    case "":
        v8 := []int{-1, 128, 24, -2147483648}
        _ = v8
    default:
        p0 = (42 > 1)
    }
    var v9 uint = 4294967295
    v10 := []int{422, 602, 1}
    for fi11 := 0; fi11 < 17; fi11++ {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi11
    }
    _ = p0
    _ = v10
    _ = v9
    return 0
}
func genFn1(p0 int8, p1 int, p2 int32) bool {
    v12 := []int{}
    if (128 != 1) {
        for fi13 := 0; fi13 < 15; fi13++ {
            _ = len(v12)
            _ = fi13
        }
        switch "a" {
        case "hello":
            switch p1 {
            case -2147483648:
                switch "a" {
                case "":
                    _ = len(v12)
                case "hello":
                    for fi14 := 0; fi14 < 15; fi14++ {
                        p2 += -989
                        v15 := GenS1{F0: "hello", F1: 255}
                        v12 = append(v12, (fi14 - (fi14 & (-(p1)))))
                        _ = fi14
                        _ = v15
                    }
                default:
                    fmt.Println(p2)
                }
            case 2147483647:
                for fi16 := 0; fi16 < 14; fi16++ {
                    p1 -= (fi16 ^ p1)
                    _ = fi16
                }
            }
        case "\n\t":
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
        var v17 int16 = -1
        _ = v17
    } else {
        v12 = append(v12, p1)
    }
    var v18 uint16 = 42
    if len(v12) > 0 {
        v12[0] = p1
    }
    var v19 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v18
    _ = v19
    return true
}

func main() {
    var v1 string = "hello"
    var v2 bool = (1 <= -32768)
    var v3 uint8 = 255
    v4 := make(map[int]uint32)
    v20 := []int{-718}
    v3 += v3
    if (0 >= 2) {
        for fi21 := 0; fi21 < 17; fi21++ {
            _ = len(v4)
            _ = fi21
        }
    } else {
        fmt.Println(v2, v2, v3)
    }
    if v2 {
        v22 := GenS1{F0: v1, F1: 255}
        _ = v22
    }
    v3 = 127
    v2 = v2
    var v23 uint32 = 42
    _ = len(v4)
    if v2 {
        var v24 uint8 = v3
        v25 := GenS2{F0: 127, F1: 1, F2: 255}
        _ = v24
        _ = v25
    }
    var v26 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
