package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64; F1 bool }
func genFn0(p0 uint16, p1 int64) int16 {
    switch 127 {
    case 32768:
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 12; fi8++ {
                switch (-(fi8)) {
                case -2147483648:
                    var v9 uint = 0
                    _ = v9
                }
                p0 = 42
                _ = fi8
            }
            _ = fi7
        }
    default:
        if true {
            v10 := []int{215, -715, -479, 1}
            v11 := GenS1{F0: "foo bar", F1: ""}
            _ = v10
            _ = v11
        } else {
            v12 := GenS1{F0: "foo bar", F1: "\n\t"}
            fmt.Println(p1)
            v13 := []int{397, 714, 9223372036854775807, 2147483647}
            _ = v12
            _ = v13
        }
    }
    fmt.Println(p1, p1)
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS2{F0: 2147483647, F1: -2147483648}
        p0 -= p0
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 65535, F1: 264}
    _ = p0
    _ = p1
    _ = v16
    return -122
}

func main() {
    v1 := []uint32{435, 4294967295, 44}
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 float64 = (0.1 * float64((((float32((((1e10 / -2.25) / -2.25) + 3.14159)) - ((3.14159 * ((1e10 / 1.0) - 1.5)) + 3.14159)) - 0.1) - -2.25)))
    var v5 int8 = -1
    var v6 uint16 = 32767
    switch 2147483647 {
    case 127:
        v5 += 1
    case 0:
        v17 := GenS3{F0: 932, F1: true}
        _ = v17
    }
    if ((0 < 0) || true) {
        v6 -= v6
        for fi18 := 0; fi18 < 2; fi18++ {
            v4 += v4
            switch (fi18 | (fi18 / ((127) | 1))) {
            case -9223372036854775808:
                var v19 int = (fi18 | fi18)
                _ = v19
            default:
                if true {
                    for fi20 := 0; fi20 < 14; fi20++ {
                        if (true && ((199 >= 128) || (1e10 >= ((0.0 * v4) * 0.0)))) {
                            v4 = (0.1 + -2.25)
                            v21 := GenS4{F0: 4294967295}
                            v6 += (v6 << ((18446744073709551615) & 63))
                            _ = v21
                        }
                        _ = fi20
                    }
                    v22 := GenS2{F0: 42, F1: 2147483647}
                    _ = v22
                }
            }
            _ = fi18
        }
        switch -2147483648 {
        case 705:
            v4 = ((v4 * 3.14159) - 1.5)
        case -375:
            v23 := v2[127]
            _ = v23
        case 1:
            var v24 uint8 = 139
            _ = v24
        default:
            v2[32767] = -1
        }
    }
    switch -1 {
    case -256:
        for fi25 := 0; fi25 < 13; fi25++ {
            v6 -= (v6 - 1)
            _ = fi25
        }
    case -9223372036854775808:
        v6 += (v6 - v6)
    case -1:
        var v26 float32 = 0.1
        _ = v26
    default:
        v27 := GenS5{F0: 65535, F1: 1000, F2: 0}
        _ = v27
    }
    if false {
        fmt.Println(v5, v6, v6)
        if true {
            v28 := GenS5{F0: 418, F1: -339, F2: 1}
            v3[""] = false
            delete(v3, "foo bar")
            _ = v28
        }
        v29 := GenS6{F0: "\n\t", F1: uint64(v6)}
        _ = v29
    }
    v30 := GenS7{F0: 2147483647}
    v6 -= 1
    if len(v1) > 0 {
        v1[0] = 65535
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        v1 = append(v1, 204)
        v32 := GenS8{F0: -1, F1: true}
        _ = len(v1)
        _ = fi31
        _ = v32
    }
    v6 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
