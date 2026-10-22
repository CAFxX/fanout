package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int64; F1 int; F2 int64 }
func genFn0(p0 int64, p1 uint16, p2 int16) bool {
    var v5 int8 = 0
    if (p1 <= (p1 & (p1 | (^(p1))))) {
        for fi6 := 0; fi6 < 3; fi6++ {
            p1 = (^(p1))
            v7 := GenS1{F0: 620, F1: 4294967295}
            _ = fi6
            _ = v7
        }
        var v8 uint8 = 170
        var v9 float32 = (float32((((float64(1.5) / 1e10) + ((-2.25 * 3.14159) - (0.0 + 1e10))) / 3.14159)) + (1.5 * float32(float64(3.14159))))
        _ = v8
        _ = v9
    } else {
        p0 += p0
        var v10 int16 = 882
        var v11 bool = (3.14159 >= 0.0)
        _ = v10
        _ = v11
    }
    p2 -= p2
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (2147483648 == 861)
}

func main() {
    var v1 uint16 = 1
    v2 := []int{-270, -9223372036854775808, -2147483648, 225}
    var v3 bool = true
    var v4 int16 = 32767
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 float64 = 1e10
        for fi14 := 0; fi14 < 2; fi14++ {
            v3 = true
            switch "hello" {
            case "hello":
                v15 := GenS2{F0: 9223372036854775807, F1: -2147483648, F2: -9223372036854775808}
                _ = v15
            default:
                switch (fi14 + 0) {
                case 32767:
                    v3 = v3
                case 9223372036854775807:
                    v16 := GenS1{F0: 2, F1: 0}
                    _ = v16
                }
            }
            _ = len(v2)
            _ = fi14
        }
        for fi17 := 0; fi17 < 15; fi17++ {
            switch fi12 {
            case 128:
                v1 = v1
            case 1:
                switch fi12 {
                case 128:
                    var v18 int16 = v4
                    _ = v18
                case -9223372036854775808:
                    v19 := GenS1{F0: 4294967295, F1: 0}
                    _ = v19
                case -1:
                    v3 = v3
                default:
                    var v20 float64 = 1e10
                    _ = v20
                }
            case -9223372036854775808:
                v3 = (!(("" < "\n\t") || (18446744073709551615 == 18446744073709551615)))
            }
            _ = fi17
        }
        _ = fi12
        _ = v13
    }
    v1 += v1
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint16 = (^(v1))
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 float64 = ((-2.25 + ((0.0 * 0.0) / 3.14159)) + (1e10 + (float64((-2.25 - (float32(1.5) * -2.25))) + (((3.14159 + -2.25) / 1.5) / (0.1 / ((float64(1e10) - float64(3.14159)) * float64((0.1 / 1.0))))))))
        v4 += 0
        _ = fi23
        _ = v24
    }
    fmt.Println(v4, v3)
    if (!(128 < -864)) {
        var v25 string = ("a" + "")
        _ = v25
    }
    var v26 float64 = 1.5
    v27 := GenS2{F0: 1, F1: 483, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
