package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int16; F2 string }
type GenS3 struct { F0 uint64; F1 uint32; F2 int32 }
func genFn0(p0 uint32, p1 int8, p2 uint) uint32 {
    p1 = (p1 >> (((^(p2))) & 63))
    v6 := []int{4294967295, -480, 1}
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int16 = 1
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int16 = -1
            switch 32767 {
            case -2147483648:
                _ = len(v6)
            case 42:
                var v11 bool = ((fi7 >> ((p2) & 63)) <= fi9)
                _ = v11
            case 1:
                var v12 uint16 = 42
                _ = v12
            default:
                if (-2.25 >= 0.0) {
                    if false {
                        if (!true) {
                            p1 = (p1 + -85)
                            _ = len(v6)
                        } else {
                            if len(v6) > 0 {
                                v6[0] = fi9
                            }
                            if false {
                                var v13 uint32 = p0
                                v8 += (v10 / (((-(v8))) | 1))
                                v14 := GenS1{F0: false}
                                _ = v13
                                _ = v14
                            } else {
                                v10 += v10
                            }
                        }
                    } else {
                        fmt.Println(v8)
                        v8 = v10
                    }
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v15 := GenS2{F0: 0, F1: (v10 + (v8 * -826)), F2: ""}
                    v16 := GenS2{F0: 375, F1: int16(p2), F2: ("a" + "foo bar")}
                    _ = v15
                    _ = v16
                }
            }
            switch (fi7 % ((fi7) | 1)) {
            case 32768:
                fmt.Println(p0, fi7)
            default:
                v17 := GenS2{F0: 42, F1: v10, F2: ""}
                _ = v17
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (^(p0))
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 1
    var v3 int32 = 208
    var v4 uint16 = 2
    var v5 int32 = v2
    var v18 uint64 = 42
    var v19 int8 = 0
    if true {
        v20 := GenS2{F0: 975, F1: -32768, F2: ""}
        var v21 string = "a"
        var v22 float32 = 3.14159
        _ = v20
        _ = v21
        _ = v22
    } else {
        var v23 float64 = 1.5
        if false {
            for fi24 := 0; fi24 < 6; fi24++ {
                v4 = (^(v4))
                var v25 int64 = -9223372036854775808
                v4 += (v4 ^ v4)
                _ = fi24
                _ = v25
            }
        } else {
            v26 := []int{42, 0, -1}
            v2 += v3
            fmt.Println(v4, v2)
            _ = v26
        }
        var v27 int16 = -1
        _ = v23
        _ = v27
    }
    v28 := GenS3{F0: 1, F1: genFn0(2147483647, (v19 | v19), 0), F2: (v3 >> ((0) & 63))}
    var v29 string = "a"
    var v30 int = 0
    fmt.Println(v5)
    switch v30 {
    case -2147483648:
        v19 -= v19
    case 32768:
        v1["hello"] = v29
    default:
        for fi31 := 0; fi31 < 10; fi31++ {
            for fi32 := 0; fi32 < 9; fi32++ {
                v2 -= (^(v3))
                _ = fi32
            }
            _ = fi31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
