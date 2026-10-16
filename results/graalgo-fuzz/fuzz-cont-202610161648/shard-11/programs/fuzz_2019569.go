package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint64 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int8; F1 uint; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint32 }
func genFn0() uint16 {
    v7 := []int{127, -762, 0, -1}
    var v8 string = "\n\t"
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 string = (v8 + "hello")
        fmt.Println(fi9, v10, v8)
        _ = fi9
        _ = v10
    }
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 string) uint8 {
    if true {
        fmt.Println(p0)
    }
    fmt.Println(p0)
    p0 = ""
    switch -818 {
    case 42:
        fmt.Println(p0)
    case -810:
        var v11 int = 0
        _ = v11
    default:
        switch p0 {
        case "a":
            p0 = (p0 + p0)
        case "foo bar":
            p0 = p0
        }
    }
    _ = p0
    return 103
}

func main() {
    var v1 int8 = -1
    var v2 uint16 = 255
    var v3 int = 9223372036854775807
    var v4 float64 = -2.25
    var v5 float64 = (v4 / -2.25)
    v6 := make(map[int]string)
    for fi12 := 0; fi12 < 7; fi12++ {
        if true {
            if false {
                v13 := GenS1{F0: false, F1: v1, F2: 4294967295}
                var v14 uint = 597
                v1 -= 0
                _ = v13
                _ = v14
            } else {
                var v15 uint16 = v2
                switch ((("" + "") + "") + ("\n\t" + ((((("a" + "\n\t") + ("a" + "foo bar")) + "foo bar") + "\n\t") + "a"))) {
                case "a":
                    v2 += v15
                case "foo bar":
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v17 := GenS2{F0: false, F1: true}
                        v18 := GenS2{F0: true, F1: (v15 >= 17)}
                        fmt.Println(v3, fi16)
                        fmt.Println(v3, v1)
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
                if (!(18446744073709551615 >= 4294967295)) {
                    if true {
                        var v19 uint64 = 352
                        if false {
                            v19 += 42
                        }
                        var v20 int = (fi12 % ((fi12) | 1))
                        _ = v19
                        _ = v20
                    }
                    v15 -= (v2 - v2)
                    v21 := v6[127]
                    _ = v21
                }
                _ = v15
            }
            for fi22 := 0; fi22 < 5; fi22++ {
                v23 := GenS2{F0: (-335 < 2147483647), F1: false}
                v5 = v5
                var v24 int64 = -1
                _ = fi22
                _ = v23
                _ = v24
            }
            v25 := GenS1{F0: true, F1: v1, F2: 656}
            _ = v25
        } else {
            var v26 int16 = 149
            var v27 uint64 = 450
            var v28 float32 = 0.0
            _ = v26
            _ = v27
            _ = v28
        }
        v29 := GenS1{F0: false, F1: v1, F2: 4294967295}
        _ = fi12
        _ = v29
    }
    switch "hello" {
    case "\n\t":
        if (!false) {
            v3 = (v3 / ((2147483647) | 1))
        }
    default:
        v2 = 32767
    }
    if true {
        fmt.Println(v1)
    } else {
        var v30 float64 = v4
        _ = v30
    }
    v31 := v6[v3]
    v31 = "\n\t"
    switch "hello" {
    case "a":
        var v32 int32 = 0
        _ = v32
    default:
        if false {
            for fi33 := 0; fi33 < 15; fi33++ {
                switch v3 {
                case -2147483648:
                    v34 := GenS1{F0: (42 <= v2), F1: v1, F2: 2}
                    _ = v34
                case 0:
                    v5 -= ((v4 / (v4 - 3.14159)) / (1.5 - (v4 / (v4 / -2.25))))
                }
                v35 := GenS3{F0: int8(v2), F1: 65535, F2: 899}
                var v36 int = (fi33 - 38)
                _ = fi33
                _ = v35
                _ = v36
            }
        }
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        fmt.Println(v5)
        _ = fi37
    }
    v38 := v6[32768]
    if (245 == 1) {
        v5 -= (v5 / v4)
        if false {
            switch v38 {
            case "a":
                for fi39 := 0; fi39 < 2; fi39++ {
                    delete(v6, (v3 % ((2147483647) | 1)))
                    _ = fi39
                }
            default:
                for fi40 := 0; fi40 < 10; fi40++ {
                    var v41 int8 = v1
                    v42 := GenS3{F0: v1, F1: 738, F2: -1}
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
        }
        v1 -= v1
    }
    v43 := GenS3{F0: v1, F1: 26, F2: -32768}
    fmt.Println(v5)
    if false {
        v31 = ((v38 + ("" + (("" + v38) + "hello"))) + "foo bar")
    } else {
        if false {
            fmt.Println(v3)
            var v44 int8 = v1
            for fi45 := 0; fi45 < 7; fi45++ {
                v46 := []int{(fi45 + v3)}
                var v47 int = v3
                v38 = "\n\t"
                _ = fi45
                _ = v46
                _ = v47
            }
            _ = v44
        } else {
            if (false && false) {
                var v48 int32 = -1
                switch ((v38 + v38) + v31) {
                case "\n\t":
                    if ((!((false && (genFn1("") >= 167)) && ((13 == 985) && true))) || (((-54 > v1) && (105 == 0)) && true)) {
                        fmt.Println(v48, v5, v31)
                    }
                default:
                    for fi49 := 0; fi49 < 13; fi49++ {
                        v50 := []int{(v3 << ((42) & 63))}
                        fmt.Println(v2)
                        v3 = -2147483648
                        _ = fi49
                        _ = v50
                    }
                }
                _ = v48
            }
            v51 := GenS2{F0: false, F1: (127 < 128)}
            for fi52 := 0; fi52 < 1; fi52++ {
                v31 = v31
                _ = fi52
            }
            _ = v51
        }
    }
    if (1 != 0) {
        fmt.Println(v4, v3)
    } else {
        fmt.Println(v31, v4, v2)
        v5 -= v5
    }
    var v53 int16 = -1
    v54 := GenS4{F0: 0, F1: 395}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
