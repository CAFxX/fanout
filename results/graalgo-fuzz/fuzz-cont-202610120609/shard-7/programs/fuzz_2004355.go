package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 string; F1 uint16; F2 uint16 }
func genFn0(p0 int8, p1 bool) int8 {
    v4 := []int{}
    for fi5 := 0; fi5 < 3; fi5++ {
        fmt.Println(p1, p1)
        if false {
            if (!true) {
                p0 += p0
                var v6 int8 = p0
                var v7 float64 = (float64(-2.25) + -2.25)
                _ = v6
                _ = v7
            } else {
                fmt.Println(p1, p1, fi5)
                var v8 uint8 = 42
                if p1 {
                    var v9 int32 = 2147483647
                    var v10 uint8 = (v8 + 2)
                    if (!(p0 == (p0 / ((p0) | 1)))) {
                        var v11 int = fi5
                        var v12 uint64 = 269
                        _ = v11
                        _ = v12
                    }
                    _ = v10
                    _ = v9
                }
                _ = v8
            }
            for fi13 := 0; fi13 < 8; fi13++ {
                p1 = (1 == 42)
                fmt.Println(fi13, fi13)
                var v14 int64 = 1
                var v15 uint = 42
                var v16 uint8 = 42
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        }
        var v17 int32 = -1
        _ = fi5
        _ = v17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 uint = 1
        v4 = append(v4, (fi18 * (fi18 & (fi18 | (fi18 | fi18)))))
        for fi20 := 0; fi20 < 9; fi20++ {
            switch "\n\t" {
            case "a":
                v21 := GenS1{F0: 0, F1: -358}
                _ = v21
            }
            var v22 float64 = (((((0.0 * 1.5) - ((1.5 + (float64(-2.25) / 0.1)) / 0.1)) / (0.1 * 1.5)) * (((0.1 + -2.25) / (float64(0.0) + 1e10)) * ((1.5 - -2.25) - 1e10))) / -2.25)
            _ = fi20
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v4
    return p0
}
func genFn1() int8 {
    var v23 int8 = 0
    v24 := []int{797}
    _ = v23
    _ = v24
    return v23
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 4294967295
    var v3 float64 = float64(0.1)
    switch "foo bar" {
    case "a":
        var v25 uint8 = 0
        _ = v25
    case "":
        var v26 float32 = ((3.14159 * 1e10) - 3.14159)
        _ = v26
    }
    var v27 uint8 = 136
    if (v3 >= ((((1e10 / float64(-2.25)) / 1e10) - v3) - float64(0.1))) {
        if false {
            if (true || (0.0 < (1e10 * float32((0.1 / (v3 - (float64(-2.25) + ((0.0 * v3) / v3)))))))) {
                if false {
                    for fi28 := 0; fi28 < 13; fi28++ {
                        var v29 float32 = 3.14159
                        _ = fi28
                        _ = v29
                    }
                } else {
                    var v30 uint32 = 2147483648
                    if (true || true) {
                        v31 := GenS1{F0: 18446744073709551615, F1: 42}
                        _ = v31
                    } else {
                        v3 = ((1.5 + (1.5 + (v3 / ((((v3 * (((1e10 - (0.0 + v3)) / 1.0) - (v3 - (v3 - 0.0)))) / 1.5) + (v3 / (v3 - float64(((((float32((float64(-2.25) * 3.14159)) * -2.25) + 3.14159) / -2.25) / 3.14159))))) + (v3 - (v3 - ((-2.25 / (((v3 * v3) - float64(1e10)) * (v3 + 1e10))) / ((v3 - -2.25) / (v3 - 1.5))))))))) * v3)
                    }
                    var v32 float64 = v3
                    _ = v30
                    _ = v32
                }
                v2 -= (^(v2))
                if (((v3 - v3) > v3) || true) {
                    fmt.Println(v3)
                }
            }
        }
        v33 := GenS1{F0: 2, F1: -243}
        _ = v33
    }
    v3 = 0.1
    v2 -= v2
    v2 -= 1
    v34 := GenS2{F0: "\n\t", F1: 42, F2: 42}
    switch -268 {
    case 0:
        delete(v1, "\n\t")
    default:
        fmt.Println(v27, v3)
    }
    var v35 int16 = 562
    v3 = v3
    _ = len(v1)
    var v36 uint16 = 2
    _ = len(v1)
    v36 += 0
    v2 += 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
