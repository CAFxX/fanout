package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
func genFn0(p0 uint) int8 {
    if ((128 != -174) || (-9223372036854775808 < -1)) {
        var v7 bool = false
        switch 109 {
        case 2147483647:
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := GenS1{F0: 42, F1: 4294967295, F2: (fi8 + fi8)}
                _ = fi8
                _ = v9
            }
        case 128:
            switch 128 {
            case 2147483647:
                v10 := []int{}
                _ = v10
            }
        case 0:
            p0 += 42
        default:
            p0 = p0
        }
        _ = v7
    }
    p0 = p0
    _ = p0
    return -128
}

func main() {
    v1 := []string{("hello" + "foo bar"), "a"}
    v2 := make(map[int]int)
    var v3 uint32 = 132
    var v4 int16 = 1
    var v5 float64 = 1e10
    var v6 float32 = 3.14159
    v2[0] = 42
    fmt.Println(v3)
    for fi11 := 0; fi11 < 12; fi11++ {
        if (1 >= 0) {
            v6 -= -2.25
            var v12 float32 = (v6 + ((v6 - ((((-2.25 + v6) + float32((((0.0 - 1e10) / v5) + v5))) - -2.25) / float32(float64(v6)))) - 0.1))
            v6 += 0.0
            _ = v12
        } else {
            _ = len(v2)
        }
        var v13 bool = (!("a" != "hello"))
        _ = fi11
        _ = v13
    }
    var v14 int16 = -225
    v15 := GenS1{F0: 2, F1: 4294967295, F2: 0}
    fmt.Println(v6, v3, v14)
    switch "hello" {
    case "hello":
        for fi16 := 0; fi16 < 4; fi16++ {
            if (!false) {
                v1 = append(v1, "hello")
            }
            _ = fi16
        }
    case "":
        var v17 int8 = genFn0(152)
        _ = v17
    case "a":
        v6 += ((v6 / (0.1 / 0.1)) + ((v6 - 1e10) * v6))
    }
    var v18 int = 89
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 uint = 65535
        fmt.Println(v14, v5, fi19)
        v5 = v5
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: 155, F1: 42, F2: v18}
    for fi22 := 0; fi22 < 2; fi22++ {
        v1 = append(v1, (("hello" + "\n\t") + "hello"))
        _ = fi22
    }
    var v23 uint32 = 65535
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
