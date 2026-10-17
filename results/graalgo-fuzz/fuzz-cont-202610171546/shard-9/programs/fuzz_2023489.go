package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 bool; F1 bool; F2 uint8 }
type GenS3 struct { F0 bool; F1 int16; F2 bool }
type GenS4 struct { F0 uint16; F1 string; F2 int8 }
type GenS5 struct { F0 int64; F1 string }
func genFn0(p0 uint32) uint64 {
    for fi4 := 0; fi4 < 10; fi4++ {
        p0 = 112
        _ = fi4
    }
    switch "" {
    case "\n\t":
        v5 := []int{-1, 42, -2147483648, 128}
        _ = v5
    case "hello":
        var v6 uint16 = 2
        _ = v6
    }
    v7 := []int{1, 1}
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 int32, p2 uint16) string {
    p2 += (p2 & 42)
    switch 4294967295 {
    case 32768:
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: 42, F1: 100}
            if (!(!(!true))) {
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 uint = p0
                    var v12 uint16 = p2
                    var v13 int16 = 1
                    if (!((!true) || (p1 <= 0))) {
                        switch fi8 {
                        case 2147483647:
                            v14 := GenS1{F0: uint32(v12), F1: 100}
                            _ = v14
                        case -758:
                            v15 := GenS1{F0: 272, F1: 104}
                            _ = v15
                        default:
                            fmt.Println(v13)
                        }
                    } else {
                        v16 := GenS1{F0: 2147483647, F1: 0}
                        var v17 int8 = -91
                        p2 += 255
                        _ = v16
                        _ = v17
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 int64 = 0
                    v19 -= (v19 * v19)
                    _ = fi18
                    _ = v19
                }
                p2 = p2
            }
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 string = "\n\t"
                v22 := []int{(fi20 ^ fi20), fi20, (fi20 + fi20), (fi20 | 42)}
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                v23 := GenS1{F0: 187, F1: 1}
                fmt.Println(v21)
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            _ = fi8
            _ = v9
        }
    case 127:
        if ((true && true) && false) {
            if (!false) {
                var v24 int64 = -1
                v25 := []int{2147483647, 127}
                _ = v24
                _ = v25
            } else {
                fmt.Println(p1)
                p0 = 1
            }
        } else {
            var v26 uint8 = 128
            v27 := []int{-240, 9223372036854775807, -154}
            v28 := GenS2{F0: true, F1: (true || false), F2: v26}
            _ = v26
            _ = v27
            _ = v28
        }
    case -2147483648:
        var v29 string = ("" + ("" + ("hello" + "\n\t")))
        _ = v29
    }
    if true {
        p1 += (^(p1))
    }
    fmt.Println(p1, p2, p0)
    p1 -= (p1 | p1)
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 42
    var v3 uint8 = 42
    delete(v1, 670)
    fmt.Println(v3, v3)
    for fi30 := 0; fi30 < 12; fi30++ {
        v2 += uint(v3)
        _ = fi30
    }
    switch "a" {
    case "\n\t":
        v31 := []int{1, -2147483648}
        _ = v31
    case "a":
        var v32 uint8 = (v3 << (((v2 / ((v2) | 1))) & 63))
        _ = v32
    case "":
        v2 += 113
    default:
        v3 -= v3
    }
    switch ("" + "hello") {
    case "foo bar":
        for fi33 := 0; fi33 < 8; fi33++ {
            v34 := GenS3{F0: true, F1: -58, F2: (65535 == 2147483648)}
            if ((^(v2)) >= 1) {
                var v35 uint8 = (v3 ^ 164)
                _ = v35
            }
            _ = fi33
            _ = v34
        }
    default:
        if ((!(!("a" < "\n\t"))) || true) {
            for fi36 := 0; fi36 < 13; fi36++ {
                v3 = (^(v3))
                v2 = (v2 & (^(v2)))
                _ = fi36
            }
            for fi37 := 0; fi37 < 6; fi37++ {
                v38 := []int{fi37, 179, fi37}
                var v39 float32 = 1e10
                _ = fi37
                _ = v38
                _ = v39
            }
        } else {
            v40 := []int{0}
            v41 := GenS4{F0: 32767, F1: "foo bar", F2: 127}
            for fi42 := 0; fi42 < 2; fi42++ {
                v2 += v2
                v43 := v1[2]
                var v44 int8 = 100
                _ = fi42
                _ = v43
                _ = v44
            }
            _ = v40
            _ = v41
        }
    }
    for fi45 := 0; fi45 < 6; fi45++ {
        v3 = (v3 / ((v3) | 1))
        var v46 uint8 = v3
        for fi47 := 0; fi47 < 7; fi47++ {
            v48 := GenS3{F0: ((^(v2)) != 1), F1: 0, F2: false}
            v49 := GenS4{F0: 2, F1: "foo bar", F2: 1}
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = fi45
        _ = v46
    }
    for fi50 := 0; fi50 < 12; fi50++ {
        var v51 uint = v2
        _ = fi50
        _ = v51
    }
    if true {
        v2 = 255
        v52 := []int{32768, -665, 1}
        _ = v52
    }
    var v53 int = -784
    if false {
        if true {
            v53 = v53
            v54 := []int{(v53 % (((-(v53))) | 1)), (v53 >> (((v2 | (v2 % (((v2 ^ (v2 >> (((v2 | (^(v2)))) & 63)))) | 1)))) & 63)), v53}
            _ = v54
        } else {
            v55 := []int{}
            v56 := GenS1{F0: 65535, F1: -128}
            var v57 string = "foo bar"
            _ = v55
            _ = v56
            _ = v57
        }
    }
    var v58 int16 = 40
    var v59 int8 = -1
    v60 := GenS5{F0: -9223372036854775808, F1: ("hello" + "hello")}
    v3 = v3
    var v61 float64 = ((-2.25 / 1.0) - 3.14159)
    var v62 int64 = -628
    fmt.Println(v58, v53)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v58
    _ = v59
    _ = v60
    _ = v61
    _ = v62
}
