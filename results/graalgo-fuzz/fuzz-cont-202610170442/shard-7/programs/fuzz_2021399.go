package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint, p1 uint) uint8 {
    var v4 uint16 = 65535
    v5 := GenS1{F0: 184, F1: 31, F2: 1}
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 66
}
func genFn1(p0 uint8, p1 bool) uint64 {
    fmt.Println(p0, p0)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 float64 = -2.25
        var v9 bool = p1
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 394
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]bool)
    var v3 float64 = -2.25
    v3 = 0.1
    v1 += 312
    var v10 uint8 = genFn0(1, 65535)
    v3 += ((((float64((0.1 / 3.14159)) + 0.1) / (((float64(1.5) + v3) + (((0.1 * v3) * (float64(0.1) / (1e10 - 1.5))) - 0.1)) / (float64(((3.14159 + 0.1) * 1e10)) - ((float64(0.0) + (1e10 - (1e10 / (((v3 * float64((1.5 / 1.0))) / v3) * (v3 - ((v3 + 1e10) - 0.1)))))) * (1.5 / (v3 + (float64(0.1) - v3))))))) + v3) - (float64(1.5) / v3))
    v3 -= float64(3.14159)
    switch 4294967295 {
    case 42:
        if true {
            fmt.Println(v10, v10)
        } else {
            if (0 < 42) {
                v11 := GenS2{F0: 52}
                v1 += 56
                _ = v11
            } else {
                v10 += v10
            }
            if ((!(false || true)) && ((^(v10)) > v10)) {
                for fi12 := 0; fi12 < 12; fi12++ {
                    v13 := GenS3{F0: 1}
                    _ = fi12
                    _ = v13
                }
                v14 := GenS1{F0: (^(v10)), F1: v10, F2: genFn1((v10 | (v10 >> ((705) & 63))), true)}
                _ = v14
            }
            v15 := []int{}
            _ = v15
        }
    case 496:
        switch ((("a" + ("\n\t" + "foo bar")) + "foo bar") + "\n\t") {
        case "\n\t":
            v1 = (v1 >> ((42) & 63))
        default:
            var v16 uint32 = 384
            _ = v16
        }
    default:
        var v17 string = ""
        _ = v17
    }
    fmt.Println(v1, v3)
    v18 := GenS3{F0: 1}
    fmt.Println(v1, v10, v10)
    var v19 int = -2147483648
    var v20 bool = false
    for fi21 := 0; fi21 < 7; fi21++ {
        if (100 != 53) {
            v22 := GenS1{F0: 253, F1: 0, F2: v1}
            _ = v22
        }
        delete(v2, ("a" + ""))
        var v23 int8 = 0
        _ = fi21
        _ = v23
    }
    var v24 int32 = 1
    v25 := GenS2{F0: genFn0(4294967295, 255)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
}
