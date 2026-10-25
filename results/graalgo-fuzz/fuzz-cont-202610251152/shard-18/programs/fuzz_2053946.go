package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint16; F1 int32; F2 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 int16; F2 int8 }
func genFn0(p0 int) uint16 {
    v4 := []int{(-(p0))}
    v5 := GenS1{F0: 100, F1: -518}
    p0 = 4294967295
    var v6 string = "\n\t"
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 925
}
func genFn1(p0 int16, p1 string) uint64 {
    if (!(2 <= 127)) {
        p1 = ("hello" + (p1 + p1))
        var v7 bool = (false || (0.1 <= (0.1 + ((float64(0.0) - 0.1) * (((-2.25 * (float64((0.1 * (-2.25 / 1e10))) / 1.5)) * float64(3.14159)) / -2.25)))))
        _ = v7
    } else {
        if (((1 <= 0) || true) && false) {
            switch p1 {
            case "hello":
                v8 := []int{}
                _ = v8
            case "a":
                if (1.5 != (0.1 / float64(-2.25))) {
                    v9 := []int{}
                    var v10 bool = (!true)
                    _ = v10
                    _ = v9
                } else {
                    var v11 uint = 2
                    var v12 bool = true
                    if ((("" + "a") + p1) != (p1 + p1)) {
                        v12 = (42 != 32767)
                        v13 := []int{}
                        _ = v13
                    }
                    _ = v11
                    _ = v12
                }
            case "":
                for fi14 := 0; fi14 < 9; fi14++ {
                    if ((3.14159 / (float32(0.1) * (1e10 * (3.14159 / (float32(1.5) / float32(((0.0 + 3.14159) / (-2.25 / 1e10)))))))) < -2.25) {
                        v15 := GenS2{F0: "a", F1: (-(p0))}
                        v16 := []int{}
                        v17 := GenS2{F0: p1, F1: p0}
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
            default:
                if false {
                    switch (p1 + "a") {
                    case "hello":
                        switch (p1 + p1) {
                        case "hello":
                            p0 -= (p0 << ((65535) & 63))
                        default:
                            var v18 bool = (1 != 65535)
                            _ = v18
                        }
                    case "\n\t":
                        v19 := []int{128, 42}
                        _ = v19
                    default:
                        v20 := GenS3{F0: 6, F1: -2147483648}
                        _ = v20
                    }
                    var v21 string = (((p1 + p1) + p1) + p1)
                    p1 = ("a" + v21)
                    _ = v21
                } else {
                    p1 = "hello"
                }
            }
            var v22 int32 = 1
            p1 = "foo bar"
            _ = v22
        }
    }
    p0 -= -956
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 32767
    var v3 int64 = 1
    if false {
        var v23 uint64 = 2
        _ = v23
    } else {
        if true {
            switch "" {
            case "foo bar":
                v24 := []int{42, 951, 32767}
                _ = v24
            case "\n\t":
                v25 := GenS3{F0: 65535, F1: 0}
                _ = v25
            }
            for fi26 := 0; fi26 < 14; fi26++ {
                v1[""] = 2
                var v27 bool = (false && false)
                v28 := []int{}
                _ = fi26
                _ = v27
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 bool = ((42 >= 42) && true)
            var v31 int = fi29
            _ = fi29
            _ = v30
            _ = v31
        }
        if false {
            if true {
                if true {
                    for fi32 := 0; fi32 < 19; fi32++ {
                        for fi33 := 0; fi33 < 7; fi33++ {
                            v34 := GenS4{F0: (v2 >> ((0) & 63)), F1: -1, F2: (genFn1(-200, ("" + ("hello" + "hello"))) > 586)}
                            var v35 bool = true
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        v3 = (-(v3))
                        fmt.Println(v3, v3, fi32)
                        _ = fi32
                    }
                } else {
                    var v36 uint32 = 532
                    _ = v36
                }
                v37 := GenS2{F0: "hello", F1: 1}
                v38 := GenS2{F0: ("foo bar" + "a"), F1: 1}
                _ = v37
                _ = v38
            } else {
                for fi39 := 0; fi39 < 9; fi39++ {
                    var v40 float64 = ((1e10 / float64(1e10)) + 0.1)
                    var v41 int64 = v3
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                v2 = v2
                var v42 int8 = -69
                _ = v42
            }
        } else {
            v43 := GenS5{F0: 1}
            _ = v43
        }
    }
    v3 = v3
    v2 = 1
    fmt.Println(v3, v3)
    v2 = (v2 ^ 0)
    var v44 uint = 924
    v2 = v2
    var v45 uint = 42
    if ((((-2.25 / (0.0 + (float64(1e10) - (float64((1e10 - ((-2.25 + -2.25) + 1e10))) * float64(1e10))))) / 1.5) * (float64(1e10) * 0.1)) >= 1e10) {
        if (424 < 18446744073709551615) {
            v46 := GenS6{F0: 4294967295, F1: 42, F2: 95}
            v44 = (v44 >> ((v45) & 63))
            _ = v46
        } else {
            for fi47 := 0; fi47 < 19; fi47++ {
                if (!false) {
                    v2 -= v2
                } else {
                    v2 = (v2 ^ v2)
                    v2 -= 255
                }
                v45 -= 65535
                v48 := v1["\n\t"]
                _ = fi47
                _ = v48
            }
            v2 -= (^(v2))
        }
        var v49 uint16 = v2
        if (true || (!(1e10 < -2.25))) {
            v50 := GenS4{F0: (v2 << ((v45) & 63)), F1: -950, F2: (127 <= 32767)}
            var v51 uint = (^(v44))
            _ = v50
            _ = v51
        }
        _ = v49
    } else {
        v52 := []int{32768, -451, -2147483648, -1}
        var v53 int32 = -864
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 2; fi54++ {
        switch ("a" + "\n\t") {
        case "hello":
            var v55 int = fi54
            _ = v55
        case "a":
            fmt.Println(v3)
        }
        var v56 uint16 = (v2 & v2)
        delete(v1, ("foo bar" + ("" + "\n\t")))
        _ = fi54
        _ = v56
    }
    v45 += v44
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
}
