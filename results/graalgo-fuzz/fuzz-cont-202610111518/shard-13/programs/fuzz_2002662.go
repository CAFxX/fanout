package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0() int32 {
    switch "foo bar" {
    case "hello":
        if false {
            switch 42 {
            case 0:
                var v6 uint = 1
                _ = v6
            case -9223372036854775808:
                for fi7 := 0; fi7 < 10; fi7++ {
                    var v8 uint64 = 42
                    v8 = (v8 / ((v8) | 1))
                    var v9 uint = 383
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case 9223372036854775807:
                var v10 bool = ((595 > 1) && (-1 != 127))
                _ = v10
            }
            v11 := GenS1{F0: 72, F1: 983, F2: 0}
            v12 := GenS2{F0: 896, F1: (!(0 < 0))}
            _ = v11
            _ = v12
        }
    case "":
        v13 := GenS2{F0: 42, F1: false}
        _ = v13
    default:
        var v14 int16 = 0
        _ = v14
    }
    switch 0 {
    case 2147483647:
        v15 := []int{-2147483648, -1, -934, -1}
        _ = v15
    case -9223372036854775808:
        v16 := []int{-888, 9223372036854775807, 2147483647, 927}
        _ = v16
    }
    var v17 string = "foo bar"
    v18 := GenS1{F0: 0, F1: 4294967295, F2: 100}
    _ = v17
    _ = v18
    return -340
}
func genFn1() int8 {
    if ((false || true) || false) {
        var v19 uint = 453
        for fi20 := 0; fi20 < 15; fi20++ {
            for fi21 := 0; fi21 < 15; fi21++ {
                fmt.Println(fi21, fi21, fi21)
                v22 := GenS1{F0: 127, F1: 18446744073709551615, F2: 103}
                var v23 uint = 844
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        var v24 uint16 = 0
        _ = v19
        _ = v24
    }
    var v25 float64 = 0.0
    _ = v25
    return 127
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]string)
    var v3 uint16 = 42
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    var v26 int32 = v1
    for fi27 := 0; fi27 < 3; fi27++ {
        if (false && true) {
            if (true && false) {
                for fi28 := 0; fi28 < 7; fi28++ {
                    var v29 uint16 = (v3 * v3)
                    var v30 int64 = 42
                    v1 -= v26
                    fmt.Println(fi27, v26)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                fmt.Println(fi27)
            } else {
                var v31 int = fi27
                var v32 float32 = -2.25
                _ = v31
                _ = v32
            }
            switch (fi27 | (fi27 >> ((18446744073709551615) & 63))) {
            case 0:
                var v33 uint16 = (v3 << ((951) & 63))
                _ = v33
            case -44:
                fmt.Println(v1, v3, fi27)
            case 127:
                v34 := GenS2{F0: genFn0(), F1: true}
                _ = v34
            }
        } else {
            delete(v2, (fi27 | fi27))
        }
        var v35 int = fi27
        v36 := GenS1{F0: 127, F1: 0, F2: 127}
        _ = fi27
        _ = v35
        _ = v36
    }
    if (0 <= 42) {
        v3 = v3
        delete(v4, 686)
        for fi37 := 0; fi37 < 3; fi37++ {
            v3 = v3
            _ = fi37
        }
    }
    v38 := v4[-2147483648]
    v1 += genFn0()
    v3 += (v3 | v3)
    fmt.Println(v26, v26, v3)
    var v39 int = 0
    v3 = v3
    fmt.Println(v1, v26, v1)
    switch v38 {
    case 9223372036854775807:
        v38 += (v38 / (((v39 & (-(v38)))) | 1))
    case -2147483648:
        var v40 int64 = 1
        _ = v40
    case 32767:
        v5[""] = (v38 * v39)
    }
    if false {
        v26 = (v1 << ((373) & 63))
        v2[v39] = ("hello" + "")
        v3 = v3
    } else {
        switch (^(v38)) {
        case 9223372036854775807:
            for fi41 := 0; fi41 < 18; fi41++ {
                v42 := v5["hello"]
                if (v3 > (v3 % ((2) | 1))) {
                    v43 := GenS1{F0: 255, F1: 2, F2: 1}
                    _ = v43
                } else {
                    v44 := []int{v38}
                    v45 := v4[fi41]
                    _ = v44
                    _ = v45
                }
                v42 += fi41
                _ = fi41
                _ = v42
            }
        case 127:
            v39 = 32768
        case 4294967295:
            var v46 uint64 = 1
            _ = v46
        default:
            v26 += v1
        }
        v26 = (-(v1))
        if false {
            v47 := GenS1{F0: 128, F1: 4294967295, F2: 39}
            switch v38 {
            case 818:
                v48 := GenS1{F0: 181, F1: 1, F2: 32}
                _ = v48
            case -9223372036854775808:
                switch ("foo bar" + "") {
                case "a":
                    var v49 uint64 = 167
                    _ = v49
                default:
                    fmt.Println(v26)
                }
            default:
                v50 := GenS2{F0: 42, F1: (1 <= 9223372036854775807)}
                _ = v50
            }
            v26 = genFn0()
            _ = v47
        }
    }
    for fi51 := 0; fi51 < 12; fi51++ {
        v52 := GenS2{F0: v26, F1: (0 < 344)}
        _ = fi51
        _ = v52
    }
    for fi53 := 0; fi53 < 16; fi53++ {
        for fi54 := 0; fi54 < 16; fi54++ {
            switch v38 {
            case 1:
                if true {
                    var v55 int64 = -1
                    v56 := v4[fi54]
                    _ = v55
                    _ = v56
                } else {
                    var v57 uint16 = (^(v3))
                    v38 = 42
                    _ = v57
                }
            case -2147483648:
                v58 := []int{(^(v38)), fi53, fi53, (v38 << ((260) & 63))}
                _ = v58
            case 32768:
                v59 := GenS3{F0: 127, F1: v1}
                _ = v59
            default:
                switch "" {
                case "foo bar":
                    if (655 != v3) {
                        v60 := v2[(fi54 >> ((575) & 63))]
                        _ = v60
                    }
                case "\n\t":
                    delete(v5, "")
                case "hello":
                    var v61 int16 = -32768
                    _ = v61
                default:
                    var v62 int16 = -1
                    _ = v62
                }
            }
            fmt.Println(v1, fi54)
            var v63 int16 = 347
            _ = fi54
            _ = v63
        }
        _ = fi53
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
