package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint }
func genFn0(p0 int16, p1 uint8, p2 int64) uint32 {
    fmt.Println(p2, p1, p2)
    switch "" {
    case "a":
        p0 += (p0 & (p0 - 0))
    default:
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 61, F1: 1}
            if false {
                p1 -= 1
                p0 -= (p0 ^ -1)
            }
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS2{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 917
}

func main() {
    v1 := []string{"\n\t", ("\n\t" + "a"), ("a" + "hello")}
    var v2 uint8 = 1
    var v3 int32 = -209
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    var v6 uint8 = v2
    v2 -= 42
    var v10 bool = true
    var v11 int64 = -1
    var v12 float64 = 3.14159
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            if v10 {
                v5[fi13] = 65535
                v12 = float64(float32(float64((0.1 * ((-2.25 / -2.25) + 0.0)))))
                var v15 int8 = -56
                _ = v15
            }
            v16 := v4[("foo bar" + "")]
            _ = fi14
            _ = v16
        }
        _ = fi13
    }
    var v17 uint8 = (v2 | v2)
    var v18 uint8 = (v6 + v2)
    _ = len(v1)
    if (true && v10) {
        v6 = v6
        fmt.Println(v12, v3, v12)
    } else {
        v12 = v12
        v17 = (v6 / ((v17) | 1))
    }
    v2 += 127
    v19 := GenS1{F0: v11, F1: 18446744073709551615}
    var v20 int16 = 32767
    for fi21 := 0; fi21 < 15; fi21++ {
        v18 = v18
        var v22 int16 = 32767
        _ = fi21
        _ = v22
    }
    if v10 {
        if v10 {
            if len(v1) > 0 {
                v1[0] = "a"
            }
            v18 = 255
            for fi23 := 0; fi23 < 1; fi23++ {
                v18 -= (^(v17))
                v12 -= v12
                v18 += 2
                _ = fi23
            }
        }
        switch 0 {
        case -1:
            var v24 int8 = -128
            _ = v24
        case 1:
            for fi25 := 0; fi25 < 16; fi25++ {
                fmt.Println(v10)
                v26 := GenS2{F0: 437}
                _ = fi25
                _ = v26
            }
        }
        var v27 uint = 2
        _ = v27
    }
    if (!v10) {
        switch 2147483647 {
        case -909:
            v28 := GenS2{F0: 647}
            _ = v28
        case -2147483648:
            v29 := GenS2{F0: 4294967295}
            _ = v29
        case 128:
            v30 := GenS1{F0: 42, F1: 42}
            _ = v30
        }
        if (v20 < v20) {
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            fmt.Println(v6, v3)
            for fi31 := 0; fi31 < 9; fi31++ {
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v2, fi31, v18)
                    case "a":
                        v11 = (v11 >> ((255) & 63))
                    }
                    v3 = v3
                }
                fmt.Println(v6)
                _ = fi31
            }
        } else {
            var v32 float64 = v12
            _ = v32
        }
        var v33 int32 = v3
        _ = v33
    }
    v34 := GenS2{F0: 668}
    v35 := GenS2{F0: 18446744073709551615}
    v36 := GenS2{F0: 926}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
