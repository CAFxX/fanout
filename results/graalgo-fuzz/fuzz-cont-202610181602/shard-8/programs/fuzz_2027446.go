package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int16; F2 uint8 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint32, p1 int32, p2 int32) uint {
    var v6 uint8 = 205
    var v7 int64 = -854
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint32, p1 int) uint {
    var v8 uint8 = 42
    var v9 bool = (58 == 42)
    p0 = 820
    if true {
        if (0.0 > float32((1.5 * (1.5 - 0.1)))) {
            v10 := []int{}
            _ = v10
        } else {
            v11 := GenS1{F0: "a", F1: 1, F2: -2147483648}
            v9 = v9
            v12 := []int{-1, p1, (p1 / (((p1 & (p1 % (((p1 ^ p1)) | 1)))) | 1)), p1}
            _ = v11
            _ = v12
        }
        p1 += (-(p1))
    } else {
        if (v9 && (!(0 >= 4294967295))) {
            if v9 {
                v8 = (^(v8))
                var v13 int = p1
                _ = v13
            }
            p0 -= (^(p0))
        }
        var v14 int16 = 32767
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint8 = 9
            switch ("" + "hello") {
            case "a":
                if true {
                    if v9 {
                        var v17 uint64 = 0
                        for fi18 := 0; fi18 < 2; fi18++ {
                            if v9 {
                                v19 := []int{fi15, fi15, fi15}
                                _ = v19
                            }
                            var v20 int32 = -2147483648
                            if true {
                                var v21 uint32 = 2147483648
                                switch p1 {
                                case -835:
                                    var v22 int32 = v20
                                    v14 += v14
                                    var v23 uint = uint(p0)
                                    _ = v22
                                    _ = v23
                                case -577:
                                    var v24 uint8 = 128
                                    _ = v24
                                default:
                                    var v25 uint = 263
                                    v16 = 42
                                    if (v20 <= (-(v20))) {
                                        v26 := []int{}
                                        v26 = append(v26, -275)
                                        var v27 int8 = 56
                                        _ = v26
                                        _ = v27
                                    } else {
                                        v28 := []int{(p1 << (((v25 * (^(v25)))) & 63)), (fi15 / ((fi18) | 1))}
                                        _ = v28
                                    }
                                    _ = v25
                                }
                                switch fi15 {
                                case -1:
                                    v16 += (v16 * v16)
                                default:
                                    var v29 int16 = v14
                                    _ = v29
                                }
                                p0 -= 4294967295
                                _ = v21
                            } else {
                                var v30 uint8 = 1
                                v20 = v20
                                _ = v30
                            }
                            _ = fi18
                            _ = v20
                        }
                        fmt.Println(p0, fi15, v17)
                        _ = v17
                    } else {
                        if v9 {
                            if (p0 >= 2147483647) {
                                v8 = (^(v8))
                            }
                            v16 = (^(v16))
                            var v31 int32 = 246
                            _ = v31
                        } else {
                            fmt.Println(v16, fi15, p1)
                            v32 := GenS1{F0: "\n\t", F1: 50, F2: 1}
                            _ = v32
                        }
                        v9 = (255 <= 32767)
                        v33 := []int{42}
                        _ = v33
                    }
                } else {
                    v34 := GenS1{F0: "foo bar", F1: genFn0(p0, 147, 1), F2: -296}
                    var v35 uint32 = (p0 | p0)
                    _ = v34
                    _ = v35
                }
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 int8 = -128
    var v2 int8 = (v1 / ((0) | 1))
    var v3 int32 = -1
    v4 := make(map[int]int)
    var v5 string = "\n\t"
    var v36 uint64 = 1
    fmt.Println(v5, v1)
    fmt.Println(v5, v3, v3)
    if true {
        if true {
            v37 := GenS1{F0: ((v5 + v5) + (v5 + v5)), F1: 42, F2: (-(v3))}
            var v38 uint64 = v36
            _ = v37
            _ = v38
        }
        var v39 uint = 42
        _ = v39
    } else {
        v2 = (v2 / ((v2) | 1))
    }
    v40 := GenS1{F0: "a", F1: 2, F2: v3}
    if true {
        v1 -= (v1 - (v1 - (v1 | (v2 & v1))))
        if (0 > 1) {
            var v41 uint16 = 255
            v42 := GenS2{F0: (v2 % ((127) | 1))}
            fmt.Println(v5, v2)
            _ = v41
            _ = v42
        } else {
            v43 := GenS1{F0: v5, F1: 51, F2: 2147483647}
            delete(v4, 42)
            _ = v43
        }
        var v44 int64 = 1
        _ = v44
    }
    for fi45 := 0; fi45 < 9; fi45++ {
        for fi46 := 0; fi46 < 7; fi46++ {
            var v47 int32 = v3
            var v48 uint16 = 65535
            v49 := GenS3{F0: (^(v2)), F1: 1000, F2: 1}
            switch v5 {
            case "foo bar":
                var v50 int = fi45
                _ = v50
            case "\n\t":
                v1 = (v1 + v2)
            }
            _ = fi46
            _ = v47
            _ = v48
            _ = v49
        }
        _ = fi45
    }
    _ = len(v4)
    v51 := []int{-9223372036854775808, -314}
    v36 -= v36
    var v52 int32 = v3
    v53 := GenS4{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v51
    _ = v52
    _ = v53
}
