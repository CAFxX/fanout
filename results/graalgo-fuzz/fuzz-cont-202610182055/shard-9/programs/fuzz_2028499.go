package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int32 }
type GenS2 struct { F0 int8; F1 uint; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint; F2 int8 }
func genFn0(p0 int32, p1 int8) int16 {
    v7 := []int{9223372036854775807, 0}
    v8 := GenS1{F0: 42, F1: 1, F2: (p0 ^ p0)}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v9 uint32 = 2147483648
    if len(v7) > 0 {
        v7[0] = 128
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1() int {
    v10 := []int{128, 2147483647}
    for fi11 := 0; fi11 < 18; fi11++ {
        _ = len(v10)
        _ = fi11
    }
    _ = v10
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 42
    var v2 int = 625
    var v3 uint64 = 1
    var v4 int = (v2 << ((45) & 63))
    var v5 int8 = 23
    v6 := make(map[int]bool)
    v12 := GenS1{F0: 177, F1: 1, F2: 2147483647}
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 float32 = 1e10
        fmt.Println(v14, v2)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 uint16 = 255
        for fi17 := 0; fi17 < 1; fi17++ {
            switch ("hello" + "foo bar") {
            case "":
                var v18 float64 = 1.5
                _ = v18
            case "\n\t":
                fmt.Println(v3, v5)
            case "hello":
                var v19 float64 = -2.25
                _ = v19
            default:
                switch fi15 {
                case 32767:
                    v2 = (v2 / (((-(fi17))) | 1))
                }
            }
            v20 := GenS2{F0: v5, F1: 719, F2: (v1 & (^(v1)))}
            _ = fi17
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v5 -= v5
        _ = fi21
    }
    if (!false) {
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 int16 = 0
            switch "a" {
            case "hello":
                if true {
                    v5 -= v5
                }
            case "a":
                v24 := GenS1{F0: 42, F1: 255, F2: 42}
                _ = v24
            case "\n\t":
                var v25 uint16 = (^(v1))
                _ = v25
            default:
                switch (("a" + (("foo bar" + "\n\t") + ("foo bar" + "hello"))) + "") {
                case "":
                    _ = len(v6)
                case "foo bar":
                    for fi26 := 0; fi26 < 9; fi26++ {
                        fmt.Println(v5, v23, v1)
                        v27 := GenS2{F0: (v5 * (v5 * v5)), F1: 194, F2: 65535}
                        v5 -= (v5 ^ (v5 - 1))
                        _ = fi26
                        _ = v27
                    }
                default:
                    var v28 int64 = -9223372036854775808
                    _ = v28
                }
            }
            _ = fi22
            _ = v23
        }
    } else {
        if ((v2 ^ (v4 * v2)) > (v2 | (v4 << ((18446744073709551615) & 63)))) {
            v1 -= v1
        } else {
            switch "" {
            case "a":
                if false {
                    v5 += v5
                    fmt.Println(v3, v3, v1)
                    var v29 uint = 693
                    _ = v29
                } else {
                    v2 -= 32767
                    var v30 uint8 = 128
                    _ = v30
                }
            }
            for fi31 := 0; fi31 < 3; fi31++ {
                fmt.Println(v1)
                if (!(false && (false || true))) {
                    v32 := v6[(v4 >> ((uint(v1)) & 63))]
                    _ = v32
                } else {
                    v33 := GenS1{F0: 521, F1: 0, F2: 321}
                    v34 := GenS3{F0: -103, F1: 4294967295, F2: (v5 | (v5 << ((2) & 63)))}
                    _ = v33
                    _ = v34
                }
                v1 = 255
                _ = fi31
            }
            if true {
                v35 := []int{v4}
                if false {
                    fmt.Println(v2, v3)
                    var v36 int32 = 174
                    for fi37 := 0; fi37 < 18; fi37++ {
                        v36 += (v36 ^ v36)
                        v4 += genFn1()
                        var v38 uint64 = v3
                        _ = fi37
                        _ = v38
                    }
                    _ = v36
                } else {
                    delete(v6, 2147483647)
                    var v39 int16 = -187
                    _ = v39
                }
                v6[(v2 << ((255) & 63))] = (((-2.25 - (0.0 + -2.25)) == 0.0) && false)
                _ = v35
            }
        }
        fmt.Println(v4, v5)
    }
    v40 := GenS3{F0: 2147483647, F1: 18446744073709551615, F2: v5}
    v41 := []int{}
    var v42 uint64 = 1
    v43 := GenS3{F0: -1, F1: 1, F2: v5}
    delete(v6, -9223372036854775808)
    if true {
        v44 := GenS3{F0: -866, F1: 299, F2: (^(v5))}
        var v45 bool = false
        _ = v44
        _ = v45
    } else {
        var v46 bool = true
        _ = v46
    }
    for fi47 := 0; fi47 < 20; fi47++ {
        v2 = v2
        var v48 uint32 = 911
        for fi49 := 0; fi49 < 17; fi49++ {
            var v50 uint32 = 971
            if (!true) {
                var v51 uint32 = 42
                v4 += (v4 ^ (v2 << ((587) & 63)))
                _ = v51
            } else {
                if true {
                    v52 := GenS3{F0: 0, F1: 18446744073709551615, F2: (v5 | 127)}
                    var v53 float64 = 0.1
                    var v54 int = fi47
                    var v55 uint8 = 13
                    var v56 int64 = 1
                    _ = v52
                    _ = v53
                    _ = v54
                    _ = v55
                    _ = v56
                }
            }
            var v57 uint64 = (v42 * (v3 | (v3 >> ((2) & 63))))
            _ = fi49
            _ = v50
            _ = v57
        }
        _ = fi47
        _ = v48
    }
    var v58 uint8 = 255
    switch ((("foo bar" + "foo bar") + "") + ("a" + "foo bar")) {
    case "foo bar":
        switch (v4 % ((v4) | 1)) {
        case 1:
            switch -1 {
            case 127:
                v59 := GenS1{F0: -92, F1: (v58 >> ((2) & 63)), F2: 42}
                _ = v59
            case 2147483647:
                v60 := GenS3{F0: -1, F1: 255, F2: (v5 & 1)}
                _ = v60
            default:
                if false {
                    v58 = v58
                    fmt.Println(v58, v42, v58)
                    fmt.Println(v1)
                } else {
                    v5 -= 100
                    for fi61 := 0; fi61 < 18; fi61++ {
                        v3 += v3
                        _ = fi61
                    }
                    v62 := v6[v2]
                    _ = v62
                }
            }
        case 42:
            v63 := GenS3{F0: 2147483647, F1: 855, F2: v5}
            _ = v63
        case 9223372036854775807:
            if true {
                fmt.Println(v4, v58)
            } else {
                var v64 float32 = (0.0 - (0.1 * 1.5))
                _ = v64
            }
        }
    default:
        v58 = (v58 + (v58 ^ v58))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v58
    _ = v6
}
