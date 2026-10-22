package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int8 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int; F2 int32 }
func genFn0() int16 {
    var v7 string = "hello"
    var v8 uint16 = 1
    fmt.Println(v8, v7)
    if true {
        for fi9 := 0; fi9 < 4; fi9++ {
            if false {
                var v10 float32 = 1.5
                v7 = (v7 + (v7 + "hello"))
                _ = v10
            }
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := []int{(fi11 + fi9)}
                v12 = append(v12, -1)
                _ = fi11
                _ = v12
            }
            _ = fi9
        }
    } else {
        for fi13 := 0; fi13 < 9; fi13++ {
            v8 += v8
            _ = fi13
        }
        if ((-128 != 42) && (("hello" != v7) && false)) {
            v7 = v7
            var v14 uint64 = 0
            var v15 bool = true
            _ = v14
            _ = v15
        }
    }
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int16) int16 {
    if ((65535 >= 303) && (9223372036854775807 == 0)) {
        var v16 int = 42
        _ = v16
    } else {
        var v17 int8 = 127
        v17 = (v17 | (v17 >> ((335) & 63)))
        _ = v17
    }
    fmt.Println(p0)
    var v18 int64 = -9
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(p0, p0)
        v20 := GenS1{F0: 1, F1: v18, F2: 67}
        _ = fi19
        _ = v20
    }
    var v21 int8 = 30
    _ = p0
    _ = v18
    _ = v21
    return genFn0()
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = (0.1 * float64(3.14159))
    var v3 uint16 = 179
    var v4 uint = 558
    var v5 bool = false
    var v6 uint64 = 736
    var v22 uint = (v4 ^ 42)
    v23 := GenS1{F0: v6, F1: 9223372036854775807, F2: 0}
    fmt.Println(v2, v3, v6)
    for fi24 := 0; fi24 < 7; fi24++ {
        switch fi24 {
        case 457:
            _ = len(v1)
        case 9223372036854775807:
            v1["a"] = (("\n\t" + "foo bar") + "hello")
        default:
            v25 := GenS2{F0: 694, F1: 1, F2: (v22 + v4)}
            _ = v25
        }
        var v26 int32 = 1
        _ = fi24
        _ = v26
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v22, v5, v2)
    case "a":
        for fi27 := 0; fi27 < 3; fi27++ {
            for fi28 := 0; fi28 < 4; fi28++ {
                var v29 float64 = 0.0
                _ = fi28
                _ = v29
            }
            v30 := []int{fi27, (-(fi27)), fi27}
            _ = fi27
            _ = v30
        }
    case "\n\t":
        if (!(128 != 141)) {
            if (v5 && (!v5)) {
                if (1 < 1) {
                    v5 = v5
                    var v31 float64 = v2
                    var v32 uint16 = (^(v3))
                    _ = v31
                    _ = v32
                } else {
                    if v5 {
                        if ((v6 >> ((18446744073709551615) & 63)) > 450) {
                            fmt.Println(v2, v4, v22)
                        } else {
                            var v33 uint32 = 42
                            switch "" {
                            case "\n\t":
                                if ((0 <= v6) || v5) {
                                    if v5 {
                                        v34 := GenS3{F0: (v22 * 65535)}
                                        _ = v34
                                    } else {
                                        var v35 uint16 = (v3 ^ (v3 % (((v3 | (v3 % ((2) | 1)))) | 1)))
                                        v5 = (!v5)
                                        _ = v35
                                    }
                                    v6 += (v6 % ((199) | 1))
                                }
                            }
                            if (true && ("\n\t" == ("\n\t" + "foo bar"))) {
                                if v5 {
                                    var v36 int16 = 651
                                    var v37 int16 = 642
                                    v4 += v4
                                    _ = v36
                                    _ = v37
                                }
                                var v38 int16 = 32767
                                _ = v38
                            }
                            _ = v33
                        }
                        v6 = (v6 ^ 2)
                        v39 := GenS3{F0: v4}
                        _ = v39
                    }
                    fmt.Println(v2)
                    v4 = uint(v6)
                }
                var v40 bool = ((-1 >= -1) || false)
                v41 := v1["\n\t"]
                _ = v40
                _ = v41
            } else {
                v3 = (v3 + v3)
            }
        } else {
            var v42 uint16 = 42
            v3 += v3
            _ = v42
        }
    }
    var v43 uint8 = 128
    v44 := []int{1, -9223372036854775808, 506}
    if true {
        v2 = v2
    }
    var v45 uint8 = 42
    v46 := GenS4{F0: ((("hello" + "hello") + ((((("foo bar" + "") + "\n\t") + ("foo bar" + "foo bar")) + "a") + "a")) + ("foo bar" + "")), F1: 1, F2: 1}
    for fi47 := 0; fi47 < 2; fi47++ {
        for fi48 := 0; fi48 < 18; fi48++ {
            var v49 int16 = 1
            if ((!false) || (v2 >= 1.5)) {
                v50 := GenS2{F0: 65535, F1: (v49 | (v49 >> ((v22) & 63))), F2: v4}
                _ = v50
            } else {
                if (!(!v5)) {
                    fmt.Println(v4)
                } else {
                    v6 -= (v6 << ((v22) & 63))
                }
            }
            _ = fi48
            _ = v49
        }
        _ = fi47
    }
    v45 = v43
    switch "\n\t" {
    case "foo bar":
        if (!true) {
            v1["foo bar"] = "\n\t"
            if v5 {
                var v51 uint16 = 196
                _ = v51
            } else {
                v4 = v4
            }
        }
    default:
        var v52 float64 = 1.5
        _ = v52
    }
    if (v5 || (v5 && v5)) {
        v1[""] = "hello"
        v53 := GenS4{F0: (("" + "hello") + ("foo bar" + "hello")), F1: -2147483648, F2: 1}
        _ = v53
    } else {
        fmt.Println(v6)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
