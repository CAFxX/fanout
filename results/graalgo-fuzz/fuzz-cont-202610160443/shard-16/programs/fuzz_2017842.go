package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint16; F1 bool; F2 int64 }
type GenS5 struct { F0 string; F1 int64 }
func genFn0(p0 int16, p1 uint16, p2 uint8) int32 {
    v7 := []int{9223372036854775807, 904, 2147483647, -2147483648}
    var v8 bool = (11 != p1)
    var v9 bool = ((p2 ^ (p2 >> ((874) & 63))) < (p2 / (((p2 & p2)) | 1)))
    v8 = (3.14159 < float32((-2.25 * -2.25)))
    if (4294967295 >= 256) {
        fmt.Println(p2)
        if false {
            if false {
                var v10 uint64 = 426
                _ = v10
            } else {
                if (!(-2.25 != float64((((0.0 * 3.14159) / 1.5) / 3.14159)))) {
                    v8 = (!v9)
                    var v11 float64 = 0.1
                    for fi12 := 0; fi12 < 8; fi12++ {
                        v8 = false
                        _ = fi12
                    }
                    _ = v11
                } else {
                    switch "a" {
                    case "":
                        if (0 >= 42) {
                            if len(v7) > 0 {
                                v7[0] = 32767
                            }
                            _ = len(v7)
                        } else {
                            var v13 uint = 113
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            v14 := GenS1{F0: 4294967295, F1: "\n\t", F2: 396}
                            _ = v13
                            _ = v14
                        }
                    }
                    fmt.Println(p2, p2, p2)
                }
            }
        }
    } else {
        var v15 float64 = 1.5
        var v16 uint8 = (p2 + p2)
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -754
}

func main() {
    var v1 uint = 42
    var v2 int8 = 127
    var v3 string = "\n\t"
    var v4 bool = false
    var v5 string = "foo bar"
    var v6 int32 = int32(v2)
    v17 := []int{489, 1, 297}
    var v18 uint16 = 2
    v6 -= (-(v6))
    if false {
        if v4 {
            var v19 uint = v1
            _ = v19
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS2{F0: 86}
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := GenS3{F0: v6, F1: 491}
            if (v4 && (452 < v6)) {
                if v4 {
                    var v24 string = "hello"
                    if ((36 < 541) && v4) {
                        fmt.Println(v6, v24)
                    }
                    for fi25 := 0; fi25 < 9; fi25++ {
                        v26 := GenS3{F0: 2147483647, F1: 0}
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                }
                fmt.Println(v2, v1)
            } else {
                for fi27 := 0; fi27 < 20; fi27++ {
                    v28 := GenS2{F0: 9223372036854775807}
                    v29 := GenS1{F0: (v1 / ((4294967295) | 1)), F1: "", F2: uint32(v18)}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            if (127 == 9) {
                v30 := GenS1{F0: 846, F1: v5, F2: 742}
                fmt.Println(fi22)
                _ = len(v17)
                _ = v30
            }
            _ = fi22
            _ = v23
        }
    }
    if v4 {
        for fi31 := 0; fi31 < 3; fi31++ {
            switch fi31 {
            case 4294967295:
                v4 = v4
            case 128:
                v32 := GenS4{F0: (^(v18)), F1: (v4 || (!true)), F2: 0}
                _ = v32
            }
            if (1.5 <= -2.25) {
                v2 = v2
                v33 := GenS5{F0: (v3 + ("foo bar" + ("" + (v5 + "hello")))), F1: -9223372036854775808}
                _ = v33
            } else {
                if v4 {
                    var v34 uint8 = uint8(v1)
                    var v35 uint16 = v18
                    var v36 int8 = v2
                    _ = v34
                    _ = v35
                    _ = v36
                } else {
                    var v37 int32 = v6
                    _ = v37
                }
                _ = len(v17)
            }
            var v38 float64 = 3.14159
            _ = fi31
            _ = v38
        }
    } else {
        if (0.1 != 1e10) {
            if len(v17) > 0 {
                v17[0] = 549
            }
            for fi39 := 0; fi39 < 1; fi39++ {
                for fi40 := 0; fi40 < 10; fi40++ {
                    fmt.Println(v3)
                    var v41 bool = v4
                    var v42 uint = v1
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                if v4 {
                    fmt.Println(v5)
                }
                _ = fi39
            }
            v2 = v2
        }
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    var v43 uint64 = 0
    v44 := GenS4{F0: (v18 ^ v18), F1: v4, F2: 9223372036854775807}
    var v45 float32 = 1e10
    var v46 int32 = (v6 * -1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
