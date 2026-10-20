package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint8, p2 string) string {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 6; fi7++ {
        p2 = p2
        _ = fi7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        if true {
            var v9 int64 = -800
            v10 := GenS1{F0: 0}
            fmt.Println(fi8, p2, p0)
            _ = v10
            _ = v9
        } else {
            v11 := GenS1{F0: 0}
            if (false && true) {
                var v12 uint32 = 0
                for fi13 := 0; fi13 < 7; fi13++ {
                    if true {
                        v14 := []int{(^(fi8)), (fi13 | fi13), (fi13 ^ 128), (fi13 - 128)}
                        switch 0 {
                        case 4294967295:
                            switch fi8 {
                            case 0:
                                if (0.0 < (float64(1.5) + 0.1)) {
                                    p1 = (^(p1))
                                }
                            case 128:
                                var v15 int16 = 42
                                _ = v15
                            default:
                                v14 = append(v14, fi8)
                            }
                        case -1:
                            fmt.Println(v12, fi8)
                        default:
                            var v16 uint16 = 1
                            _ = v16
                        }
                        _ = v14
                    } else {
                        var v17 string = (p2 + p2)
                        var v18 string = v17
                        v19 := GenS1{F0: 72}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    var v20 int = 42
                    _ = fi13
                    _ = v20
                }
                _ = v12
            }
            _ = v11
        }
        _ = fi8
    }
    var v21 float64 = ((3.14159 + float64(0.1)) + 0.1)
    v22 := GenS2{F0: -2147483648, F1: (p0 % ((p0) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return "\n\t"
}

func main() {
    var v1 int32 = -2147483648
    var v2 int = 4294967295
    var v3 int64 = -1
    var v4 int64 = v3
    v5 := make(map[uint32]bool)
    var v6 int = v2
    v2 -= v2
    var v23 int16 = 0
    v6 = v6
    for fi24 := 0; fi24 < 8; fi24++ {
        switch (^(v2)) {
        case 1:
            v25 := GenS2{F0: 62, F1: (v4 >> ((0) & 63))}
            _ = v25
        case 4294967295:
            if true {
                v26 := GenS3{F0: (fi24 >> ((2) & 63)), F1: 370, F2: 4294967295}
                _ = v26
            } else {
                if (!(true || (v2 < (v2 % (((fi24 >> ((2) & 63))) | 1))))) {
                    v1 = v1
                }
                var v27 bool = ((1.5 - (float32(-2.25) - ((float32((1.5 - 1.5)) - ((float32(1e10) * 1.5) - 1.5)) / -2.25))) < (3.14159 - 3.14159))
                _ = v27
            }
        case -1:
            switch "foo bar" {
            case "hello":
                v4 = 855
            case "a":
                v23 = (v23 * v23)
            case "foo bar":
                var v28 uint8 = 42
                _ = v28
            default:
                v6 = (fi24 << ((2) & 63))
            }
        }
        v29 := GenS2{F0: v6, F1: (^(v4))}
        _ = fi24
        _ = v29
    }
    v30 := GenS3{F0: (v2 * (v2 + v6)), F1: 4294967295, F2: 42}
    fmt.Println(v4, v3, v2)
    v31 := []int{}
    if false {
        v23 += v23
        if true {
            v32 := GenS1{F0: 18446744073709551615}
            _ = v32
        }
        if true {
            for fi33 := 0; fi33 < 15; fi33++ {
                v34 := GenS2{F0: (fi33 - (v6 * (v6 + (v6 % ((fi33) | 1))))), F1: v3}
                _ = fi33
                _ = v34
            }
            var v35 uint8 = 125
            _ = v35
        }
    } else {
        v36 := GenS4{F0: 2}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
