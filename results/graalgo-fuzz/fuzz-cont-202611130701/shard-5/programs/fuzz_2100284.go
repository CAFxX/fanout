package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 uint }
func genFn0() int {
    var v7 uint32 = 2147483648
    var v8 uint64 = 42
    if true {
        v8 = 1
        v7 = (^(v7))
    } else {
        v7 -= 2
        var v9 float64 = 1.5
        _ = v9
    }
    var v10 uint16 = 0
    _ = v10
    _ = v7
    _ = v8
    return 32767
}
func genFn1() uint64 {
    var v11 int8 = 1
    var v12 int64 = 9223372036854775807
    var v13 uint8 = 173
    _ = v11
    _ = v12
    _ = v13
    return 163
}

func main() {
    var v1 uint16 = 2
    var v2 int32 = 42
    v3 := make(map[int]bool)
    var v4 float64 = (0.1 * (1e10 - 0.1))
    v5 := []int{32768, -414, int(v2)}
    v6 := make(map[int]uint32)
    var v14 bool = (!true)
    fmt.Println(v4)
    for fi15 := 0; fi15 < 2; fi15++ {
        if true {
            switch genFn0() {
            case -368:
                if len(v5) > 0 {
                    v5[0] = fi15
                }
            case 32768:
                for fi16 := 0; fi16 < 17; fi16++ {
                    v4 = v4
                    v14 = (-636 >= 0)
                    _ = fi16
                }
            case 127:
                var v17 int = (fi15 + -649)
                _ = v17
            }
            var v18 int = 127
            if ((v4 / v4) > 1e10) {
                var v19 string = "a"
                _ = v19
            }
            _ = v18
        } else {
            var v20 uint8 = 42
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 int8 = -64
                _ = fi21
                _ = v22
            }
            _ = v20
        }
        v14 = v14
        var v23 int64 = 13
        _ = fi15
        _ = v23
    }
    v24 := GenS1{F0: v2, F1: 128, F2: (v1 % ((1) | 1))}
    v25 := GenS1{F0: (-(v2)), F1: 586, F2: 1}
    if (100 >= 127) {
        v26 := GenS1{F0: (v2 % (((v2 & 1)) | 1)), F1: -577, F2: v1}
        _ = v26
    }
    if len(v5) > 0 {
        v5[0] = -1
    }
    var v27 uint16 = v1
    switch 9223372036854775807 {
    case -1:
        for fi28 := 0; fi28 < 2; fi28++ {
            fmt.Println(v1)
            _ = fi28
        }
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        if true {
            var v30 bool = false
            if v14 {
                if ((681 < (v2 * (v2 >> ((370) & 63)))) && v14) {
                    delete(v3, (-(fi29)))
                } else {
                    v31 := GenS1{F0: v2, F1: genFn0(), F2: (^(v27))}
                    v32 := GenS1{F0: v2, F1: fi29, F2: v1}
                    _ = v31
                    _ = v32
                }
            }
            v33 := GenS2{F0: fi29, F1: int64(fi29), F2: 0}
            _ = v30
            _ = v33
        } else {
            if false {
                fmt.Println(v4, v2)
            } else {
                fmt.Println(v2)
            }
            if (!v14) {
                for fi34 := 0; fi34 < 2; fi34++ {
                    fmt.Println(v2, fi34, v27)
                    var v35 int16 = -79
                    var v36 int = (fi34 | fi34)
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                var v37 uint32 = 610
                v38 := v3[fi29]
                _ = v37
                _ = v38
            } else {
                var v39 uint64 = 1
                var v40 float64 = float64(3.14159)
                _ = v39
                _ = v40
            }
            var v41 int64 = -9223372036854775808
            _ = v41
        }
        var v42 uint8 = 1
        _ = fi29
        _ = v42
    }
    fmt.Println(v4)
    _ = len(v3)
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
