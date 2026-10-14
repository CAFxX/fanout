package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int16; F2 uint }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint; F2 int64 }
func genFn0(p0 int32, p1 uint32, p2 int32) uint32 {
    if true {
        var v7 float64 = float64(0.0)
        var v8 string = ""
        p1 -= (p1 - 2147483647)
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 134}
        _ = v9
    }
    p1 = p1
    v10 := GenS1{F0: 63}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"a", "foo bar"}
    var v3 bool = true
    var v4 int32 = 42
    var v5 float64 = (float64(1e10) * (((float64(3.14159) / (1.5 * ((1.5 / 1.5) - (1.5 - (0.0 + -2.25))))) / 1e10) + 3.14159))
    v6 := make(map[uint32]bool)
    switch 0 {
    case 128:
        v11 := GenS1{F0: 1}
        _ = v11
    case -2147483648:
        v12 := v6[genFn0((v4 + (v4 * 42)), 2147483647, -1)]
        _ = v12
    default:
        if v3 {
            for fi13 := 0; fi13 < 19; fi13++ {
                if (0 <= 983) {
                    v3 = true
                    v3 = v3
                } else {
                    v5 += 1e10
                    for fi14 := 0; fi14 < 11; fi14++ {
                        var v15 uint = 1
                        _ = fi14
                        _ = v15
                    }
                }
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 uint64 = 4294967295
                    var v18 uint16 = 2
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi13
            }
            var v19 int8 = -113
            _ = v19
        }
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    if (false && (genFn0(v4, 807, (v4 * -2147483648)) > 845)) {
        v20 := GenS1{F0: 42}
        _ = v20
    } else {
        v21 := v6[2147483648]
        delete(v6, 20)
        _ = v21
    }
    if ((("hello" + "") + ((("\n\t" + "hello") + ((("\n\t" + "foo bar") + "\n\t") + ("" + "foo bar"))) + "foo bar")) != "foo bar") {
        var v22 uint16 = 46
        if (!v3) {
            v3 = v3
        } else {
            v22 += v22
        }
        v23 := GenS2{F0: 42, F1: 0, F2: 978}
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: 100, F1: 51, F2: 18446744073709551615}
    v4 -= -439
    switch 32767 {
    case -58:
        v6[1] = v3
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := GenS3{F0: 390, F1: (^(v4))}
        if (2147483647 != fi25) {
            if true {
                var v27 uint32 = 65535
                fmt.Println(v3)
                fmt.Println(v27)
                _ = v27
            } else {
                if false {
                    fmt.Println(v4, v5)
                    v5 += (1e10 - v5)
                    if false {
                        var v28 uint32 = 1
                        v29 := GenS2{F0: 0, F1: 0, F2: 2}
                        v30 := GenS2{F0: 1, F1: 676, F2: 512}
                        _ = v28
                        _ = v29
                        _ = v30
                    } else {
                        if v3 {
                            for fi31 := 0; fi31 < 13; fi31++ {
                                v32 := GenS3{F0: 65535, F1: (-(v4))}
                                switch (fi31 % ((4294967295) | 1)) {
                                case 2147483647:
                                    if true {
                                        fmt.Println(v3)
                                        fmt.Println(v5, fi25, v3)
                                        if v3 {
                                            fmt.Println(fi25, v3)
                                            var v33 bool = (!(!false))
                                            _ = v33
                                        }
                                        var v34 int = -50
                                        v4 = 0
                                        _ = v34
                                    }
                                case 4294967295:
                                    if ((true || v3) && v3) {
                                        delete(v1, 4294967295)
                                        var v35 int64 = 174
                                        v36 := GenS4{F0: 392}
                                        _ = v35
                                        _ = v36
                                    } else {
                                        fmt.Println(v5)
                                        var v37 bool = (!v3)
                                        _ = v37
                                    }
                                default:
                                    var v38 uint32 = 2147483647
                                    _ = v38
                                }
                                _ = fi31
                                _ = v32
                            }
                            v39 := GenS5{F0: v3, F1: 65535, F2: 1}
                            _ = v39
                        } else {
                            v40 := GenS2{F0: 17, F1: -32768, F2: 65535}
                            switch (^(fi25)) {
                            case 2147483647:
                                fmt.Println(fi25)
                            case 647:
                                var v41 uint8 = 165
                                _ = v41
                            }
                            _ = v40
                        }
                        var v42 bool = true
                        var v43 uint64 = 4294967295
                        _ = v42
                        _ = v43
                    }
                }
                var v44 string = ((("a" + "a") + (("\n\t" + "a") + "")) + "a")
                _ = v44
            }
            if v3 {
                v4 -= v4
            }
        }
        _ = fi25
        _ = v26
    }
    v45 := GenS3{F0: 1, F1: (v4 * v4)}
    for fi46 := 0; fi46 < 12; fi46++ {
        var v47 float32 = (float32(v5) * 1e10)
        _ = fi46
        _ = v47
    }
    v48 := GenS4{F0: 0}
    switch 4294967295 {
    case 32768:
        v49 := GenS3{F0: 672, F1: (-(v4))}
        _ = v49
    default:
        fmt.Println(v4, v5, v4)
    }
    switch "a" {
    case "\n\t":
        for fi50 := 0; fi50 < 4; fi50++ {
            var v51 bool = (!(!v3))
            _ = fi50
            _ = v51
        }
    }
    v3 = ((v3 || (((1.5 > float32(v5)) || true) || v3)) && (true && false))
    v52 := GenS4{F0: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
    _ = v52
    _ = v6
}
