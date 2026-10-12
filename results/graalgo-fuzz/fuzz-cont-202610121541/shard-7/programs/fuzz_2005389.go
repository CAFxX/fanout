package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int16; F2 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16; F1 uint16; F2 bool }
type GenS10 struct { F0 uint64; F1 bool }

func main() {
    var v1 string = (("a" + (("\n\t" + "") + ("\n\t" + "a"))) + "hello")
    v2 := []uint32{4294967295}
    v3 := []int{-2147483648, 32768, -1}
    var v4 int16 = 0
    v5 := GenS1{F0: 18446744073709551615, F1: 550}
    for fi6 := 0; fi6 < 1; fi6++ {
        v4 -= 0
        _ = fi6
    }
    v7 := GenS1{F0: 2, F1: 0}
    for fi8 := 0; fi8 < 8; fi8++ {
        if true {
            fmt.Println(fi8)
            fmt.Println(v4, v1)
            v4 -= (v4 / ((v4) | 1))
        } else {
            v1 = "hello"
        }
        var v9 int = fi8
        switch (v1 + "hello") {
        case "hello":
            v10 := GenS2{F0: (((127 == 100) || true) || (!(!true)))}
            _ = v10
        case "\n\t":
            v9 = v9
        case "a":
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 string = v1
                v13 := GenS3{F0: 127}
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            for fi14 := 0; fi14 < 9; fi14++ {
                if len(v3) > 0 {
                    v3[0] = fi14
                }
                v4 = -52
                var v15 bool = false
                v16 := GenS2{F0: (!(394 == 2147483648))}
                if (4294967295 != 779) {
                    if ("\n\t" <= ("hello" + v1)) {
                        var v17 int32 = 1
                        _ = v17
                    } else {
                        fmt.Println(v1, fi8, v9)
                    }
                    v18 := GenS4{F0: v1, F1: v4, F2: -2147483648}
                    _ = v18
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        _ = fi8
        _ = v9
    }
    v1 = ((v1 + "foo bar") + "hello")
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint16 = 1
        for fi21 := 0; fi21 < 20; fi21++ {
            switch v1 {
            case "\n\t":
                fmt.Println(fi21, v4)
            case "a":
                v1 = (((v1 + v1) + (v1 + v1)) + v1)
            default:
                if false {
                    v4 = v4
                    if false {
                        switch "" {
                        case "foo bar":
                            v22 := GenS4{F0: "a", F1: 1000, F2: -1}
                            _ = v22
                        case "a":
                            var v23 int64 = -1
                            v24 := GenS1{F0: 18446744073709551615, F1: 941}
                            _ = v23
                            _ = v24
                        case "":
                            var v25 float64 = 0.1
                            _ = v25
                        default:
                            v26 := GenS1{F0: 18446744073709551615, F1: 2147483647}
                            _ = v26
                        }
                        var v27 uint16 = v20
                        _ = v27
                    } else {
                        v28 := GenS5{F0: 127}
                        if true {
                            switch fi21 {
                            case 0:
                                v1 = "\n\t"
                            case 128:
                                v29 := GenS6{F0: -868}
                                _ = v29
                            case -543:
                                var v30 uint16 = (^(v20))
                                _ = v30
                            default:
                                if false {
                                    fmt.Println(v20, fi21)
                                }
                            }
                        }
                        var v31 uint16 = (^(v20))
                        _ = v28
                        _ = v31
                    }
                    v32 := GenS7{F0: (18446744073709551615 == 830)}
                    fmt.Println(v20)
                    _ = v32
                } else {
                    fmt.Println(v1)
                    var v33 uint16 = (v20 % (((v20 % (((v20 >> ((42) & 63))) | 1))) | 1))
                    var v34 float64 = 1.5
                    _ = v33
                    _ = v34
                }
            }
            if (-1 == -1) {
                if (3.14159 < 1.5) {
                    var v35 uint16 = (v20 + 32767)
                    fmt.Println(fi21)
                    v35 -= (v35 ^ v35)
                    _ = v35
                } else {
                    v36 := GenS8{F0: 42}
                    _ = v36
                }
                var v37 int = fi19
                v37 -= v37
                _ = v37
            }
            v38 := GenS9{F0: -32768, F1: (v20 - v20), F2: false}
            _ = fi21
            _ = v38
        }
        for fi39 := 0; fi39 < 8; fi39++ {
            fmt.Println(fi39)
            var v40 bool = (!(1.5 > float32(1e10)))
            var v41 int = -1
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = fi19
        _ = v20
    }
    for fi42 := 0; fi42 < 12; fi42++ {
        var v43 int = (^(fi42))
        _ = fi42
        _ = v43
    }
    v44 := GenS10{F0: 2, F1: ((0 <= 127) || false)}
    if true {
        v1 = (v1 + ("a" + v1))
    } else {
        if (v1 <= v1) {
            if true {
                var v45 uint = 193
                _ = v45
            }
            var v46 float64 = (float64(1.5) - float64(0.1))
            _ = v46
        }
    }
    if ((32767 > 0) && (1 < 986)) {
        var v47 int16 = v4
        _ = v47
    } else {
        for fi48 := 0; fi48 < 1; fi48++ {
            var v49 int32 = -1
            v50 := GenS6{F0: 1}
            _ = fi48
            _ = v49
            _ = v50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v7
}
