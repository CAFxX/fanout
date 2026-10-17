package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 string; F1 uint; F2 int8 }
type GenS4 struct { F0 int; F1 int16 }
func genFn0(p0 uint32, p1 uint16, p2 int64) int8 {
    p0 += p0
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{(fi7 >> ((1) & 63)), (fi7 / ((fi7) | 1)), (fi7 * (fi7 >> ((18446744073709551615) & 63))), 127}
        p0 += (p0 ^ (p0 * p0))
        v9 := GenS1{F0: p2, F1: ("" + "\n\t"), F2: -1}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    var v3 int8 = -1
    v4 := make(map[uint32]string)
    v5 := []uint32{2, 4294967295, 0, 2147483647}
    v6 := []int{1, 2147483647}
    fmt.Println(v3)
    if false {
        var v10 int32 = 0
        fmt.Println(v3, v10)
        var v11 uint16 = 255
        _ = v10
        _ = v11
    }
    var v12 float32 = (0.0 + ((0.0 - (float32(3.14159) - (1e10 + float32(1.5)))) * 0.1))
    var v13 uint16 = 255
    fmt.Println(v12, v3)
    switch 331 {
    case -590:
        for fi14 := 0; fi14 < 4; fi14++ {
            if false {
                if false {
                    var v15 int64 = 9223372036854775807
                    v13 = (^(v13))
                    for fi16 := 0; fi16 < 20; fi16++ {
                        var v17 bool = (v13 != 42)
                        if ((!v17) || (!false)) {
                            var v18 int64 = (v15 << ((42) & 63))
                            _ = v18
                        }
                        v19 := GenS2{F0: 574, F1: 0}
                        v5 = append(v5, 227)
                        var v20 uint32 = 88
                        v21 := v2["hello"]
                        _ = fi16
                        _ = v17
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    _ = v15
                } else {
                    var v22 uint = 255
                    fmt.Println(v3)
                    _ = v22
                }
                var v23 float32 = v12
                var v24 uint64 = 0
                _ = v23
                _ = v24
            }
            v12 = (1.5 / float32(float64(3.14159)))
            switch "" {
            case "a":
                if (fi14 <= (fi14 | (-(fi14)))) {
                    v25 := GenS3{F0: ("hello" + "\n\t"), F1: uint(v13), F2: (v3 | v3)}
                    var v26 float64 = 0.1
                    _ = v25
                    _ = v26
                } else {
                    switch ("foo bar" + "a") {
                    case "a":
                        v13 = 404
                    }
                    var v27 uint32 = 2
                    v5 = append(v5, 2)
                    _ = v27
                }
            case "":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi14
        }
    default:
        var v28 int32 = -2147483648
        _ = v28
    }
    v13 = 932
    fmt.Println(v12, v3)
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 string = "hello"
        _ = fi29
        _ = v30
    }
    var v31 uint64 = 18446744073709551615
    for fi32 := 0; fi32 < 6; fi32++ {
        var v33 uint8 = 127
        var v34 uint8 = v33
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS4{F0: 807, F1: 1000}
    var v36 uint8 = 0
    v36 -= v36
    if (!(0 < 42)) {
        var v37 uint8 = 1
        var v38 string = (("a" + "\n\t") + (("\n\t" + "\n\t") + "\n\t"))
        var v39 float64 = 0.1
        _ = v37
        _ = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        switch fi40 {
        case -9223372036854775808:
            v41 := GenS4{F0: 991, F1: -1}
            _ = v41
        case 32768:
            _ = len(v4)
        }
        fmt.Println(v13)
        _ = fi40
    }
    _ = len(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
