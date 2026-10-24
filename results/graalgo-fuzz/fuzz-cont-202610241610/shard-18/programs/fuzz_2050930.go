package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint; F2 uint64 }
type GenS3 struct { F0 int; F1 int32; F2 uint }
type GenS4 struct { F0 int32; F1 uint32; F2 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 uint; F1 uint32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16; F1 string }
func genFn0(p0 int32, p1 int16, p2 int8) int16 {
    p1 += -1
    v7 := GenS1{F0: -9223372036854775808}
    if (255 != 5) {
        if (!false) {
            p1 = p1
            v8 := []int{-1, 9223372036854775807, 42}
            _ = v8
        }
    }
    switch -1 {
    case -2147483648:
        p2 = 42
    case 9223372036854775807:
        if false {
            v9 := GenS2{F0: p0, F1: 42, F2: 1}
            v10 := GenS3{F0: 105, F1: -1, F2: 4294967295}
            p2 += -1
            _ = v10
            _ = v9
        }
    case 128:
        var v11 uint8 = 0
        _ = v11
    default:
        var v12 int16 = p1
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1(p0 int, p1 string) string {
    if true {
        for fi13 := 0; fi13 < 14; fi13++ {
            switch 128 {
            case 1:
                fmt.Println(fi13, p0)
            case -370:
                for fi14 := 0; fi14 < 13; fi14++ {
                    p1 = p1
                    if true {
                        v15 := GenS4{F0: 2147483647, F1: 473, F2: 0}
                        _ = v15
                    } else {
                        p1 = p1
                    }
                    fmt.Println(p0, p1)
                    _ = fi14
                }
            default:
                v16 := []int{}
                _ = v16
            }
            _ = fi13
        }
    }
    v17 := GenS2{F0: -182, F1: 0, F2: 4294967295}
    _ = p0
    _ = p1
    _ = v17
    return ("\n\t" + p1)
}

func main() {
    var v1 uint64 = 0
    var v2 string = "a"
    var v3 int8 = -1
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]bool)
    v6 := make(map[string]int)
    v3 += 1
    var v18 int64 = 931
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int64 = (v18 << ((0) & 63))
        if ((false && true) || true) {
            switch "a" {
            case "\n\t":
                switch (fi19 ^ (fi19 % ((-386) | 1))) {
                case 32767:
                    if (!(1 > int64(fi19))) {
                        v2 = ""
                    } else {
                        v21 := []int{}
                        _ = v21
                    }
                case 0:
                    v22 := GenS5{F0: v3}
                    _ = v22
                default:
                    var v23 string = (v2 + v2)
                    _ = v23
                }
            case "foo bar":
                if false {
                    v2 = (v2 + v2)
                    switch v2 {
                    case "":
                        v24 := GenS6{F0: v2, F1: -128}
                        _ = v24
                    case "foo bar":
                        var v25 string = v2
                        _ = v25
                    default:
                        fmt.Println(v3)
                    }
                } else {
                    v26 := GenS3{F0: (fi19 / (((fi19 & (-(fi19)))) | 1)), F1: 95, F2: 65535}
                    _ = v26
                }
            default:
                v18 += v20
            }
            var v27 uint64 = v1
            v4[2] = (-2147483648 != -1)
            _ = v27
        } else {
            v28 := GenS7{F0: 65535, F1: 2147483648}
            var v29 int = (fi19 * -223)
            v2 = v2
            _ = v28
            _ = v29
        }
        _ = fi19
        _ = v20
    }
    v30 := GenS8{F0: (v1 + v1)}
    if true {
        fmt.Println(v18, v1)
    }
    if (true && true) {
        v31 := GenS2{F0: 1, F1: 42, F2: (v1 * (v1 & 328))}
        if (32767 <= 0) {
            switch ((v2 + "foo bar") + ("a" + "hello")) {
            case "":
                v1 = 0
            case "foo bar":
                v4[2147483648] = false
            case "a":
                fmt.Println(v2, v18)
            default:
                v32 := GenS9{F0: 42, F1: v2}
                _ = v32
            }
        }
        _ = v31
    }
    v33 := GenS1{F0: (v18 * (v18 / ((v18) | 1)))}
    v34 := []int{}
    v2 = (v2 + "a")
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
