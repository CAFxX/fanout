package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 uint64; F2 uint16 }
func genFn0() int {
    if true {
        var v7 string = "hello"
        fmt.Println(v7)
        _ = v7
    }
    v8 := []int{0, 783, 169}
    v9 := GenS1{F0: 65535, F1: 802}
    var v10 int64 = 1
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            v10 += (v10 - v10)
            v8 = append(v8, (-(fi11)))
            if false {
                v13 := GenS1{F0: 2, F1: 253}
                var v14 int = (fi11 | (fi11 << ((65535) & 63)))
                v14 = 0
                _ = v13
                _ = v14
            }
            _ = fi12
        }
        _ = fi11
    }
    _ = v10
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1() int16 {
    var v15 bool = false
    var v16 float64 = 3.14159
    v17 := []int{-429}
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 6; fi19++ {
            if v15 {
                var v20 string = ("hello" + ("a" + "hello"))
                if (-128 <= 0) {
                    var v21 float64 = (v16 / v16)
                    _ = v21
                }
                var v22 bool = (255 <= 128)
                _ = v20
                _ = v22
            } else {
                var v23 int64 = -9223372036854775808
                switch ("hello" + (("" + (("a" + "") + ("hello" + "foo bar"))) + ("a" + ""))) {
                case "foo bar":
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                case "hello":
                    var v24 float32 = 1e10
                    _ = v24
                }
                _ = v23
            }
            if true {
                if (0 < -98) {
                    var v25 uint16 = 32767
                    _ = v25
                } else {
                    v17 = append(v17, (fi19 << ((18446744073709551615) & 63)))
                    v26 := GenS2{F0: 214, F1: 9223372036854775807}
                    _ = v26
                }
            } else {
                v15 = v15
                fmt.Println(v15)
                if (1 <= 1) {
                    _ = len(v17)
                } else {
                    var v27 int64 = -929
                    fmt.Println(fi19, v15, v16)
                    v16 = 0.0
                    var v28 int = (fi18 << ((2) & 63))
                    if len(v17) > 0 {
                        v17[0] = v28
                    }
                    _ = len(v17)
                    v29 := GenS3{F0: 881}
                    _ = v27
                    _ = v28
                    _ = v29
                }
                if (1 == 42) {
                    if v15 {
                        var v30 uint8 = 42
                        var v31 int64 = -1
                        var v32 float64 = (0.1 / v16)
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    fmt.Println(fi18, v15)
                    v33 := GenS2{F0: 2, F1: 806}
                    _ = v33
                }
            }
            var v34 uint = 16
            _ = fi19
            _ = v34
        }
        fmt.Println(v16, fi18)
        _ = fi18
    }
    _ = v15
    _ = v16
    _ = v17
    return -163
}

func main() {
    var v1 uint = 255
    v2 := []string{"foo bar", "foo bar", ""}
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    var v5 float32 = (1e10 / 1.5)
    v6 := make(map[int]bool)
    for fi35 := 0; fi35 < 6; fi35++ {
        for fi36 := 0; fi36 < 4; fi36++ {
            v1 = (^(v1))
            _ = fi36
        }
        _ = fi35
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        for fi38 := 0; fi38 < 12; fi38++ {
            v5 += 3.14159
            _ = fi38
        }
        _ = len(v3)
        for fi39 := 0; fi39 < 16; fi39++ {
            var v40 int = 681
            _ = fi39
            _ = v40
        }
        _ = fi37
    }
    var v41 int = 4294967295
    if false {
        v5 = (1e10 * (1e10 + v5))
    } else {
        for fi42 := 0; fi42 < 8; fi42++ {
            fmt.Println(v41, v41, v5)
            _ = fi42
        }
    }
    fmt.Println(v41, v41, v41)
    v5 += (float32((1e10 + -2.25)) + -2.25)
    var v43 int8 = 127
    var v44 uint = 434
    _ = len(v4)
    delete(v3, v41)
    v5 = (v5 / v5)
    fmt.Println(v1)
    var v45 int64 = 0
    for fi46 := 0; fi46 < 18; fi46++ {
        switch ("hello" + "foo bar") {
        case "a":
            for fi47 := 0; fi47 < 15; fi47++ {
                v45 = (v45 & 42)
                var v48 int64 = int64(fi47)
                _ = fi47
                _ = v48
            }
        case "foo bar":
            v49 := GenS4{F0: 42}
            _ = v49
        case "":
            for fi50 := 0; fi50 < 19; fi50++ {
                delete(v6, 42)
                _ = fi50
            }
        default:
            for fi51 := 0; fi51 < 13; fi51++ {
                var v52 float64 = 1.5
                switch v41 {
                case 2147483647:
                    _ = len(v2)
                case -774:
                    v41 -= (fi46 / ((2147483647) | 1))
                default:
                    v2 = append(v2, "")
                }
                _ = fi51
                _ = v52
            }
        }
        if false {
            var v53 bool = true
            var v54 int = genFn0()
            v55 := GenS5{F0: -9223372036854775808, F1: 4294967295, F2: 32767}
            _ = v53
            _ = v54
            _ = v55
        } else {
            v56 := GenS1{F0: 231, F1: 125}
            v57 := GenS4{F0: 1}
            _ = v56
            _ = v57
        }
        _ = fi46
    }
    for fi58 := 0; fi58 < 5; fi58++ {
        var v59 uint8 = 2
        var v60 bool = false
        _ = fi58
        _ = v59
        _ = v60
    }
    v5 = 1e10
    if (0 == 42) {
        v41 -= v41
        fmt.Println(v45)
    } else {
        v61 := v6[(v41 % (((-(v41))) | 1))]
        v45 += v45
        switch v41 {
        case -1:
            if (42 != 390) {
                v43 = (v43 / (((^(v43))) | 1))
                switch "" {
                case "a":
                    var v62 string = "hello"
                    _ = v62
                case "hello":
                    var v63 int16 = -32768
                    _ = v63
                default:
                    var v64 uint8 = 42
                    _ = v64
                }
            }
        case 961:
            v45 = v45
        default:
            v65 := GenS3{F0: 842}
            _ = v65
        }
        _ = v61
    }
    fmt.Println(v43)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
