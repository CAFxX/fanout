package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint64; F1 int16; F2 int32 }
type GenS5 struct { F0 int8; F1 uint64; F2 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int }
func genFn0(p0 bool, p1 int32, p2 int8) uint16 {
    if p0 {
        p1 = 179
        v6 := GenS1{F0: 4294967295}
        var v7 uint8 = 42
        _ = v6
        _ = v7
    }
    var v8 int32 = (p1 - 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 926
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 42
    v3 := []string{("a" + "\n\t"), "foo bar", "\n\t"}
    v4 := make(map[uint32]int)
    var v5 uint16 = (v2 % (((v2 >> ((772) & 63))) | 1))
    for fi9 := 0; fi9 < 10; fi9++ {
        v2 = (v2 | (v2 << ((772) & 63)))
        _ = fi9
    }
    v10 := GenS2{F0: 255, F1: 42}
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int32 = 2147483647
            v14 := GenS3{F0: int8(v5), F1: (fi12 + fi11)}
            switch 4294967295 {
            case -485:
                v15 := GenS3{F0: 127, F1: fi11}
                _ = v15
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        if false {
            if true {
                fmt.Println(v5, v5)
                switch ("" + "\n\t") {
                case "hello":
                    fmt.Println(fi11, v2)
                case "foo bar":
                    if len(v3) > 0 {
                        v3[0] = ("foo bar" + "\n\t")
                    }
                case "":
                    fmt.Println(v2)
                default:
                    v16 := GenS4{F0: uint64(v5), F1: 1, F2: 42}
                    _ = v16
                }
            }
        }
        fmt.Println(fi11, fi11, v5)
        _ = fi11
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            if false {
                if (482 != 42) {
                    v19 := v1[42]
                    if (4294967295 == 18446744073709551615) {
                        var v20 int16 = -160
                        v5 = (v2 / ((v2) | 1))
                        v20 += v20
                        _ = v20
                    }
                    _ = v19
                }
            } else {
                v21 := GenS3{F0: 100, F1: fi17}
                v5 += (^(v5))
                var v22 string = "foo bar"
                _ = v21
                _ = v22
            }
            _ = fi18
        }
        v23 := GenS5{F0: 1, F1: 1, F2: 9}
        _ = fi17
        _ = v23
    }
    var v24 string = "hello"
    var v25 uint16 = genFn0(true, 1, 42)
    var v26 int64 = -851
    v5 -= v25
    if false {
        v2 -= (v2 + 735)
        if (32767 != -187) {
            switch -960 {
            case -440:
                for fi27 := 0; fi27 < 18; fi27++ {
                    for fi28 := 0; fi28 < 18; fi28++ {
                        v29 := GenS6{F0: v26}
                        if true {
                            v25 = (v25 & 421)
                            v24 = ("\n\t" + "")
                            if ((97 != 951) && false) {
                                var v30 float64 = 3.14159
                                _ = v30
                            } else {
                                fmt.Println(v25, v2)
                                v26 += (-(v26))
                                v24 = "hello"
                            }
                        }
                        _ = fi28
                        _ = v29
                    }
                    _ = fi27
                }
            case 31:
                var v31 bool = ((false || false) && true)
                _ = v31
            }
            var v32 uint16 = v5
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        fmt.Println(fi33)
        _ = fi33
    }
    var v34 bool = true
    v24 = "a"
    v35 := GenS7{F0: false}
    for fi36 := 0; fi36 < 10; fi36++ {
        for fi37 := 0; fi37 < 12; fi37++ {
            var v38 int16 = -32768
            v39 := GenS8{F0: (fi37 ^ (fi37 / (((fi36 / (((fi37 % (((fi37 / (((^(fi36))) | 1))) | 1))) | 1))) | 1)))}
            _ = fi37
            _ = v38
            _ = v39
        }
        v40 := v4[2147483648]
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi36
        _ = v40
    }
    v41 := v1[42]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v41
    _ = v5
}
