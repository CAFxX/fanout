package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0() uint {
    var v6 uint64 = 935
    v6 += 18446744073709551615
    fmt.Println(v6)
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            var v8 string = ""
            _ = v8
        }
        v9 := []int{(fi7 << ((65535) & 63)), fi7, (fi7 & 9223372036854775807)}
        _ = fi7
        _ = v9
    }
    _ = v6
    return 842
}
func genFn1(p0 uint, p1 bool) int64 {
    p0 = genFn0()
    p0 = (p0 % ((genFn0()) | 1))
    p1 = (0.0 >= 1.5)
    v10 := []int{32768}
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    var v1 float32 = ((0.0 + (1.5 / 1e10)) - (-2.25 * 0.0))
    var v2 int8 = -128
    var v3 uint = 402
    v4 := make(map[int]bool)
    var v5 int = -161
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 uint = (v3 / ((255) | 1))
        var v13 string = "hello"
        _ = fi11
        _ = v12
        _ = v13
    }
    v3 -= genFn0()
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            v3 = v3
            switch (fi15 ^ 113) {
            case -9223372036854775808:
                var v16 uint16 = 42
                _ = v16
            }
            v5 = fi15
            _ = fi15
        }
        _ = fi14
    }
    v17 := []int{v5, v5, (v5 * (v5 * (v5 | (v5 & 270))))}
    v2 -= v2
    v18 := GenS1{F0: true, F1: v2, F2: 42}
    if true {
        v3 = (v3 >> ((genFn0()) & 63))
    } else {
        if true {
            if (!(-123 <= int8(v5))) {
                if true {
                    v3 = genFn0()
                    var v19 uint32 = 4294967295
                    _ = v19
                } else {
                    var v20 int = (v5 % (((v5 | -429)) | 1))
                    _ = v20
                }
                fmt.Println(v5)
            } else {
                v21 := GenS1{F0: (57 > v2), F1: v2, F2: 0}
                _ = v21
            }
            switch "a" {
            case "foo bar":
                if true {
                    v22 := GenS1{F0: true, F1: v2, F2: 1}
                    _ = v22
                } else {
                    v5 = 29
                    var v23 bool = (228 >= genFn1(v3, false))
                    for fi24 := 0; fi24 < 4; fi24++ {
                        v3 -= (v3 << (((v3 << ((1) & 63))) & 63))
                        _ = fi24
                    }
                    _ = v23
                }
            case "":
                for fi25 := 0; fi25 < 4; fi25++ {
                    v2 = 1
                    var v26 int32 = 0
                    var v27 bool = true
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            default:
                switch "foo bar" {
                case "foo bar":
                    var v28 int8 = (v2 | (v2 & (v2 | v2)))
                    _ = v28
                case "":
                    var v29 uint16 = 0
                    _ = v29
                default:
                    v30 := GenS1{F0: (false || (("" == (("" + "hello") + "")) || ((!(!((v3 + v3) < v3))) || (42 != 46)))), F1: 127, F2: 93}
                    _ = v30
                }
            }
            var v31 uint8 = 1
            _ = v31
        } else {
            if false {
                v32 := GenS1{F0: false, F1: 0, F2: 0}
                _ = v32
            } else {
                v33 := GenS1{F0: (v5 > v5), F1: v2, F2: 128}
                fmt.Println(v2, v1, v2)
                _ = v33
            }
        }
        v2 += -100
        v34 := GenS1{F0: true, F1: (v2 * (v2 & 42)), F2: 255}
        _ = v34
    }
    v35 := GenS1{F0: (2147483647 < 2147483647), F1: -13, F2: 99}
    v36 := v4[v5]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
