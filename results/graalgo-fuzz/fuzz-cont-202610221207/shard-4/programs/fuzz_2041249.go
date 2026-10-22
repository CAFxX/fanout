package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int32; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint32; F1 int16 }
func genFn0(p0 int8, p1 int16) int8 {
    if false {
        for fi4 := 0; fi4 < 4; fi4++ {
            if ((-(fi4)) != fi4) {
                if false {
                    p1 = (p1 % ((42) | 1))
                    for fi5 := 0; fi5 < 3; fi5++ {
                        var v6 int32 = 2147483647
                        if ((p1 + (p1 | p1)) > p1) {
                            v7 := []int{128, (fi5 << ((1) & 63)), fi4}
                            _ = len(v7)
                            if (((0.1 - (float64(1e10) * -2.25)) == float64((0.0 + -2.25))) && (65535 != 1)) {
                                if len(v7) > 0 {
                                    v7[0] = (fi5 * 2147483647)
                                }
                            } else {
                                var v8 uint = 1
                                _ = v8
                            }
                            _ = v7
                        }
                        var v9 uint16 = 42
                        var v10 int64 = 61
                        var v11 bool = true
                        switch 4294967295 {
                        case 4294967295:
                            var v12 int16 = (p1 | (p1 % ((p1) | 1)))
                            _ = v12
                        case 1:
                            if false {
                                fmt.Println(fi5, fi4)
                            }
                        default:
                            if v11 {
                                var v13 uint64 = 42
                                var v14 int32 = 129
                                _ = v13
                                _ = v14
                            }
                        }
                        fmt.Println(v10, v9)
                        _ = fi5
                        _ = v10
                        _ = v11
                        _ = v6
                        _ = v9
                    }
                    fmt.Println(fi4)
                } else {
                    fmt.Println(fi4, p0)
                }
                switch ("hello" + "") {
                case "":
                    var v15 int64 = -180
                    _ = v15
                case "a":
                    v16 := []int{}
                    _ = v16
                }
                fmt.Println(p1, p1)
            } else {
                p1 = (p1 * p1)
                var v17 int16 = p1
                v18 := GenS1{F0: 2}
                _ = v17
                _ = v18
            }
            v19 := []int{(fi4 * (fi4 | fi4))}
            _ = fi4
            _ = v19
        }
    } else {
        var v20 uint = 2
        p1 = -1
        _ = v20
    }
    v21 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = v21
    return p0
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 0
    var v3 int32 = -1
    var v22 int16 = -32768
    v23 := GenS2{F0: 18446744073709551615, F1: 0}
    var v24 uint = 1
    for fi25 := 0; fi25 < 7; fi25++ {
        v26 := []int{42}
        v22 = v22
        _ = fi25
        _ = v26
    }
    v27 := []int{}
    fmt.Println(v24)
    v28 := GenS3{F0: (v3 % (((v3 * -2147483648)) | 1)), F1: 1, F2: v2}
    var v29 bool = true
    var v30 uint8 = 2
    v31 := GenS4{F0: 4294967295, F1: (v22 + v22)}
    fmt.Println(v22, v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
