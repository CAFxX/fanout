package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int8; F2 int64 }
type GenS5 struct { F0 uint; F1 int16 }
func genFn0(p0 int, p1 uint32) int32 {
    v4 := GenS1{F0: -122, F1: (p1 ^ p1)}
    switch (^(p0)) {
    case 32767:
        switch (p0 % ((4294967295) | 1)) {
        case 4294967295:
            v5 := []int{-1, (-(p0))}
            _ = v5
        default:
            for fi6 := 0; fi6 < 1; fi6++ {
                v7 := GenS1{F0: 127, F1: (^(p1))}
                _ = fi6
                _ = v7
            }
        }
    default:
        v8 := []int{}
        _ = v8
    }
    var v9 float64 = 3.14159
    p1 = (p1 - (p1 / ((65535) | 1)))
    p0 -= (p0 % (((p0 / (((p0 - (p0 | 127))) | 1))) | 1))
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return -911
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    var v3 string = ""
    if false {
        v10 := GenS1{F0: 42, F1: 2}
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 float32 = (3.14159 + 0.0)
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v13 int = 978
    v3 = v3
    var v14 bool = (0.1 == (1e10 + 0.1))
    v15 := GenS1{F0: -119, F1: 4294967295}
    switch (v3 + (v3 + "\n\t")) {
    case "a":
        var v16 uint64 = 1
        _ = v16
    case "":
        v17 := GenS1{F0: -54, F1: 2147483648}
        _ = v17
    default:
        if false {
            v13 = (^(v13))
            fmt.Println(v3)
            v1 = append(v1, 488)
        }
    }
    v18 := GenS2{F0: 128, F1: v14, F2: 22}
    v19 := GenS2{F0: 255, F1: (42 == -32768), F2: 1}
    v13 -= v13
    var v20 uint8 = 255
    v21 := GenS3{F0: 362}
    v22 := GenS4{F0: 4294967295, F1: int8(v20), F2: -9223372036854775808}
    if ((3.14159 + 1e10) < 0.1) {
        v23 := GenS1{F0: -37, F1: 42}
        _ = v23
    }
    if (!(!v14)) {
        fmt.Println(v3, v14, v14)
        for fi24 := 0; fi24 < 10; fi24++ {
            _ = len(v1)
            v2[104] = 1
            _ = fi24
        }
        if v14 {
            var v25 uint8 = 214
            _ = v25
        } else {
            switch 717 {
            case 128:
                for fi26 := 0; fi26 < 19; fi26++ {
                    var v27 int16 = 32767
                    _ = fi26
                    _ = v27
                }
            default:
                if true {
                    var v28 int64 = -1
                    _ = len(v1)
                    _ = v28
                } else {
                    v13 += -424
                    var v29 uint = 873
                    _ = v29
                }
            }
            if ((-2.25 / 1e10) != -2.25) {
                var v30 int = (v13 - 747)
                for fi31 := 0; fi31 < 12; fi31++ {
                    for fi32 := 0; fi32 < 8; fi32++ {
                        var v33 string = v3
                        fmt.Println(v14, fi31, v13)
                        _ = fi32
                        _ = v33
                    }
                    if false {
                        v34 := GenS5{F0: 937, F1: 1000}
                        _ = v34
                    }
                    _ = fi31
                }
                _ = v30
            }
        }
    } else {
        _ = len(v1)
        var v35 string = "a"
        _ = v35
    }
    var v36 int8 = 42
    v3 = v3
    v13 = (v13 % ((442) | 1))
    if (v3 != v3) {
        var v37 int8 = (v36 + 100)
        v38 := GenS5{F0: 0, F1: 143}
        _ = v37
        _ = v38
    } else {
        switch 127 {
        case 493:
            v39 := v2[65535]
            _ = v39
        }
        v13 = (v13 * v13)
        switch (v13 ^ (v13 & v13)) {
        case -2147483648:
            fmt.Println(v3, v36, v20)
        case 42:
            var v40 int = v13
            _ = v40
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v36
}
