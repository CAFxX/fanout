package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 int8, p1 uint64, p2 uint32) int64 {
    v6 := []int{32767, 127}
    v6 = append(v6, 42)
    fmt.Println(p0, p2)
    var v7 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 uint) string {
    if (32767 <= 32767) {
        p1 += p0
        v8 := GenS1{F0: 1, F1: "foo bar"}
        var v9 uint8 = 42
        _ = v8
        _ = v9
    }
    p1 += p0
    var v10 uint8 = 0
    for fi11 := 0; fi11 < 11; fi11++ {
        if (!(42 != -313)) {
            var v12 float64 = ((0.0 + (-2.25 * float64(0.0))) - -2.25)
            var v13 float32 = 1.5
            fmt.Println(v10, v12)
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return ("a" + "foo bar")
}

func main() {
    var v1 int8 = -128
    var v2 float32 = 1e10
    v3 := make(map[string]int)
    var v4 uint = 255
    var v5 int = 9223372036854775807
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 float32 = (v2 * ((v2 / v2) / 1.0))
        for fi16 := 0; fi16 < 12; fi16++ {
            if ("hello" != ("\n\t" + "\n\t")) {
                v5 += (fi16 - fi16)
                v2 -= v15
                if (false || ((v5 << ((v4) & 63)) <= 1)) {
                    v17 := v3["hello"]
                    v18 := []int{}
                    var v19 int32 = -1
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    v5 += fi14
                    v5 = -1
                    v15 = (1.5 / v2)
                    v20 := []int{-664}
                    _ = v20
                }
                v21 := GenS1{F0: fi16, F1: ""}
                v22 := GenS1{F0: (fi16 & (v5 ^ (fi16 + v5))), F1: ("\n\t" + "\n\t")}
                _ = v21
                _ = v22
            } else {
                v1 += v1
                v23 := []int{v5, v5, fi14, fi14}
                v1 = (v1 >> ((2) & 63))
                _ = v23
            }
            v24 := GenS2{F0: -104}
            v4 -= (v4 * (v4 + (v4 >> ((2) & 63))))
            _ = fi16
            _ = v24
        }
        _ = fi14
        _ = v15
    }
    v1 = (v1 | v1)
    var v25 float32 = v2
    v26 := GenS3{F0: 127, F1: 888}
    v4 += 65535
    var v27 uint32 = 2147483647
    var v28 int8 = -34
    v27 = (^(v27))
    fmt.Println(v27)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
