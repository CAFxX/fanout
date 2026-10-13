package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 string; F1 uint; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint; F2 uint }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]bool)
    v3 := []string{(v1 + v1), (v1 + v1)}
    v4 := make(map[int]bool)
    var v5 int8 = 0
    v6 := GenS1{F0: 0}
    fmt.Println(v1, v1)
    if (724 > 65535) {
        v5 = -59
        fmt.Println(v1)
        v1 = ((v1 + "foo bar") + v1)
    } else {
        fmt.Println(v5, v1)
    }
    fmt.Println(v1)
    _ = len(v3)
    var v7 uint32 = 42
    var v8 uint8 = 168
    if (v8 >= v8) {
        if len(v3) > 0 {
            v3[0] = "a"
        }
    } else {
        if (100 != v5) {
            if (!(!false)) {
                var v9 float64 = 1.5
                v10 := GenS1{F0: 0}
                _ = v10
                _ = v9
            }
        } else {
            v5 = v5
        }
        switch 86 {
        case 128:
            fmt.Println(v7, v1, v5)
        default:
            if true {
                v7 = v7
                v11 := GenS1{F0: 1}
                var v12 uint = 4294967295
                _ = v11
                _ = v12
            } else {
                var v13 float64 = 1.5
                for fi14 := 0; fi14 < 4; fi14++ {
                    v13 += -2.25
                    v5 = (v5 >> ((639) & 63))
                    switch -557 {
                    case 128:
                        v15 := GenS1{F0: 1}
                        _ = v15
                    default:
                        v7 = v7
                    }
                    _ = fi14
                }
                var v16 uint8 = v8
                _ = v13
                _ = v16
            }
        }
        fmt.Println(v5, v7)
    }
    var v17 int16 = 1
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 bool = false
        if false {
            var v20 string = ""
            if v19 {
                var v21 int = 0
                for fi22 := 0; fi22 < 9; fi22++ {
                    switch "foo bar" {
                    case "hello":
                        v23 := GenS1{F0: 323}
                        _ = v23
                    case "":
                        fmt.Println(fi22)
                    case "foo bar":
                        v24 := GenS1{F0: 1}
                        _ = v24
                    default:
                        v21 = (fi18 * v21)
                    }
                    _ = fi22
                }
                _ = v21
            }
            if v19 {
                if false {
                    if (true || ((v5 != v5) || (v7 != (v7 * (^(v7)))))) {
                        for fi25 := 0; fi25 < 3; fi25++ {
                            if (v17 <= v17) {
                                switch fi18 {
                                case 32767:
                                    v19 = (v1 == v20)
                                }
                                if ((v17 / ((v17) | 1)) >= 1) {
                                    if (v19 || true) {
                                        v5 = 1
                                    } else {
                                        var v26 bool = v19
                                        _ = v26
                                    }
                                    if (1 > 309) {
                                        if v19 {
                                            var v27 uint = 18446744073709551615
                                            if (fi18 != fi18) {
                                                var v28 string = "a"
                                                _ = v28
                                            } else {
                                                fmt.Println(fi25, v17, fi18)
                                            }
                                            v29 := GenS1{F0: 0}
                                            _ = len(v2)
                                            fmt.Println(v7, v5)
                                            _ = v27
                                            _ = v29
                                        }
                                    }
                                    var v30 int8 = 127
                                    _ = v30
                                } else {
                                    v17 = v17
                                }
                            }
                            var v31 int8 = (v5 >> ((455) & 63))
                            _ = fi25
                            _ = v31
                        }
                        var v32 float64 = 0.0
                        v20 = v1
                        _ = v32
                    } else {
                        fmt.Println(v17)
                    }
                    var v33 int = -2147483648
                    _ = v33
                }
            } else {
                var v34 int = fi18
                _ = v34
            }
            _ = v20
        } else {
            var v35 uint8 = 0
            fmt.Println(v17)
            _ = v35
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v1, v1, v1)
    for fi36 := 0; fi36 < 5; fi36++ {
        v8 -= v8
        if (!true) {
            v7 = v7
        } else {
            var v37 float32 = 0.0
            _ = v37
        }
        _ = fi36
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v38 int8 = 100
    for fi39 := 0; fi39 < 4; fi39++ {
        if ((false || true) && true) {
            v8 += v8
            switch v1 {
            case "hello":
                v40 := GenS2{F0: 42, F1: v7}
                _ = v40
            case "foo bar":
                fmt.Println(v38, v5)
            }
            var v41 float32 = (((0.0 * (1.5 + ((1e10 / 0.1) + (float32(0.1) - (1.5 - -2.25))))) + (-2.25 - 0.1)) - 1e10)
            _ = v41
        } else {
            for fi42 := 0; fi42 < 10; fi42++ {
                v43 := GenS1{F0: 42}
                v5 -= (v38 - v38)
                switch (v1 + ("\n\t" + (v1 + "foo bar"))) {
                case "":
                    if (false || true) {
                        if false {
                            v7 -= v7
                            v38 += (v38 % ((0) | 1))
                        } else {
                            var v44 int32 = 1
                            v8 += v8
                            _ = v44
                        }
                        v45 := GenS3{F0: ("" + v1), F1: 633, F2: 255}
                        _ = v45
                    }
                default:
                    switch (fi39 >> ((75) & 63)) {
                    case -375:
                        v46 := GenS4{F0: -914, F1: uint(v8), F2: 18446744073709551615}
                        _ = v46
                    case 128:
                        v17 -= (^(v17))
                    case 300:
                        var v47 uint = 255
                        _ = v47
                    default:
                        v17 += (v17 % ((v17) | 1))
                    }
                }
                _ = fi42
                _ = v43
            }
            v48 := GenS4{F0: -627, F1: uint(v7), F2: 18446744073709551615}
            var v49 int = fi39
            _ = v48
            _ = v49
        }
        _ = fi39
    }
    v38 = (v38 & (v38 + int8(v8)))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
