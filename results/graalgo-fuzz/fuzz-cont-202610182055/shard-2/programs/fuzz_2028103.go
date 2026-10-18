package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 bool }
func genFn0(p0 bool, p1 uint64, p2 int) int {
    p2 = p2
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 32767
}
func genFn1(p0 int16, p1 int32, p2 uint32) int16 {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 string = ""
        if (p2 <= p2) {
            for fi8 := 0; fi8 < 3; fi8++ {
                if (49 == 109) {
                    v7 = v7
                    v7 = v7
                    p2 += (p2 + 1)
                    v9 := []int{genFn0(((0.1 - 0.0) == 1.5), 713, (fi6 * (fi8 + (fi6 * (fi6 - genFn0(false, 0, genFn0(true, 4294967295, 32768))))))), (-(fi8))}
                    _ = v9
                } else {
                    var v10 int = (fi8 - (-(fi6)))
                    _ = v10
                }
                if true {
                    v11 := GenS1{F0: 549, F1: 362, F2: (1.5 != -2.25)}
                    _ = v11
                } else {
                    v12 := GenS1{F0: 2, F1: 2, F2: false}
                    var v13 int8 = 100
                    _ = v12
                    _ = v13
                }
                var v14 int32 = p1
                _ = fi8
                _ = v14
            }
            var v15 uint64 = 817
            _ = v15
        } else {
            p0 = (p0 >> ((42) & 63))
            p1 += p1
            for fi16 := 0; fi16 < 15; fi16++ {
                if (!(2 > 255)) {
                    var v17 uint64 = 42
                    switch fi16 {
                    case 0:
                        var v18 bool = ((p2 % (((p2 + (^(p2)))) | 1)) <= 2)
                        _ = v18
                    case -414:
                        fmt.Println(p1)
                    case -899:
                        if ((255 >= 954) && (1 == (-(p1)))) {
                            switch v7 {
                            case "a":
                                switch "" {
                                case "hello":
                                    var v19 uint64 = 1
                                    _ = v19
                                case "a":
                                    var v20 int = (fi6 & (-(fi16)))
                                    _ = v20
                                }
                                v21 := []int{42}
                                var v22 uint16 = 0
                                fmt.Println(p1)
                                _ = v21
                                _ = v22
                            default:
                                var v23 uint64 = v17
                                _ = v23
                            }
                        } else {
                            p2 += (p2 / ((p2) | 1))
                        }
                    }
                    _ = v17
                }
                var v24 int = (fi6 & -9223372036854775808)
                _ = fi16
                _ = v24
            }
        }
        v25 := []int{genFn0((true && (1 > 942)), 0, 9223372036854775807), (-(fi6)), genFn0(true, 42, fi6)}
        _ = fi6
        _ = v25
        _ = v7
    }
    v26 := GenS1{F0: 171, F1: 255, F2: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return p0
}

func main() {
    v1 := []string{"\n\t", "\n\t"}
    v2 := make(map[int]int)
    var v3 float32 = -2.25
    var v4 uint = 65535
    var v5 uint16 = 1
    if true {
        fmt.Println(v5)
        _ = len(v1)
        switch 155 {
        case -154:
            v2[32767] = 408
        case 32767:
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        case 223:
            var v27 float64 = 1.5
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        if ((!false) && false) {
            v29 := GenS1{F0: (^(v5)), F1: (v4 | 0), F2: true}
            for fi30 := 0; fi30 < 14; fi30++ {
                _ = len(v1)
                _ = fi30
            }
            v4 = v4
            _ = v29
        }
        _ = fi28
    }
    var v31 uint32 = 65535
    v3 = v3
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 bool = false
        for fi34 := 0; fi34 < 20; fi34++ {
            fmt.Println(fi34)
            if (v33 && ((float32(0.1) / v3) < (v3 - float32(0.1)))) {
                v35 := GenS1{F0: v5, F1: 203, F2: v33}
                _ = v35
            }
            _ = fi34
        }
        _ = fi32
        _ = v33
    }
    var v36 int32 = 0
    v37 := GenS1{F0: 255, F1: 319, F2: true}
    v3 += float32(0.1)
    var v38 uint8 = 253
    for fi39 := 0; fi39 < 4; fi39++ {
        var v40 bool = (!(18446744073709551615 <= (v4 << (((v4 * 255)) & 63))))
        switch (fi39 - 32768) {
        case 0:
            var v41 uint64 = 757
            _ = v41
        default:
            v38 = 0
        }
        _ = fi39
        _ = v40
    }
    var v42 uint16 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
}
