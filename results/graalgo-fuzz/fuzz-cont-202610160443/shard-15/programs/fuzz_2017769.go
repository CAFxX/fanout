package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 string; F2 int8 }
func genFn0(p0 uint16, p1 int64, p2 bool) bool {
    switch "" {
    case "\n\t":
        p1 = -292
    case "":
        var v6 int16 = -864
        _ = v6
    case "a":
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(p2, p0)
            v8 := GenS1{F0: p1, F1: true}
            if false {
                var v9 int16 = 1
                p2 = p2
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS2{F0: 0}
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS1{F0: p1, F1: ((!(!p2)) || (((1e10 * 3.14159) * 1e10) > (-2.25 + 0.1)))}
                p1 += p1
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    }
    v14 := []int{42, 127}
    p0 = (p0 / (((p0 | (^(p0)))) | 1))
    p0 -= (p0 ^ p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return ((float32(0.0) - -2.25) >= ((1.5 * 1.5) / float32(float64(3.14159))))
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]bool)
    var v3 uint = 485
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    fmt.Println(v1, v1)
    v3 = (v3 + (v3 * v3))
    v5[1] = "a"
    var v15 int8 = 1
    v16 := []int{}
    var v17 float32 = 1e10
    var v18 string = (("foo bar" + "hello") + ("\n\t" + "foo bar"))
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 uint = v3
        fmt.Println(v20, v18, v1)
        _ = fi19
        _ = v20
    }
    v3 = (v3 | v3)
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            var v23 uint64 = 2
            var v24 int = fi21
            if (false && true) {
                v24 = fi22
                if (!true) {
                    var v25 int16 = 0
                    if ((v3 * v3) <= 2) {
                        v26 := GenS1{F0: -9223372036854775808, F1: (false || false)}
                        v17 += (0.1 + (v17 + v17))
                        var v27 uint16 = 0
                        fmt.Println(v23, v18, v3)
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS3{F0: false, F1: v18, F2: (v15 + v15)}
                    _ = v25
                    _ = v28
                } else {
                    var v29 float64 = v1
                    _ = v29
                }
                v30 := v2[323]
                v24 -= v24
                _ = v30
            } else {
                switch v18 {
                case "":
                    v31 := GenS2{F0: (^(v3))}
                    _ = v31
                case "foo bar":
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                default:
                    fmt.Println(fi21, v24)
                }
                if true {
                    v15 -= v15
                    if (false || ((18446744073709551615 >= (^(v23))) && true)) {
                        var v32 uint8 = 227
                        _ = v32
                    } else {
                        v3 = (v3 | v3)
                        if (!(v18 >= "a")) {
                            var v33 int32 = 2147483647
                            _ = v33
                        }
                        fmt.Println(fi21, fi21)
                    }
                    _ = len(v16)
                    var v34 int8 = (v15 ^ v15)
                    _ = v34
                } else {
                    v23 += v23
                    var v35 uint32 = 1
                    _ = v35
                }
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        fmt.Println(fi21, v18)
        v36 := GenS3{F0: (true && ((((!false) || (973 != fi21)) || (true && false)) || (0 <= 1000))), F1: v18, F2: 0}
        _ = fi21
        _ = v36
    }
    v17 += v17
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
