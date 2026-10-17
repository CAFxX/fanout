package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 bool; F1 uint; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 uint32 }

func main() {
    var v1 float64 = (1e10 + float64(-2.25))
    var v2 float64 = v1
    var v3 float32 = 0.1
    v4 := []int{32767}
    for fi5 := 0; fi5 < 16; fi5++ {
        v3 += (v3 * 0.1)
        _ = fi5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        v1 = v1
        v7 := GenS1{F0: fi6, F1: false}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        switch "" {
        case "hello":
            if true {
                var v9 int16 = -578
                if (128 > 2) {
                    v1 += v1
                    var v10 int64 = 1
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        if len(v4) > 0 {
                            v4[0] = -35
                        }
                        _ = fi11
                    }
                    var v12 string = (("" + "hello") + "a")
                    v13 := GenS1{F0: 982, F1: (!(((-2.25 / -2.25) <= 0.1) && (false && false)))}
                    _ = v12
                    _ = v13
                }
                _ = v9
            }
        case "foo bar":
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 bool = false
                _ = len(v4)
                if (fi14 > fi14) {
                    v16 := GenS1{F0: (fi14 ^ -815), F1: ((v15 || false) && (((((float64((((-2.25 - (v3 * v3)) + 0.0) + (0.1 * (1e10 + v3)))) * (v1 * v1)) * 0.1) / 1.0) != v2) || (true || false)))}
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
        case "":
            v17 := GenS1{F0: fi8, F1: (0 <= 1)}
            _ = v17
        default:
            for fi18 := 0; fi18 < 20; fi18++ {
                if false {
                    var v19 int32 = 148
                    v20 := GenS1{F0: (^(fi8)), F1: false}
                    if false {
                        v4 = append(v4, (fi18 + 4294967295))
                        v4 = append(v4, fi18)
                        v21 := GenS2{F0: -128, F1: fi8}
                        _ = v21
                    } else {
                        v3 = float32(float64(float32((0.1 + (v1 / v1)))))
                        if len(v4) > 0 {
                            v4[0] = fi18
                        }
                        v2 += v1
                    }
                    _ = v19
                    _ = v20
                } else {
                    var v22 int16 = 414
                    if len(v4) > 0 {
                        v4[0] = (fi18 & fi8)
                    }
                    _ = v22
                }
                v3 -= 1e10
                var v23 uint16 = 661
                var v24 uint8 = 171
                _ = fi18
                _ = v23
                _ = v24
            }
        }
        v4 = append(v4, (^(fi8)))
        fmt.Println(v2)
        _ = fi8
    }
    v2 += ((float64((v3 * float32(1.5))) * (3.14159 + v1)) / 3.14159)
    v25 := GenS1{F0: -9223372036854775808, F1: (!false)}
    v3 = 0.1
    v3 -= v3
    var v26 uint64 = 2
    v4 = append(v4, -200)
    fmt.Println(v1, v1, v3)
    for fi27 := 0; fi27 < 10; fi27++ {
        if (true || (v3 >= 1.5)) {
            v1 += 1e10
        }
        _ = fi27
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        v1 = v2
        _ = fi28
    }
    if true {
        switch -449 {
        case -9223372036854775808:
            if ((3.14159 + (v3 * v3)) > v3) {
                _ = len(v4)
                for fi29 := 0; fi29 < 6; fi29++ {
                    v3 -= 0.1
                    _ = fi29
                }
                var v30 int32 = 1
                _ = v30
            } else {
                switch "" {
                case "hello":
                    var v31 int8 = 127
                    _ = v31
                case "\n\t":
                    if (2 == 2147483647) {
                        for fi32 := 0; fi32 < 1; fi32++ {
                            _ = len(v4)
                            v33 := GenS3{F0: true, F1: 86, F2: 42}
                            _ = fi32
                            _ = v33
                        }
                        v34 := GenS4{F0: 42}
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 18; fi35++ {
                            v36 := GenS2{F0: -2, F1: (fi35 | (^(fi35)))}
                            _ = fi35
                            _ = v36
                        }
                        var v37 uint32 = 993
                        v38 := GenS5{F0: v26}
                        _ = v37
                        _ = v38
                    }
                default:
                    v39 := GenS4{F0: 1}
                    _ = v39
                }
                var v40 float64 = (((0.1 + (v2 + v1)) + (v1 / v1)) / v2)
                v41 := GenS6{F0: -1, F1: 0}
                _ = v40
                _ = v41
            }
        }
    } else {
        v3 = 1e10
    }
    for fi42 := 0; fi42 < 4; fi42++ {
        v1 = v1
        v26 = 834
        var v43 float64 = v1
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
