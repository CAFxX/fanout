package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int32, p1 uint16, p2 int) string {
    p2 = (p2 << ((989) & 63))
    if true {
        if (false && false) {
            var v6 bool = (false || (100 == 42))
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := GenS1{F0: 1}
                for fi9 := 0; fi9 < 13; fi9++ {
                    v10 := GenS1{F0: 42}
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 4; fi11++ {
                    v12 := GenS1{F0: 542}
                    var v13 int32 = -859
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi7
                _ = v8
            }
            if (18446744073709551615 > 2) {
                var v14 int64 = -9223372036854775808
                _ = v14
            } else {
                fmt.Println(v6, p0)
                var v15 uint8 = 255
                _ = v15
            }
            _ = v6
        } else {
            for fi16 := 0; fi16 < 4; fi16++ {
                fmt.Println(p0, p1, p1)
                _ = fi16
            }
            p0 = -1
            v17 := GenS1{F0: 42}
            _ = v17
        }
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    var v1 float64 = -2.25
    var v2 int32 = 1
    var v3 float32 = 0.1
    var v4 int32 = (v2 / ((v2) | 1))
    var v5 int = -389
    v18 := []int{}
    if true {
        v18 = append(v18, (v5 ^ v5))
        fmt.Println(v1, v2, v4)
    } else {
        v4 -= v2
        v3 += v3
        switch "a" {
        case "foo bar":
            fmt.Println(v5, v5)
        }
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        v3 += ((float32(v1) + (3.14159 + (v3 * 1e10))) * 0.1)
        if true {
            var v20 int32 = v4
            _ = v20
        } else {
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 uint64 = 378
            v1 = 0.0
            v18 = append(v18, (v5 & (fi21 & fi21)))
            _ = fi21
            _ = v22
        }
        _ = fi19
    }
    var v23 uint16 = 1
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 int = (-(fi24))
        v26 := GenS1{F0: 255}
        _ = fi24
        _ = v25
        _ = v26
    }
    if false {
        var v27 float64 = 0.0
        if (1 <= 42) {
            v27 -= (v27 - v27)
            var v28 string = "hello"
            _ = v28
        }
        _ = v27
    }
    fmt.Println(v23, v23, v2)
    var v29 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
