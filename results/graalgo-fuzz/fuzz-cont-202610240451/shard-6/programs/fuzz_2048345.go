package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 string; F1 uint8; F2 bool }
type GenS6 struct { F0 uint32; F1 uint; F2 int8 }
func genFn0() uint8 {
    var v7 int16 = 0
    if (!(160 <= 2)) {
        for fi8 := 0; fi8 < 11; fi8++ {
            v7 -= v7
            v7 = v7
            _ = fi8
        }
    }
    v7 += (v7 - 1)
    v7 = v7
    var v9 int = 4294967295
    _ = v7
    _ = v9
    return 42
}

func main() {
    var v1 int8 = -10
    var v2 uint = 65535
    v3 := make(map[string]bool)
    var v4 uint16 = 1
    v5 := make(map[string]string)
    var v6 int8 = (^(v1))
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := []int{9223372036854775807, (fi10 * -203)}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 18446744073709551615}
    var v13 int16 = -1
    if ((0.0 <= 0.1) || (1 < genFn0())) {
        if true {
            var v14 int16 = v13
            v1 = (v1 * 100)
            v15 := GenS1{F0: 1}
            _ = v14
            _ = v15
        }
        var v16 int = -1
        _ = v16
    }
    delete(v5, ("" + ("foo bar" + "\n\t")))
    v13 += v13
    for fi17 := 0; fi17 < 10; fi17++ {
        if false {
            switch 32767 {
            case 415:
                v18 := GenS2{F0: 449}
                _ = v18
            }
        } else {
            fmt.Println(v13, v13, v2)
            v19 := GenS1{F0: 18446744073709551615}
            if (32768 < (fi17 ^ -1)) {
                v20 := GenS2{F0: (fi17 + fi17)}
                fmt.Println(v4, fi17)
                for fi21 := 0; fi21 < 3; fi21++ {
                    v4 = v4
                    switch ("hello" + ((("a" + "foo bar") + ("" + "\n\t")) + "")) {
                    case "a":
                        fmt.Println(fi17)
                    case "":
                        var v22 uint = 0
                        _ = v22
                    }
                    _ = fi21
                }
                _ = v20
            } else {
                switch fi17 {
                case 127:
                    v23 := GenS3{F0: -1}
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 13; fi24++ {
                        var v25 bool = true
                        v6 = (v1 ^ v1)
                        _ = fi24
                        _ = v25
                    }
                }
                if true {
                    var v26 int16 = v13
                    _ = v26
                }
                v6 += v1
            }
            _ = v19
        }
        var v27 uint16 = v4
        v28 := GenS4{F0: (0.0 > 3.14159), F1: (v4 | (^(v4)))}
        _ = fi17
        _ = v27
        _ = v28
    }
    if false {
        var v29 uint32 = uint32(v2)
        var v30 uint16 = (v4 | (^(v4)))
        _ = v29
        _ = v30
    } else {
        v31 := []int{443}
        v13 = v13
        for fi32 := 0; fi32 < 4; fi32++ {
            v13 -= (v13 % ((v13) | 1))
            switch "hello" {
            case "hello":
                var v33 float64 = 3.14159
                _ = v33
            case "foo bar":
                if false {
                    v34 := GenS1{F0: 18446744073709551615}
                    _ = v34
                }
            case "":
                v35 := GenS4{F0: false, F1: 65535}
                _ = v35
            default:
                v36 := GenS3{F0: -956}
                _ = v36
            }
            for fi37 := 0; fi37 < 18; fi37++ {
                var v38 uint64 = 2
                var v39 uint32 = 888
                v40 := GenS5{F0: ("foo bar" + "\n\t"), F1: 88, F2: ((-2147483648 < -2147483648) || true)}
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
            _ = fi32
        }
        _ = v31
    }
    if false {
        if (v4 == v4) {
            for fi41 := 0; fi41 < 17; fi41++ {
                switch (fi41 | fi41) {
                case 42:
                    fmt.Println(v13)
                default:
                    v42 := GenS4{F0: false, F1: (v4 ^ v4)}
                    _ = v42
                }
                v43 := GenS5{F0: ("foo bar" + ""), F1: genFn0(), F2: (false || (-1 == 1))}
                _ = fi41
                _ = v43
            }
            var v44 uint8 = 42
            var v45 uint = 252
            _ = v44
            _ = v45
        }
    }
    v46 := GenS6{F0: 65535, F1: v2, F2: v1}
    for fi47 := 0; fi47 < 8; fi47++ {
        for fi48 := 0; fi48 < 10; fi48++ {
            var v49 uint16 = (v4 & (v4 | 2))
            v1 -= int8(v49)
            _ = fi48
            _ = v49
        }
        var v50 uint = (v2 - v2)
        _ = fi47
        _ = v50
    }
    delete(v5, "a")
    if (int(v4) < 4294967295) {
        var v51 int = 42
        v52 := []int{v51}
        for fi53 := 0; fi53 < 12; fi53++ {
            var v54 int = (fi53 >> (((v2 >> ((v2) & 63))) & 63))
            _ = fi53
            _ = v54
        }
        _ = v51
        _ = v52
    }
    if (false || false) {
        var v55 uint = 667
        var v56 uint64 = 0
        _ = v55
        _ = v56
    } else {
        v13 -= 1
        for fi57 := 0; fi57 < 11; fi57++ {
            var v58 int = fi57
            for fi59 := 0; fi59 < 4; fi59++ {
                v13 = (v13 / ((1) | 1))
                v2 -= (^(v2))
                _ = fi59
            }
            delete(v3, ("hello" + "foo bar"))
            _ = fi57
            _ = v58
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
