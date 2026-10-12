package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint8; F1 uint8; F2 string }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 uint16; F2 int32 }
func genFn0(p0 int8, p1 uint8, p2 int16) int8 {
    switch ("hello" + ((("\n\t" + ("foo bar" + "foo bar")) + "\n\t") + "hello")) {
    case "\n\t":
        p2 = p2
    case "hello":
        v4 := GenS1{F0: 862, F1: 42}
        _ = v4
    default:
        p2 -= -347
    }
    var v5 string = ""
    p2 += 0
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}
func genFn1(p0 uint16) int {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{fi6}
        var v8 int32 = -122
        if true {
            p0 = (p0 * 42)
            p0 = 65535
        } else {
            v8 = (v8 << ((824) & 63))
            var v9 uint64 = 629
            if true {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        var v10 uint16 = (p0 / (((^(p0))) | 1))
        switch "\n\t" {
        case "hello":
            v10 = 2
        case "a":
            for fi11 := 0; fi11 < 20; fi11++ {
                if false {
                    switch "hello" {
                    case "hello":
                        var v12 uint = 42
                        _ = v12
                    case "foo bar":
                        for fi13 := 0; fi13 < 11; fi13++ {
                            var v14 int64 = 766
                            v10 += 0
                            var v15 uint8 = 128
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    default:
                        var v16 int8 = genFn0(-1, 127, -1)
                        _ = v16
                    }
                }
                _ = fi11
            }
        case "\n\t":
            v17 := []int{}
            _ = v17
        default:
            fmt.Println(v10, p0)
        }
        _ = v10
    } else {
        var v18 uint8 = 128
        v18 -= v18
        v19 := GenS2{F0: ("a" + ("hello" + "foo bar")), F1: 32768}
        _ = v18
        _ = v19
    }
    fmt.Println(p0)
    _ = p0
    return 2147483647
}

func main() {
    var v1 string = ("hello" + (("\n\t" + (("hello" + "") + ("hello" + ("foo bar" + ("" + "hello"))))) + ("\n\t" + "foo bar")))
    var v2 int8 = -1
    var v3 bool = (32767 != 42)
    var v20 int = genFn1(956)
    v21 := []int{genFn1(148)}
    v22 := GenS2{F0: v1, F1: (v20 << ((681) & 63))}
    var v23 uint32 = 2147483647
    v2 += v2
    if v3 {
        v2 = (v2 + (v2 >> ((724) & 63)))
        switch v1 {
        case "foo bar":
            var v24 int8 = (v2 << ((18446744073709551615) & 63))
            _ = v24
        default:
            v2 -= (v2 ^ 100)
        }
    } else {
        v25 := GenS2{F0: (v1 + ((v1 + ("" + "hello")) + v1)), F1: v20}
        var v26 int = 32768
        var v27 uint16 = 42
        _ = v25
        _ = v26
        _ = v27
    }
    _ = len(v21)
    v28 := GenS3{F0: 128, F1: 42, F2: v1}
    v3 = (float64(0.1) <= 0.0)
    for fi29 := 0; fi29 < 3; fi29++ {
        if v3 {
            _ = len(v21)
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 uint16 = 1
                v23 = (v23 >> ((42) & 63))
                var v32 int = fi29
                _ = fi30
                _ = v31
                _ = v32
            }
            if false {
                v23 -= (^(v23))
                v23 += (^(v23))
            } else {
                var v33 string = v1
                var v34 uint = 1
                _ = v33
                _ = v34
            }
        }
        if v3 {
            fmt.Println(v3)
            v20 += v20
        }
        v3 = (32767 <= 32767)
        _ = fi29
    }
    if v3 {
        var v35 string = "\n\t"
        for fi36 := 0; fi36 < 18; fi36++ {
            v2 = (-(v2))
            v3 = v3
            var v37 uint16 = 0
            _ = fi36
            _ = v37
        }
        _ = v35
    } else {
        if (1.5 >= 1.5) {
            var v38 int32 = 302
            v23 = v23
            for fi39 := 0; fi39 < 15; fi39++ {
                switch (v20 << ((0) & 63)) {
                case 9223372036854775807:
                    v40 := GenS4{F0: 89, F1: 127}
                    _ = v40
                case 452:
                    switch (v1 + (v1 + (v1 + (v1 + v1)))) {
                    case "\n\t":
                        var v41 uint16 = 581
                        _ = v41
                    case "a":
                        var v42 int8 = v2
                        _ = v42
                    case "hello":
                        v1 = "foo bar"
                    }
                }
                v43 := GenS5{F0: (!true)}
                _ = fi39
                _ = v43
            }
            _ = v38
        }
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
    }
    fmt.Println(v23, v2)
    fmt.Println(v3)
    if v3 {
        fmt.Println(v2, v3, v3)
        for fi44 := 0; fi44 < 14; fi44++ {
            for fi45 := 0; fi45 < 1; fi45++ {
                fmt.Println(v2, v3)
                v46 := GenS3{F0: 255, F1: 112, F2: (("hello" + (v1 + v1)) + v1)}
                _ = fi45
                _ = v46
            }
            _ = fi44
        }
        switch v20 {
        case 89:
            if true {
                if ((v20 + 0) != (v20 | v20)) {
                    fmt.Println(v2, v20)
                    var v47 uint16 = 831
                    _ = v47
                } else {
                    if (42 > 18446744073709551615) {
                        if (!(v20 > (v20 ^ (v20 << ((991) & 63))))) {
                            v2 = v2
                            var v48 bool = ("a" >= (v1 + v1))
                            v48 = v3
                            _ = v48
                        } else {
                            var v49 int64 = int64(v20)
                            var v50 int8 = -128
                            _ = v49
                            _ = v50
                        }
                    }
                    if (430 >= 11) {
                        if (v20 != 9223372036854775807) {
                            v23 -= v23
                            v2 -= (v2 % ((1) | 1))
                        } else {
                            _ = len(v21)
                            v3 = (!v3)
                            if v3 {
                                v21 = append(v21, 445)
                            } else {
                                var v51 uint64 = 0
                                _ = v51
                            }
                        }
                        v2 += v2
                    }
                    v1 = v1
                }
                for fi52 := 0; fi52 < 14; fi52++ {
                    fmt.Println(fi52, v3, v3)
                    var v53 float64 = ((1.5 / 1.0) * 1e10)
                    fmt.Println(v2, v20, fi52)
                    _ = fi52
                    _ = v53
                }
                for fi54 := 0; fi54 < 20; fi54++ {
                    _ = len(v21)
                    for fi55 := 0; fi55 < 15; fi55++ {
                        v56 := GenS3{F0: 127, F1: 250, F2: v1}
                        switch -2147483648 {
                        case 127:
                            v57 := GenS6{F0: 255, F1: 211, F2: -989}
                            _ = v57
                        case 128:
                            v3 = true
                        case -1:
                            if len(v21) > 0 {
                                v21[0] = (v20 / ((fi54) | 1))
                            }
                        }
                        _ = fi55
                        _ = v56
                    }
                    _ = fi54
                }
            } else {
                var v58 int64 = -1
                fmt.Println(v23, v58)
                _ = v58
            }
        case 4294967295:
            var v59 int8 = v2
            _ = v59
        }
    } else {
        v20 = (v20 / ((4294967295) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
}
