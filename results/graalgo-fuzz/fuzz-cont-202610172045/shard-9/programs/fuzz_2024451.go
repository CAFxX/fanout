package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int8, p1 int) int64 {
    p0 = 0
    var v7 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]int)
    var v3 uint64 = 42
    v4 := make(map[string]bool)
    v5 := []uint32{}
    v6 := []int{736}
    fmt.Println(v3)
    v3 = (v3 * (^(v3)))
    if false {
        if true {
            for fi8 := 0; fi8 < 17; fi8++ {
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 int8 = -128
                    _ = fi9
                    _ = v10
                }
                var v11 int64 = 9223372036854775807
                _ = fi8
                _ = v11
            }
        }
    }
    v12 := GenS1{F0: 1}
    v13 := v2[-2147483648]
    v14 := GenS2{F0: 14, F1: 2, F2: -32768}
    v1 = v1
    switch -2147483648 {
    case -2147483648:
        for fi15 := 0; fi15 < 18; fi15++ {
            _ = len(v2)
            _ = fi15
        }
    default:
        v1 = 2
    }
    if false {
        switch ("" + "\n\t") {
        case "foo bar":
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
        case "hello":
            _ = len(v5)
        case "":
            v3 -= v3
        }
        v16 := GenS2{F0: 2, F1: 0, F2: int16(v3)}
        v17 := GenS2{F0: 255, F1: 42, F2: -1}
        _ = v16
        _ = v17
    } else {
        var v18 uint = 456
        if (0 >= (^(v1))) {
            for fi19 := 0; fi19 < 5; fi19++ {
                for fi20 := 0; fi20 < 8; fi20++ {
                    fmt.Println(v18)
                    var v21 uint16 = 798
                    v18 = 42
                    if (((v21 & 505) >= v21) && (-1 < 1)) {
                        var v22 uint = v18
                        var v23 float32 = ((3.14159 - -2.25) - ((float32(((-2.25 * 3.14159) - ((0.1 + float64(3.14159)) - (1.5 * (1.5 * (0.1 - float64(1.5))))))) / (float32(0.1) * 1e10)) + 1e10))
                        _ = v22
                        _ = v23
                    } else {
                        var v24 int16 = 1
                        var v25 bool = (true && true)
                        v26 := GenS3{F0: -1, F1: 9223372036854775807}
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    _ = fi20
                    _ = v21
                }
                _ = fi19
            }
            v4["hello"] = (0.0 >= 1e10)
            v27 := GenS2{F0: 65535, F1: 42, F2: 1}
            _ = v27
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            fmt.Println(v18)
        }
        _ = v18
    }
    if (85 >= 1) {
        for fi28 := 0; fi28 < 1; fi28++ {
            v13 = 1
            v13 -= fi28
            v5 = append(v5, 4294967295)
            _ = fi28
        }
        v29 := GenS4{F0: 65535}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
