package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint; F2 uint32 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 uint8; F1 int16; F2 int16 }

func main() {
    var v1 uint32 = 2147483647
    var v2 int8 = int8(v1)
    var v3 float64 = 1.5
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]bool)
    if true {
        v1 = v1
        if (!("\n\t" == "hello")) {
            if (("\n\t" + "a") <= "") {
                var v6 float64 = v3
                if true {
                    var v7 string = ("" + ("a" + ("hello" + "")))
                    var v8 int64 = 9223372036854775807
                    if (((-2147483648 != 342) || (-1 < v8)) && false) {
                        _ = len(v4)
                    }
                    _ = v7
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 9; fi9++ {
                        fmt.Println(v6)
                        _ = fi9
                    }
                    v10 := GenS1{F0: "hello", F1: -32768}
                    if (0 <= 32767) {
                        for fi11 := 0; fi11 < 18; fi11++ {
                            if (false || true) {
                                var v12 string = "foo bar"
                                _ = v12
                            } else {
                                var v13 int = 0
                                if false {
                                    for fi14 := 0; fi14 < 18; fi14++ {
                                        var v15 uint = 856
                                        v13 += 127
                                        _ = fi14
                                        _ = v15
                                    }
                                    switch 2147483647 {
                                    case 9223372036854775807:
                                        v1 = v1
                                    case 1:
                                        v4[v1] = true
                                    case 4294967295:
                                        fmt.Println(v13, v6)
                                    default:
                                        if (false || true) {
                                            v13 = -1
                                        }
                                    }
                                    if ((65535 != (v1 - v1)) || true) {
                                        var v16 int32 = -2147483648
                                        v6 = float64(0.0)
                                        _ = v16
                                    } else {
                                        if (true && false) {
                                            switch 955 {
                                            case 42:
                                                v6 = float64(0.0)
                                            case 1:
                                                if (false && true) {
                                                    var v17 int32 = 0
                                                    v18 := GenS1{F0: "", F1: -1}
                                                    var v19 uint64 = 2
                                                    _ = v17
                                                    _ = v18
                                                    _ = v19
                                                } else {
                                                    fmt.Println(v1)
                                                    if (!false) {
                                                        fmt.Println(v2, v6, v13)
                                                        var v20 string = (("hello" + "\n\t") + "hello")
                                                        _ = v20
                                                    } else {
                                                        if ((18446744073709551615 <= 396) && false) {
                                                            if false {
                                                                v6 += v3
                                                            }
                                                            if ((^(fi11)) != (^(v13))) {
                                                                switch "hello" {
                                                                case "":
                                                                    var v21 uint16 = 1
                                                                    _ = v21
                                                                case "a":
                                                                    var v22 uint64 = 18446744073709551615
                                                                    _ = v22
                                                                }
                                                                v3 -= (0.1 + (float64(-2.25) + 0.0))
                                                                v23 := []int{}
                                                                _ = v23
                                                            }
                                                        } else {
                                                            v3 = (float64(1.5) / (v3 - v6))
                                                        }
                                                    }
                                                    fmt.Println(v3, v3, v1)
                                                }
                                            case -869:
                                                if (v13 != (v13 + (v13 * (v13 * (v13 % ((v13) | 1)))))) {
                                                    var v24 uint32 = (v1 | v1)
                                                    v25 := GenS1{F0: "", F1: 32767}
                                                    _ = v24
                                                    _ = v25
                                                }
                                            }
                                        }
                                        var v26 uint8 = 127
                                        _ = v26
                                    }
                                } else {
                                    v27 := GenS1{F0: "a", F1: 42}
                                    v4[v1] = true
                                    _ = v27
                                }
                                v28 := GenS1{F0: "hello", F1: 0}
                                _ = v13
                                _ = v28
                            }
                            v29 := GenS2{F0: 2, F1: (fi11 << ((558) & 63))}
                            if ((88 != (v1 | 2147483648)) || false) {
                                v1 += v1
                            } else {
                                v30 := GenS3{F0: "foo bar"}
                                _ = v30
                            }
                            _ = fi11
                            _ = v29
                        }
                    }
                    _ = v10
                }
                for fi31 := 0; fi31 < 10; fi31++ {
                    fmt.Println(v2, v3, v6)
                    switch "hello" {
                    case "a":
                        v32 := GenS4{F0: -9223372036854775808, F1: 71, F2: 468}
                        _ = v32
                    case "foo bar":
                        v33 := GenS5{F0: ("hello" + "\n\t"), F1: (-(fi31))}
                        _ = v33
                    case "":
                        var v34 uint8 = 128
                        _ = v34
                    default:
                        var v35 int8 = v2
                        _ = v35
                    }
                    var v36 uint = 1
                    _ = fi31
                    _ = v36
                }
                _ = v6
            }
        } else {
            for fi37 := 0; fi37 < 19; fi37++ {
                v38 := GenS4{F0: 0, F1: 1, F2: v1}
                var v39 float32 = ((1.5 / 3.14159) - -2.25)
                var v40 string = ("hello" + ("\n\t" + ""))
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
            if true {
                v41 := GenS5{F0: "hello", F1: 900}
                var v42 uint16 = 42
                fmt.Println(v1)
                _ = v41
                _ = v42
            }
        }
    } else {
        v43 := GenS6{F0: 255, F1: 0}
        v44 := []int{-710, 127, 32767}
        _ = v43
        _ = v44
    }
    fmt.Println(v2, v1)
    if (633 >= 871) {
        v1 = 329
        v3 -= -2.25
        if (32767 >= 1) {
            var v45 float64 = v3
            fmt.Println(v2, v2)
            _ = v45
        } else {
            v1 -= 535
            for fi46 := 0; fi46 < 14; fi46++ {
                for fi47 := 0; fi47 < 14; fi47++ {
                    v1 = v1
                    v48 := v4[v1]
                    _ = fi47
                    _ = v48
                }
                _ = fi46
            }
        }
    } else {
        if (!(!true)) {
            var v49 uint = 71
            var v50 uint8 = 255
            _ = v49
            _ = v50
        }
    }
    var v51 int64 = -130
    if true {
        fmt.Println(v3)
    } else {
        v52 := GenS4{F0: -929, F1: 0, F2: v1}
        v53 := GenS1{F0: "foo bar", F1: -32768}
        v2 -= (v2 >> ((909) & 63))
        _ = v52
        _ = v53
    }
    switch 571 {
    case 32768:
        v54 := GenS7{F0: 124, F1: 1, F2: -1}
        _ = v54
    case 128:
        for fi55 := 0; fi55 < 1; fi55++ {
            v2 += 100
            _ = fi55
        }
    }
    var v56 int8 = v2
    var v57 int32 = 573
    v58 := []int{-2147483648}
    v57 -= v57
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v56
    _ = v57
    _ = v58
}
