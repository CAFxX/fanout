package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int64; F2 int32 }
func genFn0(p0 string, p1 uint8) uint8 {
    v6 := []int{127}
    if false {
        var v7 int = 127
        _ = v7
    } else {
        for fi8 := 0; fi8 < 4; fi8++ {
            p0 = p0
            _ = fi8
        }
        var v9 uint64 = 721
        _ = v9
    }
    p0 = p0
    var v10 uint = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return 12
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (((32767 <= 1000) || false) || false)
    var v3 int8 = 100
    v4 := make(map[int]bool)
    var v5 uint32 = 2
    var v11 uint = 18446744073709551615
    var v12 int64 = -9223372036854775808
    var v13 int32 = 0
    v14 := []int{-137, 127, 128}
    if (v3 < 42) {
        v15 := GenS1{F0: -1}
        v2 = (false && (42 <= 18446744073709551615))
        var v16 uint64 = uint64(v11)
        _ = v15
        _ = v16
    }
    v2 = (2 < 2)
    if v2 {
        fmt.Println(v3, v2)
        fmt.Println(v2, v13)
        var v17 uint64 = 417
        _ = v17
    } else {
        v18 := GenS1{F0: 512}
        _ = v18
    }
    if v2 {
        if ((-(v12)) != (v12 * v12)) {
            v5 += (^(v5))
            if v2 {
                if true {
                    v12 = v12
                    var v19 int = 4294967295
                    _ = v19
                }
                var v20 int8 = v3
                _ = v20
            } else {
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 int16 = -461
                    _ = fi21
                    _ = v22
                }
                v3 = 127
                v23 := GenS1{F0: v12}
                _ = v23
            }
            if ((float64(0.1) * 0.0) == 0.0) {
                v24 := GenS1{F0: v12}
                var v25 float64 = -2.25
                _ = v24
                _ = v25
            }
        } else {
            fmt.Println(v5, v3)
            for fi26 := 0; fi26 < 4; fi26++ {
                _ = len(v14)
                _ = fi26
            }
        }
        var v27 int = -1
        var v28 uint64 = 562
        _ = v27
        _ = v28
    } else {
        fmt.Println(v13, v2, v13)
        v29 := GenS2{F0: (v5 & v5), F1: (v12 >> ((256) & 63)), F2: (v13 ^ v13)}
        _ = v29
    }
    fmt.Println(v11, v11, v5)
    if (!v2) {
        var v30 uint8 = uint8(v11)
        _ = v30
    } else {
        if (!v2) {
            var v31 bool = (-2.25 <= float32(0.0))
            var v32 uint64 = 971
            _ = v31
            _ = v32
        } else {
            if (float32(0.0) == (1e10 - ((float32((((1.5 * (0.0 / ((1e10 * 0.1) / (-2.25 * float64(1e10))))) * 1e10) + 0.1)) + -2.25) - 1.5))) {
                for fi33 := 0; fi33 < 13; fi33++ {
                    for fi34 := 0; fi34 < 11; fi34++ {
                        var v35 uint16 = 607
                        var v36 uint32 = uint32(v35)
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    _ = fi33
                }
            }
            switch "\n\t" {
            case "a":
                var v37 uint16 = 0
                _ = v37
            default:
                if false {
                    v38 := GenS1{F0: (v12 >> ((4294967295) & 63))}
                    for fi39 := 0; fi39 < 10; fi39++ {
                        var v40 int32 = v13
                        var v41 int = (fi39 % (((fi39 | -338)) | 1))
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    var v42 uint64 = 2
                    _ = v38
                    _ = v42
                } else {
                    fmt.Println(v2, v11)
                    for fi43 := 0; fi43 < 10; fi43++ {
                        var v44 float32 = 1e10
                        v4[fi43] = (!((v3 ^ v3) != (-(v3))))
                        var v45 int64 = v12
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    v5 -= 111
                }
            }
            var v46 bool = (102 <= 37)
            _ = v46
        }
        v2 = true
    }
    var v47 uint16 = 42
    if v2 {
        fmt.Println(v47)
    }
    if len(v14) > 0 {
        v14[0] = 2147483647
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
