package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int16; F2 int16 }
func genFn0(p0 uint64, p1 string) int64 {
    p0 = p0
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    var v1 int64 = 0
    var v2 uint8 = 7
    var v3 int16 = 0
    var v4 int8 = 42
    v2 += v2
    v5 := GenS1{F0: (v1 << ((4294967295) & 63))}
    var v6 bool = (!false)
    if v6 {
        if v6 {
            v7 := []int{127, 127, 32768}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 2; fi8++ {
                for fi9 := 0; fi9 < 9; fi9++ {
                    v10 := []int{(fi8 % ((0) | 1)), fi9, fi9}
                    v1 = genFn0(1, "foo bar")
                    _ = fi9
                    _ = v10
                }
                v2 += (v2 | 169)
                _ = fi8
            }
            switch "hello" {
            case "foo bar":
                for fi11 := 0; fi11 < 13; fi11++ {
                    if (2 > 0) {
                        for fi12 := 0; fi12 < 10; fi12++ {
                            if (342 != 845) {
                                if (1 >= 18446744073709551615) {
                                    v13 := GenS1{F0: (-(v1))}
                                    var v14 float32 = (0.0 / ((-2.25 + 1e10) - 1.5))
                                    _ = v13
                                    _ = v14
                                }
                                var v15 uint32 = 775
                                if (v6 || v6) {
                                    if (!((42 > 960) && v6)) {
                                        var v16 bool = ("a" < "hello")
                                        v16 = v6
                                        var v17 bool = false
                                        _ = v16
                                        _ = v17
                                    }
                                } else {
                                    v18 := []int{fi12, fi11}
                                    _ = len(v18)
                                    var v19 uint16 = 587
                                    _ = v18
                                    _ = v19
                                }
                                var v20 int64 = genFn0(351, "\n\t")
                                _ = v15
                                _ = v20
                            } else {
                                v2 = 1
                                v21 := GenS1{F0: -435}
                                _ = v21
                            }
                            v22 := GenS2{F0: 804}
                            _ = fi12
                            _ = v22
                        }
                        var v23 int16 = v3
                        _ = v23
                    } else {
                        var v24 uint16 = 274
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 1; fi25++ {
                        var v26 int8 = v4
                        v26 = 100
                        _ = fi25
                        _ = v26
                    }
                    _ = fi11
                }
            }
        }
    } else {
        v1 += v1
        v27 := []int{47}
        _ = v27
    }
    if (!(395 != 2)) {
        var v28 uint = 42
        _ = v28
    }
    var v29 uint32 = 2
    for fi30 := 0; fi30 < 17; fi30++ {
        v2 = 0
        for fi31 := 0; fi31 < 11; fi31++ {
            fmt.Println(v1, v2, v29)
            _ = fi31
        }
        if ((v2 / ((v2) | 1)) == (v2 * v2)) {
            v6 = (!v6)
            v32 := GenS3{F0: v6}
            fmt.Println(v1, v6, fi30)
            _ = v32
        }
        _ = fi30
    }
    v1 = v1
    v33 := GenS3{F0: ((v2 <= v2) || v6)}
    v6 = true
    for fi34 := 0; fi34 < 8; fi34++ {
        var v35 int32 = 1
        var v36 uint = 983
        var v37 int64 = (v1 / ((genFn0(742, ("foo bar" + (((("hello" + "") + "\n\t") + "hello") + "foo bar")))) | 1))
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        v39 := GenS3{F0: (65535 >= 0)}
        _ = fi38
        _ = v39
    }
    v29 = v29
    v4 = (v4 >> ((65535) & 63))
    for fi40 := 0; fi40 < 18; fi40++ {
        if (!(-2.25 <= float64((0.0 * ((-2.25 * 0.0) - (1e10 + 3.14159)))))) {
            v41 := GenS3{F0: v6}
            _ = v41
        }
        for fi42 := 0; fi42 < 13; fi42++ {
            if v6 {
                v43 := []int{}
                var v44 int32 = -1
                var v45 int32 = v44
                _ = v43
                _ = v44
                _ = v45
            } else {
                if v6 {
                    switch ("" + "a") {
                    case "\n\t":
                        v1 = (v1 / (((v1 % ((genFn0(4294967295, ("" + "foo bar"))) | 1))) | 1))
                    }
                    switch fi42 {
                    case -823:
                        var v46 uint16 = 255
                        _ = v46
                    }
                } else {
                    v1 += v1
                    var v47 uint64 = 1
                    fmt.Println(v4)
                    _ = v47
                }
                if v6 {
                    v48 := GenS4{F0: (v2 & v2), F1: (v3 % ((-1) | 1)), F2: (v3 & v3)}
                    var v49 uint64 = 430
                    fmt.Println(v1, fi40)
                    _ = v48
                    _ = v49
                }
            }
            v50 := []int{fi42, (fi42 >> ((42) & 63)), (fi40 - (fi42 - fi40)), int(v4)}
            v6 = true
            v3 -= v3
            _ = fi42
            _ = v50
        }
        v51 := GenS4{F0: (v2 & (v2 >> ((519) & 63))), F1: v3, F2: (v3 ^ -318)}
        _ = fi40
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
