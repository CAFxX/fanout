package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 string; F1 int64; F2 uint8 }
func genFn0(p0 int32, p1 bool, p2 uint8) int16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 790, F1: 1}
        if p1 {
            var v9 float64 = 0.0
            switch fi7 {
            case 42:
                v10 := []int{fi7}
                _ = v10
            }
            switch "a" {
            case "\n\t":
                v11 := GenS1{F0: 1, F1: 65535}
                _ = v11
            default:
                var v12 int32 = 42
                _ = v12
            }
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        p1 = p1
        switch int(p2) {
        case 0:
            v14 := GenS2{F0: ("" + "hello"), F1: -300, F2: p2}
            _ = v14
        case 4294967295:
            switch "a" {
            case "":
                fmt.Println(fi13)
            case "\n\t":
                v15 := GenS2{F0: "hello", F1: 42, F2: (p2 >> ((uint(p2)) & 63))}
                _ = v15
            case "a":
                v16 := GenS1{F0: 32, F1: 1}
                _ = v16
            default:
                var v17 float64 = 1.5
                _ = v17
            }
        case 128:
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 uint64 = 109
                fmt.Println(p1, fi18, p1)
                fmt.Println(fi13, p1, fi13)
                var v20 string = "foo bar"
                _ = fi18
                _ = v19
                _ = v20
            }
        default:
            switch -910 {
            case -509:
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 float32 = 1e10
                    _ = fi21
                    _ = v22
                }
            case -2147483648:
                p1 = (59 >= -89)
            case -518:
                var v23 uint = 118
                _ = v23
            default:
                var v24 int = fi13
                _ = v24
            }
        }
        _ = fi13
    }
    p2 = 177
    p1 = ((("a" + ((("\n\t" + "a") + "hello") + ("" + ""))) + "\n\t") > "foo bar")
    _ = p0
    _ = p1
    _ = p2
    return 749
}
func genFn1(p0 int64, p1 int16) int {
    if false {
        p1 += p1
        var v25 float32 = 1.5
        var v26 uint8 = 128
        _ = v25
        _ = v26
    } else {
        v27 := []int{176, 0, 172}
        v28 := GenS1{F0: 42, F1: 0}
        _ = v27
        _ = v28
    }
    var v29 string = ((("" + "foo bar") + "") + "")
    p1 = (p1 % ((p1) | 1))
    var v30 int64 = (p0 - (p0 | (p0 | -1)))
    _ = p0
    _ = p1
    _ = v29
    _ = v30
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 65535
    var v2 float32 = (float32((1.5 - ((-2.25 * ((((0.0 + (1.5 + 1.5)) * ((1.5 + (3.14159 / 1.5)) - (0.0 + (1.5 * float64(3.14159))))) / (1e10 - float64(1.5))) / (1e10 + (float64((1e10 / (0.1 - 1.5))) / -2.25)))) / 1.5))) / (((1e10 + ((1e10 - float32((1.5 + 1e10))) + ((-2.25 + (0.1 + 0.1)) * 0.1))) + 0.1) * 1.5))
    var v3 uint16 = 255
    var v4 int16 = 1000
    var v5 int32 = 0
    var v6 int16 = 1
    v2 -= (v2 * v2)
    for fi31 := 0; fi31 < 11; fi31++ {
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 int8 = 100
            v1 += (^(v1))
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 float64 = -2.25
        _ = fi34
        _ = v35
    }
    switch "" {
    case "foo bar":
        var v36 uint16 = 2
        _ = v36
    default:
        for fi37 := 0; fi37 < 13; fi37++ {
            v1 = 1
            v4 += (v6 * (v4 % ((genFn0((^(v5)), (true || false), 42)) | 1)))
            if false {
                var v38 int8 = 40
                if (false || (((4294967295 >= 42) && (true || false)) || (!(!false)))) {
                    fmt.Println(v1, v5, v1)
                } else {
                    var v39 int16 = (^(v4))
                    _ = v39
                }
                _ = v38
            } else {
                v40 := []int{(fi37 ^ (^(fi37))), (fi37 & (fi37 << ((504) & 63))), (fi37 ^ genFn1(1, v4)), genFn1(42, 0)}
                _ = v40
            }
            _ = fi37
        }
    }
    if (2 <= v3) {
        v41 := GenS1{F0: 42, F1: (^(v3))}
        _ = v41
    }
    var v42 int16 = 0
    fmt.Println(v2)
    if false {
        v43 := GenS2{F0: "foo bar", F1: -219, F2: 127}
        v1 = v1
        v6 -= v6
        _ = v43
    }
    var v44 float32 = (v2 / (1.5 - v2))
    for fi45 := 0; fi45 < 2; fi45++ {
        v46 := []int{-9223372036854775808, genFn1(42, v6), (^(fi45))}
        if false {
            var v47 bool = false
            _ = v47
        }
        _ = fi45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v5
    _ = v6
}
