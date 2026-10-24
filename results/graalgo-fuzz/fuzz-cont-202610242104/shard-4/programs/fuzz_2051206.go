package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint8; F2 bool }
func genFn0(p0 int8, p1 int64) string {
    v5 := GenS1{F0: p0}
    p1 = p1
    p0 = 1
    _ = p0
    _ = p1
    _ = v5
    return ""
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    var v3 uint8 = 121
    var v4 int16 = -511
    for fi6 := 0; fi6 < 7; fi6++ {
        v4 += v4
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: -1}
            v9 := GenS1{F0: 100}
            var v10 uint32 = 2147483647
            var v11 uint16 = 157
            if false {
                v2[fi6] = 839
            }
            _ = fi7
            _ = v10
            _ = v11
            _ = v8
            _ = v9
        }
        fmt.Println(fi6, fi6, v3)
        _ = fi6
    }
    var v12 float64 = (-2.25 * float64(1e10))
    v13 := v2[-212]
    if ((239 != v3) || true) {
        v12 -= 0.0
        v12 = float64((0.1 - (1e10 / 1.5)))
        v14 := GenS1{F0: -120}
        _ = v14
    } else {
        for fi15 := 0; fi15 < 17; fi15++ {
            if (("\n\t" + "") <= ("foo bar" + "hello")) {
                v2[4294967295] = v13
                v16 := GenS1{F0: -1}
                _ = v16
            }
            v17 := GenS1{F0: -1}
            _ = fi15
            _ = v17
        }
        if false {
            var v18 int64 = 0
            switch "\n\t" {
            case "\n\t":
                var v19 int64 = v18
                _ = v19
            default:
                var v20 float32 = (1e10 + float32((3.14159 - 1e10)))
                _ = v20
            }
            _ = v18
        }
    }
    v21 := []int{1}
    v4 += v4
    _ = len(v21)
    for fi22 := 0; fi22 < 9; fi22++ {
        v12 = v12
        fmt.Println(v13, v3, v12)
        _ = fi22
    }
    v13 -= v13
    v13 = 2147483647
    fmt.Println(v4)
    switch 32768 {
    case 71:
        for fi23 := 0; fi23 < 14; fi23++ {
            switch (("" + (("hello" + "a") + "hello")) + "hello") {
            case "":
                if ((!false) && (!true)) {
                    if false {
                        for fi24 := 0; fi24 < 13; fi24++ {
                            var v25 uint16 = 2
                            _ = fi24
                            _ = v25
                        }
                        v26 := GenS2{F0: 32767, F1: (^(v3)), F2: (255 <= 65535)}
                        _ = v26
                    } else {
                        var v27 int = fi23
                        var v28 int64 = 897
                        _ = v27
                        _ = v28
                    }
                    var v29 bool = false
                    v3 = 255
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 10; fi30++ {
                        var v31 bool = true
                        v13 += (^(v13))
                        _ = fi30
                        _ = v31
                    }
                    var v32 int64 = -411
                    _ = v32
                }
            default:
                switch ("" + "") {
                case "foo bar":
                    fmt.Println(v4, v13, v12)
                }
            }
            var v33 int16 = (v4 + 799)
            switch 203 {
            case 32768:
                for fi34 := 0; fi34 < 7; fi34++ {
                    var v35 int = 952
                    _ = fi34
                    _ = v35
                }
            case -9223372036854775808:
                var v36 int16 = (v4 ^ (v4 >> ((42) & 63)))
                _ = v36
            }
            _ = fi23
            _ = v33
        }
    case -2147483648:
        var v37 string = "hello"
        _ = v37
    case -866:
        v13 += v13
    default:
        var v38 uint16 = 1
        _ = v38
    }
    v13 -= v13
    v2[-1] = v13
    var v39 uint8 = (v3 & v3)
    switch 32768 {
    case 53:
        switch "a" {
        case "a":
            v12 -= (3.14159 / (float64(1e10) * 0.1))
        case "\n\t":
            var v40 int64 = 1
            _ = v40
        case "hello":
            fmt.Println(v12, v39, v12)
        default:
            if (uint16(v13) != 1) {
                var v41 float32 = 3.14159
                _ = v41
            }
        }
    case 42:
        _ = len(v21)
    default:
        v42 := GenS2{F0: 2147483647, F1: v39, F2: false}
        _ = v42
    }
    if false {
        v4 -= (v4 << ((65535) & 63))
        var v43 uint8 = v39
        _ = v43
    } else {
        v12 += (v12 + (((v12 + -2.25) / 1.5) * v12))
        if (848 != 0) {
            v4 += -1
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v39
    _ = v4
}
