package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool, p1 uint, p2 int16) int64 {
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 bool = ((-(p2)) > 42)
            _ = fi6
            _ = v7
        }
        var v8 uint8 = 128
        _ = fi5
        _ = v8
    }
    var v9 int64 = 9223372036854775807
    fmt.Println(p1)
    var v10 string = (("a" + "foo bar") + "hello")
    var v11 float64 = (1e10 - 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return v9
}

func main() {
    var v1 bool = false
    var v2 uint32 = 2147483648
    var v3 bool = true
    v4 := []uint32{(v2 ^ (v2 * (v2 % ((v2) | 1)))), v2}
    for fi12 := 0; fi12 < 20; fi12++ {
        if len(v4) > 0 {
            v4[0] = v2
        }
        _ = fi12
    }
    v13 := GenS1{F0: 42, F1: 208}
    fmt.Println(v2)
    v1 = (!(!true))
    for fi14 := 0; fi14 < 4; fi14++ {
        if (!(1.5 >= 0.0)) {
            switch fi14 {
            case 0:
                v1 = v1
            case 1:
                var v15 int = (fi14 / ((fi14) | 1))
                _ = v15
            case 2147483647:
                var v16 bool = v1
                _ = v16
            default:
                v3 = (v3 && (v3 && v3))
            }
        } else {
            var v17 float32 = 1.5
            _ = v17
        }
        v18 := GenS2{F0: 127}
        _ = fi14
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 bool = v3
        _ = fi19
        _ = v20
    }
    var v21 float32 = 0.1
    var v22 uint16 = 2
    if true {
        _ = len(v4)
        var v23 int = 127
        if false {
            fmt.Println(v21)
            for fi24 := 0; fi24 < 10; fi24++ {
                v21 += float32(3.14159)
                var v25 int = v23
                _ = len(v4)
                _ = fi24
                _ = v25
            }
        } else {
            v23 = v23
            var v26 string = "hello"
            _ = v26
        }
        _ = v23
    }
    fmt.Println(v3)
    var v27 uint64 = 0
    v4 = append(v4, (v2 - (v2 + (v2 % (((^(v2))) | 1)))))
    v2 = (v2 ^ uint32(v22))
    v3 = (0.1 < (((0.0 + 0.1) / (v21 / (float32(((0.0 + 0.0) - 3.14159)) / -2.25))) - v21))
    v2 = (v2 ^ 499)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
}
