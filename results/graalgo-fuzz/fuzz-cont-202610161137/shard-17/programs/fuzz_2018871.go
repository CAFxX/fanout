package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint8, p1 int64, p2 uint32) int64 {
    var v6 float64 = ((1.5 / -2.25) - (1.5 * 3.14159))
    fmt.Println(p0, p0)
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint8 = p0
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        p2 = 0
        p2 = p2
        p2 -= p2
        _ = fi9
    }
    v10 := []int{42, -737}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return (p1 | -1)
}

func main() {
    var v1 uint32 = 4294967295
    var v2 bool = (!(false || (true && (9223372036854775807 <= 9223372036854775807))))
    var v3 int64 = 1
    var v4 uint8 = 125
    var v5 uint32 = v1
    var v11 bool = (42 == 42)
    v12 := []int{337}
    var v13 uint8 = (^(v4))
    if len(v12) > 0 {
        v12[0] = 2147483647
    }
    var v14 int = 1
    v15 := GenS1{F0: 747}
    var v16 uint16 = 32767
    v16 = (v16 / ((v16) | 1))
    if (!v2) {
        if ((!true) && (65535 == 2)) {
            v17 := GenS2{F0: v14, F1: 9223372036854775807}
            for fi18 := 0; fi18 < 7; fi18++ {
                v19 := GenS3{F0: 9223372036854775807}
                var v20 uint = 351
                _ = fi18
                _ = v19
                _ = v20
            }
            if (false && v11) {
                if v2 {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        if ((v16 ^ v16) == (v16 + v16)) {
                            if false {
                                var v22 bool = false
                                fmt.Println(v5, v5)
                                v3 += (v3 * v3)
                                _ = v22
                            }
                        } else {
                            v2 = true
                            v16 -= (v16 + (v16 % (((v16 | 2)) | 1)))
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        }
                        for fi23 := 0; fi23 < 6; fi23++ {
                            v24 := GenS4{F0: (true || ("hello" >= "foo bar")), F1: 536}
                            var v25 uint = 0
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        var v26 uint64 = 0
                        _ = fi21
                        _ = v26
                    }
                    var v27 uint8 = v13
                    var v28 int16 = -32768
                    _ = v27
                    _ = v28
                } else {
                    fmt.Println(v2, v16, v11)
                }
                v29 := GenS3{F0: (^(v3))}
                var v30 int8 = 42
                _ = v29
                _ = v30
            } else {
                v5 -= (v1 / ((v5) | 1))
                var v31 int64 = (v3 & genFn0(v13, 949, v5))
                for fi32 := 0; fi32 < 4; fi32++ {
                    v5 -= (v5 ^ uint32(v4))
                    _ = fi32
                }
                _ = v31
            }
            _ = v17
        }
        fmt.Println(v16, v3)
        v13 += (^(v4))
    } else {
        v3 -= v3
        var v33 float32 = (0.1 / 1e10)
        _ = v33
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v11 = v2
    for fi34 := 0; fi34 < 12; fi34++ {
        v35 := GenS4{F0: v11, F1: 2}
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
