package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint32) uint64 {
    v7 := []int{9223372036854775807}
    p0 -= p0
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        p0 = (p0 + (p0 + (p0 - (p0 * p0))))
        _ = fi8
    }
    _ = p0
    _ = v7
    return 882
}
func genFn1(p0 uint, p1 uint64) int16 {
    p1 -= 18446744073709551615
    v9 := []int{}
    v10 := GenS1{F0: 42, F1: p1}
    v11 := GenS1{F0: 1, F1: p1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 32767
}

func main() {
    var v1 int8 = 0
    var v2 int64 = 601
    var v3 string = "foo bar"
    v4 := make(map[int]string)
    var v5 float32 = 0.1
    v6 := make(map[uint32]int)
    _ = len(v6)
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS2{F0: (v2 >> ((65535) & 63)), F1: 2147483648}
        _ = fi12
        _ = v13
    }
    var v14 int8 = v1
    fmt.Println(v1, v1)
    var v15 bool = (!(false || false))
    v2 = 9223372036854775807
    var v16 string = v3
    v17 := GenS1{F0: 1000, F1: 2}
    v14 -= v1
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 float64 = (float64(float32((float64((0.1 - (-2.25 / 1.0))) - 0.1))) / 0.1)
        fmt.Println(v14, v19)
        var v20 float64 = v19
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        if (!(127 <= 0)) {
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 uint8 = 197
                v24 := GenS2{F0: v2, F1: 2}
                v25 := GenS2{F0: v2, F1: 4294967295}
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
        }
        if (!(!v15)) {
            v26 := GenS2{F0: (v2 | (v2 % ((v2) | 1))), F1: 703}
            _ = v26
        } else {
            if (!((846 == 255) && v15)) {
                v16 = (("" + v3) + v3)
                switch "hello" {
                case "foo bar":
                    v27 := GenS1{F0: 0, F1: 4294967295}
                    _ = v27
                }
            } else {
                v1 = (v14 | v14)
                var v28 int = (fi21 + 0)
                _ = v28
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                v1 = v14
                if true {
                    if v15 {
                        v30 := []int{}
                        if false {
                            v16 = v16
                            var v31 float64 = (3.14159 / ((1.5 - 1.5) + float64(v5)))
                            _ = v31
                        } else {
                            v32 := GenS3{F0: -1}
                            delete(v4, fi21)
                            if true {
                                var v33 uint32 = 626
                                _ = v33
                            } else {
                                if len(v30) > 0 {
                                    v30[0] = fi29
                                }
                                v34 := GenS4{F0: -32768, F1: 102}
                                _ = v34
                            }
                            fmt.Println(v2, v1, v2)
                            _ = v32
                        }
                        _ = v30
                    }
                    v16 = v16
                }
                _ = fi29
            }
        }
        _ = fi21
    }
    fmt.Println(v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
