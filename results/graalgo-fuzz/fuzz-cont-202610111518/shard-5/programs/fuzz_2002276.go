package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int16; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint16, p1 int, p2 int) int8 {
    if (1 > 1) {
        fmt.Println(p0)
        var v7 uint32 = 929
        _ = v7
    }
    v8 := GenS1{F0: -89, F1: 2147483648, F2: 449}
    v9 := GenS1{F0: 127, F1: 2147483648, F2: 691}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -79
}

func main() {
    var v1 float64 = 0.0
    var v2 uint8 = 127
    var v3 float32 = 1.5
    var v4 float32 = (1.5 + 1e10)
    var v5 string = "hello"
    v6 := []uint32{}
    var v10 uint = 42
    var v11 bool = true
    fmt.Println(v4, v11, v10)
    fmt.Println(v3)
    var v12 int64 = -120
    switch "foo bar" {
    case "":
        fmt.Println(v2, v12)
    default:
        switch 42 {
        case 0:
            var v13 uint32 = 369
            _ = v13
        default:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    }
    if (63 < int8(v12)) {
        v12 = v12
        v1 += 0.1
        v10 = (^(v10))
    } else {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 float64 = v1
            fmt.Println(v4, v12)
            _ = fi14
            _ = v15
        }
        v6 = append(v6, 941)
        if v11 {
            v2 = v2
        } else {
            _ = len(v6)
        }
    }
    var v16 float64 = v1
    for fi17 := 0; fi17 < 14; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 float32 = (v3 * -2.25)
            var v20 int16 = 1000
            var v21 int = 127
            if len(v6) > 0 {
                v6[0] = 1
            }
            if v11 {
                v20 += 1
                v6 = append(v6, 42)
                if v11 {
                    fmt.Println(v5, v3)
                    v22 := GenS2{F0: 0, F1: 100, F2: 184}
                    fmt.Println(v5)
                    if true {
                        if len(v6) > 0 {
                            v6[0] = 156
                        }
                        var v23 int8 = genFn0(0, (fi18 + 365), (-(fi17)))
                        switch (("\n\t" + ((v5 + "\n\t") + ("hello" + (v5 + v5)))) + v5) {
                        case "\n\t":
                            var v24 uint8 = (^(v2))
                            _ = v24
                        case "":
                            v16 -= 3.14159
                            v6 = append(v6, 0)
                            v25 := GenS2{F0: (v20 + v20), F1: v23, F2: 2}
                            v26 := GenS2{F0: (v20 >> ((850) & 63)), F1: (v23 - v23), F2: 18446744073709551615}
                            _ = v25
                            _ = v26
                        case "hello":
                            fmt.Println(v16, v5)
                        }
                        _ = v23
                    }
                    _ = v22
                }
            }
            if (v5 >= v5) {
                v27 := GenS3{F0: v10}
                switch v5 {
                case "foo bar":
                    var v28 uint = v10
                    _ = v28
                default:
                    fmt.Println(v10, v10, v20)
                }
                switch "a" {
                case "\n\t":
                    var v29 int32 = 674
                    _ = v29
                case "":
                    switch v5 {
                    case "":
                        var v30 uint16 = 32767
                        _ = v30
                    case "foo bar":
                        if false {
                            v12 += int64(v2)
                            v11 = ((v2 | v2) < 1)
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if (v5 <= "\n\t") {
                                v6 = append(v6, 335)
                            } else {
                                var v31 int32 = -590
                                _ = v31
                            }
                            v6 = append(v6, 65535)
                            var v32 uint = (v10 % ((18446744073709551615) | 1))
                            _ = v32
                        }
                    }
                case "hello":
                    if true {
                        if v11 {
                            _ = len(v6)
                            if false {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v33 int32 = 42
                                v20 = (v20 * v20)
                                _ = v33
                            } else {
                                v20 = (v20 + v20)
                                var v34 uint = v10
                                v2 -= (^(v2))
                                _ = v34
                            }
                            var v35 uint32 = uint32(v10)
                            _ = v35
                        }
                        var v36 uint8 = v2
                        if len(v6) > 0 {
                            v6[0] = 346
                        }
                        _ = v36
                    }
                }
                v37 := GenS4{F0: 18446744073709551615, F1: 4294967295}
                fmt.Println(v16, fi17)
                var v38 int32 = 1
                var v39 float64 = 0.1
                _ = v27
                _ = v37
                _ = v38
                _ = v39
            }
            var v40 int8 = 100
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
            _ = v40
        }
        _ = fi17
    }
    if (!v11) {
        var v41 float64 = float64((v4 - (v4 + 3.14159)))
        _ = v41
    } else {
        v1 = (v1 / (v1 - 0.1))
    }
    var v42 uint8 = v2
    v43 := GenS4{F0: 395, F1: 559}
    if v11 {
        var v44 int16 = -623
        var v45 float32 = float32(float64(3.14159))
        var v46 string = v5
        _ = v44
        _ = v45
        _ = v46
    } else {
        fmt.Println(v16, v2)
        var v47 uint = 2
        v6 = append(v6, 861)
        _ = v47
    }
    switch v5 {
    case "\n\t":
        v16 += float64(float32(v1))
    case "a":
        if len(v6) > 0 {
            v6[0] = 820
        }
    default:
        var v48 float64 = (0.1 + v16)
        _ = v48
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
