package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int64; F2 int32 }
type GenS4 struct { F0 bool; F1 bool; F2 int16 }
type GenS5 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint64; F1 int8; F2 string }
func genFn0(p0 int32, p1 bool) uint {
    if true {
        if p1 {
            var v5 uint16 = 255
            _ = v5
        } else {
            var v6 uint8 = 255
            _ = v6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 bool = true
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS1{F0: 0, F1: 1, F2: 0}
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    } else {
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := []int{fi11, fi11}
            var v13 uint32 = 2
            var v14 uint32 = 36
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        var v15 int64 = -544
        fmt.Println(p1)
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 int = fi16
        v18 := GenS2{F0: 1}
        var v19 float32 = -2.25
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := []int{-668, -9223372036854775808}
    p1 = (2 > 42)
    _ = p0
    _ = p1
    _ = v20
    return 255
}

func main() {
    v1 := []uint32{1}
    var v2 uint32 = 65535
    var v3 float64 = 1.5
    var v4 int64 = -9223372036854775808
    switch "foo bar" {
    case "a":
        var v21 int16 = 1000
        _ = v21
    case "hello":
        v3 = (v3 / float64((1e10 - 0.1)))
    case "\n\t":
        if true {
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 uint = 0
                _ = fi22
                _ = v23
            }
            fmt.Println(v2)
            var v24 int8 = -1
            _ = v24
        }
    default:
        v2 = 468
    }
    v25 := GenS3{F0: 42, F1: v4, F2: -985}
    if (false && (v2 < (^(v2)))) {
        if len(v1) > 0 {
            v1[0] = (v2 + (v2 | 59))
        }
        v26 := GenS4{F0: (-1 == 2147483647), F1: (-9223372036854775808 < -1), F2: -32768}
        var v27 bool = false
        _ = v26
        _ = v27
    }
    var v28 uint64 = 0
    if len(v1) > 0 {
        v1[0] = v2
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        if (true || false) {
            var v30 int64 = (^(v4))
            if true {
                v3 = (3.14159 + 1e10)
            }
            v28 -= (v28 - v28)
            _ = v30
        }
        fmt.Println(v4, v3)
        _ = fi29
    }
    v31 := GenS1{F0: 282, F1: 241, F2: 0}
    if true {
        if false {
            var v32 int = 128
            _ = v32
        } else {
            var v33 int32 = 42
            if false {
                v4 = 9223372036854775807
            } else {
                v4 += v4
                var v34 uint32 = 4294967295
                _ = v34
            }
            if (!((v28 << ((820) & 63)) == 0)) {
                var v35 uint64 = (v28 | v28)
                v33 -= (v33 + 0)
                if true {
                    v28 += (^(v28))
                    for fi36 := 0; fi36 < 9; fi36++ {
                        for fi37 := 0; fi37 < 12; fi37++ {
                            v1 = append(v1, v2)
                            _ = fi37
                        }
                        _ = len(v1)
                        _ = fi36
                    }
                    if (v4 <= -1) {
                        for fi38 := 0; fi38 < 9; fi38++ {
                            var v39 int8 = -128
                            if true {
                                v40 := GenS1{F0: 2, F1: 65535, F2: int16(v35)}
                                _ = v40
                            } else {
                                for fi41 := 0; fi41 < 15; fi41++ {
                                    var v42 int16 = -948
                                    switch (fi41 & -924) {
                                    case 939:
                                        v42 = 42
                                    case 713:
                                        fmt.Println(fi41, v42, fi41)
                                    default:
                                        v2 += v2
                                    }
                                    v43 := GenS5{F0: v39, F1: 255, F2: 313}
                                    _ = fi41
                                    _ = v42
                                    _ = v43
                                }
                            }
                            switch fi38 {
                            case -2147483648:
                                switch fi38 {
                                case 42:
                                    _ = len(v1)
                                case 32767:
                                    for fi44 := 0; fi44 < 11; fi44++ {
                                        var v45 uint64 = 4294967295
                                        var v46 int = 127
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                    }
                                default:
                                    v33 = (v33 + (v33 ^ (v33 & v33)))
                                }
                            case -1:
                                v47 := GenS5{F0: v39, F1: 198, F2: 2}
                                _ = v47
                            case 674:
                                var v48 int8 = v39
                                _ = v48
                            default:
                                v39 = v39
                            }
                            _ = fi38
                            _ = v39
                        }
                        v35 = v35
                        var v49 float64 = v3
                        _ = v49
                    }
                } else {
                    _ = len(v1)
                    var v50 uint = 791
                    v51 := GenS2{F0: 2}
                    _ = v50
                    _ = v51
                }
                _ = v35
            }
            _ = v33
        }
        v4 += (v4 << ((471) & 63))
    }
    v52 := GenS5{F0: 0, F1: 0, F2: 403}
    if ((42 != 715) && (!false)) {
        if len(v1) > 0 {
            v1[0] = 930
        }
        if (true && true) {
            fmt.Println(v4, v3, v2)
        }
    }
    v3 -= (((v3 * (3.14159 * float64((float32(0.0) - (3.14159 + (float32(0.0) / (-2.25 + (-2.25 / -2.25)))))))) / (v3 - (-2.25 * v3))) - v3)
    if (!(v2 <= 42)) {
        var v53 int16 = 1
        var v54 int = int(v28)
        _ = v53
        _ = v54
    }
    fmt.Println(v3, v2, v3)
    v28 = (v28 | 300)
    v55 := GenS6{F0: v28, F1: 64, F2: "hello"}
    var v56 bool = (false || false)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v52
    _ = v55
    _ = v56
}
