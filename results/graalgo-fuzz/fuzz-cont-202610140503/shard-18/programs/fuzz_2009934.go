package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0(p0 uint, p1 uint64, p2 uint32) bool {
    var v6 uint64 = 400
    switch 9223372036854775807 {
    case 1:
        fmt.Println(p1, v6)
    case -507:
        v7 := GenS1{F0: (p0 < (p0 - 0)), F1: p0, F2: 605}
        _ = v7
    default:
        p1 += 42
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return true
}
func genFn1(p0 uint8, p1 string) int8 {
    p0 = p0
    v8 := []int{}
    _ = len(v8)
    _ = p0
    _ = p1
    _ = v8
    return -1
}

func main() {
    var v1 uint = 589
    var v2 int8 = 0
    var v3 uint16 = 838
    var v4 int = 42
    v5 := make(map[int]uint32)
    v1 += (v1 * v1)
    switch "foo bar" {
    case "hello":
        var v9 int64 = -1
        _ = v9
    default:
        v10 := []int{v4, (v4 % ((v4) | 1)), v4}
        _ = v10
    }
    var v11 float32 = (-2.25 / ((0.0 * float32(0.1)) + 1e10))
    v12 := []int{128, 1, (v4 * 753), v4}
    switch "\n\t" {
    case "foo bar":
        var v13 uint32 = 654
        _ = v13
    default:
        v5[v4] = 4294967295
    }
    v11 = v11
    for fi14 := 0; fi14 < 11; fi14++ {
        v12 = append(v12, 2147483647)
        var v15 uint64 = 42
        for fi16 := 0; fi16 < 15; fi16++ {
            if ((("" + "hello") >= "hello") || (!true)) {
                var v17 int16 = 42
                var v18 uint = v1
                _ = v17
                _ = v18
            }
            v19 := GenS1{F0: false, F1: 564, F2: -715}
            var v20 uint64 = v15
            _ = fi16
            _ = v19
            _ = v20
        }
        _ = fi14
        _ = v15
    }
    v1 += (v1 << ((v1) & 63))
    fmt.Println(v2)
    for fi21 := 0; fi21 < 2; fi21++ {
        if (!true) {
            for fi22 := 0; fi22 < 18; fi22++ {
                v23 := GenS2{F0: 617, F1: -1}
                _ = fi22
                _ = v23
            }
            var v24 uint16 = v3
            _ = v24
        }
        if true {
            fmt.Println(v1)
            fmt.Println(v2, v3, v4)
        } else {
            if false {
                if (((((v11 - 1.5) + 1.5) + 0.1) * (float32((-2.25 - ((float64(v11) * 1.5) / 3.14159))) - float32(1e10))) > (0.0 * v11)) {
                    v2 -= (v2 * (v2 * (-(v2))))
                    var v25 uint16 = (v3 / ((2) | 1))
                    _ = v25
                } else {
                    v3 = v3
                    v2 += v2
                    var v26 int8 = (-(v2))
                    _ = v26
                }
                v3 = (v3 ^ (^(v3)))
            }
            for fi27 := 0; fi27 < 14; fi27++ {
                switch fi27 {
                case 42:
                    fmt.Println(fi21, v4, v2)
                case 4294967295:
                    var v28 int32 = 405
                    _ = v28
                case 127:
                    fmt.Println(fi27)
                }
                _ = fi27
            }
        }
        _ = len(v5)
        _ = fi21
    }
    fmt.Println(v2)
    if true {
        fmt.Println(v4, v2, v1)
    }
    v29 := GenS2{F0: -366, F1: -2147483648}
    switch (v4 + v4) {
    case 32767:
        var v30 bool = (!(255 >= 1))
        _ = v30
    case -671:
        fmt.Println(v11, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
