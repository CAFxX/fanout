package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 string; F2 int64 }
type GenS7 struct { F0 int64; F1 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string; F1 int8; F2 int8 }
func genFn0(p0 int) int16 {
    v7 := []int{}
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: 2147483647}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 1000
}

func main() {
    var v1 int64 = 590
    v2 := []uint32{103, 65535, 2, 889}
    var v3 int = 32767
    var v4 uint32 = 1
    var v5 int32 = 316
    var v6 int64 = v1
    v4 -= v4
    for fi11 := 0; fi11 < 2; fi11++ {
        v5 -= (v5 / ((v5) | 1))
        v12 := GenS2{F0: (v6 / (((v1 << ((255) & 63))) | 1)), F1: 31}
        _ = fi11
        _ = v12
    }
    v13 := GenS3{F0: v3, F1: 1}
    v6 = v1
    var v14 uint64 = 1
    if true {
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := GenS4{F0: (^(v5)), F1: 1}
            v14 -= v14
            _ = fi15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 1; fi17++ {
            fmt.Println(v5)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (-277 > genFn0((fi17 * 128))) {
                var v18 int32 = (v5 - v5)
                _ = v18
            }
            _ = fi17
        }
        if (!(4294967295 <= 457)) {
            v19 := GenS5{F0: 100}
            _ = v19
        }
    }
    v14 = 4294967295
    v20 := GenS1{F0: v5}
    if false {
        v21 := GenS6{F0: "a", F1: ("hello" + "foo bar"), F2: v1}
        _ = v21
    }
    if (!(!true)) {
        if (!true) {
            v4 = (v4 % ((v4) | 1))
            v22 := GenS7{F0: v1, F1: 107}
            _ = v22
        } else {
            var v23 int8 = 127
            switch v3 {
            case 814:
                if (!(!(v6 < (^(v1))))) {
                    switch -740 {
                    case 141:
                        var v24 uint32 = v4
                        _ = v24
                    case 42:
                        fmt.Println(v3, v14, v4)
                    case 32768:
                        if false {
                            v4 = 2
                            var v25 float64 = (-2.25 * 0.1)
                            _ = v25
                        } else {
                            if len(v2) > 0 {
                                v2[0] = (v4 << ((4294967295) & 63))
                            }
                            if (!(!true)) {
                                switch "foo bar" {
                                case "hello":
                                    for fi26 := 0; fi26 < 17; fi26++ {
                                        switch "a" {
                                        case "foo bar":
                                            for fi27 := 0; fi27 < 7; fi27++ {
                                                v1 += v6
                                                v23 += (v23 >> ((52) & 63))
                                                _ = fi27
                                            }
                                        case "\n\t":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        fmt.Println(v5, v14)
                                        if (v14 <= v14) {
                                            if (true || true) {
                                                v14 = 4294967295
                                                v1 = 0
                                                if false {
                                                    fmt.Println(fi26, fi26)
                                                    var v28 int = (-(v3))
                                                    _ = v28
                                                } else {
                                                    v29 := GenS3{F0: 127, F1: 165}
                                                    _ = v29
                                                }
                                            }
                                        } else {
                                            var v30 float64 = 0.1
                                            _ = v30
                                        }
                                        _ = fi26
                                    }
                                case "a":
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v1 -= (^(v1))
                                        fmt.Println(v23)
                                    }
                                case "":
                                    fmt.Println(v3, v23, v6)
                                }
                            } else {
                                fmt.Println(v1, v3)
                            }
                        }
                    default:
                        var v31 int16 = 0
                        _ = v31
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v32 float32 = 1e10
                    _ = v32
                }
            case 128:
                if true {
                    if true {
                        switch (v3 + (v3 & (v3 ^ -1))) {
                        case -9223372036854775808:
                            if (("\n\t" > "foo bar") || (v3 == v3)) {
                                fmt.Println(v14)
                            }
                        }
                        if (("foo bar" > "foo bar") || false) {
                            v5 = 0
                            v1 += 42
                            var v33 uint8 = 2
                            _ = v33
                        } else {
                            fmt.Println(v5, v6)
                        }
                        var v34 uint8 = 67
                        _ = v34
                    } else {
                        v35 := GenS8{F0: v14}
                        _ = v35
                    }
                    v2 = append(v2, (^(v4)))
                }
            case 137:
                for fi36 := 0; fi36 < 19; fi36++ {
                    var v37 uint64 = (v14 * 996)
                    v14 = v14
                    if (32767 == 42) {
                        switch fi36 {
                        case -2147483648:
                            if true {
                                v38 := GenS9{F0: ((("\n\t" + (("hello" + "") + "a")) + ("\n\t" + ("hello" + "hello"))) + ("foo bar" + "")), F1: (v23 ^ -15), F2: -12}
                                v39 := GenS6{F0: ("" + ""), F1: ("\n\t" + "hello"), F2: v6}
                                _ = v38
                                _ = v39
                            } else {
                                var v40 uint64 = v14
                                _ = v40
                            }
                        case 0:
                            var v41 bool = false
                            _ = v41
                        }
                    } else {
                        for fi42 := 0; fi42 < 5; fi42++ {
                            v1 = v1
                            _ = fi42
                        }
                    }
                    _ = fi36
                    _ = v37
                }
            }
            var v43 int = v3
            _ = v23
            _ = v43
        }
        for fi44 := 0; fi44 < 10; fi44++ {
            _ = len(v2)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi44
        }
    } else {
        switch v3 {
        case 32768:
            v2 = append(v2, v4)
        case 597:
            var v45 bool = false
            _ = v45
        case 42:
            var v46 bool = false
            _ = v46
        }
        for fi47 := 0; fi47 < 18; fi47++ {
            v14 = (v14 / ((167) | 1))
            _ = fi47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
