package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint = 914
    var v2 float64 = (-2.25 * 3.14159)
    v3 := make(map[int]bool)
    var v4 uint32 = 1
    v5 := make(map[int]bool)
    var v6 int64 = 769
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            v1 = 0
            if (!true) {
                var v8 uint32 = (^(v4))
                for fi9 := 0; fi9 < 2; fi9++ {
                    if true {
                        v2 = float64((((((1e10 - ((1.5 / 3.14159) / (0.1 / 1.0))) + float32(float64(-2.25))) / 1e10) - 1.5) / ((1.5 * (float32(v2) * 1.5)) - 1.5)))
                    } else {
                        var v10 float32 = 1.5
                        _ = v10
                    }
                    v11 := GenS1{F0: "hello"}
                    if (v2 >= (v2 - float64((float32(float64(-2.25)) - 0.1)))) {
                        if (true || (!true)) {
                            v12 := v3[fi9]
                            _ = v12
                        } else {
                            v13 := GenS1{F0: "foo bar"}
                            var v14 int16 = -32768
                            v6 -= 0
                            _ = v13
                            _ = v14
                        }
                    }
                    var v15 int64 = v6
                    if (18446744073709551615 == 582) {
                        switch "\n\t" {
                        case "hello":
                            if true {
                                fmt.Println(v2, v1)
                            } else {
                                if false {
                                    _ = len(v3)
                                    v16 := GenS1{F0: "foo bar"}
                                    _ = v16
                                } else {
                                    v15 -= 839
                                }
                                switch 128 {
                                case 128:
                                    delete(v3, fi7)
                                case 881:
                                    var v17 int = (fi9 | (^(fi7)))
                                    _ = v17
                                }
                                v18 := GenS2{F0: 32767, F1: (!true)}
                                _ = v18
                            }
                            var v19 uint64 = 517
                            _ = v19
                        default:
                            var v20 bool = (v8 <= (^(v4)))
                            _ = v20
                        }
                        fmt.Println(v4, v4)
                        var v21 int64 = (v15 + -133)
                        _ = v21
                    } else {
                        v4 -= (v8 ^ v8)
                    }
                    _ = fi9
                    _ = v11
                    _ = v15
                }
                for fi22 := 0; fi22 < 2; fi22++ {
                    v5[fi22] = true
                    _ = fi22
                }
                _ = v8
            } else {
                v2 += (1e10 + v2)
                var v23 int64 = v6
                _ = v23
            }
            if (!(("\n\t" + ("foo bar" + "hello")) <= ("\n\t" + "a"))) {
                fmt.Println(v4, v1, v6)
                v2 += v2
            }
        }
        _ = fi7
    }
    delete(v3, 9223372036854775807)
    var v24 float64 = v2
    switch 161 {
    case 127:
        if (v6 <= v6) {
            var v25 int8 = -1
            _ = v25
        }
    case 128:
        v2 = float64(0.1)
    case 9223372036854775807:
        fmt.Println(v6)
    }
    var v26 uint16 = uint16(v4)
    var v27 float32 = (0.0 / 1.5)
    for fi28 := 0; fi28 < 13; fi28++ {
        for fi29 := 0; fi29 < 1; fi29++ {
            if (218 >= (v1 >> ((255) & 63))) {
                var v30 uint16 = (v26 * (v26 << ((v1) & 63)))
                v5[(^(fi28))] = true
                _ = v30
            } else {
                var v31 uint32 = 65535
                var v32 uint8 = 212
                var v33 int16 = 1
                _ = v31
                _ = v32
                _ = v33
            }
            delete(v5, fi29)
            _ = fi29
        }
        if ((v6 % ((v6) | 1)) != 0) {
            v34 := []int{(fi28 ^ -718)}
            for fi35 := 0; fi35 < 6; fi35++ {
                fmt.Println(v24, fi28, fi28)
                v36 := GenS1{F0: ("hello" + ("foo bar" + "a"))}
                v26 = v26
                _ = fi35
                _ = v36
            }
            _ = v34
        } else {
            v24 = (float64(v27) - ((float64((v27 * v27)) * v2) + (1.5 + (0.1 * v24))))
        }
        _ = fi28
    }
    v26 += 1
    v37 := GenS3{F0: false}
    var v38 uint8 = 127
    for fi39 := 0; fi39 < 2; fi39++ {
        _ = len(v3)
        var v40 uint16 = v26
        for fi41 := 0; fi41 < 2; fi41++ {
            v26 -= 65535
            _ = fi41
        }
        _ = fi39
        _ = v40
    }
    switch -9223372036854775808 {
    case 42:
        var v42 int32 = 2147483647
        _ = v42
    default:
        v24 = v24
    }
    var v43 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
