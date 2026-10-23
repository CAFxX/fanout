package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 bool; F2 int32 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint8; F2 bool }
func genFn0(p0 int, p1 bool, p2 uint) int32 {
    p0 += (p0 >> ((p2) & 63))
    v4 := []int{(p0 + 0), (p0 << ((p2) & 63)), 9223372036854775807}
    p2 = p2
    v5 := GenS1{F0: 2147483647}
    if p1 {
        _ = len(v4)
    } else {
        v6 := GenS1{F0: 251}
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS2{F0: 42, F1: p1, F2: 2147483647}
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 string, p1 int8) int32 {
    if (true && (!true)) {
        v9 := GenS1{F0: 2147483647}
        _ = v9
    }
    var v10 int = 1
    v10 = v10
    v11 := []int{v10}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -827
}

func main() {
    var v1 int8 = -43
    var v2 int32 = 1
    var v3 uint16 = 924
    var v12 int64 = 472
    var v13 int16 = 1000
    var v14 int16 = v13
    v3 += (v3 + 486)
    v15 := []int{42}
    var v16 int32 = genFn0(1, true, 18446744073709551615)
    var v17 int8 = v1
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v2 = v16
    if true {
        var v18 int8 = v1
        _ = v18
    } else {
        if false {
            var v19 int = 179
            v20 := GenS2{F0: v1, F1: true, F2: (v2 / ((v2) | 1))}
            _ = v19
            _ = v20
        } else {
            if (false && false) {
                var v21 uint64 = 1
                v22 := GenS3{F0: false, F1: v21, F2: 1}
                if (false && false) {
                    var v23 uint64 = 42
                    _ = v23
                } else {
                    v1 = (^(v17))
                }
                _ = v21
                _ = v22
            } else {
                v16 = genFn1(("foo bar" + "foo bar"), v17)
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            var v24 int32 = (-(v2))
            _ = v24
        }
        if true {
            v25 := GenS4{F0: int16(v16), F1: 128, F2: true}
            _ = v25
        }
        if ((v12 * (^(v12))) == v12) {
            if true {
                if (v14 != v14) {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        if (0 < genFn0((fi26 % ((int(v16)) | 1)), false, 1)) {
                            var v27 int16 = v13
                            for fi28 := 0; fi28 < 8; fi28++ {
                                if true {
                                    var v29 float64 = 3.14159
                                    _ = v29
                                }
                                v15 = append(v15, (^(fi26)))
                                _ = fi28
                            }
                            _ = v27
                        } else {
                            for fi30 := 0; fi30 < 18; fi30++ {
                                v31 := GenS1{F0: 4294967295}
                                v32 := GenS4{F0: -1, F1: 56, F2: true}
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                        _ = fi26
                    }
                }
                v33 := GenS4{F0: -1, F1: 160, F2: (true || true)}
                v12 = (v12 / ((321) | 1))
                _ = v33
            }
            v2 = v2
        }
    }
    if (1e10 <= (float64(1e10) / 0.1)) {
        v34 := GenS4{F0: (v13 >> ((4294967295) & 63)), F1: 185, F2: true}
        _ = v34
    }
    switch "foo bar" {
    case "a":
        var v35 uint64 = 42
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
