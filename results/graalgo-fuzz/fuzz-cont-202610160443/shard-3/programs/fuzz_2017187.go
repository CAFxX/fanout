package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0(p0 int16, p1 string, p2 int64) int32 {
    if (true || (!(42 == 526))) {
        fmt.Println(p2)
        for fi7 := 0; fi7 < 20; fi7++ {
            if ((0 == 18446744073709551615) && true) {
                p1 = p1
                if false {
                    var v8 uint = 4294967295
                    if false {
                        var v9 int64 = p2
                        _ = v9
                    } else {
                        v10 := []int{(fi7 >> ((v8) & 63)), (fi7 % ((fi7) | 1)), 42}
                        _ = v10
                    }
                    _ = v8
                } else {
                    for fi11 := 0; fi11 < 15; fi11++ {
                        var v12 bool = (float64(0.1) <= 1e10)
                        _ = fi11
                        _ = v12
                    }
                    var v13 int32 = -291
                    v14 := []int{(fi7 * fi7), fi7, fi7}
                    _ = v13
                    _ = v14
                }
                var v15 bool = true
                _ = v15
            }
            p1 = ""
            var v16 int64 = p2
            _ = fi7
            _ = v16
        }
    } else {
        var v17 bool = true
        p0 = (p0 >> ((42) & 63))
        _ = v17
    }
    p0 += p0
    p1 = p1
    for fi18 := 0; fi18 < 10; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            if true {
                var v20 uint16 = 42
                p2 += p2
                v21 := []int{(fi18 | -707), (-(fi18)), 0}
                _ = v20
                _ = v21
            } else {
                if (0 >= 223) {
                    switch (fi18 << ((508) & 63)) {
                    case -2147483648:
                        v22 := GenS1{F0: p1, F1: 42, F2: -736}
                        _ = v22
                    }
                } else {
                    p0 += (^(p0))
                }
            }
            v23 := GenS1{F0: p1, F1: 62, F2: 2147483647}
            _ = fi19
            _ = v23
        }
        if true {
            v24 := GenS2{F0: 42, F1: (p0 + 1000)}
            var v25 uint16 = 429
            fmt.Println(v25)
            _ = v24
            _ = v25
        }
        p0 += p0
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int32, p1 string, p2 int32) uint8 {
    v26 := GenS1{F0: p1, F1: 128, F2: p0}
    v27 := GenS3{F0: 2147483648, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return 198
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 424
    v3 := make(map[int]int)
    var v4 int32 = 0
    v5 := []uint32{803}
    var v6 bool = true
    if (1.5 <= (((0.1 + 0.1) * 0.0) * float64(float32((0.1 * 1e10))))) {
        v2 = v2
        v2 -= 1
        for fi28 := 0; fi28 < 8; fi28++ {
            v29 := GenS3{F0: 1, F1: 0}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi28
            _ = v29
        }
    } else {
        v5 = append(v5, 42)
    }
    _ = len(v3)
    v6 = (985 != 380)
    v2 += (^(v2))
    fmt.Println(v4, v6, v4)
    switch 32768 {
    case 869:
        var v30 uint = 208
        _ = v30
    case 128:
        for fi31 := 0; fi31 < 6; fi31++ {
            if v6 {
                v5 = append(v5, 209)
                v2 -= 18446744073709551615
            } else {
                switch ("" + "a") {
                case "hello":
                    fmt.Println(fi31, fi31)
                case "\n\t":
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 int = 1
                        if (v4 < v4) {
                            v4 = (v4 << ((2) & 63))
                            switch v33 {
                            case 168:
                                var v34 float32 = -2.25
                                _ = v34
                            case 128:
                                var v35 int16 = 0
                                _ = v35
                            case -261:
                                var v36 float64 = 0.1
                                _ = v36
                            default:
                                fmt.Println(v6, fi32)
                            }
                        }
                        delete(v3, fi32)
                        _ = fi32
                        _ = v33
                    }
                }
            }
            _ = fi31
        }
    case 42:
        delete(v3, 128)
    default:
        for fi37 := 0; fi37 < 6; fi37++ {
            if true {
                fmt.Println(v2, v4)
                var v38 int64 = -759
                v38 += (v38 % ((1) | 1))
                _ = v38
            }
            var v39 uint8 = 1
            fmt.Println(v2)
            _ = fi37
            _ = v39
        }
    }
    v4 += v4
    v2 = v2
    var v40 bool = (1 != 948)
    if false {
        var v41 uint32 = 1
        v40 = (0 >= -128)
        _ = v41
    } else {
        v6 = true
        delete(v1, 2)
        var v42 int = 128
        _ = v42
    }
    fmt.Println(v4, v40)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
