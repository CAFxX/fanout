package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int; F1 int16; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int; F2 int32 }
func genFn0(p0 int8, p1 uint, p2 uint8) uint64 {
    p1 += p1
    v6 := GenS1{F0: 663, F1: 315, F2: 4294967295}
    switch 127 {
    case 697:
        p2 = p2
    default:
        v7 := []int{-9223372036854775808}
        _ = v7
    }
    var v8 uint32 = 0
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return 678
}
func genFn1(p0 int, p1 bool, p2 uint16) int {
    p2 -= p2
    if p1 {
        if p1 {
            v9 := []int{p0, (p0 & p0), p0}
            if (true && p1) {
                var v10 uint32 = 42
                _ = v10
            }
            var v11 uint32 = 2
            _ = v11
            _ = v9
        } else {
            p1 = ((0 >= 42) && p1)
            p1 = (0 != 100)
            p1 = false
        }
        var v12 string = ""
        p1 = ((^(p0)) > p0)
        _ = v12
    }
    var v13 uint8 = 42
    switch "" {
    case "\n\t":
        var v14 uint = 42
        _ = v14
    case "":
        v15 := []int{p0, 32768, p0, p0}
        _ = v15
    case "foo bar":
        if p1 {
            fmt.Println(p2, p2, p0)
            var v16 int32 = -310
            for fi17 := 0; fi17 < 10; fi17++ {
                p0 = (p0 & -9223372036854775808)
                var v18 uint8 = (v13 | 248)
                p2 -= p2
                _ = fi17
                _ = v18
            }
            _ = v16
        } else {
            p1 = p1
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 int32 = -665
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 int8 = -1
                    _ = fi21
                    _ = v22
                }
                if (fi19 <= 1) {
                    var v23 string = "\n\t"
                    if p1 {
                        p2 = 311
                        var v24 int8 = 0
                        _ = v24
                    } else {
                        p2 += p2
                    }
                    p1 = (!(!(fi19 > 241)))
                    _ = v23
                }
                _ = fi19
                _ = v20
            }
        }
    }
    var v25 int64 = -256
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v25
    return -1
}

func main() {
    var v1 int16 = 0
    var v2 uint32 = 2147483647
    v3 := make(map[string]uint32)
    var v4 uint16 = 194
    var v5 uint = 1
    switch 42 {
    case -744:
        fmt.Println(v5)
    case 42:
        if false {
            switch 32767 {
            case -666:
                v26 := GenS1{F0: 1, F1: (v2 >> ((v5) & 63)), F2: genFn0(-1, (v5 / (((^(v5))) | 1)), 128)}
                _ = v26
            case 4294967295:
                v5 -= v5
            }
            var v27 string = ("" + "foo bar")
            v27 = "a"
            _ = v27
        }
    }
    v1 += v1
    delete(v3, ("hello" + "hello"))
    if ((v1 * 1) > -32768) {
        switch "\n\t" {
        case "hello":
            var v28 uint64 = 1
            _ = v28
        default:
            v4 -= v4
        }
        if false {
            fmt.Println(v5, v4)
            var v29 float64 = ((((0.0 / 1.0) - 0.0) - 3.14159) * float64(float32((1e10 - 1.5))))
            v30 := v3["a"]
            _ = v29
            _ = v30
        }
        if (!false) {
            for fi31 := 0; fi31 < 2; fi31++ {
                switch "\n\t" {
                case "":
                    v3[(("hello" + ("\n\t" + "foo bar")) + "a")] = (^(v2))
                case "hello":
                    if true {
                        v5 = uint(v4)
                    }
                default:
                    v1 = (v1 - (^(v1)))
                }
                for fi32 := 0; fi32 < 4; fi32++ {
                    fmt.Println(fi32, v4)
                    fmt.Println(v5, v5)
                    if true {
                        var v33 int16 = (v1 | (v1 & (v1 * v1)))
                        var v34 int8 = 42
                        _ = v33
                        _ = v34
                    }
                    _ = fi32
                }
                _ = fi31
            }
            var v35 uint16 = v4
            _ = v35
        } else {
            var v36 float32 = 3.14159
            for fi37 := 0; fi37 < 3; fi37++ {
                v38 := GenS1{F0: (-(v1)), F1: v2, F2: 18446744073709551615}
                var v39 string = ""
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = v36
        }
    }
    v40 := GenS2{F0: genFn1(-501, false, v4), F1: v1, F2: v1}
    var v41 uint32 = v2
    if ((v5 << ((0) & 63)) >= 414) {
        var v42 int16 = -1
        v43 := GenS3{F0: 515}
        _ = v42
        _ = v43
    }
    var v44 float32 = ((0.1 - (0.1 * 1.5)) - -2.25)
    for fi45 := 0; fi45 < 17; fi45++ {
        v46 := GenS4{F0: (("a" + ("foo bar" + (("a" + ("hello" + ("foo bar" + "a"))) + "foo bar"))) + (("foo bar" + "\n\t") + "\n\t")), F1: (^(fi45)), F2: -1}
        var v47 uint32 = 519
        fmt.Println(v41, v44)
        _ = fi45
        _ = v46
        _ = v47
    }
    fmt.Println(v4)
    fmt.Println(v2, v2)
    switch -81 {
    case -9223372036854775808:
        v3["hello"] = v41
    case 449:
        var v48 uint8 = 33
        _ = v48
    default:
        fmt.Println(v41, v2, v4)
    }
    switch -584 {
    case 1:
        for fi49 := 0; fi49 < 6; fi49++ {
            v50 := v3[""]
            _ = fi49
            _ = v50
        }
    case 32768:
        v4 -= (v4 / ((v4) | 1))
    case 502:
        v51 := GenS1{F0: v1, F1: v2, F2: 42}
        _ = v51
    }
    switch "a" {
    case "foo bar":
        for fi52 := 0; fi52 < 20; fi52++ {
            fmt.Println(v4)
            switch "" {
            case "":
                if (v5 < v5) {
                    v5 = (v5 >> (((^(v5))) & 63))
                }
            }
            _ = fi52
        }
    case "":
        var v53 int16 = (-(v1))
        _ = v53
    default:
        v54 := GenS1{F0: v1, F1: (v41 - (^(v41))), F2: 538}
        _ = v54
    }
    v55 := []int{32768, -9223372036854775808, 1}
    fmt.Println(v2, v41)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v55:", len(v55), v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v5
    _ = v55
}
