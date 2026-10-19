package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int16; F1 int32; F2 uint64 }
func genFn0(p0 uint, p1 int16) uint16 {
    var v6 float32 = float32(1e10)
    v7 := []int{128, 2147483647, 595, 127}
    v8 := GenS1{F0: -280}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int16, p1 int, p2 uint64) uint16 {
    var v9 int32 = -940
    var v10 uint8 = 42
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS1{F0: v9}
        switch p1 {
        case 2147483647:
            v13 := GenS1{F0: v9}
            _ = v13
        case 127:
            v14 := GenS1{F0: (v9 | (v9 / ((v9) | 1)))}
            _ = v14
        default:
            p2 = p2
        }
        _ = fi11
        _ = v12
    }
    v15 := GenS1{F0: (v9 * 528)}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v15
    _ = v9
    return genFn0(65535, p0)
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int = 0
    var v3 uint = 324
    v4 := []uint32{42, 2147483647, 0}
    var v5 string = "hello"
    fmt.Println(v1, v1, v3)
    var v16 uint32 = 65535
    v5 = "hello"
    var v17 int64 = 4
    var v18 uint32 = v16
    for fi19 := 0; fi19 < 5; fi19++ {
        fmt.Println(v3, v16, v18)
        if true {
            fmt.Println(v18, v16)
            var v20 uint = (v3 >> (((v3 - v3)) & 63))
            _ = v20
        } else {
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 float32 = 0.1
                if true {
                    v18 = (^(v16))
                    v22 -= (float32((1.5 + float64(-2.25))) + 1e10)
                } else {
                    fmt.Println(v1)
                    v23 := GenS1{F0: 2147483647}
                    if true {
                        v1 = (v1 + v17)
                        if (v5 > v5) {
                            v3 -= 65535
                            fmt.Println(v17, v22)
                            v17 = v17
                        }
                    }
                    switch v5 {
                    case "":
                        switch "a" {
                        case "\n\t":
                            if (((v3 == v3) && false) || (42 >= 515)) {
                                if ((v3 << ((v3) & 63)) >= (v3 * (v3 * (v3 / (((^(v3))) | 1))))) {
                                    v5 = ("foo bar" + v5)
                                    v24 := GenS2{F0: false, F1: 0, F2: genFn0(v3, 1)}
                                    v25 := GenS1{F0: 0}
                                    _ = v24
                                    _ = v25
                                }
                                v4 = append(v4, v18)
                            }
                        case "":
                            v26 := GenS3{F0: 42, F1: 161}
                            _ = v26
                        case "a":
                            if (3.14159 <= (v22 + v22)) {
                                v27 := GenS3{F0: 255, F1: uint8(v3)}
                                if true {
                                    v28 := GenS3{F0: 1, F1: 99}
                                    v1 = (v1 | 42)
                                    _ = v28
                                }
                                _ = v27
                            }
                        default:
                            fmt.Println(v17, v3)
                        }
                    case "a":
                        var v29 int32 = -820
                        _ = v29
                    case "foo bar":
                        v30 := GenS2{F0: true, F1: 442, F2: 1}
                        _ = v30
                    default:
                        if false {
                            var v31 float64 = (3.14159 * -2.25)
                            _ = v31
                        }
                    }
                    _ = v23
                }
                var v32 float64 = 0.0
                _ = fi21
                _ = v22
                _ = v32
            }
        }
        _ = fi19
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        fmt.Println(v1)
        _ = fi33
    }
    if (false && true) {
        switch v5 {
        case "hello":
            v34 := GenS4{F0: int16(v17), F1: 2147483647, F2: 877}
            _ = v34
        case "foo bar":
            var v35 bool = (v18 != 2)
            _ = v35
        case "":
            for fi36 := 0; fi36 < 15; fi36++ {
                var v37 uint32 = (v16 / ((65535) | 1))
                if (true && true) {
                    var v38 int16 = 380
                    var v39 int8 = -5
                    if (-2.25 <= float64(((((0.1 * 1e10) + 0.0) + 0.0) * 0.1))) {
                        v16 = (^(v16))
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch 32768 {
                        case 986:
                            v40 := GenS1{F0: 0}
                            _ = v40
                        case -935:
                            fmt.Println(v5)
                        case 128:
                            var v41 uint32 = (v16 + v16)
                            _ = v41
                        }
                        for fi42 := 0; fi42 < 6; fi42++ {
                            var v43 uint32 = (v16 + v16)
                            var v44 uint64 = 746
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                    }
                    _ = v38
                    _ = v39
                } else {
                    v45 := GenS1{F0: 124}
                    _ = v45
                }
                if false {
                    for fi46 := 0; fi46 < 1; fi46++ {
                        var v47 uint32 = (v18 & v18)
                        var v48 uint16 = 0
                        fmt.Println(fi46, v16)
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    v5 = (v5 + ((v5 + v5) + ("foo bar" + "hello")))
                    var v49 bool = (false || true)
                    _ = v49
                }
                _ = fi36
                _ = v37
            }
        }
        fmt.Println(v17, v2, v3)
        var v50 int8 = -49
        _ = v50
    } else {
        switch v5 {
        case "":
            v16 += v18
        case "hello":
            for fi51 := 0; fi51 < 8; fi51++ {
                var v52 bool = true
                v17 = v1
                _ = fi51
                _ = v52
            }
        case "a":
            for fi53 := 0; fi53 < 18; fi53++ {
                v54 := GenS2{F0: (!(65535 <= 643)), F1: 65535, F2: 277}
                if true {
                    var v55 uint64 = 1
                    _ = v55
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi53
                _ = v54
            }
        }
        for fi56 := 0; fi56 < 14; fi56++ {
            var v57 uint32 = (^(v16))
            switch (v2 | 376) {
            case 229:
                v16 += v18
            default:
                var v58 uint8 = 187
                _ = v58
            }
            _ = fi56
            _ = v57
        }
        var v59 uint8 = 255
        _ = v59
    }
    var v60 float32 = (3.14159 / ((0.0 * (3.14159 - 1e10)) + float32(1e10)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
}
