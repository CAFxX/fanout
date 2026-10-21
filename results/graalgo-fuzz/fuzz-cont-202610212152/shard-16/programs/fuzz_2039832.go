package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int16, p1 uint) bool {
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 bool = false
        var v7 string = "hello"
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: -1}
            fmt.Println(fi5, p0)
            var v10 uint8 = 2
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        p1 -= p1
        p1 -= p1
        _ = fi11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS1{F0: -30}
        _ = fi12
        _ = v13
    }
    p1 -= p1
    _ = p0
    _ = p1
    return (true && (1.5 != 0.1))
}
func genFn1(p0 uint32) int64 {
    switch 128 {
    case -297:
        v14 := GenS1{F0: 1}
        _ = v14
    case -691:
        if true {
            fmt.Println(p0)
            v15 := []int{636, 252, 474, 435}
            _ = v15
        }
    case 9223372036854775807:
        var v16 int8 = 99
        _ = v16
    }
    switch -9223372036854775808 {
    case 2147483647:
        switch ("\n\t" + "foo bar") {
        case "a":
            v17 := []int{}
            _ = v17
        case "hello":
            var v18 uint16 = 1
            _ = v18
        }
    default:
        if true {
            var v19 int = 0
            v20 := GenS1{F0: 100}
            _ = v19
            _ = v20
        } else {
            switch "" {
            case "\n\t":
                fmt.Println(p0)
            default:
                if (false && true) {
                    v21 := []int{-39}
                    switch "foo bar" {
                    case "hello":
                        if len(v21) > 0 {
                            v21[0] = 188
                        }
                    default:
                        for fi22 := 0; fi22 < 16; fi22++ {
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            _ = fi22
                        }
                    }
                    p0 = p0
                    _ = v21
                } else {
                    v23 := []int{-382, 52, 9223372036854775807, -1}
                    _ = v23
                }
            }
            p0 -= (p0 & p0)
        }
    }
    switch "\n\t" {
    case "\n\t":
        var v24 string = "\n\t"
        _ = v24
    case "":
        var v25 uint32 = 2147483647
        _ = v25
    }
    _ = p0
    return 910
}

func main() {
    var v1 uint8 = 2
    v2 := []uint32{2147483647, 4294967295, 0}
    var v3 bool = true
    v4 := make(map[uint32]int)
    for fi26 := 0; fi26 < 11; fi26++ {
        if len(v2) > 0 {
            v2[0] = 184
        }
        if ((!(2 >= v1)) || (!v3)) {
            var v27 int64 = genFn1(2147483648)
            var v28 uint64 = 42
            var v29 uint64 = 1
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi26
    }
    v30 := GenS1{F0: -1}
    for fi31 := 0; fi31 < 4; fi31++ {
        v3 = (42 > 65535)
        _ = fi31
    }
    if ("" < ("foo bar" + "\n\t")) {
        v32 := GenS1{F0: 42}
        _ = v32
    } else {
        v33 := GenS1{F0: -128}
        v34 := GenS1{F0: 127}
        v1 -= v1
        _ = v33
        _ = v34
    }
    switch "\n\t" {
    case "hello":
        if (false || (1 <= 4294967295)) {
            var v35 int8 = 1
            var v36 uint32 = 2147483647
            switch "\n\t" {
            case "hello":
                switch int(v1) {
                case -860:
                    for fi37 := 0; fi37 < 4; fi37++ {
                        v2 = append(v2, (v36 % ((v36) | 1)))
                        _ = len(v4)
                        _ = fi37
                    }
                default:
                    v38 := v4[(^(v36))]
                    _ = v38
                }
            case "foo bar":
                fmt.Println(v35, v3)
            default:
                for fi39 := 0; fi39 < 17; fi39++ {
                    for fi40 := 0; fi40 < 14; fi40++ {
                        v41 := GenS1{F0: (v35 << ((886) & 63))}
                        v42 := GenS2{F0: v1, F1: 242, F2: -887}
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    v3 = ((0 > 957) && (!v3))
                    _ = fi39
                }
            }
            _ = v35
            _ = v36
        } else {
            v1 -= (v1 & (v1 + (^(v1))))
            v43 := GenS2{F0: v1, F1: (v1 + (^(v1))), F2: -1}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v43
        }
    case "":
        var v44 uint8 = v1
        _ = v44
    case "\n\t":
        for fi45 := 0; fi45 < 1; fi45++ {
            v1 += v1
            _ = fi45
        }
    default:
        _ = len(v2)
    }
    var v46 int32 = 0
    var v47 int8 = 4
    if true {
        var v48 float64 = 3.14159
        _ = v48
    }
    switch "hello" {
    case "":
        fmt.Println(v46)
    default:
        v49 := v4[42]
        _ = v49
    }
    var v50 int = 420
    v51 := GenS3{F0: 2147483648}
    fmt.Println(v46, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v46
    _ = v47
    _ = v50
    _ = v51
}
