package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int16; F2 string }
func genFn0(p0 int64, p1 int8, p2 int) uint16 {
    p1 -= p1
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := []int{}
        for fi6 := 0; fi6 < 8; fi6++ {
            switch (-(fi4)) {
            case 42:
                p0 = (p0 | (p0 & 42))
            case 0:
                if len(v5) > 0 {
                    v5[0] = fi4
                }
            default:
                v5 = append(v5, (fi6 - p2))
            }
            fmt.Println(p0, fi6, fi6)
            _ = fi6
        }
        _ = fi4
        _ = v5
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    return 65535
}
func genFn1(p0 int16, p1 bool) string {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int32 = -1
        _ = fi7
        _ = v8
    }
    switch 4294967295 {
    case 0:
        if p1 {
            v9 := []int{-1, 1, 127, -9223372036854775808}
            for fi10 := 0; fi10 < 15; fi10++ {
                if (("\n\t" + (("" + ("a" + "a")) + "")) == "foo bar") {
                    var v11 uint64 = 2
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = v11
                } else {
                    p1 = false
                    if len(v9) > 0 {
                        v9[0] = 4294967295
                    }
                    if (735 == 0) {
                        p1 = (128 > 127)
                        v12 := GenS1{F0: 2, F1: 255}
                        var v13 bool = true
                        _ = v12
                        _ = v13
                    } else {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        _ = len(v9)
                        p0 = p0
                    }
                }
                _ = fi10
            }
            fmt.Println(p1, p1)
            _ = v9
        }
    }
    var v14 int64 = -685
    _ = p0
    _ = p1
    _ = v14
    return ""
}

func main() {
    var v1 int8 = 0
    v2 := []int{-9223372036854775808, -975}
    var v3 int8 = (v1 * (v1 + v1))
    _ = len(v2)
    fmt.Println(v1, v1)
    for fi15 := 0; fi15 < 7; fi15++ {
        fmt.Println(v1, fi15)
        _ = fi15
    }
    v1 -= -1
    var v16 uint16 = 0
    fmt.Println(v1, v3)
    var v17 uint64 = 446
    v18 := GenS1{F0: v17, F1: 127}
    fmt.Println(v17, v1)
    if (370 < v17) {
        v3 = -31
        v1 = v1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        if true {
            var v19 uint16 = (v16 / ((v16) | 1))
            _ = v19
        } else {
            if (4294967295 > v17) {
                switch -36 {
                case 9223372036854775807:
                    fmt.Println(v3)
                default:
                    v20 := GenS2{F0: 0}
                    _ = v20
                }
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 int64 = 1
                    v17 = v17
                    switch int(v17) {
                    case 0:
                        v22 = -1
                    case -2147483648:
                        fmt.Println(v16, v17, v17)
                    case 32767:
                        v23 := GenS3{F0: 2, F1: -429, F2: ""}
                        _ = v23
                    }
                    _ = fi21
                    _ = v22
                }
            } else {
                if len(v2) > 0 {
                    v2[0] = 42
                }
            }
        }
        v24 := GenS1{F0: uint64(v16), F1: 42}
        if len(v2) > 0 {
            v2[0] = 261
        }
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
