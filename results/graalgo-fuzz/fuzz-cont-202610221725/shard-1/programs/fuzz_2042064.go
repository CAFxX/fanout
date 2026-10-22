package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16; F1 uint8; F2 int8 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 uint64; F2 int }
type GenS7 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS8 struct { F0 int16; F1 int32; F2 uint16 }
func genFn0(p0 int8) bool {
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            switch "a" {
            case "":
                p0 += 100
            case "foo bar":
                var v9 uint16 = 32767
                v10 := GenS1{F0: 42}
                fmt.Println(v9, p0, fi8)
                _ = v10
                _ = v9
            case "\n\t":
                p0 += (p0 * p0)
                if true {
                    var v11 uint16 = 255
                    v12 := []int{fi7, fi7, (fi7 | fi7)}
                    v13 := GenS2{F0: 1, F1: 65535}
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    fmt.Println(p0, p0, fi7)
                    fmt.Println(fi8, p0)
                    var v14 uint = 2
                    _ = v14
                }
                var v15 bool = false
                _ = v15
            default:
                switch (^(fi8)) {
                case 815:
                    var v16 int = (fi7 / (((fi8 >> ((270) & 63))) | 1))
                    _ = v16
                case -941:
                    var v17 string = ("" + (("\n\t" + "hello") + "hello"))
                    _ = v17
                case 630:
                    v18 := []int{}
                    _ = v18
                }
            }
            _ = fi8
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 uint = 42
            v21 := []int{2147483647, fi7, 32767}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi7
    }
    var v22 int64 = -9223372036854775808
    var v23 uint64 = 42
    if true {
        var v24 int16 = 32767
        _ = v24
    } else {
        v23 += (^(v23))
    }
    v25 := []int{-1, 42, 127}
    _ = p0
    _ = v22
    _ = v23
    _ = v25
    return ((!false) && true)
}
func genFn1() int32 {
    var v26 int32 = -35
    v27 := []int{32768, -851, -2147483648}
    _ = v26
    _ = v27
    return (^(v26))
}

func main() {
    var v1 int16 = -188
    v2 := []uint32{2, 2147483647, 924}
    var v3 uint32 = 42
    var v4 uint16 = 42
    var v5 float64 = 0.1
    v6 := []string{"\n\t", "\n\t"}
    switch 0 {
    case 32767:
        for fi28 := 0; fi28 < 8; fi28++ {
            v5 = -2.25
            v3 -= (v3 ^ (v3 * (v3 ^ v3)))
            for fi29 := 0; fi29 < 12; fi29++ {
                var v30 int = (fi29 * fi28)
                var v31 int32 = 42
                v3 = 204
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
    default:
        for fi32 := 0; fi32 < 6; fi32++ {
            fmt.Println(v5, v5, v4)
            _ = fi32
        }
    }
    if true {
        var v33 uint = 798
        v34 := GenS3{F0: (^(v1)), F1: 16, F2: -39}
        _ = v33
        _ = v34
    } else {
        _ = len(v6)
        switch -9223372036854775808 {
        case -1:
            v1 += (v1 << ((2) & 63))
        }
    }
    var v35 int64 = 9223372036854775807
    if (4294967295 <= 4294967295) {
        for fi36 := 0; fi36 < 17; fi36++ {
            v35 = (v35 << ((42) & 63))
            _ = fi36
        }
        v4 -= v4
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        for fi38 := 0; fi38 < 18; fi38++ {
            var v39 uint64 = 4294967295
            switch (fi38 << ((597) & 63)) {
            case 42:
                v5 = (v5 + v5)
            case 32767:
                if (-40 != 44) {
                    if (v5 > v5) {
                        v5 = ((v5 - v5) + v5)
                    }
                }
            }
            _ = fi38
            _ = v39
        }
        for fi40 := 0; fi40 < 15; fi40++ {
            if (6 <= -1) {
                fmt.Println(v4, fi40)
                _ = len(v6)
                if len(v2) > 0 {
                    v2[0] = 65535
                }
            }
            if (false || (-128 < 42)) {
                v1 += (v1 ^ v1)
                v41 := GenS3{F0: -48, F1: 128, F2: int8(v35)}
                var v42 int = (fi37 ^ (fi37 ^ fi37))
                _ = v41
                _ = v42
            } else {
                v43 := GenS4{F0: 42, F1: 444, F2: 179}
                if false {
                    switch "" {
                    case "":
                        var v44 int64 = (v35 * -9223372036854775808)
                        _ = v44
                    case "\n\t":
                        v45 := GenS5{F0: true}
                        _ = v45
                    }
                    var v46 uint = 18446744073709551615
                    _ = v46
                } else {
                    switch "hello" {
                    case "":
                        if len(v2) > 0 {
                            v2[0] = 63
                        }
                    case "hello":
                        v47 := GenS6{F0: 2, F1: 0, F2: fi37}
                        _ = v47
                    default:
                        var v48 int = fi40
                        _ = v48
                    }
                }
                fmt.Println(v5, v1, fi37)
                _ = v43
            }
            v5 = (v5 / 0.1)
            _ = fi40
        }
        if (((!(((0.1 - ((float32(0.1) * 3.14159) / 0.1)) - 1.5) > (3.14159 + -2.25))) && (genFn1() < 1)) && (true || ((2147483647 >= v3) && true))) {
            v3 = (^(v3))
            for fi49 := 0; fi49 < 8; fi49++ {
                v50 := GenS5{F0: false}
                v3 += v3
                _ = fi49
                _ = v50
            }
        } else {
            switch ("hello" + (("\n\t" + ("\n\t" + "")) + ("" + ""))) {
            case "hello":
                if true {
                    fmt.Println(v35, v35, v5)
                    switch "" {
                    case "a":
                        var v51 uint = uint(v3)
                        _ = v51
                    }
                }
            case "\n\t":
                v52 := GenS5{F0: true}
                _ = v52
            default:
                v53 := GenS7{F0: (v3 >> ((0) & 63)), F1: -700, F2: (v3 & 42)}
                _ = v53
            }
            if (!(-818 <= 42)) {
                for fi54 := 0; fi54 < 10; fi54++ {
                    v55 := GenS1{F0: 42}
                    _ = fi54
                    _ = v55
                }
            }
        }
        _ = fi37
    }
    var v56 int16 = int16(v4)
    var v57 int64 = 9223372036854775807
    if (!true) {
        v1 = 1
        if len(v6) > 0 {
            v6[0] = ""
        }
        var v58 int = -715
        _ = v58
    } else {
        fmt.Println(v57, v3)
        for fi59 := 0; fi59 < 3; fi59++ {
            var v60 bool = (1 != 251)
            _ = fi59
            _ = v60
        }
    }
    v61 := GenS8{F0: (v1 / ((-32768) | 1)), F1: -552, F2: v4}
    v3 += (v3 / ((v3) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v6
    _ = v61
}
