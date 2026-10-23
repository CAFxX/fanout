package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint16 }
func genFn0(p0 uint8, p1 bool) string {
    v4 := []int{}
    p1 = (814 < 255)
    _ = p0
    _ = p1
    _ = v4
    return ""
}

func main() {
    var v1 uint32 = 65535
    var v2 uint8 = 0
    v3 := make(map[uint32]int)
    v5 := GenS1{F0: true, F1: "", F2: 65535}
    switch 42 {
    case 927:
        if (133 < 591) {
            v2 = (v2 ^ (v2 | 0))
            fmt.Println(v2, v2)
        }
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 float64 = 1e10
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 uint16 = 340
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: true, F1: "foo bar", F2: 42}
        _ = fi6
        _ = v11
    }
    if true {
        if false {
            v3[v1] = 127
            v12 := GenS1{F0: ((483 == 18446744073709551615) || true), F1: "hello", F2: 953}
            fmt.Println(v1, v1)
            _ = v12
        } else {
            var v13 float64 = -2.25
            v14 := GenS2{F0: -949, F1: 2}
            for fi15 := 0; fi15 < 6; fi15++ {
                v13 = v13
                for fi16 := 0; fi16 < 19; fi16++ {
                    v1 = (^(v1))
                    if true {
                        var v17 uint32 = (v1 - (v1 & v1))
                        v1 -= (v17 | (v17 & (v1 + (^(v17)))))
                        var v18 uint32 = (^(v1))
                        _ = v17
                        _ = v18
                    }
                    _ = fi16
                }
                _ = fi15
            }
            _ = v13
            _ = v14
        }
    } else {
        var v19 uint8 = v2
        var v20 int16 = 1
        _ = v19
        _ = v20
    }
    v2 = 1
    for fi21 := 0; fi21 < 2; fi21++ {
        v2 -= 198
        if false {
            for fi22 := 0; fi22 < 8; fi22++ {
                v23 := []int{0, fi22, (fi22 & (fi21 & (fi22 | fi22))), fi21}
                if len(v23) > 0 {
                    v23[0] = 4294967295
                }
                _ = fi22
                _ = v23
            }
            if ((!(2 > 42)) || ((false && (544 >= uint(v1))) || true)) {
                var v24 int16 = 1
                _ = v24
            } else {
                for fi25 := 0; fi25 < 13; fi25++ {
                    v1 -= v1
                    if false {
                        v2 -= v2
                        switch "" {
                        case "":
                            v26 := []int{(fi21 + (fi21 ^ (fi25 ^ (fi25 & -9223372036854775808))))}
                            _ = v26
                        default:
                            switch (fi21 - -933) {
                            case 4294967295:
                                switch (fi21 - fi21) {
                                case -433:
                                    fmt.Println(fi25, fi25)
                                default:
                                    v27 := GenS3{F0: -324}
                                    _ = v27
                                }
                            case -648:
                                var v28 uint8 = uint8(v1)
                                _ = v28
                            }
                        }
                        v29 := []int{fi25}
                        _ = v29
                    } else {
                        v1 = (v1 - 807)
                        v30 := GenS4{F0: uint16(v1), F1: (v1 % ((v1) | 1)), F2: 255}
                        v31 := GenS1{F0: false, F1: ("hello" + "a"), F2: 42}
                        _ = v30
                        _ = v31
                    }
                    _ = fi25
                }
                var v32 bool = (127 <= 35)
                _ = v32
            }
            delete(v3, 644)
        } else {
            v33 := []int{fi21, fi21, 128}
            fmt.Println(v2, v1)
            _ = v33
        }
        _ = fi21
    }
    v34 := []int{128, 775, 127, -590}
    switch 268 {
    case -2147483648:
        if ((37 >= 908) || (("hello" + "hello") == "\n\t")) {
            var v35 uint8 = (^(v2))
            _ = v35
        } else {
            switch "foo bar" {
            case "":
                _ = len(v34)
            }
        }
    }
    fmt.Println(v1, v2)
    switch "a" {
    case "foo bar":
        fmt.Println(v1)
    case "\n\t":
        for fi36 := 0; fi36 < 16; fi36++ {
            v1 = v1
            _ = fi36
        }
    case "a":
        var v37 uint32 = (^(v1))
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v5
}
