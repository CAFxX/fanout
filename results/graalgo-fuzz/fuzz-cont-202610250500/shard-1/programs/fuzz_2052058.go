package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int32 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint32 }
func genFn0(p0 int64, p1 uint8) int {
    v7 := []int{-918}
    v8 := GenS1{F0: 42, F1: 1, F2: 2147483647}
    var v9 int16 = 0
    if true {
        var v10 float32 = ((-2.25 / (0.1 / 1.5)) / ((0.1 - 1.5) / -2.25))
        p1 = p1
        v10 = v10
        _ = v10
    } else {
        var v11 int = 4294967295
        v12 := GenS1{F0: 1, F1: 255, F2: 0}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = v11
        _ = v12
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 32768
}
func genFn1(p0 uint, p1 uint64) uint {
    var v13 uint64 = (p1 & (^(p1)))
    var v14 int64 = -244
    var v15 uint8 = 2
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return uint(v15)
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[uint32]int)
    var v3 uint = 143
    v4 := make(map[int]uint32)
    var v5 float64 = -2.25
    v6 := make(map[uint32]uint32)
    if true {
        for fi16 := 0; fi16 < 1; fi16++ {
            if (1 > 4) {
                if true {
                    var v17 uint16 = 0
                    fmt.Println(v17, v3, fi16)
                    _ = v17
                } else {
                    v18 := GenS2{F0: 32767, F1: 65535}
                    var v19 bool = (-9223372036854775808 > 0)
                    _ = v18
                    _ = v19
                }
                fmt.Println(v3, fi16, v1)
                v20 := GenS2{F0: 46, F1: 42}
                _ = v20
            } else {
                if (true && (!true)) {
                    fmt.Println(v3, fi16, v1)
                    v21 := GenS2{F0: -1, F1: 1}
                    _ = v21
                } else {
                    v22 := GenS3{F0: -1}
                    _ = len(v4)
                    var v23 uint8 = 128
                    _ = v22
                    _ = v23
                }
                switch fi16 {
                case 32767:
                    if false {
                        var v24 float64 = 0.1
                        _ = v24
                    } else {
                        if true {
                            for fi25 := 0; fi25 < 8; fi25++ {
                                v26 := GenS2{F0: -32768, F1: 2}
                                var v27 float32 = 3.14159
                                var v28 string = ("" + "hello")
                                v29 := []int{(fi25 & 9223372036854775807)}
                                fmt.Println(v3, fi25, v28)
                                if len(v29) >= 2 {
                                    v29 = v29[1:len(v29)]
                                }
                                v29 = append(v29, (v1 + (v1 - genFn0(1, 127))))
                                _ = fi25
                                _ = v26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                        }
                        var v30 int8 = 0
                        _ = v30
                    }
                case 1:
                    delete(v6, 0)
                default:
                    switch 0 {
                    case 504:
                        for fi31 := 0; fi31 < 16; fi31++ {
                            if (genFn1(v3, 1) != (v3 + 244)) {
                                fmt.Println(fi16)
                                switch -2147483648 {
                                case 32767:
                                    v5 -= 1.5
                                case 743:
                                    v32 := []int{fi31}
                                    _ = v32
                                case -535:
                                    var v33 int64 = -9223372036854775808
                                    _ = v33
                                default:
                                    fmt.Println(fi31, v5)
                                }
                                if (!(genFn1((v3 % ((v3) | 1)), 42) > (v3 << ((v3) & 63)))) {
                                    fmt.Println(v5, v3)
                                    fmt.Println(v1, v1, v5)
                                    v34 := []int{(^(fi31)), (fi16 | 42)}
                                    var v35 uint32 = 42
                                    _ = v34
                                    _ = v35
                                }
                            } else {
                                v36 := GenS1{F0: (v3 - (v3 / (((^(v3))) | 1))), F1: (v3 % ((898) | 1)), F2: 690}
                                var v37 float64 = v5
                                _ = v36
                                _ = v37
                            }
                            v4[32768] = 65535
                            _ = len(v6)
                            v5 = -2.25
                            _ = fi31
                        }
                    case 2147483647:
                        var v38 int16 = 982
                        _ = v38
                    default:
                        var v39 bool = ((v3 << (((v3 | v3)) & 63)) < 989)
                        _ = v39
                    }
                }
            }
            _ = fi16
        }
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        if false {
            for fi41 := 0; fi41 < 18; fi41++ {
                fmt.Println(fi41)
                v42 := GenS4{F0: 708, F1: 527}
                _ = fi41
                _ = v42
            }
        } else {
            var v43 uint64 = 42
            _ = v43
        }
        _ = fi40
    }
    v5 = (v5 * v5)
    v44 := GenS4{F0: -2147483648, F1: 4294967295}
    v3 = 1
    var v45 int8 = -1
    for fi46 := 0; fi46 < 7; fi46++ {
        v1 += 32768
        switch (fi46 | (v1 >> ((v3) & 63))) {
        case 128:
            v47 := GenS4{F0: -2147483648, F1: 2147483647}
            _ = v47
        case 9223372036854775807:
            if false {
                var v48 bool = (4294967295 <= 2147483647)
                v49 := GenS2{F0: -1, F1: 2}
                _ = v48
                _ = v49
            }
        default:
            var v50 uint64 = 1
            _ = v50
        }
        _ = fi46
    }
    v51 := GenS1{F0: (v3 + 2), F1: genFn1(554, 0), F2: 486}
    var v52 bool = true
    v52 = true
    v1 = (v1 & v1)
    for fi53 := 0; fi53 < 8; fi53++ {
        fmt.Println(v1, v1)
        var v54 int = fi53
        var v55 uint16 = 2
        _ = fi53
        _ = v54
        _ = v55
    }
    fmt.Println(v52, v3, v1)
    var v56 int = (-(v1))
    v1 += -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v51
    _ = v52
    _ = v56
    _ = v6
}
