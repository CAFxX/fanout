package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int16; F2 uint32 }
func genFn0(p0 int, p1 uint8) string {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{p0, fi5}
        switch "hello" {
        case "a":
            v7 := GenS1{F0: 295}
            _ = v7
        case "\n\t":
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS1{F0: 0}
                switch fi8 {
                case 633:
                    fmt.Println(fi8)
                case 747:
                    v10 := GenS1{F0: 977}
                    _ = v10
                case -57:
                    if len(v6) > 0 {
                        v6[0] = (-(p0))
                    }
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                fmt.Println(p0, p0)
                _ = fi8
                _ = v9
            }
        case "hello":
            if (!true) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                p0 += fi5
                p1 = (p1 + p1)
            } else {
                if false {
                    var v11 bool = true
                    _ = v11
                }
            }
        }
        _ = len(v6)
        _ = fi5
        _ = v6
    }
    var v12 uint32 = 251
    for fi13 := 0; fi13 < 6; fi13++ {
        v12 += v12
        if (9223372036854775807 != 905) {
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 int8 = 1
                var v16 int16 = -1
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        _ = fi13
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v12
    return ("\n\t" + "a")
}

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{}
    var v3 int16 = 32767
    var v4 float64 = (-2.25 * float64(0.0))
    var v17 bool = ((!false) || true)
    var v18 float64 = (1e10 - ((0.0 + 0.1) / v4))
    v2 = append(v2, uint32(v1))
    fmt.Println(v18)
    fmt.Println(v1)
    var v19 uint = 489
    if (42 < 0) {
        v20 := GenS2{F0: 660}
        v18 -= v18
        _ = v20
    } else {
        v21 := GenS2{F0: 42}
        fmt.Println(v17, v4, v19)
        v18 -= v18
        _ = v21
    }
    switch "" {
    case "":
        v4 = (v18 / float64(3.14159))
    case "foo bar":
        v19 -= (^(v19))
    }
    if (1 <= 42) {
        if true {
            v22 := GenS2{F0: -2147483648}
            for fi23 := 0; fi23 < 1; fi23++ {
                if v17 {
                    v19 = (v19 & (v19 ^ v19))
                    if v17 {
                        var v24 string = ("a" + "hello")
                        v25 := GenS2{F0: 1}
                        fmt.Println(v1)
                        _ = v24
                        _ = v25
                    }
                } else {
                    v1 = (v1 - 512)
                    _ = len(v2)
                    var v26 int16 = v3
                    _ = v26
                }
                v19 += v19
                _ = fi23
            }
            _ = v22
        }
        if v17 {
            _ = len(v2)
            var v27 uint8 = 19
            _ = v27
        } else {
            v2 = append(v2, 2)
            v28 := GenS2{F0: 298}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        if true {
            v30 := GenS1{F0: 42}
            var v31 bool = (2147483647 >= (fi29 >> ((v19) & 63)))
            _ = v30
            _ = v31
        }
        _ = fi29
    }
    fmt.Println(v1)
    v32 := GenS3{F0: ("hello" + "hello"), F1: v3, F2: 4294967295}
    v17 = (267 >= 639)
    fmt.Println(v17, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
