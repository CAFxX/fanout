package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int16; F1 bool }
func genFn0() int {
    var v5 int32 = -321
    v6 := GenS1{F0: 1, F1: 299}
    for fi7 := 0; fi7 < 6; fi7++ {
        if (!true) {
            v8 := []int{fi7, fi7, fi7}
            _ = len(v8)
            _ = v8
        } else {
            if (42 == 1) {
                v9 := GenS1{F0: 57, F1: 9223372036854775807}
                fmt.Println(v5, v5)
                v5 = v5
                _ = v9
            } else {
                var v10 int64 = -9223372036854775808
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 int32 = v5
                    v10 -= (v10 / ((0) | 1))
                    fmt.Println(v5)
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
            var v13 uint8 = 1
            v5 = 42
            _ = v13
        }
        switch (fi7 | 0) {
        case 0:
            for fi14 := 0; fi14 < 18; fi14++ {
                fmt.Println(fi7)
                _ = fi14
            }
        case 4294967295:
            v15 := []int{(fi7 >> ((4294967295) & 63))}
            _ = v15
        case 42:
            v5 += (v5 - (v5 ^ v5))
        default:
            v5 = (v5 + (^(v5)))
        }
        _ = fi7
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        if (!(0 == 104)) {
            v17 := []int{(fi16 * -1), (fi16 * fi16), fi16}
            v5 -= 2147483647
            if len(v17) > 0 {
                v17[0] = 992
            }
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            v5 = 785
            _ = fi18
        }
        _ = fi16
    }
    _ = v5
    _ = v6
    return 745
}

func main() {
    var v1 float64 = 3.14159
    var v2 float64 = (1.5 / -2.25)
    v3 := []int{-229, 1, 9223372036854775807, -2147483648}
    v4 := make(map[string]uint32)
    if true {
        v4["foo bar"] = 42
        v19 := v4[("a" + "")]
        _ = v19
    } else {
        v1 = v2
        var v20 bool = true
        v1 += v1
        _ = v20
    }
    v2 = (3.14159 - float64(0.1))
    fmt.Println(v2, v2)
    fmt.Println(v2)
    var v21 uint8 = 151
    var v22 uint16 = 255
    _ = len(v4)
    for fi23 := 0; fi23 < 2; fi23++ {
        v24 := GenS1{F0: -1, F1: (fi23 & fi23)}
        if ((100 == 127) && true) {
            var v25 int16 = 42
            _ = v25
        } else {
            var v26 uint8 = v21
            if (true && true) {
                v27 := GenS2{F0: 0, F1: (29 >= 4294967295)}
                if false {
                    var v28 int32 = -2147483648
                    _ = v28
                } else {
                    if true {
                        var v29 int8 = 100
                        _ = v29
                    }
                    v1 += (0.1 / v1)
                }
                _ = v27
            }
            _ = v26
        }
        if (v22 <= 670) {
            if (!false) {
                v30 := GenS1{F0: 100, F1: (-(fi23))}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch ((("hello" + "") + "a") + "foo bar") {
                case "hello":
                    if true {
                        _ = len(v3)
                        if (190 > 1) {
                            fmt.Println(v2)
                            switch "" {
                            case "foo bar":
                                v31 := v4[""]
                                _ = v31
                            }
                        } else {
                            fmt.Println(v2, v21)
                            if (!(0.0 >= float64(0.0))) {
                                v32 := GenS2{F0: 32767, F1: false}
                                _ = v32
                            } else {
                                var v33 int = fi23
                                _ = v33
                            }
                        }
                    }
                case "a":
                    for fi34 := 0; fi34 < 13; fi34++ {
                        if (1000 > -911) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            v35 := GenS2{F0: 1, F1: false}
                            var v36 int = 128
                            _ = v35
                            _ = v36
                        }
                        v21 = v21
                        v22 -= v22
                        _ = fi34
                    }
                case "":
                    for fi37 := 0; fi37 < 8; fi37++ {
                        var v38 uint8 = 1
                        v21 = v38
                        _ = fi37
                        _ = v38
                    }
                default:
                    if (1 > 9223372036854775807) {
                        v2 += v2
                        if len(v3) > 0 {
                            v3[0] = fi23
                        }
                        switch int(v22) {
                        case 127:
                            v21 -= (v21 ^ (v21 << ((2) & 63)))
                        case 42:
                            v22 = 65535
                        }
                    } else {
                        v2 = v2
                    }
                }
                _ = v30
            } else {
                switch genFn0() {
                case -10:
                    fmt.Println(v1, v1, v2)
                }
                for fi39 := 0; fi39 < 15; fi39++ {
                    v2 -= v1
                    _ = fi39
                }
                for fi40 := 0; fi40 < 20; fi40++ {
                    v1 += v2
                    v3 = append(v3, 2147483647)
                    _ = fi40
                }
            }
        }
        _ = fi23
        _ = v24
    }
    if true {
        var v41 uint = 1
        var v42 int16 = 32767
        _ = v41
        _ = v42
    } else {
        for fi43 := 0; fi43 < 19; fi43++ {
            fmt.Println(v22, v1, v21)
            _ = fi43
        }
        v21 += v21
        _ = len(v4)
    }
    v2 -= (v1 - (1e10 + v1))
    v22 = (v22 + 694)
    v1 = 1.5
    var v44 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v44
}
