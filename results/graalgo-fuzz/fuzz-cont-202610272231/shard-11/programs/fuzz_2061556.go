package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 bool }
func genFn0() uint8 {
    var v4 float64 = -2.25
    v5 := []int{128, 0, 9223372036854775807}
    var v6 uint16 = 42
    for fi7 := 0; fi7 < 5; fi7++ {
        switch 42 {
        case 2147483647:
            switch fi7 {
            case 42:
                v6 = 1
            case 127:
                if false {
                    for fi8 := 0; fi8 < 7; fi8++ {
                        var v9 uint8 = 1
                        v9 -= 110
                        v10 := GenS1{F0: 2}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
            case 9223372036854775807:
                for fi11 := 0; fi11 < 1; fi11++ {
                    v4 += 3.14159
                    v12 := GenS1{F0: 42}
                    _ = fi11
                    _ = v12
                }
            }
        case 42:
            switch ("\n\t" + ("foo bar" + ("\n\t" + "foo bar"))) {
            case "hello":
                switch "\n\t" {
                case "":
                    v13 := GenS2{F0: 0}
                    _ = v13
                case "\n\t":
                    v4 -= float64((1e10 + 1.5))
                case "a":
                    switch 4294967295 {
                    case 44:
                        var v14 int8 = 1
                        _ = v14
                    default:
                        var v15 bool = true
                        _ = v15
                    }
                default:
                    for fi16 := 0; fi16 < 17; fi16++ {
                        switch fi16 {
                        case 386:
                            v17 := GenS3{F0: v6, F1: 42, F2: "\n\t"}
                            _ = v17
                        default:
                            v18 := GenS2{F0: -1}
                            _ = v18
                        }
                        v4 = ((1.5 - v4) / -2.25)
                        _ = fi16
                    }
                }
            case "a":
                for fi19 := 0; fi19 < 14; fi19++ {
                    var v20 uint = 1
                    var v21 float64 = float64(0.1)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            default:
                switch fi7 {
                case 127:
                    if len(v5) > 0 {
                        v5[0] = (fi7 + (-(fi7)))
                    }
                case 32767:
                    var v22 int32 = -2147483648
                    _ = v22
                case -1:
                    v4 += v4
                }
            }
        }
        _ = fi7
    }
    v23 := GenS3{F0: (v6 * (^(v6))), F1: 0, F2: "hello"}
    _ = v23
    _ = v4
    _ = v5
    _ = v6
    return uint8(v6)
}
func genFn1(p0 uint8) int {
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := []int{fi24, int(p0), fi24, fi24}
        _ = fi24
        _ = v25
    }
    p0 = p0
    _ = p0
    return -1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 int = 910
    var v26 int16 = -32768
    if (false || true) {
        v26 += (v26 + v26)
        for fi27 := 0; fi27 < 6; fi27++ {
            fmt.Println(fi27, v3)
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 int = (v3 / (((fi27 % ((426) | 1))) | 1))
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
        if true {
            switch v3 {
            case 32767:
                var v30 string = ("hello" + "")
                _ = v30
            case 128:
                v31 := GenS2{F0: -1}
                _ = v31
            default:
                fmt.Println(v26, v3)
            }
        } else {
            fmt.Println(v26, v3)
        }
    }
    v26 = (v26 % ((-161) | 1))
    if (42 >= 116) {
        v3 += (v3 | v3)
        for fi32 := 0; fi32 < 17; fi32++ {
            v3 += genFn1(127)
            _ = fi32
        }
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        if true {
            if (!(!(!((42 != 42) || (!true))))) {
                v34 := GenS2{F0: -1}
                v35 := []int{1, (^(v3)), (-(v3))}
                _ = v34
                _ = v35
            } else {
                if true {
                    var v36 bool = false
                    _ = v36
                }
                switch "hello" {
                case "foo bar":
                    var v37 bool = ((v26 >> ((0) & 63)) == v26)
                    _ = v37
                case "\n\t":
                    fmt.Println(v26, v3)
                }
            }
            fmt.Println(v3, fi33)
            for fi38 := 0; fi38 < 19; fi38++ {
                _ = len(v2)
                if false {
                    if false {
                        var v39 uint64 = 42
                        _ = v39
                    } else {
                        v26 = (v26 | (v26 & 1000))
                    }
                }
                _ = fi38
            }
        }
        _ = fi33
    }
    v26 += v26
    v3 += (v3 + genFn1(153))
    fmt.Println(v3)
    if true {
        for fi40 := 0; fi40 < 7; fi40++ {
            switch 781 {
            case -151:
                v41 := []int{(v3 ^ (fi40 & -427)), genFn1(42)}
                _ = v41
            case 371:
                var v42 uint16 = 282
                _ = v42
            case 42:
                v43 := GenS2{F0: -57}
                _ = v43
            }
            var v44 int32 = 2147483647
            v26 -= (v26 + (^(v26)))
            _ = fi40
            _ = v44
        }
    } else {
        if (!true) {
            if (false && false) {
                v45 := GenS4{F0: -1}
                v46 := GenS5{F0: 65535}
                _ = v45
                _ = v46
            }
            v47 := GenS6{F0: 311, F1: 0}
            v3 = genFn1(1)
            _ = v47
        }
        v3 += v3
        for fi48 := 0; fi48 < 15; fi48++ {
            var v49 uint16 = 121
            v26 -= (v26 ^ 1)
            _ = fi48
            _ = v49
        }
    }
    v50 := GenS5{F0: 641}
    fmt.Println(v3, v26)
    switch ("\n\t" + "\n\t") {
    case "":
        var v51 string = ""
        _ = v51
    }
    var v52 uint64 = 18446744073709551615
    v53 := GenS7{F0: true}
    if true {
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v50
    _ = v52
    _ = v53
}
