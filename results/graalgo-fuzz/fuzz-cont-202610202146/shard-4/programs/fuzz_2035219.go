package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64, p1 int) string {
    v4 := GenS1{F0: 65535, F1: ("foo bar" < "a")}
    p0 -= -911
    p0 -= (p0 % ((p0) | 1))
    var v5 int = p1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "foo bar"
}

func main() {
    var v1 bool = false
    var v2 int16 = 1
    var v3 int64 = -9223372036854775808
    v6 := GenS2{F0: 0}
    v7 := []int{0, 128, -1, 825}
    var v8 bool = (!(((-2.25 - (3.14159 / 1.0)) / (3.14159 / 0.1)) == float64((0.1 / float32(float64(((0.1 + 1.5) * float32(float64(1e10)))))))))
    for fi9 := 0; fi9 < 13; fi9++ {
        _ = len(v7)
        _ = fi9
    }
    var v10 uint16 = 2
    v11 := GenS2{F0: 489}
    if (248 != -1) {
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS2{F0: 4294967295}
            _ = fi12
            _ = v13
        }
        if false {
            _ = len(v7)
        } else {
            var v14 uint32 = 2147483647
            v15 := GenS1{F0: 2, F1: v1}
            _ = v14
            _ = v15
        }
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
    }
    if (0 > v2) {
        var v16 bool = ((780 >= 65535) && v1)
        v8 = (v16 && ("" != "a"))
        v16 = ((!false) && (!(((1 < -567) && v8) || (v1 || (float32(-2.25) > 0.0)))))
        _ = v16
    } else {
        var v17 float32 = 1.5
        _ = v17
    }
    switch -470 {
    case 9223372036854775807:
        _ = len(v7)
    default:
        v8 = v1
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        if v8 {
            for fi19 := 0; fi19 < 3; fi19++ {
                fmt.Println(fi18, fi19, fi19)
                v10 = (v10 * v10)
                _ = fi19
            }
            var v20 bool = (!v1)
            var v21 bool = v1
            _ = v20
            _ = v21
        } else {
            fmt.Println(v1, v2, v2)
        }
        _ = fi18
    }
    var v22 int8 = -1
    v7 = append(v7, 1)
    var v23 bool = (!false)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
