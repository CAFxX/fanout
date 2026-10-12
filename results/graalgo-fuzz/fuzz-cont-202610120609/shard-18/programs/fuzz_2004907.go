package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint8; F1 int8; F2 int64 }
type GenS4 struct { F0 bool; F1 int64; F2 string }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 int32 }
func genFn0(p0 int16, p1 int) int16 {
    v6 := []int{p1}
    v7 := GenS1{F0: -9223372036854775808, F1: ("a" + "\n\t")}
    if len(v6) > 0 {
        v6[0] = 42
    }
    switch "\n\t" {
    case "a":
        v8 := GenS2{F0: (p1 - 32767), F1: 2147483647}
        _ = v8
    case "\n\t":
        if true {
            p1 = p1
            if false {
                if (true || false) {
                    if (!true) {
                        v9 := GenS1{F0: 1, F1: ""}
                        _ = v9
                    } else {
                        switch "a" {
                        case "hello":
                            var v10 int64 = 0
                            _ = v10
                        case "":
                            if (851 > 18446744073709551615) {
                                _ = len(v6)
                            } else {
                                switch (p1 * (p1 + p1)) {
                                case 374:
                                    _ = len(v6)
                                default:
                                    p1 = 9223372036854775807
                                }
                            }
                        case "a":
                            if true {
                                p1 = 32768
                            }
                        default:
                            var v11 string = (("foo bar" + "\n\t") + "")
                            _ = v11
                        }
                    }
                }
                v12 := GenS1{F0: 0, F1: "\n\t"}
                var v13 float32 = ((1.5 * 1e10) / 1e10)
                _ = v12
                _ = v13
            }
            p1 = p1
        } else {
            var v14 int = -329
            v14 += (^(p1))
            v15 := GenS1{F0: -1, F1: ""}
            _ = v14
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 1; fi16++ {
            p1 = (fi16 >> ((255) & 63))
            switch (p1 - (fi16 / (((-(p1))) | 1))) {
            case 32768:
                p1 = fi16
            case 332:
                p1 = p1
            case -938:
                fmt.Println(p0)
            default:
                p0 = p0
            }
            var v17 float64 = -2.25
            _ = fi16
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}

func main() {
    v1 := []int{-829, 208, 42, -2147483648}
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v5 uint16 = 522
    var v18 string = ("\n\t" + "")
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 bool = true
        _ = fi19
        _ = v20
    }
    v18 = (("foo bar" + (v18 + v18)) + "\n\t")
    var v21 uint32 = 2147483648
    v22 := GenS3{F0: 1, F1: 78, F2: 57}
    v23 := GenS4{F0: false, F1: -496, F2: v18}
    v24 := GenS5{F0: 0, F1: 207}
    var v25 uint32 = uint32(v5)
    var v26 uint32 = 2147483647
    v27 := GenS5{F0: 32767, F1: 59}
    v21 += v25
    v21 = v21
    v28 := GenS6{F0: 0}
    v21 -= 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
