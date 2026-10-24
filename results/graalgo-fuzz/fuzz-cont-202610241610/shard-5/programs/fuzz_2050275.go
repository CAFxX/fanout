package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 string; F1 uint64; F2 uint64 }
func genFn0(p0 int64, p1 bool, p2 uint16) int {
    var v6 int64 = (p0 << ((750) & 63))
    v7 := []int{42, 127, -742, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -568
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"foo bar", "a"}
    v3 := make(map[string]string)
    var v4 int64 = 42
    v5 := make(map[int]uint32)
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int16 = 145
        v5[genFn0(v4, (true || (4294967295 >= 65535)), 255)] = 545
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 int64 = v4
            if false {
                v12 := GenS1{F0: (fi8 ^ (fi8 << ((2) & 63)))}
                v9 += 0
                if true {
                    fmt.Println(fi8, fi8)
                    fmt.Println(fi10)
                } else {
                    if false {
                        var v13 uint16 = 42
                        switch genFn0(v4, (((v4 & (v4 % (((^(v4))) | 1))) > 9223372036854775807) && true), v13) {
                        case 32768:
                            v2 = append(v2, ("\n\t" + ""))
                        case -347:
                            if (!((v9 & (v9 >> ((1) & 63))) < (v9 % ((v9) | 1)))) {
                                v14 := GenS2{F0: 2147483648}
                                _ = v14
                            }
                        case -9223372036854775808:
                            switch ((("foo bar" + "a") + (("" + "") + "hello")) + ("hello" + ("" + "hello"))) {
                            case "":
                                v11 -= (-(v4))
                            }
                        default:
                            fmt.Println(v9, v13, v11)
                        }
                        var v15 float32 = (0.1 / (0.1 * 3.14159))
                        _ = v13
                        _ = v15
                    } else {
                        if (v4 <= (v11 & (v11 << ((936) & 63)))) {
                            var v16 uint16 = 1
                            var v17 uint32 = 586
                            _ = v16
                            _ = v17
                        } else {
                            v18 := GenS2{F0: 2}
                            var v19 int = (fi8 / ((fi10) | 1))
                            if ((v9 >> ((0) & 63)) > v9) {
                                v4 += v4
                                switch ((("a" + "") + ("foo bar" + (("a" + ("\n\t" + "")) + ("\n\t" + ("foo bar" + "a"))))) + ("\n\t" + (("" + "\n\t") + "foo bar"))) {
                                case "\n\t":
                                    v20 := GenS3{F0: 65535, F1: 997}
                                    _ = v20
                                }
                                v21 := GenS2{F0: 42}
                                _ = v21
                            } else {
                                var v22 uint64 = 4294967295
                                v2 = append(v2, "\n\t")
                                if true {
                                    fmt.Println(v11)
                                    v23 := GenS2{F0: 1}
                                    _ = v23
                                } else {
                                    v11 += 9223372036854775807
                                    v24 := GenS1{F0: (-(fi10))}
                                    v4 += -63
                                    _ = v24
                                }
                                v25 := GenS4{F0: ("\n\t" + ("\n\t" + "hello")), F1: v22, F2: v22}
                                v22 += v22
                                _ = v22
                                _ = v25
                            }
                            v9 = v9
                            if (true || (true || (1e10 > 3.14159))) {
                                v19 = (v19 >> ((4294967295) & 63))
                                v9 = (v9 % ((v9) | 1))
                            }
                            fmt.Println(fi8, v9)
                            _ = v18
                            _ = v19
                        }
                        _ = len(v2)
                        fmt.Println(fi10, v9)
                    }
                }
                _ = v12
            }
            v11 = 42
            fmt.Println(v4)
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v26 := GenS4{F0: (("\n\t" + "") + "\n\t"), F1: 2, F2: 0}
    v27 := GenS4{F0: (("hello" + ("hello" + ("" + (("foo bar" + "\n\t") + "a")))) + "\n\t"), F1: 0, F2: 2}
    for fi28 := 0; fi28 < 18; fi28++ {
        switch "foo bar" {
        case "foo bar":
            if (-1 <= 1) {
                v4 = v4
                switch (fi28 - -556) {
                case -2147483648:
                    if (255 != 925) {
                        var v29 int16 = 845
                        if true {
                            v30 := GenS4{F0: "hello", F1: 0, F2: 536}
                            _ = v30
                        }
                        _ = v29
                    }
                }
            } else {
                switch genFn0(v4, true, 1) {
                case -1:
                    v2 = append(v2, "")
                case -138:
                    for fi31 := 0; fi31 < 14; fi31++ {
                        fmt.Println(fi28, fi28, v4)
                        var v32 string = ("a" + "hello")
                        _ = fi31
                        _ = v32
                    }
                default:
                    for fi33 := 0; fi33 < 20; fi33++ {
                        v34 := GenS1{F0: genFn0(v4, (0.1 > 1.5), 255)}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v2 = append(v2, "hello")
                        fmt.Println(v4, fi28)
                        _ = fi33
                        _ = v34
                    }
                }
            }
        case "a":
            v35 := GenS4{F0: "", F1: 2, F2: 2}
            _ = v35
        }
        var v36 uint8 = 255
        for fi37 := 0; fi37 < 16; fi37++ {
            var v38 string = "hello"
            _ = fi37
            _ = v38
        }
        _ = fi28
        _ = v36
    }
    _ = len(v5)
    v4 = (v4 >> ((0) & 63))
    v39 := GenS2{F0: 42}
    var v40 uint16 = 255
    fmt.Println(v4)
    fmt.Println(v4)
    v41 := GenS2{F0: 65535}
    fmt.Println(v40)
    var v42 uint16 = (^(v40))
    var v43 float64 = -2.25
    var v44 int8 = -100
    for fi45 := 0; fi45 < 11; fi45++ {
        if true {
            fmt.Println(v43)
        } else {
            v43 += 3.14159
        }
        fmt.Println(v4)
        _ = fi45
    }
    var v46 int = 9223372036854775807
    if true {
        switch ("foo bar" + ((("" + "\n\t") + ("\n\t" + ("hello" + "hello"))) + ("\n\t" + ("a" + "\n\t")))) {
        case "foo bar":
            var v47 int = 4294967295
            _ = v47
        case "a":
            var v48 uint32 = 0
            _ = v48
        default:
            var v49 uint = 350
            _ = v49
        }
        var v50 int64 = v4
        v43 -= float64((1.5 * (0.1 + (0.1 / 0.1))))
        _ = v50
    } else {
        v4 = v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v46
    _ = v5
}
