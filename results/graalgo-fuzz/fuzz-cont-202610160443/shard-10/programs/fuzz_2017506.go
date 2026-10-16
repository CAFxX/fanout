package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int64, p1 int64) uint8 {
    if true {
        v4 := GenS1{F0: (p0 >> ((1) & 63))}
        if true {
            p1 = p0
            var v5 uint = 18446744073709551615
            if true {
                var v6 uint = (^(v5))
                _ = v6
            }
            _ = v5
        } else {
            if true {
                switch 112 {
                case 128:
                    fmt.Println(p0, p0)
                case 2147483647:
                    var v7 bool = (2147483647 > 2147483647)
                    _ = v7
                case 0:
                    fmt.Println(p1, p1)
                default:
                    v8 := []int{32768, -9223372036854775808, -137}
                    _ = v8
                }
            } else {
                var v9 bool = true
                for fi10 := 0; fi10 < 14; fi10++ {
                    v9 = ((p1 / ((-9223372036854775808) | 1)) <= p1)
                    _ = fi10
                }
                p0 -= (p0 | -1)
                _ = v9
            }
        }
        _ = v4
    }
    switch 2147483647 {
    case -474:
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := GenS1{F0: 0}
            if true {
                v13 := GenS2{F0: -111, F1: 32767}
                _ = v13
            } else {
                var v14 uint32 = 2147483647
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    case 0:
        var v15 bool = (p1 == p0)
        _ = v15
    default:
        var v16 float32 = 1.5
        _ = v16
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 int16 = -754
    var v2 uint16 = 577
    v3 := make(map[int]bool)
    v2 = 255
    var v17 float64 = 3.14159
    switch -9223372036854775808 {
    case 32767:
        fmt.Println(v17, v2, v2)
    case -1:
        switch 42 {
        case -2147483648:
            var v18 uint = 4294967295
            _ = v18
        default:
            for fi19 := 0; fi19 < 9; fi19++ {
                if true {
                    if true {
                        for fi20 := 0; fi20 < 7; fi20++ {
                            fmt.Println(v2, v17)
                            _ = len(v3)
                            if ((true && true) || (707 <= 905)) {
                                v21 := GenS1{F0: 0}
                                v17 -= 0.0
                                var v22 uint32 = 0
                                _ = v21
                                _ = v22
                            } else {
                                var v23 int32 = 42
                                fmt.Println(v1, v1)
                                if true {
                                    v24 := GenS3{F0: (v2 << ((382) & 63))}
                                    switch 42 {
                                    case -2147483648:
                                        v1 = (v1 | v1)
                                        v25 := []int{fi19}
                                        _ = v25
                                    case 2147483647:
                                        v26 := GenS2{F0: 100, F1: -334}
                                        var v27 uint16 = (v2 | v2)
                                        _ = v26
                                        _ = v27
                                    case 994:
                                        var v28 int32 = 930
                                        _ = v28
                                    }
                                    _ = v24
                                } else {
                                    fmt.Println(fi19, fi20, v23)
                                    v2 += 65535
                                    v2 -= (v2 % (((v2 & (v2 | (v2 & v2)))) | 1))
                                }
                                _ = v23
                            }
                            _ = fi20
                        }
                        var v29 uint16 = 1
                        v2 = v29
                        _ = v29
                    } else {
                        fmt.Println(v2)
                    }
                    fmt.Println(v17)
                }
                _ = fi19
            }
        }
    case 128:
        v1 += (v1 & 0)
    }
    var v30 int64 = int64(v1)
    v1 = (v1 / (((v1 * 1)) | 1))
    v31 := GenS4{F0: (^(v2)), F1: -710}
    v30 = (v30 - 0)
    v32 := []int{-1, -618}
    v33 := GenS5{F0: 42}
    if true {
        if false {
            var v34 float64 = v17
            _ = v34
        } else {
            fmt.Println(v30, v2, v17)
        }
        v30 = (v30 / (((v30 << ((366) & 63))) | 1))
    } else {
        var v35 string = (("a" + "\n\t") + "foo bar")
        v1 += v1
        v30 = (v30 << ((787) & 63))
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
