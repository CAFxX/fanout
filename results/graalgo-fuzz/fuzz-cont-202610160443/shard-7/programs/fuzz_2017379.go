package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0() string {
    v4 := GenS1{F0: 42}
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: 2147483647}
        var v7 uint8 = 187
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = v4
    return "foo bar"
}
func genFn1(p0 uint8) string {
    if (-1 < 1) {
        var v8 int64 = 1
        var v9 float32 = float32((0.1 + 1.5))
        _ = v8
        _ = v9
    }
    var v10 uint = uint(p0)
    v10 -= uint(p0)
    for fi11 := 0; fi11 < 10; fi11++ {
        switch fi11 {
        case 2147483647:
            var v12 int = (-(fi11))
            _ = v12
        }
        var v13 uint = 1
        var v14 float64 = (float64(float32(1e10)) - 3.14159)
        _ = fi11
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v10
    return ("" + "")
}

func main() {
    var v1 int64 = -741
    var v2 bool = ((0 != -741) && (!false))
    var v3 int = -9223372036854775808
    v15 := GenS2{F0: v1, F1: 255}
    switch v3 {
    case -78:
        if (v1 == v1) {
            switch (v3 >> ((545) & 63)) {
            case 1:
                if v2 {
                    v16 := []int{-597, 0}
                    fmt.Println(v3)
                    var v17 float32 = (((-2.25 * 1.5) / (float32((3.14159 / -2.25)) * ((1e10 * (1e10 + 0.0)) - float32((1.5 / 0.1))))) / (-2.25 / 0.1))
                    _ = v16
                    _ = v17
                }
            case 128:
                v2 = v2
            case 32768:
                var v18 float64 = ((0.1 + (float64(1e10) - -2.25)) * float64(((0.0 / 3.14159) * 0.0)))
                _ = v18
            }
            for fi19 := 0; fi19 < 8; fi19++ {
                v3 += (fi19 + (v3 | fi19))
                v20 := []int{}
                _ = fi19
                _ = v20
            }
        }
    case 403:
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint64 = 226
            v23 := []int{42, fi21, (fi21 << ((uint(v22)) & 63)), v3}
            _ = fi21
            _ = v22
            _ = v23
        }
    default:
        if false {
            v2 = true
        } else {
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 uint64 = 349
                v2 = v2
                for fi26 := 0; fi26 < 3; fi26++ {
                    fmt.Println(v1)
                    var v27 bool = (65535 == 25)
                    _ = fi26
                    _ = v27
                }
                _ = fi24
                _ = v25
            }
            v2 = (0.1 > 3.14159)
        }
    }
    if (!v2) {
        if (v2 && (v2 && false)) {
            var v28 int16 = -1
            _ = v28
        }
    }
    v29 := []int{v3, v3, 127, (v3 >> ((2) & 63))}
    switch (v3 >> ((2) & 63)) {
    case -2147483648:
        v30 := GenS2{F0: 0, F1: 139}
        _ = v30
    default:
        for fi31 := 0; fi31 < 17; fi31++ {
            v32 := GenS3{F0: 992, F1: 127}
            for fi33 := 0; fi33 < 8; fi33++ {
                v3 = fi33
                _ = fi33
            }
            if v2 {
                var v34 int = v3
                var v35 int16 = 301
                var v36 int = 357
                _ = v34
                _ = v35
                _ = v36
            }
            _ = fi31
            _ = v32
        }
    }
    switch 127 {
    case 9223372036854775807:
        for fi37 := 0; fi37 < 10; fi37++ {
            v3 = (^(v3))
            var v38 int64 = -9223372036854775808
            _ = fi37
            _ = v38
        }
    case 1:
        if v2 {
            var v39 float64 = float64(1e10)
            _ = v39
        } else {
            if ((v3 | 127) <= 9223372036854775807) {
                var v40 uint32 = 4294967295
                _ = v40
            }
            fmt.Println(v2, v3, v1)
        }
    case 0:
        v29 = append(v29, -9223372036854775808)
    }
    v41 := GenS2{F0: (v1 << ((638) & 63)), F1: 255}
    v1 = (v1 / ((v1) | 1))
    for fi42 := 0; fi42 < 20; fi42++ {
        v3 -= (fi42 ^ 95)
        fmt.Println(v2, v2)
        fmt.Println(v2, fi42, v3)
        _ = fi42
    }
    fmt.Println(v3)
    v2 = false
    for fi43 := 0; fi43 < 19; fi43++ {
        fmt.Println(fi43)
        _ = fi43
    }
    v44 := GenS2{F0: v1, F1: 127}
    var v45 int32 = -402
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v41
    _ = v44
    _ = v45
}
