package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32, p1 int32) uint {
    if true {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 int16 = 0
            _ = fi7
            _ = v8
        }
        if (!(true || (false || false))) {
            var v9 uint = 1
            v10 := []int{-9223372036854775808}
            _ = v10
            _ = v9
        } else {
            fmt.Println(p1, p1)
            v11 := []int{-432, 897, 42}
            var v12 uint8 = 255
            _ = v11
            _ = v12
        }
        var v13 uint64 = 884
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 float32 = float32((-2.25 - float64((float32(0.0) - float32(1e10)))))
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    return 563
}
func genFn1(p0 int) string {
    var v16 uint32 = 0
    if (!true) {
        v17 := []int{1, p0, p0, p0}
        _ = v17
    }
    v18 := []int{(-(p0)), 32767, p0, (p0 % ((p0) | 1))}
    var v19 int32 = 0
    _ = p0
    _ = v16
    _ = v18
    _ = v19
    return ""
}

func main() {
    var v1 int16 = 0
    var v2 uint = 4294967295
    var v3 bool = (0 != 127)
    v4 := make(map[int]bool)
    var v5 float64 = float64(3.14159)
    var v6 int = 42
    if true {
        v3 = v3
        if (v2 == (^(v2))) {
            v20 := GenS1{F0: 255, F1: false}
            _ = v20
        } else {
            var v21 int32 = -1
            for fi22 := 0; fi22 < 15; fi22++ {
                v23 := []int{fi22, (-(fi22)), int(v21), (-(fi22))}
                _ = fi22
                _ = v23
            }
            v24 := GenS2{F0: v2}
            _ = v21
            _ = v24
        }
        fmt.Println(v1)
    }
    v1 += v1
    v6 = v6
    var v25 uint = 171
    var v26 int64 = 990
    var v27 int64 = (v26 >> ((v25) & 63))
    for fi28 := 0; fi28 < 11; fi28++ {
        fmt.Println(v5)
        _ = fi28
    }
    var v29 uint64 = 42
    v30 := []int{v6, (v6 & int(v26)), v6}
    switch v6 {
    case 127:
        if false {
            for fi31 := 0; fi31 < 8; fi31++ {
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
                _ = fi31
            }
            v30 = append(v30, (v6 & (v6 << (((^(v25))) & 63))))
        } else {
            if (!true) {
                if true {
                    if (((-2.25 * 0.0) - 1.5) > 1.5) {
                        v32 := v4[(v6 * (v6 ^ v6))]
                        if len(v30) >= 2 {
                            v30 = v30[1:len(v30)]
                        }
                        v4[v6] = (v32 && (4294967295 == 47))
                        _ = v32
                    }
                }
            }
            if len(v30) > 0 {
                v30[0] = v6
            }
            v6 = (v6 << (((v25 << (((v25 / (((v2 * 42)) | 1))) & 63))) & 63))
        }
    case -550:
        switch "\n\t" {
        case "\n\t":
            if v3 {
                var v33 string = ("a" + (("\n\t" + ("\n\t" + ("hello" + "\n\t"))) + "a"))
                fmt.Println(v25, v6)
                var v34 uint32 = 4294967295
                _ = v33
                _ = v34
            } else {
                var v35 bool = v3
                fmt.Println(v35, v26)
                _ = v35
            }
        case "a":
            if len(v30) > 0 {
                v30[0] = v6
            }
        default:
            v4[0] = (1.5 <= (0.1 + 1.5))
        }
    default:
        fmt.Println(v25, v5, v5)
    }
    var v36 float32 = -2.25
    var v37 uint64 = 307
    _ = len(v30)
    var v38 int16 = v1
    if ("\n\t" <= "\n\t") {
        v27 -= 97
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
