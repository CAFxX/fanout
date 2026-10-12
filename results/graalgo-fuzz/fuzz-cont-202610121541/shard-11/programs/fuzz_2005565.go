package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int32; F1 uint8; F2 int16 }
type GenS3 struct { F0 bool }
func genFn0() uint64 {
    var v5 uint64 = 2
    v5 -= (^(v5))
    if false {
        var v6 int32 = -2147483648
        for fi7 := 0; fi7 < 17; fi7++ {
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := GenS1{F0: "", F1: 0}
                if ("hello" >= "\n\t") {
                    v6 -= v6
                }
                var v10 int32 = 42
                _ = fi8
                _ = v10
                _ = v9
            }
            v6 -= 0
            _ = fi7
        }
        _ = v6
    } else {
        var v11 int16 = -32768
        var v12 float64 = 0.0
        _ = v11
        _ = v12
    }
    _ = v5
    return v5
}
func genFn1(p0 uint32) uint32 {
    var v13 uint = 18446744073709551615
    if (0 > 100) {
        v14 := GenS2{F0: -1, F1: 127, F2: 409}
        switch 127 {
        case -2147483648:
            var v15 uint8 = 206
            _ = v15
        case 32768:
            v13 += 1
        case 127:
            switch "hello" {
            case "":
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := []int{fi16}
                    var v18 uint = v13
                    v17 = append(v17, (fi16 % (((-(fi16))) | 1)))
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            case "foo bar":
                v19 := []int{944, 127}
                _ = v19
            case "\n\t":
                v13 = v13
            }
        default:
            v20 := GenS2{F0: 1, F1: 2, F2: -21}
            _ = v20
        }
        _ = v14
    }
    v21 := []int{2147483647, 42, 128, 9223372036854775807}
    _ = p0
    _ = v13
    _ = v21
    return (^(p0))
}

func main() {
    v1 := []int{1, -106, -1, 1}
    var v2 int32 = 525
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    fmt.Println(v2)
    for fi22 := 0; fi22 < 11; fi22++ {
        if false {
            var v23 int16 = -32768
            for fi24 := 0; fi24 < 5; fi24++ {
                fmt.Println(fi22)
                var v25 float32 = (-2.25 - 1e10)
                v26 := GenS1{F0: ("hello" + ("foo bar" + "\n\t")), F1: (^(v23))}
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v23
        }
        switch "\n\t" {
        case "a":
            var v27 uint8 = 42
            _ = v27
        case "foo bar":
            var v28 uint = 198
            _ = v28
        case "hello":
            var v29 int64 = 0
            _ = v29
        default:
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 int32 = 945
                v32 := GenS2{F0: v2, F1: 128, F2: -1}
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        _ = fi22
    }
    v4[196] = 0
    v1 = append(v1, 9223372036854775807)
    var v33 uint8 = 0
    for fi34 := 0; fi34 < 1; fi34++ {
        v35 := GenS3{F0: false}
        fmt.Println(v2, fi34, fi34)
        for fi36 := 0; fi36 < 14; fi36++ {
            fmt.Println(fi34, v33, v33)
            var v37 int32 = (-(v2))
            fmt.Println(fi36, fi34)
            _ = fi36
            _ = v37
        }
        _ = fi34
        _ = v35
    }
    var v38 int32 = v2
    fmt.Println(v2, v2)
    for fi39 := 0; fi39 < 17; fi39++ {
        switch -328 {
        case -764:
            v33 -= 195
        case -2147483648:
            if false {
                for fi40 := 0; fi40 < 13; fi40++ {
                    if true {
                        var v41 uint64 = 0
                        v38 = 0
                        _ = v41
                    }
                    _ = fi40
                }
            }
        default:
            v33 -= 0
        }
        var v42 float32 = -2.25
        _ = fi39
        _ = v42
    }
    fmt.Println(v33, v33, v2)
    _ = len(v3)
    delete(v4, 198)
    for fi43 := 0; fi43 < 17; fi43++ {
        v33 = (^(v33))
        v38 = (v2 - v2)
        _ = fi43
    }
    delete(v4, 128)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v4
}
