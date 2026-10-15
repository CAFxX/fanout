package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }

func main() {
    var v1 int8 = 42
    v2 := []string{"hello"}
    var v3 float32 = ((3.14159 - ((float32(1e10) + 1.5) - 0.1)) + (((3.14159 / 1e10) * 3.14159) + 1.5))
    var v4 int64 = 1
    v5 := []uint32{1, 4294967295}
    if (true && true) {
        for fi6 := 0; fi6 < 5; fi6++ {
            for fi7 := 0; fi7 < 18; fi7++ {
                if (false && false) {
                    v8 := GenS1{F0: -2147483648, F1: 755, F2: 2}
                    v4 += v4
                    v9 := GenS1{F0: 689, F1: 42, F2: 89}
                    _ = v8
                    _ = v9
                }
                _ = fi7
            }
            _ = len(v2)
            v10 := GenS1{F0: -2147483648, F1: 255, F2: 2}
            _ = fi6
            _ = v10
        }
        v3 = -2.25
    } else {
        var v11 int64 = v4
        var v12 int64 = 9223372036854775807
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: -2147483648, F1: 4294967295, F2: 128}
    var v14 float64 = 3.14159
    v1 -= 21
    if (!false) {
        _ = len(v2)
        if (2147483647 <= 32767) {
            var v15 bool = (true || (v4 > (-(v4))))
            _ = v15
        } else {
            for fi16 := 0; fi16 < 15; fi16++ {
                v1 = v1
                if len(v5) > 0 {
                    v5[0] = 445
                }
                if (float32((-2.25 / (v14 + v14))) >= v3) {
                    v3 = float32(-2.25)
                }
                _ = fi16
            }
            v17 := GenS1{F0: 0, F1: 255, F2: 128}
            fmt.Println(v3, v14, v4)
            _ = v17
        }
        v1 -= (v1 / ((v1) | 1))
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        v3 -= 0.0
        fmt.Println(v14)
        _ = fi18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        _ = len(v2)
        _ = fi19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 float32 = 0.0
        var v22 string = "hello"
        _ = fi20
        _ = v21
        _ = v22
    }
    switch "hello" {
    case "a":
        for fi23 := 0; fi23 < 16; fi23++ {
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 float32 = v3
                var v26 int = (fi23 >> ((350) & 63))
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi23
        }
    case "":
        v14 = v14
    }
    _ = len(v5)
    var v27 int32 = 1
    for fi28 := 0; fi28 < 15; fi28++ {
        switch "foo bar" {
        case "foo bar":
            var v29 uint16 = 766
            _ = v29
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
