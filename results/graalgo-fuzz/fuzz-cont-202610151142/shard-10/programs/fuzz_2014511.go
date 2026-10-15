package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 bool, p1 uint64) bool {
    v5 := GenS1{F0: 42, F1: 2}
    p0 = p0
    v6 := GenS2{F0: 18446744073709551615, F1: 0}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (0 <= 65535)
}
func genFn1(p0 uint8) int64 {
    if (false && false) {
        var v7 float64 = 1.5
        var v8 int8 = -109
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 179, F1: 819}
    _ = p0
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 32767
    var v3 string = ""
    v4 := []string{(v3 + (v3 + v3)), (("" + ((v3 + "\n\t") + (v3 + (v3 + (v3 + ("foo bar" + "foo bar")))))) + "\n\t")}
    fmt.Println(v2, v2)
    switch 42 {
    case 299:
        for fi10 := 0; fi10 < 16; fi10++ {
            v2 -= v2
            fmt.Println(v3, v2, fi10)
            _ = fi10
        }
    case -531:
        if (true && (!(("hello" + "hello") < "a"))) {
            v4 = append(v4, ("\n\t" + v3))
            for fi11 := 0; fi11 < 13; fi11++ {
                switch fi11 {
                case 32768:
                    if (false && true) {
                        for fi12 := 0; fi12 < 5; fi12++ {
                            fmt.Println(fi11, fi12, v2)
                            _ = fi12
                        }
                    }
                default:
                    _ = len(v4)
                }
                _ = fi11
            }
            var v13 uint64 = 321
            _ = v13
        } else {
            v2 += (v2 | 32767)
            v2 = (^(v2))
        }
    default:
        fmt.Println(v3, v3)
    }
    if (9223372036854775807 < -617) {
        switch 127 {
        case -1:
            fmt.Println(v2, v2)
        case 32768:
            _ = len(v4)
        case 4294967295:
            if (1e10 == (-2.25 + 1e10)) {
                v3 = ("a" + "a")
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 bool = true
                    _ = fi14
                    _ = v15
                }
                var v16 uint16 = (v2 / ((908) | 1))
                _ = v16
            }
        default:
            var v17 int64 = -880
            _ = v17
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v3 = ((v3 + ("foo bar" + "")) + ("" + ((v3 + (v3 + ("" + ((v3 + (v3 + v3)) + "\n\t")))) + v3)))
        _ = fi18
    }
    if (2147483648 != 4294967295) {
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 int32 = 1
            v3 = "\n\t"
            _ = fi19
            _ = v20
        }
    } else {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 uint16 = (^(v2))
            var v23 uint32 = 744
            if true {
                var v24 uint8 = 127
                v23 += (v23 >> ((415) & 63))
                _ = v24
            }
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    if (!((!true) && (!true))) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v25 := GenS2{F0: 42, F1: 159}
        _ = v25
    }
    v26 := GenS3{F0: (v2 & v2), F1: (-14 <= -961)}
    v27 := v1[90]
    var v28 int = 32768
    if true {
        for fi29 := 0; fi29 < 9; fi29++ {
            v30 := GenS2{F0: 46, F1: 4294967295}
            var v31 uint16 = v2
            v1[118] = (v3 + v27)
            _ = fi29
            _ = v30
            _ = v31
        }
    } else {
        var v32 int8 = 100
        var v33 bool = true
        var v34 uint32 = 2147483648
        _ = v32
        _ = v33
        _ = v34
    }
    _ = len(v1)
    if false {
        if (true || false) {
            v2 = v2
        } else {
            switch "a" {
            case "foo bar":
                var v35 uint32 = 65535
                _ = v35
            case "":
                for fi36 := 0; fi36 < 18; fi36++ {
                    switch v3 {
                    case "":
                        v28 = v28
                    case "a":
                        _ = len(v4)
                    default:
                        v27 = (v3 + ("\n\t" + "foo bar"))
                    }
                    var v37 uint16 = v2
                    _ = fi36
                    _ = v37
                }
            case "a":
                v3 = v27
            default:
                var v38 uint64 = 2
                _ = v38
            }
        }
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        if true {
            var v40 int8 = 127
            _ = v40
        }
        var v41 float64 = -2.25
        _ = fi39
        _ = v41
    }
    var v42 float32 = -2.25
    for fi43 := 0; fi43 < 17; fi43++ {
        var v44 uint8 = 2
        for fi45 := 0; fi45 < 3; fi45++ {
            v44 -= (v44 ^ (v44 % ((v44) | 1)))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v3 = v27
            _ = fi45
        }
        _ = fi43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v42
}
