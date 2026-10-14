package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int16; F1 int32; F2 int64 }
type GenS3 struct { F0 int32; F1 uint8; F2 int }
func genFn0(p0 uint16) uint64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int8 = 66
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 string = "foo bar"
        if false {
            if true {
                var v11 string = (v10 + v10)
                fmt.Println(v11, v11, v11)
                _ = v11
            } else {
                v10 = v10
                if false {
                    if false {
                        switch 110 {
                        case 805:
                            var v12 uint32 = 65535
                            _ = v12
                        case 110:
                            fmt.Println(p0)
                        }
                        v13 := GenS1{F0: 128, F1: (p0 | p0)}
                        v14 := GenS2{F0: 0, F1: 42, F2: -353}
                        _ = v13
                        _ = v14
                    } else {
                        if false {
                            p0 -= (p0 % ((p0) | 1))
                        }
                        p0 += 704
                    }
                    var v15 bool = (false && false)
                    var v16 int16 = 0
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint32 = 42
                    if false {
                        p0 += 1
                        v18 := []int{32768, (fi9 << ((18446744073709551615) & 63))}
                        _ = v18
                    } else {
                        if (0 <= 1000) {
                            v19 := []int{fi9}
                            _ = v19
                        } else {
                            p0 += (^(p0))
                            for fi20 := 0; fi20 < 12; fi20++ {
                                v21 := []int{32767, fi9, fi9, fi9}
                                var v22 string = "a"
                                v21 = append(v21, (fi20 ^ (fi20 << ((526) & 63))))
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        }
                        v23 := GenS2{F0: -676, F1: 42, F2: -9223372036854775808}
                        var v24 float64 = 3.14159
                        _ = v23
                        _ = v24
                    }
                    _ = v17
                }
            }
            v10 = "hello"
        }
        v25 := []int{(fi9 * (-(fi9))), (fi9 << ((42) & 63))}
        _ = fi9
        _ = v10
        _ = v25
    }
    _ = p0
    return 883
}

func main() {
    var v1 uint16 = 42
    var v2 string = "a"
    v3 := make(map[int]int)
    var v4 uint = 42
    var v5 float64 = (0.0 / (0.0 - (1e10 + (0.0 * 0.1))))
    v6 := make(map[uint32]bool)
    var v26 int = 4294967295
    var v27 int64 = -241
    v1 = 993
    for fi28 := 0; fi28 < 20; fi28++ {
        v2 = ((v2 + v2) + "\n\t")
        var v29 int64 = (v27 << (((^(v4))) & 63))
        _ = fi28
        _ = v29
    }
    v30 := GenS2{F0: 815, F1: -1, F2: v27}
    v31 := GenS2{F0: 1000, F1: 0, F2: v27}
    if (1 > 495) {
        v32 := GenS2{F0: 42, F1: 1, F2: (v27 * (v27 | v27))}
        delete(v6, 2)
        _ = v32
    } else {
        switch (v26 + (-(v26))) {
        case 0:
            v26 = v26
        default:
            v33 := []int{9223372036854775807}
            _ = v33
        }
        if (!true) {
            _ = len(v3)
            v34 := GenS3{F0: -2147483648, F1: 103, F2: v26}
            if (691 == v27) {
                if true {
                    v2 = ((v2 + v2) + v2)
                }
            } else {
                if (!true) {
                    v26 = int(v27)
                    v35 := GenS1{F0: (v26 / ((-660) | 1)), F1: (v1 & (^(v1)))}
                    if true {
                        v26 = (^(v26))
                    }
                    _ = v35
                }
            }
            _ = v34
        } else {
            v2 = v2
            fmt.Println(v2, v27)
            for fi36 := 0; fi36 < 11; fi36++ {
                fmt.Println(v27)
                _ = fi36
            }
        }
        for fi37 := 0; fi37 < 1; fi37++ {
            fmt.Println(v27, v27)
            _ = fi37
        }
    }
    var v38 float64 = float64((float32(v5) / 1.5))
    v39 := GenS3{F0: 2147483647, F1: 255, F2: -2147483648}
    v27 = (v27 & v27)
    if false {
        fmt.Println(v5)
        var v40 uint32 = 171
        _ = v40
    }
    for fi41 := 0; fi41 < 3; fi41++ {
        var v42 string = "a"
        _ = fi41
        _ = v42
    }
    v26 = v26
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
