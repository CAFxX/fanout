package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int32; F1 uint32; F2 int64 }
func genFn0(p0 int32) uint32 {
    v7 := GenS1{F0: 0, F1: 1, F2: 2}
    v8 := []int{42, 32767, 128, 32767}
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []uint32{2147483647, 480, 2147483647, 138}
    v2 := make(map[int]uint32)
    var v3 int = 9223372036854775807
    var v4 uint8 = 127
    v5 := make(map[string]uint32)
    v6 := []int{2147483647, (v3 | (v3 & v3)), 0, v3}
    if false {
        if (false || (1 != 4294967295)) {
            if (-1 > 32767) {
                v4 -= 2
                var v9 float32 = 1.5
                _ = v9
            }
            var v10 int16 = 1000
            _ = v10
        }
        var v11 int = (v3 >> ((1) & 63))
        _ = v11
    } else {
        v4 = v4
        if true {
            _ = len(v2)
        } else {
            _ = len(v2)
        }
    }
    v1 = append(v1, 549)
    if true {
        v3 += (^(v3))
        v5["a"] = 167
        fmt.Println(v4, v3)
    }
    v12 := GenS1{F0: 1, F1: 65535, F2: 2}
    if true {
        v3 -= (^(v3))
        fmt.Println(v3, v3)
    }
    v13 := GenS2{F0: (v4 / ((v4) | 1)), F1: (-(v3))}
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint = 4294967295
        for fi16 := 0; fi16 < 15; fi16++ {
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                if ((^(fi14)) >= (fi14 | fi14)) {
                    switch (fi16 & (^(fi16))) {
                    case 32768:
                        v15 += v15
                    }
                    fmt.Println(v3)
                    fmt.Println(fi14, fi14)
                }
            case "a":
                v17 := GenS2{F0: (v4 ^ (v4 * (v4 % ((v4) | 1)))), F1: fi14}
                _ = v17
            case "foo bar":
                _ = len(v1)
            default:
                var v18 int = (v3 | (-(fi14)))
                _ = v18
            }
            fmt.Println(v15, fi16, v15)
            var v19 uint = (v15 / ((v15) | 1))
            if (v3 == (fi16 % ((fi14) | 1))) {
                v20 := GenS1{F0: 32767, F1: 712, F2: 4294967295}
                _ = v20
            }
            v21 := GenS3{F0: 2147483647, F1: 4294967295, F2: 0}
            _ = fi16
            _ = v19
            _ = v21
        }
        _ = fi14
        _ = v15
    }
    var v22 uint = 312
    if ((-128 < 0) && (!true)) {
        v4 = (v4 % ((v4) | 1))
        for fi23 := 0; fi23 < 19; fi23++ {
            v22 += (v22 % ((v22) | 1))
            fmt.Println(v4, v22)
            _ = fi23
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := GenS3{F0: -1, F1: 1, F2: 814}
            fmt.Println(v22, v22)
            if (800 < -1) {
                var v26 bool = (!false)
                _ = v26
            } else {
                if (1.5 != 3.14159) {
                    v27 := GenS2{F0: 110, F1: v3}
                    v4 = 42
                    _ = v27
                } else {
                    v28 := v2[fi24]
                    _ = v28
                }
            }
            _ = fi24
            _ = v25
        }
    }
    fmt.Println(v3, v4, v22)
    v29 := GenS3{F0: -850, F1: 916, F2: 1}
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 int8 = 0
        for fi32 := 0; fi32 < 16; fi32++ {
            v33 := GenS3{F0: -1, F1: 2147483648, F2: 58}
            v34 := GenS1{F0: 0, F1: 42, F2: 666}
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi30
        _ = v31
    }
    var v35 uint64 = 582
    var v36 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
