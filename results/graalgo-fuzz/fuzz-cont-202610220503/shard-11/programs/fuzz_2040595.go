package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint; F2 int8 }
func genFn0(p0 bool, p1 int8, p2 int8) uint {
    v4 := GenS1{F0: -2147483648}
    v5 := []int{-9223372036854775808, -1}
    fmt.Println(p1)
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        p0 = (42 < 909)
        if p0 {
            p1 = p1
            if p0 {
                if p0 {
                    var v7 uint16 = 992
                    _ = v7
                } else {
                    var v8 uint8 = 42
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v8
                }
                v9 := GenS1{F0: 1}
                var v10 uint64 = 42
                _ = v10
                _ = v9
            }
        } else {
            fmt.Println(p0)
            v11 := GenS2{F0: p2, F1: 552, F2: 100}
            if false {
                switch fi6 {
                case -76:
                    fmt.Println(p1, p0)
                }
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = v11
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 816
}
func genFn1(p0 uint64, p1 uint8) uint32 {
    v12 := []int{179, -2147483648, 1}
    fmt.Println(p1, p0)
    if false {
        var v13 int16 = 462
        v12 = append(v12, 519)
        v13 -= (v13 << ((416) & 63))
        _ = v13
    }
    if (-2.25 > float32(1e10)) {
        v14 := GenS1{F0: 474}
        _ = v14
    } else {
        if false {
            if false {
                fmt.Println(p0, p1)
                if true {
                    v15 := GenS1{F0: 1}
                    _ = v15
                }
            }
            if true {
                _ = len(v12)
            }
            switch "\n\t" {
            case "":
                fmt.Println(p0, p1)
            case "hello":
                if false {
                    v16 := GenS2{F0: 1, F1: 65535, F2: 1}
                    _ = v16
                }
            }
        } else {
            if false {
                var v17 int8 = 100
                var v18 uint32 = 894
                p0 -= 42
                _ = v17
                _ = v18
            }
            if (-128 > -128) {
                for fi19 := 0; fi19 < 9; fi19++ {
                    if false {
                        for fi20 := 0; fi20 < 16; fi20++ {
                            var v21 float32 = 0.0
                            var v22 uint32 = 2147483648
                            p1 = p1
                            _ = len(v12)
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                    if len(v12) > 0 {
                        v12[0] = (fi19 & (fi19 | fi19))
                    }
                    if (1 != 42) {
                        v23 := GenS2{F0: 1, F1: 255, F2: int8(p1)}
                        _ = v23
                    } else {
                        var v24 float64 = (0.1 / -2.25)
                        v25 := GenS1{F0: 623}
                        _ = v24
                        _ = v25
                    }
                    _ = fi19
                }
                if (((p0 ^ p0) <= p0) && false) {
                    p1 = (p1 - (p1 >> ((42) & 63)))
                    if true {
                        var v26 int32 = 1
                        p0 += 0
                        _ = v26
                    }
                }
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
        }
    }
    if (p1 > p1) {
        switch 1 {
        case 866:
            fmt.Println(p0, p0)
        case 9223372036854775807:
            var v27 uint32 = 4294967295
            _ = v27
        }
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = v12
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    var v28 int8 = -72
    fmt.Println(v28)
    v29 := GenS2{F0: (v28 << ((1) & 63)), F1: genFn0(true, (v28 % ((v28) | 1)), v28), F2: (v28 * (v28 % (((v28 & (v28 - v28))) | 1)))}
    var v30 uint64 = 0
    var v31 bool = (!(!(0.1 != (float32(((0.0 / ((-2.25 / 1.0) / (3.14159 / 1.5))) - 0.1)) * 1e10))))
    for fi32 := 0; fi32 < 4; fi32++ {
        var v33 int64 = -520
        for fi34 := 0; fi34 < 9; fi34++ {
            if ((fi34 ^ (^(fi32))) <= fi32) {
                v35 := []int{(fi34 << ((0) & 63)), fi32}
                if len(v35) >= 2 {
                    v35 = v35[1:len(v35)]
                }
                _ = v35
            } else {
                var v36 int64 = v33
                _ = v36
            }
            _ = fi34
        }
        var v37 uint32 = 148
        _ = fi32
        _ = v33
        _ = v37
    }
    var v38 int8 = v28
    var v39 uint = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
}
