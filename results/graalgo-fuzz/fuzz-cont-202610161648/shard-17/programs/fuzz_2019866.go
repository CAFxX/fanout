package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int64; F2 int }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int64; F1 int16; F2 uint32 }
type GenS6 struct { F0 int16; F1 int64 }
func genFn0(p0 uint16, p1 int64, p2 uint32) uint16 {
    fmt.Println(p2, p1)
    var v7 int32 = 1
    var v8 int64 = 277
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 bool, p1 uint) uint32 {
    var v9 string = "hello"
    var v10 uint32 = 931
    v11 := GenS1{F0: 128, F1: 42}
    var v12 int16 = 1
    v13 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v9
    return v10
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 86
    var v3 uint16 = 2
    var v4 float32 = (0.0 * (0.0 - (1.5 - 3.14159)))
    var v5 int = 32768
    var v6 int16 = 871
    v4 = float32(1.5)
    v4 = v4
    if false {
        v5 = (v5 / (((v5 - 170)) | 1))
        if true {
            v1 = 42
            v14 := []int{v5, v5, v5}
            _ = v14
        } else {
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := GenS2{F0: 0}
                for fi17 := 0; fi17 < 20; fi17++ {
                    var v18 string = ("a" + "\n\t")
                    _ = fi17
                    _ = v18
                }
                v19 := GenS2{F0: 587}
                _ = fi15
                _ = v16
                _ = v19
            }
            var v20 int8 = int8(v2)
            switch "a" {
            case "":
                if false {
                    for fi21 := 0; fi21 < 5; fi21++ {
                        switch (v5 * 2147483647) {
                        case 0:
                            fmt.Println(v20, v6, v6)
                        default:
                            for fi22 := 0; fi22 < 13; fi22++ {
                                v2 = (v2 | (^(v2)))
                                if ((v2 * v2) > 2) {
                                    var v23 int64 = 666
                                    _ = v23
                                } else {
                                    if ((fi21 >> ((65535) & 63)) > v5) {
                                        var v24 uint32 = 1
                                        v25 := GenS3{F0: 0, F1: -9223372036854775808, F2: v5}
                                        v6 = 1000
                                        _ = v24
                                        _ = v25
                                    }
                                    v26 := []int{(v5 / ((fi21) | 1)), (fi21 + (fi22 | (fi21 >> ((685) & 63)))), (fi21 + (^(v5)))}
                                    var v27 float32 = float32(1e10)
                                    var v28 uint64 = 98
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                _ = fi22
                            }
                        }
                        _ = fi21
                    }
                    fmt.Println(v6, v5, v2)
                } else {
                    if (291 == 42) {
                        switch 42 {
                        case 127:
                            var v29 int8 = v20
                            _ = v29
                        case 128:
                            v30 := []int{-991}
                            _ = v30
                        case -981:
                            v31 := []int{-2147483648, (-(v5))}
                            _ = v31
                        }
                        switch (v5 >> ((18446744073709551615) & 63)) {
                        case 466:
                            var v32 uint64 = 4294967295
                            _ = v32
                        case 42:
                            var v33 int8 = (v20 % (((v20 | (v20 | 0))) | 1))
                            _ = v33
                        case 127:
                            switch (v5 % (((v5 >> ((65535) & 63))) | 1)) {
                            case -407:
                                var v34 int16 = (v6 & (v6 | (v6 % ((v6) | 1))))
                                _ = v34
                            case -284:
                                if (false && true) {
                                    fmt.Println(v2, v6, v3)
                                    v35 := []int{(v5 - v5), v5, (v5 >> ((1) & 63))}
                                    if false {
                                        v36 := GenS4{F0: 2, F1: 0}
                                        if (!true) {
                                            v6 -= (^(v6))
                                            if (true && true) {
                                                switch v5 {
                                                case 9223372036854775807:
                                                    var v37 int64 = int64(v2)
                                                    _ = v37
                                                case 0:
                                                    fmt.Println(v4, v5)
                                                default:
                                                    v6 -= -555
                                                }
                                                var v38 uint64 = 42
                                                _ = v38
                                            }
                                        }
                                        for fi39 := 0; fi39 < 10; fi39++ {
                                            var v40 int8 = v20
                                            _ = fi39
                                            _ = v40
                                        }
                                        _ = v36
                                    } else {
                                        v41 := GenS5{F0: 535, F1: (v6 >> ((65535) & 63)), F2: 2147483648}
                                        v3 += v3
                                        v42 := GenS2{F0: 902}
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = v35
                                } else {
                                    var v43 int64 = -1
                                    _ = v43
                                }
                            default:
                                fmt.Println(v2, v2, v4)
                            }
                        }
                    } else {
                        var v44 string = "foo bar"
                        _ = v44
                    }
                    v2 = (v2 - (^(v2)))
                    v45 := []int{(^(v5)), (v5 - v5)}
                    _ = v45
                }
            }
            _ = v20
        }
        v46 := []int{v5}
        _ = v46
    }
    switch v5 {
    case -372:
        fmt.Println(v3, v6, v6)
    case 1:
        v47 := GenS6{F0: (v6 + v6), F1: -1}
        _ = v47
    case 828:
        v48 := []int{-1}
        _ = v48
    }
    fmt.Println(v5, v6, v4)
    v49 := []int{127}
    v1 = v1
    var v50 uint8 = v2
    fmt.Println(v5, v3)
    if len(v49) >= 2 {
        v49 = v49[1:len(v49)]
    }
    var v51 float64 = float64((((v4 - v4) * 0.1) * (v4 / v4)))
    v4 = v4
    if (!true) {
        v4 = v4
    }
    var v52 int = (^(v5))
    var v53 uint64 = 1
    var v54 int16 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
