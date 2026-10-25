package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 8; fi7++ {
        if (!(true || true)) {
            var v8 uint = 2
            _ = v8
        }
        if (!true) {
            var v9 uint16 = 1
            for fi10 := 0; fi10 < 16; fi10++ {
                if true {
                    fmt.Println(fi7)
                    v11 := []int{(fi7 ^ fi7), 4294967295, fi10}
                    switch "" {
                    case "\n\t":
                        if len(v11) > 0 {
                            v11[0] = fi7
                        }
                    default:
                        v9 += 311
                    }
                    _ = v11
                }
                if (!true) {
                    v12 := GenS1{F0: 1}
                    v9 -= (^(v9))
                    switch fi10 {
                    case -1:
                        v9 = v9
                        fmt.Println(fi10, fi7, v9)
                        v9 += (v9 * v9)
                        var v13 uint32 = 999
                        _ = v13
                    default:
                        v14 := []int{128, (fi7 >> ((761) & 63)), fi7}
                        _ = v14
                    }
                    fmt.Println(fi7, v9, fi7)
                    _ = v12
                } else {
                    var v15 uint32 = 489
                    _ = v15
                }
                v16 := []int{(fi10 ^ fi10), fi10, (fi10 + (fi10 | fi10)), fi7}
                _ = fi10
                _ = v16
            }
            fmt.Println(fi7)
            _ = v9
        } else {
            var v17 uint32 = 0
            var v18 float32 = 1e10
            _ = v17
            _ = v18
        }
        _ = fi7
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v20 := []int{}
        _ = fi19
        _ = v20
    }
    var v21 float64 = 0.0
    fmt.Println(v21)
    v21 = v21
    _ = v21
    return 1
}

func main() {
    var v1 string = "foo bar"
    var v2 uint = 65535
    var v3 bool = true
    var v4 int64 = 553
    var v5 string = (v1 + "")
    var v6 uint64 = 1
    v4 += int64(v2)
    for fi22 := 0; fi22 < 13; fi22++ {
        v5 = v5
        _ = fi22
    }
    v23 := GenS2{F0: "", F1: -958}
    v24 := GenS2{F0: ((("\n\t" + ("" + ("hello" + (v1 + (v5 + ("" + v1)))))) + v1) + v1), F1: 521}
    v25 := []int{-1, -696}
    v26 := GenS2{F0: v5, F1: 287}
    v27 := GenS1{F0: -2147483648}
    if (!(!(genFn0() == 921))) {
        switch -561 {
        case -9223372036854775808:
            if (v3 && (1.5 < 0.0)) {
                var v28 uint16 = 65535
                v29 := GenS2{F0: v1, F1: 1}
                _ = v28
                _ = v29
            } else {
                for fi30 := 0; fi30 < 18; fi30++ {
                    v4 = -9223372036854775808
                    _ = fi30
                }
                v2 += uint(v6)
            }
        case -1:
            v5 = v1
        case -2147483648:
            if false {
                v3 = v3
            }
        }
        var v31 uint64 = v6
        v31 -= v6
        _ = v31
    } else {
        var v32 bool = (1 == -32768)
        v4 = int64(v6)
        fmt.Println(v2, v6, v4)
        _ = v32
    }
    v6 = (v6 / ((v6) | 1))
    if (0.0 != 3.14159) {
        var v33 uint = v2
        v1 = ""
        if (v33 <= 665) {
            v5 = v5
        } else {
            v34 := GenS1{F0: 1}
            v33 = (v33 % (((v2 + 963)) | 1))
            var v35 int64 = v4
            _ = v34
            _ = v35
        }
        _ = v33
    } else {
        v2 = (v2 ^ v2)
    }
    var v36 int32 = -476
    var v37 bool = (!true)
    v5 = v1
    fmt.Println(v37, v1)
    v38 := GenS3{F0: 0, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
