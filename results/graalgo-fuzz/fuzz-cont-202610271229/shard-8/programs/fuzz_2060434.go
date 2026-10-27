package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint64, p1 int32, p2 bool) int64 {
    var v4 uint64 = 745
    v5 := []int{int(p1)}
    switch "hello" {
    case "":
        for fi6 := 0; fi6 < 4; fi6++ {
            if ((1.5 / ((3.14159 / 1e10) / 1.0)) != 1.5) {
                var v7 int32 = (p1 | (^(p1)))
                _ = v7
            }
            _ = fi6
        }
    case "a":
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                v5 = append(v5, (fi8 + fi8))
                v4 -= (v4 & (v4 << ((4294967295) & 63)))
                switch ("\n\t" + "hello") {
                case "":
                    v10 := GenS1{F0: 18446744073709551615}
                    _ = v10
                default:
                    v11 := GenS1{F0: v4}
                    _ = v11
                }
                _ = fi9
            }
            if true {
                v12 := GenS1{F0: v4}
                _ = v12
            } else {
                var v13 int32 = p1
                v14 := GenS1{F0: 1}
                fmt.Println(p0, p1)
                _ = v13
                _ = v14
            }
            fmt.Println(v4, p0, v4)
            _ = fi8
        }
    case "hello":
        var v15 int32 = (p1 ^ (p1 % (((p1 - p1)) | 1)))
        _ = v15
    default:
        p2 = (v4 > 2)
    }
    var v16 float32 = 1.5
    var v17 uint32 = 986
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 int = -2147483648
    var v2 int64 = -603
    v3 := make(map[uint32]string)
    switch "a" {
    case "\n\t":
        var v18 string = ""
        _ = v18
    default:
        _ = len(v3)
    }
    var v19 int8 = 90
    var v20 uint = 2
    if false {
        v20 += v20
    } else {
        v21 := GenS1{F0: uint64(v20)}
        _ = v21
    }
    v2 += v2
    var v22 bool = true
    var v23 uint = 42
    for fi24 := 0; fi24 < 8; fi24++ {
        if v22 {
            v20 = v20
            if (386 == 65535) {
                if (65535 <= 56) {
                    v25 := GenS2{F0: v2}
                    var v26 float64 = -2.25
                    _ = v25
                    _ = v26
                } else {
                    if true {
                        var v27 uint32 = 82
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            fmt.Println(fi28, v2, fi24)
                            v22 = v22
                            _ = fi28
                        }
                    }
                }
            }
        } else {
            var v29 uint32 = 105
            v22 = (v29 > (v29 % (((^(v29))) | 1)))
            for fi30 := 0; fi30 < 19; fi30++ {
                v1 -= -194
                var v31 uint64 = 42
                var v32 int32 = -1
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v29
        }
        var v33 uint16 = 2
        _ = fi24
        _ = v33
    }
    v23 += v23
    if v22 {
        var v34 float64 = float64(float32(float64(0.0)))
        _ = v34
    }
    v35 := []int{v1, (v1 + (v1 - v1))}
    v36 := GenS2{F0: (v2 - v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v36
}
