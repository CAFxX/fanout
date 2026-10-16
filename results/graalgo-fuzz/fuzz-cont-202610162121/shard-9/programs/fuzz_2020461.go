package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint32 }
func genFn0() int16 {
    var v5 int16 = -268
    if true {
        v6 := GenS1{F0: -126, F1: (v5 + (v5 % ((v5) | 1))), F2: 2147483647}
        for fi7 := 0; fi7 < 20; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                fmt.Println(fi8, v5, fi7)
                v9 := []int{fi8, fi8}
                v10 := GenS1{F0: -401, F1: (v5 + (v5 << ((589) & 63))), F2: 2147483647}
                _ = fi8
                _ = v10
                _ = v9
            }
            switch (fi7 >> ((255) & 63)) {
            case -2147483648:
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 int8 = 1
                    v13 := []int{(^(fi11)), (^(fi11)), (fi11 / (((fi11 ^ -9223372036854775808)) | 1))}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 1:
                v5 += v5
            default:
                fmt.Println(v5, v5)
            }
            _ = fi7
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            switch 973 {
            case 123:
                var v15 string = "hello"
                _ = v15
            case 0:
                if (0 < -770) {
                    fmt.Println(v5, fi14)
                    v5 = (v5 >> ((424) & 63))
                    for fi16 := 0; fi16 < 16; fi16++ {
                        var v17 int8 = 14
                        _ = fi16
                        _ = v17
                    }
                }
            default:
                v18 := []int{(fi14 % ((-968) | 1)), fi14}
                _ = v18
            }
            v19 := GenS1{F0: 9223372036854775807, F1: -32768, F2: -133}
            _ = fi14
            _ = v19
        }
        _ = v6
    } else {
        var v20 float32 = 3.14159
        _ = v20
    }
    v21 := GenS2{F0: -819, F1: 42}
    switch 2147483647 {
    case 999:
        v22 := []int{}
        _ = v22
    case -1:
        if (false && (false && true)) {
            var v23 int = -1
            v24 := GenS2{F0: 0, F1: 14}
            v5 = (^(v5))
            _ = v23
            _ = v24
        }
    case 32768:
        fmt.Println(v5)
    }
    _ = v21
    _ = v5
    return (v5 * (v5 * (v5 & 32767)))
}

func main() {
    v1 := []int{}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 uint16 = 1
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    delete(v3, 4294967295)
    v4 = v4
    var v25 uint16 = v4
    if true {
        var v26 uint64 = 211
        var v27 float64 = 3.14159
        _ = v26
        _ = v27
    }
    var v28 uint = 1
    v28 = 137
    fmt.Println(v28, v25, v4)
    if (0 > 0) {
        v29 := GenS3{F0: 2147483647}
        v30 := v2[("hello" + "\n\t")]
        _ = v29
        _ = v30
    }
    var v31 int32 = 2147483647
    if (!((4294967295 == 668) && false)) {
        for fi32 := 0; fi32 < 2; fi32++ {
            fmt.Println(v28, v25, v31)
            v31 += (-(v31))
            var v33 int64 = 1
            _ = fi32
            _ = v33
        }
        var v34 float32 = 1.5
        _ = v34
    } else {
        for fi35 := 0; fi35 < 11; fi35++ {
            var v36 string = ("" + "\n\t")
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 11; fi37++ {
            v25 -= 32767
            var v38 float64 = (3.14159 / (3.14159 * 0.1))
            for fi39 := 0; fi39 < 10; fi39++ {
                switch (fi37 ^ fi37) {
                case 4294967295:
                    if (!(-1 != 110)) {
                        v31 = (^(v31))
                        v2[""] = 65535
                    } else {
                        v25 = 2
                    }
                }
                fmt.Println(v31)
                var v40 float32 = (1e10 * 0.0)
                var v41 int = fi37
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = fi37
            _ = v38
        }
    }
    if (!true) {
        v4 = v25
    } else {
        var v42 uint16 = v4
        _ = v42
    }
    var v43 uint = (v28 - v28)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v43
}
