package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 string }
func genFn0() uint8 {
    var v6 float32 = 0.0
    if true {
        v6 += v6
        v6 += -2.25
        if ((3.14159 / (v6 / 3.14159)) < 0.1) {
            v7 := []int{-9223372036854775808}
            switch -771 {
            case 1:
                var v8 uint64 = 0
                _ = v8
            case -9223372036854775808:
                fmt.Println(v6)
            }
            _ = v7
        }
    }
    v6 = (3.14159 * (-2.25 * 3.14159))
    var v9 float64 = ((float64((v6 / ((v6 / 0.1) - v6))) * (1.5 * -2.25)) * float64(float32((((float64(1e10) / float64((float32((3.14159 * float64((-2.25 + 3.14159)))) * v6))) + float64(v6)) * float64((1.5 * ((v6 + (v6 / v6)) * (v6 + 3.14159))))))))
    var v10 uint64 = 688
    _ = v10
    _ = v6
    _ = v9
    return 128
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 1
    v3 := make(map[int]string)
    var v4 uint8 = 255
    v5 := make(map[string]int)
    v11 := []int{127}
    v12 := GenS1{F0: 42, F1: 185}
    fmt.Println(v4, v2)
    v2 = v2
    v4 += v4
    for fi13 := 0; fi13 < 17; fi13++ {
        v4 -= 42
        v14 := v5["\n\t"]
        _ = fi13
        _ = v14
    }
    var v15 bool = true
    v2 = -128
    var v16 uint16 = 2
    v11 = append(v11, -9223372036854775808)
    var v17 float32 = 1e10
    fmt.Println(v17)
    if (0 >= 0) {
        for fi18 := 0; fi18 < 14; fi18++ {
            for fi19 := 0; fi19 < 15; fi19++ {
                v4 -= (v4 * 42)
                _ = fi19
            }
            v15 = (149 >= 469)
            v16 -= (v16 / ((v16) | 1))
            _ = fi18
        }
        v20 := GenS1{F0: 65535, F1: -130}
        _ = v20
    } else {
        v21 := GenS2{F0: (("hello" + "hello") + "")}
        v22 := GenS2{F0: "\n\t"}
        switch "" {
        case "":
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        default:
            v23 := GenS1{F0: 4294967295, F1: -307}
            _ = v23
        }
        _ = v21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
