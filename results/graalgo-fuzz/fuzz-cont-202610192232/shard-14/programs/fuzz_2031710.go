package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 string; F1 int64; F2 int }
type GenS3 struct { F0 string }
func genFn0(p0 string) uint16 {
    if (false || (true || (882 >= 42))) {
        if false {
            var v6 uint32 = 131
            _ = v6
        } else {
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := GenS1{F0: 940, F1: 2}
                for fi9 := 0; fi9 < 3; fi9++ {
                    v10 := GenS1{F0: 65535, F1: 42}
                    var v11 string = ""
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = fi7
                _ = v8
            }
            var v12 uint32 = 128
            if (((1.5 * (((0.0 / 1.5) * (0.0 + -2.25)) / float64(1e10))) / (1.5 - 0.1)) != (-2.25 + 1.5)) {
                v12 -= (^(v12))
                var v13 uint = 635
                _ = v13
            } else {
                v14 := []int{}
                _ = v14
            }
            _ = v12
        }
        if true {
            if true {
                p0 = p0
                v15 := []int{2147483647}
                _ = v15
            }
            v16 := GenS2{F0: "", F1: 1, F2: -942}
            _ = v16
        }
    }
    var v17 int8 = 42
    _ = p0
    _ = v17
    return 255
}

func main() {
    var v1 float64 = (0.0 * (((3.14159 + float64((3.14159 + -2.25))) * 1.5) + -2.25))
    var v2 int16 = 0
    v3 := make(map[string]string)
    var v4 int8 = 78
    var v5 int32 = -2147483648
    switch 2147483647 {
    case 688:
        v5 += (v5 % ((v5) | 1))
    case 517:
        v4 -= (v4 / (((v4 * (-(v4)))) | 1))
    }
    var v18 uint64 = 0
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS3{F0: "a"}
        _ = fi19
        _ = v20
    }
    switch 767 {
    case 127:
        v21 := []int{32767, 9223372036854775807, 0, 42}
        _ = v21
    case -673:
        var v22 uint32 = 2
        _ = v22
    default:
        switch (("\n\t" + ("\n\t" + "a")) + "") {
        case "":
            v23 := GenS1{F0: 0, F1: v18}
            _ = v23
        }
    }
    fmt.Println(v2, v2, v4)
    var v24 int64 = 1
    var v25 int16 = v2
    var v26 uint = 572
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
