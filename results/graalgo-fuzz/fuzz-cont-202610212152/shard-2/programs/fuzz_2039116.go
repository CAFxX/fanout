package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint }
type GenS2 struct { F0 uint32; F1 uint; F2 string }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int32; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint; F1 int }
func genFn0() int64 {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int8 = -1
        switch fi5 {
        case 42:
            var v7 bool = true
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 string = "\n\t"
        _ = fi8
        _ = v9
    }
    var v10 uint = 126
    _ = v10
    return 825
}

func main() {
    v1 := []int{893, -9223372036854775808, -15}
    var v2 uint16 = 32767
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v11 := GenS1{F0: 18446744073709551615, F1: 993, F2: 4294967295}
    for fi12 := 0; fi12 < 8; fi12++ {
        v2 = v2
        v2 -= 42
        _ = fi12
    }
    v13 := v3[("a" + "\n\t")]
    if v13 {
        switch 42 {
        case 32767:
            v2 -= 42
        }
        v14 := GenS1{F0: 2, F1: 362, F2: 18446744073709551615}
        v13 = (-1 < -330)
        _ = v14
    }
    v2 -= 2
    fmt.Println(v2, v13)
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS1{F0: 255, F1: 1, F2: 42}
        _ = fi15
        _ = v16
    }
    var v17 int32 = -850
    var v18 uint8 = 2
    if (123 > 2) {
        _ = len(v4)
        for fi19 := 0; fi19 < 12; fi19++ {
            if (3.14159 != 3.14159) {
                var v20 int = fi19
                var v21 float64 = -2.25
                _ = v20
                _ = v21
            }
            _ = fi19
        }
    }
    v22 := GenS2{F0: 65535, F1: 822, F2: "foo bar"}
    v23 := v3[("foo bar" + ("hello" + "foo bar"))]
    v24 := GenS1{F0: 435, F1: 2147483647, F2: 255}
    var v25 int64 = 1
    var v26 int64 = 0
    if (0 == 42) {
        v2 = uint16(v18)
        v27 := GenS3{F0: 4294967295, F1: "a"}
        _ = v27
    } else {
        if (true && (!((42 == -1) || false))) {
            if len(v1) > 0 {
                v1[0] = 492
            }
            if (false || (v13 && (0.1 < (((float64(float32(((1e10 / 1e10) / 1.5))) / (1.5 - (1.5 / -2.25))) - float64(((-2.25 * 1e10) / 1.0))) / 1.0)))) {
                for fi28 := 0; fi28 < 15; fi28++ {
                    v2 = (v2 ^ 970)
                    if v23 {
                        v18 -= v18
                        for fi29 := 0; fi29 < 11; fi29++ {
                            var v30 float32 = 1e10
                            _ = fi29
                            _ = v30
                        }
                    }
                    _ = fi28
                }
            }
        } else {
            for fi31 := 0; fi31 < 18; fi31++ {
                v23 = true
                var v32 int16 = -987
                _ = fi31
                _ = v32
            }
            if false {
                v33 := GenS3{F0: 4294967295, F1: "foo bar"}
                if false {
                    v26 += (v25 & (v26 ^ -1))
                    if v13 {
                        fmt.Println(v25)
                        for fi34 := 0; fi34 < 18; fi34++ {
                            v35 := GenS4{F0: int32(v18), F1: (v26 ^ (v25 ^ (v26 & genFn0()))), F2: (v2 >> ((148) & 63))}
                            _ = fi34
                            _ = v35
                        }
                        for fi36 := 0; fi36 < 15; fi36++ {
                            v2 += v2
                            if false {
                                for fi37 := 0; fi37 < 14; fi37++ {
                                    v26 -= v25
                                    switch (fi36 << ((65535) & 63)) {
                                    case 2147483647:
                                        v38 := GenS5{F0: 255, F1: fi37}
                                        _ = v38
                                    case 32767:
                                        fmt.Println(v2, v2, v18)
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if (!((!v23) || v13)) {
                                        delete(v4, fi37)
                                        var v39 int16 = 42
                                        var v40 uint32 = 2
                                        _ = v39
                                        _ = v40
                                    }
                                    _ = fi37
                                }
                            } else {
                                v41 := v3[((("foo bar" + "foo bar") + "") + "a")]
                                if len(v1) > 0 {
                                    v1[0] = fi36
                                }
                                if v41 {
                                    v26 = (v25 - (v25 + int64(v17)))
                                    var v42 int8 = 100
                                    _ = v42
                                }
                                _ = v41
                            }
                            var v43 uint = 250
                            _ = fi36
                            _ = v43
                        }
                    }
                }
                v25 += 0
                _ = v33
            }
        }
        for fi44 := 0; fi44 < 2; fi44++ {
            var v45 int32 = v17
            var v46 uint64 = 0
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = len(v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
