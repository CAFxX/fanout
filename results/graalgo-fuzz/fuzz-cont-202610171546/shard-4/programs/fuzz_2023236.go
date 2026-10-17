package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint32) int {
    var v7 uint64 = 42
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    return -2147483648
}

func main() {
    var v1 int32 = -135
    var v2 uint64 = 1
    v3 := []int{4294967295}
    var v4 string = "foo bar"
    var v5 float64 = 3.14159
    var v6 int = 0
    for fi8 := 0; fi8 < 18; fi8++ {
        v6 = 2147483647
        v3 = append(v3, (fi8 % ((genFn0(42, uint32(v2))) | 1)))
        if len(v3) > 0 {
            v3[0] = (fi8 | fi8)
        }
        _ = fi8
    }
    var v9 int8 = 12
    fmt.Println(v1, v1, v1)
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(v6)
        _ = fi10
    }
    var v11 int8 = (v9 + -128)
    v4 = v4
    for fi12 := 0; fi12 < 14; fi12++ {
        if true {
            var v13 int = genFn0(124, 1)
            var v14 string = v4
            if true {
                v14 = "hello"
                if (false && ((126 != 127) || true)) {
                    v15 := GenS1{F0: (v4 + v14), F1: -1}
                    _ = v15
                } else {
                    var v16 int32 = v1
                    var v17 uint16 = 32767
                    v18 := GenS2{F0: 128, F1: v17}
                    _ = v16
                    _ = v17
                    _ = v18
                }
                if len(v3) > 0 {
                    v3[0] = genFn0(127, 2147483647)
                }
            } else {
                v2 -= (^(v2))
                v19 := GenS3{F0: 463, F1: 0}
                _ = v19
            }
            _ = v13
            _ = v14
        } else {
            v4 = "foo bar"
        }
        v9 -= v9
        var v20 float32 = 0.0
        _ = fi12
        _ = v20
    }
    var v21 uint64 = 0
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 bool = true
            switch (v4 + "a") {
            case "a":
                switch genFn0(128, 336) {
                case 128:
                    fmt.Println(v11)
                }
            case "hello":
                v25 := GenS1{F0: (((v4 + v4) + v4) + v4), F1: -1}
                _ = v25
            case "foo bar":
                v9 -= 127
                v5 += v5
            default:
                if (!v24) {
                    v11 += 1
                    fmt.Println(v9, v9, v11)
                    _ = len(v3)
                    var v26 uint16 = 32767
                    if v24 {
                        var v27 string = v4
                        _ = v27
                    }
                    _ = v26
                }
            }
            var v28 string = v4
            _ = fi23
            _ = v24
            _ = v28
        }
        _ = fi22
    }
    v29 := GenS4{F0: (^(v1))}
    fmt.Println(v6)
    switch v6 {
    case 127:
        v5 += ((-2.25 + ((v5 * v5) - v5)) * 3.14159)
    default:
        var v30 string = "a"
        _ = v30
    }
    _ = len(v3)
    fmt.Println(v6)
    var v31 int = genFn0(163, 973)
    if false {
        if (!true) {
            for fi32 := 0; fi32 < 6; fi32++ {
                var v33 int32 = (-(v1))
                _ = fi32
                _ = v33
            }
            switch v4 {
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = (v6 << ((65535) & 63))
                }
            case "a":
                v4 = "\n\t"
            case "":
                v34 := GenS2{F0: 217, F1: 255}
                _ = v34
            default:
                v11 -= v11
            }
            var v35 string = v4
            _ = v35
        }
        if ((2 != 106) || (((true || (v1 == v1)) || ((32767 > 65535) && false)) && ((v21 * (v2 + v2)) != v2))) {
            var v36 float64 = v5
            v11 -= int8(v2)
            _ = v36
        }
        var v37 int8 = 0
        _ = v37
    } else {
        for fi38 := 0; fi38 < 3; fi38++ {
            var v39 int64 = -1
            _ = fi38
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
