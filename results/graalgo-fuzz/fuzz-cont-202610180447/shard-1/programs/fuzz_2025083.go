package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 string; F2 uint8 }
type GenS5 struct { F0 int64; F1 uint }
func genFn0(p0 int, p1 int32, p2 int64) uint64 {
    p2 = int64(p1)
    p0 -= p0
    v5 := []int{(p0 | -1), p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    var v1 string = "hello"
    var v2 string = v1
    var v3 int64 = 1
    v4 := make(map[int]string)
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{(^(fi6)), 288}
        v3 += (v3 ^ v3)
        _ = fi6
        _ = v7
    }
    if true {
        if (356 >= genFn0(457, -1, v3)) {
            for fi8 := 0; fi8 < 14; fi8++ {
                v3 = (^(v3))
                _ = fi8
            }
            v9 := GenS1{F0: 1, F1: 65535}
            _ = v9
        } else {
            v10 := GenS2{F0: 179, F1: 104}
            switch "\n\t" {
            case "hello":
                if true {
                    v2 = v2
                    for fi11 := 0; fi11 < 16; fi11++ {
                        v3 = v3
                        var v12 int = fi11
                        var v13 uint = 255
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                } else {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        fmt.Println(v2, v3, fi14)
                        v2 = ((("" + "foo bar") + "\n\t") + (v2 + v1))
                        var v15 string = "hello"
                        _ = fi14
                        _ = v15
                    }
                    v4[32768] = ((("foo bar" + v2) + v1) + ((v1 + v1) + ""))
                }
            case "\n\t":
                v16 := []int{}
                _ = v16
            case "a":
                var v17 uint = 0
                _ = v17
            }
            v3 -= v3
            _ = v10
        }
        switch v2 {
        case "foo bar":
            v3 = v3
        case "\n\t":
            v1 = ""
        }
        var v18 uint32 = 0
        _ = v18
    } else {
        var v19 uint8 = 0
        v20 := GenS3{F0: (-(v3))}
        v21 := GenS4{F0: v19, F1: "hello", F2: v19}
        _ = v19
        _ = v20
        _ = v21
    }
    fmt.Println(v1, v2)
    for fi22 := 0; fi22 < 1; fi22++ {
        switch fi22 {
        case -542:
            v23 := GenS4{F0: 255, F1: "foo bar", F2: 196}
            _ = v23
        default:
            if (-128 > 0) {
                v2 = v2
                v1 = v1
            }
        }
        _ = fi22
    }
    v1 = "foo bar"
    v24 := []int{-102, 934, -2147483648, -618}
    v25 := GenS5{F0: (v3 - (-(v3))), F1: 4294967295}
    switch v1 {
    case "\n\t":
        fmt.Println(v1)
    case "":
        var v26 float64 = 0.0
        _ = v26
    case "foo bar":
        if true {
            v1 = (v2 + v2)
            var v27 int8 = 100
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            _ = v27
        } else {
            if false {
                for fi28 := 0; fi28 < 15; fi28++ {
                    var v29 int32 = -1
                    v29 = (v29 << ((0) & 63))
                    v3 = (v3 & 731)
                    _ = fi28
                    _ = v29
                }
                v1 = v2
            } else {
                if true {
                    v2 = ("foo bar" + v1)
                    if (false && (!false)) {
                        if ((false && false) && (true || true)) {
                            v30 := v4[42]
                            _ = v30
                        }
                        fmt.Println(v2)
                        _ = len(v24)
                    }
                    v1 = (v2 + (v2 + "\n\t"))
                }
            }
            fmt.Println(v1)
            v2 = v2
        }
    default:
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 string = (v2 + "a")
            _ = fi31
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
