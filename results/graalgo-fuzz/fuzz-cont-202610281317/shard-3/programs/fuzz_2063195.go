package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint8; F2 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 uint }
type GenS9 struct { F0 int8; F1 uint8; F2 int }
func genFn0() int32 {
    if true {
        v7 := GenS1{F0: 39}
        v8 := []int{-9223372036854775808, 128}
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 int = (^(fi9))
            v10 = -835
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    } else {
        v11 := GenS2{F0: 83}
        for fi12 := 0; fi12 < 16; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 bool = (!true)
                v14 = ((118 < 31) && (!v14))
                fmt.Println(v14, fi13, fi13)
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        var v15 bool = (!(!true))
        _ = v11
        _ = v15
    }
    var v16 uint64 = 4294967295
    var v17 float32 = 0.1
    if (!true) {
        v18 := GenS3{F0: -920}
        fmt.Println(v17, v16)
        for fi19 := 0; fi19 < 7; fi19++ {
            for fi20 := 0; fi20 < 11; fi20++ {
                if false {
                    var v21 uint32 = 2147483647
                    v22 := []int{fi20, (fi19 / ((fi20) | 1))}
                    fmt.Println(v21, v21)
                    _ = v21
                    _ = v22
                }
                switch "\n\t" {
                case "foo bar":
                    var v23 int8 = -42
                    _ = v23
                default:
                    v17 += v17
                }
                var v24 bool = true
                _ = fi20
                _ = v24
            }
            v25 := []int{(fi19 ^ -2147483648), fi19}
            fmt.Println(v17, v17)
            _ = fi19
            _ = v25
        }
        _ = v18
    } else {
        if true {
            if true {
                if (((!("hello" <= "foo bar")) && (-193 == -2147483648)) && (5 < 1000)) {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        if true {
                            if (!false) {
                                v27 := []int{(fi26 | 4294967295)}
                                var v28 int32 = 42
                                var v29 int64 = -9223372036854775808
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                        } else {
                            switch (fi26 >> ((109) & 63)) {
                            case 32768:
                                fmt.Println(v16, v17, fi26)
                            default:
                                var v30 bool = ("hello" != "a")
                                _ = v30
                            }
                            v31 := []int{-693, fi26}
                            _ = v31
                        }
                        var v32 float64 = (0.0 + ((1.5 + float64(v17)) + 0.1))
                        _ = fi26
                        _ = v32
                    }
                    if true {
                        v17 = v17
                    } else {
                        v33 := GenS3{F0: -113}
                        _ = v33
                    }
                    v34 := GenS4{F0: 125, F1: 32767}
                    _ = v34
                }
            } else {
                var v35 uint16 = 0
                if (343 < 1) {
                    fmt.Println(v35, v17)
                }
                _ = v35
            }
            if (v16 > v16) {
                for fi36 := 0; fi36 < 7; fi36++ {
                    v37 := GenS1{F0: 2}
                    v16 -= v16
                    _ = fi36
                    _ = v37
                }
            }
            v38 := GenS5{F0: -1}
            _ = v38
        }
        for fi39 := 0; fi39 < 13; fi39++ {
            v40 := []int{fi39, fi39, fi39, fi39}
            var v41 int64 = 1
            _ = fi39
            _ = v40
            _ = v41
        }
        for fi42 := 0; fi42 < 19; fi42++ {
            v16 = (v16 ^ v16)
            _ = fi42
        }
    }
    _ = v16
    _ = v17
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 42
    v3 := []int{-1, (v2 % ((v2) | 1))}
    var v4 float64 = -2.25
    var v5 bool = ((((-2.25 / 3.14159) < 1.5) || true) && (9223372036854775807 != 42))
    v6 := make(map[string]int)
    fmt.Println(v2, v5)
    fmt.Println(v2, v2, v2)
    switch v2 {
    case 2147483647:
        fmt.Println(v5)
    case 128:
        v43 := GenS3{F0: 0}
        _ = v43
    default:
        if true {
            var v44 uint32 = 42
            fmt.Println(v44, v4, v5)
            _ = v44
        } else {
            v4 += 0.0
        }
    }
    if v5 {
        var v45 uint32 = 0
        _ = v45
    } else {
        for fi46 := 0; fi46 < 19; fi46++ {
            switch (fi46 ^ fi46) {
            case 882:
                var v47 bool = v5
                _ = v47
            default:
                v5 = true
            }
            v2 += v2
            _ = fi46
        }
    }
    var v48 int = (^(v2))
    v49 := v1["a"]
    var v50 int16 = -83
    for fi51 := 0; fi51 < 20; fi51++ {
        switch 128 {
        case -602:
            var v52 uint = 52
            _ = v52
        }
        for fi53 := 0; fi53 < 8; fi53++ {
            if v5 {
                v54 := GenS6{F0: 0, F1: 128, F2: (v50 / ((522) | 1))}
                _ = v54
            } else {
                v55 := GenS7{F0: v49}
                v56 := GenS6{F0: -12, F1: 3, F2: (v50 << ((0) & 63))}
                _ = v55
                _ = v56
            }
            v48 -= (fi51 ^ (fi51 / (((fi51 - (-(v48)))) | 1)))
            _ = fi53
        }
        _ = fi51
    }
    if v49 {
        v48 += 918
        switch v48 {
        case -512:
            v49 = false
        default:
            fmt.Println(v2, v48)
        }
    } else {
        if ((((3.14159 * 1.5) * (float32(v4) - 1e10)) + (1e10 + float32(v4))) < 1e10) {
            if (42 < 1) {
                v4 = (v4 - (((v4 + v4) - v4) * ((3.14159 / v4) - -2.25)))
                v57 := GenS8{F0: 6, F1: 1}
                if len(v3) > 0 {
                    v3[0] = v2
                }
                _ = v57
            } else {
                v49 = v5
                var v58 int = 791
                _ = v58
            }
            var v59 float64 = float64(float32(((float64(0.0) / v4) + (v4 - -2.25))))
            _ = v59
        }
        for fi60 := 0; fi60 < 6; fi60++ {
            v61 := GenS9{F0: -93, F1: 221, F2: fi60}
            _ = fi60
            _ = v61
        }
        var v62 float32 = 1e10
        _ = v62
    }
    v48 += (v48 << ((2) & 63))
    if (("" + ("a" + "a")) == "\n\t") {
        var v63 bool = (!true)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v64 := v6[((("a" + "\n\t") + "hello") + "a")]
        _ = v63
        _ = v64
    } else {
        var v65 uint8 = 128
        _ = v65
    }
    delete(v1, "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
