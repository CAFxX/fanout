package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 bool; F1 int; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int64; F1 int }
func genFn0(p0 int64, p1 int8) uint16 {
    var v4 float32 = 0.0
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := []int{4294967295, (fi5 * (fi5 & fi5)), 32767, 42}
        fmt.Println(p1, fi5)
        switch fi5 {
        case 1:
            _ = len(v6)
        }
        _ = fi5
        _ = v6
    }
    p0 += (-(p0))
    v7 := GenS1{F0: 9223372036854775807, F1: false}
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 2
}

func main() {
    v1 := []int{}
    v2 := make(map[int]bool)
    var v3 bool = false
    fmt.Println(v3)
    var v8 bool = v3
    v9 := GenS2{F0: (4294967295 >= 4294967295), F1: 4294967295, F2: 42}
    var v10 uint64 = 1
    if ((42 > 18446744073709551615) && (!((0 < 127) || (!(v3 || true))))) {
        var v11 float64 = 0.1
        v12 := GenS1{F0: -976, F1: (v3 && false)}
        v13 := GenS3{F0: 0}
        _ = v11
        _ = v12
        _ = v13
    } else {
        v3 = v3
    }
    var v14 uint64 = (v10 - v10)
    _ = len(v2)
    v3 = (-2147483648 == 1)
    v8 = v3
    v14 = 42
    if (uint32(v14) != 269) {
        switch ("a" + "a") {
        case "\n\t":
            switch ("\n\t" + "hello") {
            case "\n\t":
                v14 -= 941
            default:
                v2[128] = (((1e10 - 0.1) <= 3.14159) && ((42 <= 0) || ((((0.1 - (1.5 * 1e10)) + float64(-2.25)) * 1.5) == -2.25)))
            }
        case "hello":
            var v15 int = 32768
            _ = v15
        default:
            for fi16 := 0; fi16 < 13; fi16++ {
                switch fi16 {
                case 32768:
                    v17 := GenS4{F0: "foo bar"}
                    _ = v17
                case 42:
                    v14 = (v14 >> ((42) & 63))
                case 127:
                    for fi18 := 0; fi18 < 13; fi18++ {
                        _ = len(v1)
                        var v19 float32 = float32(((1.5 / float64((1.5 + float32(float64(1.5))))) * float64((1e10 * 1.5))))
                        _ = fi18
                        _ = v19
                    }
                }
                switch "a" {
                case "\n\t":
                    v20 := GenS5{F0: 255, F1: (!v8)}
                    _ = v20
                case "":
                    v21 := GenS6{F0: 1, F1: 683}
                    _ = v21
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    v3 = v8
                }
                _ = fi16
            }
        }
    } else {
        v14 = v10
        if v8 {
            v8 = (-1 <= -2147483648)
            v1 = append(v1, 32768)
            v22 := GenS6{F0: 1, F1: 0}
            _ = v22
        } else {
            v23 := GenS7{F0: -9223372036854775808, F1: -344}
            if len(v1) > 0 {
                v1[0] = 554
            }
            _ = v23
        }
    }
    if (!(!v8)) {
        if len(v1) > 0 {
            v1[0] = 923
        }
    }
    _ = len(v2)
    var v24 int32 = -814
    for fi25 := 0; fi25 < 16; fi25++ {
        v26 := GenS2{F0: true, F1: 944, F2: v10}
        _ = fi25
        _ = v26
    }
    var v27 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v8
    _ = v9
}
