package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int32; F1 string; F2 int8 }
type GenS5 struct { F0 int; F1 int64; F2 uint32 }
type GenS6 struct { F0 int32; F1 int64; F2 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int32; F1 uint8; F2 int64 }
func genFn0() string {
    v6 := GenS1{F0: 42, F1: 0}
    v7 := []int{0, -1, -169}
    _ = v6
    _ = v7
    return ("" + ((("hello" + "\n\t") + "foo bar") + ("a" + ("\n\t" + "foo bar"))))
}
func genFn1(p0 int, p1 int) uint8 {
    if false {
        var v8 float64 = (1e10 + ((0.1 + 1e10) + 3.14159))
        var v9 uint32 = 65535
        _ = v8
        _ = v9
    } else {
        v10 := GenS2{F0: 0, F1: 2147483647}
        _ = v10
    }
    if false {
        v11 := []int{p0, p1, p0, (p0 & p1)}
        v12 := GenS1{F0: 65535, F1: 229}
        for fi13 := 0; fi13 < 16; fi13++ {
            if (true || (371 == 2147483648)) {
                var v14 int32 = -2147483648
                _ = v14
            }
            if (!("a" <= "foo bar")) {
                var v15 uint = 185
                _ = v15
            }
            var v16 bool = false
            _ = fi13
            _ = v16
        }
        _ = v11
        _ = v12
    } else {
        v17 := []int{(p1 << ((255) & 63))}
        var v18 uint64 = 2
        if false {
            switch p0 {
            case 9223372036854775807:
                for fi19 := 0; fi19 < 6; fi19++ {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        var v21 int = fi19
                        _ = fi20
                        _ = v21
                    }
                    _ = len(v17)
                    _ = fi19
                }
            default:
                if (!false) {
                    switch (p1 * p1) {
                    case 410:
                        var v22 uint = 65535
                        _ = v22
                    default:
                        for fi23 := 0; fi23 < 16; fi23++ {
                            _ = len(v17)
                            var v24 uint8 = 2
                            _ = fi23
                            _ = v24
                        }
                    }
                } else {
                    p1 = p1
                    var v25 int16 = 32767
                    _ = v25
                }
            }
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        }
        _ = v17
        _ = v18
    }
    v26 := []int{(^(p1))}
    for fi27 := 0; fi27 < 9; fi27++ {
        _ = len(v26)
        p0 -= 9223372036854775807
        _ = fi27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 int32 = 983
        var v30 uint64 = 18446744073709551615
        _ = fi28
        _ = v29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v26
    return 2
}

func main() {
    var v1 float64 = 0.1
    v2 := []int{128, 794, 77}
    var v3 uint32 = 1
    var v4 int8 = int8(v3)
    var v5 bool = false
    if v5 {
        var v31 uint64 = 0
        switch 127 {
        case 0:
            for fi32 := 0; fi32 < 2; fi32++ {
                v1 = v1
                _ = fi32
            }
        default:
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
        }
        var v33 int64 = 42
        _ = v31
        _ = v33
    }
    for fi34 := 0; fi34 < 20; fi34++ {
        var v35 uint64 = 350
        for fi36 := 0; fi36 < 18; fi36++ {
            v37 := GenS3{F0: 1, F1: 1}
            _ = fi36
            _ = v37
        }
        var v38 uint8 = 15
        _ = fi34
        _ = v35
        _ = v38
    }
    var v39 string = "a"
    var v40 uint16 = 842
    for fi41 := 0; fi41 < 12; fi41++ {
        for fi42 := 0; fi42 < 5; fi42++ {
            if (v5 || v5) {
                var v43 float64 = v1
                _ = v43
            }
            _ = fi42
        }
        switch v39 {
        case "\n\t":
            var v44 int16 = 0
            _ = v44
        case "a":
            v1 = (v1 / ((3.14159 + v1) - float64(1.5)))
        case "foo bar":
            var v45 bool = (v4 <= v4)
            _ = v45
        default:
            var v46 uint16 = (^(v40))
            _ = v46
        }
        if v5 {
            if len(v2) > 0 {
                v2[0] = (fi41 - (fi41 % ((-144) | 1)))
            }
            v47 := GenS4{F0: -1, F1: ("a" + ((((v39 + v39) + "foo bar") + v39) + v39)), F2: (v4 << ((42) & 63))}
            _ = v47
        } else {
            for fi48 := 0; fi48 < 16; fi48++ {
                var v49 float64 = float64((0.0 / 0.1))
                fmt.Println(v5)
                if v5 {
                    var v50 uint64 = 42
                    _ = v50
                }
                switch fi48 {
                case -2147483648:
                    switch -197 {
                    case -2147483648:
                        v51 := GenS3{F0: 1, F1: 339}
                        v2 = append(v2, (fi48 & (fi48 % (((fi48 % ((4294967295) | 1))) | 1))))
                        var v52 uint8 = 2
                        _ = v51
                        _ = v52
                    }
                default:
                    v53 := GenS5{F0: (-(fi41)), F1: 0, F2: v3}
                    _ = v53
                }
                switch ("" + "") {
                case "\n\t":
                    v54 := GenS6{F0: 963, F1: -1, F2: 0}
                    v55 := GenS7{F0: v5}
                    _ = v54
                    _ = v55
                }
                _ = fi48
                _ = v49
            }
            if (fi41 < (fi41 % ((576) | 1))) {
                var v56 int16 = -764
                v2 = append(v2, fi41)
                for fi57 := 0; fi57 < 7; fi57++ {
                    var v58 int16 = -2
                    if (!v5) {
                        fmt.Println(v39)
                    }
                    _ = fi57
                    _ = v58
                }
                _ = v56
            }
        }
        _ = fi41
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ("foo bar" > "a") {
        var v59 uint = 65535
        v40 += 0
        if (((1e10 - 1.5) / (1e10 - ((((float32(v1) + 0.1) - (1.5 / -2.25)) + (0.1 - 0.0)) + 0.1))) <= 0.0) {
            for fi60 := 0; fi60 < 3; fi60++ {
                v61 := GenS2{F0: 129, F1: 1}
                if v5 {
                    var v62 uint = (v59 ^ v59)
                    v63 := GenS8{F0: -1, F1: genFn1((fi60 << ((4294967295) & 63)), fi60), F2: 1}
                    v5 = v5
                    _ = v62
                    _ = v63
                }
                _ = fi60
                _ = v61
            }
        } else {
            var v64 uint64 = 54
            v4 = v4
            for fi65 := 0; fi65 < 1; fi65++ {
                var v66 string = "\n\t"
                _ = fi65
                _ = v66
            }
            _ = v64
        }
        _ = v59
    }
    fmt.Println(v1, v39)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
