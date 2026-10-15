package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 bool }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint32) int {
    v6 := GenS1{F0: p0, F1: 203, F2: (!true)}
    var v7 int = 42
    p0 = 588
    _ = p0
    _ = v6
    _ = v7
    return (v7 * 6)
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    var v3 uint32 = 835
    v4 := make(map[int]uint32)
    var v5 string = "\n\t"
    v5 = v5
    if false {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int64 = 0
            v10 := GenS1{F0: v3, F1: 2, F2: (1 <= 0)}
            v9 += v9
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    if (-128 >= 101) {
        v3 = (v3 ^ v3)
    } else {
        v11 := v2[2147483647]
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := []int{}
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int8 = -1
                fmt.Println(fi12)
                var v16 uint32 = (^(v3))
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    v3 += (v3 | v3)
    var v17 int16 = 989
    var v18 uint64 = 42
    if ((!(v18 <= 2)) || false) {
        switch 4294967295 {
        case 0:
            delete(v1, 9223372036854775807)
        default:
            v19 := GenS1{F0: 42, F1: 0, F2: false}
            _ = v19
        }
    }
    v20 := []int{128, 2147483647, 127}
    v17 -= (v17 - v17)
    switch v5 {
    case "":
        var v21 int32 = -2147483648
        _ = v21
    case "\n\t":
        for fi22 := 0; fi22 < 16; fi22++ {
            v5 = ""
            if true {
                v18 = 774
                switch (v5 + v5) {
                case "foo bar":
                    v23 := GenS1{F0: 4294967295, F1: 2, F2: false}
                    _ = v23
                case "hello":
                    if len(v20) > 0 {
                        v20[0] = fi22
                    }
                case "":
                    v3 -= v3
                }
                switch v5 {
                case "a":
                    var v24 string = (v5 + "foo bar")
                    _ = v24
                case "":
                    var v25 uint16 = 604
                    _ = v25
                }
            }
            fmt.Println(fi22, v18, v3)
            _ = fi22
        }
    default:
        v26 := GenS1{F0: (v3 >> ((255) & 63)), F1: 121, F2: false}
        _ = v26
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        _ = len(v2)
        v28 := GenS1{F0: (v3 << ((42) & 63)), F1: 0, F2: ((3.14159 / 1.0) < 0.1)}
        v18 = 4294967295
        _ = fi27
        _ = v28
    }
    var v29 uint = 244
    for fi30 := 0; fi30 < 9; fi30++ {
        for fi31 := 0; fi31 < 2; fi31++ {
            if false {
                switch (v5 + (((v5 + (v5 + (("hello" + v5) + "a"))) + "hello") + "")) {
                case "foo bar":
                    v20 = append(v20, genFn0(v3))
                case "a":
                    fmt.Println(fi31, v18)
                default:
                    v18 += (v18 + v18)
                }
                _ = len(v1)
                v32 := GenS2{F0: false, F1: 255}
                _ = v32
            } else {
                var v33 int64 = 566
                v3 = (v3 << ((v29) & 63))
                _ = v33
            }
            if true {
                var v34 int64 = 1000
                fmt.Println(v5, v34, v17)
                _ = v34
            }
            _ = fi31
        }
        fmt.Println(v3)
        _ = fi30
    }
    fmt.Println(v5, v3)
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 int64 = 426
        if (44 >= 128) {
            v5 = "\n\t"
            fmt.Println(v29, fi35)
        } else {
            v37 := GenS3{F0: 42}
            _ = v37
        }
        switch 32767 {
        case 128:
            var v38 float32 = 1e10
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
