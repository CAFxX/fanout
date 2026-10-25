package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 uint64; F2 int }
func genFn0(p0 uint) uint32 {
    var v4 float32 = ((0.1 / 1.0) + (1.5 + (3.14159 * 1.5)))
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int = (-(fi5))
        v4 = v4
        v4 = (v4 * 0.0)
        _ = fi5
        _ = v6
    }
    if (false && ((p0 + p0) != (p0 - (p0 & (p0 % ((2) | 1)))))) {
        v7 := GenS1{F0: p0, F1: 0}
        _ = v7
    } else {
        switch "foo bar" {
        case "a":
            if false {
                for fi8 := 0; fi8 < 20; fi8++ {
                    fmt.Println(v4, v4)
                    _ = fi8
                }
            } else {
                for fi9 := 0; fi9 < 12; fi9++ {
                    v10 := GenS2{F0: 0}
                    var v11 int = 42
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                v4 = ((v4 / v4) * v4)
            }
        case "":
            fmt.Println(v4, v4)
        case "foo bar":
            var v12 bool = (814 < 2147483647)
            _ = v12
        default:
            if false {
                switch ("a" + "hello") {
                case "":
                    v4 = v4
                case "foo bar":
                    v4 -= 0.0
                }
                var v13 float64 = -2.25
                var v14 float32 = 0.1
                _ = v13
                _ = v14
            }
        }
        fmt.Println(v4, p0)
    }
    _ = p0
    _ = v4
    return 315
}

func main() {
    v1 := []int{}
    var v2 uint64 = 2
    v3 := []string{("hello" + "hello"), (("hello" + ("\n\t" + "")) + "\n\t"), "hello"}
    v15 := GenS1{F0: 42, F1: 42}
    v16 := GenS3{F0: 0}
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            v1 = append(v1, fi17)
            fmt.Println(fi18, fi18, v2)
            _ = fi18
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            if len(v1) > 0 {
                v1[0] = fi17
            }
            if (false || true) {
                if true {
                    v2 -= (v2 ^ 0)
                    _ = len(v1)
                    if (!(true && true)) {
                        if false {
                            v2 -= v2
                            if true {
                                if (220 > 1) {
                                    switch fi17 {
                                    case 0:
                                        switch (-(fi19)) {
                                        case 128:
                                            if len(v1) > 0 {
                                                v1[0] = 567
                                            }
                                        case 32767:
                                            v20 := GenS4{F0: 54, F1: (v2 | v2), F2: (-(fi17))}
                                            _ = v20
                                        case -2147483648:
                                            var v21 uint16 = 0
                                            _ = v21
                                        default:
                                            if false {
                                                v2 += (^(v2))
                                                v22 := GenS3{F0: -1}
                                                v23 := GenS4{F0: 114, F1: 1, F2: (fi17 * fi17)}
                                                _ = v22
                                                _ = v23
                                            } else {
                                                v24 := GenS3{F0: -683}
                                                _ = len(v3)
                                                _ = v24
                                            }
                                        }
                                    default:
                                        var v25 float64 = (1e10 / 3.14159)
                                        _ = v25
                                    }
                                    v2 -= 2
                                    fmt.Println(v2, fi19, fi19)
                                }
                                fmt.Println(v2)
                                v2 = 335
                            }
                        }
                        switch fi19 {
                        case -1:
                            var v26 float64 = -2.25
                            _ = v26
                        case 127:
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                        default:
                            if (0.0 == -2.25) {
                                var v27 float32 = 0.1
                                _ = v27
                            } else {
                                var v28 float32 = -2.25
                                v29 := GenS3{F0: -9223372036854775808}
                                _ = v28
                                _ = v29
                            }
                            var v30 float64 = 0.0
                            _ = v30
                        }
                        var v31 int8 = 42
                        v2 += 2
                        _ = v31
                    }
                }
            } else {
                switch "hello" {
                case "\n\t":
                    var v32 bool = (0 == 134)
                    _ = v32
                case "hello":
                    switch fi17 {
                    case 32768:
                        v2 = (v2 * (^(v2)))
                        fmt.Println(fi17, fi19)
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(fi17)
                        }
                    }
                case "foo bar":
                    if false {
                        v33 := GenS1{F0: 0, F1: genFn0(1)}
                        v34 := GenS1{F0: 42, F1: 647}
                        _ = v33
                        _ = v34
                    }
                default:
                    fmt.Println(v2, fi19)
                }
                v2 = v2
            }
            _ = fi19
        }
        _ = fi17
    }
    switch 42 {
    case 127:
        v35 := GenS2{F0: 9223372036854775807}
        _ = v35
    case 4294967295:
        if (true && true) {
            v2 -= (v2 % ((717) | 1))
        }
    default:
        for fi36 := 0; fi36 < 6; fi36++ {
            for fi37 := 0; fi37 < 20; fi37++ {
                v2 = v2
                v2 = v2
                _ = fi37
            }
            _ = fi36
        }
    }
    fmt.Println(v2)
    if true {
        v2 -= (v2 + v2)
        v2 += (v2 ^ 50)
    } else {
        switch "foo bar" {
        case "foo bar":
            var v38 uint16 = 648
            _ = v38
        case "a":
            var v39 int16 = -768
            _ = v39
        case "\n\t":
            if (!(!((!false) || (!false)))) {
                v2 = 0
                var v40 uint16 = 255
                if ((!((!false) && true)) && false) {
                    v41 := GenS1{F0: 159, F1: 707}
                    var v42 float32 = 1.5
                    fmt.Println(v42, v42)
                    _ = v41
                    _ = v42
                } else {
                    fmt.Println(v2, v2)
                    v43 := GenS4{F0: 42, F1: uint64(v40), F2: 870}
                    _ = v43
                }
                _ = v40
            }
        default:
            fmt.Println(v2)
        }
    }
    v2 = 0
    switch (("" + "\n\t") + "") {
    case "\n\t":
        var v44 uint32 = 795
        _ = v44
    }
    var v45 uint16 = 32767
    v45 = 255
    v2 = (v2 - v2)
    var v46 float32 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v45
    _ = v46
}
