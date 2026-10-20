package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 string, p1 uint, p2 uint8) int {
    var v5 uint32 = 2
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int64 = 756
        if false {
            if false {
                var v8 uint8 = (p2 + 186)
                v9 := GenS1{F0: p1}
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: (p1 - p1)}
            v7 = 300
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 52
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = 0.1
    v3 := make(map[int]bool)
    var v4 float32 = (v2 - v2)
    if (9223372036854775807 <= 1) {
        fmt.Println(v2, v1)
        if false {
            if ((201 == 0) && ((((1e10 - v4) * 0.0) - ((v2 - v4) / float32((3.14159 + 1e10)))) < (float32((-2.25 + ((((((0.1 * 3.14159) * 1e10) * float64(v4)) - 3.14159) * 3.14159) + (0.0 + 0.0)))) / (v2 * (0.1 - v2))))) {
                v1 = "hello"
                v1 = ((v1 + ("hello" + v1)) + (v1 + v1))
            }
            fmt.Println(v2)
        } else {
            var v11 float64 = -2.25
            _ = v11
        }
        var v12 string = v1
        _ = v12
    }
    var v13 int = 9223372036854775807
    v13 = v13
    var v14 int = 4294967295
    v15 := GenS1{F0: 2}
    var v16 uint32 = 2147483647
    v1 = v1
    switch (v1 + (v1 + v1)) {
    case "foo bar":
        delete(v3, (v13 & (v13 / (((^(v14))) | 1))))
    default:
        if (!true) {
            fmt.Println(v14)
            for fi17 := 0; fi17 < 11; fi17++ {
                v4 = 1.5
                if false {
                    fmt.Println(v4, v16)
                }
                _ = fi17
            }
            var v18 int64 = 569
            _ = v18
        } else {
            v19 := GenS2{F0: 1, F1: 42}
            _ = v19
        }
    }
    _ = len(v3)
    if ((80 != 9223372036854775807) || ((^(v16)) <= (v16 % (((v16 ^ v16)) | 1)))) {
        var v20 int16 = -32768
        if false {
            v4 = (v4 - v4)
            v21 := GenS1{F0: 4294967295}
            var v22 uint16 = 42
            _ = v21
            _ = v22
        }
        v4 -= v2
        _ = v20
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(v2, v14)
        v16 -= v16
        _ = fi23
    }
    v24 := GenS1{F0: 141}
    fmt.Println(v2)
    v2 += v2
    var v25 int32 = -1
    var v26 uint16 = 812
    var v27 uint16 = v26
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
