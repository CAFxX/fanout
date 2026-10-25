package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint; F2 string }
type GenS3 struct { F0 string; F1 string; F2 uint64 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 uint8; F2 uint32 }
type GenS9 struct { F0 int32; F1 int; F2 uint64 }
func genFn0(p0 uint, p1 int16, p2 int16) uint {
    switch "foo bar" {
    case "a":
        p0 -= p0
    case "hello":
        for fi6 := 0; fi6 < 2; fi6++ {
            if false {
                fmt.Println(p2, p2, p2)
                v7 := GenS1{F0: 873, F1: 71, F2: -9223372036854775808}
                _ = v7
            }
            if (42 < -1) {
                fmt.Println(p1, p0, p0)
                switch ("\n\t" + "a") {
                case "a":
                    var v8 int32 = 897
                    _ = v8
                case "foo bar":
                    var v9 string = ("hello" + "hello")
                    _ = v9
                case "":
                    p1 -= 0
                default:
                    var v10 string = ""
                    _ = v10
                }
                p2 -= (p1 << ((18446744073709551615) & 63))
            } else {
                if false {
                    p0 += (p0 - (p0 / ((583) | 1)))
                    p0 -= (p0 + (p0 | 255))
                }
                v11 := GenS2{F0: 2147483648, F1: 65535, F2: ((("foo bar" + "hello") + "hello") + "")}
                _ = v11
            }
            _ = fi6
        }
    }
    fmt.Println(p1, p1, p1)
    v12 := []int{-1, int(p1), 0}
    var v13 uint64 = 42
    switch -521 {
    case 2147483647:
        v12 = append(v12, 127)
    case 9223372036854775807:
        fmt.Println(p2)
    case 42:
        var v14 int32 = -1
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return p0
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    var v4 int8 = -1
    var v5 int64 = -1
    fmt.Println(v1)
    v15 := []int{}
    v16 := GenS3{F0: "", F1: "", F2: 1}
    fmt.Println(v1, v5, v1)
    switch 0 {
    case 32768:
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := GenS4{F0: v4, F1: "foo bar"}
            if true {
                v15 = append(v15, int(v1))
                var v19 bool = (-841 != -32768)
                _ = v19
            }
            _ = fi17
            _ = v18
        }
    }
    if true {
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 uint = 255
            fmt.Println(v5)
            _ = fi20
            _ = v21
        }
        var v22 uint64 = 18446744073709551615
        _ = v22
    } else {
        v23 := GenS5{F0: 0, F1: "\n\t"}
        v24 := GenS1{F0: 65535, F1: v4, F2: (v5 / (((v5 >> ((0) & 63))) | 1))}
        _ = v23
        _ = v24
    }
    v5 = (v5 / (((v5 >> ((42) & 63))) | 1))
    if (float32(((0.1 - (((1e10 / 1.0) * 1.5) * (3.14159 * -2.25))) + ((-2.25 * -2.25) - -2.25))) > ((1.5 / (1.5 - (0.0 / (1e10 + 0.0)))) + 0.0)) {
        var v25 float32 = 1.5
        var v26 int16 = 0
        v3[uint32(v1)] = 1
        _ = v25
        _ = v26
    }
    var v27 uint32 = 709
    v28 := GenS6{F0: (v5 % ((9223372036854775807) | 1))}
    fmt.Println(v27, v4)
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 int = -689
        v27 += v27
        fmt.Println(v5)
        _ = fi29
        _ = v30
    }
    v1 -= v1
    for fi31 := 0; fi31 < 18; fi31++ {
        var v32 int = (fi31 ^ fi31)
        for fi33 := 0; fi33 < 17; fi33++ {
            if true {
                v34 := GenS7{F0: 0}
                _ = v34
            }
            if (!true) {
                v15 = append(v15, fi31)
                v35 := v3[v27]
                v36 := GenS8{F0: int(v4), F1: 42, F2: v27}
                var v37 int8 = (v4 / ((v4) | 1))
                var v38 uint = 268
                v4 -= (-(v4))
                var v39 uint16 = 42
                _ = v35
                _ = v36
                _ = v37
                _ = v38
                _ = v39
            } else {
                v40 := GenS9{F0: 0, F1: fi31, F2: v1}
                v2[(v27 % (((^(v27))) | 1))] = -2147483648
                _ = v40
            }
            fmt.Println(fi33, v32, v32)
            v1 = v1
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
