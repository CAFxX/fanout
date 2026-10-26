package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 string; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint16; F1 string }
func genFn0(p0 int32, p1 uint8, p2 uint64) uint32 {
    v6 := GenS1{F0: (p1 << ((724) & 63))}
    p0 += p0
    for fi7 := 0; fi7 < 8; fi7++ {
        switch fi7 {
        case 127:
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 float32 = 1.5
                v10 := GenS2{F0: uint16(p2), F1: (("\n\t" + "\n\t") + ("\n\t" + "foo bar")), F2: 2147483648}
                v11 := []int{(fi8 * (fi8 % (((fi7 % (((fi7 / ((fi8) | 1))) | 1))) | 1))), (fi7 % ((1) | 1)), (fi7 >> ((18446744073709551615) & 63))}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        case 42:
            var v12 float64 = (-2.25 * ((0.0 + (1e10 - 0.0)) - 1.5))
            _ = v12
        default:
            v13 := GenS1{F0: (^(p1))}
            _ = v13
        }
        p0 = int32(p2)
        var v14 uint16 = 830
        _ = fi7
        _ = v14
    }
    if true {
        var v15 uint64 = 42
        for fi16 := 0; fi16 < 1; fi16++ {
            switch 32767 {
            case -480:
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 string = "a"
                    var v19 uint8 = (p1 + (p1 & p1))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            fmt.Println(p0, v15)
            _ = fi16
        }
        _ = v15
    } else {
        if false {
            var v20 int8 = -1
            var v21 int64 = -9223372036854775808
            fmt.Println(p1, p0, p0)
            _ = v20
            _ = v21
        } else {
            v22 := GenS3{F0: 0, F1: 2147483648, F2: p1}
            _ = v22
        }
        fmt.Println(p0, p2, p1)
        v23 := []int{-531, -9223372036854775808, -292}
        _ = v23
    }
    v24 := []int{0, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v6
    return 42
}

func main() {
    var v1 float64 = 0.1
    v2 := []uint32{489, 2, 2}
    var v3 uint16 = 1
    v4 := make(map[int]int)
    v5 := []string{"hello"}
    switch -2147483648 {
    case -864:
        if (!(1 > 1)) {
            fmt.Println(v3, v1)
        }
    case 2147483647:
        var v25 uint32 = 2147483647
        _ = v25
    case 529:
        for fi26 := 0; fi26 < 9; fi26++ {
            for fi27 := 0; fi27 < 20; fi27++ {
                fmt.Println(fi27, v3, v1)
                var v28 uint8 = 127
                _ = fi27
                _ = v28
            }
            var v29 string = ""
            for fi30 := 0; fi30 < 2; fi30++ {
                var v31 int32 = int32(fi30)
                var v32 int8 = 89
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi26
            _ = v29
        }
    default:
        v1 = v1
    }
    v33 := GenS4{F0: (^(v3)), F1: "a"}
    v3 = v3
    v1 = 3.14159
    var v34 uint32 = 839
    var v35 int32 = 0
    var v36 string = ""
    fmt.Println(v36, v36, v36)
    if true {
        v2 = append(v2, uint32(v3))
    }
    fmt.Println(v35, v36, v35)
    if false {
        delete(v4, 0)
        var v37 string = v36
        _ = v37
    } else {
        delete(v4, 32768)
        v34 -= 2
        for fi38 := 0; fi38 < 14; fi38++ {
            for fi39 := 0; fi39 < 13; fi39++ {
                fmt.Println(v34, v1, v3)
                if (!((fi38 * 340) <= (-(fi39)))) {
                    v35 += (v35 >> ((42) & 63))
                }
                var v40 int32 = (v35 / (((v35 % ((-1) | 1))) | 1))
                _ = fi39
                _ = v40
            }
            v36 = v36
            var v41 int8 = -128
            _ = fi38
            _ = v41
        }
    }
    v4[2147483647] = 42
    v35 -= (-(v35))
    if (true || false) {
        switch ("foo bar" + ((v36 + ("foo bar" + ("a" + v36))) + ((("foo bar" + "") + "\n\t") + v36))) {
        case "":
            v1 = v1
        case "\n\t":
            if true {
                if false {
                    v42 := GenS1{F0: 102}
                    _ = v42
                }
                if (751 <= 0) {
                    v43 := v4[-2147483648]
                    var v44 int64 = 0
                    for fi45 := 0; fi45 < 3; fi45++ {
                        v1 += v1
                        _ = len(v4)
                        _ = fi45
                    }
                    _ = v43
                    _ = v44
                }
                var v46 string = v36
                _ = v46
            }
        case "foo bar":
            v47 := GenS4{F0: (v3 * v3), F1: ((("hello" + "\n\t") + v36) + v36)}
            _ = v47
        default:
            var v48 uint64 = 255
            _ = v48
        }
        var v49 bool = true
        var v50 uint16 = 540
        _ = v49
        _ = v50
    }
    v1 = 1e10
    fmt.Println(v35, v35, v3)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
