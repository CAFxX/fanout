package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 string }
func genFn0(p0 string, p1 string, p2 uint64) int16 {
    var v4 uint32 = 4294967295
    v5 := []int{128, 2147483647, -378, -1}
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int8 = -1
            v9 := GenS1{F0: (p2 == p2), F1: fi6}
            switch 4294967295 {
            case -2147483648:
                switch fi7 {
                case 639:
                    p2 += (^(p2))
                case 32768:
                    v4 -= 0
                default:
                    switch p0 {
                    case "foo bar":
                        v10 := GenS1{F0: (true || (42 > 0)), F1: 351}
                        _ = v10
                    default:
                        var v11 int16 = -1
                        _ = v11
                    }
                }
            case 474:
                v5 = append(v5, 32767)
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(p2, p1)
        var v13 float32 = 1e10
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := []int{32767, 4294967295}
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    if (false && true) {
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: (!(((((0.0 / float32(0.1)) + (((1e10 - 1.5) / 0.1) * 3.14159)) / -2.25) * 1e10) < -2.25)), F1: fi14}
            v16 := GenS1{F0: (1e10 >= (float64(1.5) / 1e10)), F1: fi14}
            _ = fi14
            _ = v15
            _ = v16
        }
        v17 := GenS1{F0: (2147483647 != 936), F1: 128}
        if false {
            var v18 int = 42
            var v19 int32 = -1
            var v20 int = v18
            _ = v18
            _ = v19
            _ = v20
        } else {
            var v21 int8 = 71
            _ = v21
        }
        _ = v17
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 uint = 4294967295
        _ = fi22
        _ = v23
    }
    var v24 int = 2147483647
    v25 := v2[2]
    switch (v24 - (v24 >> ((2) & 63))) {
    case 1:
        if true {
            v25 = v25
            var v26 uint32 = 479
            fmt.Println(v24, v26)
            _ = v26
        }
    case 32768:
        v25 = v25
    default:
        v24 = (v24 & 32768)
    }
    fmt.Println(v24, v24)
    for fi27 := 0; fi27 < 4; fi27++ {
        v25 = v25
        _ = fi27
    }
    v28 := GenS2{F0: ((v25 + "") + ((v25 + "") + ""))}
    fmt.Println(v25, v24)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
}
