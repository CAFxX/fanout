package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int64 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint8; F1 int32; F2 string }
type GenS5 struct { F0 uint64; F1 uint }
func genFn0() int64 {
    v7 := GenS1{F0: 87, F1: 1, F2: -114}
    var v8 float32 = ((0.0 * float32((((1.5 / 3.14159) * (((1.5 - (-2.25 - (float64(1e10) * ((0.1 * -2.25) * (float64(float32((1.5 + 0.0))) / float64(1.5)))))) * 0.0) - 0.1)) * (0.0 + 0.0)))) - (0.1 + (-2.25 - 1e10)))
    if false {
        switch -853 {
        case 357:
            v9 := GenS1{F0: 1, F1: 886, F2: -1}
            _ = v9
        case 32768:
            for fi10 := 0; fi10 < 10; fi10++ {
                v11 := []int{}
                fmt.Println(v8, v8)
                v12 := GenS2{F0: 805, F1: "hello"}
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    } else {
        switch "foo bar" {
        case "hello":
            for fi13 := 0; fi13 < 4; fi13++ {
                if (!((1 >= 42) && ("\n\t" == ("foo bar" + ("foo bar" + ((("\n\t" + "foo bar") + "foo bar") + "")))))) {
                    var v14 int = 1
                    _ = v14
                }
                v15 := []int{127, (fi13 - fi13)}
                _ = fi13
                _ = v15
            }
        case "\n\t":
            v8 = v8
        case "":
            for fi16 := 0; fi16 < 6; fi16++ {
                v8 += (v8 / (v8 + (0.1 / v8)))
                v8 += 0.1
                _ = fi16
            }
        }
        v8 -= v8
    }
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 0
    v3 := []uint32{(v2 << ((173) & 63))}
    var v4 int32 = 42
    var v5 int = 32768
    v6 := make(map[uint32]uint32)
    if false {
        v2 = (v2 / (((^(v2))) | 1))
        fmt.Println(v2, v5, v2)
        fmt.Println(v4, v5)
    }
    v17 := GenS1{F0: -128, F1: v4, F2: 0}
    v4 += (v4 >> ((165) & 63))
    v18 := GenS3{F0: 99, F1: -869}
    v4 = v4
    if (3.14159 > (-2.25 * (((1e10 * (1e10 + (1.5 * -2.25))) / (0.0 + ((3.14159 / (1e10 / 0.1)) * (((float64((float32(-2.25) - 3.14159)) * -2.25) * (1.5 + 1e10)) / ((((float64(0.0) * 3.14159) - 1.5) * -2.25) * -2.25))))) - ((1.5 / ((1.5 - 0.0) * (3.14159 + 1e10))) / -2.25)))) {
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS4{F0: 2, F1: v4, F2: "foo bar"}
            _ = fi19
            _ = v20
        }
        delete(v6, v2)
        v21 := GenS5{F0: 2, F1: 300}
        _ = v21
    }
    _ = len(v3)
    _ = len(v1)
    for fi22 := 0; fi22 < 8; fi22++ {
        v2 = v2
        var v23 uint16 = 65535
        if len(v3) > 0 {
            v3[0] = v2
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
