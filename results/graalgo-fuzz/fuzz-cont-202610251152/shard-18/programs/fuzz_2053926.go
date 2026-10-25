package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 int32 }
func genFn0(p0 int8, p1 string, p2 int8) bool {
    var v5 uint8 = 165
    var v6 float32 = (((-2.25 - 0.0) - -2.25) + 0.1)
    var v7 uint32 = 1
    v8 := []int{int(p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 int64, p1 int8) int32 {
    var v9 int8 = p1
    fmt.Println(v9)
    _ = p0
    _ = p1
    _ = v9
    return 560
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float64 = 0.0
    var v3 int8 = 100
    var v4 uint8 = 127
    var v10 int64 = v1
    v4 = v4
    switch 647 {
    case 4294967295:
        v11 := []int{1, 128, int(v3)}
        _ = v11
    case -9223372036854775808:
        var v12 int8 = v3
        _ = v12
    case 1:
        if true {
            v1 -= v10
        } else {
            v13 := GenS1{F0: (((0.0 != ((v2 / (((float64(-2.25) * v2) - (1e10 - 0.1)) * 0.1)) - (1e10 / float64(1.5)))) && true) || true), F1: v1, F2: 1}
            v2 += v2
            v10 = v10
            _ = v13
        }
    default:
        v14 := GenS2{F0: 187}
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        if ("" <= "hello") {
            switch (-(fi15)) {
            case 541:
                v3 = 127
            case 2147483647:
                fmt.Println(v10)
            default:
                v10 = v1
            }
            var v16 int32 = -2147483648
            _ = v16
        } else {
            for fi17 := 0; fi17 < 14; fi17++ {
                v18 := GenS3{F0: 0}
                var v19 uint = 1
                v20 := []int{(fi17 | -2147483648), (-(fi17))}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            if true {
                v21 := GenS3{F0: 2}
                if false {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        switch "\n\t" {
                        case "":
                            fmt.Println(v3, v4)
                            if (false && false) {
                                v1 = (v10 % ((-9223372036854775808) | 1))
                                var v23 bool = true
                                _ = v23
                            }
                            switch "\n\t" {
                            case "foo bar":
                                v2 -= ((-2.25 / (v2 * 3.14159)) / (-2.25 * v2))
                            case "\n\t":
                                var v24 int = fi22
                                _ = v24
                            }
                        }
                        _ = fi22
                    }
                    v25 := []int{(fi15 * fi15), (fi15 % ((-2147483648) | 1)), fi15, (fi15 + (fi15 - (fi15 % ((4294967295) | 1))))}
                    var v26 int8 = v3
                    _ = v25
                    _ = v26
                } else {
                    fmt.Println(v3)
                    if true {
                        v3 = v3
                    }
                }
                if false {
                    v27 := GenS4{F0: (v4 * v4)}
                    _ = v27
                }
                _ = v21
            }
        }
        v10 = v1
        _ = fi15
    }
    var v28 bool = (((((v3 * -1) <= v3) || true) || false) && (-723 > 172))
    v29 := GenS5{F0: (v3 * (v3 ^ 42))}
    v30 := GenS5{F0: (-(v3))}
    v31 := GenS1{F0: (v10 >= (^(v1))), F1: (-(v10)), F2: 0}
    v32 := GenS3{F0: 7}
    v33 := []int{-514, -1}
    v34 := GenS6{F0: "\n\t", F1: -29}
    for fi35 := 0; fi35 < 5; fi35++ {
        v10 = v1
        v3 -= v3
        switch 2147483647 {
        case 9223372036854775807:
            if (65535 == 1) {
                v36 := GenS5{F0: -1}
                v37 := GenS7{F0: 65535, F1: fi35}
                _ = v36
                _ = v37
            }
        default:
            var v38 uint = 65535
            _ = v38
        }
        _ = fi35
    }
    v39 := GenS8{F0: -349}
    switch ("hello" + "a") {
    case "a":
        var v40 bool = false
        _ = v40
    default:
        for fi41 := 0; fi41 < 3; fi41++ {
            for fi42 := 0; fi42 < 17; fi42++ {
                if len(v33) >= 2 {
                    v33 = v33[1:len(v33)]
                }
                _ = fi42
            }
            _ = fi41
        }
    }
    for fi43 := 0; fi43 < 13; fi43++ {
        var v44 int8 = (v3 + 19)
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        var v45 int32 = -106
        _ = fi43
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        var v47 int16 = 1000
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
}
