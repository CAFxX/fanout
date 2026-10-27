package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int16; F2 int32 }

func main() {
    v1 := []string{"", "\n\t", ("foo bar" + "")}
    v2 := make(map[string]string)
    var v3 bool = false
    v4 := make(map[int]int)
    var v5 bool = true
    var v6 string = "\n\t"
    var v7 bool = v3
    if (9223372036854775807 <= 42) {
        switch ("\n\t" + (v6 + v6)) {
        case "":
            v7 = (2147483647 >= -705)
        }
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint32 = 554
            var v11 uint64 = 871
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    var v12 int64 = -1
    var v13 float64 = (((((((3.14159 / (float64(1e10) / 1e10)) + float64(float32(1.5))) / (1e10 / (1.5 * 1e10))) + 1.5) * 3.14159) + ((-2.25 - 0.0) + 1e10)) * 3.14159)
    if v5 {
        var v14 uint64 = 18446744073709551615
        if v5 {
            var v15 float32 = (0.0 - -2.25)
            if (v5 || true) {
                v7 = (v5 || (0 < 18))
            } else {
                var v16 uint64 = v14
                _ = v16
            }
            v17 := GenS1{F0: 42}
            _ = v15
            _ = v17
        }
        _ = v14
    }
    v18 := GenS1{F0: 0}
    v19 := GenS2{F0: "hello", F1: 619, F2: 42}
    for fi20 := 0; fi20 < 3; fi20++ {
        v2["foo bar"] = "a"
        v13 = -2.25
        v6 = (("" + ("foo bar" + v6)) + (v6 + (v6 + v6)))
        _ = fi20
    }
    fmt.Println(v5, v3, v6)
    v21 := GenS2{F0: (v6 + "\n\t"), F1: 1000, F2: -2147483648}
    fmt.Println(v13, v5)
    if (-1 < -973) {
        v4[127] = -94
        v7 = (2 < 172)
        v6 = ("hello" + ("foo bar" + ((v6 + (v6 + (v6 + (v6 + v6)))) + "a")))
    }
    if (42 < 0) {
        var v22 uint8 = 1
        if v7 {
            v5 = (v13 >= (((float64(1e10) / (((v13 - ((3.14159 * v13) + ((v13 / v13) - 1.5))) + v13) + 3.14159)) * -2.25) + v13))
        } else {
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 bool = false
                v24 = (18446744073709551615 > 0)
                switch fi23 {
                case 876:
                    v3 = (893 == 219)
                }
                _ = fi23
                _ = v24
            }
        }
        if (1 <= 788) {
            v25 := GenS1{F0: 0}
            if len(v1) > 0 {
                v1[0] = v6
            }
            _ = v25
        } else {
            for fi26 := 0; fi26 < 16; fi26++ {
                if v5 {
                    if (1000 != -270) {
                        if false {
                            var v27 int16 = 1000
                            v28 := GenS2{F0: "", F1: (v27 - (v27 % (((v27 * (v27 | (v27 * -433)))) | 1))), F2: 0}
                            _ = v27
                            _ = v28
                        }
                    }
                } else {
                    var v29 uint8 = v22
                    _ = v29
                }
                var v30 uint16 = 739
                _ = fi26
                _ = v30
            }
            for fi31 := 0; fi31 < 1; fi31++ {
                v13 -= 1e10
                v32 := GenS1{F0: 245}
                var v33 float64 = float64(1e10)
                _ = fi31
                _ = v32
                _ = v33
            }
            fmt.Println(v12)
        }
        _ = v22
    } else {
        switch -1 {
        case 127:
            for fi34 := 0; fi34 < 8; fi34++ {
                v7 = (128 > 2)
                v3 = ((fi34 == fi34) && false)
                _ = fi34
            }
        case 42:
            fmt.Println(v6, v6, v13)
        case 1:
            for fi35 := 0; fi35 < 2; fi35++ {
                var v36 int8 = -121
                var v37 uint = 780
                _ = fi35
                _ = v36
                _ = v37
            }
        default:
            fmt.Println(v13, v6)
        }
        fmt.Println(v5, v13)
    }
    v12 -= -185
    if v3 {
        var v38 int8 = 0
        var v39 bool = ((-2.25 < 0.0) && true)
        v1 = append(v1, "\n\t")
        _ = v38
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
