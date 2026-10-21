package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int8; F1 int16; F2 bool }
type GenS4 struct { F0 uint; F1 int16 }
func genFn0(p0 int8, p1 int8) uint16 {
    var v4 float64 = 0.0
    p0 = (p0 * (p0 / (((-(p1))) | 1)))
    if false {
        p1 = p0
    } else {
        var v5 uint = 805
        _ = v5
    }
    p0 = (p1 | (p1 - p1))
    if false {
        fmt.Println(p1, p0, p1)
        for fi6 := 0; fi6 < 10; fi6++ {
            fmt.Println(p0)
            v7 := GenS1{F0: 2, F1: 1000, F2: -9223372036854775808}
            _ = fi6
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 2
}

func main() {
    v1 := []int{1}
    var v2 string = "hello"
    var v3 uint32 = 983
    if true {
        var v8 float32 = 0.1
        fmt.Println(v3, v3, v2)
        _ = v8
    }
    v3 = v3
    v2 = (v2 + "hello")
    v9 := GenS2{F0: 779, F1: 2147483647}
    var v10 uint64 = 2
    v3 = 4294967295
    _ = len(v1)
    var v11 float64 = (((0.1 - 0.0) / float64(-2.25)) / (1e10 + ((0.0 * 3.14159) * 1e10)))
    v10 -= (v10 * (^(v10)))
    if true {
        v3 -= (v3 & v3)
        for fi12 := 0; fi12 < 9; fi12++ {
            fmt.Println(fi12)
            v2 = ""
            _ = fi12
        }
        if true {
            v13 := GenS3{F0: 25, F1: 1, F2: (128 <= 1)}
            switch v2 {
            case "foo bar":
                var v14 int64 = -9223372036854775808
                _ = v14
            case "hello":
                if true {
                    if (!(!false)) {
                        if (((65535 != 42) || false) || (v11 < v11)) {
                            var v15 float64 = float64(0.1)
                            var v16 uint64 = v10
                            _ = v15
                            _ = v16
                        } else {
                            var v17 bool = true
                            for fi18 := 0; fi18 < 18; fi18++ {
                                for fi19 := 0; fi19 < 8; fi19++ {
                                    v17 = (v17 && false)
                                    v17 = v17
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi19
                                }
                                _ = fi18
                            }
                            _ = v17
                        }
                        var v20 uint16 = genFn0(42, 1)
                        v21 := GenS3{F0: 42, F1: -1, F2: true}
                        _ = v20
                        _ = v21
                    }
                    var v22 bool = false
                    var v23 int = -1
                    _ = v22
                    _ = v23
                }
            case "":
                fmt.Println(v2, v10)
            default:
                v10 = (v10 | 423)
            }
            _ = v13
        } else {
            if (128 > 0) {
                if len(v1) > 0 {
                    v1[0] = 32767
                }
                var v24 float32 = float32((v11 - (v11 / 0.1)))
                v25 := GenS2{F0: 4294967295, F1: 32767}
                _ = v24
                _ = v25
            }
            var v26 float64 = v11
            _ = v26
        }
    }
    if true {
        var v27 int32 = 2147483647
        v11 = v11
        v10 = 2
        _ = v27
    } else {
        for fi28 := 0; fi28 < 5; fi28++ {
            v29 := GenS1{F0: 804, F1: 508, F2: 42}
            _ = fi28
            _ = v29
        }
        var v30 int64 = -472
        _ = v30
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        if ((v10 != 1) && true) {
            v1 = append(v1, (fi31 % ((632) | 1)))
            var v32 int = fi31
            switch v32 {
            case 2147483647:
                v2 = ("a" + v2)
            default:
                v3 += (v3 | v3)
            }
            _ = v32
        }
        if (false && (int16(v3) <= -528)) {
            var v33 uint16 = 997
            v34 := GenS4{F0: uint(v10), F1: 439}
            var v35 uint16 = v33
            _ = v33
            _ = v34
            _ = v35
        }
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
