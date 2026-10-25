package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 int16; F1 uint16; F2 int64 }
func genFn0(p0 uint64, p1 int, p2 string) uint32 {
    switch "" {
    case "hello":
        if true {
            v6 := []int{(p1 * int(p0)), p1, (p1 % ((p1) | 1))}
            if (255 > 255) {
                var v7 float32 = (((0.0 - 3.14159) / -2.25) / 3.14159)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v7
            }
            for fi8 := 0; fi8 < 1; fi8++ {
                for fi9 := 0; fi9 < 4; fi9++ {
                    v6 = append(v6, p1)
                    _ = fi9
                }
                p0 -= p0
                _ = fi8
            }
            _ = v6
        }
    case "a":
        fmt.Println(p1, p1)
    case "":
        var v10 float64 = (0.0 - (((-2.25 / 1.0) + (-2.25 + -2.25)) / (3.14159 + 1.5)))
        _ = v10
    default:
        switch (("\n\t" + "foo bar") + "") {
        case "foo bar":
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := GenS1{F0: 0, F1: 738}
                for fi13 := 0; fi13 < 7; fi13++ {
                    v14 := GenS2{F0: 42, F1: -32768}
                    if (false && true) {
                        v15 := GenS3{F0: -1, F1: (1 < 0)}
                        v16 := GenS4{F0: 42}
                        p1 += (fi13 << ((657) & 63))
                        v17 := GenS1{F0: uint(p0), F1: 904}
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    if true {
                        var v18 uint32 = 0
                        _ = v18
                    }
                    _ = fi13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
        case "":
            if (42 != 2147483648) {
                v19 := GenS5{F0: (^(p0)), F1: false, F2: 255}
                p2 = p2
                p1 += p1
                _ = v19
            }
        default:
            var v20 string = (p2 + "foo bar")
            _ = v20
        }
    }
    var v21 uint8 = 1
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            p1 += (fi23 - -2147483648)
            var v24 uint = 42
            var v25 int32 = 1
            _ = fi23
            _ = v24
            _ = v25
        }
        v26 := GenS3{F0: -1, F1: (false || true)}
        v27 := []int{(p1 << ((2) & 63)), 4294967295, (^(p1)), p1}
        _ = fi22
        _ = v26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 65535
}

func main() {
    var v1 int64 = 1
    v2 := make(map[int]int)
    var v3 int8 = -54
    var v4 uint64 = 18446744073709551615
    var v5 uint32 = 203
    if false {
        var v28 bool = true
        _ = v28
    }
    var v29 int32 = 2147483647
    v30 := []int{-448, 42, 563}
    var v31 int = -186
    v4 = 4294967295
    for fi32 := 0; fi32 < 6; fi32++ {
        if ((((!(v5 <= 2147483648)) && true) && (520 < 18446744073709551615)) || ((v3 << ((0) & 63)) == v3)) {
            v33 := GenS1{F0: 4294967295, F1: genFn0(v4, (v31 ^ (-(v31))), "hello")}
            if (("\n\t" >= ("a" + ("foo bar" + ""))) || true) {
                if ((v1 & 42) == v1) {
                    var v34 uint = 376
                    var v35 float32 = -2.25
                    _ = v34
                    _ = v35
                } else {
                    var v36 float64 = 0.1
                    if true {
                        var v37 float32 = float32(1e10)
                        _ = v37
                    }
                    _ = v36
                }
                var v38 uint = 835
                switch (v31 << (((v38 % ((v38) | 1))) & 63)) {
                case 1:
                    var v39 bool = (3.14159 >= 3.14159)
                    _ = v39
                case 9223372036854775807:
                    var v40 int16 = int16(v1)
                    _ = v40
                case 42:
                    var v41 int16 = -411
                    _ = v41
                }
                _ = v38
            } else {
                v2[1] = -901
            }
            switch (^(v31)) {
            case 2147483647:
                v5 = v5
            case 0:
                fmt.Println(v4)
            case -603:
                v1 = (^(v1))
            }
            _ = v33
        }
        _ = fi32
    }
    v4 = v4
    for fi42 := 0; fi42 < 15; fi42++ {
        var v43 int32 = (v29 * v29)
        _ = fi42
        _ = v43
    }
    var v44 bool = true
    if (255 <= 128) {
        if true {
            var v45 uint8 = 71
            if v44 {
                v4 += (v4 * 0)
            }
            _ = v45
        } else {
            v46 := GenS6{F0: 42, F1: 18446744073709551615}
            var v47 uint8 = 36
            _ = v46
            _ = v47
        }
        var v48 uint = 0
        _ = v48
    } else {
        v44 = v44
    }
    v49 := GenS5{F0: 4294967295, F1: true, F2: 65535}
    v50 := GenS7{F0: v29, F1: v4}
    delete(v2, (^(v31)))
    var v51 int = v31
    if (!((4294967295 > (v51 + 9223372036854775807)) || (!v44))) {
        v52 := GenS4{F0: v1}
        switch "" {
        case "\n\t":
            for fi53 := 0; fi53 < 19; fi53++ {
                for fi54 := 0; fi54 < 2; fi54++ {
                    v55 := GenS6{F0: 4294967295, F1: 116}
                    if (((v5 & (^(v5))) <= v5) && (1000 <= 1000)) {
                        fmt.Println(fi54)
                        v44 = ((1 <= 4294967295) && v44)
                    }
                    _ = fi54
                    _ = v55
                }
                var v56 string = "a"
                var v57 int = v51
                _ = fi53
                _ = v56
                _ = v57
            }
        }
        _ = v52
    } else {
        for fi58 := 0; fi58 < 4; fi58++ {
            v4 = uint64(v5)
            v29 = v29
            if ((v4 != v4) && v44) {
                if (!(4294967295 < v5)) {
                    _ = len(v30)
                    var v59 uint64 = v4
                    v4 += v4
                    _ = v59
                }
                v29 += (v29 + 0)
                v60 := GenS8{F0: -1, F1: 255, F2: (v1 << ((2) & 63))}
                _ = v60
            } else {
                var v61 string = "foo bar"
                v5 = (v5 - 42)
                _ = v61
            }
            _ = fi58
        }
        var v62 int8 = (v3 | (v3 + 42))
        v31 = 9223372036854775807
        _ = v62
    }
    var v63 int8 = (v3 % ((v3) | 1))
    v2[v51] = v51
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v44
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v63
}
