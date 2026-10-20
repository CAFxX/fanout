package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint8 }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: -128, F1: (fi7 & fi7)}
        var v9 int = 127
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{935, 142, 1, -2147483648}
    for fi11 := 0; fi11 < 11; fi11++ {
        if (fi11 >= fi11) {
            v12 := GenS1{F0: -37, F1: (fi11 & (fi11 / ((fi11) | 1)))}
            var v13 int16 = -877
            _ = v12
            _ = v13
        } else {
            var v14 int8 = -78
            _ = v14
        }
        _ = fi11
    }
    _ = v10
    return 876
}

func main() {
    var v1 float32 = -2.25
    var v2 int = 520
    var v3 float64 = (0.0 + (0.1 + 1.5))
    var v4 int64 = 9223372036854775807
    var v5 uint = 4294967295
    v6 := []uint32{1, 42, 4294967295}
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint16 = 42
        _ = fi15
        _ = v16
    }
    var v17 uint32 = 65535
    for fi18 := 0; fi18 < 8; fi18++ {
        v17 = 1
        fmt.Println(v3, v2)
        var v19 bool = false
        _ = fi18
        _ = v19
    }
    v2 += (v2 - (v2 & v2))
    var v20 int32 = 1
    switch (("foo bar" + "foo bar") + "a") {
    case "":
        switch ("hello" + "foo bar") {
        case "hello":
            for fi21 := 0; fi21 < 5; fi21++ {
                if (int(v4) >= -887) {
                    var v22 uint16 = 42
                    _ = v22
                } else {
                    var v23 int64 = (v4 + (v4 | (v4 & (v4 / (((v4 << (((v5 % ((v5) | 1))) & 63))) | 1)))))
                    v3 += (v3 * (float64((3.14159 / float32(v3))) / v3))
                    fmt.Println(v4, v4, v17)
                    _ = v23
                }
                _ = fi21
            }
        default:
            v24 := GenS1{F0: int8(v20), F1: 9}
            _ = v24
        }
    case "foo bar":
        if (-1 > 42) {
            switch v2 {
            case 32768:
                if ((v20 ^ 42) > (v20 | (v20 | 2147483647))) {
                    if true {
                        v25 := GenS1{F0: 65, F1: v2}
                        v26 := GenS1{F0: 83, F1: (v2 % ((371) | 1))}
                        v20 -= -1
                        _ = v25
                        _ = v26
                    }
                    if (-46 == -1) {
                        v4 += 42
                        if false {
                            var v27 float64 = 3.14159
                            _ = len(v6)
                            _ = v27
                        } else {
                            v3 -= (v3 * v3)
                            v2 = (-(v2))
                        }
                    } else {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            if false {
                                v29 := GenS1{F0: 100, F1: fi28}
                                var v30 float32 = -2.25
                                _ = v29
                                _ = v30
                            } else {
                                switch v2 {
                                case 0:
                                    _ = len(v6)
                                case -26:
                                    for fi31 := 0; fi31 < 2; fi31++ {
                                        var v32 float64 = -2.25
                                        _ = fi31
                                        _ = v32
                                    }
                                case 1:
                                    var v33 uint16 = genFn0()
                                    _ = v33
                                default:
                                    var v34 float32 = v1
                                    _ = v34
                                }
                            }
                            _ = fi28
                        }
                        for fi35 := 0; fi35 < 12; fi35++ {
                            var v36 int8 = -128
                            var v37 int16 = 0
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    }
                    var v38 bool = true
                    _ = v38
                }
            default:
                switch "foo bar" {
                case "a":
                    v39 := GenS1{F0: 42, F1: -2147483648}
                    _ = v39
                default:
                    if len(v6) > 0 {
                        v6[0] = 4294967295
                    }
                }
            }
            v3 = (float64(0.0) / 1e10)
        }
    case "a":
        fmt.Println(v20)
    default:
        for fi40 := 0; fi40 < 13; fi40++ {
            if true {
                v3 -= v3
                var v41 string = ("foo bar" + "hello")
                switch fi40 {
                case 42:
                    for fi42 := 0; fi42 < 7; fi42++ {
                        v43 := GenS1{F0: -61, F1: 2147483647}
                        if true {
                            var v44 uint32 = v17
                            if len(v6) > 0 {
                                v6[0] = v44
                            }
                            if ((("foo bar" + v41) + ("a" + v41)) > "hello") {
                                var v45 uint64 = 274
                                var v46 uint8 = 128
                                v1 = v1
                                _ = v45
                                _ = v46
                            }
                            _ = v44
                        }
                        v47 := GenS2{F0: 17}
                        _ = fi42
                        _ = v43
                        _ = v47
                    }
                case 572:
                    switch v41 {
                    case "a":
                        var v48 uint32 = v17
                        _ = v48
                    default:
                        var v49 int = (-(fi40))
                        _ = v49
                    }
                case 1:
                    v50 := GenS1{F0: -35, F1: v2}
                    _ = v50
                default:
                    switch v41 {
                    case "foo bar":
                        if len(v6) > 0 {
                            v6[0] = (v17 & (^(v17)))
                        }
                    }
                }
                _ = v41
            }
            _ = fi40
        }
    }
    v5 = v5
    var v51 uint16 = 255
    v52 := GenS1{F0: -128, F1: -349}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v53 int16 = -441
    v5 += v5
    v17 = 2147483647
    var v54 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
