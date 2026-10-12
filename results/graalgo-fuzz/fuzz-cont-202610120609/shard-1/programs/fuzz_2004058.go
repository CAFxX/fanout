package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint32; F1 int64; F2 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 string; F2 uint8 }
type GenS8 struct { F0 uint16 }
func genFn0() string {
    var v6 bool = ((!false) || ((0 >= 42) || true))
    switch -1 {
    case 4294967295:
        var v7 int8 = 0
        _ = v7
    case 127:
        switch "\n\t" {
        case "foo bar":
            v6 = false
        case "\n\t":
            v8 := GenS1{F0: 42, F1: -52}
            _ = v8
        default:
            if (-128 == -1) {
                v9 := []int{}
                _ = v9
            } else {
                for fi10 := 0; fi10 < 4; fi10++ {
                    v6 = v6
                    switch (fi10 * (fi10 * 128)) {
                    case -383:
                        var v11 int = 32767
                        _ = v11
                    case 1:
                        var v12 bool = v6
                        _ = v12
                    }
                    fmt.Println(v6)
                    _ = fi10
                }
                v13 := []int{880}
                _ = v13
            }
        }
    case 32767:
        v14 := GenS1{F0: 2, F1: 1}
        _ = v14
    }
    fmt.Println(v6)
    if false {
        v6 = false
        var v15 int16 = -32768
        _ = v15
    }
    _ = v6
    return "\n\t"
}

func main() {
    var v1 uint32 = 2147483648
    var v2 bool = false
    var v3 string = (("a" + "") + ("foo bar" + "\n\t"))
    var v4 int8 = -128
    var v5 float32 = 0.0
    v16 := GenS2{F0: (!false)}
    if true {
        var v17 float32 = 0.0
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 uint64 = 238
            _ = fi18
            _ = v19
        }
        var v20 uint32 = (v1 % ((v1) | 1))
        _ = v17
        _ = v20
    }
    var v21 uint16 = 1
    switch 77 {
    case 32768:
        v22 := []int{276, 1}
        _ = v22
    case 30:
        var v23 uint16 = (v21 / ((v21) | 1))
        _ = v23
    case -356:
        v24 := []int{127, -291}
        _ = v24
    default:
        if (((v4 - (v4 * (^(v4)))) >= v4) || true) {
            v25 := GenS3{F0: 0}
            _ = v25
        } else {
            v26 := []int{}
            for fi27 := 0; fi27 < 12; fi27++ {
                for fi28 := 0; fi28 < 5; fi28++ {
                    v3 = v3
                    _ = fi28
                }
                for fi29 := 0; fi29 < 7; fi29++ {
                    v4 = (^(v4))
                    if v2 {
                        switch v3 {
                        case "hello":
                            v5 += v5
                        case "\n\t":
                            v1 -= (v1 / ((v1) | 1))
                        default:
                            if len(v26) > 0 {
                                v26[0] = fi27
                            }
                        }
                    } else {
                        v2 = true
                        _ = len(v26)
                        v5 = (0.1 + v5)
                    }
                    v21 += v21
                    _ = fi29
                }
                v2 = true
                _ = fi27
            }
            _ = v26
        }
    }
    fmt.Println(v2, v2)
    if v2 {
        var v30 bool = (true || v2)
        v31 := GenS4{F0: (^(v21)), F1: 1}
        _ = v30
        _ = v31
    } else {
        switch "foo bar" {
        case "\n\t":
            switch v3 {
            case "hello":
                var v32 uint64 = uint64(v21)
                _ = v32
            case "":
                if ((v3 + v3) <= v3) {
                    v33 := GenS3{F0: 448}
                    _ = v33
                } else {
                    v34 := GenS5{F0: 0, F1: 0, F2: 9223372036854775807}
                    if v2 {
                        switch "hello" {
                        case "hello":
                            v35 := []int{}
                            _ = v35
                        case "\n\t":
                            v36 := GenS6{F0: v2}
                            _ = v36
                        case "":
                            v4 += v4
                        default:
                            v37 := []int{32767, 9223372036854775807, 9223372036854775807}
                            _ = v37
                        }
                        switch 9223372036854775807 {
                        case -416:
                            if (!v2) {
                                fmt.Println(v2, v3, v3)
                                v4 -= 100
                            } else {
                                var v38 int16 = -32768
                                _ = v38
                            }
                        case -635:
                            v5 = 1e10
                        case -537:
                            var v39 int64 = 42
                            _ = v39
                        }
                    } else {
                        v5 -= 1.5
                        var v40 bool = (!true)
                        _ = v40
                    }
                    _ = v34
                }
            case "a":
                if (v5 == (0.0 + (3.14159 * ((1.5 / (0.1 / v5)) + v5)))) {
                    if v2 {
                        if (128 < 9223372036854775807) {
                            v3 = ("hello" + v3)
                            fmt.Println(v21, v3, v1)
                            v41 := GenS4{F0: (v21 | (^(v21))), F1: 249}
                            _ = v41
                        }
                        v42 := GenS6{F0: ((65535 != uint(v21)) || v2)}
                        v21 = (v21 >> ((255) & 63))
                        _ = v42
                    }
                    fmt.Println(v2)
                    v3 = (v3 + "a")
                } else {
                    v21 -= 1
                }
            }
        case "hello":
            v3 = "\n\t"
        default:
            switch v3 {
            case "hello":
                var v43 string = v3
                _ = v43
            }
        }
        v5 = (((float32(-2.25) + 1e10) * 0.0) - (v5 / 1.0))
    }
    var v44 int = 0
    v45 := GenS5{F0: 768, F1: 42, F2: 0}
    v1 = v1
    if v2 {
        v46 := GenS4{F0: v21, F1: 4294967295}
        _ = v46
    } else {
        fmt.Println(v1, v1, v1)
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        v48 := GenS4{F0: (^(v21)), F1: 0}
        if true {
            var v49 uint8 = 141
            v21 -= (v21 << ((18446744073709551615) & 63))
            _ = v49
        }
        _ = fi47
        _ = v48
    }
    for fi50 := 0; fi50 < 17; fi50++ {
        v51 := []int{int(v4), v44, (v44 * v44), v44}
        var v52 int16 = 1000
        _ = fi50
        _ = v51
        _ = v52
    }
    v53 := GenS7{F0: v44, F1: (((("hello" + v3) + "a") + ("foo bar" + ("\n\t" + v3))) + "foo bar"), F2: uint8(v1)}
    v54 := []int{(v44 / ((v44) | 1)), 2147483647}
    v4 = -93
    v55 := GenS8{F0: v21}
    if len(v54) >= 2 {
        v54 = v54[1:len(v54)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v54:", len(v54), v54)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v53
    _ = v54
    _ = v55
}
