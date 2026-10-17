package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int32; F1 bool; F2 uint8 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int64; F1 int64; F2 uint8 }
type GenS8 struct { F0 string }

func main() {
    var v1 string = "a"
    var v2 string = (("hello" + (v1 + "a")) + v1)
    v3 := make(map[uint32]bool)
    var v4 float64 = float64(0.1)
    v5 := []int{-90}
    v6 := make(map[int]uint32)
    var v7 int16 = 1000
    switch 1 {
    case 32767:
        v8 := GenS1{F0: 707, F1: 65535, F2: 255}
        _ = v8
    case 128:
        if false {
            if true {
                for fi9 := 0; fi9 < 3; fi9++ {
                    fmt.Println(v2, v4)
                    _ = fi9
                }
            }
            var v10 uint = 255
            if true {
                for fi11 := 0; fi11 < 16; fi11++ {
                    if true {
                        fmt.Println(v1, v4)
                        var v12 uint = v10
                        _ = v12
                    }
                    v3[715] = false
                    _ = fi11
                }
                v13 := GenS2{F0: 42, F1: -1}
                _ = len(v6)
                _ = v13
            } else {
                _ = len(v5)
                for fi14 := 0; fi14 < 8; fi14++ {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        var v16 int64 = 914
                        var v17 string = "a"
                        fmt.Println(v16, v4)
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    fmt.Println(v2)
                    _ = fi14
                }
                if false {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    switch (v1 + "a") {
                    case "a":
                        for fi18 := 0; fi18 < 15; fi18++ {
                            if false {
                                v19 := GenS1{F0: 457, F1: 0, F2: 2}
                                var v20 uint16 = 65535
                                _ = v19
                                _ = v20
                            }
                            v5 = append(v5, 0)
                            v4 = (((v4 * (1e10 + v4)) / (v4 - (v4 + float64(3.14159)))) * 1e10)
                            _ = fi18
                        }
                    case "\n\t":
                        if false {
                            if len(v5) > 0 {
                                v5[0] = 108
                            }
                            v4 = (v4 - (v4 - 0.1))
                        } else {
                            if true {
                                v4 += ((v4 * 0.1) * (1e10 * 1e10))
                                var v21 int32 = -1
                                v22 := GenS1{F0: 2147483648, F1: 65535, F2: 0}
                                _ = v21
                                _ = v22
                            } else {
                                v23 := GenS3{F0: (v10 << ((1) & 63)), F1: v7}
                                if true {
                                    var v24 uint64 = 18446744073709551615
                                    if (true || ((true && (false && false)) || false)) {
                                        v4 = v4
                                        var v25 uint = v10
                                        v26 := GenS1{F0: 42, F1: 32767, F2: 128}
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = v24
                                } else {
                                    var v27 uint8 = 127
                                    var v28 int64 = 1
                                    _ = v27
                                    _ = v28
                                }
                                _ = v23
                            }
                        }
                    }
                } else {
                    v7 -= v7
                    var v29 float32 = 0.0
                    v1 = v1
                    _ = v29
                }
            }
            _ = v10
        } else {
            var v30 uint8 = 127
            delete(v6, -362)
            _ = v30
        }
    }
    fmt.Println(v2, v4, v7)
    var v31 int64 = -501
    var v32 string = "hello"
    v33 := GenS4{F0: -1, F1: false, F2: 1}
    v34 := v6[2147483647]
    var v35 int16 = int16(v34)
    var v36 bool = false
    v37 := GenS2{F0: 2147483647, F1: 0}
    switch v1 {
    case "foo bar":
        v2 = v32
    default:
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
    }
    switch -812 {
    case 32768:
        fmt.Println(v2, v31)
    case 32767:
        switch 32767 {
        case 271:
            var v38 float32 = -2.25
            _ = v38
        }
    }
    switch 4294967295 {
    case 2147483647:
        var v39 string = "\n\t"
        _ = v39
    case 42:
        fmt.Println(v34, v36, v31)
    }
    var v40 uint = 0
    if (false && ((42 < -2147483648) && ("hello" > ("" + ("foo bar" + (v32 + (v1 + "foo bar"))))))) {
        if (3.14159 > -2.25) {
            var v41 string = v2
            v32 = v1
            if v36 {
                fmt.Println(v1, v40, v2)
                switch "\n\t" {
                case "a":
                    for fi42 := 0; fi42 < 1; fi42++ {
                        var v43 int16 = v7
                        if v36 {
                            v43 += v7
                        }
                        var v44 float32 = 0.1
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                case "":
                    var v45 uint8 = 175
                    _ = v45
                case "foo bar":
                    for fi46 := 0; fi46 < 5; fi46++ {
                        var v47 int8 = 45
                        var v48 int64 = -1
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                }
                v41 = v32
            } else {
                if len(v5) > 0 {
                    v5[0] = -2147483648
                }
                v1 = ""
            }
            _ = v41
        }
        _ = len(v5)
    } else {
        if (1 != -9223372036854775808) {
            v49 := GenS5{F0: 0, F1: ("\n\t" + v32)}
            _ = v49
        } else {
            v50 := GenS6{F0: v2, F1: 2}
            _ = v50
        }
        var v51 int32 = -2147483648
        for fi52 := 0; fi52 < 19; fi52++ {
            v53 := GenS7{F0: (v31 % ((v31) | 1)), F1: 479, F2: 189}
            _ = len(v6)
            _ = fi52
            _ = v53
        }
        _ = v51
    }
    switch v2 {
    case "hello":
        if v36 {
            var v54 string = (((("foo bar" + (v1 + v2)) + ((("foo bar" + (v2 + ("hello" + "foo bar"))) + ("\n\t" + ("" + ((v1 + "foo bar") + "hello")))) + ((v2 + "\n\t") + ("\n\t" + v1)))) + v2) + v2)
            var v55 uint = v40
            _ = v54
            _ = v55
        } else {
            v4 = ((float64(-2.25) * v4) * (3.14159 * v4))
            for fi56 := 0; fi56 < 3; fi56++ {
                switch (fi56 & (^(fi56))) {
                case 32768:
                    if (v35 < 1) {
                        fmt.Println(v4)
                    } else {
                        if (v36 || (127 >= 1)) {
                            v57 := GenS8{F0: (v1 + v1)}
                            switch 127 {
                            case -9223372036854775808:
                                var v58 int = (-(fi56))
                                _ = v58
                            }
                            _ = v57
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v59 uint32 = (v34 ^ v34)
                            _ = v59
                        }
                        v60 := v6[-1]
                        v2 = v32
                        _ = v60
                    }
                default:
                    var v61 float64 = v4
                    _ = v61
                }
                var v62 float64 = (((1e10 * v4) - v4) / (v4 / v4))
                _ = fi56
                _ = v62
            }
        }
    case "a":
        var v63 uint = (v40 * v40)
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
