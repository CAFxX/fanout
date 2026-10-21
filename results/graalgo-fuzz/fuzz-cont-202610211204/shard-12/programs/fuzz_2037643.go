package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint64 }
func genFn0(p0 int32) int32 {
    var v4 uint = 42
    if (32767 > -574) {
        var v5 int8 = 100
        if false {
            p0 = -600
            var v6 uint = (v4 - (v4 + v4))
            _ = v6
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            if (!true) {
                for fi8 := 0; fi8 < 16; fi8++ {
                    if ((^(v4)) <= 271) {
                        v9 := GenS1{F0: -893, F1: 1, F2: (127 > (v5 << ((255) & 63)))}
                        _ = v9
                    }
                    p0 += (p0 - (p0 + 2147483647))
                    v10 := []int{(fi7 - 505), 232, fi8}
                    _ = fi8
                    _ = v10
                }
                v11 := GenS2{F0: fi7}
                switch fi7 {
                case 0:
                    v5 = v5
                case 42:
                    var v12 int = int(p0)
                    _ = v12
                case 128:
                    v13 := []int{(^(fi7)), (fi7 >> ((v4) & 63)), 128, fi7}
                    _ = v13
                default:
                    p0 += 759
                }
                _ = v11
            } else {
                v5 = v5
            }
            switch fi7 {
            case 32767:
                v4 -= v4
            default:
                var v14 uint = 2
                _ = v14
            }
            _ = fi7
        }
        _ = v5
    } else {
        var v15 int32 = p0
        _ = v15
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v4)
    default:
        var v16 uint64 = 42
        _ = v16
    }
    v17 := []int{4294967295}
    _ = p0
    _ = v17
    _ = v4
    return p0
}

func main() {
    v1 := []string{("a" + ((((("a" + ("\n\t" + "\n\t")) + "hello") + "foo bar") + ("hello" + "")) + "a")), "hello"}
    v2 := []int{128, 0}
    var v3 float64 = 0.1
    v18 := GenS1{F0: -1, F1: 1, F2: false}
    v3 += ((v3 - (v3 + (v3 - float64(float32((v3 / ((v3 - (v3 * 0.0)) * (v3 / (v3 - float64(-2.25)))))))))) + 1.5)
    var v19 int64 = -895
    v20 := GenS1{F0: (v19 % ((0) | 1)), F1: v19, F2: (1.5 != 0.0)}
    v21 := GenS3{F0: 32767, F1: 572}
    var v22 float64 = ((v3 + (v3 - v3)) * float64(1e10))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v23 int8 = -40
    if (!(128 <= 42)) {
        var v24 uint16 = 711
        v24 = v24
        _ = v24
    }
    var v25 int16 = -32768
    v26 := GenS2{F0: -171}
    var v27 int32 = 687
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
