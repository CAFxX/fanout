package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0(p0 int, p1 int16, p2 uint16) bool {
    if ((p2 >> ((2) & 63)) == 485) {
        p1 += p1
        if (!(75 >= 0)) {
            v4 := []int{0}
            var v5 float64 = (float64((-2.25 - float32(0.0))) * 1.5)
            for fi6 := 0; fi6 < 2; fi6++ {
                v5 -= v5
                for fi7 := 0; fi7 < 18; fi7++ {
                    v8 := GenS1{F0: 2}
                    if ((!true) && (42 > 18446744073709551615)) {
                        var v9 string = ""
                        var v10 int64 = 932
                        _ = v10
                        _ = v9
                    } else {
                        var v11 int8 = 100
                        if false {
                            if false {
                                p0 -= 724
                                if false {
                                    var v12 uint = 65535
                                    _ = v12
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            } else {
                                switch "" {
                                case "\n\t":
                                    p0 -= 389
                                default:
                                    if false {
                                        switch 127 {
                                        case 42:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        case 9223372036854775807:
                                            var v13 int64 = 42
                                            _ = v13
                                        case 339:
                                            v14 := GenS1{F0: 2}
                                            _ = v14
                                        }
                                    }
                                }
                                v15 := GenS1{F0: 42}
                                _ = v15
                            }
                            p2 = p2
                        }
                        switch (^(p0)) {
                        case 308:
                            v4 = append(v4, fi7)
                        case 522:
                            p0 -= (p0 << ((0) & 63))
                        case 708:
                            switch "\n\t" {
                            case "\n\t":
                                var v16 uint32 = 65535
                                _ = v16
                            case "hello":
                                p0 = (fi6 ^ fi7)
                            case "a":
                                v11 = (v11 << ((1) & 63))
                            }
                        }
                        _ = v11
                    }
                    switch "a" {
                    case "":
                        p0 = p0
                    case "\n\t":
                        var v17 int32 = -481
                        _ = v17
                    }
                    _ = fi7
                    _ = v8
                }
                _ = fi6
            }
            _ = v4
            _ = v5
        } else {
            p1 += int16(p0)
        }
    }
    p2 = 66
    p1 = 625
    v18 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return (!false)
}
func genFn1(p0 uint32, p1 int16) uint16 {
    var v19 int16 = p1
    if true {
        v20 := GenS1{F0: p0}
        p0 += p0
        _ = v20
    } else {
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := []int{(-(fi21))}
            _ = fi21
            _ = v22
        }
    }
    v23 := []int{-241, -757, -177, -906}
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 13; fi25++ {
            if (!true) {
                v26 := GenS1{F0: p0}
                var v27 int32 = 734
                _ = v26
                _ = v27
            }
            _ = fi25
        }
        _ = fi24
    }
    fmt.Println(v19, v19)
    _ = p0
    _ = p1
    _ = v19
    _ = v23
    return 824
}

func main() {
    var v1 uint16 = 2
    v2 := []string{""}
    var v3 float32 = 1.5
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    switch "foo bar" {
    case "":
        v1 += (^(v1))
    case "foo bar":
        var v28 bool = false
        _ = v28
    default:
        if ((!(!false)) && true) {
            for fi29 := 0; fi29 < 8; fi29++ {
                for fi30 := 0; fi30 < 8; fi30++ {
                    if (!(!((0 <= 0) && (1.5 >= 3.14159)))) {
                        fmt.Println(v3)
                    } else {
                        v1 = v1
                        v31 := GenS2{F0: (fi29 >> ((0) & 63)), F1: v1}
                        _ = v31
                    }
                    fmt.Println(fi30, fi29)
                    v32 := GenS1{F0: 627}
                    _ = fi30
                    _ = v32
                }
                _ = fi29
            }
        }
    }
    _ = len(v2)
    var v33 uint64 = 1
    v3 = (0.1 + v3)
    v1 = 65535
    v33 -= v33
    var v34 uint16 = genFn1(42, -1)
    v2 = append(v2, "hello")
    var v35 int = -545
    v36 := GenS1{F0: 2}
    v35 = v35
    v2 = append(v2, "a")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
